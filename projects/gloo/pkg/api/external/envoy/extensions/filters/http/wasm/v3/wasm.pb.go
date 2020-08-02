// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/wasm/v3/wasm.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/wasm/v3"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type Wasm struct {
	// General Plugin configuration.
	Config               *v3.PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Wasm) Reset()         { *m = Wasm{} }
func (m *Wasm) String() string { return proto.CompactTextString(m) }
func (*Wasm) ProtoMessage()    {}
func (*Wasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4a50b3626ddabde, []int{0}
}
func (m *Wasm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wasm.Unmarshal(m, b)
}
func (m *Wasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wasm.Marshal(b, m, deterministic)
}
func (m *Wasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wasm.Merge(m, src)
}
func (m *Wasm) XXX_Size() int {
	return xxx_messageInfo_Wasm.Size(m)
}
func (m *Wasm) XXX_DiscardUnknown() {
	xxx_messageInfo_Wasm.DiscardUnknown(m)
}

var xxx_messageInfo_Wasm proto.InternalMessageInfo

func (m *Wasm) GetConfig() *v3.PluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Wasm)(nil), "envoy.extensions.filters.http.wasm.v3.Wasm")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/wasm/v3/wasm.proto", fileDescriptor_a4a50b3626ddabde)
}

var fileDescriptor_a4a50b3626ddabde = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x09, 0x1c, 0x07, 0xae, 0xdd, 0x21, 0x28, 0x57, 0x88, 0x28, 0x8a, 0x8d, 0x19, 0x70,
	0x7b, 0x8b, 0x13, 0xac, 0x17, 0x1b, 0x41, 0x6c, 0x72, 0x6b, 0x2e, 0x37, 0x9a, 0xcd, 0x84, 0x64,
	0x6e, 0xdd, 0x7b, 0xa3, 0x7b, 0x04, 0x9f, 0xc7, 0x77, 0xb0, 0x97, 0xcd, 0x46, 0x2c, 0xb6, 0x39,
	0xac, 0xf2, 0xcf, 0x9f, 0xf9, 0x86, 0x7f, 0x98, 0xa2, 0x36, 0xc8, 0xeb, 0xcd, 0x52, 0xd6, 0xd4,
	0x40, 0x24, 0x4b, 0x37, 0x48, 0x60, 0x2c, 0x11, 0xf8, 0x40, 0x6f, 0xba, 0xe6, 0x38, 0x54, 0xca,
	0x23, 0xe8, 0x8e, 0x75, 0x70, 0xca, 0x82, 0x76, 0x2d, 0x6d, 0x53, 0xe9, 0x22, 0x92, 0x8b, 0xb0,
	0x42, 0xcb, 0x3a, 0x44, 0x58, 0x33, 0x7b, 0xf8, 0x50, 0xb1, 0x81, 0xb6, 0x4c, 0xaf, 0xf4, 0x81,
	0x98, 0x66, 0x97, 0x89, 0x90, 0x7f, 0x84, 0xcc, 0x84, 0xec, 0x09, 0x99, 0x3a, 0xdb, 0x72, 0x7e,
	0x31, 0x1a, 0x3c, 0x9e, 0x35, 0x3f, 0x6e, 0x95, 0xc5, 0x57, 0xc5, 0x1a, 0x7e, 0x45, 0xfe, 0x38,
	0x32, 0x64, 0x28, 0x49, 0xe8, 0x55, 0x76, 0x67, 0xba, 0xe3, 0xc1, 0xd4, 0x1d, 0x0f, 0xde, 0xf9,
	0x43, 0x31, 0x79, 0x52, 0xb1, 0x99, 0xdd, 0x15, 0xd3, 0x9a, 0xdc, 0x0a, 0xcd, 0x89, 0x38, 0x13,
	0xd7, 0x87, 0xb7, 0x57, 0x72, 0x94, 0x33, 0x47, 0x93, 0x95, 0xdd, 0x18, 0x74, 0xf7, 0xa9, 0xfb,
	0x31, 0x53, 0x8b, 0x9d, 0xf8, 0xfc, 0x9e, 0x88, 0xdd, 0xd7, 0xa9, 0x28, 0x4a, 0xa4, 0x01, 0xf6,
	0x81, 0xba, 0xad, 0xdc, 0x6b, 0xdf, 0xc5, 0x41, 0x9f, 0xa2, 0xea, 0x23, 0x55, 0xe2, 0xf9, 0x65,
	0xbf, 0x43, 0xf8, 0x77, 0xf3, 0x8f, 0x63, 0x2c, 0xa7, 0x69, 0xf3, 0xf2, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x6e, 0xce, 0xf6, 0x99, 0xef, 0x01, 0x00, 0x00,
}

func (this *Wasm) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Wasm)
	if !ok {
		that2, ok := that.(Wasm)
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
	if !this.Config.Equal(that1.Config) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
