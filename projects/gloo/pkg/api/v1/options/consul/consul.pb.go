// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/consul/consul.proto

package consul

import (
	reflect "reflect"
	sync "sync"

	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
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

// Upstream Spec for Consul Upstreams
// consul Upstreams represent a set of one or more addressable pods for a consul Service
// the Gloo consul Upstream maps to a single service port. Because consul Services support multiple ports,
// Gloo requires that a different upstream be created for each port
// consul Upstreams are typically generated automatically by Gloo from the consul API
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Consul Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Deprecated: This field was renamed to subset_tags. If subset_tags is used, this field is
	// ignored. Otherwise, the behavior is the same as subset_tags field below.
	ServiceTags []string `protobuf:"bytes,2,rep,name=service_tags,json=serviceTags,proto3" json:"service_tags,omitempty"`
	// Gloo will segment instances based off of these tags. This allows you to set routes that route
	// to a subset of the instances of the service
	SubsetTags []string `protobuf:"bytes,6,rep,name=subset_tags,json=subsetTags,proto3" json:"subset_tags,omitempty"`
	// The list of service tags Gloo should search for on a service instance
	// before deciding whether or not to include the instance as part of this
	// upstream. Empty list means that all service instances with the same service name will be
	// included. When not empty, only service instances that match all of the tags (subset match) will be selected
	// for this upstream.
	InstanceTags []string `protobuf:"bytes,7,rep,name=instance_tags,json=instanceTags,proto3" json:"instance_tags,omitempty"`
	// The opposite of instanceTags, this is a list of service tags that gloo should ensure are not
	// in a service instance before including it in an upstream.
	InstanceBlacklistTags []string `protobuf:"bytes,8,rep,name=instance_blacklist_tags,json=instanceBlacklistTags,proto3" json:"instance_blacklist_tags,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *options.ServiceSpec `protobuf:"bytes,3,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// Sets the consistency mode. The default is DefaultMode.
	//
	// Note: Gloo handles staleness well (as it runs update loops ~ once/second) but makes many requests
	// to get consul endpoints so users may want to opt into stale reads once the implications are understood.
	ConsistencyMode ConsulConsistencyModes `protobuf:"varint,9,opt,name=consistencyMode,proto3,enum=consul.options.gloo.solo.io.ConsulConsistencyModes" json:"consistencyMode,omitempty"`
	// QueryOptions are the query options to use for all Consul queries.
	QueryOptions *QueryOptions `protobuf:"bytes,10,opt,name=query_options,json=queryOptions,proto3" json:"query_options,omitempty"`
	// Is this consul service connect enabled.
	ConnectEnabled bool `protobuf:"varint,4,opt,name=connect_enabled,json=connectEnabled,proto3" json:"connect_enabled,omitempty"`
	// The data centers in which the service instance represented by this upstream is registered.
	DataCenters []string `protobuf:"bytes,5,rep,name=data_centers,json=dataCenters,proto3" json:"data_centers,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *UpstreamSpec) GetServiceTags() []string {
	if x != nil {
		return x.ServiceTags
	}
	return nil
}

func (x *UpstreamSpec) GetSubsetTags() []string {
	if x != nil {
		return x.SubsetTags
	}
	return nil
}

func (x *UpstreamSpec) GetInstanceTags() []string {
	if x != nil {
		return x.InstanceTags
	}
	return nil
}

func (x *UpstreamSpec) GetInstanceBlacklistTags() []string {
	if x != nil {
		return x.InstanceBlacklistTags
	}
	return nil
}

func (x *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if x != nil {
		return x.ServiceSpec
	}
	return nil
}

func (x *UpstreamSpec) GetConsistencyMode() ConsulConsistencyModes {
	if x != nil {
		return x.ConsistencyMode
	}
	return ConsulConsistencyModes_DefaultMode
}

func (x *UpstreamSpec) GetQueryOptions() *QueryOptions {
	if x != nil {
		return x.QueryOptions
	}
	return nil
}

func (x *UpstreamSpec) GetConnectEnabled() bool {
	if x != nil {
		return x.ConnectEnabled
	}
	return false
}

func (x *UpstreamSpec) GetDataCenters() []string {
	if x != nil {
		return x.DataCenters
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x04, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61,
	0x67, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x62,
	0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x15, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x5d, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d,
	0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x43, 0x6f,
	0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x4e, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x42, 0x4d, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_goTypes = []interface{}{
	(*UpstreamSpec)(nil),        // 0: consul.options.gloo.solo.io.UpstreamSpec
	(*options.ServiceSpec)(nil), // 1: options.gloo.solo.io.ServiceSpec
	(ConsulConsistencyModes)(0), // 2: consul.options.gloo.solo.io.ConsulConsistencyModes
	(*QueryOptions)(nil),        // 3: consul.options.gloo.solo.io.QueryOptions
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_depIdxs = []int32{
	1, // 0: consul.options.gloo.solo.io.UpstreamSpec.service_spec:type_name -> options.gloo.solo.io.ServiceSpec
	2, // 1: consul.options.gloo.solo.io.UpstreamSpec.consistencyMode:type_name -> consul.options.gloo.solo.io.ConsulConsistencyModes
	3, // 2: consul.options.gloo.solo.io.UpstreamSpec.query_options:type_name -> consul.options.gloo.solo.io.QueryOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_query_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSpec); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_consul_consul_proto_depIdxs = nil
}
