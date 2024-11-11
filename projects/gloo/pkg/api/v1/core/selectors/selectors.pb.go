// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/core/selectors/selectors.proto

package selectors

import (
	reflect "reflect"
	sync "sync"

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

// Selector expression operator, while the set-based syntax differs from Kubernetes (kubernetes: `key: !mylabel`, gloo: `key: mylabel, operator: "!"` | kubernetes: `key: mylabel`, gloo: `key: mylabel, operator: exists`), the functionality remains the same.
type Selector_Expression_Operator int32

const (
	// =
	Selector_Expression_Equals Selector_Expression_Operator = 0
	// ==
	Selector_Expression_DoubleEquals Selector_Expression_Operator = 1
	// !=
	Selector_Expression_NotEquals Selector_Expression_Operator = 2
	// in
	Selector_Expression_In Selector_Expression_Operator = 3
	// notin
	Selector_Expression_NotIn Selector_Expression_Operator = 4
	// exists
	Selector_Expression_Exists Selector_Expression_Operator = 5
	// !
	Selector_Expression_DoesNotExist Selector_Expression_Operator = 6
	// gt
	Selector_Expression_GreaterThan Selector_Expression_Operator = 7
	// lt
	Selector_Expression_LessThan Selector_Expression_Operator = 8
)

// Enum value maps for Selector_Expression_Operator.
var (
	Selector_Expression_Operator_name = map[int32]string{
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
	Selector_Expression_Operator_value = map[string]int32{
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

func (x Selector_Expression_Operator) Enum() *Selector_Expression_Operator {
	p := new(Selector_Expression_Operator)
	*p = x
	return p
}

func (x Selector_Expression_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Selector_Expression_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_enumTypes[0].Descriptor()
}

func (Selector_Expression_Operator) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_enumTypes[0]
}

func (x Selector_Expression_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Selector_Expression_Operator.Descriptor instead.
func (Selector_Expression_Operator) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescGZIP(), []int{0, 1, 0}
}

type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []string          `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Labels     map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Expressions allow for more flexible Route Tables label matching, such as equality-based requirements, set-based requirements, or a combination of both.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#equality-based-requirement
	Expressions []*Selector_Expression `protobuf:"bytes,3,rep,name=expressions,proto3" json:"expressions,omitempty"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescGZIP(), []int{0}
}

func (x *Selector) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *Selector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Selector) GetExpressions() []*Selector_Expression {
	if x != nil {
		return x.Expressions
	}
	return nil
}

type Selector_Expression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes label key, must conform to Kubernetes syntax requirements
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#syntax-and-character-set
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The operator can only be in, notin, =, ==, !=, exists, ! (DoesNotExist), gt (GreaterThan), lt (LessThan).
	Operator Selector_Expression_Operator `protobuf:"varint,2,opt,name=operator,proto3,enum=selectors.core.gloo.solo.io.Selector_Expression_Operator" json:"operator,omitempty"`
	Values   []string                     `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Selector_Expression) Reset() {
	*x = Selector_Expression{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Selector_Expression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector_Expression) ProtoMessage() {}

func (x *Selector_Expression) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector_Expression.ProtoReflect.Descriptor instead.
func (*Selector_Expression) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Selector_Expression) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Selector_Expression) GetOperator() Selector_Expression_Operator {
	if x != nil {
		return x.Operator
	}
	return Selector_Expression_Equals
}

func (x *Selector_Expression) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e,
	0x04, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x52, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x97, 0x02, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x55, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x49, 0x6e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x49, 0x6e,
	0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x05, 0x12, 0x10,
	0x0a, 0x0c, 0x44, 0x6f, 0x65, 0x73, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x06,
	0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x54, 0x68, 0x61, 0x6e, 0x10,
	0x07, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x65, 0x73, 0x73, 0x54, 0x68, 0x61, 0x6e, 0x10, 0x08, 0x42,
	0x4d, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_goTypes = []any{
	(Selector_Expression_Operator)(0), // 0: selectors.core.gloo.solo.io.Selector.Expression.Operator
	(*Selector)(nil),                  // 1: selectors.core.gloo.solo.io.Selector
	nil,                               // 2: selectors.core.gloo.solo.io.Selector.LabelsEntry
	(*Selector_Expression)(nil),       // 3: selectors.core.gloo.solo.io.Selector.Expression
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_depIdxs = []int32{
	2, // 0: selectors.core.gloo.solo.io.Selector.labels:type_name -> selectors.core.gloo.solo.io.Selector.LabelsEntry
	3, // 1: selectors.core.gloo.solo.io.Selector.expressions:type_name -> selectors.core.gloo.solo.io.Selector.Expression
	0, // 2: selectors.core.gloo.solo.io.Selector.Expression.operator:type_name -> selectors.core.gloo.solo.io.Selector.Expression.Operator
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_core_selectors_selectors_proto_depIdxs = nil
}
