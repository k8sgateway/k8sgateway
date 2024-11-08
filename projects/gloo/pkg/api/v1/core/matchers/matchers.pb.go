// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.8
// source: github.com/solo-io/gloo/projects/gloo/api/v1/core/matchers/matchers.proto

package matchers

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Parameters for matching routes to requests received by a Gloo-managed proxy
type Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PathSpecifier:
	//
	//	*Matcher_Prefix
	//	*Matcher_Exact
	//	*Matcher_Regex
	//	*Matcher_ConnectMatcher_
	PathSpecifier isMatcher_PathSpecifier `protobuf_oneof:"path_specifier"`
	// Indicates that prefix/path matching should be case sensitive. The default is true.
	CaseSensitive *wrappers.BoolValue `protobuf:"bytes,4,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
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
	Methods []string `protobuf:"bytes,8,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescGZIP(), []int{0}
}

func (m *Matcher) GetPathSpecifier() isMatcher_PathSpecifier {
	if m != nil {
		return m.PathSpecifier
	}
	return nil
}

func (x *Matcher) GetPrefix() string {
	if x, ok := x.GetPathSpecifier().(*Matcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *Matcher) GetExact() string {
	if x, ok := x.GetPathSpecifier().(*Matcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *Matcher) GetRegex() string {
	if x, ok := x.GetPathSpecifier().(*Matcher_Regex); ok {
		return x.Regex
	}
	return ""
}

func (x *Matcher) GetConnectMatcher() *Matcher_ConnectMatcher {
	if x, ok := x.GetPathSpecifier().(*Matcher_ConnectMatcher_); ok {
		return x.ConnectMatcher
	}
	return nil
}

func (x *Matcher) GetCaseSensitive() *wrappers.BoolValue {
	if x != nil {
		return x.CaseSensitive
	}
	return nil
}

func (x *Matcher) GetHeaders() []*HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Matcher) GetQueryParameters() []*QueryParameterMatcher {
	if x != nil {
		return x.QueryParameters
	}
	return nil
}

func (x *Matcher) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type isMatcher_PathSpecifier interface {
	isMatcher_PathSpecifier()
}

type Matcher_Prefix struct {
	// If specified, the route is a prefix rule meaning that the prefix must
	// match the beginning of the *:path* header.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3,oneof"`
}

type Matcher_Exact struct {
	// If specified, the route is an exact path rule meaning that the path must
	// exactly match the *:path* header once the query string is removed.
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof"`
}

type Matcher_Regex struct {
	// If specified, the route is a regular expression rule meaning that the
	// regex must match the *:path* header once the query string is removed. The entire path
	// (without the query string) must match the regex. The rule will not match if only a
	// sub-sequence of the *:path* header matches the regex. The regex grammar is defined `here
	// <http://en.cppreference.com/w/cpp/regex/ecmascript>`_.
	//
	// Examples:<br/>
	//
	// * The regex */b[io]t* matches the path */bit*<br/>
	// * The regex */b[io]t* matches the path */bot*<br/>
	// * The regex */b[io]t* does not match the path */bite*<br/>
	// * The regex */b[io]t* does not match the path */bit/bot*<br/><br/>
	//
	// Note that the complexity of the regex is constrained by the regex engine's "program size" setting.
	// If your regex is too complex, you may need to adjust the `regexMaxProgramSize` field
	// in the `GlooOptions` section of your `Settings` resource (The gloo default is 1024)
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

type Matcher_ConnectMatcher_ struct {
	// If this is used as the matcher, the matcher will only match CONNECT requests.
	// Note that this will not match HTTP/2 upgrade-style CONNECT requests
	// (WebSocket and the like) as they are normalized in Envoy as HTTP/1.1 style
	// upgrades.
	// This is the only way to match CONNECT requests for HTTP/1.1. For HTTP/2,
	// where CONNECT requests may have a path, the path matchers will work if
	// there is a path present.
	// Note that CONNECT support is currently considered alpha in Envoy.
	ConnectMatcher *Matcher_ConnectMatcher `protobuf:"bytes,9,opt,name=connect_matcher,json=connectMatcher,proto3,oneof"`
}

func (*Matcher_Prefix) isMatcher_PathSpecifier() {}

func (*Matcher_Exact) isMatcher_PathSpecifier() {}

func (*Matcher_Regex) isMatcher_PathSpecifier() {}

func (*Matcher_ConnectMatcher_) isMatcher_PathSpecifier() {}

// Internally, Gloo always uses the HTTP/2 *:authority* header to represent the HTTP/1 *Host* header.
// Thus, if attempting to match on *Host*, match on *:authority* instead.
type HeaderMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the name of the header in the request.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the value of the header. If the value is absent a request that
	// has the name header will match, regardless of the header’s value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Specifies whether the header value should be treated as regex or not.
	Regex bool `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
	// If set to true, the result of the match will be inverted. Defaults to false.
	//
	// Examples:
	// * name=foo, invert_match=true: matches if no header named `foo` is present
	// * name=foo, value=bar, invert_match=true: matches if no header named `foo` with value `bar` is present
	// * name=foo, value=“\d{3}“, regex=true, invert_match=true: matches if no header named `foo` with a value consisting of three integers is present
	InvertMatch bool `protobuf:"varint,4,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
}

func (x *HeaderMatcher) Reset() {
	*x = HeaderMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMatcher) ProtoMessage() {}

func (x *HeaderMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMatcher.ProtoReflect.Descriptor instead.
func (*HeaderMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeaderMatcher) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *HeaderMatcher) GetRegex() bool {
	if x != nil {
		return x.Regex
	}
	return false
}

func (x *HeaderMatcher) GetInvertMatch() bool {
	if x != nil {
		return x.InvertMatch
	}
	return false
}

// Query parameter matching treats the query string of a request's :path header
// as an ampersand-separated list of keys and/or key=value elements.
type QueryParameterMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	Regex bool `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (x *QueryParameterMatcher) Reset() {
	*x = QueryParameterMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParameterMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParameterMatcher) ProtoMessage() {}

func (x *QueryParameterMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParameterMatcher.ProtoReflect.Descriptor instead.
func (*QueryParameterMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescGZIP(), []int{2}
}

func (x *QueryParameterMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryParameterMatcher) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *QueryParameterMatcher) GetRegex() bool {
	if x != nil {
		return x.Regex
	}
	return false
}

type Matcher_ConnectMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Matcher_ConnectMatcher) Reset() {
	*x = Matcher_ConnectMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher_ConnectMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher_ConnectMatcher) ProtoMessage() {}

func (x *Matcher_ConnectMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher_ConnectMatcher.ProtoReflect.Descriptor instead.
func (*Matcher_ConnectMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescGZIP(), []int{0, 0}
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x03, 0x0a, 0x07,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x12, 0x5d, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x41, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x10, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x1a, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x10, 0x0a, 0x0e, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x57, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x42, 0x4c, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_goTypes = []interface{}{
	(*Matcher)(nil),                // 0: matchers.core.gloo.solo.io.Matcher
	(*HeaderMatcher)(nil),          // 1: matchers.core.gloo.solo.io.HeaderMatcher
	(*QueryParameterMatcher)(nil),  // 2: matchers.core.gloo.solo.io.QueryParameterMatcher
	(*Matcher_ConnectMatcher)(nil), // 3: matchers.core.gloo.solo.io.Matcher.ConnectMatcher
	(*wrappers.BoolValue)(nil),     // 4: google.protobuf.BoolValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_depIdxs = []int32{
	3, // 0: matchers.core.gloo.solo.io.Matcher.connect_matcher:type_name -> matchers.core.gloo.solo.io.Matcher.ConnectMatcher
	4, // 1: matchers.core.gloo.solo.io.Matcher.case_sensitive:type_name -> google.protobuf.BoolValue
	1, // 2: matchers.core.gloo.solo.io.Matcher.headers:type_name -> matchers.core.gloo.solo.io.HeaderMatcher
	2, // 3: matchers.core.gloo.solo.io.Matcher.query_parameters:type_name -> matchers.core.gloo.solo.io.QueryParameterMatcher
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParameterMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher_ConnectMatcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Matcher_Prefix)(nil),
		(*Matcher_Exact)(nil),
		(*Matcher_Regex)(nil),
		(*Matcher_ConnectMatcher_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_matchers_matchers_proto_depIdxs = nil
}
