// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc_web/grpc_web.proto

package grpc_web

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GrpcWeb support is enabled be default. Use this extension to disable it.
type GrpcWeb struct {
	// Disable grpc web support.
	Disable              bool     `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcWeb) Reset()         { *m = GrpcWeb{} }
func (m *GrpcWeb) String() string { return proto.CompactTextString(m) }
func (*GrpcWeb) ProtoMessage()    {}
func (*GrpcWeb) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73ce11c1cb701fb, []int{0}
}
func (m *GrpcWeb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcWeb.Unmarshal(m, b)
}
func (m *GrpcWeb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcWeb.Marshal(b, m, deterministic)
}
func (m *GrpcWeb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcWeb.Merge(m, src)
}
func (m *GrpcWeb) XXX_Size() int {
	return xxx_messageInfo_GrpcWeb.Size(m)
}
func (m *GrpcWeb) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcWeb.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcWeb proto.InternalMessageInfo

func (m *GrpcWeb) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcWeb)(nil), "grpc_web.plugins.gloo.solo.io.GrpcWeb")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/grpc_web/grpc_web.proto", fileDescriptor_b73ce11c1cb701fb)
}

var fileDescriptor_b73ce11c1cb701fb = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0xd3, 0x8b, 0x0a,
	0x92, 0xe3, 0xcb, 0x53, 0x93, 0xe0, 0x0c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x04,
	0x1f, 0xa2, 0x52, 0x0f, 0xa4, 0x5b, 0x0f, 0x64, 0xb0, 0x5e, 0x66, 0xbe, 0x94, 0x48, 0x7a, 0x7e,
	0x7a, 0x3e, 0x58, 0xa5, 0x3e, 0x88, 0x05, 0xd1, 0x24, 0x15, 0x42, 0x96, 0x13, 0x4a, 0x8a, 0x12,
	0xf3, 0x8a, 0xd3, 0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x0b, 0x12, 0x8b, 0x12,
	0x73, 0x53, 0x4b, 0x52, 0x8b, 0x8a, 0x21, 0xa6, 0x2a, 0x29, 0x73, 0xb1, 0xbb, 0x17, 0x15, 0x24,
	0x87, 0xa7, 0x26, 0x09, 0x49, 0x70, 0xb1, 0xa7, 0x64, 0x16, 0x27, 0x26, 0xe5, 0xa4, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xb8, 0x4e, 0xee, 0x2b, 0x1e, 0xc9, 0x31, 0x46, 0x39, 0x12,
	0xe7, 0x80, 0x82, 0xec, 0x74, 0x5c, 0xe1, 0x90, 0xc4, 0x06, 0xb6, 0xd4, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0x85, 0x1e, 0x1f, 0x4f, 0x01, 0x00, 0x00,
}

func (this *GrpcWeb) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcWeb)
	if !ok {
		that2, ok := that.(GrpcWeb)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Disable != that1.Disable {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
