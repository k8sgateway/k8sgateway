// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/datadog.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for the Remote Configuration feature.
type DatadogRemoteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Frequency at which new configuration updates are queried.
	// If no value is provided, the default value is delegated to the Datadog tracing library.
	PollingInterval *durationpb.Duration `protobuf:"bytes,1,opt,name=polling_interval,json=pollingInterval,proto3" json:"polling_interval,omitempty"`
	// Disabled remote config.
	// This field does not exist in envoy's config but allow us to preserve the default behavior
	// when upgrading to envoy v1.31
	Disabled *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=disabled,proto3" json:"disabled,omitempty"`
}

func (x *DatadogRemoteConfig) Reset() {
	*x = DatadogRemoteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatadogRemoteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatadogRemoteConfig) ProtoMessage() {}

func (x *DatadogRemoteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatadogRemoteConfig.ProtoReflect.Descriptor instead.
func (*DatadogRemoteConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescGZIP(), []int{0}
}

func (x *DatadogRemoteConfig) GetPollingInterval() *durationpb.Duration {
	if x != nil {
		return x.PollingInterval
	}
	return nil
}

func (x *DatadogRemoteConfig) GetDisabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

// Configuration for the Datadog tracer.
// [#extension: envoy.tracers.datadog]
type DatadogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cluster to use for submitting traces to the Datadog agent.
	//
	// Types that are assignable to CollectorCluster:
	//
	//	*DatadogConfig_CollectorUpstreamRef
	//	*DatadogConfig_ClusterName
	CollectorCluster isDatadogConfig_CollectorCluster `protobuf_oneof:"collector_cluster"`
	// The name used for the service when traces are generated by envoy.
	ServiceName *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Optional hostname to use when sending spans to the collector_cluster. Useful for collectors
	// that require a specific hostname. Defaults to collector_cluster above.
	CollectorHostname string `protobuf:"bytes,4,opt,name=collector_hostname,json=collectorHostname,proto3" json:"collector_hostname,omitempty"`
	// Configures remote configuration.
	// Remote Configuration allows to configure the tracer from Datadog's user interface.
	// This feature can drastically increase the number of connections to the Datadog Agent.
	// Each tracer regularly polls for configuration updates, and the number of tracers is the product
	// of the number of listeners and worker threads.
	RemoteConfig *DatadogRemoteConfig `protobuf:"bytes,5,opt,name=remote_config,json=remoteConfig,proto3" json:"remote_config,omitempty"`
}

func (x *DatadogConfig) Reset() {
	*x = DatadogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatadogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatadogConfig) ProtoMessage() {}

func (x *DatadogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatadogConfig.ProtoReflect.Descriptor instead.
func (*DatadogConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescGZIP(), []int{1}
}

func (m *DatadogConfig) GetCollectorCluster() isDatadogConfig_CollectorCluster {
	if m != nil {
		return m.CollectorCluster
	}
	return nil
}

func (x *DatadogConfig) GetCollectorUpstreamRef() *core.ResourceRef {
	if x, ok := x.GetCollectorCluster().(*DatadogConfig_CollectorUpstreamRef); ok {
		return x.CollectorUpstreamRef
	}
	return nil
}

func (x *DatadogConfig) GetClusterName() string {
	if x, ok := x.GetCollectorCluster().(*DatadogConfig_ClusterName); ok {
		return x.ClusterName
	}
	return ""
}

func (x *DatadogConfig) GetServiceName() *wrapperspb.StringValue {
	if x != nil {
		return x.ServiceName
	}
	return nil
}

func (x *DatadogConfig) GetCollectorHostname() string {
	if x != nil {
		return x.CollectorHostname
	}
	return ""
}

func (x *DatadogConfig) GetRemoteConfig() *DatadogRemoteConfig {
	if x != nil {
		return x.RemoteConfig
	}
	return nil
}

type isDatadogConfig_CollectorCluster interface {
	isDatadogConfig_CollectorCluster()
}

type DatadogConfig_CollectorUpstreamRef struct {
	// The upstream to use for submitting traces to the Datadog agent.
	CollectorUpstreamRef *core.ResourceRef `protobuf:"bytes,1,opt,name=collector_upstream_ref,json=collectorUpstreamRef,proto3,oneof"`
}

type DatadogConfig_ClusterName struct {
	// The name of the cluster to use for submitting traces to the Datadog agent. Note that the
	// cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v3.Bootstrap.StaticResources.clusters>`.
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3,oneof"`
}

func (*DatadogConfig_CollectorUpstreamRef) isDatadogConfig_CollectorCluster() {}

func (*DatadogConfig_ClusterName) isDatadogConfig_CollectorCluster() {}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x10,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x9b, 0x03, 0x0a, 0x0d, 0x44,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x16,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x66, 0x12,
	0x23, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d,
	0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a,
	0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x2b, 0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x25, 0x0a, 0x23,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x13, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0xcf, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0,
	0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x82, 0x8a, 0xd7, 0xad, 0x04, 0x2a, 0x12, 0x28, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a,
	0x2b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x44, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_goTypes = []any{
	(*DatadogRemoteConfig)(nil),    // 0: solo.io.envoy.config.trace.v3.DatadogRemoteConfig
	(*DatadogConfig)(nil),          // 1: solo.io.envoy.config.trace.v3.DatadogConfig
	(*durationpb.Duration)(nil),    // 2: google.protobuf.Duration
	(*wrapperspb.BoolValue)(nil),   // 3: google.protobuf.BoolValue
	(*core.ResourceRef)(nil),       // 4: core.solo.io.ResourceRef
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_depIdxs = []int32{
	2, // 0: solo.io.envoy.config.trace.v3.DatadogRemoteConfig.polling_interval:type_name -> google.protobuf.Duration
	3, // 1: solo.io.envoy.config.trace.v3.DatadogRemoteConfig.disabled:type_name -> google.protobuf.BoolValue
	4, // 2: solo.io.envoy.config.trace.v3.DatadogConfig.collector_upstream_ref:type_name -> core.solo.io.ResourceRef
	5, // 3: solo.io.envoy.config.trace.v3.DatadogConfig.service_name:type_name -> google.protobuf.StringValue
	0, // 4: solo.io.envoy.config.trace.v3.DatadogConfig.remote_config:type_name -> solo.io.envoy.config.trace.v3.DatadogRemoteConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DatadogRemoteConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DatadogConfig); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes[1].OneofWrappers = []any{
		(*DatadogConfig_CollectorUpstreamRef)(nil),
		(*DatadogConfig_ClusterName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_trace_v3_datadog_proto_depIdxs = nil
}
