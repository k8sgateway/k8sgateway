// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/caching/caching.proto

package caching

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The basic reference for the service. Details name and namespace.
	CachingServiceRef *core.ResourceRef `protobuf:"bytes,1,opt,name=caching_service_ref,json=cachingServiceRef,proto3" json:"caching_service_ref,omitempty"`
	// A list of string matchers that state what headers are allowed to vary
	// and still be cached. Per upstream envoy allowed vary headers.
	AllowedVaryHeaders []*v3.StringMatcher `protobuf:"bytes,2,rep,name=allowed_vary_headers,json=allowedVaryHeaders,proto3" json:"allowed_vary_headers,omitempty"`
	// Connection timeout  for retrieval from an sync cache
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Max payload size to cache. If unset defaults to a reasonable value.
	// If explicitly set to 0 will prevent anything with a body from
	// being cached.
	MaxPayloadSize *wrappers.UInt64Value `protobuf:"bytes,4,opt,name=max_payload_size,json=maxPayloadSize,proto3" json:"max_payload_size,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_msgTypes[0]
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
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetCachingServiceRef() *core.ResourceRef {
	if x != nil {
		return x.CachingServiceRef
	}
	return nil
}

func (x *Settings) GetAllowedVaryHeaders() []*v3.StringMatcher {
	if x != nil {
		return x.AllowedVaryHeaders
	}
	return nil
}

func (x *Settings) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Settings) GetMaxPayloadSize() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxPayloadSize
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDesc = []byte{
	0x0a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2,
	0x02, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x49, 0x0a, 0x13, 0x63,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x52, 0x11, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x5e, 0x0a, 0x14, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x56, 0x61, 0x72, 0x79, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x10, 0x6d,
	0x61, 0x78, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x69, 0x7a, 0x65, 0x42, 0x59, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_goTypes = []interface{}{
	(*Settings)(nil),             // 0: caching.options.gloo.solo.io.Settings
	(*core.ResourceRef)(nil),     // 1: core.solo.io.ResourceRef
	(*v3.StringMatcher)(nil),     // 2: solo.io.envoy.type.matcher.v3.StringMatcher
	(*duration.Duration)(nil),    // 3: google.protobuf.Duration
	(*wrappers.UInt64Value)(nil), // 4: google.protobuf.UInt64Value
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_depIdxs = []int32{
	1, // 0: caching.options.gloo.solo.io.Settings.caching_service_ref:type_name -> core.solo.io.ResourceRef
	2, // 1: caching.options.gloo.solo.io.Settings.allowed_vary_headers:type_name -> solo.io.envoy.type.matcher.v3.StringMatcher
	3, // 2: caching.options.gloo.solo.io.Settings.timeout:type_name -> google.protobuf.Duration
	4, // 3: caching.options.gloo.solo.io.Settings.max_payload_size:type_name -> google.protobuf.UInt64Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_caching_caching_proto_depIdxs = nil
}
