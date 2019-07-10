// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/ec2/aws_ec2.proto

package glooec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	// The AWS Region where the desired EC2 instances exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  ```
	SecretRef core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref"`
	Filters   []*Filter        `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
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
	return fileDescriptor_fc1fd6f1173c4563, []int{0}
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

func (m *UpstreamSpec) GetSecretRef() core.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return core.ResourceRef{}
}

func (m *UpstreamSpec) GetFilters() []*Filter {
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

type Filter struct {
	// Types that are valid to be assigned to Spec:
	//	*Filter_Key
	//	*Filter_KvPair_
	Spec                 isFilter_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{1}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

type isFilter_Spec interface {
	isFilter_Spec()
	Equal(interface{}) bool
}

type Filter_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}
type Filter_KvPair_ struct {
	KvPair *Filter_KvPair `protobuf:"bytes,2,opt,name=kv_pair,json=kvPair,proto3,oneof"`
}

func (*Filter_Key) isFilter_Spec()     {}
func (*Filter_KvPair_) isFilter_Spec() {}

func (m *Filter) GetSpec() isFilter_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Filter) GetKey() string {
	if x, ok := m.GetSpec().(*Filter_Key); ok {
		return x.Key
	}
	return ""
}

func (m *Filter) GetKvPair() *Filter_KvPair {
	if x, ok := m.GetSpec().(*Filter_KvPair_); ok {
		return x.KvPair
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Filter_OneofMarshaler, _Filter_OneofUnmarshaler, _Filter_OneofSizer, []interface{}{
		(*Filter_Key)(nil),
		(*Filter_KvPair_)(nil),
	}
}

func _Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Filter)
	// spec
	switch x := m.Spec.(type) {
	case *Filter_Key:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Key)
	case *Filter_KvPair_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KvPair); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Filter.Spec has unexpected type %T", x)
	}
	return nil
}

func _Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Filter)
	switch tag {
	case 1: // spec.key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Spec = &Filter_Key{x}
		return true, err
	case 2: // spec.kv_pair
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Filter_KvPair)
		err := b.DecodeMessage(msg)
		m.Spec = &Filter_KvPair_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Filter)
	// spec
	switch x := m.Spec.(type) {
	case *Filter_Key:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Key)))
		n += len(x.Key)
	case *Filter_KvPair_:
		s := proto.Size(x.KvPair)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Filter_KvPair struct {
	// keys are not case-sensitive, as with AWS Condition Keys
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter_KvPair) Reset()         { *m = Filter_KvPair{} }
func (m *Filter_KvPair) String() string { return proto.CompactTextString(m) }
func (*Filter_KvPair) ProtoMessage()    {}
func (*Filter_KvPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc1fd6f1173c4563, []int{1, 0}
}
func (m *Filter_KvPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter_KvPair.Unmarshal(m, b)
}
func (m *Filter_KvPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter_KvPair.Marshal(b, m, deterministic)
}
func (m *Filter_KvPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter_KvPair.Merge(m, src)
}
func (m *Filter_KvPair) XXX_Size() int {
	return xxx_messageInfo_Filter_KvPair.Size(m)
}
func (m *Filter_KvPair) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter_KvPair.DiscardUnknown(m)
}

var xxx_messageInfo_Filter_KvPair proto.InternalMessageInfo

func (m *Filter_KvPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Filter_KvPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "aws_ec2.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*Filter)(nil), "aws_ec2.plugins.gloo.solo.io.Filter")
	proto.RegisterType((*Filter_KvPair)(nil), "aws_ec2.plugins.gloo.solo.io.Filter.KvPair")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/aws/ec2/aws_ec2.proto", fileDescriptor_fc1fd6f1173c4563)
}

var fileDescriptor_fc1fd6f1173c4563 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xad, 0x69, 0x36, 0xbb, 0xf1, 0x82, 0x84, 0xac, 0x15, 0x0a, 0x05, 0x41, 0xb4, 0xe2, 0x10,
	0x09, 0xb0, 0x21, 0xdc, 0xf7, 0x50, 0xa1, 0xd5, 0x52, 0x2e, 0xc8, 0x88, 0x0b, 0x97, 0x28, 0xb5,
	0x26, 0xc1, 0x24, 0xed, 0x58, 0xb6, 0x13, 0xc4, 0xcf, 0x70, 0xe6, 0x53, 0xf8, 0x03, 0x6e, 0x1c,
	0xf8, 0x12, 0x94, 0xb8, 0x45, 0x1c, 0x2a, 0xd4, 0xd3, 0xcc, 0x1b, 0xbf, 0x37, 0x6f, 0x46, 0x1e,
	0xba, 0x6a, 0xb4, 0xff, 0xd4, 0xaf, 0xb9, 0xc2, 0x8d, 0x70, 0xd8, 0xe1, 0x73, 0x8d, 0xa2, 0xe9,
	0x10, 0x85, 0xb1, 0xf8, 0x19, 0x94, 0x77, 0x01, 0x55, 0x46, 0x8b, 0xe1, 0xa5, 0x30, 0x5d, 0xdf,
	0xe8, 0xad, 0x13, 0xd5, 0x17, 0x27, 0x40, 0x15, 0x63, 0x2c, 0x41, 0x15, 0xdc, 0x58, 0xf4, 0xc8,
	0x1e, 0xfe, 0x85, 0x81, 0xc6, 0x47, 0x29, 0x1f, 0xbb, 0x72, 0x8d, 0x8b, 0x8b, 0x06, 0x1b, 0x9c,
	0x88, 0x62, 0xcc, 0x82, 0x66, 0xf1, 0xec, 0x80, 0xff, 0x14, 0x5b, 0xed, 0xf7, 0xae, 0x16, 0xea,
	0xc0, 0xbe, 0xfc, 0x49, 0xe8, 0xed, 0x0f, 0xc6, 0x79, 0x0b, 0xd5, 0xe6, 0xbd, 0x01, 0xc5, 0xee,
	0xd1, 0xd8, 0x42, 0xa3, 0x71, 0x9b, 0x92, 0x8c, 0xe4, 0x89, 0xdc, 0x21, 0x76, 0x45, 0xa9, 0x03,
	0x65, 0xc1, 0x97, 0x16, 0xea, 0xf4, 0x56, 0x46, 0xf2, 0xf3, 0xe2, 0x3e, 0x57, 0x68, 0x61, 0x3f,
	0x0f, 0x97, 0xe0, 0xb0, 0xb7, 0x0a, 0x24, 0xd4, 0xcb, 0xe8, 0xc7, 0xaf, 0xc7, 0x33, 0x99, 0x04,
	0x89, 0x84, 0x9a, 0x5d, 0xd1, 0xd3, 0x5a, 0x77, 0x1e, 0xac, 0x4b, 0xe7, 0xd9, 0x3c, 0x3f, 0x2f,
	0x9e, 0xf0, 0xff, 0x2d, 0xc7, 0xaf, 0x27, 0xb2, 0xdc, 0x8b, 0xd8, 0x03, 0x9a, 0x98, 0x7e, 0xdd,
	0x69, 0x55, 0x6a, 0x93, 0x46, 0x19, 0xc9, 0xcf, 0xe4, 0x59, 0x28, 0xbc, 0x31, 0x8c, 0xd1, 0xc8,
	0xa0, 0xf5, 0xe9, 0x49, 0x46, 0xf2, 0x3b, 0x72, 0xca, 0x2f, 0xbf, 0x11, 0x1a, 0x87, 0x26, 0x8c,
	0xd1, 0x79, 0x0b, 0x5f, 0xc3, 0x42, 0x37, 0x33, 0x39, 0x02, 0x76, 0x4d, 0x4f, 0xdb, 0xa1, 0x34,
	0x95, 0xb6, 0xbb, 0x65, 0x9e, 0x1e, 0x33, 0x0f, 0x7f, 0x3b, 0xbc, 0xab, 0xb4, 0xbd, 0x99, 0xc9,
	0xb8, 0x9d, 0xb2, 0xc5, 0x0b, 0x1a, 0x87, 0x1a, 0xbb, 0xfb, 0x8f, 0x4b, 0xf0, 0xb8, 0xa0, 0x27,
	0x43, 0xd5, 0xf5, 0x30, 0x39, 0x24, 0x32, 0x80, 0x65, 0x4c, 0x23, 0x67, 0x40, 0x2d, 0x57, 0xdf,
	0x7f, 0x3f, 0x22, 0x1f, 0x5f, 0x1f, 0x77, 0x2e, 0xa6, 0x6d, 0x0e, 0x9d, 0xcc, 0xf8, 0x06, 0xaa,
	0x58, 0xc7, 0xd3, 0x6f, 0xbe, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x73, 0x42, 0x2d, 0xab, 0x7d,
	0x02, 0x00, 0x00,
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
func (this *Filter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Filter)
	if !ok {
		that2, ok := that.(Filter)
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
func (this *Filter_Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Filter_Key)
	if !ok {
		that2, ok := that.(Filter_Key)
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
func (this *Filter_KvPair_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Filter_KvPair_)
	if !ok {
		that2, ok := that.(Filter_KvPair_)
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
func (this *Filter_KvPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Filter_KvPair)
	if !ok {
		that2, ok := that.(Filter_KvPair)
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
