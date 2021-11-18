// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxyprotocol/proxyprotocol.proto

package proxy_protocol

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CustomProxyProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rules to apply to requests.
	Rules []*CustomProxyProtocol_Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// Allow requests through that don't use proxy protocol. Defaults to false.
	//
	// .. attention::
	//
	//   This breaks conformance with the specification.
	//   Only enable if ALL traffic to the listener comes from a trusted source.
	//   For more information on the security implications of this feature, see
	//   https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt
	//
	AllowRequestsWithoutProxyProtocol bool `protobuf:"varint,2,opt,name=allow_requests_without_proxy_protocol,json=allowRequestsWithoutProxyProtocol,proto3" json:"allow_requests_without_proxy_protocol,omitempty"`
}

func (x *CustomProxyProtocol) Reset() {
	*x = CustomProxyProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomProxyProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomProxyProtocol) ProtoMessage() {}

func (x *CustomProxyProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomProxyProtocol.ProtoReflect.Descriptor instead.
func (*CustomProxyProtocol) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescGZIP(), []int{0}
}

func (x *CustomProxyProtocol) GetRules() []*CustomProxyProtocol_Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *CustomProxyProtocol) GetAllowRequestsWithoutProxyProtocol() bool {
	if x != nil {
		return x.AllowRequestsWithoutProxyProtocol
	}
	return false
}

type CustomProxyProtocol_KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CustomProxyProtocol_KeyValuePair) Reset() {
	*x = CustomProxyProtocol_KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomProxyProtocol_KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomProxyProtocol_KeyValuePair) ProtoMessage() {}

func (x *CustomProxyProtocol_KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomProxyProtocol_KeyValuePair.ProtoReflect.Descriptor instead.
func (*CustomProxyProtocol_KeyValuePair) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CustomProxyProtocol_KeyValuePair) GetMetadataNamespace() string {
	if x != nil {
		return x.MetadataNamespace
	}
	return ""
}

func (x *CustomProxyProtocol_KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// A Rule defines what metadata to apply when a header is present or missing.
type CustomProxyProtocol_Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type that triggers the rule - required
	// TLV type is defined as uint8_t in proxy protocol. See `the spec
	// <https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt>`_ for details.
	TlvType uint32 `protobuf:"varint,1,opt,name=tlv_type,json=tlvType,proto3" json:"tlv_type,omitempty"`
	// If the TLV type is present, apply this metadata KeyValuePair.
	OnTlvPresent *CustomProxyProtocol_KeyValuePair `protobuf:"bytes,2,opt,name=on_tlv_present,json=onTlvPresent,proto3" json:"on_tlv_present,omitempty"`
}

func (x *CustomProxyProtocol_Rule) Reset() {
	*x = CustomProxyProtocol_Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomProxyProtocol_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomProxyProtocol_Rule) ProtoMessage() {}

func (x *CustomProxyProtocol_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomProxyProtocol_Rule.ProtoReflect.Descriptor instead.
func (*CustomProxyProtocol_Rule) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CustomProxyProtocol_Rule) GetTlvType() uint32 {
	if x != nil {
		return x.TlvType
	}
	return 0
}

func (x *CustomProxyProtocol_Rule) GetOnTlvPresent() *CustomProxyProtocol_KeyValuePair {
	if x != nil {
		return x.OnTlvPresent
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDesc = []byte{
	0x0a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x03, 0x0a, 0x13, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x57, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x25, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x6f, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x21, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x58, 0x0a, 0x0c, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0x9c, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x23,
	0x0a, 0x08, 0x74, 0x6c, 0x76, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x2a, 0x03, 0x10, 0x80, 0x02, 0x52, 0x07, 0x74, 0x6c, 0x76, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x6f, 0x6e, 0x5f, 0x74, 0x6c, 0x76, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0c, 0x6f, 0x6e, 0x54, 0x6c, 0x76, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x42, 0xb9, 0x01, 0x0a, 0x41, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_goTypes = []interface{}{
	(*CustomProxyProtocol)(nil),              // 0: envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol
	(*CustomProxyProtocol_KeyValuePair)(nil), // 1: envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol.KeyValuePair
	(*CustomProxyProtocol_Rule)(nil),         // 2: envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol.Rule
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_depIdxs = []int32{
	2, // 0: envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol.rules:type_name -> envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol.Rule
	1, // 1: envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol.Rule.on_tlv_present:type_name -> envoy.config.listener.proxy_protocol.v3.CustomProxyProtocol.KeyValuePair
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomProxyProtocol); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomProxyProtocol_KeyValuePair); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomProxyProtocol_Rule); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxyprotocol_proxyprotocol_proto_depIdxs = nil
}
