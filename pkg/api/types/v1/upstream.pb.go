// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upstream.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// *
// Upstream represents a destination for routing. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v1/cluster_manager/cluster.html?highlight=cluster) in Envoy terminology.
// Upstreams can take a variety of types<!--(TODO)--> in gloo. Language extensions known as plugins<!--(TODO)--> allow the addition of new
// types of upstreams. <!--See [upstream types](TODO) for a detailed description of available upstream types.-->
type Upstream struct {
	// Name of the upstream. Names must be unique and follow the following syntax rules:
	// One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type indicates the type of the upstream. Examples include service<!--(TODO)-->, kubernetes<!--(TODO)-->, and [aws](../plugins/aws.md)
	// Types are defined by the plugin<!--(TODO)--> associated with them.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Connection Timeout tells gloo to set a timeout for unresponsive connections created to this upstream.
	// If not provided by the user, it will set to a default value
	ConnectionTimeout time.Duration `protobuf:"bytes,3,opt,name=connection_timeout,json=connectionTimeout,stdduration" json:"connection_timeout"`
	// Spec contains properties that are specific to the upstream type. The spec is always required, but
	// the expected content is specified by the [upstream plugin] for the given upstream type.
	// Most often the upstream spec will be a map<string, string>
	Spec *google_protobuf.Struct `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	// Certain upstream types support (and may require) [functions](../introduction/concepts.md#Functions).
	// Functions allow function-level routing to be done. For example, the [AWS lambda](../plugins/aws.md) upstream type
	// Permits routing to AWS lambda function].
	// [routes](virtualservice.md#Route) on virtualservices can specify function destinations to route to specific functions.
	Functions []*Function `protobuf:"bytes,5,rep,name=functions" json:"functions,omitempty"`
	// Service Info contains information about the service running on the upstream
	// Service Info is optional, but is used by certain plugins (such as the gRPC plugin)
	// as well as discovery services to provide sophistocated routing features for well-known
	// types of services
	ServiceInfo *ServiceInfo `protobuf:"bytes,8,opt,name=service_info,json=serviceInfo" json:"service_info,omitempty"`
	// Status indicates the validation status of the upstream resource. Status is read-only by clients, and set by gloo during validation
	Status *Status `protobuf:"bytes,6,opt,name=status" json:"status,omitempty" testdiff:"ignore"`
	// Metadata contains the resource metadata for the upstream
	Metadata *Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Upstream) Reset()                    { *m = Upstream{} }
func (m *Upstream) String() string            { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()               {}
func (*Upstream) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{0} }

func (m *Upstream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Upstream) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Upstream) GetConnectionTimeout() time.Duration {
	if m != nil {
		return m.ConnectionTimeout
	}
	return 0
}

func (m *Upstream) GetSpec() *google_protobuf.Struct {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Upstream) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *Upstream) GetServiceInfo() *ServiceInfo {
	if m != nil {
		return m.ServiceInfo
	}
	return nil
}

func (m *Upstream) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Upstream) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ServiceInfo struct {
	// Type indicates the type of service running on the upstream.
	// Current options include `REST`, `gRPC`, and `NATS`
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Properties contains properties that describe the service. The spec may be required
	// by the Upstream Plugin that handles the given Service Type
	// Most often the service properties will be a map<string, string>
	Properties *google_protobuf.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *ServiceInfo) Reset()                    { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string            { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()               {}
func (*ServiceInfo) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{1} }

func (m *ServiceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ServiceInfo) GetProperties() *google_protobuf.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Function struct {
	// Name of the function. Functions are referenced by name from routes and therefore must be unique within an upstream
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Spec for the function. Like [upstream specs](TODO), the content of function specs is specified by the [upstream plugin](TODO) for the upstream's type.
	Spec *google_protobuf.Struct `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
}

func (m *Function) Reset()                    { *m = Function{} }
func (m *Function) String() string            { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()               {}
func (*Function) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{2} }

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetSpec() *google_protobuf.Struct {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*Upstream)(nil), "v1.Upstream")
	proto.RegisterType((*ServiceInfo)(nil), "v1.ServiceInfo")
	proto.RegisterType((*Function)(nil), "v1.Function")
}
func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.ConnectionTimeout != that1.ConnectionTimeout {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if len(this.Functions) != len(that1.Functions) {
		return false
	}
	for i := range this.Functions {
		if !this.Functions[i].Equal(that1.Functions[i]) {
			return false
		}
	}
	if !this.ServiceInfo.Equal(that1.ServiceInfo) {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}
func (this *ServiceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceInfo)
	if !ok {
		that2, ok := that.(ServiceInfo)
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
	if this.Type != that1.Type {
		return false
	}
	if !this.Properties.Equal(that1.Properties) {
		return false
	}
	return true
}
func (this *Function) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Function)
	if !ok {
		that2, ok := that.(Function)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("upstream.proto", fileDescriptorUpstream) }

var fileDescriptorUpstream = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xbe, 0x69, 0x73, 0x7b, 0xd3, 0x49, 0xe9, 0xa5, 0xc3, 0xbd, 0xdc, 0xb9, 0x45, 0xda, 0x92,
	0x55, 0x50, 0x48, 0x69, 0x5d, 0x88, 0x2e, 0x8b, 0x08, 0x22, 0x6e, 0xa6, 0xba, 0x71, 0x53, 0xd2,
	0x74, 0x12, 0x02, 0x66, 0x26, 0xcc, 0x4c, 0x0a, 0xbe, 0x89, 0x8f, 0xe0, 0xa3, 0xf8, 0x14, 0x15,
	0x7c, 0x04, 0x57, 0x2e, 0x25, 0x27, 0xd3, 0x1f, 0x54, 0x04, 0x77, 0xe7, 0x7c, 0x3f, 0x33, 0x1f,
	0xdf, 0x41, 0xed, 0x22, 0x57, 0x5a, 0xb2, 0x30, 0x0b, 0x72, 0x29, 0xb4, 0xc0, 0xb5, 0xe5, 0xa8,
	0xbb, 0x97, 0x08, 0x91, 0xdc, 0xb2, 0x21, 0x20, 0xf3, 0x22, 0x1e, 0x2a, 0x2d, 0x8b, 0x48, 0x57,
	0x8a, 0x6e, 0xef, 0x3d, 0xbb, 0x28, 0x64, 0xa8, 0x53, 0xc1, 0x0d, 0xff, 0x27, 0x11, 0x89, 0x80,
	0x71, 0x58, 0x4e, 0x06, 0x6d, 0x29, 0x1d, 0xea, 0x42, 0x99, 0xad, 0x9d, 0x31, 0x1d, 0x2e, 0x42,
	0x1d, 0x56, 0xbb, 0xf7, 0x5a, 0x43, 0xce, 0xb5, 0x09, 0x82, 0x31, 0xb2, 0x79, 0x98, 0x31, 0x62,
	0x0d, 0x2c, 0xbf, 0x49, 0x61, 0x2e, 0x31, 0x7d, 0x97, 0x33, 0x52, 0xab, 0xb0, 0x72, 0xc6, 0x14,
	0xe1, 0x48, 0x70, 0xce, 0xa2, 0xf2, 0xf3, 0x99, 0x4e, 0x33, 0x26, 0x0a, 0x4d, 0xea, 0x03, 0xcb,
	0x77, 0xc7, 0xff, 0x83, 0x2a, 0x65, 0xb0, 0x4e, 0x19, 0x9c, 0x9a, 0x94, 0x13, 0xe7, 0x71, 0xd5,
	0xff, 0x71, 0xff, 0xd4, 0xb7, 0x68, 0x67, 0x6b, 0xbf, 0xaa, 0xdc, 0xf8, 0x00, 0xd9, 0x2a, 0x67,
	0x11, 0xb1, 0xe1, 0x95, 0x7f, 0x1f, 0x5e, 0x99, 0x42, 0x13, 0x14, 0x44, 0x78, 0x1f, 0x35, 0xe3,
	0x82, 0x83, 0x5f, 0x91, 0x9f, 0x83, 0xba, 0xef, 0x8e, 0x5b, 0xc1, 0x72, 0x14, 0x9c, 0x19, 0x90,
	0x6e, 0x69, 0x3c, 0x46, 0x2d, 0xc5, 0xe4, 0x32, 0x8d, 0xd8, 0x2c, 0xe5, 0xb1, 0x20, 0x0e, 0x7c,
	0xf0, 0xbb, 0x94, 0x4f, 0x2b, 0xfc, 0x9c, 0xc7, 0x82, 0xba, 0x6a, 0xbb, 0xe0, 0x63, 0xd4, 0xa8,
	0x5a, 0x23, 0x0d, 0x50, 0x23, 0x50, 0x03, 0x32, 0xf9, 0xfb, 0xb2, 0xea, 0x77, 0x34, 0x53, 0x7a,
	0x91, 0xc6, 0xf1, 0x89, 0x97, 0x26, 0x5c, 0x48, 0xe6, 0x51, 0x63, 0xc0, 0x3e, 0x72, 0xd6, 0x15,
	0x93, 0x5f, 0x60, 0x86, 0x64, 0x97, 0x06, 0xa3, 0x1b, 0xd6, 0xbb, 0x41, 0xee, 0x4e, 0x80, 0x4d,
	0xd1, 0xd6, 0x4e, 0xd1, 0x47, 0x08, 0xe5, 0x52, 0xe4, 0x4c, 0xea, 0x94, 0x29, 0x38, 0xc1, 0x17,
	0xd5, 0xec, 0x48, 0xbd, 0x0b, 0xe4, 0xac, 0xbb, 0xf8, 0xf4, 0xaa, 0xdf, 0x69, 0x7b, 0x62, 0x3f,
	0x3c, 0xf7, 0xac, 0x79, 0x03, 0xc8, 0xc3, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x65, 0x90,
	0xd3, 0xb8, 0x02, 0x00, 0x00,
}
