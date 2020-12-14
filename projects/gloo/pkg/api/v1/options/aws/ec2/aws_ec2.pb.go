// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/ec2/aws_ec2.proto

package ec2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The AWS Region where the desired EC2 instances exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Optional, if not set, Gloo will try to use the default AWS secret specified by environment variables.
	// If a secret is not provided, the environment must specify both the AWS access key and secret.
	// The environment variables used to indicate the AWS account can be:
	// - for the access key: "AWS_ACCESS_KEY_ID" or "AWS_ACCESS_KEY"
	// - for the secret: "AWS_SECRET_ACCESS_KEY" or "AWS_SECRET_KEY"
	// If set, a [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	// Gloo will create an EC2 API client with this credential. You may choose to use a credential with limited access
	// in conjunction with a list of Roles, specified by their Amazon Resource Number (ARN).
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// Optional, Amazon Resource Number (ARN) referring to IAM Role that should be assumed when the Upstream
	// queries for eligible EC2 instances.
	// If provided, Gloo will create an EC2 API client with the provided role. If not provided, Gloo will not assume
	// a role.
	RoleArn string `protobuf:"bytes,7,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
	// List of tag filters for selecting instances
	// An instance must match all the filters in order to be selected
	// Filter keys are not case-sensitive
	Filters []*TagFilter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// If set, will use the EC2 public IP address. Defaults to the private IP address.
	PublicIp bool `protobuf:"varint,4,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	// If set, will use this port on EC2 instances. Defaults to port 80.
	Port uint32 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[0]
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
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if x != nil {
		return x.SecretRef
	}
	return nil
}

func (x *UpstreamSpec) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

func (x *UpstreamSpec) GetFilters() []*TagFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *UpstreamSpec) GetPublicIp() bool {
	if x != nil {
		return x.PublicIp
	}
	return false
}

func (x *UpstreamSpec) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type TagFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Spec:
	//	*TagFilter_Key
	//	*TagFilter_KvPair_
	Spec isTagFilter_Spec `protobuf_oneof:"spec"`
}

func (x *TagFilter) Reset() {
	*x = TagFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagFilter) ProtoMessage() {}

func (x *TagFilter) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagFilter.ProtoReflect.Descriptor instead.
func (*TagFilter) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescGZIP(), []int{1}
}

func (m *TagFilter) GetSpec() isTagFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (x *TagFilter) GetKey() string {
	if x, ok := x.GetSpec().(*TagFilter_Key); ok {
		return x.Key
	}
	return ""
}

func (x *TagFilter) GetKvPair() *TagFilter_KvPair {
	if x, ok := x.GetSpec().(*TagFilter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

type isTagFilter_Spec interface {
	isTagFilter_Spec()
}

type TagFilter_Key struct {
	// if set, only instances that have a tag with this key will be matched
	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

type TagFilter_KvPair_ struct {
	// if set, only instances that have a tag with this key and value
	KvPair *TagFilter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof"`
}

func (*TagFilter_Key) isTagFilter_Spec() {}

func (*TagFilter_KvPair_) isTagFilter_Spec() {}

type TagFilter_KvPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values are case-sensitive
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TagFilter_KvPair) Reset() {
	*x = TagFilter_KvPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagFilter_KvPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagFilter_KvPair) ProtoMessage() {}

func (x *TagFilter_KvPair) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagFilter_KvPair.ProtoReflect.Descriptor instead.
func (*TagFilter_KvPair) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TagFilter_KvPair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TagFilter_KvPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDesc = []byte{
	0x0a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x63, 0x32, 0x2f, 0x61,
	0x77, 0x73, 0x5f, 0x65, 0x63, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x77,
	0x73, 0x5f, 0x65, 0x63, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a,
	0x0c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x41, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x77,
	0x73, 0x5f, 0x65, 0x63, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xa4,
	0x01, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x49, 0x0a, 0x07, 0x6b, 0x76, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x63, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x54, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x76, 0x50, 0x61, 0x69,
	0x72, 0x48, 0x00, 0x52, 0x06, 0x6b, 0x76, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x30, 0x0a, 0x06, 0x4b,
	0x76, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x42, 0x4a, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x65, 0x63, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_goTypes = []interface{}{
	(*UpstreamSpec)(nil),     // 0: aws_ec2.options.gloo.solo.io.UpstreamSpec
	(*TagFilter)(nil),        // 1: aws_ec2.options.gloo.solo.io.TagFilter
	(*TagFilter_KvPair)(nil), // 2: aws_ec2.options.gloo.solo.io.TagFilter.KvPair
	(*core.ResourceRef)(nil), // 3: core.solo.io.ResourceRef
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_depIdxs = []int32{
	3, // 0: aws_ec2.options.gloo.solo.io.UpstreamSpec.secret_ref:type_name -> core.solo.io.ResourceRef
	1, // 1: aws_ec2.options.gloo.solo.io.UpstreamSpec.filters:type_name -> aws_ec2.options.gloo.solo.io.TagFilter
	2, // 2: aws_ec2.options.gloo.solo.io.TagFilter.kv_pair:type_name -> aws_ec2.options.gloo.solo.io.TagFilter.KvPair
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagFilter); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagFilter_KvPair); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TagFilter_Key)(nil),
		(*TagFilter_KvPair_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_ec2_aws_ec2_proto_depIdxs = nil
}
