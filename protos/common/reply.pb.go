// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/reply.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SimpleReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleReply) Reset()         { *m = SimpleReply{} }
func (m *SimpleReply) String() string { return proto.CompactTextString(m) }
func (*SimpleReply) ProtoMessage()    {}
func (*SimpleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b4de264dfe8be32, []int{0}
}

func (m *SimpleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleReply.Unmarshal(m, b)
}
func (m *SimpleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleReply.Marshal(b, m, deterministic)
}
func (m *SimpleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleReply.Merge(m, src)
}
func (m *SimpleReply) XXX_Size() int {
	return xxx_messageInfo_SimpleReply.Size(m)
}
func (m *SimpleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleReply.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleReply proto.InternalMessageInfo

func (m *SimpleReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SimpleReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SimpleReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SimpleReply)(nil), "common.SimpleReply")
}

func init() { proto.RegisterFile("common/reply.proto", fileDescriptor_2b4de264dfe8be32) }

var fileDescriptor_2b4de264dfe8be32 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2f, 0x4a, 0x2d, 0xc8, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x83, 0x88, 0x29, 0x85, 0x72, 0x71, 0x07, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0x06, 0x81, 0x24, 0x85,
	0x24, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x82, 0x60, 0x5c, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0x30, 0x1b, 0xa4, 0x3a, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x59, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xc6, 0x75, 0xd2, 0x8d, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd, 0x49, 0xcb, 0x2f, 0x4a, 0x07, 0x51, 0x05, 0x39, 0x89, 0x25,
	0x69, 0xf9, 0x45, 0xb9, 0xfa, 0x60, 0x57, 0x14, 0xeb, 0x43, 0x5c, 0x91, 0xc4, 0x06, 0xe6, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x86, 0x4c, 0x83, 0xaa, 0x00, 0x00, 0x00,
}