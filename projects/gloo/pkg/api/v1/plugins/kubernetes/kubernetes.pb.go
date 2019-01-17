// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/kubernetes/kubernetes.proto

package kubernetes // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"

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

// Upstream Spec for Kubernetes Upstreams
// Kubernetes Upstreams represent a set of one or more addressable pods for a Kubernetes Service
// the Gloo Kubernetes Upstream maps to a single service port. Because Kubernetes Services support mulitple ports,
// Gloo requires that a different upstream be created for each port
// Kubernetes Upstreams are typically generated automatically by Gloo from the Kubernetes API
type UpstreamSpec struct {
	// The name of the Kubernetes Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The namespace where the Service lives
	ServiceNamespace string `protobuf:"bytes,2,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
	// The access port port of the kubernetes service is listening. This port is used by Gloo to look up the corresponding
	// port on the pod for routing.
	ServicePort uint32 `protobuf:"varint,3,opt,name=service_port,json=servicePort,proto3" json:"service_port,omitempty"`
	// Allows finer-grained filtering of pods for the Upstream. Gloo will select pods based on their labels if
	// any are provided here.
	// (see [Kubernetes labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//     An optional Service Spec describing the service listening at this address
	ServiceSpec          *plugins.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_kubernetes_b49d4d3b931cb1a7, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (dst *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(dst, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceNamespace() string {
	if m != nil {
		return m.ServiceNamespace
	}
	return ""
}

func (m *UpstreamSpec) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *UpstreamSpec) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *plugins.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "kubernetes.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterMapType((map[string]string)(nil), "kubernetes.plugins.gloo.solo.io.UpstreamSpec.SelectorEntry")
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
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.ServiceNamespace != that1.ServiceNamespace {
		return false
	}
	if this.ServicePort != that1.ServicePort {
		return false
	}
	if len(this.Selector) != len(that1.Selector) {
		return false
	}
	for i := range this.Selector {
		if this.Selector[i] != that1.Selector[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/kubernetes/kubernetes.proto", fileDescriptor_kubernetes_b49d4d3b931cb1a7)
}

var fileDescriptor_kubernetes_b49d4d3b931cb1a7 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x49, 0x6b, 0x45, 0xb7, 0x2d, 0xd4, 0xa5, 0x87, 0xd0, 0x83, 0xa6, 0x9e, 0x02, 0xe2,
	0x2e, 0xd6, 0x8b, 0xd8, 0x9b, 0x3f, 0x88, 0x17, 0x29, 0x2d, 0x22, 0x78, 0x91, 0x74, 0x19, 0x62,
	0x6c, 0x92, 0x59, 0x76, 0x37, 0x85, 0xbe, 0x91, 0x2f, 0xe5, 0xc5, 0x27, 0x91, 0x4d, 0xda, 0xb8,
	0x82, 0xa2, 0x78, 0x9b, 0x19, 0xbe, 0xfd, 0x66, 0x32, 0x13, 0x32, 0x89, 0x13, 0xf3, 0x5c, 0xcc,
	0x99, 0xc0, 0x8c, 0x6b, 0x4c, 0xf1, 0x38, 0x41, 0x1e, 0xa7, 0x88, 0x5c, 0x2a, 0x7c, 0x01, 0x61,
	0x74, 0x95, 0x45, 0x32, 0xe1, 0xcb, 0x13, 0x2e, 0xd3, 0x22, 0x4e, 0x72, 0xcd, 0x17, 0xc5, 0x1c,
	0x54, 0x0e, 0x06, 0xdc, 0x90, 0x49, 0x85, 0x06, 0xe9, 0x81, 0x5b, 0xa9, 0x78, 0x66, 0x1d, 0xcc,
	0xea, 0x59, 0x82, 0x83, 0x7e, 0x8c, 0x31, 0x96, 0x2c, 0xb7, 0x51, 0xf5, 0x6c, 0x70, 0xf3, 0xaf,
	0x41, 0x34, 0xa8, 0x65, 0x22, 0xe0, 0x49, 0x4b, 0x10, 0x95, 0xe8, 0xf0, 0xad, 0x41, 0x3a, 0xf7,
	0x52, 0x1b, 0x05, 0x51, 0x36, 0x93, 0x20, 0xe8, 0x90, 0x74, 0x36, 0x58, 0x1e, 0x65, 0xe0, 0x7b,
	0x81, 0x17, 0xee, 0x4e, 0xdb, 0xeb, 0xda, 0x5d, 0x94, 0x01, 0x3d, 0x22, 0x7b, 0x2e, 0xa2, 0x65,
	0x24, 0xc0, 0x6f, 0x94, 0x5c, 0xcf, 0xe1, 0xca, 0xba, 0xeb, 0x93, 0xa8, 0x8c, 0xdf, 0x0c, 0xbc,
	0xb0, 0x5b, 0xfb, 0x26, 0xa8, 0x0c, 0x7d, 0x20, 0x3b, 0x1a, 0x52, 0x10, 0x06, 0x95, 0xbf, 0x15,
	0x34, 0xc3, 0xf6, 0x68, 0xcc, 0x7e, 0x59, 0x0b, 0x73, 0x67, 0x66, 0xb3, 0xf5, 0xeb, 0xeb, 0xdc,
	0xa8, 0xd5, 0xb4, 0x96, 0xd1, 0xab, 0xcf, 0xde, 0xf6, 0x93, 0xfd, 0x56, 0xe0, 0x85, 0xed, 0xd1,
	0xf0, 0x7b, 0xe3, 0xac, 0x22, 0xad, 0xb0, 0x1e, 0xcf, 0x26, 0x83, 0x31, 0xe9, 0x7e, 0x69, 0x40,
	0x7b, 0xa4, 0xb9, 0x80, 0xd5, 0x7a, 0x33, 0x36, 0xa4, 0x7d, 0xd2, 0x5a, 0x46, 0x69, 0xb1, 0xd9,
	0x42, 0x95, 0x9c, 0x37, 0xce, 0xbc, 0x8b, 0xdb, 0xd7, 0xf7, 0x7d, 0xef, 0xf1, 0xf2, 0x6f, 0xe7,
	0x92, 0x8b, 0xf8, 0xe7, 0x7f, 0x67, 0xbe, 0x5d, 0x5e, 0xec, 0xf4, 0x23, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0xec, 0xa4, 0x1d, 0x85, 0x02, 0x00, 0x00,
}
