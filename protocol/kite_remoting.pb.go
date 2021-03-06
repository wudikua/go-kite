// Code generated by protoc-gen-go.
// source: kite_remoting.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type HeartBeatPacket_SerializeType int32

const (
	HeartBeatPacket_JSON    HeartBeatPacket_SerializeType = 0
	HeartBeatPacket_MSGPACK HeartBeatPacket_SerializeType = 1
	HeartBeatPacket_PROTOC  HeartBeatPacket_SerializeType = 2
)

var HeartBeatPacket_SerializeType_name = map[int32]string{
	0: "JSON",
	1: "MSGPACK",
	2: "PROTOC",
}
var HeartBeatPacket_SerializeType_value = map[string]int32{
	"JSON":    0,
	"MSGPACK": 1,
	"PROTOC":  2,
}

func (x HeartBeatPacket_SerializeType) Enum() *HeartBeatPacket_SerializeType {
	p := new(HeartBeatPacket_SerializeType)
	*p = x
	return p
}
func (x HeartBeatPacket_SerializeType) String() string {
	return proto.EnumName(HeartBeatPacket_SerializeType_name, int32(x))
}
func (x *HeartBeatPacket_SerializeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HeartBeatPacket_SerializeType_value, data, "HeartBeatPacket_SerializeType")
	if err != nil {
		return err
	}
	*x = HeartBeatPacket_SerializeType(value)
	return nil
}

// 心跳请求包
type HeartBeatPacket struct {
	Id               *int64                         `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type             *HeartBeatPacket_SerializeType `protobuf:"varint,2,req,name=type,enum=protocol.HeartBeatPacket_SerializeType,def=0" json:"type,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *HeartBeatPacket) Reset()         { *m = HeartBeatPacket{} }
func (m *HeartBeatPacket) String() string { return proto.CompactTextString(m) }
func (*HeartBeatPacket) ProtoMessage()    {}

const Default_HeartBeatPacket_Type HeartBeatPacket_SerializeType = HeartBeatPacket_JSON

func (m *HeartBeatPacket) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HeartBeatPacket) GetType() HeartBeatPacket_SerializeType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_HeartBeatPacket_Type
}

// 连接的Meta数据包
type ConnectioMetaPacket struct {
	GroupId          *string `protobuf:"bytes,1,req,name=groupId" json:"groupId,omitempty"`
	SecretKey        *string `protobuf:"bytes,2,opt,name=secretKey" json:"secretKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConnectioMetaPacket) Reset()         { *m = ConnectioMetaPacket{} }
func (m *ConnectioMetaPacket) String() string { return proto.CompactTextString(m) }
func (*ConnectioMetaPacket) ProtoMessage()    {}

func (m *ConnectioMetaPacket) GetGroupId() string {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return ""
}

func (m *ConnectioMetaPacket) GetSecretKey() string {
	if m != nil && m.SecretKey != nil {
		return *m.SecretKey
	}
	return ""
}

// 心跳响应包
type HeartBeatACKPacket struct {
	AckId            *int64 `protobuf:"varint,1,req,name=ackId" json:"ackId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *HeartBeatACKPacket) Reset()         { *m = HeartBeatACKPacket{} }
func (m *HeartBeatACKPacket) String() string { return proto.CompactTextString(m) }
func (*HeartBeatACKPacket) ProtoMessage()    {}

func (m *HeartBeatACKPacket) GetAckId() int64 {
	if m != nil && m.AckId != nil {
		return *m.AckId
	}
	return 0
}

// 事务确认数据包
type TranscationACKPacket struct {
	AckId            *int64  `protobuf:"varint,1,req,name=ackId" json:"ackId,omitempty"`
	MessageId        *string `protobuf:"bytes,2,req,name=messageId" json:"messageId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TranscationACKPacket) Reset()         { *m = TranscationACKPacket{} }
func (m *TranscationACKPacket) String() string { return proto.CompactTextString(m) }
func (*TranscationACKPacket) ProtoMessage()    {}

func (m *TranscationACKPacket) GetAckId() int64 {
	if m != nil && m.AckId != nil {
		return *m.AckId
	}
	return 0
}

func (m *TranscationACKPacket) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

type Header struct {
	MessageId        *string `protobuf:"bytes,1,req,name=messageId" json:"messageId,omitempty"`
	Topic            *string `protobuf:"bytes,2,req,name=topic" json:"topic,omitempty"`
	MessageType      *string `protobuf:"bytes,3,req,name=messageType" json:"messageType,omitempty"`
	ExpiredTime      *int64  `protobuf:"varint,4,req,name=expiredTime" json:"expiredTime,omitempty"`
	GroupId          *string `protobuf:"bytes,5,req,name=groupId" json:"groupId,omitempty"`
	Commited         *bool   `protobuf:"varint,6,req,name=commited" json:"commited,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}

func (m *Header) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Header) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *Header) GetMessageType() string {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return ""
}

func (m *Header) GetExpiredTime() int64 {
	if m != nil && m.ExpiredTime != nil {
		return *m.ExpiredTime
	}
	return 0
}

func (m *Header) GetGroupId() string {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return ""
}

func (m *Header) GetCommited() bool {
	if m != nil && m.Commited != nil {
		return *m.Commited
	}
	return false
}

// byte类消息
type BytesMessage struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             []byte  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BytesMessage) Reset()         { *m = BytesMessage{} }
func (m *BytesMessage) String() string { return proto.CompactTextString(m) }
func (*BytesMessage) ProtoMessage()    {}

func (m *BytesMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BytesMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// string类型的message
type StringMessage struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             *string `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}

func (m *StringMessage) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *StringMessage) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.HeartBeatPacket_SerializeType", HeartBeatPacket_SerializeType_name, HeartBeatPacket_SerializeType_value)
}
