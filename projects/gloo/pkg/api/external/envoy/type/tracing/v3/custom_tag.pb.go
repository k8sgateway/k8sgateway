// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/tracing/v3/custom_tag.proto

package v3

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/metadata/v3"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Describes custom tags for the active span.
// [#next-free-field: 6]
type CustomTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to populate the tag name.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Used to specify what kind of custom tag.
	//
	// Types that are assignable to Type:
	//	*CustomTag_Literal_
	//	*CustomTag_Environment_
	//	*CustomTag_RequestHeader
	//	*CustomTag_Metadata_
	Type isCustomTag_Type `protobuf_oneof:"type"`
}

func (x *CustomTag) Reset() {
	*x = CustomTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTag) ProtoMessage() {}

func (x *CustomTag) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTag.ProtoReflect.Descriptor instead.
func (*CustomTag) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescGZIP(), []int{0}
}

func (x *CustomTag) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (m *CustomTag) GetType() isCustomTag_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *CustomTag) GetLiteral() *CustomTag_Literal {
	if x, ok := x.GetType().(*CustomTag_Literal_); ok {
		return x.Literal
	}
	return nil
}

func (x *CustomTag) GetEnvironment() *CustomTag_Environment {
	if x, ok := x.GetType().(*CustomTag_Environment_); ok {
		return x.Environment
	}
	return nil
}

func (x *CustomTag) GetRequestHeader() *CustomTag_Header {
	if x, ok := x.GetType().(*CustomTag_RequestHeader); ok {
		return x.RequestHeader
	}
	return nil
}

func (x *CustomTag) GetMetadata() *CustomTag_Metadata {
	if x, ok := x.GetType().(*CustomTag_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

type isCustomTag_Type interface {
	isCustomTag_Type()
}

type CustomTag_Literal_ struct {
	// A literal custom tag.
	Literal *CustomTag_Literal `protobuf:"bytes,2,opt,name=literal,proto3,oneof"`
}

type CustomTag_Environment_ struct {
	// An environment custom tag.
	Environment *CustomTag_Environment `protobuf:"bytes,3,opt,name=environment,proto3,oneof"`
}

type CustomTag_RequestHeader struct {
	// A request header custom tag.
	RequestHeader *CustomTag_Header `protobuf:"bytes,4,opt,name=request_header,json=requestHeader,proto3,oneof"`
}

type CustomTag_Metadata_ struct {
	// A custom tag to obtain tag value from the metadata.
	Metadata *CustomTag_Metadata `protobuf:"bytes,5,opt,name=metadata,proto3,oneof"`
}

func (*CustomTag_Literal_) isCustomTag_Type() {}

func (*CustomTag_Environment_) isCustomTag_Type() {}

func (*CustomTag_RequestHeader) isCustomTag_Type() {}

func (*CustomTag_Metadata_) isCustomTag_Type() {}

// Literal type custom tag with static value for the tag value.
type CustomTag_Literal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Static literal value to populate the tag value.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CustomTag_Literal) Reset() {
	*x = CustomTag_Literal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTag_Literal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTag_Literal) ProtoMessage() {}

func (x *CustomTag_Literal) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTag_Literal.ProtoReflect.Descriptor instead.
func (*CustomTag_Literal) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CustomTag_Literal) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Environment type custom tag with environment name and default value.
type CustomTag_Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Environment variable name to obtain the value to populate the tag value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the environment variable is not found,
	// the tag value will be populated with this default value if specified,
	// otherwise no tag will be populated.
	DefaultValue string `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (x *CustomTag_Environment) Reset() {
	*x = CustomTag_Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTag_Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTag_Environment) ProtoMessage() {}

func (x *CustomTag_Environment) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTag_Environment.ProtoReflect.Descriptor instead.
func (*CustomTag_Environment) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CustomTag_Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomTag_Environment) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

// Header type custom tag with header name and default value.
type CustomTag_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name to obtain the value to populate the tag value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the header does not exist,
	// the tag value will be populated with this default value if specified,
	// otherwise no tag will be populated.
	DefaultValue string `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (x *CustomTag_Header) Reset() {
	*x = CustomTag_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTag_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTag_Header) ProtoMessage() {}

func (x *CustomTag_Header) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTag_Header.ProtoReflect.Descriptor instead.
func (*CustomTag_Header) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescGZIP(), []int{0, 2}
}

func (x *CustomTag_Header) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomTag_Header) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

// Metadata type custom tag using
// :ref:`MetadataKey <envoy_api_msg_type.metadata.v3.MetadataKey>` to retrieve the protobuf value
// from :ref:`Metadata <envoy_api_msg_config.core.v3.Metadata>`, and populate the tag value with
// `the canonical JSON <https://developers.google.com/protocol-buffers/docs/proto3#json>`_
// representation of it.
type CustomTag_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify what kind of metadata to obtain tag value from.
	Kind *v3.MetadataKind `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata key to define the path to retrieve the tag value.
	MetadataKey *v3.MetadataKey `protobuf:"bytes,2,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	// When no valid metadata is found,
	// the tag value would be populated with this default value if specified,
	// otherwise no tag would be populated.
	DefaultValue string `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (x *CustomTag_Metadata) Reset() {
	*x = CustomTag_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTag_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTag_Metadata) ProtoMessage() {}

func (x *CustomTag_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTag_Metadata.ProtoReflect.Descriptor instead.
func (*CustomTag_Metadata) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescGZIP(), []int{0, 3}
}

func (x *CustomTag_Metadata) GetKind() *v3.MetadataKind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *CustomTag_Metadata) GetMetadataKey() *v3.MetadataKey {
	if x != nil {
		return x.MetadataKey
	}
	return nil
}

func (x *CustomTag_Metadata) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDesc = []byte{
	0x0a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x08, 0x0a, 0x09, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61,
	0x67, 0x12, 0x19, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x4c, 0x0a, 0x07,
	0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x58, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4f,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x59, 0x0a, 0x07, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x2f, 0x8a, 0xc8, 0xde, 0x8e, 0x04,
	0x29, 0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54,
	0x61, 0x67, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x1a, 0x84, 0x01, 0x0a, 0x0b, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x33, 0x8a, 0xc8,
	0xde, 0x8e, 0x04, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x54, 0x61, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x80, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72,
	0x08, 0x20, 0x01, 0xc0, 0x01, 0x01, 0xc8, 0x01, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x2e, 0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x28, 0x0a, 0x26, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x1a, 0xf3, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x40, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x4e, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x30, 0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x2a,
	0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61,
	0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x27, 0x8a, 0xc8, 0xde, 0x8e,
	0x04, 0x21, 0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x54, 0x61, 0x67, 0x42, 0x0b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x42, 0xa1, 0x01, 0x0a, 0x2b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33,
	0x42, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33,
	0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01,
	0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_goTypes = []interface{}{
	(*CustomTag)(nil),             // 0: solo.io.envoy.type.tracing.v3.CustomTag
	(*CustomTag_Literal)(nil),     // 1: solo.io.envoy.type.tracing.v3.CustomTag.Literal
	(*CustomTag_Environment)(nil), // 2: solo.io.envoy.type.tracing.v3.CustomTag.Environment
	(*CustomTag_Header)(nil),      // 3: solo.io.envoy.type.tracing.v3.CustomTag.Header
	(*CustomTag_Metadata)(nil),    // 4: solo.io.envoy.type.tracing.v3.CustomTag.Metadata
	(*v3.MetadataKind)(nil),       // 5: solo.io.envoy.type.metadata.v3.MetadataKind
	(*v3.MetadataKey)(nil),        // 6: solo.io.envoy.type.metadata.v3.MetadataKey
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_depIdxs = []int32{
	1, // 0: solo.io.envoy.type.tracing.v3.CustomTag.literal:type_name -> solo.io.envoy.type.tracing.v3.CustomTag.Literal
	2, // 1: solo.io.envoy.type.tracing.v3.CustomTag.environment:type_name -> solo.io.envoy.type.tracing.v3.CustomTag.Environment
	3, // 2: solo.io.envoy.type.tracing.v3.CustomTag.request_header:type_name -> solo.io.envoy.type.tracing.v3.CustomTag.Header
	4, // 3: solo.io.envoy.type.tracing.v3.CustomTag.metadata:type_name -> solo.io.envoy.type.tracing.v3.CustomTag.Metadata
	5, // 4: solo.io.envoy.type.tracing.v3.CustomTag.Metadata.kind:type_name -> solo.io.envoy.type.metadata.v3.MetadataKind
	6, // 5: solo.io.envoy.type.tracing.v3.CustomTag.Metadata.metadata_key:type_name -> solo.io.envoy.type.metadata.v3.MetadataKey
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTag); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTag_Literal); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTag_Environment); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTag_Header); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTag_Metadata); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CustomTag_Literal_)(nil),
		(*CustomTag_Environment_)(nil),
		(*CustomTag_RequestHeader)(nil),
		(*CustomTag_Metadata_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_type_tracing_v3_custom_tag_proto_depIdxs = nil
}
