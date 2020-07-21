// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/backoff.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
<<<<<<< HEAD
	_ "github.com/solo-io/protoc-gen-ext/extproto"
=======
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
>>>>>>> master
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

// Configuration defining a jittered exponential back off strategy.
type BackoffStrategy struct {
	// The base interval to be used for the next back off computation. It should
	// be greater than zero and less than or equal to :ref:`max_interval
	// <envoy_api_field_config.core.v3.BackoffStrategy.max_interval>`.
	BaseInterval *types.Duration `protobuf:"bytes,1,opt,name=base_interval,json=baseInterval,proto3" json:"base_interval,omitempty"`
	// Specifies the maximum interval between retries. This parameter is optional,
	// but must be greater than or equal to the :ref:`base_interval
	// <envoy_api_field_config.core.v3.BackoffStrategy.base_interval>` if set. The default
	// is 10 times the :ref:`base_interval
	// <envoy_api_field_config.core.v3.BackoffStrategy.base_interval>`.
	MaxInterval          *types.Duration `protobuf:"bytes,2,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BackoffStrategy) Reset()         { *m = BackoffStrategy{} }
func (m *BackoffStrategy) String() string { return proto.CompactTextString(m) }
func (*BackoffStrategy) ProtoMessage()    {}
func (*BackoffStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f962803a44bd07, []int{0}
}
func (m *BackoffStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackoffStrategy.Unmarshal(m, b)
}
func (m *BackoffStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackoffStrategy.Marshal(b, m, deterministic)
}
func (m *BackoffStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackoffStrategy.Merge(m, src)
}
func (m *BackoffStrategy) XXX_Size() int {
	return xxx_messageInfo_BackoffStrategy.Size(m)
}
func (m *BackoffStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BackoffStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BackoffStrategy proto.InternalMessageInfo

func (m *BackoffStrategy) GetBaseInterval() *types.Duration {
	if m != nil {
		return m.BaseInterval
	}
	return nil
}

func (m *BackoffStrategy) GetMaxInterval() *types.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*BackoffStrategy)(nil), "envoy.config.core.v3.BackoffStrategy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/backoff.proto", fileDescriptor_83f962803a44bd07)
}

var fileDescriptor_83f962803a44bd07 = []byte{
<<<<<<< HEAD
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0xff, 0xd7, 0x7f, 0x95, 0x92, 0x56, 0x2c, 0xa1, 0xa0, 0x76, 0x28, 0xd2, 0x49, 0x04,
	0xef, 0x85, 0x76, 0x76, 0x09, 0x0e, 0x8a, 0x4b, 0xa9, 0xb8, 0xb8, 0xc8, 0x25, 0x7d, 0x73, 0x9e,
	0x4d, 0xf3, 0x86, 0xeb, 0x35, 0x5c, 0x77, 0x17, 0xbf, 0x89, 0x14, 0xdc, 0xc5, 0xa9, 0x9f, 0xc5,
	0xcd, 0x0f, 0xe0, 0x2e, 0xc9, 0xa5, 0xb8, 0x08, 0x75, 0x7b, 0xf2, 0xe4, 0xb9, 0xdf, 0x3d, 0xf7,
	0xbe, 0xde, 0xad, 0x54, 0xe6, 0x61, 0x11, 0xf2, 0x88, 0x66, 0x30, 0xa7, 0x84, 0xce, 0x14, 0x81,
	0x4c, 0x88, 0x20, 0xd3, 0xf4, 0x88, 0x91, 0x99, 0xbb, 0x2f, 0x91, 0x29, 0x40, 0x6b, 0x50, 0xa7,
	0x22, 0x01, 0x4c, 0x73, 0x5a, 0x42, 0x44, 0x69, 0xac, 0x24, 0x44, 0xa4, 0x11, 0xf2, 0x21, 0x84,
	0x22, 0x9a, 0x52, 0x1c, 0xf3, 0x4c, 0x93, 0x21, 0xbf, 0x53, 0x66, 0xb8, 0xcb, 0xf0, 0x22, 0xc3,
	0xf3, 0x61, 0xb7, 0x27, 0x89, 0x64, 0x82, 0x50, 0x66, 0xc2, 0x45, 0x0c, 0x93, 0x85, 0x16, 0x46,
	0x51, 0xea, 0x4e, 0x75, 0x0f, 0x72, 0x91, 0xa8, 0x89, 0x30, 0x08, 0x1b, 0x51, 0xfd, 0xe8, 0x48,
	0x92, 0x54, 0x4a, 0x28, 0x54, 0xe5, 0xfa, 0x68, 0x8d, 0x33, 0xd1, 0x1a, 0xe7, 0xf5, 0x5f, 0x99,
	0xb7, 0x1f, 0xb8, 0x2a, 0x37, 0x46, 0x0b, 0x83, 0x72, 0xe9, 0x8f, 0xbc, 0xbd, 0x50, 0xcc, 0xf1,
	0x5e, 0xa5, 0x06, 0x75, 0x2e, 0x92, 0x43, 0x76, 0xcc, 0x4e, 0x9a, 0x83, 0x23, 0xee, 0xea, 0xf0,
	0x4d, 0x1d, 0x7e, 0x51, 0xd5, 0x09, 0xda, 0xef, 0x9f, 0xeb, 0xff, 0xcd, 0x15, 0x6b, 0x34, 0xd8,
	0xa0, 0xde, 0x5e, 0x3f, 0x9d, 0x8f, 0x5b, 0x05, 0xe1, 0xaa, 0x02, 0xf8, 0x97, 0x5e, 0x6b, 0x26,
	0xec, 0x0f, 0xb0, 0xb6, 0x0d, 0xe8, 0x15, 0xc0, 0x9d, 0x15, 0xab, 0x9d, 0xfe, 0x1b, 0x37, 0x67,
	0xc2, 0x6e, 0x48, 0xc1, 0x33, 0x7b, 0xfb, 0xaa, 0xb3, 0x97, 0x8f, 0x1e, 0xf3, 0xfa, 0x8a, 0x78,
	0x39, 0xb6, 0x4c, 0x93, 0x5d, 0xf2, 0xdf, 0x26, 0x18, 0xb4, 0xaa, 0xb7, 0x8d, 0x8a, 0x1b, 0x46,
	0xec, 0xee, 0xfa, 0x6f, 0xeb, 0xcb, 0xa6, 0x72, 0xfb, 0x0a, 0xc3, 0xdd, 0xb2, 0xf7, 0xf0, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x70, 0x13, 0x67, 0xa7, 0x14, 0x02, 0x00, 0x00,
=======
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xbe, 0x13, 0xee, 0x2d, 0x25, 0xed, 0xd5, 0x52, 0x0a, 0x6a, 0xc1, 0x62, 0xeb, 0xa6, 0x08,
	0xce, 0x40, 0xbb, 0x13, 0xdc, 0x04, 0x11, 0x44, 0x90, 0x52, 0x71, 0xe3, 0x46, 0x26, 0xe9, 0x64,
	0x1c, 0x9b, 0xce, 0x09, 0x93, 0x49, 0x48, 0x77, 0x2e, 0x5c, 0xb8, 0xf1, 0x05, 0x7c, 0x02, 0xe9,
	0x23, 0xb8, 0x72, 0x23, 0xb8, 0x15, 0xdf, 0xc0, 0xc7, 0xe8, 0x4a, 0x92, 0x49, 0x11, 0x54, 0xa8,
	0xbb, 0x33, 0xe7, 0xfb, 0xe1, 0x3b, 0x67, 0x8e, 0x7d, 0xc6, 0x85, 0xbe, 0x8c, 0x5d, 0xec, 0xc1,
	0x84, 0x44, 0x10, 0xc0, 0xae, 0x00, 0xc2, 0x03, 0x00, 0x12, 0x2a, 0xb8, 0x62, 0x9e, 0x8e, 0xcc,
	0x8b, 0x86, 0x82, 0xb0, 0x54, 0x33, 0x25, 0x69, 0x40, 0x98, 0x4c, 0x60, 0x4a, 0x3c, 0x90, 0xbe,
	0xe0, 0xc4, 0x03, 0xc5, 0x48, 0xd2, 0x27, 0x2e, 0xf5, 0xc6, 0xe0, 0xfb, 0x38, 0x54, 0xa0, 0xa1,
	0xde, 0xc8, 0x39, 0xd8, 0x70, 0x70, 0xc6, 0xc1, 0x49, 0xbf, 0xd9, 0xe2, 0x00, 0x3c, 0x60, 0x24,
	0xe7, 0xb8, 0xb1, 0x4f, 0x46, 0xb1, 0xa2, 0x5a, 0x80, 0x34, 0xaa, 0xe6, 0x66, 0x3c, 0x0a, 0x29,
	0xa1, 0x52, 0x82, 0xce, 0xdb, 0x11, 0x89, 0x34, 0xd5, 0x71, 0x54, 0xc0, 0xed, 0x6f, 0x70, 0xc2,
	0x54, 0x24, 0x40, 0x0a, 0xc9, 0x0b, 0xca, 0x5a, 0x42, 0x03, 0x31, 0xa2, 0x9a, 0x91, 0x45, 0x51,
	0x00, 0x0d, 0x0e, 0x1c, 0xf2, 0x92, 0x64, 0x95, 0xe9, 0x76, 0xde, 0x90, 0xbd, 0xea, 0x98, 0xe0,
	0xa7, 0x5a, 0x51, 0xcd, 0xf8, 0xb4, 0x7e, 0x62, 0xff, 0x77, 0x69, 0xc4, 0x2e, 0x84, 0xd4, 0x4c,
	0x25, 0x34, 0x58, 0x47, 0x5b, 0xa8, 0x5b, 0xe9, 0x6d, 0x60, 0x13, 0x1e, 0x2f, 0xc2, 0xe3, 0x83,
	0x22, 0xbc, 0xb3, 0x32, 0x77, 0x2a, 0x33, 0x54, 0x2e, 0xa3, 0xde, 0xdf, 0xda, 0xd3, 0xcd, 0xfe,
	0xb0, 0x9a, 0xe9, 0x8f, 0x0a, 0x79, 0xfd, 0xd0, 0xae, 0x4e, 0x68, 0xfa, 0x69, 0x67, 0x2d, 0xb3,
	0x2b, 0xcf, 0x9d, 0x7f, 0x33, 0x64, 0xed, 0xfc, 0x19, 0x56, 0x26, 0x34, 0x5d, 0xf8, 0xec, 0x75,
	0xef, 0x9f, 0x6f, 0x5b, 0xdb, 0x76, 0xdb, 0x6c, 0x96, 0x86, 0x02, 0x27, 0x3d, 0xb3, 0xd9, 0x2f,
	0x13, 0x38, 0x77, 0xe8, 0xe1, 0xbd, 0x85, 0x1e, 0xaf, 0x5f, 0x5e, 0x4b, 0x56, 0xcd, 0xb2, 0x3b,
	0x02, 0x70, 0x2e, 0x0a, 0x15, 0xa4, 0x53, 0xfc, 0xd3, 0xcf, 0x38, 0xd5, 0xc2, 0x63, 0x90, 0xa5,
	0x19, 0xa0, 0xf3, 0xe3, 0xdf, 0x9d, 0x45, 0x38, 0xe6, 0xcb, 0x4f, 0xc3, 0x2d, 0xe5, 0x33, 0xf6,
	0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0x42, 0xac, 0x10, 0x6c, 0x02, 0x00, 0x00,
>>>>>>> master
}

func (this *BackoffStrategy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BackoffStrategy)
	if !ok {
		that2, ok := that.(BackoffStrategy)
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
	if !this.BaseInterval.Equal(that1.BaseInterval) {
		return false
	}
	if !this.MaxInterval.Equal(that1.MaxInterval) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
