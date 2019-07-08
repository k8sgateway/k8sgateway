// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/glootest.proto

package glootest

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestRequest struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48a95305bb6f580, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

type TestResponse struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d48a95305bb6f580, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "glootest.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "glootest.TestResponse")
}

func init() { proto.RegisterFile("protos/glootest.proto", fileDescriptor_d48a95305bb6f580) }

var fileDescriptor_d48a95305bb6f580 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xcf, 0xc9, 0xcf, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0xf3, 0x85, 0x38,
	0x60, 0x7c, 0x25, 0x79, 0x2e, 0xee, 0x90, 0xd4, 0xe2, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x01, 0x2e, 0xe6, 0xe2, 0x92, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10,
	0x53, 0x49, 0x81, 0x8b, 0x07, 0xa2, 0xa0, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0x53, 0x85, 0x91,
	0x0f, 0xc4, 0x88, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x5b, 0x2e, 0x2e, 0x10, 0xd7,
	0x37, 0xb5, 0x24, 0x23, 0x3f, 0x45, 0x48, 0x54, 0x0f, 0x6e, 0x35, 0x92, 0x3d, 0x52, 0x62, 0xe8,
	0xc2, 0x10, 0xd3, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x2e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x20, 0x13, 0xa8, 0x45, 0xba, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	TestMethod(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestMethod(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/glootest.TestService/TestMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	TestMethod(context.Context, *TestRequest) (*TestResponse, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_TestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glootest.TestService/TestMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestMethod(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glootest.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMethod",
			Handler:    _TestService_TestMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/glootest.proto",
}
