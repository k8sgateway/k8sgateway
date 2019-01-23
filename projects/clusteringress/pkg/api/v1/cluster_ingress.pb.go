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
	Status   core.Status   `protobuf:"bytes,4,opt,name=status" json:"status"`
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
	return fileDescriptor_cluster_ingress_4c6203f93b3010e9, []int{0}
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

func (m *ClusterIngress) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
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
	if !this.Status.Equal(&that1.Status) {
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
	proto.RegisterFile("github.com/solo-io/gloo/projects/clusteringress/api/v1/cluster_ingress.proto", fileDescriptor_cluster_ingress_4c6203f93b3010e9)
}

var fileDescriptor_cluster_ingress_4c6203f93b3010e9 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xff, 0xdd, 0xf2, 0x1f, 0x12, 0x41, 0x24, 0x94, 0x51, 0x27, 0xa8, 0x08, 0x82, 0x37,
	0x26, 0x6c, 0xde, 0xc8, 0xae, 0x74, 0x82, 0x20, 0x28, 0xc8, 0x76, 0xe7, 0xcd, 0xe8, 0x62, 0x8c,
	0x71, 0xdd, 0x4e, 0x69, 0x52, 0x61, 0x77, 0xb2, 0xa7, 0xf1, 0x2d, 0xbc, 0xf5, 0x29, 0xbc, 0xf0,
	0x0d, 0xfa, 0x06, 0xb2, 0x24, 0x15, 0x3a, 0x15, 0xdc, 0x55, 0x9b, 0x9c, 0xef, 0xf7, 0xf5, 0x3b,
	0xa7, 0x07, 0x5f, 0x49, 0x65, 0x1e, 0xf2, 0x11, 0xe5, 0x30, 0x61, 0x1a, 0x12, 0x38, 0x52, 0xc0,
	0x64, 0x02, 0xc0, 0xd2, 0x0c, 0x1e, 0x05, 0x37, 0x9a, 0xf1, 0x24, 0xd7, 0x46, 0x64, 0x6a, 0x2a,
	0x33, 0xa1, 0x35, 0x8b, 0x53, 0xc5, 0x9e, 0xda, 0xe5, 0xed, 0xd0, 0x5f, 0xd3, 0x34, 0x03, 0x03,
	0x64, 0xbb, 0x2a, 0xa6, 0x0b, 0x27, 0xba, 0xb0, 0xa5, 0x0a, 0x5a, 0xa1, 0x04, 0x09, 0x56, 0xc7,
	0x16, 0x6f, 0x0e, 0x69, 0x6d, 0x49, 0x00, 0x99, 0x08, 0x66, 0x4f, 0xa3, 0xfc, 0x9e, 0xc5, 0xd3,
	0x99, 0x2f, 0xb5, 0x7f, 0xc8, 0x66, 0x9f, 0x63, 0x65, 0xca, 0x1c, 0x13, 0x61, 0xe2, 0xbb, 0xd8,
	0xc4, 0x1e, 0x61, 0x7f, 0x40, 0xb4, 0x89, 0x4d, 0xae, 0x57, 0xf8, 0x46, 0x79, 0x76, 0xc8, 0xfe,
	0x6b, 0x0d, 0x6f, 0x9c, 0xbb, 0x3e, 0x2f, 0x5d, 0x9f, 0xe4, 0x04, 0xaf, 0x95, 0x41, 0xa2, 0x60,
	0x2f, 0x38, 0x5c, 0xef, 0x34, 0x29, 0x87, 0x4c, 0x94, 0xbd, 0xd3, 0x6b, 0x5f, 0xed, 0xa1, 0xb7,
	0xf7, 0xdd, 0x7f, 0xfd, 0x2f, 0x35, 0xe9, 0xe0, 0x86, 0xcb, 0x13, 0x21, 0xcb, 0x85, 0x55, 0x6e,
	0x60, 0x6b, 0x9e, 0xf2, 0x4a, 0x72, 0x81, 0xc3, 0xa5, 0xf1, 0x0f, 0x75, 0x2a, 0x78, 0x54, 0xf3,
	0x0e, 0x6e, 0xa2, 0xb4, 0x9c, 0x28, 0x3d, 0x9b, 0xce, 0xfa, 0x84, 0x57, 0x12, 0x0f, 0x52, 0xc1,
	0xc9, 0x0d, 0x6e, 0x7e, 0xf3, 0x71, 0x59, 0xea, 0xbf, 0x3b, 0xf5, 0xd0, 0x73, 0x81, 0x82, 0x7e,
	0xb8, 0xe4, 0x67, 0xb9, 0xee, 0xc1, 0xbc, 0x40, 0xff, 0x71, 0x9d, 0x2b, 0x39, 0x2f, 0x10, 0x21,
	0x9b, 0xd5, 0x75, 0x10, 0x7a, 0x5e, 0xa0, 0x5a, 0x14, 0xf4, 0x4e, 0x5f, 0x3e, 0x76, 0x82, 0xdb,
	0xee, 0xaa, 0xab, 0x97, 0x8e, 0xa5, 0xff, 0x25, 0xa3, 0x86, 0x8d, 0x74, 0xfc, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x05, 0xd1, 0x0c, 0xdf, 0xbf, 0x02, 0x00, 0x00,
}
