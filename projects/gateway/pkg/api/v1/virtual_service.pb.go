// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto

package v1 // import "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=vs
// @solo-kit:resource.plural_name=virtual_services
// @solo-kit:resource.resource_groups=api.gateway.solo.io
// A virtual service describes the set of routes to match for a set of domains.
// Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
type VirtualService struct {
	VirtualHost *v1.VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost" json:"virtual_host,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic for this set of routes
	SslConfig *v1.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig" json:"ssl_config,omitempty"`
	// Since virtual services are long-lived the purpose of a virtual service may change over time.
	// Users may like to update the name to better reflect the purpose of the virtual service.
	// Since metadata.name is associated with the resource id, it is not possible to change that
	// value without deleting and recreating a resource.
	// If additional fields such as this are needed, consider implementing a separate "auxillary data" object
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VirtualService) Reset()         { *m = VirtualService{} }
func (m *VirtualService) String() string { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()    {}
func (*VirtualService) Descriptor() ([]byte, []int) {
	return fileDescriptor_virtual_service_2e2e6bb5e6dc7f41, []int{0}
}
func (m *VirtualService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualService.Unmarshal(m, b)
}
func (m *VirtualService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualService.Marshal(b, m, deterministic)
}
func (dst *VirtualService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualService.Merge(dst, src)
}
func (m *VirtualService) XXX_Size() int {
	return xxx_messageInfo_VirtualService.Size(m)
}
func (m *VirtualService) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualService.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualService proto.InternalMessageInfo

func (m *VirtualService) GetVirtualHost() *v1.VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetSslConfig() *v1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *VirtualService) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *VirtualService) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *VirtualService) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
}
func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
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
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service.proto", fileDescriptor_virtual_service_2e2e6bb5e6dc7f41)
}

var fileDescriptor_virtual_service_2e2e6bb5e6dc7f41 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0xb9, 0x37, 0x5c, 0x19, 0x88, 0xc6, 0x86, 0x68, 0x61, 0x21, 0xc8, 0x8a, 0x8d,
	0x9d, 0x20, 0x89, 0x41, 0xe3, 0x0a, 0x63, 0x74, 0xa3, 0x8b, 0x92, 0xb8, 0x70, 0x43, 0x86, 0x32,
	0x0c, 0x23, 0x2d, 0xa7, 0xe9, 0x1c, 0xaa, 0xbc, 0x91, 0x0b, 0x1f, 0xc4, 0xa7, 0x60, 0xe1, 0x23,
	0xf8, 0x04, 0xa6, 0xd3, 0xa9, 0x84, 0xc4, 0x44, 0x56, 0x3d, 0xa7, 0xff, 0xf9, 0xfe, 0x93, 0xff,
	0x64, 0xc8, 0xb5, 0x90, 0x38, 0x5d, 0x8c, 0x5c, 0x1f, 0x42, 0xaa, 0x20, 0x80, 0x13, 0x09, 0x54,
	0x04, 0x00, 0x34, 0x8a, 0xe1, 0x89, 0xfb, 0xa8, 0xa8, 0x60, 0xc8, 0x9f, 0xd9, 0x92, 0xb2, 0x48,
	0xd2, 0xa4, 0x43, 0x13, 0x19, 0xe3, 0x82, 0x05, 0x43, 0xc5, 0xe3, 0x44, 0xfa, 0xdc, 0x8d, 0x62,
	0x40, 0xb0, 0xf7, 0xcc, 0x94, 0x9b, 0x7a, 0xb8, 0x12, 0xea, 0x55, 0x01, 0x02, 0xb4, 0x46, 0xd3,
	0x2a, 0x1b, 0xab, 0x77, 0x7e, 0xd8, 0xa6, 0xbf, 0x33, 0x89, 0xf9, 0x82, 0x90, 0x23, 0x1b, 0x33,
	0x64, 0x06, 0xa1, 0x5b, 0x20, 0x0a, 0x19, 0x2e, 0x94, 0x01, 0x7a, 0xbf, 0x27, 0x4a, 0x3b, 0x83,
	0x46, 0x31, 0xbc, 0x2c, 0x33, 0xb2, 0xf5, 0x56, 0x20, 0xbb, 0x0f, 0x59, 0xbc, 0x41, 0x96, 0xce,
	0xbe, 0x24, 0x95, 0x3c, 0xf0, 0x14, 0x14, 0x3a, 0x56, 0xd3, 0x6a, 0x97, 0x4f, 0x6b, 0x6e, 0x6a,
	0x91, 0x67, 0x75, 0x0d, 0x73, 0x0b, 0x0a, 0xbd, 0x72, 0xb2, 0x6e, 0xec, 0x33, 0x42, 0x94, 0x0a,
	0x86, 0x3e, 0xcc, 0x27, 0x52, 0x38, 0x05, 0xcd, 0x1e, 0x6e, 0xb2, 0x03, 0x15, 0x5c, 0x69, 0xd9,
	0x2b, 0xa9, 0xbc, 0xb4, 0x8f, 0x49, 0x65, 0x2c, 0x55, 0x14, 0xb0, 0xe5, 0x70, 0xce, 0x42, 0xee,
	0xfc, 0x6d, 0x5a, 0xed, 0x92, 0x57, 0x36, 0xff, 0xee, 0x59, 0xc8, 0xed, 0x1b, 0x52, 0xcc, 0x52,
	0x3b, 0x45, 0x6d, 0x5b, 0x75, 0x7d, 0x88, 0xf9, 0xda, 0x56, 0x6b, 0xfd, 0xda, 0xfb, 0xaa, 0xf1,
	0xe7, 0x73, 0xd5, 0xd8, 0x47, 0xae, 0x70, 0x2c, 0x27, 0x93, 0x8b, 0x96, 0x14, 0x73, 0x88, 0x79,
	0xcb, 0x33, 0xb8, 0xdd, 0x23, 0x3b, 0xf9, 0xc5, 0x9d, 0xff, 0xda, 0xea, 0x60, 0xd3, 0xea, 0xce,
	0xa8, 0xfd, 0x7f, 0xa9, 0x99, 0xf7, 0x3d, 0xdd, 0x3f, 0x7f, 0xfd, 0x38, 0xb2, 0x1e, 0xbb, 0x5b,
	0x3f, 0xa0, 0x68, 0x26, 0xcc, 0xd5, 0x47, 0x45, 0x7d, 0xf0, 0xee, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0xfd, 0x7f, 0x19, 0x7e, 0x02, 0x00, 0x00,
}
