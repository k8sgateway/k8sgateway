// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/cors/cors.proto

package cors

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

// CorsPolicy defines Cross-Origin Resource Sharing for a virtual service.
type CorsPolicy struct {
	// Specifies the origins that will be allowed to make CORS requests.
	//
	// An origin is allowed if either allow_origin or allow_origin_regex match.
	AllowOrigin []string `protobuf:"bytes,1,rep,name=allow_origin,json=allowOrigin,proto3" json:"allow_origin,omitempty"`
	// Specifies regex patterns that match origins that will be allowed to make
	// CORS requests.
	//
	// An origin is allowed if either allow_origin or allow_origin_regex match.
	AllowOriginRegex []string `protobuf:"bytes,2,rep,name=allow_origin_regex,json=allowOriginRegex,proto3" json:"allow_origin_regex,omitempty"`
	// Specifies the content for the *access-control-allow-methods* header.
	AllowMethods []string `protobuf:"bytes,3,rep,name=allow_methods,json=allowMethods,proto3" json:"allow_methods,omitempty"`
	// Specifies the content for the *access-control-allow-headers* header.
	AllowHeaders []string `protobuf:"bytes,4,rep,name=allow_headers,json=allowHeaders,proto3" json:"allow_headers,omitempty"`
	// Specifies the content for the *access-control-expose-headers* header.
	ExposeHeaders []string `protobuf:"bytes,5,rep,name=expose_headers,json=exposeHeaders,proto3" json:"expose_headers,omitempty"`
	// Specifies the content for the *access-control-max-age* header.
	MaxAge string `protobuf:"bytes,6,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// Specifies whether the resource allows credentials.
	AllowCredentials bool `protobuf:"varint,7,opt,name=allow_credentials,json=allowCredentials,proto3" json:"allow_credentials,omitempty"`
	// Optional, only applies to route-specific CORS Policies, defaults to false.
	// If set, the CORS Policy (specified on the virtual host) will be disabled for this route.
	DisableForRoute      bool     `protobuf:"varint,8,opt,name=disable_for_route,json=disableForRoute,proto3" json:"disable_for_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CorsPolicy) Reset()         { *m = CorsPolicy{} }
func (m *CorsPolicy) String() string { return proto.CompactTextString(m) }
func (*CorsPolicy) ProtoMessage()    {}
func (*CorsPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_742a11da6057c3fe, []int{0}
}
func (m *CorsPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorsPolicy.Unmarshal(m, b)
}
func (m *CorsPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorsPolicy.Marshal(b, m, deterministic)
}
func (m *CorsPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorsPolicy.Merge(m, src)
}
func (m *CorsPolicy) XXX_Size() int {
	return xxx_messageInfo_CorsPolicy.Size(m)
}
func (m *CorsPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CorsPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CorsPolicy proto.InternalMessageInfo

func (m *CorsPolicy) GetAllowOrigin() []string {
	if m != nil {
		return m.AllowOrigin
	}
	return nil
}

func (m *CorsPolicy) GetAllowOriginRegex() []string {
	if m != nil {
		return m.AllowOriginRegex
	}
	return nil
}

func (m *CorsPolicy) GetAllowMethods() []string {
	if m != nil {
		return m.AllowMethods
	}
	return nil
}

func (m *CorsPolicy) GetAllowHeaders() []string {
	if m != nil {
		return m.AllowHeaders
	}
	return nil
}

func (m *CorsPolicy) GetExposeHeaders() []string {
	if m != nil {
		return m.ExposeHeaders
	}
	return nil
}

func (m *CorsPolicy) GetMaxAge() string {
	if m != nil {
		return m.MaxAge
	}
	return ""
}

func (m *CorsPolicy) GetAllowCredentials() bool {
	if m != nil {
		return m.AllowCredentials
	}
	return false
}

func (m *CorsPolicy) GetDisableForRoute() bool {
	if m != nil {
		return m.DisableForRoute
	}
	return false
}

func init() {
	proto.RegisterType((*CorsPolicy)(nil), "cors.plugins.gloo.solo.io.CorsPolicy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/cors/cors.proto", fileDescriptor_742a11da6057c3fe)
}

var fileDescriptor_742a11da6057c3fe = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc7, 0x71, 0xb2, 0xe5, 0x43, 0x5b, 0xb6, 0x45, 0x0c, 0xe6, 0xe5, 0x22, 0x64, 0x1b, 0x83,
	0xb0, 0x0f, 0x8b, 0xb1, 0xeb, 0x5e, 0xb4, 0x29, 0xa5, 0x37, 0xa5, 0xc5, 0x97, 0xbd, 0x31, 0xb2,
	0x7d, 0xa2, 0xa8, 0x95, 0x7d, 0x84, 0x64, 0x37, 0xee, 0x1b, 0xf5, 0x31, 0xfa, 0x2c, 0x7d, 0x92,
	0x62, 0x29, 0x69, 0x02, 0xed, 0x45, 0x6f, 0x8c, 0xf4, 0xfb, 0xff, 0xce, 0x39, 0x86, 0x23, 0x72,
	0x2c, 0x64, 0xb5, 0xaa, 0xd3, 0x28, 0xc3, 0x82, 0x59, 0x54, 0xf8, 0x57, 0x22, 0x13, 0x0a, 0x91,
	0x69, 0x83, 0x57, 0x90, 0x55, 0xd6, 0xdf, 0xb8, 0x96, 0xec, 0xe6, 0x1f, 0xd3, 0xaa, 0x16, 0xb2,
	0xb4, 0x2c, 0x43, 0xe3, 0x3f, 0x91, 0x36, 0x58, 0x21, 0xfd, 0xea, 0xcf, 0x3e, 0x8d, 0xda, 0x8a,
	0xa8, 0x6d, 0x16, 0x49, 0x9c, 0x7c, 0x16, 0x28, 0xd0, 0x59, 0xac, 0x3d, 0xf9, 0x82, 0xc9, 0x54,
	0x20, 0x0a, 0x05, 0xcc, 0xdd, 0xd2, 0x7a, 0xc9, 0xd6, 0x86, 0x6b, 0x0d, 0xdb, 0x86, 0xcf, 0xf3,
	0xbc, 0x36, 0xbc, 0x92, 0x58, 0xfa, 0xfc, 0xfb, 0x7d, 0x87, 0x90, 0x05, 0x1a, 0x7b, 0x81, 0x4a,
	0x66, 0xb7, 0xf4, 0x1b, 0x79, 0xcf, 0x95, 0xc2, 0x75, 0x82, 0x46, 0x0a, 0x59, 0x86, 0xc1, 0xac,
	0x3b, 0x1f, 0xc6, 0xef, 0x1c, 0x3b, 0x77, 0x88, 0xfe, 0x21, 0x74, 0x5f, 0x49, 0x0c, 0x08, 0x68,
	0xc2, 0x8e, 0x13, 0x3f, 0xed, 0x89, 0x71, 0xcb, 0xe9, 0x0f, 0x32, 0xf2, 0x76, 0x01, 0xd5, 0x0a,
	0x73, 0x1b, 0x76, 0x9d, 0xe8, 0xa7, 0x9c, 0x79, 0xb6, 0x93, 0x56, 0xc0, 0x73, 0x30, 0x36, 0x7c,
	0xb3, 0x27, 0x9d, 0x7a, 0x46, 0x7f, 0x92, 0x0f, 0xd0, 0x68, 0xb4, 0xf0, 0x64, 0xbd, 0x75, 0xd6,
	0xc8, 0xd3, 0xad, 0xf6, 0x85, 0xf4, 0x0b, 0xde, 0x24, 0x5c, 0x40, 0xd8, 0x9b, 0x05, 0xf3, 0x61,
	0xdc, 0x2b, 0x78, 0x73, 0x28, 0x80, 0xfe, 0x26, 0x63, 0x3f, 0x24, 0x33, 0x90, 0x43, 0x59, 0x49,
	0xae, 0x6c, 0xd8, 0x9f, 0x05, 0xf3, 0xc1, 0xe6, 0xb7, 0x17, 0x3b, 0x4e, 0x7f, 0x91, 0x71, 0x2e,
	0x2d, 0x4f, 0x15, 0x24, 0x4b, 0x34, 0x89, 0xc1, 0xba, 0x82, 0x70, 0xe0, 0xe4, 0x8f, 0x9b, 0xe0,
	0x04, 0x4d, 0xdc, 0xe2, 0xa3, 0xc5, 0xdd, 0xc3, 0x34, 0xb8, 0x3c, 0x78, 0xdd, 0xfe, 0xf5, 0xb5,
	0x78, 0xe9, 0x0d, 0xa4, 0x3d, 0xb7, 0x8e, 0xff, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0xbf,
	0x15, 0xde, 0x47, 0x02, 0x00, 0x00,
}

func (this *CorsPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CorsPolicy)
	if !ok {
		that2, ok := that.(CorsPolicy)
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
	if len(this.AllowOrigin) != len(that1.AllowOrigin) {
		return false
	}
	for i := range this.AllowOrigin {
		if this.AllowOrigin[i] != that1.AllowOrigin[i] {
			return false
		}
	}
	if len(this.AllowOriginRegex) != len(that1.AllowOriginRegex) {
		return false
	}
	for i := range this.AllowOriginRegex {
		if this.AllowOriginRegex[i] != that1.AllowOriginRegex[i] {
			return false
		}
	}
	if len(this.AllowMethods) != len(that1.AllowMethods) {
		return false
	}
	for i := range this.AllowMethods {
		if this.AllowMethods[i] != that1.AllowMethods[i] {
			return false
		}
	}
	if len(this.AllowHeaders) != len(that1.AllowHeaders) {
		return false
	}
	for i := range this.AllowHeaders {
		if this.AllowHeaders[i] != that1.AllowHeaders[i] {
			return false
		}
	}
	if len(this.ExposeHeaders) != len(that1.ExposeHeaders) {
		return false
	}
	for i := range this.ExposeHeaders {
		if this.ExposeHeaders[i] != that1.ExposeHeaders[i] {
			return false
		}
	}
	if this.MaxAge != that1.MaxAge {
		return false
	}
	if this.AllowCredentials != that1.AllowCredentials {
		return false
	}
	if this.DisableForRoute != that1.DisableForRoute {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
