// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/retries/retries.proto

package retries

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This specifies the retry policy interval for backoffs. Note that if the base interval provided is larger than the maximum interval OR if any of the durations passed are <= 0 MS, there will be an error.
type RetryBackOff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the base interval for a retry
	BaseInterval *duration.Duration `protobuf:"bytes,1,opt,name=base_interval,json=baseInterval,proto3" json:"base_interval,omitempty"`
	// Specifies the max interval for a retry
	MaxInterval *duration.Duration `protobuf:"bytes,2,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
}

func (x *RetryBackOff) Reset() {
	*x = RetryBackOff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryBackOff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryBackOff) ProtoMessage() {}

func (x *RetryBackOff) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryBackOff.ProtoReflect.Descriptor instead.
func (*RetryBackOff) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP(), []int{0}
}

func (x *RetryBackOff) GetBaseInterval() *duration.Duration {
	if x != nil {
		return x.BaseInterval
	}
	return nil
}

func (x *RetryBackOff) GetMaxInterval() *duration.Duration {
	if x != nil {
		return x.MaxInterval
	}
	return nil
}

// Retry Policy applied at the Route and/or Virtual Hosts levels.
type RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the conditions under which retry takes place. These are the same
	// conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/v1.14.1/configuration/http/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	RetryOn string `protobuf:"bytes,1,opt,name=retry_on,json=retryOn,proto3" json:"retry_on,omitempty"`
	// Specifies the allowed number of retries. This parameter is optional and
	// defaults to 1. These are the same conditions [documented for Envoy](https://www.envoyproxy.io/docs/envoy/v1.14.1/configuration/http/http_filters/router_filter#config-http-filters-router-x-envoy-retry-on)
	NumRetries uint32 `protobuf:"varint,2,opt,name=num_retries,json=numRetries,proto3" json:"num_retries,omitempty"`
	// Specifies a non-zero upstream timeout per retry attempt. This parameter is optional.
	PerTryTimeout *duration.Duration `protobuf:"bytes,3,opt,name=per_try_timeout,json=perTryTimeout,proto3" json:"per_try_timeout,omitempty"`
	// Specifies the retry policy interval
	RetryBackOff *RetryBackOff `protobuf:"bytes,4,opt,name=retry_back_off,json=retryBackOff,proto3" json:"retry_back_off,omitempty"`
	// Types that are assignable to PriorityPredicate:
	//
	//	*RetryPolicy_PreviousPriorities_
	PriorityPredicate isRetryPolicy_PriorityPredicate `protobuf_oneof:"priority_predicate"`
}

func (x *RetryPolicy) Reset() {
	*x = RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy) ProtoMessage() {}

func (x *RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP(), []int{1}
}

func (x *RetryPolicy) GetRetryOn() string {
	if x != nil {
		return x.RetryOn
	}
	return ""
}

func (x *RetryPolicy) GetNumRetries() uint32 {
	if x != nil {
		return x.NumRetries
	}
	return 0
}

func (x *RetryPolicy) GetPerTryTimeout() *duration.Duration {
	if x != nil {
		return x.PerTryTimeout
	}
	return nil
}

func (x *RetryPolicy) GetRetryBackOff() *RetryBackOff {
	if x != nil {
		return x.RetryBackOff
	}
	return nil
}

func (m *RetryPolicy) GetPriorityPredicate() isRetryPolicy_PriorityPredicate {
	if m != nil {
		return m.PriorityPredicate
	}
	return nil
}

func (x *RetryPolicy) GetPreviousPriorities() *RetryPolicy_PreviousPriorities {
	if x, ok := x.GetPriorityPredicate().(*RetryPolicy_PreviousPriorities_); ok {
		return x.PreviousPriorities
	}
	return nil
}

type isRetryPolicy_PriorityPredicate interface {
	isRetryPolicy_PriorityPredicate()
}

type RetryPolicy_PreviousPriorities_ struct {
	// Specify the previous priorities.
	// For more information about previous priorities, see the [Envoy docs](https://www.envoyproxy.io/docs/envoy/v1.30.1/api-v3/extensions/retry/priority/previous_priorities/v3/previous_priorities_config.proto#envoy-v3-api-file-envoy-extensions-retry-priority-previous-priorities-v3-previous-priorities-config-proto).
	PreviousPriorities *RetryPolicy_PreviousPriorities `protobuf:"bytes,5,opt,name=previous_priorities,json=previousPriorities,proto3,oneof"`
}

func (*RetryPolicy_PreviousPriorities_) isRetryPolicy_PriorityPredicate() {}

type RetryPolicy_PreviousPriorities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify the update frequency for the previous priorities. For more information about previous priorities, see the [Envoy docs](https://www.envoyproxy.io/docs/envoy/v1.30.1/api-v3/extensions/retry/priority/previous_priorities/v3/previous_priorities_config.proto#envoy-v3-api-file-envoy-extensions-retry-priority-previous-priorities-v3-previous-priorities-config-proto).
	// This option only works in combination with an Upstream failover policy that enables priorities.
	UpdateFrequency *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
}

func (x *RetryPolicy_PreviousPriorities) Reset() {
	*x = RetryPolicy_PreviousPriorities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy_PreviousPriorities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy_PreviousPriorities) ProtoMessage() {}

func (x *RetryPolicy_PreviousPriorities) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy_PreviousPriorities.ProtoReflect.Descriptor instead.
func (*RetryPolicy_PreviousPriorities) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RetryPolicy_PreviousPriorities) GetUpdateFrequency() *wrappers.UInt32Value {
	if x != nil {
		return x.UpdateFrequency
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x42, 0x61, 0x63, 0x6b, 0x4f, 0x66, 0x66, 0x12, 0x4e, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0xaa,
	0x01, 0x08, 0x08, 0x01, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01,
	0x02, 0x2a, 0x00, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x22, 0xc4, 0x03, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e,
	0x75, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x70, 0x65, 0x72, 0x54, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x50, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x66,
	0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b,
	0x4f, 0x66, 0x66, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x4f, 0x66,
	0x66, 0x12, 0x6f, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x12,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x1a, 0x5d, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x42, 0x14, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x4e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04,
	0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_goTypes = []any{
	(*RetryBackOff)(nil),                   // 0: retries.options.gloo.solo.io.RetryBackOff
	(*RetryPolicy)(nil),                    // 1: retries.options.gloo.solo.io.RetryPolicy
	(*RetryPolicy_PreviousPriorities)(nil), // 2: retries.options.gloo.solo.io.RetryPolicy.PreviousPriorities
	(*duration.Duration)(nil),              // 3: google.protobuf.Duration
	(*wrappers.UInt32Value)(nil),           // 4: google.protobuf.UInt32Value
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_depIdxs = []int32{
	3, // 0: retries.options.gloo.solo.io.RetryBackOff.base_interval:type_name -> google.protobuf.Duration
	3, // 1: retries.options.gloo.solo.io.RetryBackOff.max_interval:type_name -> google.protobuf.Duration
	3, // 2: retries.options.gloo.solo.io.RetryPolicy.per_try_timeout:type_name -> google.protobuf.Duration
	0, // 3: retries.options.gloo.solo.io.RetryPolicy.retry_back_off:type_name -> retries.options.gloo.solo.io.RetryBackOff
	2, // 4: retries.options.gloo.solo.io.RetryPolicy.previous_priorities:type_name -> retries.options.gloo.solo.io.RetryPolicy.PreviousPriorities
	4, // 5: retries.options.gloo.solo.io.RetryPolicy.PreviousPriorities.update_frequency:type_name -> google.protobuf.UInt32Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RetryBackOff); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RetryPolicy); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RetryPolicy_PreviousPriorities); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes[1].OneofWrappers = []any{
		(*RetryPolicy_PreviousPriorities_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_retries_retries_proto_depIdxs = nil
}
