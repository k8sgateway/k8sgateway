// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ratelimit.proto

package v1 // import "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/gogo/protobuf/gogoproto"
import ratelimit "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/ratelimit"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import bytes "bytes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Constraint struct {
	Key                  string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	RateLimit            *ratelimit.RateLimit `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit" json:"rate_limit,omitempty"`
	Constraints          []*Constraint        `protobuf:"bytes,4,rep,name=constraints" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Constraint) Reset()         { *m = Constraint{} }
func (m *Constraint) String() string { return proto.CompactTextString(m) }
func (*Constraint) ProtoMessage()    {}
func (*Constraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ea029df602ccc5b3, []int{0}
}
func (m *Constraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constraint.Unmarshal(m, b)
}
func (m *Constraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constraint.Marshal(b, m, deterministic)
}
func (dst *Constraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constraint.Merge(dst, src)
}
func (m *Constraint) XXX_Size() int {
	return xxx_messageInfo_Constraint.Size(m)
}
func (m *Constraint) XXX_DiscardUnknown() {
	xxx_messageInfo_Constraint.DiscardUnknown(m)
}

var xxx_messageInfo_Constraint proto.InternalMessageInfo

func (m *Constraint) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Constraint) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Constraint) GetRateLimit() *ratelimit.RateLimit {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

func (m *Constraint) GetConstraints() []*Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

//
// @solo-kit:xds-service=RateLimitDiscoveryService
// @solo-kit:resource.no_references
type RateLimitConfig struct {
	// @solo-kit:resource.name
	Domain               string        `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Constraints          []*Constraint `protobuf:"bytes,2,rep,name=constraints" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RateLimitConfig) Reset()         { *m = RateLimitConfig{} }
func (m *RateLimitConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfig) ProtoMessage()    {}
func (*RateLimitConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_ea029df602ccc5b3, []int{1}
}
func (m *RateLimitConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfig.Unmarshal(m, b)
}
func (m *RateLimitConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfig.Marshal(b, m, deterministic)
}
func (dst *RateLimitConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfig.Merge(dst, src)
}
func (m *RateLimitConfig) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfig.Size(m)
}
func (m *RateLimitConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfig proto.InternalMessageInfo

func (m *RateLimitConfig) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitConfig) GetConstraints() []*Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

func init() {
	proto.RegisterType((*Constraint)(nil), "gloo.solo.io.Constraint")
	proto.RegisterType((*RateLimitConfig)(nil), "gloo.solo.io.RateLimitConfig")
}
func (this *Constraint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Constraint)
	if !ok {
		that2, ok := that.(Constraint)
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
	if !this.RateLimit.Equal(that1.RateLimit) {
		return false
	}
	if len(this.Constraints) != len(that1.Constraints) {
		return false
	}
	for i := range this.Constraints {
		if !this.Constraints[i].Equal(that1.Constraints[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RateLimitConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitConfig)
	if !ok {
		that2, ok := that.(RateLimitConfig)
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
	if this.Domain != that1.Domain {
		return false
	}
	if len(this.Constraints) != len(that1.Constraints) {
		return false
	}
	for i := range this.Constraints {
		if !this.Constraints[i].Equal(that1.Constraints[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitDiscoveryServiceClient is the client API for RateLimitDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitDiscoveryServiceClient interface {
	StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error)
	IncrementalRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_IncrementalRateLimitConfigClient, error)
	FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type rateLimitDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitDiscoveryServiceClient(cc *grpc.ClientConn) RateLimitDiscoveryServiceClient {
	return &rateLimitDiscoveryServiceClient{cc}
}

func (c *rateLimitDiscoveryServiceClient) StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[0], "/gloo.solo.io.RateLimitDiscoveryService/StreamRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceStreamRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_StreamRateLimitConfigClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) IncrementalRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_IncrementalRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[1], "/gloo.solo.io.RateLimitDiscoveryService/IncrementalRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceIncrementalRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_IncrementalRateLimitConfigClient interface {
	Send(*v2.IncrementalDiscoveryRequest) error
	Recv() (*v2.IncrementalDiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceIncrementalRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigClient) Send(m *v2.IncrementalDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigClient) Recv() (*v2.IncrementalDiscoveryResponse, error) {
	m := new(v2.IncrementalDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/gloo.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitDiscoveryServiceServer is the server API for RateLimitDiscoveryService service.
type RateLimitDiscoveryServiceServer interface {
	StreamRateLimitConfig(RateLimitDiscoveryService_StreamRateLimitConfigServer) error
	IncrementalRateLimitConfig(RateLimitDiscoveryService_IncrementalRateLimitConfigServer) error
	FetchRateLimitConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterRateLimitDiscoveryServiceServer(s *grpc.Server, srv RateLimitDiscoveryServiceServer) {
	s.RegisterService(&_RateLimitDiscoveryService_serviceDesc, srv)
}

func _RateLimitDiscoveryService_StreamRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).StreamRateLimitConfig(&rateLimitDiscoveryServiceStreamRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_StreamRateLimitConfigServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_IncrementalRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).IncrementalRateLimitConfig(&rateLimitDiscoveryServiceIncrementalRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_IncrementalRateLimitConfigServer interface {
	Send(*v2.IncrementalDiscoveryResponse) error
	Recv() (*v2.IncrementalDiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceIncrementalRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigServer) Send(m *v2.IncrementalDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigServer) Recv() (*v2.IncrementalDiscoveryRequest, error) {
	m := new(v2.IncrementalDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_FetchRateLimitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gloo.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gloo.solo.io.RateLimitDiscoveryService",
	HandlerType: (*RateLimitDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRateLimitConfig",
			Handler:    _RateLimitDiscoveryService_FetchRateLimitConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_StreamRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "IncrementalRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_IncrementalRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ratelimit.proto",
}

func init() { proto.RegisterFile("ratelimit.proto", fileDescriptor_ratelimit_ea029df602ccc5b3) }

var fileDescriptor_ratelimit_ea029df602ccc5b3 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xf1, 0x2e, 0x54, 0xaa, 0x17, 0xa9, 0xc8, 0x5a, 0x50, 0x88, 0x4a, 0x59, 0xe5, 0x94,
	0x56, 0xc2, 0x86, 0x70, 0x62, 0x2f, 0x48, 0x14, 0x21, 0x90, 0x38, 0xa5, 0x07, 0x24, 0x0e, 0x20,
	0xd7, 0x1d, 0x5c, 0xd3, 0xc4, 0x13, 0x6c, 0x6f, 0xa4, 0x3d, 0xc2, 0x2b, 0xf0, 0x12, 0x3c, 0x03,
	0xaf, 0xc1, 0x1b, 0x20, 0x9e, 0x82, 0x13, 0x4a, 0xb2, 0x0d, 0xbb, 0x41, 0x48, 0xac, 0xc4, 0x29,
	0xbf, 0x3d, 0xe3, 0xf9, 0x7e, 0xc7, 0x33, 0x74, 0xcf, 0xc9, 0x00, 0x85, 0x29, 0x4d, 0xe0, 0x95,
	0xc3, 0x80, 0xec, 0xba, 0x2e, 0x10, 0xb9, 0xc7, 0x02, 0xb9, 0xc1, 0x78, 0x1f, 0x6c, 0x8d, 0x4b,
	0x21, 0x2b, 0x23, 0xea, 0x4c, 0x9c, 0x19, 0xaf, 0xb0, 0x06, 0xb7, 0xec, 0x72, 0xe3, 0x7d, 0x8d,
	0xa8, 0x0b, 0x68, 0xc3, 0xd2, 0x5a, 0x0c, 0x32, 0x18, 0xb4, 0x7e, 0x15, 0x7d, 0xa5, 0x4d, 0x38,
	0x5f, 0x9c, 0x72, 0x85, 0xa5, 0x68, 0xea, 0xdd, 0x33, 0xd8, 0x7d, 0x2b, 0x87, 0xef, 0x41, 0x05,
	0x2f, 0x7a, 0xd1, 0x30, 0x3b, 0xc8, 0x03, 0x51, 0x15, 0x0b, 0x6d, 0xac, 0x17, 0xbd, 0x31, 0x31,
	0xb0, 0x18, 0x4f, 0x35, 0x6a, 0x6c, 0xa5, 0x68, 0x54, 0xb7, 0x9b, 0x7c, 0x25, 0x94, 0x1e, 0xa3,
	0xf5, 0xc1, 0x49, 0x63, 0x03, 0xbb, 0x41, 0xc7, 0x17, 0xb0, 0x8c, 0xc8, 0x8c, 0xa4, 0xbb, 0x79,
	0x23, 0xd9, 0x94, 0x5e, 0xab, 0x65, 0xb1, 0x80, 0x68, 0xd4, 0xee, 0x75, 0x0b, 0xf6, 0x9c, 0xd2,
	0xa6, 0xfe, 0xdb, 0x16, 0x10, 0x8d, 0x67, 0x24, 0x9d, 0x64, 0x87, 0x7c, 0x0d, 0xd9, 0xd9, 0xe1,
	0xeb, 0xbf, 0x85, 0xe7, 0x32, 0xc0, 0xcb, 0x26, 0x9c, 0xef, 0xba, 0x4b, 0xc9, 0xe6, 0x74, 0xa2,
	0x7a, 0xbe, 0x8f, 0xae, 0xce, 0xc6, 0xe9, 0x24, 0x8b, 0x36, 0x0f, 0xfe, 0x36, 0x98, 0xaf, 0x27,
	0x27, 0x40, 0xf7, 0xfa, 0x9a, 0xc7, 0x68, 0xdf, 0x19, 0xcd, 0x6e, 0xd1, 0x9d, 0x33, 0x2c, 0xa5,
	0xb1, 0xab, 0x3b, 0xac, 0x56, 0x43, 0xcc, 0x68, 0x0b, 0x4c, 0xf6, 0x73, 0x44, 0x6f, 0xf7, 0x9c,
	0xa7, 0x97, 0xaf, 0x79, 0x02, 0xae, 0x36, 0x0a, 0xd8, 0x1b, 0x7a, 0xf3, 0x24, 0x38, 0x90, 0xe5,
	0xd0, 0xca, 0x01, 0x6f, 0xdb, 0x80, 0xcb, 0xca, 0xf0, 0x3a, 0xe3, 0xfd, 0xc1, 0x1c, 0x3e, 0x2c,
	0xc0, 0x87, 0xf8, 0xee, 0x5f, 0xe3, 0xbe, 0x42, 0xeb, 0x21, 0xb9, 0x92, 0x92, 0xfb, 0x84, 0x2d,
	0x69, 0xfc, 0xc2, 0x2a, 0x07, 0x25, 0xd8, 0x20, 0x8b, 0x21, 0xe4, 0x70, 0xb3, 0xc8, 0x5a, 0xe6,
	0x1f, 0xbc, 0xa3, 0x7f, 0x49, 0xdd, 0x40, 0x7f, 0x24, 0x74, 0xfa, 0x0c, 0x82, 0x3a, 0xff, 0xef,
	0x57, 0x4b, 0x3f, 0x7d, 0xfb, 0xf1, 0x79, 0x94, 0x24, 0x77, 0x36, 0x06, 0x64, 0xde, 0x37, 0x90,
	0x6a, 0x39, 0x73, 0x72, 0xf4, 0xe4, 0xf1, 0x97, 0xef, 0x07, 0xe4, 0xf5, 0xa3, 0x2d, 0xa7, 0xa2,
	0xba, 0xd0, 0xab, 0xc9, 0x38, 0xdd, 0x69, 0x1b, 0xfd, 0xe1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0x27, 0xc6, 0x6b, 0xb4, 0x03, 0x00, 0x00,
}
