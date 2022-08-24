// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformers/xslt/xslt_transformer.proto

package xslt

import (
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

// Defines an XSLT Transformation.
type XsltTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// XSLT transformation template which you want to transform requests/responses with.
	// Invalid XSLT transformation templates will result will result in a NACK during envoy configuration-time and the configuration will not be loaded.
	Xslt string `protobuf:"bytes,1,opt,name=xslt,proto3" json:"xslt,omitempty"`
	// Changes the content-type header of the HTTP request/response to what is set here.
	// This is useful in situations where an XSLT transformation is used to transform XML to JSON and the content-type
	// should be changed from `application/xml` to `application/json`.
	// If left empty, the content-type header remains unmodified by default.
	SetContentType string `protobuf:"bytes,2,opt,name=set_content_type,json=setContentType,proto3" json:"set_content_type,omitempty"`
	// This should be set to true if the content being transformed is not XML.
	// For example, if the content being transformed is from JSON to XML, this should be set to true.
	// XSLT transformations can only take valid XML as input to be transformed. If the body is not a valid XML
	// (e.g. using JSON as input in a JSON-to-XML transformation), setting `non_xml_transform` to true will allow the
	// XSLT to accept the non-XML input without throwing an error by passing the input as XML CDATA.
	// defaults to false.
	NonXmlTransform bool `protobuf:"varint,3,opt,name=non_xml_transform,json=nonXmlTransform,proto3" json:"non_xml_transform,omitempty"`
}

func (x *XsltTransformation) Reset() {
	*x = XsltTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XsltTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XsltTransformation) ProtoMessage() {}

func (x *XsltTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XsltTransformation.ProtoReflect.Descriptor instead.
func (*XsltTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescGZIP(), []int{0}
}

func (x *XsltTransformation) GetXslt() string {
	if x != nil {
		return x.Xslt
	}
	return ""
}

func (x *XsltTransformation) GetSetContentType() string {
	if x != nil {
		return x.SetContentType
	}
	return ""
}

func (x *XsltTransformation) GetNonXmlTransform() bool {
	if x != nil {
		return x.NonXmlTransform
	}
	return false
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDesc = []byte{
	0x0a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x78, 0x73, 0x6c, 0x74, 0x2f, 0x78, 0x73, 0x6c, 0x74, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x78, 0x73, 0x6c, 0x74, 0x2e, 0x76, 0x32,
	0x22, 0x7e, 0x0a, 0x12, 0x58, 0x73, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x73, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x78, 0x73, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x6f, 0x6e, 0x5f, 0x78, 0x6d, 0x6c, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x6e, 0x6f, 0x6e, 0x58, 0x6d, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xa3, 0x01, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x78, 0x73, 0x6c, 0x74,
	0x2e, 0x76, 0x32, 0x42, 0x14, 0x58, 0x73, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x78, 0x73, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_goTypes = []interface{}{
	(*XsltTransformation)(nil), // 0: envoy.config.transformer.xslt.v2.XsltTransformation
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XsltTransformation); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformers_xslt_xslt_transformer_proto_depIdxs = nil
}
