// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/advanced_http/advanced_http.proto

package advanced_http

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckResult int32

const (
	HealthCheckResult_default   HealthCheckResult = 0 // context-specific, defaults to one of the following values
	HealthCheckResult_healthy   HealthCheckResult = 1
	HealthCheckResult_degraded  HealthCheckResult = 2
	HealthCheckResult_unhealthy HealthCheckResult = 3
)

// Enum value maps for HealthCheckResult.
var (
	HealthCheckResult_name = map[int32]string{
		0: "default",
		1: "healthy",
		2: "degraded",
		3: "unhealthy",
	}
	HealthCheckResult_value = map[string]int32{
		"default":   0,
		"healthy":   1,
		"degraded":  2,
		"unhealthy": 3,
	}
)

func (x HealthCheckResult) Enum() *HealthCheckResult {
	p := new(HealthCheckResult)
	*p = x
	return p
}

func (x HealthCheckResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResult) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResult) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_enumTypes[0]
}

func (x HealthCheckResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResult.Descriptor instead.
func (HealthCheckResult) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP(), []int{0}
}

type ResponseAssertions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A bunch of match rules, the first match wins out and short-circuits
	ResponseMatchers []*ResponseMatcher `protobuf:"bytes,1,rep,name=response_matchers,json=responseMatchers,proto3" json:"response_matchers,omitempty"`
	// The default health response if none of the response health checks were matches.
	// If omitted, defaults to unhealthy.
	NoMatchHealth HealthCheckResult `protobuf:"varint,2,opt,name=no_match_health,json=noMatchHealth,proto3,enum=advancedhttp.options.gloo.solo.io.HealthCheckResult" json:"no_match_health,omitempty"`
}

func (x *ResponseAssertions) Reset() {
	*x = ResponseAssertions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAssertions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAssertions) ProtoMessage() {}

func (x *ResponseAssertions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAssertions.ProtoReflect.Descriptor instead.
func (*ResponseAssertions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseAssertions) GetResponseMatchers() []*ResponseMatcher {
	if x != nil {
		return x.ResponseMatchers
	}
	return nil
}

func (x *ResponseAssertions) GetNoMatchHealth() HealthCheckResult {
	if x != nil {
		return x.NoMatchHealth
	}
	return HealthCheckResult_default
}

// Defines a transformation template.
type ResponseMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the parameters to determine a single match
	ResponseMatch *ResponseMatch `protobuf:"bytes,1,opt,name=response_match,json=responseMatch,proto3" json:"response_match,omitempty"`
	// The health response if this response_match is a match.
	// If omitted, defaults to healthy
	MatchHealth HealthCheckResult `protobuf:"varint,2,opt,name=match_health,json=matchHealth,proto3,enum=advancedhttp.options.gloo.solo.io.HealthCheckResult" json:"match_health,omitempty"`
}

func (x *ResponseMatcher) Reset() {
	*x = ResponseMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMatcher) ProtoMessage() {}

func (x *ResponseMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMatcher.ProtoReflect.Descriptor instead.
func (*ResponseMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMatcher) GetResponseMatch() *ResponseMatch {
	if x != nil {
		return x.ResponseMatch
	}
	return nil
}

func (x *ResponseMatcher) GetMatchHealth() HealthCheckResult {
	if x != nil {
		return x.MatchHealth
	}
	return HealthCheckResult_default
}

// ResponseMatches can be used to extract information from the request/response.
type ResponseMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration to get the json key.
	// Treats the body as raw text if omitted.
	JsonKey *JsonKey `protobuf:"bytes,1,opt,name=json_key,json=jsonKey,proto3" json:"json_key,omitempty"`
	// If set to true, Envoy will not throw an exception in case the json body parsing
	// fails.
	IgnoreErrorOnParse bool `protobuf:"varint,2,opt,name=ignore_error_on_parse,json=ignoreErrorOnParse,proto3" json:"ignore_error_on_parse,omitempty"`
	// The source of the extraction
	//
	// Types that are assignable to Source:
	//
	//	*ResponseMatch_Header
	//	*ResponseMatch_Body
	Source isResponseMatch_Source `protobuf_oneof:"source"`
	// Only strings matching this regular expression will be considered a match.
	// The most simple value for this field is '.*', which matches the
	// whole source. The field is required.
	Regex string `protobuf:"bytes,5,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (x *ResponseMatch) Reset() {
	*x = ResponseMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMatch) ProtoMessage() {}

func (x *ResponseMatch) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMatch.ProtoReflect.Descriptor instead.
func (*ResponseMatch) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseMatch) GetJsonKey() *JsonKey {
	if x != nil {
		return x.JsonKey
	}
	return nil
}

func (x *ResponseMatch) GetIgnoreErrorOnParse() bool {
	if x != nil {
		return x.IgnoreErrorOnParse
	}
	return false
}

func (m *ResponseMatch) GetSource() isResponseMatch_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ResponseMatch) GetHeader() string {
	if x, ok := x.GetSource().(*ResponseMatch_Header); ok {
		return x.Header
	}
	return ""
}

func (x *ResponseMatch) GetBody() *empty.Empty {
	if x, ok := x.GetSource().(*ResponseMatch_Body); ok {
		return x.Body
	}
	return nil
}

func (x *ResponseMatch) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

type isResponseMatch_Source interface {
	isResponseMatch_Source()
}

type ResponseMatch_Header struct {
	// Extract information from headers
	Header string `protobuf:"bytes,3,opt,name=header,proto3,oneof"`
}

type ResponseMatch_Body struct {
	// Extract information from the request/response body
	Body *empty.Empty `protobuf:"bytes,4,opt,name=body,proto3,oneof"`
}

func (*ResponseMatch_Header) isResponseMatch_Source() {}

func (*ResponseMatch_Body) isResponseMatch_Source() {}

type JsonKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to retrieve the Value.
	Path []*JsonKey_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *JsonKey) Reset() {
	*x = JsonKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonKey) ProtoMessage() {}

func (x *JsonKey) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonKey.ProtoReflect.Descriptor instead.
func (*JsonKey) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP(), []int{3}
}

func (x *JsonKey) GetPath() []*JsonKey_PathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

// Specifies the segment in a path to retrieve value.
type JsonKey_PathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Segment:
	//
	//	*JsonKey_PathSegment_Key
	Segment isJsonKey_PathSegment_Segment `protobuf_oneof:"segment"`
}

func (x *JsonKey_PathSegment) Reset() {
	*x = JsonKey_PathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonKey_PathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonKey_PathSegment) ProtoMessage() {}

func (x *JsonKey_PathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonKey_PathSegment.ProtoReflect.Descriptor instead.
func (*JsonKey_PathSegment) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP(), []int{3, 0}
}

func (m *JsonKey_PathSegment) GetSegment() isJsonKey_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (x *JsonKey_PathSegment) GetKey() string {
	if x, ok := x.GetSegment().(*JsonKey_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

type isJsonKey_PathSegment_Segment interface {
	isJsonKey_PathSegment_Segment()
}

type JsonKey_PathSegment_Key struct {
	// If specified, use the key to retrieve the value.
	// If the key is not found, the value defaults to empty string.
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*JsonKey_PathSegment_Key) isJsonKey_PathSegment_Segment() {}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5f, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x5c, 0x0a, 0x0f, 0x6e, 0x6f, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x6e, 0x6f, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xf1, 0x01,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x45, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x6a,
	0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4f, 0x6e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x9b, 0x01, 0x0a, 0x07, 0x4a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x54, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x61, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x1a, 0x3a, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42,
	0x0e, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x2a,
	0x4a, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x75, 0x6e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x10, 0x03, 0x42, 0x94, 0x01, 0x0a, 0x3a,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x41, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x48, 0x74, 0x74, 0x70, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_goTypes = []interface{}{
	(HealthCheckResult)(0),      // 0: advancedhttp.options.gloo.solo.io.HealthCheckResult
	(*ResponseAssertions)(nil),  // 1: advancedhttp.options.gloo.solo.io.ResponseAssertions
	(*ResponseMatcher)(nil),     // 2: advancedhttp.options.gloo.solo.io.ResponseMatcher
	(*ResponseMatch)(nil),       // 3: advancedhttp.options.gloo.solo.io.ResponseMatch
	(*JsonKey)(nil),             // 4: advancedhttp.options.gloo.solo.io.JsonKey
	(*JsonKey_PathSegment)(nil), // 5: advancedhttp.options.gloo.solo.io.JsonKey.PathSegment
	(*empty.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_depIdxs = []int32{
	2, // 0: advancedhttp.options.gloo.solo.io.ResponseAssertions.response_matchers:type_name -> advancedhttp.options.gloo.solo.io.ResponseMatcher
	0, // 1: advancedhttp.options.gloo.solo.io.ResponseAssertions.no_match_health:type_name -> advancedhttp.options.gloo.solo.io.HealthCheckResult
	3, // 2: advancedhttp.options.gloo.solo.io.ResponseMatcher.response_match:type_name -> advancedhttp.options.gloo.solo.io.ResponseMatch
	0, // 3: advancedhttp.options.gloo.solo.io.ResponseMatcher.match_health:type_name -> advancedhttp.options.gloo.solo.io.HealthCheckResult
	4, // 4: advancedhttp.options.gloo.solo.io.ResponseMatch.json_key:type_name -> advancedhttp.options.gloo.solo.io.JsonKey
	6, // 5: advancedhttp.options.gloo.solo.io.ResponseMatch.body:type_name -> google.protobuf.Empty
	5, // 6: advancedhttp.options.gloo.solo.io.JsonKey.path:type_name -> advancedhttp.options.gloo.solo.io.JsonKey.PathSegment
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAssertions); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMatcher); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMatch); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonKey); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonKey_PathSegment); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ResponseMatch_Header)(nil),
		(*ResponseMatch_Body)(nil),
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*JsonKey_PathSegment_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_advanced_http_advanced_http_proto_depIdxs = nil
}
