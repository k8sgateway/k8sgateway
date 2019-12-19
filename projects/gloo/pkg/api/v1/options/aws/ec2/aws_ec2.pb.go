// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/ec2/aws_ec2.proto

package ec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
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
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14583d3ecc23381, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *UpstreamSpec) GetRoleArn() string {
	if m != nil {
		return m.RoleArn
	}
	return ""
}

func (m *UpstreamSpec) GetFilters() []*TagFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *UpstreamSpec) GetPublicIp() bool {
	if m != nil {
		return m.PublicIp
	}
	return false
}

func (m *UpstreamSpec) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type TagFilter struct {
	// Types that are valid to be assigned to Spec:
	//	*TagFilter_Key
	//	*TagFilter_KvPair_
	Spec                 isTagFilter_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TagFilter) Reset()         { *m = TagFilter{} }
func (m *TagFilter) String() string { return proto.CompactTextString(m) }
func (*TagFilter) ProtoMessage()    {}
func (*TagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14583d3ecc23381, []int{1}
}
func (m *TagFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter.Unmarshal(m, b)
}
func (m *TagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter.Marshal(b, m, deterministic)
}
func (m *TagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter.Merge(m, src)
}
func (m *TagFilter) XXX_Size() int {
	return xxx_messageInfo_TagFilter.Size(m)
}
func (m *TagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter proto.InternalMessageInfo

type isTagFilter_Spec interface {
	isTagFilter_Spec()
	Equal(interface{}) bool
}

type TagFilter_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty"`
}
type TagFilter_KvPair_ struct {
	KvPair *TagFilter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof" json:"kv_pair,omitempty"`
}

func (*TagFilter_Key) isTagFilter_Spec()     {}
func (*TagFilter_KvPair_) isTagFilter_Spec() {}

func (m *TagFilter) GetSpec() isTagFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TagFilter) GetKey() string {
	if x, ok := m.GetSpec().(*TagFilter_Key); ok {
		return x.Key
	}
	return ""
}

func (m *TagFilter) GetKvPair() *TagFilter_KvPair {
	if x, ok := m.GetSpec().(*TagFilter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagFilter_Key)(nil),
		(*TagFilter_KvPair_)(nil),
	}
}

type TagFilter_KvPair struct {
	// keys are not case-sensitive, as with AWS Condition Keys
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// values are case-sensitive
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFilter_KvPair) Reset()         { *m = TagFilter_KvPair{} }
func (m *TagFilter_KvPair) String() string { return proto.CompactTextString(m) }
func (*TagFilter_KvPair) ProtoMessage()    {}
func (*TagFilter_KvPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14583d3ecc23381, []int{1, 0}
}
func (m *TagFilter_KvPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFilter_KvPair.Unmarshal(m, b)
}
func (m *TagFilter_KvPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFilter_KvPair.Marshal(b, m, deterministic)
}
func (m *TagFilter_KvPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter_KvPair.Merge(m, src)
}
func (m *TagFilter_KvPair) XXX_Size() int {
	return xxx_messageInfo_TagFilter_KvPair.Size(m)
}
func (m *TagFilter_KvPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter_KvPair.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter_KvPair proto.InternalMessageInfo

func (m *TagFilter_KvPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TagFilter_KvPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws_ec2.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*TagFilter)(nil), "aws_ec2.options.gloo.solo.io.TagFilter")
	proto.RegisterType((*TagFilter_KvPair)(nil), "aws_ec2.options.gloo.solo.io.TagFilter.KvPair")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/ec2/aws_ec2.proto", fileDescriptor_b14583d3ecc23381)
}

var fileDescriptor_b14583d3ecc23381 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0x8e, 0x13, 0x31,
	0x10, 0x3d, 0x93, 0xdc, 0x26, 0xeb, 0x03, 0x09, 0x59, 0x27, 0xb4, 0x17, 0x10, 0x5a, 0x5d, 0xc3,
	0x16, 0xe0, 0x85, 0xd0, 0x50, 0x72, 0x29, 0xd0, 0x05, 0x1a, 0x64, 0xa0, 0xa1, 0x59, 0x6d, 0xac,
	0xd9, 0xc5, 0xec, 0x5e, 0xc6, 0x1a, 0x3b, 0xe1, 0xf8, 0x1f, 0x0a, 0xbe, 0x8b, 0x1f, 0xa0, 0xa5,
	0x44, 0x6b, 0x27, 0x88, 0x02, 0x4e, 0xa9, 0x66, 0xde, 0xd3, 0x9b, 0x99, 0x37, 0xf6, 0xf0, 0xd7,
	0xad, 0xf1, 0x9f, 0x36, 0x2b, 0xa9, 0xf1, 0xaa, 0x74, 0xd8, 0xe3, 0x13, 0x83, 0x65, 0xdb, 0x23,
	0x96, 0x96, 0xf0, 0x33, 0x68, 0xef, 0x22, 0xaa, 0xad, 0x29, 0xb7, 0xcf, 0x4a, 0xb4, 0xde, 0xe0,
	0xda, 0x95, 0xf5, 0x17, 0x57, 0x82, 0x9e, 0x0f, 0xb1, 0x02, 0x3d, 0x97, 0x96, 0xd0, 0xa3, 0x78,
	0xb0, 0x87, 0x3b, 0x99, 0x1c, 0x4a, 0xe5, 0xd0, 0x55, 0x1a, 0x9c, 0x9d, 0xb6, 0xd8, 0x62, 0x10,
	0x96, 0x43, 0x16, 0x6b, 0x66, 0x02, 0xae, 0x7d, 0x24, 0xe1, 0xda, 0xef, 0xb8, 0xc7, 0xff, 0xf0,
	0x14, 0x62, 0x67, 0xfc, 0xde, 0x09, 0x41, 0x13, 0xd5, 0xe7, 0x3f, 0x19, 0xbf, 0xfd, 0xc1, 0x3a,
	0x4f, 0x50, 0x5f, 0xbd, 0xb3, 0xa0, 0xc5, 0x3d, 0x9e, 0x10, 0xb4, 0x06, 0xd7, 0x19, 0xcb, 0x59,
	0x91, 0xaa, 0x1d, 0x12, 0x2f, 0x38, 0x77, 0xa0, 0x09, 0x7c, 0x45, 0xd0, 0x64, 0xb7, 0x72, 0x56,
	0x9c, 0xcc, 0xcf, 0xa4, 0x46, 0x82, 0xbd, 0x47, 0xa9, 0xc0, 0xe1, 0x86, 0x34, 0x28, 0x68, 0x54,
	0x1a, 0xc5, 0x0a, 0x1a, 0x71, 0xc6, 0xa7, 0x84, 0x3d, 0x54, 0x35, 0xad, 0xb3, 0x49, 0xe8, 0x39,
	0x19, 0xf0, 0x05, 0xad, 0xc5, 0x05, 0x9f, 0x34, 0xa6, 0xf7, 0x40, 0x2e, 0x1b, 0xe5, 0xa3, 0xe2,
	0x64, 0xfe, 0x48, 0xde, 0xf4, 0x0a, 0xf2, 0x7d, 0xdd, 0xbe, 0x0a, 0x7a, 0xb5, 0xaf, 0x13, 0xf7,
	0x79, 0x6a, 0x37, 0xab, 0xde, 0xe8, 0xca, 0xd8, 0x6c, 0x9c, 0xb3, 0x62, 0xaa, 0xa6, 0x91, 0x58,
	0x5a, 0x21, 0xf8, 0xd8, 0x22, 0xf9, 0xec, 0x38, 0x67, 0xc5, 0x1d, 0x15, 0xf2, 0xf3, 0x6f, 0x8c,
	0xa7, 0x7f, 0xfa, 0x08, 0xc1, 0x47, 0x1d, 0x7c, 0x8d, 0xbb, 0x5e, 0x1e, 0xa9, 0x01, 0x88, 0x25,
	0x9f, 0x74, 0xdb, 0xca, 0xd6, 0x86, 0x76, 0x7b, 0xca, 0x03, 0x5d, 0xc9, 0x37, 0xdb, 0xb7, 0xb5,
	0xa1, 0xcb, 0x23, 0x95, 0x74, 0x21, 0x9b, 0x3d, 0xe5, 0x49, 0xe4, 0xc4, 0xdd, 0xbf, 0x06, 0xc5,
	0x31, 0xa7, 0xfc, 0x78, 0x5b, 0xf7, 0x1b, 0x08, 0x43, 0x52, 0x15, 0xc1, 0x22, 0xe1, 0x63, 0x67,
	0x41, 0x2f, 0x96, 0xbf, 0x16, 0xec, 0xfb, 0x8f, 0x87, 0xec, 0xe3, 0xcb, 0xc3, 0x6e, 0xcc, 0x76,
	0xed, 0x7f, 0xee, 0x6c, 0x95, 0x84, 0xaf, 0x7e, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x00,
	0x26, 0xd9, 0xae, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if this.Region != that1.Region {
		return false
	}
	if !this.SecretRef.Equal(that1.SecretRef) {
		return false
	}
	if this.RoleArn != that1.RoleArn {
		return false
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if !this.Filters[i].Equal(that1.Filters[i]) {
			return false
		}
	}
	if this.PublicIp != that1.PublicIp {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter)
	if !ok {
		that2, ok := that.(TagFilter)
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
	if that1.Spec == nil {
		if this.Spec != nil {
			return false
		}
	} else if this.Spec == nil {
		return false
	} else if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TagFilter_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_Key)
	if !ok {
		that2, ok := that.(TagFilter_Key)
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
	if this.Key != that1.Key {
		return false
	}
	return true
}
func (this *TagFilter_KvPair_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair_)
	if !ok {
		that2, ok := that.(TagFilter_KvPair_)
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
	if !this.KvPair.Equal(that1.KvPair) {
		return false
	}
	return true
}
func (this *TagFilter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TagFilter_KvPair)
	if !ok {
		that2, ok := that.(TagFilter_KvPair)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
