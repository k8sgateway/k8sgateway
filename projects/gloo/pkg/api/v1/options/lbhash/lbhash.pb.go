// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/lbhash/lbhash.proto

package lbhash

import (
	reflect "reflect"
	sync "sync"

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

// Specifies the route’s hashing policy if the upstream cluster uses a hashing load balancer.
// https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route.proto#envoy-api-msg-route-routeaction-hashpolicy
type RouteActionHashConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of policies Envoy will use when generating a hash key for a hashing load balancer
	HashPolicies []*HashPolicy `protobuf:"bytes,1,rep,name=hash_policies,json=hashPolicies,proto3" json:"hash_policies,omitempty"`
}

func (x *RouteActionHashConfig) Reset() {
	*x = RouteActionHashConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteActionHashConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteActionHashConfig) ProtoMessage() {}

func (x *RouteActionHashConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteActionHashConfig.ProtoReflect.Descriptor instead.
func (*RouteActionHashConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescGZIP(), []int{0}
}

func (x *RouteActionHashConfig) GetHashPolicies() []*HashPolicy {
	if x != nil {
		return x.HashPolicies
	}
	return nil
}

// Envoy supports two types of cookie affinity:
// - Passive: Envoy reads the cookie from the headers
// - Generated: Envoy uses the cookie spec to generate a cookie
// In either case, the cookie is incorporated in the hash key.
// additional notes https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route.proto#envoy-api-msg-route-routeaction-hashpolicy-cookie
type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required, the name of the cookie to be used to obtain the hash key
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie.
	Ttl *duration.Duration `protobuf:"bytes,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// The name of the path for the cookie. If no path is specified here, no path will be set for the cookie.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescGZIP(), []int{1}
}

func (x *Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookie) GetTtl() *duration.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Specifies an element of Envoy's hashing policy for hashing load balancers
type HashPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to KeyType:
	//
	//	*HashPolicy_Header
	//	*HashPolicy_Cookie
	//	*HashPolicy_SourceIp
	KeyType isHashPolicy_KeyType `protobuf_oneof:"KeyType"`
	// If set, and a hash key is available after evaluating this policy, Envoy will skip the subsequent policies and
	// use the key as it is.
	// This is useful for defining "fallback" policies and limiting the time Envoy spends generating hash keys.
	Terminal bool `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
}

func (x *HashPolicy) Reset() {
	*x = HashPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashPolicy) ProtoMessage() {}

func (x *HashPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashPolicy.ProtoReflect.Descriptor instead.
func (*HashPolicy) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescGZIP(), []int{2}
}

func (m *HashPolicy) GetKeyType() isHashPolicy_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (x *HashPolicy) GetHeader() string {
	if x, ok := x.GetKeyType().(*HashPolicy_Header); ok {
		return x.Header
	}
	return ""
}

func (x *HashPolicy) GetCookie() *Cookie {
	if x, ok := x.GetKeyType().(*HashPolicy_Cookie); ok {
		return x.Cookie
	}
	return nil
}

func (x *HashPolicy) GetSourceIp() bool {
	if x, ok := x.GetKeyType().(*HashPolicy_SourceIp); ok {
		return x.SourceIp
	}
	return false
}

func (x *HashPolicy) GetTerminal() bool {
	if x != nil {
		return x.Terminal
	}
	return false
}

type isHashPolicy_KeyType interface {
	isHashPolicy_KeyType()
}

type HashPolicy_Header struct {
	// Use a given header's value as a component of the hashing load balancer's hash key
	Header string `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type HashPolicy_Cookie struct {
	// Use a given cookie as a component of the hashing load balancer's hash key
	Cookie *Cookie `protobuf:"bytes,2,opt,name=cookie,proto3,oneof"`
}

type HashPolicy_SourceIp struct {
	// Use the request's source IP address as a component of the hashing load balancer's hash key
	SourceIp bool `protobuf:"varint,3,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

func (*HashPolicy_Header) isHashPolicy_KeyType() {}

func (*HashPolicy_Cookie) isHashPolicy_KeyType() {}

func (*HashPolicy_SourceIp) isHashPolicy_KeyType() {}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x62, 0x68, 0x61, 0x73, 0x68, 0x2f, 0x6c, 0x62,
	0x68, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6c, 0x62, 0x68, 0x61,
	0x73, 0x68, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x15, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x62,
	0x68, 0x61, 0x73, 0x68, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x22, 0x5d, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0xab, 0x01, 0x0a, 0x0a, 0x48, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x18, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x06, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x62, 0x68,
	0x61, 0x73, 0x68, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x4d, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6c, 0x62, 0x68, 0x61,
	0x73, 0x68, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_goTypes = []interface{}{
	(*RouteActionHashConfig)(nil), // 0: lbhash.options.gloo.solo.io.RouteActionHashConfig
	(*Cookie)(nil),                // 1: lbhash.options.gloo.solo.io.Cookie
	(*HashPolicy)(nil),            // 2: lbhash.options.gloo.solo.io.HashPolicy
	(*duration.Duration)(nil),     // 3: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_depIdxs = []int32{
	2, // 0: lbhash.options.gloo.solo.io.RouteActionHashConfig.hash_policies:type_name -> lbhash.options.gloo.solo.io.HashPolicy
	3, // 1: lbhash.options.gloo.solo.io.Cookie.ttl:type_name -> google.protobuf.Duration
	1, // 2: lbhash.options.gloo.solo.io.HashPolicy.cookie:type_name -> lbhash.options.gloo.solo.io.Cookie
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteActionHashConfig); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashPolicy); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*HashPolicy_Header)(nil),
		(*HashPolicy_Cookie)(nil),
		(*HashPolicy_SourceIp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_lbhash_lbhash_proto_depIdxs = nil
}
