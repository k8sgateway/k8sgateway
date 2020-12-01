// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/proxy_protocol.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes[0].Descriptor()
}

func (ProxyProtocolConfig_Version) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes[0]
}

func (x ProxyProtocolConfig_Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyProtocolConfig_Version.Descriptor instead.
func (ProxyProtocolConfig_Version) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP(), []int{0, 0}
}

type ProxyProtocolConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The PROXY protocol version to use. See https://www.haproxy.org/download/2.1/doc/proxy-protocol.txt for details
	Version ProxyProtocolConfig_Version `protobuf:"varint,1,opt,name=version,proto3,enum=solo.io.envoy.config.core.v3.ProxyProtocolConfig_Version" json:"version,omitempty"`
}

func (x *ProxyProtocolConfig) Reset() {
	*x = ProxyProtocolConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyProtocolConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyProtocolConfig) ProtoMessage() {}

func (x *ProxyProtocolConfig) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProxyProtocolConfig.ProtoReflect.Descriptor instead.
func (*ProxyProtocolConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyProtocolConfig) GetVersion() ProxyProtocolConfig_Version {
	if x != nil {
		return x.Version
	}
	return ProxyProtocolConfig_V1
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
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x53, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x39, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x06, 0x0a, 0x02, 0x56, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x01,
	0x42, 0x9f, 0x01, 0x0a, 0x2a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42,
	0x12, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_goTypes = []interface{}{
	(ProxyProtocolConfig_Version)(0), // 0: solo.io.envoy.config.core.v3.ProxyProtocolConfig.Version
	(*ProxyProtocolConfig)(nil),      // 1: solo.io.envoy.config.core.v3.ProxyProtocolConfig
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_proxy_protocol_proto_depIdxs = []int32{
	0, // 0: solo.io.envoy.config.core.v3.ProxyProtocolConfig.version:type_name -> solo.io.envoy.config.core.v3.ProxyProtocolConfig.Version
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			NumEnums:      1,
			NumMessages:   1,
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
