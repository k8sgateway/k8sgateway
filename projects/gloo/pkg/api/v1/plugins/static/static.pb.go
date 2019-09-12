// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/static/static.proto

package static

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"
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

// Static upstreams are used to route request to services listening at fixed IP/Addresses.
// Static upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Static Upstreams must be created manually by users
type UpstreamSpec struct {
	// A list of addresses and ports
	// at least one must be specified
	Hosts []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Attempt to use outbound TLS
	// Gloo will automatically set this to true for port 443
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// Enable \ Disable auto host re-write.
	// If not set, the default behavior is to enable auto host rewrite, if the first host in the
	// list is not an IP address (i.e. it is a DNS address)
	// If enabled (explicitly or by the default) auto_host_rewrite will be automatically added
	// to routes that point to this upstream.
	AutoHostRewrite *types.BoolValue `protobuf:"bytes,6,opt,name=auto_host_rewrite,json=autoHostRewrite,proto3" json:"auto_host_rewrite,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec          *plugins.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe915d162b6f8af, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetHosts() []*Host {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *UpstreamSpec) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *UpstreamSpec) GetAutoHostRewrite() *types.BoolValue {
	if m != nil {
		return m.AutoHostRewrite
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *plugins.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

// Represents a single instance of an upstream
type Host struct {
	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffe915d162b6f8af, []int{1}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Host) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "static.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*Host)(nil), "static.plugins.gloo.solo.io.Host")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/static/static.proto", fileDescriptor_ffe915d162b6f8af)
}

var fileDescriptor_ffe915d162b6f8af = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4a, 0x2b, 0x31,
	0x18, 0x25, 0xb7, 0x3f, 0xf7, 0xde, 0xb4, 0x22, 0x0e, 0x82, 0x43, 0x85, 0x32, 0xed, 0x6a, 0x36,
	0x26, 0x58, 0x17, 0x2e, 0x85, 0xe2, 0x4f, 0xd7, 0x53, 0x75, 0xe1, 0x66, 0x98, 0x4e, 0x63, 0x1a,
	0x4d, 0xfd, 0x42, 0xbe, 0x4c, 0xfb, 0x4a, 0x3e, 0x97, 0x0f, 0xe0, 0x33, 0x48, 0x26, 0x2d, 0xba,
	0x28, 0x22, 0xae, 0xf2, 0x9d, 0xe4, 0x9c, 0x93, 0x73, 0x42, 0xe8, 0x44, 0x2a, 0xb7, 0xa8, 0x66,
	0xac, 0x84, 0x25, 0x47, 0xd0, 0x70, 0xa2, 0x80, 0x4b, 0x0d, 0xc0, 0x8d, 0x85, 0x27, 0x51, 0x3a,
	0x0c, 0xa8, 0x30, 0x8a, 0xaf, 0x4e, 0xb9, 0xd1, 0x95, 0x54, 0x2f, 0xc8, 0xd1, 0x15, 0x4e, 0x95,
	0x9b, 0x85, 0x19, 0x0b, 0x0e, 0xa2, 0xe3, 0x2d, 0x0a, 0x1c, 0xe6, 0x75, 0xcc, 0x5b, 0x32, 0x05,
	0xbd, 0x43, 0x09, 0x12, 0x6a, 0x1e, 0xf7, 0x53, 0x90, 0xf4, 0xfa, 0x12, 0x40, 0x6a, 0xc1, 0x6b,
	0x34, 0xab, 0x1e, 0xf9, 0xda, 0x16, 0xc6, 0x08, 0x8b, 0x9b, 0xf3, 0x9b, 0xdf, 0x85, 0x13, 0x76,
	0xa5, 0x4a, 0x91, 0xa3, 0x11, 0x9b, 0x6c, 0xc3, 0x77, 0x42, 0xbb, 0x77, 0x06, 0x9d, 0x15, 0xc5,
	0x72, 0x6a, 0x44, 0x19, 0x9d, 0xd3, 0xd6, 0x02, 0xd0, 0x61, 0x4c, 0x92, 0x46, 0xda, 0x19, 0x0d,
	0xd8, 0x37, 0xe1, 0xd9, 0x04, 0xd0, 0x65, 0x81, 0x1f, 0x1d, 0xd1, 0xbf, 0x15, 0x8a, 0xdc, 0x69,
	0x8c, 0x1b, 0x09, 0x49, 0xff, 0x65, 0xed, 0x0a, 0xc5, 0xad, 0xc6, 0xe8, 0x9a, 0x1e, 0x14, 0x95,
	0x83, 0xdc, 0xd3, 0x72, 0x2b, 0xd6, 0x56, 0x39, 0x11, 0xb7, 0x13, 0x92, 0x76, 0x46, 0x3d, 0x16,
	0x7a, 0xb2, 0x6d, 0x4f, 0x36, 0x06, 0xd0, 0xf7, 0x85, 0xae, 0x44, 0xb6, 0xef, 0x45, 0xf5, 0x05,
	0x41, 0x12, 0x5d, 0xd2, 0xee, 0xd7, 0x02, 0x71, 0xab, 0xb6, 0x18, 0xec, 0x4e, 0x36, 0x0d, 0x4c,
	0x5f, 0x29, 0xeb, 0xe0, 0x27, 0x18, 0x32, 0xda, 0xf4, 0xa6, 0x51, 0x44, 0x9b, 0xc5, 0x7c, 0x6e,
	0x63, 0x92, 0x90, 0xf4, 0x7f, 0x56, 0xcf, 0x7e, 0xcf, 0x80, 0x75, 0xf1, 0x9f, 0x84, 0xa4, 0x7b,
	0x59, 0x3d, 0x8f, 0xaf, 0x5e, 0xdf, 0xfa, 0xe4, 0xe1, 0xe2, 0x67, 0xef, 0x6d, 0x9e, 0xe5, 0xee,
	0x0f, 0x31, 0x6b, 0xd7, 0x0d, 0xcf, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x88, 0x40, 0x6e,
	0x56, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if !this.Hosts[i].Equal(that1.Hosts[i]) {
			return false
		}
	}
	if this.UseTls != that1.UseTls {
		return false
	}
	if !this.AutoHostRewrite.Equal(that1.AutoHostRewrite) {
		return false
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Host) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Host)
	if !ok {
		that2, ok := that.(Host)
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
	if this.Addr != that1.Addr {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
