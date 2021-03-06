// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/auth/auth_msg.proto

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthMsgTest struct {
	I                    int32    `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	S                    string   `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthMsgTest) Reset()         { *m = AuthMsgTest{} }
func (m *AuthMsgTest) String() string { return proto.CompactTextString(m) }
func (*AuthMsgTest) ProtoMessage()    {}
func (*AuthMsgTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_msg_607216fccd27f44a, []int{0}
}
func (m *AuthMsgTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthMsgTest.Unmarshal(m, b)
}
func (m *AuthMsgTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthMsgTest.Marshal(b, m, deterministic)
}
func (dst *AuthMsgTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthMsgTest.Merge(dst, src)
}
func (m *AuthMsgTest) XXX_Size() int {
	return xxx_messageInfo_AuthMsgTest.Size(m)
}
func (m *AuthMsgTest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthMsgTest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthMsgTest proto.InternalMessageInfo

func (m *AuthMsgTest) GetI() int32 {
	if m != nil {
		return m.I
	}
	return 0
}

func (m *AuthMsgTest) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type AuthMsgTestAck struct {
	IntList              []int32  `protobuf:"varint,1,rep,packed,name=intList,proto3" json:"intList,omitempty"`
	StrList              []string `protobuf:"bytes,2,rep,name=strList,proto3" json:"strList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthMsgTestAck) Reset()         { *m = AuthMsgTestAck{} }
func (m *AuthMsgTestAck) String() string { return proto.CompactTextString(m) }
func (*AuthMsgTestAck) ProtoMessage()    {}
func (*AuthMsgTestAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_msg_607216fccd27f44a, []int{1}
}
func (m *AuthMsgTestAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthMsgTestAck.Unmarshal(m, b)
}
func (m *AuthMsgTestAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthMsgTestAck.Marshal(b, m, deterministic)
}
func (dst *AuthMsgTestAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthMsgTestAck.Merge(dst, src)
}
func (m *AuthMsgTestAck) XXX_Size() int {
	return xxx_messageInfo_AuthMsgTestAck.Size(m)
}
func (m *AuthMsgTestAck) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthMsgTestAck.DiscardUnknown(m)
}

var xxx_messageInfo_AuthMsgTestAck proto.InternalMessageInfo

func (m *AuthMsgTestAck) GetIntList() []int32 {
	if m != nil {
		return m.IntList
	}
	return nil
}

func (m *AuthMsgTestAck) GetStrList() []string {
	if m != nil {
		return m.StrList
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthMsgTest)(nil), "auth.AuthMsgTest")
	proto.RegisterType((*AuthMsgTestAck)(nil), "auth.AuthMsgTestAck")
}

func init() { proto.RegisterFile("pb/auth/auth_msg.proto", fileDescriptor_auth_msg_607216fccd27f44a) }

var fileDescriptor_auth_msg_607216fccd27f44a = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x4f,
	0x2c, 0x2d, 0xc9, 0x00, 0x13, 0xf1, 0xb9, 0xc5, 0xe9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x2c, 0x20, 0xbe, 0x92, 0x26, 0x17, 0xb7, 0x63, 0x69, 0x49, 0x86, 0x6f, 0x71, 0x7a, 0x48, 0x6a,
	0x71, 0x89, 0x10, 0x0f, 0x17, 0x63, 0xa6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x63, 0x26,
	0x88, 0x57, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x58, 0xac, 0xe4, 0xc2, 0xc5, 0x87,
	0xa4, 0xd4, 0x31, 0x39, 0x5b, 0x48, 0x82, 0x8b, 0x3d, 0x33, 0xaf, 0xc4, 0x27, 0xb3, 0xb8, 0x44,
	0x82, 0x51, 0x81, 0x59, 0x83, 0x35, 0x08, 0xc6, 0x05, 0xc9, 0x14, 0x97, 0x14, 0x81, 0x65, 0x98,
	0x14, 0x98, 0x35, 0x38, 0x83, 0x60, 0xdc, 0x24, 0x36, 0xb0, 0xed, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0xa0, 0x8e, 0x5c, 0x97, 0x00, 0x00, 0x00,
}
