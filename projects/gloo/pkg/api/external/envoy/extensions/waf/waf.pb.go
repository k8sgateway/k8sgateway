// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ModSecurity struct {
	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Global rule sets for the current http connection manager
	RuleSets []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	// Custom message to display when an intervention occurs
	CustomInterventionMessage string   `protobuf:"bytes,3,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *ModSecurity) Reset()         { *m = ModSecurity{} }
func (m *ModSecurity) String() string { return proto.CompactTextString(m) }
func (*ModSecurity) ProtoMessage()    {}
func (*ModSecurity) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf8967e7dbe03c2, []int{0}
}
func (m *ModSecurity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModSecurity.Unmarshal(m, b)
}
func (m *ModSecurity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModSecurity.Marshal(b, m, deterministic)
}
func (m *ModSecurity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModSecurity.Merge(m, src)
}
func (m *ModSecurity) XXX_Size() int {
	return xxx_messageInfo_ModSecurity.Size(m)
}
func (m *ModSecurity) XXX_DiscardUnknown() {
	xxx_messageInfo_ModSecurity.DiscardUnknown(m)
}

var xxx_messageInfo_ModSecurity proto.InternalMessageInfo

func (m *ModSecurity) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ModSecurity) GetRuleSets() []*RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

func (m *ModSecurity) GetCustomInterventionMessage() string {
	if m != nil {
		return m.CustomInterventionMessage
	}
	return ""
}

type RuleSet struct {
	// String of rules which are added directly
	RuleStr string `protobuf:"bytes,1,opt,name=rule_str,json=ruleStr,proto3" json:"rule_str,omitempty"`
	// Array of files to include
	Files                []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleSet) Reset()         { *m = RuleSet{} }
func (m *RuleSet) String() string { return proto.CompactTextString(m) }
func (*RuleSet) ProtoMessage()    {}
func (*RuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf8967e7dbe03c2, []int{1}
}
func (m *RuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleSet.Unmarshal(m, b)
}
func (m *RuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleSet.Marshal(b, m, deterministic)
}
func (m *RuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleSet.Merge(m, src)
}
func (m *RuleSet) XXX_Size() int {
	return xxx_messageInfo_RuleSet.Size(m)
}
func (m *RuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_RuleSet proto.InternalMessageInfo

func (m *RuleSet) GetRuleStr() string {
	if m != nil {
		return m.RuleStr
	}
	return ""
}

func (m *RuleSet) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

type ModSecurityPerRoute struct {
	// Disable all rules on the current route
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Overwrite the global rules on this route
	RuleSets []*RuleSet `protobuf:"bytes,2,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	// Custom message to display when an intervention occurs
	CustomInterventionMessage string   `protobuf:"bytes,3,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *ModSecurityPerRoute) Reset()         { *m = ModSecurityPerRoute{} }
func (m *ModSecurityPerRoute) String() string { return proto.CompactTextString(m) }
func (*ModSecurityPerRoute) ProtoMessage()    {}
func (*ModSecurityPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf8967e7dbe03c2, []int{2}
}
func (m *ModSecurityPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModSecurityPerRoute.Unmarshal(m, b)
}
func (m *ModSecurityPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModSecurityPerRoute.Marshal(b, m, deterministic)
}
func (m *ModSecurityPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModSecurityPerRoute.Merge(m, src)
}
func (m *ModSecurityPerRoute) XXX_Size() int {
	return xxx_messageInfo_ModSecurityPerRoute.Size(m)
}
func (m *ModSecurityPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ModSecurityPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ModSecurityPerRoute proto.InternalMessageInfo

func (m *ModSecurityPerRoute) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ModSecurityPerRoute) GetRuleSets() []*RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

func (m *ModSecurityPerRoute) GetCustomInterventionMessage() string {
	if m != nil {
		return m.CustomInterventionMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ModSecurity)(nil), "envoy.config.filter.http.modsecurity.v2.ModSecurity")
	proto.RegisterType((*RuleSet)(nil), "envoy.config.filter.http.modsecurity.v2.RuleSet")
	proto.RegisterType((*ModSecurityPerRoute)(nil), "envoy.config.filter.http.modsecurity.v2.ModSecurityPerRoute")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/waf/waf.proto", fileDescriptor_aaf8967e7dbe03c2)
}

var fileDescriptor_aaf8967e7dbe03c2 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0x95, 0x1b, 0xdd, 0xdb, 0xc4, 0xdd, 0x72, 0x3b, 0xa4, 0xbd, 0xd2, 0x55, 0xd4, 0x85,
	0x2c, 0x38, 0xa8, 0x6c, 0x0c, 0x0c, 0x6c, 0x08, 0x55, 0xa2, 0xe9, 0xc6, 0x52, 0xa5, 0xc9, 0x89,
	0x6b, 0x70, 0x7c, 0x22, 0xfb, 0xa4, 0xd0, 0x37, 0xe2, 0x11, 0x18, 0x79, 0x16, 0x9e, 0x04, 0x35,
	0xae, 0x50, 0x37, 0xba, 0x32, 0x58, 0x3a, 0x9f, 0xfc, 0x3b, 0x7f, 0x3e, 0xe9, 0x63, 0x0b, 0x21,
	0x69, 0xd3, 0xad, 0x79, 0x89, 0x4d, 0x66, 0x51, 0xe1, 0xb9, 0xc4, 0x4c, 0x28, 0xc4, 0xac, 0x35,
	0xf8, 0x08, 0x25, 0x59, 0xa7, 0x8a, 0x56, 0x66, 0xf0, 0x42, 0x60, 0x74, 0xa1, 0x32, 0xd0, 0x5b,
	0xdc, 0xf5, 0x52, 0x5b, 0x89, 0xda, 0x66, 0xcf, 0x45, 0xbd, 0x7f, 0xbc, 0x35, 0x48, 0x18, 0x9d,
	0xf5, 0xff, 0xbc, 0x44, 0x5d, 0x4b, 0xc1, 0x6b, 0xa9, 0x08, 0x0c, 0xdf, 0x10, 0xb5, 0xbc, 0xc1,
	0xca, 0x42, 0xd9, 0x19, 0x49, 0x3b, 0xbe, 0x9d, 0x4d, 0x46, 0x02, 0x05, 0xf6, 0x3d, 0xd9, 0xbe,
	0x72, 0xed, 0xd3, 0x37, 0x8f, 0x0d, 0xe7, 0x58, 0x2d, 0x0f, 0x60, 0x34, 0x61, 0x41, 0x25, 0x6d,
	0xb1, 0x56, 0x50, 0xc5, 0x5e, 0xe2, 0xa5, 0x41, 0xfe, 0xa5, 0xa3, 0x39, 0x0b, 0x4d, 0xa7, 0x60,
	0x65, 0x81, 0x6c, 0xfc, 0x2b, 0xf1, 0xd3, 0xe1, 0xec, 0x82, 0x9f, 0xb8, 0x9e, 0xe7, 0x9d, 0x82,
	0x25, 0x50, 0x1e, 0x18, 0x57, 0xd8, 0xe8, 0x9a, 0xfd, 0x2b, 0x3b, 0x4b, 0xd8, 0xac, 0xa4, 0x26,
	0x30, 0x5b, 0xd0, 0x24, 0x51, 0xaf, 0x1a, 0xb0, 0xb6, 0x10, 0x10, 0xfb, 0x89, 0x97, 0x86, 0xf9,
	0xd8, 0x21, 0xb7, 0x47, 0xc4, 0xdc, 0x01, 0xd3, 0x2b, 0x36, 0x38, 0x0c, 0x8d, 0xc6, 0x2c, 0x70,
	0x97, 0x91, 0xe9, 0xaf, 0x0e, 0xf3, 0x41, 0xbf, 0x86, 0x4c, 0x34, 0x62, 0xbf, 0x6b, 0xa9, 0xc0,
	0xc6, 0x7e, 0xe2, 0xa7, 0x61, 0xee, 0xc4, 0xf4, 0xdd, 0x63, 0x7f, 0x8f, 0x6c, 0xdf, 0x83, 0xc9,
	0xb1, 0x23, 0xf8, 0x41, 0xf6, 0x6f, 0x16, 0xaf, 0x1f, 0xff, 0xbd, 0x87, 0xbb, 0xd3, 0x12, 0xd5,
	0x3e, 0x89, 0xef, 0x53, 0xb5, 0xfe, 0xd3, 0x67, 0xe2, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x40, 0xc9, 0xca, 0xa7, 0x02, 0x00, 0x00,
}

func (this *ModSecurity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModSecurity)
	if !ok {
		that2, ok := that.(ModSecurity)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if this.CustomInterventionMessage != that1.CustomInterventionMessage {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RuleSet)
	if !ok {
		that2, ok := that.(RuleSet)
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
	if this.RuleStr != that1.RuleStr {
		return false
	}
	if len(this.Files) != len(that1.Files) {
		return false
	}
	for i := range this.Files {
		if this.Files[i] != that1.Files[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ModSecurityPerRoute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ModSecurityPerRoute)
	if !ok {
		that2, ok := that.(ModSecurityPerRoute)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if this.CustomInterventionMessage != that1.CustomInterventionMessage {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
