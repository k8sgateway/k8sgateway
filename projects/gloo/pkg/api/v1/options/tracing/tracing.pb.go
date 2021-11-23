// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto

package tracing

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Contains settings for configuring Envoy's tracing capabilities at the listener level.
// See [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html) for additional information on Envoy's tracing capabilities.
// See [here](https://docs.solo.io/gloo-edge/latest/guides/observability/tracing/) for additional information about configuring tracing with Gloo Edge.
type ListenerTracingSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If specified, Envoy will include the headers and header values for any matching request headers.
	RequestHeadersForTags []string `protobuf:"bytes,1,rep,name=request_headers_for_tags,json=requestHeadersForTags,proto3" json:"request_headers_for_tags,omitempty"`
	// Optional. If true, Envoy will include logs for streaming events. Default: false.
	Verbose bool `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages *TracePercentages `protobuf:"bytes,3,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	// Optional. If not specified, no tracing will be performed
	// ProviderConfig defines the configuration for an external tracing provider.
	//
	// Types that are assignable to ProviderConfig:
	//	*ListenerTracingSettings_ZipkinConfig
	//	*ListenerTracingSettings_DatadogConfig
	ProviderConfig isListenerTracingSettings_ProviderConfig `protobuf_oneof:"provider_config"`
	// Optional. If specified, Envoy will include the environment variables with the given tag as tracing tags.
	EnvironmentVariablesForTags []*TracingTagEnvironmentVariable `protobuf:"bytes,6,rep,name=environment_variables_for_tags,json=environmentVariablesForTags,proto3" json:"environment_variables_for_tags,omitempty"`
	// Optional. If specified, Envoy will include the literals with the given tag as tracing tags.
	LiteralsForTags []*TracingTagLiteral `protobuf:"bytes,7,rep,name=literals_for_tags,json=literalsForTags,proto3" json:"literals_for_tags,omitempty"`
}

func (x *ListenerTracingSettings) Reset() {
	*x = ListenerTracingSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerTracingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerTracingSettings) ProtoMessage() {}

func (x *ListenerTracingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerTracingSettings.ProtoReflect.Descriptor instead.
func (*ListenerTracingSettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescGZIP(), []int{0}
}

func (x *ListenerTracingSettings) GetRequestHeadersForTags() []string {
	if x != nil {
		return x.RequestHeadersForTags
	}
	return nil
}

func (x *ListenerTracingSettings) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

func (x *ListenerTracingSettings) GetTracePercentages() *TracePercentages {
	if x != nil {
		return x.TracePercentages
	}
	return nil
}

func (m *ListenerTracingSettings) GetProviderConfig() isListenerTracingSettings_ProviderConfig {
	if m != nil {
		return m.ProviderConfig
	}
	return nil
}

func (x *ListenerTracingSettings) GetZipkinConfig() *v3.ZipkinConfig {
	if x, ok := x.GetProviderConfig().(*ListenerTracingSettings_ZipkinConfig); ok {
		return x.ZipkinConfig
	}
	return nil
}

func (x *ListenerTracingSettings) GetDatadogConfig() *v3.DatadogConfig {
	if x, ok := x.GetProviderConfig().(*ListenerTracingSettings_DatadogConfig); ok {
		return x.DatadogConfig
	}
	return nil
}

func (x *ListenerTracingSettings) GetEnvironmentVariablesForTags() []*TracingTagEnvironmentVariable {
	if x != nil {
		return x.EnvironmentVariablesForTags
	}
	return nil
}

func (x *ListenerTracingSettings) GetLiteralsForTags() []*TracingTagLiteral {
	if x != nil {
		return x.LiteralsForTags
	}
	return nil
}

type isListenerTracingSettings_ProviderConfig interface {
	isListenerTracingSettings_ProviderConfig()
}

type ListenerTracingSettings_ZipkinConfig struct {
	ZipkinConfig *v3.ZipkinConfig `protobuf:"bytes,4,opt,name=zipkin_config,json=zipkinConfig,proto3,oneof"`
}

type ListenerTracingSettings_DatadogConfig struct {
	DatadogConfig *v3.DatadogConfig `protobuf:"bytes,5,opt,name=datadog_config,json=datadogConfig,proto3,oneof"`
}

func (*ListenerTracingSettings_ZipkinConfig) isListenerTracingSettings_ProviderConfig() {}

func (*ListenerTracingSettings_DatadogConfig) isListenerTracingSettings_ProviderConfig() {}

// Contains settings for configuring Envoy's tracing capabilities at the route level.
// Note: must also specify ListenerTracingSettings for the associated listener.
// See [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/observability/tracing.html) for additional information on Envoy's tracing capabilities.
// See [here](https://docs.solo.io/gloo-edge/latest/guides/observability/tracing/) for additional information about configuring tracing with Gloo Edge.
type RouteTracingSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If set, will be used to identify the route that produced the trace.
	// Note that this value will be overridden if the "x-envoy-decorator-operation" header is passed.
	RouteDescriptor string `protobuf:"bytes,1,opt,name=route_descriptor,json=routeDescriptor,proto3" json:"route_descriptor,omitempty"`
	// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
	// TracePercentages defines the limits for random, forced, and overall tracing percentages.
	TracePercentages *TracePercentages `protobuf:"bytes,2,opt,name=trace_percentages,json=tracePercentages,proto3" json:"trace_percentages,omitempty"`
	// Optional. Default is true, If set to false, the tracing headers will not propagate to the upstream.
	Propagate *wrappers.BoolValue `protobuf:"bytes,3,opt,name=propagate,proto3" json:"propagate,omitempty"`
}

func (x *RouteTracingSettings) Reset() {
	*x = RouteTracingSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTracingSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTracingSettings) ProtoMessage() {}

func (x *RouteTracingSettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTracingSettings.ProtoReflect.Descriptor instead.
func (*RouteTracingSettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescGZIP(), []int{1}
}

func (x *RouteTracingSettings) GetRouteDescriptor() string {
	if x != nil {
		return x.RouteDescriptor
	}
	return ""
}

func (x *RouteTracingSettings) GetTracePercentages() *TracePercentages {
	if x != nil {
		return x.TracePercentages
	}
	return nil
}

func (x *RouteTracingSettings) GetPropagate() *wrappers.BoolValue {
	if x != nil {
		return x.Propagate
	}
	return nil
}

// Requests can produce traces by random sampling or when the `x-client-trace-id` header is provided.
// TracePercentages defines the limits for random, forced, and overall tracing percentages.
type TracePercentages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percentage of requests that should produce traces when the `x-client-trace-id` header is provided.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	ClientSamplePercentage *wrappers.FloatValue `protobuf:"bytes,1,opt,name=client_sample_percentage,json=clientSamplePercentage,proto3" json:"client_sample_percentage,omitempty"`
	// Percentage of requests that should produce traces by random sampling.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	RandomSamplePercentage *wrappers.FloatValue `protobuf:"bytes,2,opt,name=random_sample_percentage,json=randomSamplePercentage,proto3" json:"random_sample_percentage,omitempty"`
	// Overall percentage of requests that should produce traces.
	// optional, defaults to 100.0
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	OverallSamplePercentage *wrappers.FloatValue `protobuf:"bytes,3,opt,name=overall_sample_percentage,json=overallSamplePercentage,proto3" json:"overall_sample_percentage,omitempty"`
}

func (x *TracePercentages) Reset() {
	*x = TracePercentages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracePercentages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracePercentages) ProtoMessage() {}

func (x *TracePercentages) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracePercentages.ProtoReflect.Descriptor instead.
func (*TracePercentages) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescGZIP(), []int{2}
}

func (x *TracePercentages) GetClientSamplePercentage() *wrappers.FloatValue {
	if x != nil {
		return x.ClientSamplePercentage
	}
	return nil
}

func (x *TracePercentages) GetRandomSamplePercentage() *wrappers.FloatValue {
	if x != nil {
		return x.RandomSamplePercentage
	}
	return nil
}

func (x *TracePercentages) GetOverallSamplePercentage() *wrappers.FloatValue {
	if x != nil {
		return x.OverallSamplePercentage
	}
	return nil
}

// Requests can produce traces with custom tags.
// TracingTagEnvironmentVariable defines an environment variable which gets added as custom tag.
type TracingTagEnvironmentVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to populate the tag name.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Environment variable name to obtain the value to populate the tag value.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// When the environment variable is not found, the tag value will be populated with this default value if specified,
	// otherwise no tag will be populated.
	DefaultValue string `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (x *TracingTagEnvironmentVariable) Reset() {
	*x = TracingTagEnvironmentVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingTagEnvironmentVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingTagEnvironmentVariable) ProtoMessage() {}

func (x *TracingTagEnvironmentVariable) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingTagEnvironmentVariable.ProtoReflect.Descriptor instead.
func (*TracingTagEnvironmentVariable) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescGZIP(), []int{3}
}

func (x *TracingTagEnvironmentVariable) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *TracingTagEnvironmentVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TracingTagEnvironmentVariable) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

// Requests can produce traces with custom tags.
// TracingTagLiteral defines a literal which gets added as custom tag.
type TracingTagLiteral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to populate the tag name.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Static literal value to populate the tag value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TracingTagLiteral) Reset() {
	*x = TracingTagLiteral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracingTagLiteral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracingTagLiteral) ProtoMessage() {}

func (x *TracingTagLiteral) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracingTagLiteral.ProtoReflect.Descriptor instead.
func (*TracingTagLiteral) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescGZIP(), []int{4}
}

func (x *TracingTagLiteral) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *TracingTagLiteral) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x04, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x10, 0x74, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x52, 0x0a, 0x0d, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x5a, 0x69, 0x70, 0x6b, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0d,
	0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x80, 0x01,
	0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x1b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x5b, 0x0a, 0x11, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x5f, 0x66, 0x6f, 0x72,
	0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x0f, 0x6c, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x61, 0x67, 0x73, 0x42, 0x11, 0x0a,
	0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0xd8, 0x01, 0x0a, 0x14, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x10, 0x74, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x55, 0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x55, 0x0a, 0x18, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x16, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x57,
	0x0a, 0x19, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17,
	0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x22, 0x6a, 0x0a, 0x1d, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x67, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x67, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x4a, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_goTypes = []interface{}{
	(*ListenerTracingSettings)(nil),       // 0: tracing.options.gloo.solo.io.ListenerTracingSettings
	(*RouteTracingSettings)(nil),          // 1: tracing.options.gloo.solo.io.RouteTracingSettings
	(*TracePercentages)(nil),              // 2: tracing.options.gloo.solo.io.TracePercentages
	(*TracingTagEnvironmentVariable)(nil), // 3: tracing.options.gloo.solo.io.TracingTagEnvironmentVariable
	(*TracingTagLiteral)(nil),             // 4: tracing.options.gloo.solo.io.TracingTagLiteral
	(*v3.ZipkinConfig)(nil),               // 5: solo.io.envoy.config.trace.v3.ZipkinConfig
	(*v3.DatadogConfig)(nil),              // 6: solo.io.envoy.config.trace.v3.DatadogConfig
	(*wrappers.BoolValue)(nil),            // 7: google.protobuf.BoolValue
	(*wrappers.FloatValue)(nil),           // 8: google.protobuf.FloatValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_depIdxs = []int32{
	2,  // 0: tracing.options.gloo.solo.io.ListenerTracingSettings.trace_percentages:type_name -> tracing.options.gloo.solo.io.TracePercentages
	5,  // 1: tracing.options.gloo.solo.io.ListenerTracingSettings.zipkin_config:type_name -> solo.io.envoy.config.trace.v3.ZipkinConfig
	6,  // 2: tracing.options.gloo.solo.io.ListenerTracingSettings.datadog_config:type_name -> solo.io.envoy.config.trace.v3.DatadogConfig
	3,  // 3: tracing.options.gloo.solo.io.ListenerTracingSettings.environment_variables_for_tags:type_name -> tracing.options.gloo.solo.io.TracingTagEnvironmentVariable
	4,  // 4: tracing.options.gloo.solo.io.ListenerTracingSettings.literals_for_tags:type_name -> tracing.options.gloo.solo.io.TracingTagLiteral
	2,  // 5: tracing.options.gloo.solo.io.RouteTracingSettings.trace_percentages:type_name -> tracing.options.gloo.solo.io.TracePercentages
	7,  // 6: tracing.options.gloo.solo.io.RouteTracingSettings.propagate:type_name -> google.protobuf.BoolValue
	8,  // 7: tracing.options.gloo.solo.io.TracePercentages.client_sample_percentage:type_name -> google.protobuf.FloatValue
	8,  // 8: tracing.options.gloo.solo.io.TracePercentages.random_sample_percentage:type_name -> google.protobuf.FloatValue
	8,  // 9: tracing.options.gloo.solo.io.TracePercentages.overall_sample_percentage:type_name -> google.protobuf.FloatValue
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerTracingSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTracingSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracePercentages); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingTagEnvironmentVariable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracingTagLiteral); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ListenerTracingSettings_ZipkinConfig)(nil),
		(*ListenerTracingSettings_DatadogConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tracing_tracing_proto_depIdxs = nil
}
