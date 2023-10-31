// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc_web/grpc_web.proto

package grpc_web

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
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

// GrpcWeb support is enabled be default. Use this extension to disable it.
type GrpcWeb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Disable grpc web support.
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (x *GrpcWeb) Reset() {
	*x = GrpcWeb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcWeb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcWeb) ProtoMessage() {}

func (x *GrpcWeb) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcWeb.ProtoReflect.Descriptor instead.
func (*GrpcWeb) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcWeb) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65,
	0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x47, 0x72, 0x70, 0x63, 0x57,
	0x65, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x4f, 0xb8, 0xf5,
	0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x77, 0x65, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_goTypes = []interface{}{
	(*GrpcWeb)(nil), // 0: grpc_web.options.gloo.solo.io.GrpcWeb
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcWeb); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_grpc_web_grpc_web_proto_depIdxs = nil
}
