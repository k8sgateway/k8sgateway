// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/azure/azure.proto

package azure

import (
	reflect "reflect"
	sync "sync"

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

type UpstreamSpec_FunctionSpec_AuthLevel int32

const (
	UpstreamSpec_FunctionSpec_Anonymous UpstreamSpec_FunctionSpec_AuthLevel = 0
	UpstreamSpec_FunctionSpec_Function  UpstreamSpec_FunctionSpec_AuthLevel = 1
	UpstreamSpec_FunctionSpec_Admin     UpstreamSpec_FunctionSpec_AuthLevel = 2
)

// Enum value maps for UpstreamSpec_FunctionSpec_AuthLevel.
var (
	UpstreamSpec_FunctionSpec_AuthLevel_name = map[int32]string{
		0: "Anonymous",
		1: "Function",
		2: "Admin",
	}
	UpstreamSpec_FunctionSpec_AuthLevel_value = map[string]int32{
		"Anonymous": 0,
		"Function":  1,
		"Admin":     2,
	}
)

func (x UpstreamSpec_FunctionSpec_AuthLevel) Enum() *UpstreamSpec_FunctionSpec_AuthLevel {
	p := new(UpstreamSpec_FunctionSpec_AuthLevel)
	*p = x
	return p
}

func (x UpstreamSpec_FunctionSpec_AuthLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpstreamSpec_FunctionSpec_AuthLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_enumTypes[0].Descriptor()
}

func (UpstreamSpec_FunctionSpec_AuthLevel) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_enumTypes[0]
}

func (x UpstreamSpec_FunctionSpec_AuthLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpstreamSpec_FunctionSpec_AuthLevel.Descriptor instead.
func (UpstreamSpec_FunctionSpec_AuthLevel) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Upstream Spec for Azure Functions Upstreams
// Azure Upstreams represent a collection of Azure Functions for a particular Azure Account
// within a particular Function App
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Name of the Azure Function App where the functions are grouped
	FunctionAppName string `protobuf:"bytes,1,opt,name=function_app_name,json=functionAppName,proto3" json:"function_app_name,omitempty"`
	// A [Gloo Secret Ref](https://docs.solo.io/gloo-edge/latest/reference/cli/glooctl_create_secret_azure/) to an [Azure Publish Profile JSON file](https://azure.microsoft.com/en-us/downloads/publishing-profile-overview/).
	// Note that this secret is not required unless Function Discovery is enabled
	SecretRef *core.ResourceRef            `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	Functions []*UpstreamSpec_FunctionSpec `protobuf:"bytes,3,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSpec.ProtoReflect.Descriptor instead.
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetFunctionAppName() string {
	if x != nil {
		return x.FunctionAppName
	}
	return ""
}

func (x *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if x != nil {
		return x.SecretRef
	}
	return nil
}

func (x *UpstreamSpec) GetFunctions() []*UpstreamSpec_FunctionSpec {
	if x != nil {
		return x.Functions
	}
	return nil
}

type DestinationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Function Name of the FunctionSpec to be invoked.
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (x *DestinationSpec) Reset() {
	*x = DestinationSpec{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DestinationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationSpec) ProtoMessage() {}

func (x *DestinationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes[1]
	if x != nil {
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
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescGZIP(), []int{1}
}

func (x *DestinationSpec) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

// Function Spec for Functions on Azure Functions Upstreams
// The Function Spec contains data necessary for Gloo to invoke Azure functions
type UpstreamSpec_FunctionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Name of the Azure Function as it appears in the Azure Functions Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// Auth Level can bve either "anonymous" "function" or "admin"
	// See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details
	AuthLevel UpstreamSpec_FunctionSpec_AuthLevel `protobuf:"varint,2,opt,name=auth_level,json=authLevel,proto3,enum=azure.options.gloo.solo.io.UpstreamSpec_FunctionSpec_AuthLevel" json:"auth_level,omitempty"`
}

func (x *UpstreamSpec_FunctionSpec) Reset() {
	*x = UpstreamSpec_FunctionSpec{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpstreamSpec_FunctionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec_FunctionSpec) ProtoMessage() {}

func (x *UpstreamSpec_FunctionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSpec_FunctionSpec.ProtoReflect.Descriptor instead.
func (*UpstreamSpec_FunctionSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UpstreamSpec_FunctionSpec) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *UpstreamSpec_FunctionSpec) GetAuthLevel() UpstreamSpec_FunctionSpec_AuthLevel {
	if x != nil {
		return x.AuthLevel
	}
	return UpstreamSpec_FunctionSpec_Anonymous
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x53, 0x0a,
	0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0xc8, 0x01, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x33, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x02, 0x22, 0x36, 0x0a,
	0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x4c, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_goTypes = []any{
	(UpstreamSpec_FunctionSpec_AuthLevel)(0), // 0: azure.options.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel
	(*UpstreamSpec)(nil),                     // 1: azure.options.gloo.solo.io.UpstreamSpec
	(*DestinationSpec)(nil),                  // 2: azure.options.gloo.solo.io.DestinationSpec
	(*UpstreamSpec_FunctionSpec)(nil),        // 3: azure.options.gloo.solo.io.UpstreamSpec.FunctionSpec
	(*core.ResourceRef)(nil),                 // 4: core.solo.io.ResourceRef
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_depIdxs = []int32{
	4, // 0: azure.options.gloo.solo.io.UpstreamSpec.secret_ref:type_name -> core.solo.io.ResourceRef
	3, // 1: azure.options.gloo.solo.io.UpstreamSpec.functions:type_name -> azure.options.gloo.solo.io.UpstreamSpec.FunctionSpec
	0, // 2: azure.options.gloo.solo.io.UpstreamSpec.FunctionSpec.auth_level:type_name -> azure.options.gloo.solo.io.UpstreamSpec.FunctionSpec.AuthLevel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_azure_azure_proto_depIdxs = nil
}
