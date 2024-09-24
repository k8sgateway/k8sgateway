// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/circuit_breaker.proto

package v1

import (
	reflect "reflect"
	sync "sync"

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

// CircuitBreakerConfig contains the options for customizing circuit breaking behavior.
// See the [envoy docs](https://www.envoyproxy.io/docs/envoy/v1.14.1/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
// for the meaning of these values.
type CircuitBreakerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxConnections     *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests        *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries         *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
}

func (x *CircuitBreakerConfig) Reset() {
	*x = CircuitBreakerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CircuitBreakerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CircuitBreakerConfig) ProtoMessage() {}

func (x *CircuitBreakerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CircuitBreakerConfig.ProtoReflect.Descriptor instead.
func (*CircuitBreakerConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescGZIP(), []int{0}
}

func (x *CircuitBreakerConfig) GetMaxConnections() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxConnections
	}
	return nil
}

func (x *CircuitBreakerConfig) GetMaxPendingRequests() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxPendingRequests
	}
	return nil
}

func (x *CircuitBreakerConfig) GetMaxRequests() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxRequests
	}
	return nil
}

func (x *CircuitBreakerConfig) GetMaxRetries() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxRetries
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x14, 0x43, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x45, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x3e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01,
	0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_goTypes = []any{
	(*CircuitBreakerConfig)(nil), // 0: gloo.solo.io.CircuitBreakerConfig
	(*wrappers.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_depIdxs = []int32{
	1, // 0: gloo.solo.io.CircuitBreakerConfig.max_connections:type_name -> google.protobuf.UInt32Value
	1, // 1: gloo.solo.io.CircuitBreakerConfig.max_pending_requests:type_name -> google.protobuf.UInt32Value
	1, // 2: gloo.solo.io.CircuitBreakerConfig.max_requests:type_name -> google.protobuf.UInt32Value
	1, // 3: gloo.solo.io.CircuitBreakerConfig.max_retries:type_name -> google.protobuf.UInt32Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CircuitBreakerConfig); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_circuit_breaker_proto_depIdxs = nil
}
