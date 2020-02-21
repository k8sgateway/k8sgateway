// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/core/matchers.proto

package core

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Parameters for matching routes to requests received by a Gloo-managed proxy
type RouteMatcher struct {
	// Types that are valid to be assigned to PathSpecifier:
	//	*RouteMatcher_Prefix
	//	*RouteMatcher_Exact
	//	*RouteMatcher_Regex
	PathSpecifier isRouteMatcher_PathSpecifier `protobuf_oneof:"path_specifier"`
	// Specifies a set of headers that the route should match on. The router will
	// check the request’s headers against all the specified headers in the route
	// config. A match will happen if all the headers in the route are present in
	// the request with the same values (or based on presence if the value field
	// is not in the config).
	Headers []*HeaderMatcher `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
	// Specifies a set of URL query parameters on which the route should
	// match. The router will check the query string from the *path* header
	// against all the specified query parameters. If the number of specified
	// query parameters is nonzero, they all must match the *path* header's
	// query string for a match to occur.
	QueryParameters []*QueryParameterMatcher `protobuf:"bytes,7,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	// HTTP Method/Verb(s) to match on. If none specified, the matcher will ignore the HTTP Method
	Methods              []string `protobuf:"bytes,8,rep,name=methods,proto3" json:"methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteMatcher) Reset()         { *m = RouteMatcher{} }
func (m *RouteMatcher) String() string { return proto.CompactTextString(m) }
func (*RouteMatcher) ProtoMessage()    {}
func (*RouteMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_6df1264a8dd3a4b7, []int{0}
}
func (m *RouteMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatcher.Unmarshal(m, b)
}
func (m *RouteMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatcher.Marshal(b, m, deterministic)
}
func (m *RouteMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatcher.Merge(m, src)
}
func (m *RouteMatcher) XXX_Size() int {
	return xxx_messageInfo_RouteMatcher.Size(m)
}
func (m *RouteMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatcher proto.InternalMessageInfo

type isRouteMatcher_PathSpecifier interface {
	isRouteMatcher_PathSpecifier()
	Equal(interface{}) bool
}

type RouteMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}
type RouteMatcher_Exact struct {
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof" json:"exact,omitempty"`
}
type RouteMatcher_Regex struct {
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof" json:"regex,omitempty"`
}

func (*RouteMatcher_Prefix) isRouteMatcher_PathSpecifier() {}
func (*RouteMatcher_Exact) isRouteMatcher_PathSpecifier()  {}
func (*RouteMatcher_Regex) isRouteMatcher_PathSpecifier()  {}

func (m *RouteMatcher) GetPathSpecifier() isRouteMatcher_PathSpecifier {
	if m != nil {
		return m.PathSpecifier
	}
	return nil
}

func (m *RouteMatcher) GetPrefix() string {
	if x, ok := m.GetPathSpecifier().(*RouteMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *RouteMatcher) GetExact() string {
	if x, ok := m.GetPathSpecifier().(*RouteMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *RouteMatcher) GetRegex() string {
	if x, ok := m.GetPathSpecifier().(*RouteMatcher_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *RouteMatcher) GetHeaders() []*HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RouteMatcher) GetQueryParameters() []*QueryParameterMatcher {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *RouteMatcher) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatcher_Prefix)(nil),
		(*RouteMatcher_Exact)(nil),
		(*RouteMatcher_Regex)(nil),
	}
}

//   Internally, Gloo always uses the HTTP/2 *:authority* header to represent the HTTP/1 *Host*
//   header. Thus, if attempting to match on *Host*, match on *:authority* instead.
//
//   In the absence of any header match specifier, match will default to `present_match`
//   i.e, a request that has the `name` header will match, regardless of the header's
//   value.
type HeaderMatcher struct {
	// Specifies the name of the header in the request.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the value of the header. If the value is absent a request that
	// has the name header will match, regardless of the header’s value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Specifies whether the header value should be treated as regex or not.
	Regex                bool     `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderMatcher) Reset()         { *m = HeaderMatcher{} }
func (m *HeaderMatcher) String() string { return proto.CompactTextString(m) }
func (*HeaderMatcher) ProtoMessage()    {}
func (*HeaderMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_6df1264a8dd3a4b7, []int{1}
}
func (m *HeaderMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderMatcher.Unmarshal(m, b)
}
func (m *HeaderMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderMatcher.Marshal(b, m, deterministic)
}
func (m *HeaderMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderMatcher.Merge(m, src)
}
func (m *HeaderMatcher) XXX_Size() int {
	return xxx_messageInfo_HeaderMatcher.Size(m)
}
func (m *HeaderMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderMatcher proto.InternalMessageInfo

func (m *HeaderMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HeaderMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *HeaderMatcher) GetRegex() bool {
	if m != nil {
		return m.Regex
	}
	return false
}

// Query parameter matching treats the query string of a request's :path header
// as an ampersand-separated list of keys and/or key=value elements.
type QueryParameterMatcher struct {
	// Specifies the name of a key that must be present in the requested
	// *path*'s query string.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the value of the key. If the value is absent, a request
	// that contains the key in its query string will match, whether the
	// key appears with a value (e.g., "?debug=true") or not (e.g., "?debug")
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Specifies whether the query parameter value is a regular expression.
	// Defaults to false. The entire query parameter value (i.e., the part to
	// the right of the equals sign in "key=value") must match the regex.
	// E.g., the regex "\d+$" will match "123" but not "a123" or "123a".
	Regex                bool     `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryParameterMatcher) Reset()         { *m = QueryParameterMatcher{} }
func (m *QueryParameterMatcher) String() string { return proto.CompactTextString(m) }
func (*QueryParameterMatcher) ProtoMessage()    {}
func (*QueryParameterMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_6df1264a8dd3a4b7, []int{2}
}
func (m *QueryParameterMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryParameterMatcher.Unmarshal(m, b)
}
func (m *QueryParameterMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryParameterMatcher.Marshal(b, m, deterministic)
}
func (m *QueryParameterMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParameterMatcher.Merge(m, src)
}
func (m *QueryParameterMatcher) XXX_Size() int {
	return xxx_messageInfo_QueryParameterMatcher.Size(m)
}
func (m *QueryParameterMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParameterMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParameterMatcher proto.InternalMessageInfo

func (m *QueryParameterMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryParameterMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *QueryParameterMatcher) GetRegex() bool {
	if m != nil {
		return m.Regex
	}
	return false
}

func init() {
	proto.RegisterType((*RouteMatcher)(nil), "core.gloo.solo.io.RouteMatcher")
	proto.RegisterType((*HeaderMatcher)(nil), "core.gloo.solo.io.HeaderMatcher")
	proto.RegisterType((*QueryParameterMatcher)(nil), "core.gloo.solo.io.QueryParameterMatcher")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/core/matchers.proto", fileDescriptor_6df1264a8dd3a4b7)
}

var fileDescriptor_6df1264a8dd3a4b7 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0x6f, 0xdb, 0xdb, 0x7f, 0xbe, 0x17, 0x28, 0x56, 0x41, 0x11, 0x03, 0x8a, 0x3a, 0x65,
	0x21, 0x16, 0x20, 0x16, 0x58, 0x50, 0xa7, 0x2e, 0x08, 0x08, 0x03, 0x12, 0x4b, 0xe5, 0xa6, 0xa7,
	0x89, 0xa1, 0xe1, 0xb8, 0xb6, 0x53, 0x95, 0x07, 0xe1, 0x1d, 0x78, 0x2e, 0x9e, 0x04, 0x39, 0x4e,
	0x04, 0x15, 0x1d, 0x90, 0xd8, 0xfc, 0x7d, 0xe7, 0xcb, 0x2f, 0x47, 0x9f, 0x0e, 0xb9, 0x4c, 0x84,
	0x49, 0xf3, 0x49, 0x18, 0x63, 0xc6, 0x34, 0xce, 0xf1, 0x48, 0x20, 0x4b, 0xe6, 0x88, 0x4c, 0x2a,
	0x7c, 0x84, 0xd8, 0x68, 0xa7, 0xb8, 0x14, 0x6c, 0x79, 0xcc, 0x62, 0x54, 0xc0, 0x32, 0x6e, 0xe2,
	0x14, 0x94, 0x0e, 0xa5, 0x42, 0x83, 0x74, 0xd7, 0x9a, 0xa1, 0x4d, 0x85, 0x16, 0x10, 0x0a, 0x3c,
	0xe8, 0x27, 0x98, 0x60, 0x31, 0x65, 0xf6, 0xe5, 0x82, 0x83, 0xd7, 0x3a, 0xf9, 0x1f, 0x61, 0x6e,
	0xe0, 0xca, 0x01, 0xa8, 0x47, 0x5a, 0x52, 0xc1, 0x4c, 0xac, 0xbc, 0x9a, 0x5f, 0x0b, 0xba, 0xa3,
	0x3f, 0x51, 0xa9, 0xe9, 0x3e, 0x69, 0xc2, 0x8a, 0xc7, 0xc6, 0xab, 0x97, 0x03, 0x27, 0xad, 0xaf,
	0x20, 0x81, 0x95, 0xd7, 0xa8, 0xfc, 0x42, 0xd2, 0x73, 0xd2, 0x4e, 0x81, 0x4f, 0x41, 0x69, 0xaf,
	0xe5, 0x37, 0x82, 0x7f, 0x27, 0x7e, 0xf8, 0x6d, 0xab, 0x70, 0x54, 0x24, 0xca, 0x9f, 0x47, 0xd5,
	0x07, 0xf4, 0x8e, 0xf4, 0x16, 0x39, 0xa8, 0x97, 0xb1, 0xe4, 0x8a, 0x67, 0x60, 0x2c, 0xa4, 0x5d,
	0x40, 0x82, 0x0d, 0x90, 0x5b, 0x1b, 0xbd, 0xa9, 0x92, 0x15, 0x6c, 0x67, 0xb1, 0x66, 0x6b, 0xea,
	0x91, 0x76, 0x06, 0x26, 0xc5, 0xa9, 0xf6, 0x3a, 0x7e, 0x23, 0xe8, 0x46, 0x95, 0x1c, 0xf6, 0xc8,
	0xb6, 0xe4, 0x26, 0x1d, 0x6b, 0x09, 0xb1, 0x98, 0x09, 0x50, 0x83, 0x6b, 0xb2, 0xb5, 0xb6, 0x1a,
	0xa5, 0xe4, 0xef, 0x33, 0xcf, 0xc0, 0xb5, 0x12, 0x15, 0x6f, 0xda, 0x27, 0xcd, 0x25, 0x9f, 0xe7,
	0xe0, 0x1a, 0x89, 0x9c, 0xb0, 0xee, 0x67, 0x1f, 0x9d, 0xb2, 0x8d, 0xc1, 0x3d, 0xd9, 0xdb, 0xb8,
	0xe6, 0x6f, 0xc1, 0xc3, 0x8b, 0xb7, 0xf7, 0xc3, 0xda, 0xc3, 0xd9, 0xcf, 0x4e, 0x46, 0x3e, 0x25,
	0x5f, 0xcf, 0x66, 0xd2, 0x2a, 0xae, 0xe0, 0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0x11, 0xa2, 0x29,
	0xed, 0x72, 0x02, 0x00, 0x00,
}

func (this *RouteMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteMatcher)
	if !ok {
		that2, ok := that.(RouteMatcher)
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
	if that1.PathSpecifier == nil {
		if this.PathSpecifier != nil {
			return false
		}
	} else if this.PathSpecifier == nil {
		return false
	} else if !this.PathSpecifier.Equal(that1.PathSpecifier) {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if len(this.QueryParameters) != len(that1.QueryParameters) {
		return false
	}
	for i := range this.QueryParameters {
		if !this.QueryParameters[i].Equal(that1.QueryParameters[i]) {
			return false
		}
	}
	if len(this.Methods) != len(that1.Methods) {
		return false
	}
	for i := range this.Methods {
		if this.Methods[i] != that1.Methods[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteMatcher_Prefix) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteMatcher_Prefix)
	if !ok {
		that2, ok := that.(RouteMatcher_Prefix)
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
	if this.Prefix != that1.Prefix {
		return false
	}
	return true
}
func (this *RouteMatcher_Exact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteMatcher_Exact)
	if !ok {
		that2, ok := that.(RouteMatcher_Exact)
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
	if this.Exact != that1.Exact {
		return false
	}
	return true
}
func (this *RouteMatcher_Regex) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteMatcher_Regex)
	if !ok {
		that2, ok := that.(RouteMatcher_Regex)
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
	if this.Regex != that1.Regex {
		return false
	}
	return true
}
func (this *HeaderMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderMatcher)
	if !ok {
		that2, ok := that.(HeaderMatcher)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *QueryParameterMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryParameterMatcher)
	if !ok {
		that2, ok := that.(QueryParameterMatcher)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
