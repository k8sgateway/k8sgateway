// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.8
// source: github.com/solo-io/gloo/projects/ingress/api/v1/service.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A simple wrapper for a Kubernetes Service Object.
type KubeService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a raw byte representation of the kubernetes service this resource wraps
	KubeServiceSpec *anypb.Any `protobuf:"bytes,1,opt,name=kube_service_spec,json=kubeServiceSpec,proto3" json:"kube_service_spec,omitempty"`
	// a raw byte representation of the service status of the kubernetes service object
	KubeServiceStatus *anypb.Any `protobuf:"bytes,2,opt,name=kube_service_status,json=kubeServiceStatus,proto3" json:"kube_service_status,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata *core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *KubeService) Reset() {
	*x = KubeService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeService) ProtoMessage() {}

func (x *KubeService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeService.ProtoReflect.Descriptor instead.
func (*KubeService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *KubeService) GetKubeServiceSpec() *anypb.Any {
	if x != nil {
		return x.KubeServiceSpec
	}
	return nil
}

func (x *KubeService) GetKubeServiceStatus() *anypb.Any {
	if x != nil {
		return x.KubeServiceStatus
	}
	return nil
}

func (x *KubeService) GetMetadata() *core.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x4b, 0x75, 0x62, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0f, 0x6b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x44, 0x0a, 0x13, 0x6b, 0x75, 0x62, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x11, 0x6b, 0x75, 0x62, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x3a, 0x12, 0x82, 0xf1, 0x04, 0x0e, 0x0a, 0x02, 0x73, 0x76, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x41, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescData = file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_goTypes = []any{
	(*KubeService)(nil),   // 0: ingress.solo.io.KubeService
	(*anypb.Any)(nil),     // 1: google.protobuf.Any
	(*core.Metadata)(nil), // 2: core.solo.io.Metadata
}
var file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_depIdxs = []int32{
	1, // 0: ingress.solo.io.KubeService.kube_service_spec:type_name -> google.protobuf.Any
	1, // 1: ingress.solo.io.KubeService.kube_service_status:type_name -> google.protobuf.Any
	2, // 2: ingress.solo.io.KubeService.metadata:type_name -> core.solo.io.Metadata
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_init() }
func file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_init() {
	if File_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KubeService); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto = out.File
	file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_ingress_api_v1_service_proto_depIdxs = nil
}
