// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/healthcheck/healthcheck.proto

package healthcheck

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Add this config to a Listener/Gateway to Enable Envoy Health Checks on that port
type HealthCheck struct {
	// match health check requests using this exact path
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b46744a4a36ca59a, []int{0}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "healthcheck.plugins.gloo.solo.io.HealthCheck")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/healthcheck/healthcheck.proto", fileDescriptor_b46744a4a36ca59a)
}

var fileDescriptor_b46744a4a36ca59a = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0x33, 0x52, 0x13,
	0x73, 0x4a, 0x32, 0x92, 0x33, 0x52, 0x93, 0xb3, 0x91, 0xd9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x0a, 0x28, 0x42, 0x10, 0x2d, 0x7a, 0x20, 0x63, 0xf4, 0x40, 0x36, 0xe8, 0x65, 0xe6, 0x4b,
	0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x15, 0xeb, 0x83, 0x58, 0x10, 0x7d, 0x52, 0x3a, 0x58, 0xdc,
	0x02, 0xa6, 0xb3, 0x33, 0x4b, 0x60, 0x2e, 0x28, 0x4a, 0x4d, 0x83, 0xaa, 0x96, 0x4b, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0xcb, 0x8b, 0x12, 0x0b, 0x0a, 0x52,
	0x8b, 0x8a, 0x21, 0xf2, 0x4a, 0x8a, 0x5c, 0xdc, 0x1e, 0x60, 0x77, 0x38, 0x83, 0xdc, 0x21, 0x24,
	0xc4, 0xc5, 0x52, 0x90, 0x58, 0x92, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x3b,
	0x79, 0xad, 0x78, 0x24, 0xc7, 0x18, 0xe5, 0x42, 0x5c, 0x10, 0x14, 0x64, 0xa7, 0xe3, 0x09, 0x86,
	0x24, 0x36, 0xb0, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xbe, 0xac, 0x4a, 0x51,
	0x01, 0x00, 0x00,
}

func (this *HealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
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
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
