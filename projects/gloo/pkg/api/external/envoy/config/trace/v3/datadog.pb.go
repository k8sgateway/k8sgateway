// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/datadog.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	math "math"
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

// Configuration for the Datadog tracer.
// [#extension: envoy.tracers.datadog]
type DatadogConfig struct {
	// The upstream to use for submitting traces to the Datadog agent.
	CollectorUpstreamRef *core.ResourceRef `protobuf:"bytes,1,opt,name=collector_upstream_ref,json=collectorUpstreamRef,proto3" json:"collector_upstream_ref,omitempty"`
	// The name used for the service when traces are generated by envoy.
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatadogConfig) Reset()         { *m = DatadogConfig{} }
func (m *DatadogConfig) String() string { return proto.CompactTextString(m) }
func (*DatadogConfig) ProtoMessage()    {}
func (*DatadogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f692eada442da, []int{0}
}
func (m *DatadogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatadogConfig.Unmarshal(m, b)
}
func (m *DatadogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatadogConfig.Marshal(b, m, deterministic)
}
func (m *DatadogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatadogConfig.Merge(m, src)
}
func (m *DatadogConfig) XXX_Size() int {
	return xxx_messageInfo_DatadogConfig.Size(m)
}
func (m *DatadogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatadogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatadogConfig proto.InternalMessageInfo

func (m *DatadogConfig) GetCollectorUpstreamRef() *core.ResourceRef {
	if m != nil {
		return m.CollectorUpstreamRef
	}
	return nil
}

func (m *DatadogConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func init() {
	proto.RegisterType((*DatadogConfig)(nil), "envoy.config.trace.v3.DatadogConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/datadog.proto", fileDescriptor_540f692eada442da)
}

var fileDescriptor_540f692eada442da = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0x14, 0x41,
	0x10, 0x86, 0x99, 0x45, 0x23, 0x4e, 0x22, 0x84, 0x25, 0x6a, 0xb2, 0xe0, 0x12, 0x0d, 0xc2, 0xb2,
	0x60, 0x37, 0x66, 0x3d, 0x79, 0x5c, 0x3d, 0x8a, 0x86, 0x01, 0x3d, 0x78, 0x59, 0x2a, 0x3d, 0x35,
	0x9d, 0x36, 0x33, 0x5d, 0x4d, 0x75, 0x4f, 0x93, 0xdc, 0x3c, 0x7a, 0xf3, 0x2a, 0x3e, 0x81, 0x8f,
	0x20, 0x5e, 0x3c, 0x09, 0x5e, 0x7d, 0x05, 0x1f, 0xc1, 0xa3, 0x07, 0x91, 0xe9, 0x99, 0x08, 0x92,
	0x1c, 0x72, 0xab, 0x99, 0xbf, 0x3e, 0xaa, 0xff, 0xfa, 0x2b, 0x7f, 0xa5, 0x4d, 0x38, 0x6a, 0x0f,
	0x85, 0xa2, 0x46, 0x7a, 0xaa, 0xe9, 0x81, 0x21, 0xa9, 0x6b, 0x22, 0xe9, 0x98, 0xde, 0xa0, 0x0a,
	0xbe, 0xff, 0x02, 0x67, 0x24, 0x9e, 0x04, 0x64, 0x0b, 0xb5, 0x44, 0x1b, 0xe9, 0x54, 0x2a, 0xb2,
	0x95, 0xd1, 0x32, 0x30, 0x28, 0x94, 0x71, 0x21, 0x4b, 0x08, 0x50, 0x92, 0x16, 0x8e, 0x29, 0xd0,
	0xf8, 0x66, 0x6a, 0x12, 0x7d, 0x93, 0x48, 0x4d, 0x22, 0x2e, 0x26, 0xd3, 0xb6, 0x74, 0x20, 0xc1,
	0x5a, 0x0a, 0x10, 0x0c, 0x59, 0x2f, 0x1b, 0xa3, 0x19, 0x02, 0xf6, 0xd8, 0xe4, 0xce, 0x39, 0xdd,
	0x07, 0x08, 0xad, 0x1f, 0xe4, 0xbb, 0xe7, 0xe4, 0x88, 0xec, 0x0d, 0x59, 0x63, 0x87, 0xc1, 0x93,
	0xdb, 0x11, 0x6a, 0x53, 0x42, 0x40, 0x79, 0x56, 0x0c, 0xc2, 0x4e, 0xb2, 0x77, 0x6c, 0x42, 0x32,
	0x13, 0x1f, 0x4a, 0xc6, 0x6a, 0x90, 0xb6, 0x34, 0x69, 0x4a, 0xa5, 0xec, 0xaa, 0xfe, 0xef, 0xbd,
	0xcf, 0x59, 0x7e, 0xe3, 0x69, 0x6f, 0xea, 0x49, 0xb2, 0x31, 0x7e, 0x91, 0xdf, 0x52, 0x54, 0xd7,
	0xa8, 0x02, 0xf1, 0xaa, 0x75, 0x3e, 0x30, 0x42, 0xb3, 0x62, 0xac, 0xb6, 0xb3, 0xdd, 0x6c, 0xb6,
	0xbe, 0xbf, 0x23, 0x14, 0x31, 0x8a, 0x6e, 0x90, 0x30, 0x24, 0x0a, 0xf4, 0xd4, 0xb2, 0xc2, 0x02,
	0xab, 0x62, 0xeb, 0x1f, 0xf8, 0x72, 0xe0, 0x0a, 0xac, 0xc6, 0xf3, 0x7c, 0xc3, 0x23, 0x47, 0xa3,
	0x70, 0x65, 0xa1, 0xc1, 0xed, 0xd1, 0x6e, 0x36, 0xbb, 0xbe, 0xbc, 0xf6, 0x7b, 0x79, 0x85, 0x47,
	0x9b, 0x59, 0xb1, 0x3e, 0x88, 0xcf, 0xa1, 0xc1, 0xc7, 0xf3, 0x8f, 0xdf, 0xde, 0x4d, 0xef, 0xe7,
	0x7b, 0x17, 0x2d, 0x76, 0x5f, 0xfc, 0xf7, 0xd0, 0xe5, 0xd7, 0xec, 0xd3, 0xcf, 0x69, 0xf6, 0xeb,
	0xc3, 0x9f, 0xf7, 0x57, 0xe7, 0xe3, 0x59, 0x4f, 0x74, 0xf1, 0xd9, 0x6e, 0x57, 0xbe, 0xa7, 0xd8,
	0x8b, 0xb3, 0xc8, 0xe2, 0x23, 0xa8, 0xdd, 0x11, 0x7c, 0x79, 0xfb, 0xfd, 0xc7, 0xda, 0x68, 0x73,
	0x94, 0xef, 0x19, 0x12, 0x09, 0x72, 0x4c, 0x27, 0xa7, 0xe2, 0xc2, 0x28, 0x97, 0x1b, 0xc3, 0xc8,
	0x83, 0x6e, 0x59, 0x07, 0xd9, 0xeb, 0x67, 0x97, 0xbb, 0x24, 0x77, 0xac, 0x2f, 0x71, 0x4d, 0x87,
	0x6b, 0x29, 0x83, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xb9, 0xfc, 0x0e, 0xa0, 0x02,
	0x00, 0x00,
}

func (this *DatadogConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DatadogConfig)
	if !ok {
		that2, ok := that.(DatadogConfig)
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
	if !this.CollectorUpstreamRef.Equal(that1.CollectorUpstreamRef) {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
