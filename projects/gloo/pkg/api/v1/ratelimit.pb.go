// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/gloo/api/v1/ratelimit.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	ratelimit "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/ratelimit"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	RateLimit            *ratelimit.RateLimit `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	Constraints          []*Constraint        `protobuf:"bytes,4,rep,name=constraints,proto3" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Constraint) Reset()         { *m = Constraint{} }
func (m *Constraint) String() string { return proto.CompactTextString(m) }
func (*Constraint) ProtoMessage()    {}
func (*Constraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54dd703bc8e28ff, []int{0}
}
func (m *Constraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constraint.Unmarshal(m, b)
}
func (m *Constraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constraint.Marshal(b, m, deterministic)
}
func (m *Constraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constraint.Merge(m, src)
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
//@solo-kit:xds-service=RateLimitDiscoveryService
//@solo-kit:resource.no_references
type RateLimitConfig struct {
	// @solo-kit:resource.name
	Domain               string        `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Constraints          []*Constraint `protobuf:"bytes,2,rep,name=constraints,proto3" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RateLimitConfig) Reset()         { *m = RateLimitConfig{} }
func (m *RateLimitConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfig) ProtoMessage()    {}
func (*RateLimitConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54dd703bc8e28ff, []int{1}
}
func (m *RateLimitConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfig.Unmarshal(m, b)
}
func (m *RateLimitConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfig.Marshal(b, m, deterministic)
}
func (m *RateLimitConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfig.Merge(m, src)
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

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/gloo/api/v1/ratelimit.proto", fileDescriptor_f54dd703bc8e28ff)
}

var fileDescriptor_f54dd703bc8e28ff = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0xf1, 0x2e, 0x54, 0xaa, 0x17, 0x09, 0x64, 0x2d, 0x28, 0x44, 0xa5, 0xac, 0x72, 0x4a,
	0x2b, 0xe1, 0x40, 0x38, 0xb1, 0x17, 0x24, 0x0a, 0x08, 0x24, 0x4e, 0xe9, 0x01, 0x89, 0x03, 0xc8,
	0x75, 0x07, 0xd7, 0x34, 0xf1, 0x04, 0xdb, 0x1b, 0x69, 0x8f, 0xf0, 0x0a, 0xbc, 0x04, 0xcf, 0xc0,
	0x6b, 0xf0, 0x06, 0x88, 0xa7, 0xe0, 0x84, 0x92, 0x6c, 0xd3, 0x24, 0x08, 0x89, 0x22, 0x4e, 0x3b,
	0xf6, 0xfc, 0x33, 0xdf, 0xef, 0xcd, 0x0c, 0x7d, 0xaa, 0xb4, 0x3f, 0x59, 0x1d, 0x71, 0x89, 0x45,
	0xe2, 0x30, 0xc7, 0xbb, 0x1a, 0xdb, 0xdf, 0xd2, 0xe2, 0x7b, 0x90, 0xde, 0x25, 0x5d, 0xa0, 0x72,
	0xc4, 0x44, 0x94, 0x3a, 0xa9, 0xee, 0x27, 0x56, 0x78, 0xc8, 0x75, 0xa1, 0x3d, 0x2f, 0x2d, 0x7a,
	0x64, 0x57, 0xeb, 0x24, 0xaf, 0x0b, 0xb9, 0xc6, 0x70, 0x07, 0x4c, 0x85, 0xeb, 0x56, 0x9b, 0x26,
	0xc7, 0xda, 0x49, 0xac, 0xc0, 0xae, 0x5b, 0x6d, 0xb8, 0xa3, 0x10, 0x55, 0x0e, 0x4d, 0x5a, 0x18,
	0x83, 0x5e, 0x78, 0x8d, 0xc6, 0x6d, 0xb2, 0xaf, 0xfe, 0xcd, 0x50, 0x99, 0xaf, 0x94, 0x36, 0xee,
	0xdc, 0xd8, 0xd8, 0x62, 0x38, 0x57, 0xa8, 0xb0, 0x09, 0x93, 0x3a, 0x6a, 0x6f, 0xa3, 0xaf, 0x84,
	0xd2, 0x03, 0x34, 0xce, 0x5b, 0xa1, 0x8d, 0x67, 0xd7, 0xe9, 0xf4, 0x14, 0xd6, 0x01, 0x59, 0x90,
	0x78, 0x3b, 0xab, 0x43, 0x36, 0xa7, 0x57, 0x2a, 0x91, 0xaf, 0x20, 0x98, 0x34, 0x77, 0xed, 0x81,
	0x3d, 0xa7, 0xb4, 0xee, 0xff, 0xb6, 0x01, 0x04, 0xd3, 0x05, 0x89, 0x67, 0xe9, 0x1e, 0xef, 0x21,
	0x5b, 0x3b, 0xbc, 0xff, 0xb7, 0xf0, 0x4c, 0x78, 0x78, 0x59, 0xa7, 0xb3, 0x6d, 0x7b, 0x16, 0xb2,
	0x25, 0x9d, 0xc9, 0x8e, 0xef, 0x82, 0xcb, 0x8b, 0x69, 0x3c, 0x4b, 0x83, 0x61, 0xe1, 0xb9, 0xc1,
	0xac, 0x2f, 0x8e, 0x80, 0x5e, 0xeb, 0x7a, 0x1e, 0xa0, 0x79, 0xa7, 0x15, 0xbb, 0x49, 0xb7, 0x8e,
	0xb1, 0x10, 0xda, 0x6c, 0xde, 0xb0, 0x39, 0x8d, 0x31, 0x93, 0x0b, 0x60, 0xd2, 0x9f, 0x13, 0x7a,
	0xab, 0xe3, 0x3c, 0x39, 0xfb, 0x9a, 0x87, 0x60, 0x2b, 0x2d, 0x81, 0xbd, 0xa1, 0x37, 0x0e, 0xbd,
	0x05, 0x51, 0x8c, 0xad, 0xec, 0xf2, 0x66, 0x0c, 0xb8, 0x28, 0x35, 0xaf, 0x52, 0xde, 0x15, 0x66,
	0xf0, 0x61, 0x05, 0xce, 0x87, 0x77, 0xfe, 0x98, 0x77, 0x25, 0x1a, 0x07, 0xd1, 0xa5, 0x98, 0xdc,
	0x23, 0x6c, 0x4d, 0xc3, 0x17, 0x46, 0x5a, 0x28, 0xc0, 0x78, 0x91, 0x8f, 0x21, 0x7b, 0xc3, 0x26,
	0x3d, 0xe5, 0x6f, 0xbc, 0xfd, 0xbf, 0x91, 0x0e, 0xd0, 0x1f, 0x09, 0x9d, 0x3f, 0x03, 0x2f, 0x4f,
	0xfe, 0xfb, 0xd3, 0xe2, 0x4f, 0xdf, 0x7e, 0x7c, 0x9e, 0x44, 0xd1, 0xed, 0xc1, 0x82, 0x2c, 0xbb,
	0x01, 0x92, 0x0d, 0x67, 0x49, 0xf6, 0x1f, 0x3f, 0xfa, 0xf2, 0x7d, 0x97, 0xbc, 0x7e, 0x78, 0xc1,
	0xad, 0x28, 0x4f, 0xd5, 0x66, 0x33, 0x8e, 0xb6, 0x9a, 0x41, 0x7f, 0xf0, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x05, 0xcd, 0x2a, 0x04, 0xea, 0x03, 0x00, 0x00,
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
	Metadata: "github.com/solo-io/solo-projects/projects/gloo/api/v1/ratelimit.proto",
}
