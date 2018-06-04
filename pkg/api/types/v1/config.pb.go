// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	config.proto
	metadata.proto
	role.proto
	status.proto
	upstream.proto
	virtualservice.proto

It has these top-level messages:
	Config
	Metadata
	Role
	Status
	Upstream
	ServiceInfo
	Function
	VirtualService
	Route
	RequestMatcher
	EventMatcher
	WeightedDestination
	Destination
	FunctionDestination
	UpstreamDestination
	SSLConfig
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// Config is a top-level config object. It is used internally by gloo as a container for the entire set of config objects.
type Config struct {
	Upstreams       []*Upstream       `protobuf:"bytes,1,rep,name=upstreams" json:"upstreams,omitempty"`
	VirtualServices []*VirtualService `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices" json:"virtual_services,omitempty"`
	Roles           []*Role           `protobuf:"bytes,3,rep,name=roles" json:"roles,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetUpstreams() []*Upstream {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *Config) GetVirtualServices() []*VirtualService {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Config) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v1.Config")
}
func (this *Config) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
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
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(that1.VirtualServices[i]) {
			return false
		}
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if !this.Roles[i].Equal(that1.Roles[i]) {
			return false
		}
	}
	return true
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0xe2, 0x2b, 0x2d,
	0x28, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x85, 0x88, 0x49, 0x89, 0x94, 0x65, 0x16, 0x95, 0x94, 0x26,
	0xe6, 0x14, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x42, 0x45, 0xb9, 0x8a, 0xf2, 0x73, 0x60, 0x6c,
	0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x4d, 0x66, 0xe4, 0x62,
	0x73, 0x06, 0x1b, 0x2e, 0xa4, 0xc5, 0xc5, 0x09, 0x33, 0xb4, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83,
	0xdb, 0x88, 0x47, 0xaf, 0xcc, 0x50, 0x2f, 0x14, 0x2a, 0x18, 0x84, 0x90, 0x16, 0xb2, 0xe5, 0x12,
	0x80, 0x5a, 0x18, 0x0f, 0xb5, 0xb1, 0x58, 0x82, 0x09, 0xac, 0x45, 0x08, 0xa4, 0x25, 0x0c, 0x22,
	0x17, 0x0c, 0x91, 0x0a, 0xe2, 0x2f, 0x43, 0xe1, 0x17, 0x0b, 0xc9, 0x71, 0xb1, 0x82, 0x5c, 0x56,
	0x2c, 0xc1, 0x0c, 0xd6, 0xc3, 0x01, 0xd2, 0x13, 0x94, 0x9f, 0x93, 0x1a, 0x04, 0x11, 0x76, 0xd2,
	0x5b, 0xf1, 0x48, 0x8e, 0x31, 0x4a, 0x23, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0xbf, 0x38, 0x3f, 0x27, 0x5f, 0x37, 0x33, 0x5f, 0x3f, 0x3d, 0x27, 0x3f, 0x5f, 0xbf, 0x20,
	0x3b, 0x5d, 0x3f, 0xb1, 0x20, 0x53, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xbf, 0xcc, 0x30, 0x89,
	0x0d, 0xec, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x66, 0xe3, 0xc5, 0x28, 0x01,
	0x00, 0x00,
}
