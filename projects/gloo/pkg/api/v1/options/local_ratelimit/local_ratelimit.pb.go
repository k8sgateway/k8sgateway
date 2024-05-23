// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/local_ratelimit/local_ratelimit.proto

package local_ratelimit

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configures the token bucket, used for rate limiting.
// Ref. https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
type TokenBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum tokens that the bucket can hold. This is also the number of tokens that the bucket initially contains.
	// Must be greater than or equal to one.
	MaxTokens uint32 `protobuf:"varint,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// The number of tokens added to the bucket during each fill interval. If not specified, defaults to a single token.
	// Must be greater than zero.
	TokensPerFill *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=tokens_per_fill,json=tokensPerFill,proto3" json:"tokens_per_fill,omitempty"`
	// The fill interval that tokens are added to the bucket. During each fill interval tokens_per_fill are added to the bucket.
	// The bucket will never contain more than max_tokens tokens.
	// The fill_interval must be >= 50ms and defaults to 1 second.
	FillInterval *duration.Duration `protobuf:"bytes,3,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
}

func (x *TokenBucket) Reset() {
	*x = TokenBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBucket) ProtoMessage() {}

func (x *TokenBucket) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBucket.ProtoReflect.Descriptor instead.
func (*TokenBucket) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *TokenBucket) GetMaxTokens() uint32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *TokenBucket) GetTokensPerFill() *wrappers.UInt32Value {
	if x != nil {
		return x.TokensPerFill
	}
	return nil
}

func (x *TokenBucket) GetFillInterval() *duration.Duration {
	if x != nil {
		return x.FillInterval
	}
	return nil
}

// The Local Rate Limit settings define the default local rate limit token bucket to apply as well as other configurations
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token bucket configuration to use for rate limiting requests.
	// These options provide the ability to locally rate limit the connections in envoy. Each request processed by the filter consumes a single token.
	// If the token is available, the request will be allowed. If no tokens are available, the request will receive the configured rate limit status.
	// This default limit can be overridden in the vHost or route options.localRatelimit
	DefaultLimit *TokenBucket `protobuf:"bytes,1,opt,name=default_limit,json=defaultLimit,proto3" json:"default_limit,omitempty"`
	// Specifies the scope of the rate limiter’s token bucket. If set to false, the token bucket is shared across all worker threads, thus the rate limits are applied per Envoy process.
	// If set to true, a token bucket is allocated for each connection, thus the rate limits are applied per connection thereby allowing one to rate limit requests on a per connection basis.
	// This setting applies to all token buckets in the vHost and route as well.
	// Defaults to false
	LocalRateLimitPerDownstreamConnection *wrappers.BoolValue `protobuf:"bytes,2,opt,name=local_rate_limit_per_downstream_connection,json=localRateLimitPerDownstreamConnection,proto3" json:"local_rate_limit_per_downstream_connection,omitempty"`
	// Set this to true to return Envoy's X-RateLimit headers to the downstream.
	// reference docs here: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/common/ratelimit/v3/ratelimit.proto#envoy-v3-api-enum-extensions-common-ratelimit-v3-xratelimitheadersrfcversion
	// This setting applies at the vHost and route local rate limit as well
	// Defaults to false
	EnableXRatelimitHeaders *wrappers.BoolValue `protobuf:"bytes,3,opt,name=enable_x_ratelimit_headers,json=enableXRatelimitHeaders,proto3" json:"enable_x_ratelimit_headers,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescGZIP(), []int{1}
}

func (x *Settings) GetDefaultLimit() *TokenBucket {
	if x != nil {
		return x.DefaultLimit
	}
	return nil
}

func (x *Settings) GetLocalRateLimitPerDownstreamConnection() *wrappers.BoolValue {
	if x != nil {
		return x.LocalRateLimitPerDownstreamConnection
	}
	return nil
}

func (x *Settings) GetEnableXRatelimitHeaders() *wrappers.BoolValue {
	if x != nil {
		return x.EnableXRatelimitHeaders
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01,
	0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x0f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x6c, 0x12, 0x3e, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x22, 0xb2, 0x02, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x56, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x75, 0x0a, 0x2a, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x25, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x50, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57,
	0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x58, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x42, 0x56, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04,
	0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_goTypes = []interface{}{
	(*TokenBucket)(nil),          // 0: local_ratelimit.options.gloo.solo.io.TokenBucket
	(*Settings)(nil),             // 1: local_ratelimit.options.gloo.solo.io.Settings
	(*wrappers.UInt32Value)(nil), // 2: google.protobuf.UInt32Value
	(*duration.Duration)(nil),    // 3: google.protobuf.Duration
	(*wrappers.BoolValue)(nil),   // 4: google.protobuf.BoolValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_depIdxs = []int32{
	2, // 0: local_ratelimit.options.gloo.solo.io.TokenBucket.tokens_per_fill:type_name -> google.protobuf.UInt32Value
	3, // 1: local_ratelimit.options.gloo.solo.io.TokenBucket.fill_interval:type_name -> google.protobuf.Duration
	0, // 2: local_ratelimit.options.gloo.solo.io.Settings.default_limit:type_name -> local_ratelimit.options.gloo.solo.io.TokenBucket
	4, // 3: local_ratelimit.options.gloo.solo.io.Settings.local_rate_limit_per_downstream_connection:type_name -> google.protobuf.BoolValue
	4, // 4: local_ratelimit.options.gloo.solo.io.Settings.enable_x_ratelimit_headers:type_name -> google.protobuf.BoolValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBucket); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_local_ratelimit_local_ratelimit_proto_depIdxs = nil
}
