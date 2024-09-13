// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gateway/api/v1/matchable_http_gateway.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	ssl "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A MatchableHttpGateway describes a single FilterChain configured with:
// - The HttpConnectionManager NetworkFilter
// - A FilterChainMatch and TransportSocket that support TLS configuration and Source IP matching
//
// A Gateway CR may select one or more MatchableHttpGateways on a single listener.
// This enables separate teams to own Listener configuration (Gateway CR)
// and FilterChain configuration (MatchableHttpGateway CR)
type MatchableHttpGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NamespacedStatuses indicates the validation status of this resource.
	// NamespacedStatuses is read-only by clients, and set by gateway during validation
	NamespacedStatuses *core.NamespacedStatuses `protobuf:"bytes,1,opt,name=namespaced_statuses,json=namespacedStatuses,proto3" json:"namespaced_statuses,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata *core.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Matcher creates a FilterChainMatch and TransportSocket for a FilterChain
	// For each MatchableHttpGateway on a Gateway CR, the matcher must be unique.
	// If there are any identical matchers, the Gateway will be rejected.
	// An empty matcher will produce an empty FilterChainMatch (https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#envoy-v3-api-msg-config-listener-v3-filterchainmatch)
	// effectively matching all incoming connections
	Matcher *MatchableHttpGateway_Matcher `protobuf:"bytes,3,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// HttpGateway creates a FilterChain with an HttpConnectionManager
	HttpGateway *HttpGateway `protobuf:"bytes,4,opt,name=http_gateway,json=httpGateway,proto3" json:"http_gateway,omitempty"`
}

func (x *MatchableHttpGateway) Reset() {
	*x = MatchableHttpGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableHttpGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableHttpGateway) ProtoMessage() {}

func (x *MatchableHttpGateway) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableHttpGateway.ProtoReflect.Descriptor instead.
func (*MatchableHttpGateway) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *MatchableHttpGateway) GetNamespacedStatuses() *core.NamespacedStatuses {
	if x != nil {
		return x.NamespacedStatuses
	}
	return nil
}

func (x *MatchableHttpGateway) GetMetadata() *core.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MatchableHttpGateway) GetMatcher() *MatchableHttpGateway_Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *MatchableHttpGateway) GetHttpGateway() *HttpGateway {
	if x != nil {
		return x.HttpGateway
	}
	return nil
}

type MatchableHttpGateway_Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CidrRange specifies an IP Address and a prefix length to construct the subnet mask for a CIDR range.
	// See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/address.proto#envoy-v3-api-msg-config-core-v3-cidrrange
	SourcePrefixRanges []*v3.CidrRange `protobuf:"bytes,1,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// Ssl configuration applied to the FilterChain:
	//   - FilterChainMatch: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#config-listener-v3-filterchainmatch)
	//   - TransportSocket: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/base.proto#envoy-v3-api-msg-config-core-v3-transportsocket
	SslConfig *ssl.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
}

func (x *MatchableHttpGateway_Matcher) Reset() {
	*x = MatchableHttpGateway_Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchableHttpGateway_Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableHttpGateway_Matcher) ProtoMessage() {}

func (x *MatchableHttpGateway_Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableHttpGateway_Matcher.ProtoReflect.Descriptor instead.
func (*MatchableHttpGateway_Matcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MatchableHttpGateway_Matcher) GetSourcePrefixRanges() []*v3.CidrRange {
	if x != nil {
		return x.SourcePrefixRanges
	}
	return nil
}

func (x *MatchableHttpGateway_Matcher) GetSslConfig() *ssl.SslConfig {
	if x != nil {
		return x.SslConfig
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b,
	0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x6c, 0x2f, 0x73,
	0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe6, 0x03, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74,
	0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x57, 0x0a, 0x13, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x42, 0x04, 0xb8, 0xf5, 0x04, 0x01, 0x52, 0x12, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3f,
	0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a,
	0x9c, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x14, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x73, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x09, 0x73, 0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x18,
	0x82, 0xf1, 0x04, 0x14, 0x0a, 0x03, 0x68, 0x67, 0x77, 0x12, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x42, 0x41, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_goTypes = []interface{}{
	(*MatchableHttpGateway)(nil),         // 0: gateway.solo.io.MatchableHttpGateway
	(*MatchableHttpGateway_Matcher)(nil), // 1: gateway.solo.io.MatchableHttpGateway.Matcher
	(*core.NamespacedStatuses)(nil),      // 2: core.solo.io.NamespacedStatuses
	(*core.Metadata)(nil),                // 3: core.solo.io.Metadata
	(*HttpGateway)(nil),                  // 4: gateway.solo.io.HttpGateway
	(*v3.CidrRange)(nil),                 // 5: solo.io.envoy.config.core.v3.CidrRange
	(*ssl.SslConfig)(nil),                // 6: gloo.solo.io.SslConfig
}
var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_depIdxs = []int32{
	2, // 0: gateway.solo.io.MatchableHttpGateway.namespaced_statuses:type_name -> core.solo.io.NamespacedStatuses
	3, // 1: gateway.solo.io.MatchableHttpGateway.metadata:type_name -> core.solo.io.Metadata
	1, // 2: gateway.solo.io.MatchableHttpGateway.matcher:type_name -> gateway.solo.io.MatchableHttpGateway.Matcher
	4, // 3: gateway.solo.io.MatchableHttpGateway.http_gateway:type_name -> gateway.solo.io.HttpGateway
	5, // 4: gateway.solo.io.MatchableHttpGateway.Matcher.source_prefix_ranges:type_name -> solo.io.envoy.config.core.v3.CidrRange
	6, // 5: gateway.solo.io.MatchableHttpGateway.Matcher.ssl_config:type_name -> gloo.solo.io.SslConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_init() }
func file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableHttpGateway); i {
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
		file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchableHttpGateway_Matcher); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_http_gateway_proto_depIdxs = nil
}
