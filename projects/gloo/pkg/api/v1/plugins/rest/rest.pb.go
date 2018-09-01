// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rest/rest.proto

/*
Package rest is a generated protocol buffer package.

It is generated from these files:
	rest/rest.proto

It has these top-level messages:
	ServiceSpec
	DestinationSpec
*/
package rest

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/gogoproto"
import envoy_api_v2_filter_http "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/transformation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServiceSpec struct {
	Transformations map[string]*envoy_api_v2_filter_http.TransformationTemplate `protobuf:"bytes,1,rep,name=transformations" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	SwaggerInfo     *ServiceSpec_SwaggerInfo                                    `protobuf:"bytes,2,opt,name=swagger_info,json=swaggerInfo" json:"swagger_info,omitempty"`
}

func (m *ServiceSpec) Reset()                    { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string            { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()               {}
func (*ServiceSpec) Descriptor() ([]byte, []int) { return fileDescriptorRest, []int{0} }

func (m *ServiceSpec) GetTransformations() map[string]*envoy_api_v2_filter_http.TransformationTemplate {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *ServiceSpec) GetSwaggerInfo() *ServiceSpec_SwaggerInfo {
	if m != nil {
		return m.SwaggerInfo
	}
	return nil
}

type ServiceSpec_SwaggerInfo struct {
	// Types that are valid to be assigned to SwaggerSpec:
	//	*ServiceSpec_SwaggerInfo_Url
	//	*ServiceSpec_SwaggerInfo_Inline
	SwaggerSpec isServiceSpec_SwaggerInfo_SwaggerSpec `protobuf_oneof:"swagger_spec"`
}

func (m *ServiceSpec_SwaggerInfo) Reset()                    { *m = ServiceSpec_SwaggerInfo{} }
func (m *ServiceSpec_SwaggerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServiceSpec_SwaggerInfo) ProtoMessage()               {}
func (*ServiceSpec_SwaggerInfo) Descriptor() ([]byte, []int) { return fileDescriptorRest, []int{0, 1} }

type isServiceSpec_SwaggerInfo_SwaggerSpec interface {
	isServiceSpec_SwaggerInfo_SwaggerSpec()
	Equal(interface{}) bool
}

type ServiceSpec_SwaggerInfo_Url struct {
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}
type ServiceSpec_SwaggerInfo_Inline struct {
	Inline string `protobuf:"bytes,2,opt,name=inline,proto3,oneof"`
}

func (*ServiceSpec_SwaggerInfo_Url) isServiceSpec_SwaggerInfo_SwaggerSpec()    {}
func (*ServiceSpec_SwaggerInfo_Inline) isServiceSpec_SwaggerInfo_SwaggerSpec() {}

func (m *ServiceSpec_SwaggerInfo) GetSwaggerSpec() isServiceSpec_SwaggerInfo_SwaggerSpec {
	if m != nil {
		return m.SwaggerSpec
	}
	return nil
}

func (m *ServiceSpec_SwaggerInfo) GetUrl() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Url); ok {
		return x.Url
	}
	return ""
}

func (m *ServiceSpec_SwaggerInfo) GetInline() string {
	if x, ok := m.GetSwaggerSpec().(*ServiceSpec_SwaggerInfo_Inline); ok {
		return x.Inline
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServiceSpec_SwaggerInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServiceSpec_SwaggerInfo_OneofMarshaler, _ServiceSpec_SwaggerInfo_OneofUnmarshaler, _ServiceSpec_SwaggerInfo_OneofSizer, []interface{}{
		(*ServiceSpec_SwaggerInfo_Url)(nil),
		(*ServiceSpec_SwaggerInfo_Inline)(nil),
	}
}

func _ServiceSpec_SwaggerInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServiceSpec_SwaggerInfo)
	// swagger_spec
	switch x := m.SwaggerSpec.(type) {
	case *ServiceSpec_SwaggerInfo_Url:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Url)
	case *ServiceSpec_SwaggerInfo_Inline:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Inline)
	case nil:
	default:
		return fmt.Errorf("ServiceSpec_SwaggerInfo.SwaggerSpec has unexpected type %T", x)
	}
	return nil
}

func _ServiceSpec_SwaggerInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServiceSpec_SwaggerInfo)
	switch tag {
	case 1: // swagger_spec.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.SwaggerSpec = &ServiceSpec_SwaggerInfo_Url{x}
		return true, err
	case 2: // swagger_spec.inline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.SwaggerSpec = &ServiceSpec_SwaggerInfo_Inline{x}
		return true, err
	default:
		return false, nil
	}
}

func _ServiceSpec_SwaggerInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServiceSpec_SwaggerInfo)
	// swagger_spec
	switch x := m.SwaggerSpec.(type) {
	case *ServiceSpec_SwaggerInfo_Url:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case *ServiceSpec_SwaggerInfo_Inline:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Inline)))
		n += len(x.Inline)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// This is only for upstream with REST service spec
type DestinationSpec struct {
	FunctionName           string                                           `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	Parameters             *DestinationSpec_Parameters                      `protobuf:"bytes,2,opt,name=parameters" json:"parameters,omitempty"`
	ResponseTransformation *envoy_api_v2_filter_http.TransformationTemplate `protobuf:"bytes,3,opt,name=response_transformation,json=responseTransformation" json:"response_transformation,omitempty"`
}

func (m *DestinationSpec) Reset()                    { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string            { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()               {}
func (*DestinationSpec) Descriptor() ([]byte, []int) { return fileDescriptorRest, []int{1} }

func (m *DestinationSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *DestinationSpec) GetParameters() *DestinationSpec_Parameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *DestinationSpec) GetResponseTransformation() *envoy_api_v2_filter_http.TransformationTemplate {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type DestinationSpec_Parameters struct {
	// headers that will be used to extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         headers:
	//           x-user-id: { userId }
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// part of the (or the entire) path that will be used extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         path: /users/{ userId }
	Path *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *DestinationSpec_Parameters) Reset()         { *m = DestinationSpec_Parameters{} }
func (m *DestinationSpec_Parameters) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec_Parameters) ProtoMessage()    {}
func (*DestinationSpec_Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptorRest, []int{1, 0}
}

func (m *DestinationSpec_Parameters) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DestinationSpec_Parameters) GetPath() *google_protobuf.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "rest.plugins.gloo.solo.io.ServiceSpec")
	proto.RegisterType((*ServiceSpec_SwaggerInfo)(nil), "rest.plugins.gloo.solo.io.ServiceSpec.SwaggerInfo")
	proto.RegisterType((*DestinationSpec)(nil), "rest.plugins.gloo.solo.io.DestinationSpec")
	proto.RegisterType((*DestinationSpec_Parameters)(nil), "rest.plugins.gloo.solo.io.DestinationSpec.Parameters")
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
	if len(this.Transformations) != len(that1.Transformations) {
		return false
	}
	for i := range this.Transformations {
		if !this.Transformations[i].Equal(that1.Transformations[i]) {
			return false
		}
	}
	if !this.SwaggerInfo.Equal(that1.SwaggerInfo) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo)
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
	if that1.SwaggerSpec == nil {
		if this.SwaggerSpec != nil {
			return false
		}
	} else if this.SwaggerSpec == nil {
		return false
	} else if !this.SwaggerSpec.Equal(that1.SwaggerSpec) {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Url) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Url)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Url)
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
	if this.Url != that1.Url {
		return false
	}
	return true
}
func (this *ServiceSpec_SwaggerInfo_Inline) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_SwaggerInfo_Inline)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo_Inline)
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
	if this.Inline != that1.Inline {
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
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if !this.Parameters.Equal(that1.Parameters) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	return true
}
func (this *DestinationSpec_Parameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Parameters)
	if !ok {
		that2, ok := that.(DestinationSpec_Parameters)
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
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !this.Path.Equal(that1.Path) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("rest/rest.proto", fileDescriptorRest) }

var fileDescriptorRest = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x0b, 0x0c, 0xcd, 0x1d, 0x14, 0x59, 0x15, 0x94, 0x08, 0x4d, 0xd5, 0xb8, 0xe9, 0x0d,
	0xce, 0x28, 0x42, 0x42, 0xe3, 0x8a, 0x8a, 0x9f, 0x21, 0x24, 0x84, 0xd2, 0x8d, 0x0b, 0x34, 0xa9,
	0x72, 0xc3, 0x89, 0x6b, 0x9a, 0xd8, 0x96, 0xed, 0x74, 0xea, 0x4b, 0xf0, 0x1c, 0xbc, 0x0f, 0x37,
	0x5c, 0xf3, 0x24, 0xc8, 0x71, 0x4a, 0xd3, 0x68, 0xa0, 0x69, 0x37, 0xed, 0x39, 0xce, 0x39, 0xdf,
	0x77, 0xbe, 0xf3, 0x83, 0xba, 0x1a, 0x8c, 0x8d, 0xdd, 0x0f, 0x51, 0x5a, 0x5a, 0x89, 0x1f, 0x7a,
	0x3b, 0x2f, 0x19, 0x17, 0x86, 0xb0, 0x5c, 0x4a, 0x62, 0x64, 0x2e, 0x09, 0x97, 0xd1, 0x01, 0x93,
	0x92, 0xe5, 0x10, 0x57, 0x81, 0xb3, 0x32, 0x8b, 0x2f, 0x34, 0x55, 0x0a, 0xb4, 0xf1, 0xa9, 0x51,
	0x8f, 0x49, 0x26, 0x2b, 0x33, 0x76, 0x56, 0xfd, 0x7a, 0xce, 0xb8, 0x9d, 0x97, 0x33, 0x92, 0xca,
	0x22, 0x76, 0x48, 0x4f, 0xb8, 0xf4, 0xff, 0x0b, 0x6e, 0x1d, 0xd4, 0x37, 0x48, 0xad, 0x89, 0x1d,
	0x51, 0x4c, 0x15, 0x8f, 0x97, 0x4f, 0xe3, 0x9a, 0x3c, 0xb6, 0x9a, 0x0a, 0x93, 0x49, 0x5d, 0x50,
	0xcb, 0xa5, 0x68, 0xb9, 0x1e, 0xfd, 0xf0, 0x7b, 0x88, 0x3a, 0x13, 0xd0, 0x4b, 0x9e, 0xc2, 0x44,
	0x41, 0x8a, 0x01, 0x75, 0xb7, 0xe3, 0x4c, 0x3f, 0x18, 0x84, 0xc3, 0xce, 0xe8, 0x25, 0xf9, 0xa7,
	0x30, 0xd2, 0x00, 0x20, 0xa7, 0xdb, 0xd9, 0x6f, 0x84, 0xd5, 0xab, 0xa4, 0x8d, 0x89, 0xcf, 0xd0,
	0xbe, 0xb9, 0xa0, 0x8c, 0x81, 0x9e, 0x72, 0x91, 0xc9, 0xfe, 0xce, 0x20, 0x18, 0x76, 0x46, 0xa3,
	0x2b, 0x72, 0x4c, 0x7c, 0xea, 0x7b, 0x91, 0xc9, 0xa4, 0x63, 0x36, 0x4e, 0x64, 0x51, 0xef, 0x32,
	0x7e, 0x7c, 0x0f, 0x85, 0x0b, 0x58, 0xf5, 0x83, 0x41, 0x30, 0xdc, 0x4b, 0x9c, 0x89, 0xdf, 0xa2,
	0x5b, 0x4b, 0x9a, 0x97, 0x50, 0x33, 0x1f, 0x11, 0x10, 0x4b, 0xb9, 0x22, 0x54, 0x71, 0xb2, 0x1c,
	0x91, 0x8c, 0xe7, 0x16, 0x34, 0x99, 0x5b, 0xab, 0x5a, 0x82, 0x4e, 0xa1, 0x50, 0x39, 0xb5, 0x90,
	0xf8, 0xf4, 0xe3, 0x9d, 0x17, 0x41, 0xf4, 0x01, 0x75, 0x1a, 0x15, 0x61, 0x8c, 0xc2, 0x52, 0xe7,
	0x9e, 0xec, 0xe4, 0x46, 0xe2, 0x1c, 0xdc, 0x47, 0xbb, 0x5c, 0xe4, 0x5c, 0x78, 0x3e, 0xf7, 0x5c,
	0xfb, 0xe3, 0xbb, 0x9b, 0x4e, 0x18, 0x05, 0xe9, 0xe1, 0xcf, 0x10, 0x75, 0x5f, 0x83, 0xb1, 0x5c,
	0x54, 0x7c, 0xd5, 0x50, 0x1e, 0xa3, 0x3b, 0x59, 0x29, 0x52, 0xe7, 0x4f, 0x05, 0x2d, 0xa0, 0x16,
	0xb2, 0xbf, 0x7e, 0xfc, 0x48, 0x0b, 0xc0, 0x67, 0x08, 0x29, 0xaa, 0x69, 0x01, 0x16, 0xb4, 0xa9,
	0x65, 0x3d, 0xff, 0x4f, 0x43, 0x5b, 0x24, 0xe4, 0xd3, 0xdf, 0xe4, 0xa4, 0x01, 0x84, 0x39, 0x7a,
	0xa0, 0xc1, 0x28, 0x29, 0x0c, 0x4c, 0xb7, 0xa7, 0xd8, 0x0f, 0xaf, 0xd9, 0xba, 0xfb, 0x6b, 0xc0,
	0xed, 0xef, 0xd1, 0xaf, 0x00, 0xa1, 0x4d, 0x15, 0xf8, 0x1c, 0xdd, 0x9e, 0x03, 0xfd, 0xea, 0xd4,
	0xf8, 0x15, 0x1c, 0x5f, 0x4b, 0x0d, 0x39, 0xf1, 0x20, 0x7e, 0x13, 0xd7, 0x90, 0xf8, 0x08, 0xdd,
	0x54, 0xd4, 0xce, 0xeb, 0x46, 0x3d, 0x22, 0xfe, 0x36, 0xc9, 0xfa, 0x36, 0xc9, 0xc4, 0x6a, 0x2e,
	0xd8, 0x67, 0x37, 0xe4, 0xa4, 0x8a, 0x8c, 0x8e, 0xd1, 0x7e, 0x13, 0xea, 0x92, 0xa5, 0xea, 0x35,
	0x97, 0x6a, 0xaf, 0xb1, 0x22, 0xe3, 0x77, 0x3f, 0x7e, 0x1f, 0x04, 0x5f, 0x5e, 0x5d, 0xfd, 0x94,
	0xd5, 0x82, 0xb5, 0xcf, 0xd9, 0x89, 0x9f, 0xed, 0x56, 0x05, 0x3e, 0xfb, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x60, 0x47, 0xfa, 0x78, 0x04, 0x00, 0x00,
}
