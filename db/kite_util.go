package db

import (
	"encoding/binary"
	"errors"
	"github.com/xuyu/goredis"
	"log"
	"time"
)

type Type *KiteDBPage

// 一个固定长度的环状队列
type KiteRingQueue struct {
	front int // 队头指针
	rear  int // 队尾指针
	size  int // 队列最大长度
	data  []Type
}

func (self *KiteRingQueue) Len() int {
	return (self.front - self.rear + self.size) % self.size
}

func NewKiteRingQueue(size int) *KiteRingQueue {
	return &KiteRingQueue{
		size:  size,
		front: 0,
		rear:  0,
		data:  make([]Type, size+1),
	}
}

func (self *KiteRingQueue) Enqueue(e Type) error {
	//牺牲一个存储单元判断队列为满
	if (self.rear+1)%self.size == self.front {
		return errors.New("queue is full")
	}
	self.data[self.rear] = e
	self.rear = (self.rear + 1) % self.size
	return nil
}

func (self *KiteRingQueue) Dequeue() (Type, error) {
	if self.rear == self.front {
		return nil, errors.New("queue is empty")
	}
	data := self.data[self.front]
	self.front = (self.front + 1) % self.size
	return data, nil
}

type KiteBitset struct {
	numBits int
	bits    []byte
	version int
}

func NewKiteBitset() *KiteBitset {
	redis, err := goredis.Dial(&goredis.DialConfig{Address: ":6379"})
	if err != nil {
		log.Fatal("page bits redis down")
	}
	r, _ := redis.ExecuteCommand("GET", "page_bits")
	bits, _ := r.BytesValue()
	r, _ = redis.ExecuteCommand("GET", "page_bits_num")
	num_bytes, _ := r.BytesValue()
	var b *KiteBitset
	if bits != nil && len(num_bytes) > 0 {
		num := int(binary.BigEndian.Uint32(num_bytes))
		b = &KiteBitset{numBits: num, bits: bits, version: 0}
	} else {
		b = &KiteBitset{numBits: 0, bits: make([]byte, 0), version: 0}
	}

	go b.Save()
	return b
}

func (b *KiteBitset) Save() {
	saveVersion := b.version
	for {
		time.Sleep(time.Second * 3)
		if b.version > saveVersion {
			clone := &KiteBitset{numBits: b.numBits, bits: b.bits[:]}
			redis, err := goredis.Dial(&goredis.DialConfig{Address: ":6379"})
			if err != nil {
				log.Fatal("page bits redis down")
			}
			redis.ExecuteCommand("SET", "page_bits", clone.bits)
			bs := make([]byte, 4)
			binary.BigEndian.PutUint32(bs, uint32(clone.numBits))
			redis.ExecuteCommand("SET", "page_bits_num", bs)
			saveVersion = b.version
		}
	}
}

func (b *KiteBitset) AppendBytes(data []byte) {
	b.version = b.version + 1
	for _, d := range data {
		b.AppendByte(d, 8)
	}
}

func (b *KiteBitset) AppendByte(value byte, numBits int) {
	b.ensureCapacity(numBits)

	if numBits > 8 {
		log.Fatal("numBits %d out of range 0-8", numBits)
	}

	for i := numBits - 1; i >= 0; i-- {
		if value&(1<<uint(i)) != 0 {
			b.bits[b.numBits/8] |= 0x80 >> uint(b.numBits%8)
		}

		b.numBits++
	}
}

func (b *KiteBitset) ensureCapacity(numBits int) {
	numBits += b.numBits

	newNumBytes := numBits / 8
	if numBits%8 != 0 {
		newNumBytes++
	}

	if len(b.bits) >= newNumBytes {
		return
	}

	b.bits = append(b.bits, make([]byte, newNumBytes+2*len(b.bits))...)
}

func (b *KiteBitset) Len() int {
	return b.numBits
}

func (b *KiteBitset) At(index int) bool {
	b.ensureCapacity(index)

	return (b.bits[index/8] & (0x80 >> byte(index%8))) != 0
}

func (b *KiteBitset) Set(index int, status bool) {
	b.version = b.version + 1
	b.ensureCapacity(index)

	if status == true {
		b.bits[index/8] |= (0x80 >> byte(index%8))
	} else {
		b.bits[index/8] &= (0x80 >> byte(index%8)) ^ 0xf
	}
}
