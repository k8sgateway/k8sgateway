// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws.proto

package aws // import "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DestinationSpec_InvocationStyle int32

const (
	DestinationSpec_SYNC  DestinationSpec_InvocationStyle = 0
	DestinationSpec_ASYNC DestinationSpec_InvocationStyle = 1
)

var DestinationSpec_InvocationStyle_name = map[int32]string{
	0: "SYNC",
	1: "ASYNC",
}
var DestinationSpec_InvocationStyle_value = map[string]int32{
	"SYNC":  0,
	"ASYNC": 1,
}

func (x DestinationSpec_InvocationStyle) String() string {
	return proto.EnumName(DestinationSpec_InvocationStyle_name, int32(x))
}
func (DestinationSpec_InvocationStyle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aws_18e32ef2807acd5b, []int{2, 0}
}

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired Lambda Functions exxist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	SecretRef core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef" json:"secret_ref"`
	// The list of Lambda Functions contained within this region.
	// This list will be automatically populated by Gloo if discovery is enabled for AWS Lambda Functions
	LambdaFunctions      []*LambdaFunctionSpec `protobuf:"bytes,3,rep,name=lambda_functions,json=lambdaFunctions" json:"lambda_functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_aws_18e32ef2807acd5b, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (dst *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(dst, src)
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

func (m *UpstreamSpec) GetSecretRef() core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return core.ResourceRef{}
}

func (m *UpstreamSpec) GetLambdaFunctions() []*LambdaFunctionSpec {
	if m != nil {
		return m.LambdaFunctions
	}
	return nil
}

// Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions:
// - name of the function
// - qualifier for the function
type LambdaFunctionSpec struct {
	// the logical name gloo should associate with this function. if left empty, it will default to
	// lambda_function_name+qualifier
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// The Name of the Lambda Function as it appears in the AWS Lambda Portal
	LambdaFunctionName string `protobuf:"bytes,2,opt,name=lambda_function_name,json=lambdaFunctionName,proto3" json:"lambda_function_name,omitempty"`
	// The Qualifier for the Lambda Function. Qualifiers act as a kind of version
	// for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info.
	Qualifier            string   `protobuf:"bytes,3,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LambdaFunctionSpec) Reset()         { *m = LambdaFunctionSpec{} }
func (m *LambdaFunctionSpec) String() string { return proto.CompactTextString(m) }
func (*LambdaFunctionSpec) ProtoMessage()    {}
func (*LambdaFunctionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_aws_18e32ef2807acd5b, []int{1}
}
func (m *LambdaFunctionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaFunctionSpec.Unmarshal(m, b)
}
func (m *LambdaFunctionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaFunctionSpec.Marshal(b, m, deterministic)
}
func (dst *LambdaFunctionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaFunctionSpec.Merge(dst, src)
}
func (m *LambdaFunctionSpec) XXX_Size() int {
	return xxx_messageInfo_LambdaFunctionSpec.Size(m)
}
func (m *LambdaFunctionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_LambdaFunctionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_LambdaFunctionSpec proto.InternalMessageInfo

func (m *LambdaFunctionSpec) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func (m *LambdaFunctionSpec) GetLambdaFunctionName() string {
	if m != nil {
		return m.LambdaFunctionName
	}
	return ""
}

func (m *LambdaFunctionSpec) GetQualifier() string {
	if m != nil {
		return m.Qualifier
	}
	return ""
}

// Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions
type DestinationSpec struct {
	// The Logical Name of the LambdaFunctionSpec to be invoked.
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// Can be either Sync or Async.
	InvocationStyle DestinationSpec_InvocationStyle `protobuf:"varint,2,opt,name=invocation_style,json=invocationStyle,proto3,enum=aws.plugins.gloo.solo.io.DestinationSpec_InvocationStyle" json:"invocation_style,omitempty"`
	// de-jsonify response bodies returned from aws lambda
	ResponseTrasnformation bool     `protobuf:"varint,5,opt,name=response_trasnformation,json=responseTrasnformation,proto3" json:"response_trasnformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_aws_18e32ef2807acd5b, []int{2}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (dst *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(dst, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

func (m *DestinationSpec) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func (m *DestinationSpec) GetInvocationStyle() DestinationSpec_InvocationStyle {
	if m != nil {
		return m.InvocationStyle
	}
	return DestinationSpec_SYNC
}

func (m *DestinationSpec) GetResponseTrasnformation() bool {
	if m != nil {
		return m.ResponseTrasnformation
	}
	return false
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*LambdaFunctionSpec)(nil), "aws.plugins.gloo.solo.io.LambdaFunctionSpec")
	proto.RegisterType((*DestinationSpec)(nil), "aws.plugins.gloo.solo.io.DestinationSpec")
	proto.RegisterEnum("aws.plugins.gloo.solo.io.DestinationSpec_InvocationStyle", DestinationSpec_InvocationStyle_name, DestinationSpec_InvocationStyle_value)
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
	if !this.SecretRef.Equal(&that1.SecretRef) {
		return false
	}
	if len(this.LambdaFunctions) != len(that1.LambdaFunctions) {
		return false
	}
	for i := range this.LambdaFunctions {
		if !this.LambdaFunctions[i].Equal(that1.LambdaFunctions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LambdaFunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LambdaFunctionSpec)
	if !ok {
		that2, ok := that.(LambdaFunctionSpec)
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
	if this.LogicalName != that1.LogicalName {
		return false
	}
	if this.LambdaFunctionName != that1.LambdaFunctionName {
		return false
	}
	if this.Qualifier != that1.Qualifier {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if this.LogicalName != that1.LogicalName {
		return false
	}
	if this.InvocationStyle != that1.InvocationStyle {
		return false
	}
	if this.ResponseTrasnformation != that1.ResponseTrasnformation {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/aws.proto", fileDescriptor_aws_18e32ef2807acd5b)
}

var fileDescriptor_aws_18e32ef2807acd5b = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xeb, 0xa6, 0xad, 0x1a, 0xa7, 0x22, 0x91, 0x55, 0x95, 0xa5, 0x42, 0x10, 0x72, 0x40,
	0x39, 0x14, 0x2f, 0x84, 0x03, 0x42, 0x42, 0x48, 0x04, 0x84, 0x84, 0x84, 0x7a, 0xd8, 0x80, 0x10,
	0x5c, 0x56, 0x8e, 0x3b, 0xbb, 0x98, 0x7a, 0x3d, 0xc6, 0xf6, 0x16, 0xf1, 0x04, 0xbc, 0x0a, 0x6f,
	0xc1, 0x95, 0xa7, 0xe0, 0xc0, 0x6b, 0x70, 0x41, 0xeb, 0x4d, 0x40, 0x09, 0xad, 0x44, 0x0f, 0xab,
	0x1d, 0xcf, 0x7c, 0xff, 0xcc, 0x6f, 0x6b, 0xe8, 0xb4, 0x54, 0xe1, 0x7d, 0x3d, 0xe7, 0x12, 0xab,
	0xd4, 0xa3, 0xc6, 0x3b, 0x0a, 0xd3, 0x52, 0x23, 0xa6, 0xd6, 0xe1, 0x07, 0x90, 0xc1, 0xb7, 0x27,
	0x61, 0x55, 0x7a, 0x76, 0x2f, 0xb5, 0xba, 0x2e, 0x95, 0xf1, 0xa9, 0xf8, 0x14, 0x3f, 0x6e, 0x1d,
	0x06, 0x64, 0x49, 0x0c, 0xdb, 0x12, 0x6f, 0x70, 0xde, 0x74, 0xe2, 0x0a, 0x0f, 0xf7, 0x4b, 0x2c,
	0x31, 0x42, 0x69, 0x13, 0xb5, 0xfc, 0xe1, 0xd1, 0x39, 0x33, 0xe3, 0xff, 0x54, 0x85, 0xe5, 0x24,
	0x07, 0x45, 0x4b, 0x8f, 0xbe, 0x11, 0xba, 0xf7, 0xda, 0xfa, 0xe0, 0x40, 0x54, 0x33, 0x0b, 0x92,
	0x1d, 0xd0, 0x1d, 0x07, 0xa5, 0x42, 0x93, 0x90, 0x21, 0x19, 0x77, 0xb3, 0xc5, 0x89, 0x3d, 0xa6,
	0xd4, 0x83, 0x74, 0x10, 0x72, 0x07, 0x45, 0xb2, 0x39, 0x24, 0xe3, 0xde, 0xe4, 0x1a, 0x97, 0xe8,
	0x60, 0xe9, 0x87, 0x67, 0xe0, 0xb1, 0x76, 0x12, 0x32, 0x28, 0xa6, 0x5b, 0xdf, 0x7f, 0xdc, 0xdc,
	0xc8, 0xba, 0xad, 0x24, 0x83, 0x82, 0xbd, 0xa1, 0x03, 0x2d, 0xaa, 0xf9, 0x89, 0xc8, 0x8b, 0xda,
	0xc8, 0xa0, 0xd0, 0xf8, 0xa4, 0x33, 0xec, 0x8c, 0x7b, 0x93, 0x23, 0x7e, 0xd1, 0x0d, 0xf9, 0xcb,
	0xa8, 0x78, 0xbe, 0x10, 0x34, 0xfe, 0xb2, 0xbe, 0x5e, 0xc9, 0xf9, 0xd1, 0x17, 0x42, 0xd9, 0xbf,
	0x1c, 0xbb, 0x45, 0xf7, 0x34, 0x96, 0x4a, 0x0a, 0x9d, 0x1b, 0x51, 0xc1, 0xe2, 0x36, 0xbd, 0x45,
	0xee, 0x58, 0x54, 0xc0, 0xee, 0xd2, 0xfd, 0x35, 0x4b, 0x2d, 0xba, 0x19, 0x51, 0xb6, 0x3a, 0x28,
	0x2a, 0xae, 0xd3, 0xee, 0xc7, 0x5a, 0x68, 0x55, 0x28, 0x70, 0x49, 0x27, 0x62, 0x7f, 0x13, 0xa3,
	0x5f, 0x84, 0xf6, 0x9f, 0x81, 0x0f, 0xca, 0x88, 0xcb, 0xd8, 0x38, 0xa1, 0x03, 0x65, 0xce, 0x50,
	0x46, 0x51, 0xee, 0xc3, 0x67, 0xdd, 0x5a, 0xb8, 0x32, 0x79, 0x78, 0xf1, 0xcb, 0xac, 0xcd, 0xe1,
	0x2f, 0xfe, 0x74, 0x98, 0x35, 0x0d, 0xb2, 0xbe, 0x5a, 0x4d, 0xb0, 0x07, 0xf4, 0xaa, 0x03, 0x6f,
	0xd1, 0x78, 0xc8, 0x83, 0x13, 0xde, 0x14, 0xe8, 0xaa, 0x58, 0x4f, 0xb6, 0x87, 0x64, 0xbc, 0x9b,
	0x1d, 0x2c, 0xcb, 0xaf, 0x56, 0xaa, 0xa3, 0xdb, 0xb4, 0xbf, 0xd6, 0x9c, 0xed, 0xd2, 0xad, 0xd9,
	0xdb, 0xe3, 0xa7, 0x83, 0x0d, 0xd6, 0xa5, 0xdb, 0x4f, 0x62, 0x48, 0xa6, 0xd3, 0xaf, 0x3f, 0x6f,
	0x90, 0x77, 0x8f, 0xfe, 0x6f, 0xe3, 0xed, 0x69, 0x79, 0xce, 0xd6, 0xcf, 0x77, 0xe2, 0x52, 0xde,
	0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x7c, 0x67, 0x64, 0x38, 0x03, 0x00, 0x00,
}
