// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/v3/semantic_version.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
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

// Envoy uses SemVer (https://semver.org/). Major/minor versions indicate
// expected behaviors and APIs, the patch version field is used only
// for security fixes and can be generally ignored.
type SemanticVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MajorNumber uint32 `protobuf:"varint,1,opt,name=major_number,json=majorNumber,proto3" json:"major_number,omitempty"`
	MinorNumber uint32 `protobuf:"varint,2,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Patch       uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *SemanticVersion) Reset() {
	*x = SemanticVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemanticVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemanticVersion) ProtoMessage() {}

func (x *SemanticVersion) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemanticVersion.ProtoReflect.Descriptor instead.
func (*SemanticVersion) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescGZIP(), []int{0}
}

func (x *SemanticVersion) GetMajorNumber() uint32 {
	if x != nil {
		return x.MajorNumber
	}
	return 0
}

func (x *SemanticVersion) GetMinorNumber() uint32 {
	if x != nil {
		return x.MinorNumber
	}
	return 0
}

func (x *SemanticVersion) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDesc = []byte{
	0x0a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6d, 0x61, 0x6e,
	0x74, 0x69, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61,
	0x6a, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x22, 0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x1c, 0x0a, 0x1a,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e,
	0x74, 0x69, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x97, 0x01, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10,
	0x02, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x14, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_goTypes = []interface{}{
	(*SemanticVersion)(nil), // 0: solo.io.envoy.type.v3.SemanticVersion
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemanticVersion); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_v3_semantic_version_proto_depIdxs = nil
}
