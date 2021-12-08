// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/faultinjection/fault.proto

package faultinjection

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RouteAbort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percentage of requests that should be aborted, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// This should be a standard HTTP status, i.e. 503. Defaults to 0.
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
}

func (x *RouteAbort) Reset() {
	*x = RouteAbort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAbort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAbort) ProtoMessage() {}

func (x *RouteAbort) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAbort.ProtoReflect.Descriptor instead.
func (*RouteAbort) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescGZIP(), []int{0}
}

func (x *RouteAbort) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *RouteAbort) GetHttpStatus() uint32 {
	if x != nil {
		return x.HttpStatus
	}
	return 0
}

type RouteDelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percentage of requests that should be delayed, defaulting to 0.
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage float32 `protobuf:"fixed32,1,opt,name=percentage,proto3" json:"percentage,omitempty"`
	// Fixed delay, defaulting to 0.
	FixedDelay *duration.Duration `protobuf:"bytes,2,opt,name=fixed_delay,json=fixedDelay,proto3" json:"fixed_delay,omitempty"`
}

func (x *RouteDelay) Reset() {
	*x = RouteDelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteDelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteDelay) ProtoMessage() {}

func (x *RouteDelay) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteDelay.ProtoReflect.Descriptor instead.
func (*RouteDelay) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescGZIP(), []int{1}
}

func (x *RouteDelay) GetPercentage() float32 {
	if x != nil {
		return x.Percentage
	}
	return 0
}

func (x *RouteDelay) GetFixedDelay() *duration.Duration {
	if x != nil {
		return x.FixedDelay
	}
	return nil
}

type RouteFaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Abort *RouteAbort `protobuf:"bytes,1,opt,name=abort,proto3" json:"abort,omitempty"`
	Delay *RouteDelay `protobuf:"bytes,2,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *RouteFaults) Reset() {
	*x = RouteFaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteFaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteFaults) ProtoMessage() {}

func (x *RouteFaults) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteFaults.ProtoReflect.Descriptor instead.
func (*RouteFaults) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescGZIP(), []int{2}
}

func (x *RouteFaults) GetAbort() *RouteAbort {
	if x != nil {
		return x.Abort
	}
	return nil
}

func (x *RouteFaults) GetDelay() *RouteDelay {
	if x != nil {
		return x.Delay
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0a, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68,
	0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x72, 0x0a, 0x0a, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a,
	0x00, 0x52, 0x0a, 0x66, 0x69, 0x78, 0x65, 0x64, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x89, 0x01,
	0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x3c, 0x0a,
	0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x55, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x69, 0x6e, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_goTypes = []interface{}{
	(*RouteAbort)(nil),        // 0: fault.options.gloo.solo.io.RouteAbort
	(*RouteDelay)(nil),        // 1: fault.options.gloo.solo.io.RouteDelay
	(*RouteFaults)(nil),       // 2: fault.options.gloo.solo.io.RouteFaults
	(*duration.Duration)(nil), // 3: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_depIdxs = []int32{
	3, // 0: fault.options.gloo.solo.io.RouteDelay.fixed_delay:type_name -> google.protobuf.Duration
	0, // 1: fault.options.gloo.solo.io.RouteFaults.abort:type_name -> fault.options.gloo.solo.io.RouteAbort
	1, // 2: fault.options.gloo.solo.io.RouteFaults.delay:type_name -> fault.options.gloo.solo.io.RouteDelay
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAbort); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteDelay); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteFaults); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_faultinjection_fault_proto_depIdxs = nil
}
