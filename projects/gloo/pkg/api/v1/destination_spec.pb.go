// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.8
// source: github.com/solo-io/gloo/projects/gloo/api/v1/destination_spec.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"
	azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"
	grpc "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/grpc"
	rest "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/rest"
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

// Configuration for Destinations that are tied to the UpstreamSpec or ServiceSpec on that destination
type DestinationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Note to developers: new DestinationSpecs must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are assignable to DestinationType:
	//
	//	*DestinationSpec_Aws
	//	*DestinationSpec_Azure
	//	*DestinationSpec_Rest
	//	*DestinationSpec_Grpc
	DestinationType isDestinationSpec_DestinationType `protobuf_oneof:"destination_type"`
}

func (x *DestinationSpec) Reset() {
	*x = DestinationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationSpec) ProtoMessage() {}

func (x *DestinationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationSpec.ProtoReflect.Descriptor instead.
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescGZIP(), []int{0}
}

func (m *DestinationSpec) GetDestinationType() isDestinationSpec_DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return nil
}

func (x *DestinationSpec) GetAws() *aws.DestinationSpec {
	if x, ok := x.GetDestinationType().(*DestinationSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (x *DestinationSpec) GetAzure() *azure.DestinationSpec {
	if x, ok := x.GetDestinationType().(*DestinationSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

func (x *DestinationSpec) GetRest() *rest.DestinationSpec {
	if x, ok := x.GetDestinationType().(*DestinationSpec_Rest); ok {
		return x.Rest
	}
	return nil
}

func (x *DestinationSpec) GetGrpc() *grpc.DestinationSpec {
	if x, ok := x.GetDestinationType().(*DestinationSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

type isDestinationSpec_DestinationType interface {
	isDestinationSpec_DestinationType()
}

type DestinationSpec_Aws struct {
	Aws *aws.DestinationSpec `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}

type DestinationSpec_Azure struct {
	Azure *azure.DestinationSpec `protobuf:"bytes,2,opt,name=azure,proto3,oneof"`
}

type DestinationSpec_Rest struct {
	Rest *rest.DestinationSpec `protobuf:"bytes,3,opt,name=rest,proto3,oneof"`
}

type DestinationSpec_Grpc struct {
	Grpc *grpc.DestinationSpec `protobuf:"bytes,4,opt,name=grpc,proto3,oneof"`
}

func (*DestinationSpec_Aws) isDestinationSpec_DestinationType() {}

func (*DestinationSpec_Azure) isDestinationSpec_DestinationType() {}

func (*DestinationSpec_Rest) isDestinationSpec_DestinationType() {}

func (*DestinationSpec_Grpc) isDestinationSpec_DestinationType() {}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x77,
	0x73, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x03, 0x61,
	0x77, 0x73, 0x12, 0x43, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00,
	0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x48, 0x00, 0x52, 0x04, 0x72, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x42, 0x12, 0x0a, 0x10, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x3e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_goTypes = []any{
	(*DestinationSpec)(nil),       // 0: gloo.solo.io.DestinationSpec
	(*aws.DestinationSpec)(nil),   // 1: aws.options.gloo.solo.io.DestinationSpec
	(*azure.DestinationSpec)(nil), // 2: azure.options.gloo.solo.io.DestinationSpec
	(*rest.DestinationSpec)(nil),  // 3: rest.options.gloo.solo.io.DestinationSpec
	(*grpc.DestinationSpec)(nil),  // 4: grpc.options.gloo.solo.io.DestinationSpec
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_depIdxs = []int32{
	1, // 0: gloo.solo.io.DestinationSpec.aws:type_name -> aws.options.gloo.solo.io.DestinationSpec
	2, // 1: gloo.solo.io.DestinationSpec.azure:type_name -> azure.options.gloo.solo.io.DestinationSpec
	3, // 2: gloo.solo.io.DestinationSpec.rest:type_name -> rest.options.gloo.solo.io.DestinationSpec
	4, // 3: gloo.solo.io.DestinationSpec.grpc:type_name -> grpc.options.gloo.solo.io.DestinationSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DestinationSpec); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_msgTypes[0].OneofWrappers = []any{
		(*DestinationSpec_Aws)(nil),
		(*DestinationSpec_Azure)(nil),
		(*DestinationSpec_Rest)(nil),
		(*DestinationSpec_Grpc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_destination_spec_proto_depIdxs = nil
}
