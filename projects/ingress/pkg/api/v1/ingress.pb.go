// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//A simple wrapper for a Kubernetes Ingress Object.
type Ingress struct {
	// a raw byte representation of the kubernetes ingress this resource wraps
	KubeIngressSpec *types.Any `protobuf:"bytes,1,opt,name=kube_ingress_spec,json=kubeIngressSpec,proto3" json:"kube_ingress_spec,omitempty"`
	// a raw byte representation of the ingress status of the kubernetes ingress object
	KubeIngressStatus *types.Any `protobuf:"bytes,2,opt,name=kube_ingress_status,json=kubeIngressStatus,proto3" json:"kube_ingress_status,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ingress) Reset()         { *m = Ingress{} }
func (m *Ingress) String() string { return proto.CompactTextString(m) }
func (*Ingress) ProtoMessage()    {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e3857ca3a6e6b32, []int{0}
}
func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingress.Unmarshal(m, b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingress.Marshal(b, m, deterministic)
}
func (m *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(m, src)
}
func (m *Ingress) XXX_Size() int {
	return xxx_messageInfo_Ingress.Size(m)
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *Ingress) GetKubeIngressSpec() *types.Any {
	if m != nil {
		return m.KubeIngressSpec
	}
	return nil
}

func (m *Ingress) GetKubeIngressStatus() *types.Any {
	if m != nil {
		return m.KubeIngressStatus
	}
	return nil
}

func (m *Ingress) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Ingress)(nil), "ingress.solo.io.Ingress")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto", fileDescriptor_7e3857ca3a6e6b32)
}

var fileDescriptor_7e3857ca3a6e6b32 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbf, 0x4f, 0x03, 0x21,
	0x18, 0x95, 0x7a, 0xb1, 0x4a, 0x87, 0xa6, 0xd8, 0x68, 0xed, 0xa0, 0xc6, 0xc9, 0x45, 0x48, 0xed,
	0x62, 0x8c, 0x26, 0x7a, 0x5b, 0x07, 0x97, 0xba, 0xb9, 0x34, 0xdc, 0x89, 0x88, 0xfd, 0xf1, 0x91,
	0x83, 0x33, 0xed, 0xda, 0xbf, 0xc6, 0x3f, 0xc5, 0xbf, 0xc2, 0xc1, 0xcd, 0xb1, 0x5b, 0x47, 0x73,
	0x1c, 0x18, 0x4d, 0x34, 0xa9, 0x13, 0x3c, 0xbe, 0xf7, 0x3e, 0xde, 0xfb, 0x00, 0x5f, 0x4a, 0x65,
	0x1f, 0xf3, 0x84, 0xa6, 0x30, 0x66, 0x06, 0x46, 0x70, 0xa2, 0x80, 0xc9, 0x11, 0x00, 0xd3, 0x19,
	0x3c, 0x89, 0xd4, 0x1a, 0xa6, 0x26, 0x32, 0x13, 0xc6, 0x30, 0xae, 0x15, 0x7b, 0xee, 0x04, 0x48,
	0x75, 0x06, 0x16, 0x48, 0x3d, 0xc0, 0x42, 0x4b, 0x15, 0xb4, 0x9b, 0x12, 0x24, 0xb8, 0x1a, 0x2b,
	0x76, 0x25, 0xad, 0xbd, 0x27, 0x01, 0xe4, 0x48, 0x30, 0x87, 0x92, 0xfc, 0x81, 0xf1, 0xc9, 0xcc,
	0x97, 0x88, 0x98, 0xda, 0x92, 0x2f, 0xa6, 0xd6, 0x9f, 0x75, 0x7e, 0x31, 0xe5, 0xd6, 0xa1, 0xb2,
	0xc1, 0xc7, 0x58, 0x58, 0x7e, 0xcf, 0x2d, 0xf7, 0x12, 0xb6, 0x82, 0xc4, 0x58, 0x6e, 0x73, 0xf3,
	0x8f, 0x3b, 0x02, 0x2e, 0x25, 0x47, 0x1f, 0x08, 0x57, 0x7b, 0x65, 0x5e, 0x72, 0x85, 0x1b, 0xc3,
	0x3c, 0x11, 0x03, 0x9f, 0x7f, 0x60, 0xb4, 0x48, 0x5b, 0xe8, 0x10, 0x1d, 0xd7, 0x4e, 0x9b, 0xb4,
	0x4c, 0x4b, 0x43, 0x5a, 0x7a, 0x3d, 0x99, 0xf5, 0xeb, 0x05, 0xdd, 0xab, 0x6f, 0xb5, 0x48, 0x49,
	0x0f, 0x6f, 0xff, 0xec, 0xe0, 0xdc, 0xb5, 0x2a, 0x7f, 0xf7, 0x88, 0xd7, 0x97, 0x31, 0xea, 0x37,
	0xbe, 0x37, 0x72, 0x1a, 0x72, 0x86, 0x37, 0xc3, 0x38, 0x5a, 0x55, 0xa7, 0xdf, 0xa1, 0x29, 0x64,
	0x22, 0xbc, 0x0a, 0xbd, 0xf1, 0xd5, 0x38, 0x7a, 0x7d, 0x3b, 0x58, 0xeb, 0x7f, 0xb1, 0xcf, 0x77,
	0xe7, 0x8b, 0x28, 0xc2, 0x15, 0x25, 0xe7, 0x8b, 0xa8, 0x46, 0xb6, 0xbc, 0x17, 0x61, 0xe2, 0x8b,
	0x65, 0x8c, 0x5e, 0xde, 0xf7, 0xd1, 0x5d, 0x77, 0xe5, 0x0f, 0xa2, 0x87, 0xd2, 0x0f, 0x2e, 0xd9,
	0x70, 0xb6, 0xbb, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0xfe, 0x8c, 0x8f, 0x5e, 0x02, 0x00,
	0x00,
}

func (this *Ingress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ingress)
	if !ok {
		that2, ok := that.(Ingress)
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
	if !this.KubeIngressSpec.Equal(that1.KubeIngressSpec) {
		return false
	}
	if !this.KubeIngressStatus.Equal(that1.KubeIngressStatus) {
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
