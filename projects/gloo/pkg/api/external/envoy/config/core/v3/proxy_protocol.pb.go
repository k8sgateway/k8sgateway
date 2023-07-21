// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/proxy_protocol.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ProxyProtocolPassThroughTLVs_PassTLVsMatchType int32

const (
	// Pass all TLVs.
	ProxyProtocolPassThroughTLVs_INCLUDE_ALL ProxyProtocolPassThroughTLVs_PassTLVsMatchType = 0
	// Pass specific TLVs defined in tlv_type.
	ProxyProtocolPassThroughTLVs_INCLUDE ProxyProtocolPassThroughTLVs_PassTLVsMatchType = 1
)

// Enum value maps for ProxyProtocolPassThroughTLVs_PassTLVsMatchType.
var (
	ProxyProtocolPassThroughTLVs_PassTLVsMatchType_name = map[int32]string{
		0: "INCLUDE_ALL",
		1: "INCLUDE",
	}
	ProxyProtocolPassThroughTLVs_PassTLVsMatchType_value = map[string]int32{
		"INCLUDE_ALL": 0,
		"INCLUDE":     1,
	}
)

func (x ProxyProtocolPassThroughTLVs_PassTLVsMatchType) Enum() *ProxyProtocolPassThroughTLVs_PassTLVsMatchType {
	p := new(ProxyProtocolPassThroughTLVs_PassTLVsMatchType)
	*p = x
	return p
}

func (x ProxyProtocolPassThroughTLVs_PassTLVsMatchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyProtocolPassThroughTLVs_PassTLVsMatchType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes[0].Descriptor()
}

func (ProxyProtocolPassThroughTLVs_PassTLVsMatchType) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes[0]
}

func (x ProxyProtocolPassThroughTLVs_PassTLVsMatchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyProtocolPassThroughTLVs_PassTLVsMatchType.Descriptor instead.
func (ProxyProtocolPassThroughTLVs_PassTLVsMatchType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP(), []int{0, 0}
}

type ProxyProtocolConfig_Version int32

const (
	// PROXY protocol version 1. Human readable format.
	ProxyProtocolConfig_V1 ProxyProtocolConfig_Version = 0
	// PROXY protocol version 2. Binary format.
	ProxyProtocolConfig_V2 ProxyProtocolConfig_Version = 1
)

// Enum value maps for ProxyProtocolConfig_Version.
var (
	ProxyProtocolConfig_Version_name = map[int32]string{
		0: "V1",
		1: "V2",
	}
	ProxyProtocolConfig_Version_value = map[string]int32{
		"V1": 0,
		"V2": 1,
	}
)

func (x ProxyProtocolConfig_Version) Enum() *ProxyProtocolConfig_Version {
	p := new(ProxyProtocolConfig_Version)
	*p = x
	return p
}

func (x ProxyProtocolConfig_Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyProtocolConfig_Version) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes[1].Descriptor()
}

func (ProxyProtocolConfig_Version) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes[1]
}

func (x ProxyProtocolConfig_Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyProtocolConfig_Version.Descriptor instead.
func (ProxyProtocolConfig_Version) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP(), []int{1, 0}
}

type ProxyProtocolPassThroughTLVs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The strategy to pass through TLVs. Default is INCLUDE_ALL.
	// If INCLUDE_ALL is set, all TLVs will be passed through no matter the tlv_type field.
	MatchType ProxyProtocolPassThroughTLVs_PassTLVsMatchType `protobuf:"varint,1,opt,name=match_type,json=matchType,proto3,enum=solo.io.envoy.config.core.v3.ProxyProtocolPassThroughTLVs_PassTLVsMatchType" json:"match_type,omitempty"`
	// The TLV types that are applied based on match_type.
	// TLV type is defined as uint8_t in proxy protocol. See `the spec
	// <https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt>`_ for details.
	TlvType []uint32 `protobuf:"varint,2,rep,packed,name=tlv_type,json=tlvType,proto3" json:"tlv_type,omitempty"`
}

func (x *ProxyProtocolPassThroughTLVs) Reset() {
	*x = ProxyProtocolPassThroughTLVs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolPassThroughTLVs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolPassThroughTLVs) ProtoMessage() {}

func (x *ProxyProtocolPassThroughTLVs) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolPassThroughTLVs.ProtoReflect.Descriptor instead.
func (*ProxyProtocolPassThroughTLVs) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyProtocolPassThroughTLVs) GetMatchType() ProxyProtocolPassThroughTLVs_PassTLVsMatchType {
	if x != nil {
		return x.MatchType
	}
	return ProxyProtocolPassThroughTLVs_INCLUDE_ALL
}

func (x *ProxyProtocolPassThroughTLVs) GetTlvType() []uint32 {
	if x != nil {
		return x.TlvType
	}
	return nil
}

type ProxyProtocolConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The PROXY protocol version to use. See https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt for details
	Version ProxyProtocolConfig_Version `protobuf:"varint,1,opt,name=version,proto3,enum=solo.io.envoy.config.core.v3.ProxyProtocolConfig_Version" json:"version,omitempty"`
	// This config controls which TLVs can be passed to filter state if it is Proxy Protocol
	// V2 header. If there is no setting for this field, no TLVs will be passed through.
	PassThroughTlvs *ProxyProtocolPassThroughTLVs `protobuf:"bytes,2,opt,name=pass_through_tlvs,json=passThroughTlvs,proto3" json:"pass_through_tlvs,omitempty"`
}

func (x *ProxyProtocolConfig) Reset() {
	*x = ProxyProtocolConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolConfig) ProtoMessage() {}

func (x *ProxyProtocolConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyProtocolConfig.ProtoReflect.Descriptor instead.
func (*ProxyProtocolConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyProtocolConfig) GetVersion() ProxyProtocolConfig_Version {
	if x != nil {
		return x.Version
	}
	return ProxyProtocolConfig_V1
}

func (x *ProxyProtocolConfig) GetPassThroughTlvs() *ProxyProtocolPassThroughTLVs {
	if x != nil {
		return x.PassThroughTlvs
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x6f, 0x75, 0x67, 0x68, 0x54, 0x4c, 0x56, 0x73, 0x12, 0x6b, 0x0a, 0x0a, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x6f, 0x75, 0x67, 0x68, 0x54, 0x4c, 0x56, 0x73, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x54, 0x4c,
	0x56, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x74, 0x6c, 0x76, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x92, 0x01, 0x07,
	0x22, 0x05, 0x2a, 0x03, 0x10, 0x80, 0x02, 0x52, 0x07, 0x74, 0x6c, 0x76, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x31, 0x0a, 0x11, 0x50, 0x61, 0x73, 0x73, 0x54, 0x4c, 0x56, 0x73, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45,
	0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44,
	0x45, 0x10, 0x01, 0x22, 0xed, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x53, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x66, 0x0a, 0x11, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68,
	0x5f, 0x74, 0x6c, 0x76, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x54, 0x4c, 0x56, 0x73, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72,
	0x6f, 0x75, 0x67, 0x68, 0x54, 0x6c, 0x76, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56,
	0x32, 0x10, 0x01, 0x42, 0xa3, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5,
	0x04, 0x01, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a, 0x2a, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_goTypes = []interface{}{
	(ProxyProtocolPassThroughTLVs_PassTLVsMatchType)(0), // 0: solo.io.envoy.config.core.v3.ProxyProtocolPassThroughTLVs.PassTLVsMatchType
	(ProxyProtocolConfig_Version)(0),                    // 1: solo.io.envoy.config.core.v3.ProxyProtocolConfig.Version
	(*ProxyProtocolPassThroughTLVs)(nil),                // 2: solo.io.envoy.config.core.v3.ProxyProtocolPassThroughTLVs
	(*ProxyProtocolConfig)(nil),                         // 3: solo.io.envoy.config.core.v3.ProxyProtocolConfig
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_depIdxs = []int32{
	0, // 0: solo.io.envoy.config.core.v3.ProxyProtocolPassThroughTLVs.match_type:type_name -> solo.io.envoy.config.core.v3.ProxyProtocolPassThroughTLVs.PassTLVsMatchType
	1, // 1: solo.io.envoy.config.core.v3.ProxyProtocolConfig.version:type_name -> solo.io.envoy.config.core.v3.ProxyProtocolConfig.Version
	2, // 2: solo.io.envoy.config.core.v3.ProxyProtocolConfig.pass_through_tlvs:type_name -> solo.io.envoy.config.core.v3.ProxyProtocolPassThroughTLVs
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolPassThroughTLVs); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyProtocolConfig); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_depIdxs = nil
}
