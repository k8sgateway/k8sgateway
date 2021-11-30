// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/transformation_ee/transformation.proto

package transformation_ee

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	route "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/route"
	_type "github.com/solo-io/solo-kit/pkg/api/external/envoy/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FilterTransformations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies transformations based on the route matches. The first matched transformation will be
	// applied. If there are overlapped match conditions, please put the most specific match first.
	Transformations []*TransformationRule `protobuf:"bytes,1,rep,name=transformations,proto3" json:"transformations,omitempty"`
}

func (x *FilterTransformations) Reset() {
	*x = FilterTransformations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterTransformations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterTransformations) ProtoMessage() {}

func (x *FilterTransformations) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterTransformations.ProtoReflect.Descriptor instead.
func (*FilterTransformations) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{0}
}

func (x *FilterTransformations) GetTransformations() []*TransformationRule {
	if x != nil {
		return x.Transformations
	}
	return nil
}

type TransformationRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The route matching parameter. Only when the match is satisfied, the "requires" field will
	// apply.
	//
	// For example: following match will match all requests.
	//
	// .. code-block:: yaml
	//
	//    match:
	//      prefix: /
	//
	Match *route.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// transformation to perform
	RouteTransformations *RouteTransformations `protobuf:"bytes,2,opt,name=route_transformations,json=routeTransformations,proto3" json:"route_transformations,omitempty"`
}

func (x *TransformationRule) Reset() {
	*x = TransformationRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformationRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformationRule) ProtoMessage() {}

func (x *TransformationRule) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformationRule.ProtoReflect.Descriptor instead.
func (*TransformationRule) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{1}
}

func (x *TransformationRule) GetMatch() *route.RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *TransformationRule) GetRouteTransformations() *RouteTransformations {
	if x != nil {
		return x.RouteTransformations
	}
	return nil
}

type RouteTransformations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestTransformation *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	// clear the route cache if the request transformation was applied
	ClearRouteCache        bool            `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	// Apply a transformation in the onStreamComplete callback
	// (for modifying headers and dynamic metadata for access logs)
	OnStreamCompletionTransformation *Transformation `protobuf:"bytes,4,opt,name=on_stream_completion_transformation,json=onStreamCompletionTransformation,proto3" json:"on_stream_completion_transformation,omitempty"`
}

func (x *RouteTransformations) Reset() {
	*x = RouteTransformations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTransformations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTransformations) ProtoMessage() {}

func (x *RouteTransformations) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTransformations.ProtoReflect.Descriptor instead.
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{2}
}

func (x *RouteTransformations) GetRequestTransformation() *Transformation {
	if x != nil {
		return x.RequestTransformation
	}
	return nil
}

func (x *RouteTransformations) GetClearRouteCache() bool {
	if x != nil {
		return x.ClearRouteCache
	}
	return false
}

func (x *RouteTransformations) GetResponseTransformation() *Transformation {
	if x != nil {
		return x.ResponseTransformation
	}
	return nil
}

func (x *RouteTransformations) GetOnStreamCompletionTransformation() *Transformation {
	if x != nil {
		return x.OnStreamCompletionTransformation
	}
	return nil
}

type Transformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Template is in the transformed request language domain
	//
	// Types that are assignable to TransformationType:
	//	*Transformation_DlpTransformation
	TransformationType isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
}

func (x *Transformation) Reset() {
	*x = Transformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transformation) ProtoMessage() {}

func (x *Transformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transformation.ProtoReflect.Descriptor instead.
func (*Transformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{3}
}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (x *Transformation) GetDlpTransformation() *DlpTransformation {
	if x, ok := x.GetTransformationType().(*Transformation_DlpTransformation); ok {
		return x.DlpTransformation
	}
	return nil
}

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
}

type Transformation_DlpTransformation struct {
	DlpTransformation *DlpTransformation `protobuf:"bytes,1,opt,name=dlp_transformation,json=dlpTransformation,proto3,oneof"`
}

func (*Transformation_DlpTransformation) isTransformation_TransformationType() {}

type DlpTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of actions to apply
	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	// If true, headers will be transformed. Should only be true for the
	// on_stream_complete_transformation route transformation type.
	EnableHeaderTransformation bool `protobuf:"varint,2,opt,name=enable_header_transformation,json=enableHeaderTransformation,proto3" json:"enable_header_transformation,omitempty"`
	// If true, dynamic metadata will be transformed. Should only be used for the
	// on_stream_complete_transformation route transformation type.
	EnableDynamicMetadataTransformation bool `protobuf:"varint,3,opt,name=enable_dynamic_metadata_transformation,json=enableDynamicMetadataTransformation,proto3" json:"enable_dynamic_metadata_transformation,omitempty"`
}

func (x *DlpTransformation) Reset() {
	*x = DlpTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DlpTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DlpTransformation) ProtoMessage() {}

func (x *DlpTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DlpTransformation.ProtoReflect.Descriptor instead.
func (*DlpTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{4}
}

func (x *DlpTransformation) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *DlpTransformation) GetEnableHeaderTransformation() bool {
	if x != nil {
		return x.EnableHeaderTransformation
	}
	return false
}

func (x *DlpTransformation) GetEnableDynamicMetadataTransformation() bool {
	if x != nil {
		return x.EnableDynamicMetadataTransformation
	}
	return false
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier for this action.
	// Used mostly to help ID specific actions in logs.
	// If left null will default to unknown
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of regexes to apply to the response body to match data which should be masked
	// They will be applied iteratively in the order which they are specified
	Regex []string `protobuf:"bytes,2,rep,name=regex,proto3" json:"regex,omitempty"`
	// List of regexes to apply to the response body to match data which should be
	// masked. They will be applied iteratively in the order which they are
	// specified. If this field and `regex` are both provided, all the regexes will
	// be applied iteratively in the order provided, starting with the ones from `regex`
	RegexActions []*RegexAction `protobuf:"bytes,6,rep,name=regex_actions,json=regexActions,proto3" json:"regex_actions,omitempty"`
	// If specified, this rule will not actually be applied, but only logged.
	Shadow bool `protobuf:"varint,3,opt,name=shadow,proto3" json:"shadow,omitempty"`
	// The percent of the string which should be masked.
	// If not set, defaults to 75%
	Percent *_type.Percent `protobuf:"bytes,4,opt,name=percent,proto3" json:"percent,omitempty"`
	// The character which should overwrite the masked data
	// If left empty, defaults to "X"
	MaskChar string `protobuf:"bytes,5,opt,name=mask_char,json=maskChar,proto3" json:"mask_char,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{5}
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetRegex() []string {
	if x != nil {
		return x.Regex
	}
	return nil
}

func (x *Action) GetRegexActions() []*RegexAction {
	if x != nil {
		return x.RegexActions
	}
	return nil
}

func (x *Action) GetShadow() bool {
	if x != nil {
		return x.Shadow
	}
	return false
}

func (x *Action) GetPercent() *_type.Percent {
	if x != nil {
		return x.Percent
	}
	return nil
}

func (x *Action) GetMaskChar() string {
	if x != nil {
		return x.MaskChar
	}
	return ""
}

type RegexAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The regex to match for masking.
	Regex string `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	// If provided and not 0, only this specific subgroup of the regex will be masked.
	Subgroup uint32 `protobuf:"varint,2,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
}

func (x *RegexAction) Reset() {
	*x = RegexAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexAction) ProtoMessage() {}

func (x *RegexAction) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexAction.ProtoReflect.Descriptor instead.
func (*RegexAction) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP(), []int{6}
}

func (x *RegexAction) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

func (x *RegexAction) GetSubgroup() uint32 {
	if x != nil {
		return x.Subgroup
	}
	return 0
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDesc = []byte{
	0x0a, 0x6a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x6b, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd6,
	0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x78, 0x0a,
	0x15, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x14, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbf, 0x03, 0x0a, 0x14, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x74, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x76, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8c, 0x01, 0x0a, 0x23, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x20, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x0e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x12,
	0x64, 0x6c, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x6c, 0x70, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11, 0x64, 0x6c,
	0x70, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x15, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x11, 0x44, 0x6c, 0x70, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x07,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a,
	0x1c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x53, 0x0a, 0x26, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x23, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x96, 0x02, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x65, 0x67, 0x65, 0x78, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x65,
	0x78, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77,
	0x12, 0x35, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x6d, 0x61, 0x73, 0x6b, 0x5f,
	0x63, 0x68, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x6b, 0x43, 0x68, 0x61, 0x72, 0x22, 0x48, 0x0a,
	0x0b, 0x52, 0x65, 0x67, 0x65, 0x78, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x05,
	0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x75, 0x62, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0xb7, 0x01, 0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x1b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x45, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_goTypes = []interface{}{
	(*FilterTransformations)(nil), // 0: envoy.config.filter.http.transformation_ee.v2.FilterTransformations
	(*TransformationRule)(nil),    // 1: envoy.config.filter.http.transformation_ee.v2.TransformationRule
	(*RouteTransformations)(nil),  // 2: envoy.config.filter.http.transformation_ee.v2.RouteTransformations
	(*Transformation)(nil),        // 3: envoy.config.filter.http.transformation_ee.v2.Transformation
	(*DlpTransformation)(nil),     // 4: envoy.config.filter.http.transformation_ee.v2.DlpTransformation
	(*Action)(nil),                // 5: envoy.config.filter.http.transformation_ee.v2.Action
	(*RegexAction)(nil),           // 6: envoy.config.filter.http.transformation_ee.v2.RegexAction
	(*route.RouteMatch)(nil),      // 7: solo.io.envoy.api.v2.route.RouteMatch
	(*_type.Percent)(nil),         // 8: solo.io.envoy.type.Percent
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_depIdxs = []int32{
	1,  // 0: envoy.config.filter.http.transformation_ee.v2.FilterTransformations.transformations:type_name -> envoy.config.filter.http.transformation_ee.v2.TransformationRule
	7,  // 1: envoy.config.filter.http.transformation_ee.v2.TransformationRule.match:type_name -> solo.io.envoy.api.v2.route.RouteMatch
	2,  // 2: envoy.config.filter.http.transformation_ee.v2.TransformationRule.route_transformations:type_name -> envoy.config.filter.http.transformation_ee.v2.RouteTransformations
	3,  // 3: envoy.config.filter.http.transformation_ee.v2.RouteTransformations.request_transformation:type_name -> envoy.config.filter.http.transformation_ee.v2.Transformation
	3,  // 4: envoy.config.filter.http.transformation_ee.v2.RouteTransformations.response_transformation:type_name -> envoy.config.filter.http.transformation_ee.v2.Transformation
	3,  // 5: envoy.config.filter.http.transformation_ee.v2.RouteTransformations.on_stream_completion_transformation:type_name -> envoy.config.filter.http.transformation_ee.v2.Transformation
	4,  // 6: envoy.config.filter.http.transformation_ee.v2.Transformation.dlp_transformation:type_name -> envoy.config.filter.http.transformation_ee.v2.DlpTransformation
	5,  // 7: envoy.config.filter.http.transformation_ee.v2.DlpTransformation.actions:type_name -> envoy.config.filter.http.transformation_ee.v2.Action
	6,  // 8: envoy.config.filter.http.transformation_ee.v2.Action.regex_actions:type_name -> envoy.config.filter.http.transformation_ee.v2.RegexAction
	8,  // 9: envoy.config.filter.http.transformation_ee.v2.Action.percent:type_name -> solo.io.envoy.type.Percent
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterTransformations); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformationRule); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTransformations); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transformation); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DlpTransformation); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexAction); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Transformation_DlpTransformation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_transformation_ee_transformation_proto_depIdxs = nil
}
