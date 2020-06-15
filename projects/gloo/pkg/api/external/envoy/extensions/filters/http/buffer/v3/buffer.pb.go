// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/buffer/v3/buffer.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes      *types.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9fce05f80b94cf, []int{0}
}
func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buffer.Unmarshal(m, b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return xxx_messageInfo_Buffer.Size(m)
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *types.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9fce05f80b94cf, []int{1}
}
func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferPerRoute.Unmarshal(m, b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
}
func (m *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(m, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return xxx_messageInfo_BufferPerRoute.Size(m)
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
	Equal(interface{}) bool
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof" json:"disabled,omitempty"`
}
type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof" json:"buffer,omitempty"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}
func (*BufferPerRoute_Buffer) isBufferPerRoute_Override()   {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.extensions.filters.http.buffer.v3.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.extensions.filters.http.buffer.v3.BufferPerRoute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/buffer/v3/buffer.proto", fileDescriptor_7e9fce05f80b94cf)
}

var fileDescriptor_7e9fce05f80b94cf = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6e, 0xd4, 0x30,
	0x14, 0x1e, 0x67, 0xa6, 0xa3, 0xe0, 0x4a, 0xd0, 0x06, 0x24, 0xaa, 0x0a, 0x55, 0x55, 0x37, 0x45,
	0x48, 0xd8, 0x52, 0x47, 0x5c, 0xc0, 0xab, 0xc2, 0xaa, 0x8a, 0x54, 0x16, 0x48, 0x50, 0x39, 0x9d,
	0x37, 0xae, 0x8b, 0x93, 0x67, 0x6c, 0x27, 0x64, 0xae, 0xc0, 0x35, 0xd8, 0xb0, 0x62, 0x5d, 0xb1,
	0xea, 0x59, 0xd8, 0xf5, 0x00, 0xec, 0x51, 0xec, 0x94, 0x6e, 0x47, 0xec, 0x3e, 0x7f, 0xf6, 0xf7,
	0x93, 0x97, 0x47, 0x95, 0xd2, 0xe1, 0xaa, 0xad, 0xd8, 0x25, 0xd6, 0xdc, 0xa3, 0xc1, 0xd7, 0x1a,
	0xb9, 0x32, 0x88, 0xdc, 0x3a, 0xbc, 0x86, 0xcb, 0xe0, 0xd3, 0x49, 0x5a, 0xcd, 0xa1, 0x0f, 0xe0,
	0x1a, 0x69, 0x38, 0x34, 0x1d, 0xae, 0xe3, 0xb1, 0xf1, 0x1a, 0x1b, 0xcf, 0x57, 0xda, 0x04, 0x70,
	0x9e, 0x5f, 0x85, 0x60, 0x79, 0xd5, 0xae, 0x56, 0xe0, 0x78, 0xb7, 0x18, 0x11, 0xb3, 0x0e, 0x03,
	0x16, 0xc7, 0x51, 0xc5, 0x1e, 0x54, 0x6c, 0x54, 0xb1, 0x41, 0xc5, 0xc6, 0xb7, 0xdd, 0x62, 0xff,
	0x40, 0x21, 0x2a, 0x03, 0x3c, 0xca, 0xaa, 0x76, 0xc5, 0xbf, 0x3a, 0x69, 0xed, 0xf0, 0x30, 0x32,
	0xfb, 0xcf, 0x3b, 0x69, 0xf4, 0x52, 0x06, 0xe0, 0xf7, 0x60, 0xbc, 0x78, 0xa6, 0x50, 0x61, 0x84,
	0x7c, 0x40, 0x23, 0x5b, 0x40, 0x1f, 0x12, 0x09, 0x7d, 0x48, 0xdc, 0x51, 0x4d, 0xe7, 0x22, 0xe6,
	0x15, 0x1f, 0xe9, 0x6e, 0x2d, 0xfb, 0x0b, 0x07, 0x5f, 0x5a, 0xf0, 0xe1, 0xa2, 0x5a, 0x07, 0xf0,
	0x7b, 0xe4, 0x90, 0xbc, 0xdc, 0x3e, 0x79, 0xc1, 0x52, 0x11, 0x76, 0x5f, 0x84, 0x9d, 0xbf, 0x6d,
	0xc2, 0xe2, 0xe4, 0xbd, 0x34, 0x2d, 0x88, 0xa7, 0xbf, 0xee, 0x6e, 0xa7, 0xb3, 0x57, 0xd9, 0xe1,
	0x64, 0x00, 0x5b, 0xdf, 0x48, 0xb6, 0x43, 0xca, 0x27, 0xb5, 0xec, 0xcb, 0x64, 0x25, 0x06, 0xa7,
	0x77, 0xb3, 0x3c, 0xdb, 0x99, 0x1e, 0x7d, 0x27, 0xf4, 0x71, 0xca, 0x3b, 0x03, 0x57, 0x62, 0x1b,
	0xa0, 0x38, 0xa6, 0xf9, 0x52, 0x7b, 0x59, 0x19, 0x58, 0xc6, 0xb8, 0x5c, 0x3c, 0x8a, 0x86, 0xd7,
	0x59, 0x4e, 0x4e, 0x27, 0xe5, 0xbf, 0xcb, 0xe2, 0x9c, 0xce, 0xd3, 0x68, 0xf6, 0xb2, 0xd8, 0x8a,
	0xb3, 0x0d, 0xe7, 0xc8, 0x52, 0xa2, 0xa0, 0x0f, 0xfd, 0x4e, 0x27, 0xe5, 0x68, 0x26, 0x76, 0x69,
	0x8e, 0x1d, 0x38, 0xa7, 0x97, 0x50, 0x6c, 0xdd, 0xdc, 0xdd, 0x4e, 0x89, 0xf8, 0x49, 0x6e, 0xfe,
	0xcc, 0xc8, 0x8f, 0xdf, 0x07, 0x84, 0xbe, 0xd1, 0x98, 0x62, 0xac, 0xc3, 0x7e, 0xbd, 0x69, 0xa2,
	0xd8, 0x1e, 0x3f, 0x72, 0x98, 0xd7, 0x19, 0xf9, 0xf0, 0x69, 0xb3, 0xd5, 0xb2, 0x9f, 0xd5, 0x7f,
	0xad, 0x57, 0x35, 0x8f, 0x3f, 0x66, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x1d, 0x7e, 0x91,
	0xc3, 0x02, 0x00, 0x00,
}

func (this *Buffer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Buffer)
	if !ok {
		that2, ok := that.(Buffer)
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
	if !this.MaxRequestBytes.Equal(that1.MaxRequestBytes) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BufferPerRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BufferPerRoute)
	if !ok {
		that2, ok := that.(BufferPerRoute)
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
	if that1.Override == nil {
		if this.Override != nil {
			return false
		}
	} else if this.Override == nil {
		return false
	} else if !this.Override.Equal(that1.Override) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *BufferPerRoute_Disabled) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BufferPerRoute_Disabled)
	if !ok {
		that2, ok := that.(BufferPerRoute_Disabled)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	return true
}
func (this *BufferPerRoute_Buffer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BufferPerRoute_Buffer)
	if !ok {
		that2, ok := that.(BufferPerRoute_Buffer)
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
	if !this.Buffer.Equal(that1.Buffer) {
		return false
	}
	return true
}
