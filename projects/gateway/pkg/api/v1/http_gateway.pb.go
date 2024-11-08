// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.8
// source: github.com/solo-io/gloo/projects/gateway/api/v1/http_gateway.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Virtual Service Selector expression operator, while the set-based syntax differs from Kubernetes (kubernetes: `key: !mylabel`, gloo: `key: mylabel, operator: "!"` | kubernetes: `key: mylabel`, gloo: `key: mylabel, operator: exists`), the functionality remains the same.
type VirtualServiceSelectorExpressions_Expression_Operator int32

const (
	// =
	VirtualServiceSelectorExpressions_Expression_Equals VirtualServiceSelectorExpressions_Expression_Operator = 0
	// ==
	VirtualServiceSelectorExpressions_Expression_DoubleEquals VirtualServiceSelectorExpressions_Expression_Operator = 1
	// !=
	VirtualServiceSelectorExpressions_Expression_NotEquals VirtualServiceSelectorExpressions_Expression_Operator = 2
	// in
	VirtualServiceSelectorExpressions_Expression_In VirtualServiceSelectorExpressions_Expression_Operator = 3
	// notin
	VirtualServiceSelectorExpressions_Expression_NotIn VirtualServiceSelectorExpressions_Expression_Operator = 4
	// exists
	VirtualServiceSelectorExpressions_Expression_Exists VirtualServiceSelectorExpressions_Expression_Operator = 5
	// !
	VirtualServiceSelectorExpressions_Expression_DoesNotExist VirtualServiceSelectorExpressions_Expression_Operator = 6
	// gt
	VirtualServiceSelectorExpressions_Expression_GreaterThan VirtualServiceSelectorExpressions_Expression_Operator = 7
	// lt
	VirtualServiceSelectorExpressions_Expression_LessThan VirtualServiceSelectorExpressions_Expression_Operator = 8
)

// Enum value maps for VirtualServiceSelectorExpressions_Expression_Operator.
var (
	VirtualServiceSelectorExpressions_Expression_Operator_name = map[int32]string{
		0: "Equals",
		1: "DoubleEquals",
		2: "NotEquals",
		3: "In",
		4: "NotIn",
		5: "Exists",
		6: "DoesNotExist",
		7: "GreaterThan",
		8: "LessThan",
	}
	VirtualServiceSelectorExpressions_Expression_Operator_value = map[string]int32{
		"Equals":       0,
		"DoubleEquals": 1,
		"NotEquals":    2,
		"In":           3,
		"NotIn":        4,
		"Exists":       5,
		"DoesNotExist": 6,
		"GreaterThan":  7,
		"LessThan":     8,
	}
)

func (x VirtualServiceSelectorExpressions_Expression_Operator) Enum() *VirtualServiceSelectorExpressions_Expression_Operator {
	p := new(VirtualServiceSelectorExpressions_Expression_Operator)
	*p = x
	return p
}

func (x VirtualServiceSelectorExpressions_Expression_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VirtualServiceSelectorExpressions_Expression_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_enumTypes[0].Descriptor()
}

func (VirtualServiceSelectorExpressions_Expression_Operator) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_enumTypes[0]
}

func (x VirtualServiceSelectorExpressions_Expression_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VirtualServiceSelectorExpressions_Expression_Operator.Descriptor instead.
func (VirtualServiceSelectorExpressions_Expression_Operator) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescGZIP(), []int{1, 0, 0}
}

type HttpGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Names & namespace refs of the virtual services which contain the actual routes for the gateway.
	// If the list is empty, all virtual services in all namespaces that Gloo watches will apply,
	// with accordance to `ssl` flag on `Gateway` above.
	// The default namespace matching behavior can be overridden via `virtual_service_namespaces` flag below.
	// Only one of `virtualServices`, `virtualServiceExpressions` or `virtualServiceSelector` should be provided.
	// If more than one is provided only one will be checked with priority virtualServiceExpressions, virtualServiceSelector, virtualServices
	VirtualServices []*core.ResourceRef `protobuf:"bytes,1,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services,omitempty"`
	// Select virtual services by their label. If `virtual_service_namespaces` is provided below, this will apply only
	// to virtual services in the namespaces specified.
	// Only one of `virtualServices`, `virtualServiceExpressions` or `virtualServiceSelector` should be provided.
	// If more than one is provided only one will be checked with priority virtualServiceExpressions, virtualServiceSelector, virtualServices
	VirtualServiceSelector map[string]string `protobuf:"bytes,2,rep,name=virtual_service_selector,json=virtualServiceSelector,proto3" json:"virtual_service_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Select virtual services using expressions. If `virtual_service_namespaces` is provided below, this will apply only
	// to virtual services in the namespaces specified.
	// Only one of `virtualServices`, `virtualServiceExpressions` or `virtualServiceSelector` should be provided.
	// If more than one is provided only one will be checked with priority virtualServiceExpressions, virtualServiceSelector, virtualServices
	VirtualServiceExpressions *VirtualServiceSelectorExpressions `protobuf:"bytes,9,opt,name=virtual_service_expressions,json=virtualServiceExpressions,proto3" json:"virtual_service_expressions,omitempty"`
	// Restrict the search by providing a list of valid search namespaces here.
	// Setting '*' will search all namespaces, equivalent to omitting this value.
	VirtualServiceNamespaces []string `protobuf:"bytes,3,rep,name=virtual_service_namespaces,json=virtualServiceNamespaces,proto3" json:"virtual_service_namespaces,omitempty"`
	// HTTP Gateway configuration
	Options *v1.HttpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *HttpGateway) Reset() {
	*x = HttpGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpGateway) ProtoMessage() {}

func (x *HttpGateway) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpGateway.ProtoReflect.Descriptor instead.
func (*HttpGateway) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *HttpGateway) GetVirtualServices() []*core.ResourceRef {
	if x != nil {
		return x.VirtualServices
	}
	return nil
}

func (x *HttpGateway) GetVirtualServiceSelector() map[string]string {
	if x != nil {
		return x.VirtualServiceSelector
	}
	return nil
}

func (x *HttpGateway) GetVirtualServiceExpressions() *VirtualServiceSelectorExpressions {
	if x != nil {
		return x.VirtualServiceExpressions
	}
	return nil
}

func (x *HttpGateway) GetVirtualServiceNamespaces() []string {
	if x != nil {
		return x.VirtualServiceNamespaces
	}
	return nil
}

func (x *HttpGateway) GetOptions() *v1.HttpListenerOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

// Expressions to define which virtual services to select
// Example:
// expressions:
//   - key: domain
//     operator: in
//     values: example.com
type VirtualServiceSelectorExpressions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expressions allow for more flexible virtual service label matching, such as equality-based requirements, set-based requirements, or a combination of both.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#equality-based-requirement
	Expressions []*VirtualServiceSelectorExpressions_Expression `protobuf:"bytes,3,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *VirtualServiceSelectorExpressions) Reset() {
	*x = VirtualServiceSelectorExpressions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualServiceSelectorExpressions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualServiceSelectorExpressions) ProtoMessage() {}

func (x *VirtualServiceSelectorExpressions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualServiceSelectorExpressions.ProtoReflect.Descriptor instead.
func (*VirtualServiceSelectorExpressions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *VirtualServiceSelectorExpressions) GetExpressions() []*VirtualServiceSelectorExpressions_Expression {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type VirtualServiceSelectorExpressions_Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes label key, must conform to Kubernetes syntax requirements
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The operator can only be in, notin, =, ==, !=, exists, ! (DoesNotExist), gt (GreaterThan), lt (LessThan).
	Operator VirtualServiceSelectorExpressions_Expression_Operator `protobuf:"varint,2,opt,name=operator,proto3,enum=gateway.solo.io.VirtualServiceSelectorExpressions_Expression_Operator" json:"operator,omitempty"`
	Values   []string                                              `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *VirtualServiceSelectorExpressions_Expression) Reset() {
	*x = VirtualServiceSelectorExpressions_Expression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualServiceSelectorExpressions_Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualServiceSelectorExpressions_Expression) ProtoMessage() {}

func (x *VirtualServiceSelectorExpressions_Expression) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualServiceSelectorExpressions_Expression.ProtoReflect.Descriptor instead.
func (*VirtualServiceSelectorExpressions_Expression) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescGZIP(), []int{1, 0}
}

func (x *VirtualServiceSelectorExpressions_Expression) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *VirtualServiceSelectorExpressions_Expression) GetOperator() VirtualServiceSelectorExpressions_Expression_Operator {
	if x != nil {
		return x.Operator
	}
	return VirtualServiceSelectorExpressions_Expression_Equals
}

func (x *VirtualServiceSelectorExpressions_Expression) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x04, 0x0a, 0x0b,
	0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x44, 0x0a, 0x10, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66,
	0x52, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x72, 0x0a, 0x18, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x72, 0x0a, 0x1b, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x19,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x49, 0x0a, 0x1b, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xab, 0x03, 0x0a, 0x21, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xa4, 0x02, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x62, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x02, 0x12, 0x06, 0x0a,
	0x02, 0x49, 0x6e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x49, 0x6e, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c,
	0x44, 0x6f, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x06, 0x12, 0x0f,
	0x0a, 0x0b, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x10, 0x07, 0x12,
	0x0c, 0x0a, 0x08, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x10, 0x08, 0x42, 0x41, 0xb8,
	0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_goTypes = []any{
	(VirtualServiceSelectorExpressions_Expression_Operator)(0), // 0: gateway.solo.io.VirtualServiceSelectorExpressions.Expression.Operator
	(*HttpGateway)(nil),                       // 1: gateway.solo.io.HttpGateway
	(*VirtualServiceSelectorExpressions)(nil), // 2: gateway.solo.io.VirtualServiceSelectorExpressions
	nil, // 3: gateway.solo.io.HttpGateway.VirtualServiceSelectorEntry
	(*VirtualServiceSelectorExpressions_Expression)(nil), // 4: gateway.solo.io.VirtualServiceSelectorExpressions.Expression
	(*core.ResourceRef)(nil),                             // 5: core.solo.io.ResourceRef
	(*v1.HttpListenerOptions)(nil),                       // 6: gloo.solo.io.HttpListenerOptions
}
var file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_depIdxs = []int32{
	5, // 0: gateway.solo.io.HttpGateway.virtual_services:type_name -> core.solo.io.ResourceRef
	3, // 1: gateway.solo.io.HttpGateway.virtual_service_selector:type_name -> gateway.solo.io.HttpGateway.VirtualServiceSelectorEntry
	2, // 2: gateway.solo.io.HttpGateway.virtual_service_expressions:type_name -> gateway.solo.io.VirtualServiceSelectorExpressions
	6, // 3: gateway.solo.io.HttpGateway.options:type_name -> gloo.solo.io.HttpListenerOptions
	4, // 4: gateway.solo.io.VirtualServiceSelectorExpressions.expressions:type_name -> gateway.solo.io.VirtualServiceSelectorExpressions.Expression
	0, // 5: gateway.solo.io.VirtualServiceSelectorExpressions.Expression.operator:type_name -> gateway.solo.io.VirtualServiceSelectorExpressions.Expression.Operator
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_init() }
func file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HttpGateway); i {
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
		file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VirtualServiceSelectorExpressions); i {
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
		file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VirtualServiceSelectorExpressions_Expression); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway_api_v1_http_gateway_proto_depIdxs = nil
}
