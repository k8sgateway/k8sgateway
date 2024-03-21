// copied from https://github.com/envoyproxy/envoy/blob/ad89a587aa0177bfdad6b5c968a6aead5d9be7a4/api/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto

// manually updated pkg

package v3

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Control how headers and trailers are handled
type ProcessingMode_HeaderSendMode int32

const (
	// The default HeaderSendMode depends on which part of the message is being
	// processed. By default, request and response headers are sent,
	// while trailers are skipped.
	ProcessingMode_DEFAULT ProcessingMode_HeaderSendMode = 0
	// Send the header or trailer.
	ProcessingMode_SEND ProcessingMode_HeaderSendMode = 1
	// Do not send the header or trailer.
	ProcessingMode_SKIP ProcessingMode_HeaderSendMode = 2
)

// Enum value maps for ProcessingMode_HeaderSendMode.
var (
	ProcessingMode_HeaderSendMode_name = map[int32]string{
		0: "DEFAULT",
		1: "SEND",
		2: "SKIP",
	}
	ProcessingMode_HeaderSendMode_value = map[string]int32{
		"DEFAULT": 0,
		"SEND":    1,
		"SKIP":    2,
	}
)

func (x ProcessingMode_HeaderSendMode) Enum() *ProcessingMode_HeaderSendMode {
	p := new(ProcessingMode_HeaderSendMode)
	*p = x
	return p
}

func (x ProcessingMode_HeaderSendMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessingMode_HeaderSendMode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_enumTypes[0].Descriptor()
}

func (ProcessingMode_HeaderSendMode) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_enumTypes[0]
}

func (x ProcessingMode_HeaderSendMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessingMode_HeaderSendMode.Descriptor instead.
func (ProcessingMode_HeaderSendMode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescGZIP(), []int{0, 0}
}

// Control how the request and response bodies are handled
type ProcessingMode_BodySendMode int32

const (
	// Do not send the body at all. This is the default.
	ProcessingMode_NONE ProcessingMode_BodySendMode = 0
	// Stream the body to the server in pieces as they arrive at the
	// proxy.
	ProcessingMode_STREAMED ProcessingMode_BodySendMode = 1
	// Buffer the message body in memory and send the entire body at once.
	// If the body exceeds the configured buffer limit, then the
	// downstream system will receive an error.
	ProcessingMode_BUFFERED ProcessingMode_BodySendMode = 2
	// Buffer the message body in memory and send the entire body in one
	// chunk. If the body exceeds the configured buffer limit, then the body contents
	// up to the buffer limit will be sent.
	ProcessingMode_BUFFERED_PARTIAL ProcessingMode_BodySendMode = 3
)

// Enum value maps for ProcessingMode_BodySendMode.
var (
	ProcessingMode_BodySendMode_name = map[int32]string{
		0: "NONE",
		1: "STREAMED",
		2: "BUFFERED",
		3: "BUFFERED_PARTIAL",
	}
	ProcessingMode_BodySendMode_value = map[string]int32{
		"NONE":             0,
		"STREAMED":         1,
		"BUFFERED":         2,
		"BUFFERED_PARTIAL": 3,
	}
)

func (x ProcessingMode_BodySendMode) Enum() *ProcessingMode_BodySendMode {
	p := new(ProcessingMode_BodySendMode)
	*p = x
	return p
}

func (x ProcessingMode_BodySendMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessingMode_BodySendMode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_enumTypes[1].Descriptor()
}

func (ProcessingMode_BodySendMode) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_enumTypes[1]
}

func (x ProcessingMode_BodySendMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessingMode_BodySendMode.Descriptor instead.
func (ProcessingMode_BodySendMode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescGZIP(), []int{0, 1}
}

// [#next-free-field: 7]
type ProcessingMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How to handle the request header. Default is "SEND".
	RequestHeaderMode ProcessingMode_HeaderSendMode `protobuf:"varint,1,opt,name=request_header_mode,json=requestHeaderMode,proto3,enum=solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode_HeaderSendMode" json:"request_header_mode,omitempty"`
	// How to handle the response header. Default is "SEND".
	ResponseHeaderMode ProcessingMode_HeaderSendMode `protobuf:"varint,2,opt,name=response_header_mode,json=responseHeaderMode,proto3,enum=solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode_HeaderSendMode" json:"response_header_mode,omitempty"`
	// How to handle the request body. Default is "NONE".
	RequestBodyMode ProcessingMode_BodySendMode `protobuf:"varint,3,opt,name=request_body_mode,json=requestBodyMode,proto3,enum=solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode_BodySendMode" json:"request_body_mode,omitempty"`
	// How do handle the response body. Default is "NONE".
	ResponseBodyMode ProcessingMode_BodySendMode `protobuf:"varint,4,opt,name=response_body_mode,json=responseBodyMode,proto3,enum=solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode_BodySendMode" json:"response_body_mode,omitempty"`
	// How to handle the request trailers. Default is "SKIP".
	RequestTrailerMode ProcessingMode_HeaderSendMode `protobuf:"varint,5,opt,name=request_trailer_mode,json=requestTrailerMode,proto3,enum=solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode_HeaderSendMode" json:"request_trailer_mode,omitempty"`
	// How to handle the response trailers. Default is "SKIP".
	ResponseTrailerMode ProcessingMode_HeaderSendMode `protobuf:"varint,6,opt,name=response_trailer_mode,json=responseTrailerMode,proto3,enum=solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode_HeaderSendMode" json:"response_trailer_mode,omitempty"`
}

func (x *ProcessingMode) Reset() {
	*x = ProcessingMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingMode) ProtoMessage() {}

func (x *ProcessingMode) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingMode.ProtoReflect.Descriptor instead.
func (*ProcessingMode) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessingMode) GetRequestHeaderMode() ProcessingMode_HeaderSendMode {
	if x != nil {
		return x.RequestHeaderMode
	}
	return ProcessingMode_DEFAULT
}

func (x *ProcessingMode) GetResponseHeaderMode() ProcessingMode_HeaderSendMode {
	if x != nil {
		return x.ResponseHeaderMode
	}
	return ProcessingMode_DEFAULT
}

func (x *ProcessingMode) GetRequestBodyMode() ProcessingMode_BodySendMode {
	if x != nil {
		return x.RequestBodyMode
	}
	return ProcessingMode_NONE
}

func (x *ProcessingMode) GetResponseBodyMode() ProcessingMode_BodySendMode {
	if x != nil {
		return x.ResponseBodyMode
	}
	return ProcessingMode_NONE
}

func (x *ProcessingMode) GetRequestTrailerMode() ProcessingMode_HeaderSendMode {
	if x != nil {
		return x.RequestTrailerMode
	}
	return ProcessingMode_DEFAULT
}

func (x *ProcessingMode) GetResponseTrailerMode() ProcessingMode_HeaderSendMode {
	if x != nil {
		return x.ResponseTrailerMode
	}
	return ProcessingMode_DEFAULT
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDesc = []byte{
	0x0a, 0x72, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x07, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x50, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62,
	0x6f, 0x64, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e,
	0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x6f, 0x64, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x50, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x8e, 0x01, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x50, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x22, 0x31, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x4b, 0x49, 0x50, 0x10, 0x02, 0x22, 0x4a, 0x0a, 0x0c, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x42,
	0x55, 0x46, 0x46, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41, 0x4c, 0x10,
	0x03, 0x42, 0xc6, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01,
	0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76,
	0x33, 0x42, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_goTypes = []interface{}{
	(ProcessingMode_HeaderSendMode)(0), // 0: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
	(ProcessingMode_BodySendMode)(0),   // 1: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode
	(*ProcessingMode)(nil),             // 2: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_depIdxs = []int32{
	0, // 0: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.request_header_mode:type_name -> solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
	0, // 1: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.response_header_mode:type_name -> solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
	1, // 2: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.request_body_mode:type_name -> solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode
	1, // 3: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.response_body_mode:type_name -> solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode
	0, // 4: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.request_trailer_mode:type_name -> solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
	0, // 5: solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.response_trailer_mode:type_name -> solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingMode); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_ext_proc_v3_processing_mode_proto_depIdxs = nil
}
