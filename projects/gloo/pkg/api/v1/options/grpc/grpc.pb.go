// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc/grpc.proto

package grpc

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Service spec describing GRPC upstreams. This will usually be filled
// automatically via function discovery (if the upstream supports reflection).
// If your upstream service is a GRPC service, use this service spec (an empty
// spec is fine), to make sure that traffic to it is routed with http2.
type ServiceSpec struct {
	// Descriptors that contain information of the services listed below.
	// this is a serialized google.protobuf.FileDescriptorSet
	Descriptors []byte `protobuf:"bytes,1,opt,name=descriptors,proto3" json:"descriptors,omitempty"`
	// List of services used by this upstream. For a grpc upstream where you don't
	// need to use Gloo's function routing, this can be an empty list. These
	// services must be present in the descriptors.
	GrpcServices         []*ServiceSpec_GrpcService `protobuf:"bytes,2,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bddd1d7957d358a, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(m, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

func (m *ServiceSpec) GetDescriptors() []byte {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *ServiceSpec) GetGrpcServices() []*ServiceSpec_GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

// Describes a grpc service
type ServiceSpec_GrpcService struct {
	// The package of this service.
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// The service name of this service.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The functions available in this service.
	FunctionNames        []string `protobuf:"bytes,3,rep,name=function_names,json=functionNames,proto3" json:"function_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceSpec_GrpcService) Reset()         { *m = ServiceSpec_GrpcService{} }
func (m *ServiceSpec_GrpcService) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec_GrpcService) ProtoMessage()    {}
func (*ServiceSpec_GrpcService) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bddd1d7957d358a, []int{0, 0}
}
func (m *ServiceSpec_GrpcService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec_GrpcService.Unmarshal(m, b)
}
func (m *ServiceSpec_GrpcService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec_GrpcService.Marshal(b, m, deterministic)
}
func (m *ServiceSpec_GrpcService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec_GrpcService.Merge(m, src)
}
func (m *ServiceSpec_GrpcService) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec_GrpcService.Size(m)
}
func (m *ServiceSpec_GrpcService) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec_GrpcService.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec_GrpcService proto.InternalMessageInfo

func (m *ServiceSpec_GrpcService) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ServiceSpec_GrpcService) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceSpec_GrpcService) GetFunctionNames() []string {
	if m != nil {
		return m.FunctionNames
	}
	return nil
}

// This is only for upstream with Grpc service spec.
type DestinationSpec struct {
	// The proto package of the function.
	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	// The name of the service of the function.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// The name of the function.
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	// Parameters describe how to extract the function parameters from the
	// request.
	Parameters           *transformation.Parameters `protobuf:"bytes,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bddd1d7957d358a, []int{1}
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

func (m *DestinationSpec) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *DestinationSpec) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *DestinationSpec) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *DestinationSpec) GetParameters() *transformation.Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "grpc.options.gloo.solo.io.ServiceSpec")
	proto.RegisterType((*ServiceSpec_GrpcService)(nil), "grpc.options.gloo.solo.io.ServiceSpec.GrpcService")
	proto.RegisterType((*DestinationSpec)(nil), "grpc.options.gloo.solo.io.DestinationSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/grpc/grpc.proto", fileDescriptor_3bddd1d7957d358a)
}

var fileDescriptor_3bddd1d7957d358a = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0xc5, 0x71, 0x69, 0x1b, 0x39, 0x69, 0xc1, 0xf4, 0xe0, 0xfa, 0x50, 0xdc, 0x40, 0xc1, 0x97,
	0x4a, 0x34, 0x3d, 0xf7, 0x12, 0x42, 0x7b, 0x6b, 0x8b, 0x53, 0x28, 0xf4, 0x12, 0x14, 0x55, 0x51,
	0xd5, 0xc4, 0x1e, 0x21, 0x29, 0x21, 0xec, 0x79, 0x3f, 0x66, 0xcf, 0xfb, 0x49, 0xfb, 0x15, 0x7b,
	0x5c, 0x24, 0xcb, 0x89, 0x17, 0xb2, 0xb0, 0xec, 0xc5, 0xe8, 0xbd, 0x79, 0x33, 0x6f, 0x66, 0x3c,
	0x68, 0x2e, 0xa4, 0xfd, 0xb7, 0x5b, 0x61, 0x06, 0x35, 0x31, 0xb0, 0x85, 0x8f, 0x12, 0x88, 0xd8,
	0x02, 0x10, 0xa5, 0xe1, 0x3f, 0x67, 0xd6, 0xb4, 0x88, 0x2a, 0x49, 0xf6, 0x9f, 0x08, 0x28, 0x2b,
	0xa1, 0x31, 0x44, 0x68, 0xc5, 0xfc, 0x07, 0x2b, 0x0d, 0x16, 0xd2, 0xb7, 0xfe, 0x1d, 0xa2, 0xd8,
	0x65, 0x60, 0x57, 0x0c, 0x4b, 0xc8, 0xdf, 0x08, 0x10, 0xe0, 0x55, 0xc4, 0xbd, 0xda, 0x84, 0x3c,
	0xe5, 0x07, 0xdb, 0x92, 0xfc, 0x60, 0x03, 0xf7, 0xeb, 0x49, 0xad, 0x58, 0x4d, 0x1b, 0xb3, 0x06,
	0x5d, 0x53, 0x87, 0x89, 0xa2, 0x9a, 0xd6, 0xdc, 0x72, 0x6d, 0xda, 0xaa, 0x93, 0xcb, 0x01, 0x4a,
	0x16, 0x5c, 0xef, 0x25, 0xe3, 0x0b, 0xc5, 0x59, 0x5a, 0xa0, 0xe4, 0x2f, 0x37, 0x4c, 0x4b, 0x65,
	0x41, 0x9b, 0x2c, 0x2a, 0xa2, 0x72, 0x54, 0xf5, 0xa9, 0xf4, 0x37, 0x1a, 0xbb, 0x71, 0x96, 0xa6,
	0xcd, 0x32, 0xd9, 0xa0, 0x88, 0xcb, 0x64, 0x3a, 0xc5, 0x0f, 0x0e, 0x89, 0x7b, 0x06, 0xf8, 0x9b,
	0x56, 0x2c, 0xe0, 0x6a, 0x24, 0x4e, 0xc0, 0xe4, 0x17, 0x28, 0xe9, 0x05, 0xd3, 0xf7, 0x68, 0xa4,
	0x28, 0xdb, 0x50, 0xc1, 0x97, 0x0d, 0xad, 0xb9, 0x6f, 0x65, 0x58, 0x25, 0x81, 0xfb, 0x4e, 0x6b,
	0x2f, 0x09, 0x5d, 0xb4, 0x92, 0x41, 0x2b, 0x09, 0x9c, 0x97, 0x7c, 0x40, 0xaf, 0xd6, 0xbb, 0x86,
	0xb9, 0xa6, 0xbc, 0xc6, 0x64, 0x71, 0x11, 0x97, 0xc3, 0x6a, 0xdc, 0xb1, 0x4e, 0x65, 0x26, 0xd7,
	0x11, 0x7a, 0x3d, 0xe7, 0xc6, 0xca, 0xc6, 0xef, 0xc9, 0xaf, 0x22, 0x43, 0x2f, 0x82, 0x59, 0xf0,
	0xee, 0xa0, 0x8b, 0x04, 0x8f, 0x60, 0xd9, 0xc1, 0x34, 0x47, 0x2f, 0xbb, 0xc2, 0x59, 0xec, 0x43,
	0x47, 0x9c, 0xfe, 0x40, 0xe8, 0xb4, 0xfe, 0xec, 0x59, 0x11, 0x95, 0xc9, 0x94, 0xe0, 0xfb, 0x3f,
	0xe8, 0xfc, 0xfe, 0x7e, 0x1e, 0xd3, 0xaa, 0x5e, 0x89, 0xd9, 0xd7, 0xdb, 0x59, 0x74, 0x75, 0xf3,
	0x2e, 0xfa, 0xf3, 0xe5, 0x71, 0xa7, 0xa1, 0x36, 0xe2, 0xdc, 0xa5, 0xae, 0x9e, 0xfb, 0x53, 0xf8,
	0x7c, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x6b, 0x4b, 0x8c, 0xed, 0x02, 0x00, 0x00,
}

func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if !bytes.Equal(this.Descriptors, that1.Descriptors) {
		return false
	}
	if len(this.GrpcServices) != len(that1.GrpcServices) {
		return false
	}
	for i := range this.GrpcServices {
		if !this.GrpcServices[i].Equal(that1.GrpcServices[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_GrpcService)
	if !ok {
		that2, ok := that.(ServiceSpec_GrpcService)
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
	if this.PackageName != that1.PackageName {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if len(this.FunctionNames) != len(that1.FunctionNames) {
		return false
	}
	for i := range this.FunctionNames {
		if this.FunctionNames[i] != that1.FunctionNames[i] {
			return false
		}
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
	if this.Package != that1.Package {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	if this.Function != that1.Function {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
