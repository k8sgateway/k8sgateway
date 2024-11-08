// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.8
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/cors/cors.proto

package cors

import (
	reflect "reflect"
	sync "sync"

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

// CorsPolicy defines Cross-Origin Resource Sharing for a virtual service.
type CorsPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the origins that will be allowed to make CORS requests.
	//
	// An origin is allowed if either allow_origin or allow_origin_regex match.
	AllowOrigin []string `protobuf:"bytes,1,rep,name=allow_origin,json=allowOrigin,proto3" json:"allow_origin,omitempty"`
	// Specifies regex patterns that match origins that will be allowed to make
	// CORS requests.
	//
	// An origin is allowed if either allow_origin or allow_origin_regex match.
	AllowOriginRegex []string `protobuf:"bytes,2,rep,name=allow_origin_regex,json=allowOriginRegex,proto3" json:"allow_origin_regex,omitempty"`
	// Specifies the content for the *access-control-allow-methods* header.
	AllowMethods []string `protobuf:"bytes,3,rep,name=allow_methods,json=allowMethods,proto3" json:"allow_methods,omitempty"`
	// Specifies the content for the *access-control-allow-headers* header.
	AllowHeaders []string `protobuf:"bytes,4,rep,name=allow_headers,json=allowHeaders,proto3" json:"allow_headers,omitempty"`
	// Specifies the content for the *access-control-expose-headers* header.
	ExposeHeaders []string `protobuf:"bytes,5,rep,name=expose_headers,json=exposeHeaders,proto3" json:"expose_headers,omitempty"`
	// Specifies the content for the *access-control-max-age* header.
	MaxAge string `protobuf:"bytes,6,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// Specifies whether the resource allows credentials.
	AllowCredentials bool `protobuf:"varint,7,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
	// Optional, only applies to route-specific CORS Policies, defaults to false.
	// If set, the CORS Policy (specified on the virtual host) will be disabled for this route.
	DisableForRoute bool `protobuf:"varint,8,opt,name=disable_for_route,json=disableForRoute,proto3" json:"disable_for_route,omitempty"`
}

func (x *CorsPolicy) Reset() {
	*x = CorsPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CorsPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CorsPolicy) ProtoMessage() {}

func (x *CorsPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CorsPolicy.ProtoReflect.Descriptor instead.
func (*CorsPolicy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescGZIP(), []int{0}
}

func (x *CorsPolicy) GetAllowOrigin() []string {
	if x != nil {
		return x.AllowOrigin
	}
	return nil
}

func (x *CorsPolicy) GetAllowOriginRegex() []string {
	if x != nil {
		return x.AllowOriginRegex
	}
	return nil
}

func (x *CorsPolicy) GetAllowMethods() []string {
	if x != nil {
		return x.AllowMethods
	}
	return nil
}

func (x *CorsPolicy) GetAllowHeaders() []string {
	if x != nil {
		return x.AllowHeaders
	}
	return nil
}

func (x *CorsPolicy) GetExposeHeaders() []string {
	if x != nil {
		return x.ExposeHeaders
	}
	return nil
}

func (x *CorsPolicy) GetMaxAge() string {
	if x != nil {
		return x.MaxAge
	}
	return ""
}

func (x *CorsPolicy) GetAllowCredentials() bool {
	if x != nil {
		return x.AllowCredentials
	}
	return false
}

func (x *CorsPolicy) GetDisableForRoute() bool {
	if x != nil {
		return x.DisableForRoute
	}
	return false
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x72, 0x73, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x72, 0x73, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6f, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x4b, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_goTypes = []interface{}{
	(*CorsPolicy)(nil), // 0: cors.options.gloo.solo.io.CorsPolicy
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CorsPolicy); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_cors_cors_proto_depIdxs = nil
}
