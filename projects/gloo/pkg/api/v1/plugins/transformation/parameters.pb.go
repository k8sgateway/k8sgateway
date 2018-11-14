// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transformation/parameters.proto

package transformation // import "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/transformation"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

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

type Parameters struct {
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
	Path                 *types.StringValue `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Parameters) Reset()         { *m = Parameters{} }
func (m *Parameters) String() string { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()    {}
func (*Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_parameters_c161b264e05ac516, []int{0}
}
func (m *Parameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameters.Unmarshal(m, b)
}
func (m *Parameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameters.Marshal(b, m, deterministic)
}
func (dst *Parameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameters.Merge(dst, src)
}
func (m *Parameters) XXX_Size() int {
	return xxx_messageInfo_Parameters.Size(m)
}
func (m *Parameters) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameters.DiscardUnknown(m)
}

var xxx_messageInfo_Parameters proto.InternalMessageInfo

func (m *Parameters) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Parameters) GetPath() *types.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "transformation.plugins.gloo.solo.io.Parameters")
	proto.RegisterMapType((map[string]string)(nil), "transformation.plugins.gloo.solo.io.Parameters.HeadersEntry")
}
func (this *Parameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Parameters)
	if !ok {
		that2, ok := that.(Parameters)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("transformation/parameters.proto", fileDescriptor_parameters_c161b264e05ac516)
}

var fileDescriptor_parameters_c161b264e05ac516 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0xe5, 0x96, 0x3f, 0xaa, 0xcb, 0x80, 0xa2, 0x0e, 0x55, 0x85, 0x4a, 0x04, 0x4b, 0x16,
	0xee, 0xa0, 0x2c, 0xa8, 0x62, 0x42, 0x42, 0x62, 0xac, 0x82, 0x94, 0x81, 0xcd, 0x29, 0xae, 0x63,
	0x9a, 0xe4, 0x2c, 0xdb, 0x29, 0xea, 0x1b, 0xf1, 0x3c, 0x3c, 0x02, 0x4f, 0x82, 0x92, 0xd0, 0x40,
	0x37, 0x26, 0x7f, 0x3e, 0xdd, 0xf7, 0xbb, 0xef, 0xe3, 0xe7, 0xde, 0x8a, 0xd2, 0xad, 0xc8, 0x16,
	0xc2, 0x6b, 0x2a, 0xd1, 0x08, 0x2b, 0x0a, 0xe9, 0xa5, 0x75, 0x60, 0x2c, 0x79, 0x0a, 0x2e, 0xf7,
	0x17, 0xc0, 0xe4, 0x95, 0xd2, 0xa5, 0x03, 0x95, 0x13, 0x81, 0xa3, 0x9c, 0x40, 0xd3, 0x64, 0xaa,
	0x88, 0x54, 0x2e, 0xb1, 0xb1, 0xa4, 0xd5, 0x0a, 0xdf, 0xad, 0x30, 0xa6, 0x83, 0x4c, 0x46, 0x8a,
	0x14, 0x35, 0x12, 0x6b, 0xd5, 0x4e, 0x2f, 0x3e, 0x19, 0xe7, 0x8b, 0xee, 0x5e, 0x90, 0xf0, 0xe3,
	0x4c, 0x8a, 0x57, 0x69, 0xdd, 0x98, 0x85, 0xfd, 0x68, 0x38, 0xbb, 0x87, 0x7f, 0xdc, 0x86, 0x5f,
	0x02, 0x3c, 0xb5, 0xf6, 0xc7, 0xd2, 0xdb, 0x6d, 0xbc, 0x83, 0x05, 0xd7, 0xfc, 0xc0, 0x08, 0x9f,
	0x8d, 0x7b, 0x21, 0x8b, 0x86, 0xb3, 0x33, 0x68, 0xb3, 0xc2, 0x2e, 0x2b, 0x3c, 0x7b, 0xab, 0x4b,
	0x95, 0x88, 0xbc, 0x92, 0x71, 0xb3, 0x39, 0x99, 0xf3, 0x93, 0xbf, 0xa8, 0xe0, 0x94, 0xf7, 0xd7,
	0x72, 0x3b, 0x66, 0x21, 0x8b, 0x06, 0x71, 0x2d, 0x83, 0x11, 0x3f, 0xdc, 0xd4, 0x86, 0x06, 0x3a,
	0x88, 0xdb, 0xcf, 0xbc, 0x77, 0xc7, 0x1e, 0x92, 0x8f, 0xaf, 0x29, 0x7b, 0x59, 0x28, 0xed, 0xb3,
	0x2a, 0x85, 0x25, 0x15, 0x58, 0x07, 0xbd, 0xd2, 0xd4, 0xbe, 0xc6, 0xd2, 0x9b, 0x5c, 0x7a, 0x87,
	0x9d, 0xa8, 0xcb, 0xa0, 0x59, 0x2b, 0x14, 0x46, 0xe3, 0xe6, 0x06, 0x7f, 0x4a, 0xe2, 0x7e, 0xf7,
	0xf4, 0xa8, 0xc9, 0x7b, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x51, 0xf1, 0x10, 0xf2, 0xb1, 0x01,
	0x00, 0x00,
}
