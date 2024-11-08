// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.8
// source: github.com/solo-io/gloo/projects/gloo/api/v1/listener_options.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	als "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/als"
	proxy_protocol "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/proxy_protocol"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Optional, feature-specific configuration that lives on gateways.
// Each ListenerOption object contains configuration for a specific feature.
// Note to developers: new Listener plugins must be added to this struct
// to be usable by Gloo. (plugins currently need to be compiled into Gloo)
type ListenerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for access logging in a filter like the HttpConnectionManager.
	AccessLoggingService *als.AccessLoggingService `protobuf:"bytes,1,opt,name=access_logging_service,json=accessLoggingService,proto3" json:"access_logging_service,omitempty"`
	// Extensions will be passed along from Listeners, Gateways, VirtualServices, Routes, and Route tables to the
	// underlying Proxy, making them useful for controllers, validation tools, etc. which interact with kubernetes yaml.
	//
	// Some sample use cases:
	// * controllers, deployment pipelines, helm charts, etc. which wish to use extensions as a kind of opaque metadata.
	// * In the future, Gloo may support gRPC-based plugins which communicate with the Gloo translator out-of-process.
	// Opaque Extensions enables development of out-of-process plugins without requiring recompiling & redeploying Gloo's API.
	Extensions *Extensions `protobuf:"bytes,2,opt,name=extensions,proto3" json:"extensions,omitempty"`
	// Soft limit on size of the listener's new connection read and write buffers. If unspecified, defaults to 1MiB
	// For more info, check out the [Envoy docs](https://www.envoyproxy.io/docs/envoy/v1.14.1/api-v2/api/v2/listener.proto)
	PerConnectionBufferLimitBytes *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions []*core.SocketOption `protobuf:"bytes,4,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	// Enable ProxyProtocol support for this listener.
	ProxyProtocol *proxy_protocol.ProxyProtocol `protobuf:"bytes,5,opt,name=proxy_protocol,json=proxyProtocol,proto3" json:"proxy_protocol,omitempty"`
	// Configuration for listener connection balancing.
	ConnectionBalanceConfig *ConnectionBalanceConfig `protobuf:"bytes,6,opt,name=connection_balance_config,json=connectionBalanceConfig,proto3" json:"connection_balance_config,omitempty"`
	// If enabled this sets up an early access logging service for the listener.
	// Added initially to support listener level logging for HTTP listeners.
	// For more info see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-access-log
	ListenerAccessLoggingService *als.AccessLoggingService `protobuf:"bytes,7,opt,name=listener_access_logging_service,json=listenerAccessLoggingService,proto3" json:"listener_access_logging_service,omitempty"`
}

func (x *ListenerOptions) Reset() {
	*x = ListenerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerOptions) ProtoMessage() {}

func (x *ListenerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerOptions.ProtoReflect.Descriptor instead.
func (*ListenerOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescGZIP(), []int{0}
}

func (x *ListenerOptions) GetAccessLoggingService() *als.AccessLoggingService {
	if x != nil {
		return x.AccessLoggingService
	}
	return nil
}

func (x *ListenerOptions) GetExtensions() *Extensions {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *ListenerOptions) GetPerConnectionBufferLimitBytes() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PerConnectionBufferLimitBytes
	}
	return nil
}

func (x *ListenerOptions) GetSocketOptions() []*core.SocketOption {
	if x != nil {
		return x.SocketOptions
	}
	return nil
}

func (x *ListenerOptions) GetProxyProtocol() *proxy_protocol.ProxyProtocol {
	if x != nil {
		return x.ProxyProtocol
	}
	return nil
}

func (x *ListenerOptions) GetConnectionBalanceConfig() *ConnectionBalanceConfig {
	if x != nil {
		return x.ConnectionBalanceConfig
	}
	return nil
}

func (x *ListenerOptions) GetListenerAccessLoggingService() *als.AccessLoggingService {
	if x != nil {
		return x.ListenerAccessLoggingService
	}
	return nil
}

// Configuration for listener connection balancing.
type ConnectionBalanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExactBalance *ConnectionBalanceConfig_ExactBalance `protobuf:"bytes,1,opt,name=exact_balance,json=exactBalance,proto3" json:"exact_balance,omitempty"`
}

func (x *ConnectionBalanceConfig) Reset() {
	*x = ConnectionBalanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionBalanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionBalanceConfig) ProtoMessage() {}

func (x *ConnectionBalanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionBalanceConfig.ProtoReflect.Descriptor instead.
func (*ConnectionBalanceConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionBalanceConfig) GetExactBalance() *ConnectionBalanceConfig_ExactBalance {
	if x != nil {
		return x.ExactBalance
	}
	return nil
}

// A connection balancer implementation that does exact balancing. This means that a lock is
// held during balancing so that connection counts are nearly exactly balanced between worker
// threads. This is "nearly" exact in the sense that a connection might close in parallel thus
// making the counts incorrect, but this should be rectified on the next accept. This balancer
// sacrifices accept throughput for accuracy and should be used when there are a small number of
// connections that rarely cycle (e.g., service mesh gRPC egress).
type ConnectionBalanceConfig_ExactBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionBalanceConfig_ExactBalance) Reset() {
	*x = ConnectionBalanceConfig_ExactBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionBalanceConfig_ExactBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionBalanceConfig_ExactBalance) ProtoMessage() {}

func (x *ConnectionBalanceConfig_ExactBalance) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionBalanceConfig_ExactBalance.ProtoReflect.Descriptor instead.
func (*ConnectionBalanceConfig_ExactBalance) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescGZIP(), []int{1, 0}
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6c, 0x73, 0x2f, 0x61, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x05, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x64, 0x0a, 0x16, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x66, 0x0a, 0x21, 0x70, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x1d, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x4e, 0x0a, 0x0e, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x59, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x0d, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x61, 0x0a, 0x19, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x75,
	0x0a, 0x1f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x1c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x57, 0x0a, 0x0d, 0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x45, 0x78, 0x61, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x65, 0x78,
	0x61, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x0e, 0x0a, 0x0c, 0x45, 0x78,
	0x61, 0x63, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x3e, 0xb8, 0xf5, 0x04, 0x01,
	0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_goTypes = []any{
	(*ListenerOptions)(nil),                      // 0: gloo.solo.io.ListenerOptions
	(*ConnectionBalanceConfig)(nil),              // 1: gloo.solo.io.ConnectionBalanceConfig
	(*ConnectionBalanceConfig_ExactBalance)(nil), // 2: gloo.solo.io.ConnectionBalanceConfig.ExactBalance
	(*als.AccessLoggingService)(nil),             // 3: als.options.gloo.solo.io.AccessLoggingService
	(*Extensions)(nil),                           // 4: gloo.solo.io.Extensions
	(*wrapperspb.UInt32Value)(nil),               // 5: google.protobuf.UInt32Value
	(*core.SocketOption)(nil),                    // 6: solo.io.envoy.api.v2.core.SocketOption
	(*proxy_protocol.ProxyProtocol)(nil),         // 7: proxy_protocol.options.gloo.solo.io.ProxyProtocol
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_depIdxs = []int32{
	3, // 0: gloo.solo.io.ListenerOptions.access_logging_service:type_name -> als.options.gloo.solo.io.AccessLoggingService
	4, // 1: gloo.solo.io.ListenerOptions.extensions:type_name -> gloo.solo.io.Extensions
	5, // 2: gloo.solo.io.ListenerOptions.per_connection_buffer_limit_bytes:type_name -> google.protobuf.UInt32Value
	6, // 3: gloo.solo.io.ListenerOptions.socket_options:type_name -> solo.io.envoy.api.v2.core.SocketOption
	7, // 4: gloo.solo.io.ListenerOptions.proxy_protocol:type_name -> proxy_protocol.options.gloo.solo.io.ProxyProtocol
	1, // 5: gloo.solo.io.ListenerOptions.connection_balance_config:type_name -> gloo.solo.io.ConnectionBalanceConfig
	3, // 6: gloo.solo.io.ListenerOptions.listener_access_logging_service:type_name -> als.options.gloo.solo.io.AccessLoggingService
	2, // 7: gloo.solo.io.ConnectionBalanceConfig.exact_balance:type_name -> gloo.solo.io.ConnectionBalanceConfig.ExactBalance
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_extensions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListenerOptions); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionBalanceConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionBalanceConfig_ExactBalance); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_listener_options_proto_depIdxs = nil
}
