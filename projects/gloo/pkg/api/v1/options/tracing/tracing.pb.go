// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto

package tracing

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Contains settings for configuring Envoy's tracing capabilities at the listener level.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type ListenerTracingSettings struct {
	// Optional. If specified, Envoy will include the headers and header values for any matching request headers.
	RequestHeadersForTags []string `protobuf:"bytes,1,rep,name=request_headers_for_tags,json=requestHeadersForTags,proto3" json:"request_headers_for_tags,omitempty"`
	// Optional. If true, Envoy will include logs for streaming events. Default: false.
	Verbose bool `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages     *TracePercentages `protobuf:"bytes,3,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListenerTracingSettings) Reset()         { *m = ListenerTracingSettings{} }
func (m *ListenerTracingSettings) String() string { return proto.CompactTextString(m) }
func (*ListenerTracingSettings) ProtoMessage()    {}
func (*ListenerTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f134b8947c6e68, []int{0}
}
func (m *ListenerTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerTracingSettings.Unmarshal(m, b)
}
func (m *ListenerTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerTracingSettings.Marshal(b, m, deterministic)
}
func (m *ListenerTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerTracingSettings.Merge(m, src)
}
func (m *ListenerTracingSettings) XXX_Size() int {
	return xxx_messageInfo_ListenerTracingSettings.Size(m)
}
func (m *ListenerTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerTracingSettings proto.InternalMessageInfo

func (m *ListenerTracingSettings) GetRequestHeadersForTags() []string {
	if m != nil {
		return m.RequestHeadersForTags
	}
	return nil
}

func (m *ListenerTracingSettings) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *ListenerTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

// Contains settings for configuring Envoy's tracing capabilities at the route level.
// Note: must also specify ListenerTracingSettings for the associated listener.
// See here for additional information on Envoy's tracing capabilities: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html
// See here for additional information about configuring tracing with Gloo: https://gloo.solo.io/user_guides/setup_options/observability/#tracing
type RouteTracingSettings struct {
	// Optional. If set, will be used to identify the route that produced the trace.
	// Note that this value will be overridden if the "x-envoy-decorator-operation" header is passed.
	RouteDescriptor string `protobuf:"bytes,1,opt,name=route_descriptor,json=routeDescriptor,proto3" json:"route_descriptor,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages     *TracePercentages `protobuf:"bytes,2,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RouteTracingSettings) Reset()         { *m = RouteTracingSettings{} }
func (m *RouteTracingSettings) String() string { return proto.CompactTextString(m) }
func (*RouteTracingSettings) ProtoMessage()    {}
func (*RouteTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f134b8947c6e68, []int{1}
}
func (m *RouteTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTracingSettings.Unmarshal(m, b)
}
func (m *RouteTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTracingSettings.Marshal(b, m, deterministic)
}
func (m *RouteTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTracingSettings.Merge(m, src)
}
func (m *RouteTracingSettings) XXX_Size() int {
	return xxx_messageInfo_RouteTracingSettings.Size(m)
}
func (m *RouteTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTracingSettings proto.InternalMessageInfo

func (m *RouteTracingSettings) GetRouteDescriptor() string {
	if m != nil {
		return m.RouteDescriptor
	}
	return ""
}

func (m *RouteTracingSettings) GetTracePercentages() *TracePercentages {
	if m != nil {
		return m.TracePercentages
	}
	return nil
}

// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
// TracePercentages defines the limits for random, forced, and overall tracing percentages.
type TracePercentages struct {
	// Percentage of requests that should produce traces when the `x-client-trace-id` header is provided.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	ClientSamplePercentage *types.FloatValue `protobuf:"bytes,1,opt,name=client_sample_percentage,json=clientSamplePercentage,proto3" json:"client_sample_percentage,omitempty"`
	// Percentage of requests that should produce traces by random sampling.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	RandomSamplePercentage *types.FloatValue `protobuf:"bytes,2,opt,name=random_sample_percentage,json=randomSamplePercentage,proto3" json:"random_sample_percentage,omitempty"`
	// Overall percentage of requests that should produce traces.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	OverallSamplePercentage *types.FloatValue `protobuf:"bytes,3,opt,name=overall_sample_percentage,json=overallSamplePercentage,proto3" json:"overall_sample_percentage,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *TracePercentages) Reset()         { *m = TracePercentages{} }
func (m *TracePercentages) String() string { return proto.CompactTextString(m) }
func (*TracePercentages) ProtoMessage()    {}
func (*TracePercentages) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f134b8947c6e68, []int{2}
}
func (m *TracePercentages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracePercentages.Unmarshal(m, b)
}
func (m *TracePercentages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracePercentages.Marshal(b, m, deterministic)
}
func (m *TracePercentages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracePercentages.Merge(m, src)
}
func (m *TracePercentages) XXX_Size() int {
	return xxx_messageInfo_TracePercentages.Size(m)
}
func (m *TracePercentages) XXX_DiscardUnknown() {
	xxx_messageInfo_TracePercentages.DiscardUnknown(m)
}

var xxx_messageInfo_TracePercentages proto.InternalMessageInfo

func (m *TracePercentages) GetClientSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.ClientSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetRandomSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.RandomSamplePercentage
	}
	return nil
}

func (m *TracePercentages) GetOverallSamplePercentage() *types.FloatValue {
	if m != nil {
		return m.OverallSamplePercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenerTracingSettings)(nil), "tracing.options.gloo.solo.io.ListenerTracingSettings")
	proto.RegisterType((*RouteTracingSettings)(nil), "tracing.options.gloo.solo.io.RouteTracingSettings")
	proto.RegisterType((*TracePercentages)(nil), "tracing.options.gloo.solo.io.TracePercentages")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto", fileDescriptor_30f134b8947c6e68)
}

var fileDescriptor_30f134b8947c6e68 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xe5, 0x04, 0x01, 0x75, 0x0f, 0x84, 0x55, 0xa1, 0x4b, 0x41, 0x55, 0xd4, 0x53, 0x38,
	0x60, 0x8b, 0x72, 0xe0, 0x8a, 0x10, 0xaa, 0x50, 0xc5, 0x01, 0x6d, 0x0b, 0x48, 0x70, 0x58, 0x39,
	0x9b, 0xa9, 0x6b, 0x70, 0x76, 0xcc, 0x78, 0x36, 0xf4, 0x55, 0xb8, 0x70, 0xe6, 0x11, 0x78, 0x05,
	0x5e, 0x83, 0x77, 0xe0, 0x8e, 0x76, 0x9d, 0x50, 0x14, 0x52, 0xc8, 0xa1, 0xa7, 0xdd, 0xf9, 0x67,
	0xfe, 0xcf, 0xbf, 0x6d, 0x59, 0x1e, 0x5a, 0xc7, 0xa7, 0xcd, 0x58, 0x55, 0x38, 0xd5, 0x11, 0x3d,
	0x3e, 0x70, 0xa8, 0xad, 0x47, 0xd4, 0x81, 0xf0, 0x3d, 0x54, 0x1c, 0x53, 0x65, 0x82, 0xd3, 0xb3,
	0x87, 0x1a, 0x03, 0x3b, 0xac, 0xa3, 0x66, 0x32, 0x95, 0xab, 0xed, 0xe2, 0xab, 0x02, 0x21, 0x63,
	0x76, 0x6f, 0x51, 0xce, 0xc7, 0x54, 0x6b, 0x55, 0x2d, 0x55, 0x39, 0xdc, 0xd9, 0xb2, 0x68, 0xb1,
	0x1b, 0xd4, 0xed, 0x5f, 0xf2, 0xec, 0xec, 0x5a, 0x44, 0xeb, 0x41, 0x77, 0xd5, 0xb8, 0x39, 0xd1,
	0x9f, 0xc8, 0x84, 0x00, 0x14, 0x2f, 0xea, 0x4f, 0x1a, 0x32, 0x2d, 0x7d, 0xde, 0xcf, 0xe0, 0x8c,
	0x13, 0x14, 0xce, 0x38, 0x69, 0x7b, 0xdf, 0x85, 0xdc, 0x7e, 0xe1, 0x22, 0x43, 0x0d, 0x74, 0x9c,
	0x22, 0x1d, 0x01, 0xb3, 0xab, 0x6d, 0xcc, 0x1e, 0xcb, 0x9c, 0xe0, 0x63, 0x03, 0x91, 0xcb, 0x53,
	0x30, 0x13, 0xa0, 0x58, 0x9e, 0x20, 0x95, 0x6c, 0x6c, 0xcc, 0xc5, 0xb0, 0x3f, 0xda, 0x28, 0x6e,
	0xcd, 0xfb, 0xcf, 0x53, 0xfb, 0x00, 0xe9, 0xd8, 0xd8, 0x98, 0xe5, 0xf2, 0xda, 0x0c, 0x68, 0x8c,
	0x11, 0xf2, 0xde, 0x50, 0x8c, 0xae, 0x17, 0x8b, 0x32, 0x7b, 0x27, 0x6f, 0xb6, 0x1b, 0x87, 0x32,
	0x00, 0x55, 0x50, 0xb3, 0xb1, 0x10, 0xf3, 0xfe, 0x50, 0x8c, 0x36, 0xf7, 0x95, 0xfa, 0xd7, 0x91,
	0xa8, 0x36, 0x1c, 0xbc, 0x3c, 0x77, 0x15, 0x03, 0x5e, 0x52, 0xf6, 0xbe, 0x08, 0xb9, 0x55, 0x60,
	0xc3, 0xb0, 0xbc, 0x91, 0xfb, 0x72, 0x40, 0xad, 0x5e, 0x4e, 0x20, 0x56, 0xe4, 0x02, 0x23, 0xe5,
	0x62, 0x28, 0x46, 0x1b, 0xc5, 0x8d, 0x4e, 0x7f, 0xf6, 0x5b, 0x5e, 0x1d, 0xb0, 0x77, 0x49, 0x01,
	0x3f, 0xf7, 0xe4, 0x60, 0x79, 0x2c, 0x7b, 0x25, 0xf3, 0xca, 0x3b, 0xa8, 0xb9, 0x8c, 0x66, 0x1a,
	0xfc, 0x9f, 0x2b, 0x77, 0x21, 0x37, 0xf7, 0xef, 0xaa, 0x74, 0xb1, 0x6a, 0x71, 0xb1, 0xea, 0xc0,
	0xa3, 0xe1, 0xd7, 0xc6, 0x37, 0x50, 0xdc, 0x4e, 0xe6, 0xa3, 0xce, 0x7b, 0xce, 0x6d, 0xb1, 0x64,
	0xea, 0x09, 0x4e, 0x57, 0x60, 0x7b, 0x6b, 0x60, 0x93, 0xf9, 0x2f, 0xec, 0x1b, 0x79, 0x07, 0x67,
	0x40, 0xc6, 0xfb, 0x15, 0xdc, 0xfe, 0xff, 0xb9, 0xdb, 0x73, 0xf7, 0x32, 0xf8, 0xe9, 0xe1, 0xb7,
	0x9f, 0x57, 0xc4, 0xd7, 0x1f, 0xbb, 0xe2, 0xed, 0x93, 0xf5, 0x9e, 0x59, 0xf8, 0x60, 0x2f, 0x78,
	0x6a, 0xe3, 0xab, 0xdd, 0xca, 0x8f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x45, 0x09, 0x61, 0x80,
	0xb1, 0x03, 0x00, 0x00,
}

func (this *ListenerTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerTracingSettings)
	if !ok {
		that2, ok := that.(ListenerTracingSettings)
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
	if len(this.RequestHeadersForTags) != len(that1.RequestHeadersForTags) {
		return false
	}
	for i := range this.RequestHeadersForTags {
		if this.RequestHeadersForTags[i] != that1.RequestHeadersForTags[i] {
			return false
		}
	}
	if this.Verbose != that1.Verbose {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTracingSettings)
	if !ok {
		that2, ok := that.(RouteTracingSettings)
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
	if this.RouteDescriptor != that1.RouteDescriptor {
		return false
	}
	if !this.TracePercentages.Equal(that1.TracePercentages) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TracePercentages) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TracePercentages)
	if !ok {
		that2, ok := that.(TracePercentages)
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
	if !this.ClientSamplePercentage.Equal(that1.ClientSamplePercentage) {
		return false
	}
	if !this.RandomSamplePercentage.Equal(that1.RandomSamplePercentage) {
		return false
	}
	if !this.OverallSamplePercentage.Equal(that1.OverallSamplePercentage) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
