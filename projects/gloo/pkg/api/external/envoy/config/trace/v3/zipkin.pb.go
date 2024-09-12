// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/zipkin.proto

package v3

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/annotations"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Available Zipkin collector endpoint versions.
type ZipkinConfig_CollectorEndpointVersion int32

const (
	// Zipkin API v1, JSON over HTTP.
	// [#comment: The default implementation of Zipkin client before this field is added was only v1
	// and the way user configure this was by not explicitly specifying the version. Consequently,
	// before this is added, the corresponding Zipkin collector expected to receive v1 payload.
	// Hence the motivation of adding HTTP_JSON_V1 as the default is to avoid a breaking change when
	// user upgrading Envoy with this change. Furthermore, we also immediately deprecate this field,
	// since in Zipkin realm this v1 version is considered to be not preferable anymore.]
	//
	// Deprecated: Marked as deprecated in github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/zipkin.proto.
	ZipkinConfig_DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE ZipkinConfig_CollectorEndpointVersion = 0
	// Zipkin API v2, JSON over HTTP.
	ZipkinConfig_HTTP_JSON ZipkinConfig_CollectorEndpointVersion = 1
	// Zipkin API v2, protobuf over HTTP.
	ZipkinConfig_HTTP_PROTO ZipkinConfig_CollectorEndpointVersion = 2
)

// Enum value maps for ZipkinConfig_CollectorEndpointVersion.
var (
	ZipkinConfig_CollectorEndpointVersion_name = map[int32]string{
		0: "DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE",
		1: "HTTP_JSON",
		2: "HTTP_PROTO",
	}
	ZipkinConfig_CollectorEndpointVersion_value = map[string]int32{
		"DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE": 0,
		"HTTP_JSON":                             1,
		"HTTP_PROTO":                            2,
	}
)

func (x ZipkinConfig_CollectorEndpointVersion) Enum() *ZipkinConfig_CollectorEndpointVersion {
	p := new(ZipkinConfig_CollectorEndpointVersion)
	*p = x
	return p
}

func (x ZipkinConfig_CollectorEndpointVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZipkinConfig_CollectorEndpointVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_enumTypes[0].Descriptor()
}

func (ZipkinConfig_CollectorEndpointVersion) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_enumTypes[0]
}

func (x ZipkinConfig_CollectorEndpointVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZipkinConfig_CollectorEndpointVersion.Descriptor instead.
func (ZipkinConfig_CollectorEndpointVersion) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescGZIP(), []int{0, 0}
}

// Configuration for the Zipkin tracer.
// [#extension: envoy.tracers.zipkin]
// [#next-free-field: 6]
type ZipkinConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cluster that hosts the Zipkin collectors.
	//
	// Types that are assignable to CollectorCluster:
	//
	//	*ZipkinConfig_CollectorUpstreamRef
	//	*ZipkinConfig_ClusterName
	CollectorCluster isZipkinConfig_CollectorCluster `protobuf_oneof:"collector_cluster"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	// Determines whether a 128bit trace id will be used when creating a new
	// trace instance. The default value is false, which will result in a 64 bit trace id being used.
	TraceId_128Bit *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
	// Determines whether client and server spans will share the same span context.
	// The default value is true.
	SharedSpanContext *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=shared_span_context,json=sharedSpanContext,proto3" json:"shared_span_context,omitempty"`
	// Determines the selected collector endpoint version. By default, the “HTTP_JSON_V1“ will be
	// used.
	CollectorEndpointVersion ZipkinConfig_CollectorEndpointVersion `protobuf:"varint,5,opt,name=collector_endpoint_version,json=collectorEndpointVersion,proto3,enum=solo.io.envoy.config.trace.v3.ZipkinConfig_CollectorEndpointVersion" json:"collector_endpoint_version,omitempty"`
}

func (x *ZipkinConfig) Reset() {
	*x = ZipkinConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZipkinConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZipkinConfig) ProtoMessage() {}

func (x *ZipkinConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZipkinConfig.ProtoReflect.Descriptor instead.
func (*ZipkinConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescGZIP(), []int{0}
}

func (m *ZipkinConfig) GetCollectorCluster() isZipkinConfig_CollectorCluster {
	if m != nil {
		return m.CollectorCluster
	}
	return nil
}

func (x *ZipkinConfig) GetCollectorUpstreamRef() *core.ResourceRef {
	if x, ok := x.GetCollectorCluster().(*ZipkinConfig_CollectorUpstreamRef); ok {
		return x.CollectorUpstreamRef
	}
	return nil
}

func (x *ZipkinConfig) GetClusterName() string {
	if x, ok := x.GetCollectorCluster().(*ZipkinConfig_ClusterName); ok {
		return x.ClusterName
	}
	return ""
}

func (x *ZipkinConfig) GetCollectorEndpoint() string {
	if x != nil {
		return x.CollectorEndpoint
	}
	return ""
}

func (x *ZipkinConfig) GetTraceId_128Bit() *wrapperspb.BoolValue {
	if x != nil {
		return x.TraceId_128Bit
	}
	return nil
}

func (x *ZipkinConfig) GetSharedSpanContext() *wrapperspb.BoolValue {
	if x != nil {
		return x.SharedSpanContext
	}
	return nil
}

func (x *ZipkinConfig) GetCollectorEndpointVersion() ZipkinConfig_CollectorEndpointVersion {
	if x != nil {
		return x.CollectorEndpointVersion
	}
	return ZipkinConfig_DEPRECATED_AND_UNAVAILABLE_DO_NOT_USE
}

type isZipkinConfig_CollectorCluster interface {
	isZipkinConfig_CollectorCluster()
}

type ZipkinConfig_CollectorUpstreamRef struct {
	// The upstream that hosts the Zipkin collectors.
	CollectorUpstreamRef *core.ResourceRef `protobuf:"bytes,1,opt,name=collector_upstream_ref,json=collectorUpstreamRef,proto3,oneof"`
}

type ZipkinConfig_ClusterName struct {
	// The name of the cluster that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v3.Bootstrap.StaticResources.clusters>`.
	ClusterName string `protobuf:"bytes,6,opt,name=cluster_name,json=clusterName,proto3,oneof"`
}

func (*ZipkinConfig_CollectorUpstreamRef) isZipkinConfig_CollectorCluster() {}

func (*ZipkinConfig_ClusterName) isZipkinConfig_CollectorCluster() {}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDesc = []byte{
	0x0a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x7a, 0x69, 0x70, 0x6b, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x05, 0x0a, 0x0c, 0x5a, 0x69, 0x70,
	0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x16, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x66, 0x12, 0x23, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x31, 0x32, 0x38, 0x62, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x31, 0x32, 0x38, 0x62, 0x69, 0x74, 0x12, 0x4a, 0x0a,
	0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x70,
	0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x5a,
	0x69, 0x70, 0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6e,
	0x0a, 0x18, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x25, 0x44, 0x45,
	0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x55, 0x53, 0x45, 0x10, 0x00, 0x1a, 0x08, 0xa0, 0x8f, 0xa3, 0xa8, 0x05, 0x01, 0x08, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x10, 0x02, 0x3a, 0x2a,
	0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x24, 0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x5a, 0x69,
	0x70, 0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x13, 0x0a, 0x11, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42,
	0xcd, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x82, 0x8a,
	0xd7, 0xad, 0x04, 0x29, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x7a,
	0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0xb5, 0xdf,
	0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x33, 0x42, 0x0b, 0x5a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_goTypes = []interface{}{
	(ZipkinConfig_CollectorEndpointVersion)(0), // 0: solo.io.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion
	(*ZipkinConfig)(nil),                       // 1: solo.io.envoy.config.trace.v3.ZipkinConfig
	(*core.ResourceRef)(nil),                   // 2: core.solo.io.ResourceRef
	(*wrapperspb.BoolValue)(nil),               // 3: google.protobuf.BoolValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_depIdxs = []int32{
	2, // 0: solo.io.envoy.config.trace.v3.ZipkinConfig.collector_upstream_ref:type_name -> core.solo.io.ResourceRef
	3, // 1: solo.io.envoy.config.trace.v3.ZipkinConfig.trace_id_128bit:type_name -> google.protobuf.BoolValue
	3, // 2: solo.io.envoy.config.trace.v3.ZipkinConfig.shared_span_context:type_name -> google.protobuf.BoolValue
	0, // 3: solo.io.envoy.config.trace.v3.ZipkinConfig.collector_endpoint_version:type_name -> solo.io.envoy.config.trace.v3.ZipkinConfig.CollectorEndpointVersion
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZipkinConfig); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ZipkinConfig_CollectorUpstreamRef)(nil),
		(*ZipkinConfig_ClusterName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_zipkin_proto_depIdxs = nil
}
