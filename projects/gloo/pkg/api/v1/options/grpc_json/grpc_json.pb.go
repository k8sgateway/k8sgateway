// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc_json/grpc_json.proto

package grpc_json

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// [#next-free-field: 10]
type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	// Whether to route methods without the ``google.api.http`` option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     package bookstore;
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The client could ``post`` a json body ``{"shelf": 1234}`` with the path of
	// ``/bookstore.Bookstore/GetShelfRequest`` to call ``GetShelfRequest``.
	AutoMapping bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use ``ignored_query_parameters``.
	// Defaults to false.
	IgnoreUnknownQueryParameters bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take ``google.rpc.Status``
	// from the ``grpc-status-details-bin`` header and use it as JSON body.
	// If there was no such header, make ``google.rpc.Status`` out of the ``grpc-status`` and
	// ``grpc-message`` headers.
	// The error details types must be present in the ``proto_descriptor``.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//     grpc-status: 5
	//     grpc-status-details-bin:
	//         CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The ``grpc-status-details-bin`` header contains a base64-encoded protobuf message
	// ``google.rpc.Status``. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//     HTTP/1.1 404 Not Found
	//     content-type: application/json
	//
	//     {"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	//  In order to transcode the message, the ``google.rpc.RequestInfo`` type from
	//  the ``google/rpc/error_details.proto`` should be included in the configured
	//  :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus    bool     `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa94c5eabc74996, []int{0}
}
func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
	Equal(interface{}) bool
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof" json:"proto_descriptor,omitempty"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin string `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof" json:"proto_descriptor_bin,omitempty"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa94c5eabc74996, []int{0, 0}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "grpc_json.options.gloo.solo.io.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "grpc_json.options.gloo.solo.io.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc_json/grpc_json.proto", fileDescriptor_caa94c5eabc74996)
}

var fileDescriptor_caa94c5eabc74996 = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6f, 0xe4, 0x34,
	0x14, 0xc7, 0xc9, 0xf4, 0x57, 0xea, 0xfe, 0xc4, 0x94, 0x36, 0x8c, 0xda, 0x32, 0x45, 0x20, 0x8d,
	0x84, 0x48, 0x44, 0x7b, 0x81, 0xf6, 0x50, 0x1a, 0x68, 0x69, 0x91, 0x5a, 0x86, 0x00, 0x42, 0xe2,
	0x62, 0xb9, 0x89, 0x9b, 0x71, 0x3b, 0xb1, 0x5d, 0xdb, 0x49, 0x67, 0x6e, 0x1c, 0x39, 0x73, 0xe4,
	0x2f, 0xe0, 0xc0, 0x89, 0x13, 0xe2, 0xc4, 0x65, 0xa5, 0xbd, 0xee, 0xbf, 0xb0, 0xff, 0xc3, 0x4a,
	0xab, 0x3d, 0xad, 0x6c, 0xa7, 0x33, 0xb3, 0xad, 0xf6, 0x87, 0xf6, 0xf6, 0xf2, 0xbe, 0x9f, 0xf7,
	0xbe, 0xc9, 0xf3, 0x73, 0xc0, 0x59, 0x4e, 0x75, 0xb7, 0x3c, 0x0f, 0x53, 0x5e, 0x44, 0x8a, 0xf7,
	0xf8, 0x67, 0x94, 0x47, 0x79, 0x8f, 0xf3, 0x48, 0x48, 0x7e, 0x49, 0x52, 0xad, 0xdc, 0x13, 0x16,
	0x34, 0xaa, 0x3e, 0x8f, 0xb8, 0xd0, 0x94, 0x33, 0x15, 0xe5, 0x52, 0xa4, 0xe8, 0x52, 0x71, 0x36,
	0x8a, 0x42, 0x21, 0xb9, 0xe6, 0x70, 0x73, 0x94, 0xa8, 0xe1, 0xd0, 0x34, 0x08, 0x4d, 0xef, 0x90,
	0xf2, 0xe6, 0x4a, 0xce, 0x73, 0x6e, 0xd1, 0xc8, 0x44, 0xae, 0xaa, 0x09, 0x49, 0x5f, 0xbb, 0x24,
	0xe9, 0xeb, 0x3a, 0xb7, 0x51, 0x66, 0x02, 0x47, 0x98, 0x31, 0xae, 0xb1, 0x73, 0x55, 0x1a, 0xeb,
	0x52, 0xd5, 0xf2, 0xd6, 0x3d, 0xb9, 0x22, 0x52, 0x51, 0xce, 0x28, 0xcb, 0x6b, 0x64, 0xad, 0xc2,
	0x3d, 0x9a, 0x61, 0x4d, 0xa2, 0xdb, 0xc0, 0x09, 0x1f, 0xfd, 0x33, 0x03, 0xe0, 0xb7, 0x52, 0xa4,
	0xdf, 0x29, 0xce, 0x7e, 0x92, 0x98, 0xa9, 0x94, 0x67, 0x44, 0xc2, 0x4f, 0xc1, 0xb2, 0xd5, 0x51,
	0x46, 0x54, 0x2a, 0xa9, 0xd0, 0x5c, 0x06, 0x5e, 0xcb, 0x6b, 0xcf, 0x1e, 0xbf, 0x93, 0x2c, 0x59,
	0xe5, 0x9b, 0xa1, 0x00, 0xb7, 0xc1, 0xca, 0x5d, 0x18, 0x9d, 0x53, 0x16, 0x4c, 0xd6, 0x05, 0xf0,
	0x4e, 0x41, 0x4c, 0x19, 0xfc, 0x18, 0xf8, 0x8a, 0xc8, 0x8a, 0xa6, 0x44, 0x05, 0x8d, 0xd6, 0x44,
	0x7b, 0x36, 0xf6, 0x9f, 0xc5, 0x53, 0x7f, 0x78, 0x0d, 0xdf, 0x4b, 0x86, 0x0a, 0xcc, 0xc0, 0x82,
	0x90, 0x94, 0x69, 0x54, 0x0f, 0x30, 0x98, 0x68, 0x79, 0xed, 0xb9, 0xed, 0xfd, 0xf0, 0xd5, 0xa3,
	0x0d, 0xef, 0x7f, 0x51, 0xd8, 0x31, 0x7d, 0xbe, 0x77, 0x70, 0x32, 0x2f, 0xc6, 0x9e, 0xe0, 0x3e,
	0x58, 0x2f, 0xb0, 0x4e, 0xbb, 0x88, 0xb2, 0x94, 0x17, 0x94, 0xe5, 0x48, 0x92, 0xeb, 0x92, 0x28,
	0x8d, 0x24, 0x2f, 0x35, 0x09, 0xa6, 0x5a, 0x5e, 0xdb, 0x4f, 0x3e, 0xb0, 0xcc, 0x49, 0x8d, 0x24,
	0x8e, 0x48, 0x0c, 0x00, 0xbf, 0x00, 0x01, 0xcd, 0x19, 0x97, 0x24, 0x43, 0xd7, 0x25, 0x91, 0x03,
	0x24, 0xb0, 0xc4, 0x05, 0xd1, 0x44, 0xaa, 0x60, 0xda, 0x7c, 0x5c, 0xb2, 0x5a, 0xeb, 0x3f, 0x18,
	0xb9, 0x33, 0x54, 0xe1, 0x16, 0x98, 0xc7, 0xa5, 0xe6, 0xa8, 0xc0, 0x42, 0x50, 0x96, 0x07, 0x33,
	0xd6, 0x6a, 0xce, 0xe4, 0x4e, 0x5d, 0x0a, 0x1e, 0x82, 0x0f, 0x5d, 0x31, 0x2a, 0xd9, 0x15, 0xe3,
	0x37, 0xec, 0xbe, 0x87, 0x6f, 0xab, 0xd6, 0x1d, 0xf6, 0xb3, 0xa3, 0xee, 0x3a, 0x85, 0xe0, 0xbd,
	0x94, 0xb3, 0x8a, 0x48, 0x8d, 0xec, 0xf0, 0xdc, 0x06, 0x05, 0xb3, 0xb6, 0xf4, 0xdd, 0x5a, 0x32,
	0x73, 0xfb, 0xd1, 0x0a, 0xcd, 0xff, 0x1b, 0x60, 0x7e, 0x7c, 0x66, 0xf0, 0x13, 0xb0, 0x88, 0xb3,
	0x0c, 0xdd, 0x74, 0xa9, 0x26, 0x4a, 0xe0, 0x94, 0xd8, 0x85, 0xf0, 0x93, 0x05, 0x9c, 0x65, 0xbf,
	0x0c, 0x93, 0xf0, 0x00, 0x6c, 0xe0, 0xde, 0x0d, 0x1e, 0x28, 0xe4, 0x4e, 0x4e, 0x48, 0x5a, 0x50,
	0x4d, 0x2b, 0x82, 0x2e, 0x28, 0xe9, 0x65, 0xe6, 0xb4, 0x4d, 0x55, 0xd3, 0x41, 0xd6, 0xa1, 0x73,
	0x8b, 0x1c, 0x59, 0x02, 0xee, 0x82, 0xe6, 0x0b, 0x2d, 0x08, 0x2b, 0x0b, 0x85, 0xb0, 0x42, 0x94,
	0x69, 0xb7, 0x02, 0x7e, 0xb2, 0x3a, 0x56, 0x7f, 0x68, 0xf4, 0x03, 0x75, 0xc2, 0xb4, 0x82, 0x7b,
	0xa0, 0x29, 0x24, 0x31, 0x0b, 0x44, 0x90, 0x5b, 0x4a, 0x6b, 0x8b, 0x18, 0x2e, 0x88, 0xb2, 0x1b,
	0xe9, 0x27, 0x6b, 0xb7, 0x44, 0xc7, 0x00, 0xd6, 0xf4, 0xcc, 0xc8, 0xbb, 0xa7, 0x7f, 0x3e, 0xf8,
	0x7d, 0xf3, 0x18, 0x1c, 0x11, 0x56, 0xf1, 0x41, 0x98, 0x72, 0x76, 0x41, 0xf3, 0xf0, 0x82, 0xf6,
	0x34, 0x91, 0x61, 0x57, 0x6b, 0x11, 0xea, 0xd1, 0x3e, 0x55, 0xdb, 0xaf, 0xdb, 0xb2, 0xdd, 0xaf,
	0x4c, 0xbb, 0x3d, 0xf0, 0xe5, 0x5b, 0xb7, 0x8b, 0xdf, 0x07, 0x8b, 0x63, 0x77, 0x4a, 0x11, 0x0d,
	0x27, 0x9e, 0xc6, 0x5e, 0xfc, 0xb7, 0xf7, 0xef, 0x93, 0x49, 0xef, 0xaf, 0xc7, 0x9b, 0xde, 0x7f,
	0xbf, 0x3d, 0x7c, 0x34, 0xdd, 0x58, 0x6e, 0x80, 0xaf, 0x29, 0x0f, 0xad, 0x8f, 0x90, 0xbc, 0x3f,
	0x70, 0x61, 0x48, 0xfa, 0x9a, 0x30, 0x65, 0xaf, 0x87, 0xb3, 0x55, 0xce, 0x77, 0x78, 0x7b, 0xd0,
	0xf8, 0x1b, 0xec, 0xc4, 0x4b, 0x23, 0x6b, 0x3b, 0x9a, 0x8e, 0xf7, 0x6b, 0xfc, 0x66, 0x3f, 0x45,
	0x71, 0x95, 0xbf, 0xf4, 0xc7, 0x78, 0x3e, 0x6d, 0x0f, 0x62, 0xe7, 0x79, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x10, 0x65, 0x4a, 0x5d, 0x61, 0x05, 0x00, 0x00,
}

func (this *GrpcJsonTranscoder) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.DescriptorSet == nil {
		if this.DescriptorSet != nil {
			return false
		}
	} else if this.DescriptorSet == nil {
		return false
	} else if !this.DescriptorSet.Equal(that1.DescriptorSet) {
		return false
	}
	if len(this.Services) != len(that1.Services) {
		return false
	}
	for i := range this.Services {
		if this.Services[i] != that1.Services[i] {
			return false
		}
	}
	if !this.PrintOptions.Equal(that1.PrintOptions) {
		return false
	}
	if this.MatchIncomingRequestRoute != that1.MatchIncomingRequestRoute {
		return false
	}
	if len(this.IgnoredQueryParameters) != len(that1.IgnoredQueryParameters) {
		return false
	}
	for i := range this.IgnoredQueryParameters {
		if this.IgnoredQueryParameters[i] != that1.IgnoredQueryParameters[i] {
			return false
		}
	}
	if this.AutoMapping != that1.AutoMapping {
		return false
	}
	if this.IgnoreUnknownQueryParameters != that1.IgnoreUnknownQueryParameters {
		return false
	}
	if this.ConvertGrpcStatus != that1.ConvertGrpcStatus {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GrpcJsonTranscoder_ProtoDescriptor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder_ProtoDescriptor)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_ProtoDescriptor)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ProtoDescriptor != that1.ProtoDescriptor {
		return false
	}
	return true
}
func (this *GrpcJsonTranscoder_ProtoDescriptorBin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder_ProtoDescriptorBin)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_ProtoDescriptorBin)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ProtoDescriptorBin != that1.ProtoDescriptorBin {
		return false
	}
	return true
}
func (this *GrpcJsonTranscoder_PrintOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcJsonTranscoder_PrintOptions)
	if !ok {
		that2, ok := that.(GrpcJsonTranscoder_PrintOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AddWhitespace != that1.AddWhitespace {
		return false
	}
	if this.AlwaysPrintPrimitiveFields != that1.AlwaysPrintPrimitiveFields {
		return false
	}
	if this.AlwaysPrintEnumsAsInts != that1.AlwaysPrintEnumsAsInts {
		return false
	}
	if this.PreserveProtoFieldNames != that1.PreserveProtoFieldNames {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
