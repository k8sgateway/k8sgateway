// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/matcher/v3/string.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/annotations"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Specifies the way to match a string.
// [#next-free-field: 7]
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	MatchPattern isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	// If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no
	// effect for the safe_regex match.
	// For example, the matcher *data* will match both input string *Data* and *data* if set to true.
	IgnoreCase           bool     `protobuf:"varint,6,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc3a85644c05785, []int{0}
}
func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMatcher.Unmarshal(m, b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return xxx_messageInfo_StringMatcher.Size(m)
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
	Equal(interface{}) bool
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof" json:"exact,omitempty"`
}
type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}
type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof" json:"suffix,omitempty"`
}
type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof" json:"safe_regex,omitempty"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern()     {}
func (*StringMatcher_Prefix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_Suffix) isStringMatcher_MatchPattern()    {}
func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

func (m *StringMatcher) GetIgnoreCase() bool {
	if m != nil {
		return m.IgnoreCase
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdc3a85644c05785, []int{1}
}
func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStringMatcher.Unmarshal(m, b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return xxx_messageInfo_ListStringMatcher.Size(m)
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.v3.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.v3.ListStringMatcher")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/matcher/v3/string.proto", fileDescriptor_bdc3a85644c05785)
}

var fileDescriptor_bdc3a85644c05785 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xee, 0x6c, 0xb2, 0x71, 0x3b, 0xa1, 0x10, 0x97, 0xaa, 0x4b, 0xc0, 0xba, 0x49, 0x8a, 0x06,
	0xc4, 0x1d, 0x68, 0x6e, 0x3d, 0xae, 0x97, 0x52, 0xaa, 0x94, 0x15, 0x2f, 0x5e, 0xc2, 0x74, 0x33,
	0xd9, 0x8e, 0xa6, 0xfb, 0x86, 0x99, 0xc9, 0xb2, 0xb9, 0x79, 0x14, 0x11, 0x0f, 0x1e, 0xfd, 0x05,
	0xe2, 0x4f, 0xf0, 0x2e, 0x78, 0xf5, 0x2f, 0xf8, 0x2b, 0xa4, 0x27, 0x99, 0x9d, 0xb5, 0x10, 0xb2,
	0x85, 0xdc, 0xde, 0x9b, 0xef, 0xfb, 0xde, 0xfb, 0xde, 0x9b, 0x19, 0xfc, 0x3a, 0xe3, 0xfa, 0x72,
	0x79, 0x11, 0xa5, 0x70, 0x45, 0x14, 0x2c, 0xe0, 0x19, 0x07, 0x92, 0x2d, 0x00, 0x88, 0x90, 0xf0,
	0x96, 0xa5, 0x5a, 0xd9, 0x8c, 0x0a, 0x4e, 0x58, 0xa9, 0x99, 0xcc, 0xe9, 0x82, 0xb0, 0xbc, 0x80,
	0x15, 0xd1, 0x2b, 0xc1, 0xc8, 0x15, 0xd5, 0xe9, 0x25, 0x93, 0xa4, 0x98, 0x10, 0xa5, 0x25, 0xcf,
	0xb3, 0x48, 0x48, 0xd0, 0xe0, 0xdf, 0xab, 0x38, 0x91, 0xe1, 0x44, 0x35, 0x27, 0x2a, 0x26, 0xfd,
	0x41, 0xb3, 0x54, 0xb2, 0x8c, 0x95, 0x56, 0xd9, 0x1f, 0x59, 0x0a, 0xcd, 0x73, 0xd0, 0x54, 0x73,
	0xc8, 0x15, 0x99, 0x31, 0x21, 0x59, 0x5a, 0x25, 0x35, 0xe9, 0xe1, 0x72, 0x26, 0xe8, 0x1a, 0x47,
	0x69, 0xaa, 0x97, 0xaa, 0x86, 0x07, 0x1b, 0x70, 0xc1, 0xa4, 0xe2, 0x90, 0xdf, 0x18, 0xec, 0x3f,
	0x28, 0xe8, 0x82, 0xcf, 0xa8, 0x66, 0xe4, 0x7f, 0x50, 0x03, 0xfb, 0x19, 0x64, 0x50, 0x85, 0xc4,
	0x44, 0xf6, 0x74, 0xf8, 0xdd, 0xc1, 0x7b, 0xaf, 0xaa, 0x01, 0x5f, 0x58, 0xdb, 0xfe, 0x7d, 0xec,
	0xb2, 0x92, 0xa6, 0x3a, 0x40, 0x21, 0x1a, 0xef, 0x9e, 0xec, 0x24, 0x36, 0xf5, 0x07, 0xb8, 0x23,
	0x24, 0x9b, 0xf3, 0x32, 0x70, 0x0c, 0x10, 0xdf, 0xb9, 0x8e, 0xdb, 0xd2, 0x09, 0xd1, 0xc9, 0x4e,
	0x52, 0x03, 0x86, 0xa2, 0x96, 0x73, 0x43, 0x69, 0x6d, 0x50, 0x2c, 0xe0, 0xbf, 0xc4, 0x58, 0xd1,
	0x39, 0x9b, 0x56, 0x9b, 0x09, 0xdc, 0x10, 0x8d, 0xbb, 0x47, 0xa3, 0xa8, 0x71, 0xa9, 0x51, 0x62,
	0x38, 0xb5, 0xad, 0xd8, 0xbb, 0x8e, 0xdd, 0x8f, 0xc8, 0xe9, 0x99, 0x62, 0xbb, 0xa6, 0x44, 0x85,
	0xfa, 0x8f, 0x70, 0x97, 0x67, 0x39, 0x48, 0x36, 0x4d, 0xa9, 0x62, 0x41, 0x27, 0x44, 0x63, 0x2f,
	0xc1, 0xf6, 0xe8, 0x39, 0x55, 0xec, 0xf8, 0xc9, 0xd7, 0x9f, 0x1f, 0x0e, 0x86, 0x38, 0x6c, 0x68,
	0xb1, 0x36, 0x77, 0xbc, 0x8f, 0xf7, 0x2a, 0x60, 0x2a, 0xa8, 0x36, 0xef, 0xc1, 0x6f, 0xfd, 0x8d,
	0xd1, 0x69, 0xdb, 0x6b, 0xf7, 0xdc, 0xc4, 0xad, 0xec, 0x0e, 0x3f, 0x21, 0x7c, 0xf7, 0x8c, 0x2b,
	0xbd, 0xbe, 0xb0, 0x53, 0xec, 0xd5, 0x12, 0x15, 0xa0, 0xb0, 0x35, 0xee, 0x1e, 0x1d, 0xde, 0x32,
	0xd0, 0x7a, 0x43, 0x33, 0xd1, 0x17, 0xe4, 0x78, 0x28, 0xb9, 0xd1, 0x1f, 0x3f, 0x35, 0x6e, 0x1f,
	0xe3, 0xc3, 0x06, 0xfd, 0x46, 0xe3, 0xf8, 0x33, 0xfa, 0xf6, 0xe7, 0x00, 0xfd, 0x78, 0xff, 0xeb,
	0x77, 0xc7, 0xe9, 0x39, 0x78, 0xc4, 0xc1, 0xf6, 0x15, 0x12, 0xca, 0x55, 0xb3, 0x85, 0xb8, 0x6b,
	0x4b, 0x9c, 0x9b, 0xcb, 0x3f, 0x47, 0x6f, 0xce, 0xb6, 0xfb, 0x25, 0xe2, 0x5d, 0xb6, 0xc5, 0x4f,
	0xb9, 0xe8, 0x54, 0x6f, 0x6a, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x51, 0x81, 0x20, 0xc8, 0x7c,
	0x03, 0x00, 0x00,
}

func (this *StringMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringMatcher)
	if !ok {
		that2, ok := that.(StringMatcher)
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
	if that1.MatchPattern == nil {
		if this.MatchPattern != nil {
			return false
		}
	} else if this.MatchPattern == nil {
		return false
	} else if !this.MatchPattern.Equal(that1.MatchPattern) {
		return false
	}
	if this.IgnoreCase != that1.IgnoreCase {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *StringMatcher_Exact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringMatcher_Exact)
	if !ok {
		that2, ok := that.(StringMatcher_Exact)
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
func (this *StringMatcher_Prefix) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringMatcher_Prefix)
	if !ok {
		that2, ok := that.(StringMatcher_Prefix)
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
func (this *StringMatcher_Suffix) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringMatcher_Suffix)
	if !ok {
		that2, ok := that.(StringMatcher_Suffix)
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
	if this.Suffix != that1.Suffix {
		return false
	}
	return true
}
func (this *StringMatcher_SafeRegex) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StringMatcher_SafeRegex)
	if !ok {
		that2, ok := that.(StringMatcher_SafeRegex)
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
	if !this.SafeRegex.Equal(that1.SafeRegex) {
		return false
	}
	return true
}
func (this *ListStringMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListStringMatcher)
	if !ok {
		that2, ok := that.(ListStringMatcher)
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
	if len(this.Patterns) != len(that1.Patterns) {
		return false
	}
	for i := range this.Patterns {
		if !this.Patterns[i].Equal(that1.Patterns[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
