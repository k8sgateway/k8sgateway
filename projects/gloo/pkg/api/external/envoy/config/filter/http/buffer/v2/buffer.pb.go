// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/filter/http/buffer/v2/buffer.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
	return fileDescriptor_b69a0f4ebde06d40, []int{0}
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
	return fileDescriptor_b69a0f4ebde06d40, []int{1}
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
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptor_b69a0f4ebde06d40)
}

var fileDescriptor_b69a0f4ebde06d40 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6e, 0xd4, 0x30,
	0x14, 0x1e, 0x67, 0xa6, 0xa3, 0xe0, 0x4a, 0xd0, 0x06, 0x24, 0xaa, 0x0a, 0x55, 0xd5, 0x6c, 0x40,
	0x95, 0xb0, 0xd1, 0xf4, 0x06, 0x5e, 0x15, 0x56, 0xa3, 0x48, 0x20, 0x81, 0x84, 0x2a, 0x67, 0xe6,
	0xc5, 0x75, 0x71, 0xf2, 0x8c, 0xe3, 0x84, 0xcc, 0x15, 0x38, 0x04, 0x2b, 0x16, 0x1c, 0xa1, 0x62,
	0xd5, 0xb3, 0xb0, 0xeb, 0x01, 0xd8, 0xa3, 0xd8, 0x29, 0x2c, 0x69, 0x77, 0x9f, 0xdf, 0xf3, 0xf7,
	0xe3, 0xe7, 0x47, 0x0b, 0xa5, 0xfd, 0x45, 0x5b, 0xb0, 0x35, 0x56, 0xbc, 0x41, 0x83, 0x2f, 0x35,
	0x72, 0x65, 0x10, 0xb9, 0x75, 0x78, 0x09, 0x6b, 0xdf, 0xc4, 0x93, 0xb4, 0x9a, 0x43, 0xef, 0xc1,
	0xd5, 0xd2, 0x70, 0xa8, 0x3b, 0xdc, 0xf2, 0x35, 0xd6, 0xa5, 0x56, 0xbc, 0xd4, 0xc6, 0x83, 0xe3,
	0x17, 0xde, 0x5b, 0x5e, 0xb4, 0x65, 0x09, 0x8e, 0x77, 0xcb, 0x11, 0x31, 0xeb, 0xd0, 0x63, 0xb6,
	0x08, 0x04, 0x16, 0x09, 0x2c, 0x12, 0xd8, 0x40, 0x60, 0xe3, 0xb5, 0x6e, 0x79, 0x78, 0xa4, 0x10,
	0x95, 0x01, 0x1e, 0x18, 0x45, 0x5b, 0xf2, 0x2f, 0x4e, 0x5a, 0x0b, 0xae, 0x89, 0x1a, 0x87, 0x4f,
	0x3b, 0x69, 0xf4, 0x46, 0x7a, 0xe0, 0xb7, 0x60, 0x6c, 0x3c, 0x51, 0xa8, 0x30, 0x40, 0x3e, 0xa0,
	0xb1, 0x9a, 0x41, 0xef, 0x63, 0x11, 0x7a, 0x1f, 0x6b, 0x8b, 0x8a, 0xce, 0x45, 0xf0, 0xcb, 0x3e,
	0xd2, 0xfd, 0x4a, 0xf6, 0xe7, 0x0e, 0x3e, 0xb7, 0xd0, 0xf8, 0xf3, 0x62, 0xeb, 0xa1, 0x39, 0x20,
	0xc7, 0xe4, 0xc5, 0xee, 0xf2, 0x19, 0x8b, 0x41, 0xd8, 0x6d, 0x10, 0xf6, 0xf6, 0x75, 0xed, 0x4f,
	0x97, 0xef, 0xa4, 0x69, 0x41, 0x3c, 0xfe, 0x79, 0x73, 0x3d, 0x9d, 0x9d, 0x24, 0xc7, 0x93, 0x01,
	0xec, 0x7c, 0x25, 0xc9, 0x1e, 0xc9, 0x1f, 0x55, 0xb2, 0xcf, 0xa3, 0x94, 0x18, 0x94, 0xde, 0xcc,
	0xd2, 0x64, 0x6f, 0xba, 0xf8, 0x46, 0xe8, 0xc3, 0xe8, 0xb7, 0x02, 0x97, 0x63, 0xeb, 0x21, 0x7b,
	0x4e, 0xd3, 0x8d, 0x6e, 0x64, 0x61, 0x60, 0x13, 0xec, 0x52, 0xf1, 0x20, 0x08, 0x5e, 0x26, 0x29,
	0x39, 0x9b, 0xe4, 0x7f, 0x9b, 0xd9, 0x8a, 0xce, 0xe3, 0x68, 0x0e, 0x92, 0x90, 0xea, 0x84, 0xfd,
	0x7f, 0x84, 0x2c, 0x9a, 0x09, 0xfa, 0x2f, 0xda, 0xd9, 0x24, 0x1f, 0x75, 0xc4, 0x3e, 0x4d, 0xb1,
	0x03, 0xe7, 0xf4, 0x06, 0xb2, 0x9d, 0xab, 0x9b, 0xeb, 0x29, 0x11, 0xdf, 0xc9, 0xd5, 0xef, 0x19,
	0xf9, 0xf1, 0xeb, 0x88, 0xd0, 0x57, 0x1a, 0xa3, 0x83, 0x75, 0xd8, 0x6f, 0xef, 0x60, 0x26, 0x76,
	0xc7, 0xa7, 0x0d, 0x53, 0x5a, 0x91, 0x0f, 0xef, 0xef, 0xb6, 0x46, 0xf6, 0x93, 0xba, 0xef, 0x2a,
	0x15, 0xf3, 0xf0, 0x13, 0xa7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe3, 0x78, 0x88, 0xaa,
	0x02, 0x00, 0x00,
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
