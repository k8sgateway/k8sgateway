// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tap/tap.proto

package tap

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Tap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaticConfig *StaticConfig `protobuf:"bytes,1,opt,name=static_config,json=staticConfig,proto3" json:"static_config,omitempty"`
}

func (x *Tap) Reset() {
	*x = Tap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tap) ProtoMessage() {}

func (x *Tap) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tap.ProtoReflect.Descriptor instead.
func (*Tap) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescGZIP(), []int{0}
}

func (x *Tap) GetStaticConfig() *StaticConfig {
	if x != nil {
		return x.StaticConfig
	}
	return nil
}

type StaticConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputConfig *OutputConfig `protobuf:"bytes,1,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
}

func (x *StaticConfig) Reset() {
	*x = StaticConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticConfig) ProtoMessage() {}

func (x *StaticConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticConfig.ProtoReflect.Descriptor instead.
func (*StaticConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescGZIP(), []int{1}
}

func (x *StaticConfig) GetOutputConfig() *OutputConfig {
	if x != nil {
		return x.OutputConfig
	}
	return nil
}

type OutputConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sinks to which tap data should be output. Currently, only a single sink
	// is supported. TODO is there a validate rule that we can use to enforce a
	// length of 1?
	Sinks []*Sink `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
}

func (x *OutputConfig) Reset() {
	*x = OutputConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputConfig) ProtoMessage() {}

func (x *OutputConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputConfig.ProtoReflect.Descriptor instead.
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescGZIP(), []int{2}
}

func (x *OutputConfig) GetSinks() []*Sink {
	if x != nil {
		return x.Sinks
	}
	return nil
}

type Sink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the output sink to which tap data should be written
	//
	// Types that are assignable to SinkType:
	//
	//	*Sink_GrpcService
	SinkType isSink_SinkType `protobuf_oneof:"SinkType"`
}

func (x *Sink) Reset() {
	*x = Sink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink.ProtoReflect.Descriptor instead.
func (*Sink) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescGZIP(), []int{3}
}

func (m *Sink) GetSinkType() isSink_SinkType {
	if m != nil {
		return m.SinkType
	}
	return nil
}

func (x *Sink) GetGrpcService() *GrpcService {
	if x, ok := x.GetSinkType().(*Sink_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

type isSink_SinkType interface {
	isSink_SinkType()
}

type Sink_GrpcService struct {
	// Write tap data out to a GRPC service
	// .solo.io.envoy.config.core.v3.GrpcService grpc_service = 1;
	GrpcService *GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*Sink_GrpcService) isSink_SinkType() {}

// A tap sink over a GRPC service
type GrpcService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TapServer *core.ResourceRef `protobuf:"bytes,1,opt,name=tap_server,json=tapServer,proto3" json:"tap_server,omitempty"`
}

func (x *GrpcService) Reset() {
	*x = GrpcService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcService) ProtoMessage() {}

func (x *GrpcService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcService.ProtoReflect.Descriptor instead.
func (*GrpcService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescGZIP(), []int{4}
}

func (x *GrpcService) GetTapServer() *core.ResourceRef {
	if x != nil {
		return x.TapServer
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x2f, 0x74, 0x61, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x03, 0x54, 0x61, 0x70, 0x12, 0x55, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x65, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x44, 0x0a, 0x0c, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x05, 0x73, 0x69,
	0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x6b, 0x73,
	0x22, 0x63, 0x0a, 0x04, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x4a, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x74, 0x61, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x0f, 0x0a, 0x08, 0x53, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x47, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x74, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x66, 0x52, 0x09, 0x74, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x61, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_goTypes = []interface{}{
	(*Tap)(nil),              // 0: tap.options.gloo.solo.io.Tap
	(*StaticConfig)(nil),     // 1: tap.options.gloo.solo.io.StaticConfig
	(*OutputConfig)(nil),     // 2: tap.options.gloo.solo.io.OutputConfig
	(*Sink)(nil),             // 3: tap.options.gloo.solo.io.Sink
	(*GrpcService)(nil),      // 4: tap.options.gloo.solo.io.GrpcService
	(*core.ResourceRef)(nil), // 5: core.solo.io.ResourceRef
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_depIdxs = []int32{
	1, // 0: tap.options.gloo.solo.io.Tap.static_config:type_name -> tap.options.gloo.solo.io.StaticConfig
	2, // 1: tap.options.gloo.solo.io.StaticConfig.output_config:type_name -> tap.options.gloo.solo.io.OutputConfig
	3, // 2: tap.options.gloo.solo.io.OutputConfig.sinks:type_name -> tap.options.gloo.solo.io.Sink
	4, // 3: tap.options.gloo.solo.io.Sink.grpc_service:type_name -> tap.options.gloo.solo.io.GrpcService
	5, // 4: tap.options.gloo.solo.io.GrpcService.tap_server:type_name -> core.solo.io.ResourceRef
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tap); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sink); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcService); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Sink_GrpcService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_tap_tap_proto_depIdxs = nil
}
