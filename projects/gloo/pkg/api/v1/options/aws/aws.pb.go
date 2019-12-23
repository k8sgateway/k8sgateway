// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto

package aws

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
	return fileDescriptor_7e8a9525eed72921, []int{2, 0}
}

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired Lambda Functions exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// The list of Lambda Functions contained within this region.
	// This list will be automatically populated by Gloo if discovery is enabled for AWS Lambda Functions
	LambdaFunctions      []*LambdaFunctionSpec `protobuf:"bytes,3,rep,name=lambda_functions,json=lambdaFunctions,proto3" json:"lambda_functions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8a9525eed72921, []int{0}
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
	return fileDescriptor_7e8a9525eed72921, []int{1}
}
func (m *LambdaFunctionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LambdaFunctionSpec.Unmarshal(m, b)
}
func (m *LambdaFunctionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LambdaFunctionSpec.Marshal(b, m, deterministic)
}
func (m *LambdaFunctionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LambdaFunctionSpec.Merge(m, src)
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
	InvocationStyle DestinationSpec_InvocationStyle `protobuf:"varint,2,opt,name=invocation_style,json=invocationStyle,proto3,enum=aws.options.gloo.solo.io.DestinationSpec_InvocationStyle" json:"invocation_style,omitempty"`
	// de-jsonify response bodies returned from aws lambda
	ResponseTransformation bool     `protobuf:"varint,5,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8a9525eed72921, []int{2}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
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

func (m *DestinationSpec) GetResponseTransformation() bool {
	if m != nil {
		return m.ResponseTransformation
	}
	return false
}

func init() {
	proto.RegisterEnum("aws.options.gloo.solo.io.DestinationSpec_InvocationStyle", DestinationSpec_InvocationStyle_name, DestinationSpec_InvocationStyle_value)
	proto.RegisterType((*UpstreamSpec)(nil), "aws.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*LambdaFunctionSpec)(nil), "aws.options.gloo.solo.io.LambdaFunctionSpec")
	proto.RegisterType((*DestinationSpec)(nil), "aws.options.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto", fileDescriptor_7e8a9525eed72921)
}

var fileDescriptor_7e8a9525eed72921 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xd9, 0xa6, 0xad, 0x9a, 0x49, 0x45, 0xa2, 0x55, 0x55, 0x4c, 0x85, 0x50, 0xc8, 0x01,
	0xe5, 0x50, 0xd6, 0x10, 0x0e, 0x80, 0xc4, 0x85, 0x82, 0x2a, 0x21, 0xa1, 0x1e, 0x1c, 0x10, 0x82,
	0x8b, 0xb5, 0xd9, 0x8e, 0xcd, 0x52, 0xdb, 0xb3, 0xec, 0x6e, 0x4a, 0x79, 0x02, 0x5e, 0x85, 0x47,
	0x80, 0xd7, 0xe1, 0x1d, 0x38, 0x71, 0x41, 0x5e, 0x3b, 0xa0, 0x84, 0x46, 0x82, 0x83, 0xe5, 0xd9,
	0x7f, 0xbe, 0x99, 0xf9, 0xc7, 0xf2, 0xc2, 0x51, 0xae, 0xfd, 0xbb, 0xf9, 0x4c, 0x28, 0x2a, 0x63,
	0x47, 0x05, 0xdd, 0xd1, 0x14, 0xe7, 0x05, 0x51, 0x6c, 0x2c, 0xbd, 0x47, 0xe5, 0x5d, 0x73, 0x92,
	0x46, 0xc7, 0xe7, 0xf7, 0x62, 0x32, 0x5e, 0x53, 0xe5, 0x62, 0xf9, 0x31, 0x3c, 0xc2, 0x58, 0xf2,
	0xc4, 0xa3, 0x3a, 0x6c, 0x53, 0xa2, 0xc6, 0x45, 0xdd, 0x49, 0x68, 0x3a, 0xd8, 0xcb, 0x29, 0xa7,
	0x00, 0xc5, 0x75, 0xd4, 0xf0, 0x07, 0x1c, 0x2f, 0x7c, 0x23, 0xe2, 0x85, 0x6f, 0xb5, 0xc3, 0x4b,
	0x7c, 0x84, 0xf7, 0x99, 0xf6, 0x8b, 0xe9, 0x16, 0xb3, 0x86, 0x1e, 0x7d, 0x63, 0xb0, 0xfb, 0xca,
	0x38, 0x6f, 0x51, 0x96, 0x53, 0x83, 0x8a, 0xef, 0xc3, 0xb6, 0xc5, 0x5c, 0x53, 0x15, 0xb1, 0x21,
	0x1b, 0x77, 0x93, 0xf6, 0xc4, 0x1f, 0x02, 0x38, 0x54, 0x16, 0x7d, 0x6a, 0x31, 0x8b, 0x36, 0x86,
	0x6c, 0xdc, 0x9b, 0x5c, 0x17, 0x8a, 0x2c, 0x2e, 0x3c, 0x8a, 0x04, 0x1d, 0xcd, 0xad, 0xc2, 0x04,
	0xb3, 0xa4, 0xdb, 0xc0, 0x09, 0x66, 0xfc, 0x35, 0x0c, 0x0a, 0x59, 0xce, 0x4e, 0x65, 0x9a, 0xcd,
	0x2b, 0x15, 0x76, 0x8b, 0x3a, 0xc3, 0xce, 0xb8, 0x37, 0x39, 0x14, 0xeb, 0xf6, 0x15, 0x2f, 0x42,
	0xc5, 0x71, 0x5b, 0x50, 0x3b, 0x4b, 0xfa, 0xc5, 0x92, 0xe6, 0x46, 0x9f, 0x19, 0xf0, 0xbf, 0x39,
	0x7e, 0x0b, 0x76, 0x0b, 0xca, 0xb5, 0x92, 0x45, 0x5a, 0xc9, 0x12, 0xdb, 0x3d, 0x7a, 0xad, 0x76,
	0x22, 0x4b, 0xe4, 0x77, 0x61, 0x6f, 0xc5, 0x52, 0x83, 0x6e, 0x04, 0x94, 0x2f, 0x0f, 0x0a, 0x15,
	0x37, 0xa0, 0xfb, 0x61, 0x2e, 0x0b, 0x9d, 0x69, 0xb4, 0x51, 0x27, 0x60, 0x7f, 0x84, 0xd1, 0x4f,
	0x06, 0xfd, 0x67, 0xe8, 0xbc, 0xae, 0xe4, 0xff, 0xd8, 0x38, 0x85, 0x81, 0xae, 0xce, 0x49, 0x85,
	0xa2, 0xd4, 0xf9, 0x4f, 0x45, 0x63, 0xe1, 0xea, 0xe4, 0xd1, 0xfa, 0x2f, 0xb3, 0x32, 0x47, 0x3c,
	0xff, 0xdd, 0x61, 0x5a, 0x37, 0x48, 0xfa, 0x7a, 0x59, 0xe0, 0x0f, 0xe0, 0x9a, 0x45, 0x67, 0xa8,
	0x72, 0x98, 0x7a, 0x2b, 0x2b, 0x97, 0x91, 0x2d, 0x43, 0x3e, 0xda, 0x1a, 0xb2, 0xf1, 0x4e, 0xb2,
	0xbf, 0x48, 0xbf, 0x5c, 0xca, 0x8e, 0x6e, 0x43, 0x7f, 0xa5, 0x39, 0xdf, 0x81, 0xcd, 0xe9, 0x9b,
	0x93, 0xa7, 0x83, 0x2b, 0xbc, 0x0b, 0x5b, 0x4f, 0x42, 0xc8, 0x8e, 0x8e, 0xbf, 0xfe, 0xd8, 0x64,
	0x5f, 0xbe, 0xdf, 0x64, 0x6f, 0x1f, 0xff, 0xdb, 0x1d, 0x30, 0x67, 0xf9, 0x25, 0xf7, 0x60, 0xb6,
	0x1d, 0x7e, 0xc9, 0xfb, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x22, 0xe4, 0x74, 0x72, 0x4a, 0x03,
	0x00, 0x00,
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
	if this.ResponseTransformation != that1.ResponseTransformation {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
