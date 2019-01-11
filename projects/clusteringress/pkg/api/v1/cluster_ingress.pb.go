// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/clusteringress/api/v1/cluster_ingress.proto

package v1 // import "github.com/solo-io/gloo/projects/clusteringress/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
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
// A simple wrapper for a kNative ClusterIngress Object.
type ClusterIngress struct {
	Metadata core.Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata"`
	// a raw byte representation of the cluster ingress this resource wraps
	ClusterIngressSpec *types.Any `protobuf:"bytes,2,opt,name=cluster_ingress_spec,json=clusterIngressSpec" json:"cluster_ingress_spec,omitempty"`
	// a raw byte representation of the ingress status of the cluster ingress object
	ClusterIngressStatus *types.Any `protobuf:"bytes,3,opt,name=cluster_ingress_status,json=clusterIngressStatus" json:"cluster_ingress_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ClusterIngress) Reset()         { *m = ClusterIngress{} }
func (m *ClusterIngress) String() string { return proto.CompactTextString(m) }
func (*ClusterIngress) ProtoMessage()    {}
func (*ClusterIngress) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_ingress_f001847f5884e124, []int{0}
}
func (m *ClusterIngress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterIngress.Unmarshal(m, b)
}
func (m *ClusterIngress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterIngress.Marshal(b, m, deterministic)
}
func (dst *ClusterIngress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterIngress.Merge(dst, src)
}
func (m *ClusterIngress) XXX_Size() int {
	return xxx_messageInfo_ClusterIngress.Size(m)
}
func (m *ClusterIngress) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterIngress.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterIngress proto.InternalMessageInfo

func (m *ClusterIngress) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *ClusterIngress) GetClusterIngressSpec() *types.Any {
	if m != nil {
		return m.ClusterIngressSpec
	}
	return nil
}

func (m *ClusterIngress) GetClusterIngressStatus() *types.Any {
	if m != nil {
		return m.ClusterIngressStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterIngress)(nil), "clusteringress.gloo.solo.io.ClusterIngress")
}
func (this *ClusterIngress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterIngress)
	if !ok {
		that2, ok := that.(ClusterIngress)
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
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.ClusterIngressSpec.Equal(that1.ClusterIngressSpec) {
		return false
	}
	if !this.ClusterIngressStatus.Equal(that1.ClusterIngressStatus) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/clusteringress/api/v1/cluster_ingress.proto", fileDescriptor_cluster_ingress_f001847f5884e124)
}

var fileDescriptor_cluster_ingress_f001847f5884e124 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x4f, 0xce, 0x29,
	0x2d, 0x2e, 0x49, 0x2d, 0xca, 0xcc, 0x4b, 0x2f, 0x4a, 0x2d, 0x2e, 0xd6, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0x84, 0x89, 0xc6, 0x43, 0x85, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0x51,
	0x15, 0xeb, 0x81, 0x4c, 0xd2, 0x03, 0x19, 0xab, 0x97, 0x99, 0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e,
	0x0f, 0x56, 0xa7, 0x0f, 0x62, 0x41, 0xb4, 0x48, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea,
	0x83, 0x79, 0x49, 0xa5, 0x69, 0xfa, 0x89, 0x79, 0x95, 0x50, 0x29, 0x43, 0x2c, 0x6e, 0x03, 0xd3,
	0xd9, 0x99, 0x25, 0x30, 0x77, 0xe4, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0x92, 0xa0, 0x05,
	0xc6, 0x87, 0x68, 0x51, 0x6a, 0x65, 0xe2, 0xe2, 0x73, 0x86, 0x38, 0xdb, 0x13, 0xe2, 0x6c, 0x21,
	0x0b, 0x2e, 0x0e, 0x98, 0xb9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0xc9, 0xf9,
	0x45, 0xa9, 0x30, 0xaf, 0xe8, 0xf9, 0x42, 0x65, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82,
	0xab, 0x16, 0x72, 0xe3, 0x12, 0x41, 0x0b, 0x99, 0xf8, 0xe2, 0x82, 0xd4, 0x64, 0x09, 0x26, 0xb0,
	0x29, 0x22, 0x7a, 0x10, 0xcf, 0xea, 0xc1, 0x3c, 0xab, 0xe7, 0x98, 0x57, 0x19, 0x24, 0x94, 0x8c,
	0x62, 0x7b, 0x70, 0x41, 0x6a, 0xb2, 0x50, 0x00, 0x97, 0x18, 0x86, 0x39, 0x25, 0x89, 0x25, 0xa5,
	0xc5, 0x12, 0xcc, 0xb8, 0x4d, 0x72, 0x62, 0x69, 0xf8, 0xc8, 0xc2, 0x18, 0x24, 0x82, 0x66, 0x1e,
	0x58, 0x9f, 0x95, 0x7c, 0xd3, 0x47, 0x16, 0x56, 0x2e, 0xe6, 0xe4, 0xcc, 0xf4, 0xa6, 0x8f, 0x2c,
	0x42, 0x42, 0x02, 0xa8, 0x31, 0x95, 0x5a, 0xec, 0xe4, 0xb0, 0xe2, 0x91, 0x1c, 0x63, 0x94, 0x15,
	0xa9, 0xe9, 0xa1, 0x20, 0x3b, 0x1d, 0x1a, 0xb0, 0x49, 0x6c, 0x60, 0xc7, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x24, 0xe0, 0x0b, 0x00, 0x54, 0x02, 0x00, 0x00,
}
