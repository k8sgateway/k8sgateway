// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/dlp/dlp.proto

package dlp

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	matchers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/core/matchers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_type "github.com/solo-io/solo-kit/pkg/api/external/envoy/type"
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

//
//The following pre-made action types map to the following regex matchers:
//
//SSN:
//- '(?!\D)[0-9]{9}(?=\D|$)'
//- '(?!\D)[0-9]{3}\-[0-9]{2}\-[0-9]{4}(?=\D|$)'
//- '(?!\D)[0-9]{3}\ [0-9]{2}\ [0-9]{4}(?=\D|$)'
//
//MASTERCARD:
//- '(?!\D)5[1-5][0-9]{2}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//VISA:
//- '(?!\D)4[0-9]{3}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//AMEX:
//- '(?!\D)(34|37)[0-9]{2}(\ |\-|)[0-9]{6}(\ |\-|)[0-9]{5}(?=\D|$)'
//
//DISCOVER:
//- '(?!\D)6011(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//JCB:
//- '(?!\D)3[0-9]{3}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(\ |\-|)[0-9]{4}(?=\D|$)'
//- '(?!\D)(2131|1800)[0-9]{11}(?=\D|$)'
//
//DINERS_CLUB:
//- '(?!\D)30[0-5][0-9](\ |\-|)[0-9]{6}(\ |\-|)[0-9]{4}(?=\D|$)'
//- '(?!\D)(36|38)[0-9]{2}(\ |\-|)[0-9]{6}(\ |\-|)[0-9]{4}(?=\D|$)'
//
//CREDIT_CARD_TRACKERS:
//- '[1-9][0-9]{2}\-[0-9]{2}\-[0-9]{4}\^\d'
//- '(?!\D)\%?[Bb]\d{13,19}\^[\-\/\.\w\s]{2,26}\^[0-9][0-9][01][0-9][0-9]{3}'
//- '(?!\D)\;\d{13,19}\=(\d{3}|)(\d{4}|\=)'
//
//ALL_CREDIT_CARDS:
//- (All credit card related regexes from above)
//
type Action_ActionType int32

const (
	Action_CUSTOM               Action_ActionType = 0
	Action_SSN                  Action_ActionType = 1
	Action_MASTERCARD           Action_ActionType = 2
	Action_VISA                 Action_ActionType = 3
	Action_AMEX                 Action_ActionType = 4
	Action_DISCOVER             Action_ActionType = 5
	Action_JCB                  Action_ActionType = 6
	Action_DINERS_CLUB          Action_ActionType = 7
	Action_CREDIT_CARD_TRACKERS Action_ActionType = 8
	Action_ALL_CREDIT_CARDS     Action_ActionType = 9
)

var Action_ActionType_name = map[int32]string{
	0: "CUSTOM",
	1: "SSN",
	2: "MASTERCARD",
	3: "VISA",
	4: "AMEX",
	5: "DISCOVER",
	6: "JCB",
	7: "DINERS_CLUB",
	8: "CREDIT_CARD_TRACKERS",
	9: "ALL_CREDIT_CARDS",
}

var Action_ActionType_value = map[string]int32{
	"CUSTOM":               0,
	"SSN":                  1,
	"MASTERCARD":           2,
	"VISA":                 3,
	"AMEX":                 4,
	"DISCOVER":             5,
	"JCB":                  6,
	"DINERS_CLUB":          7,
	"CREDIT_CARD_TRACKERS": 8,
	"ALL_CREDIT_CARDS":     9,
}

func (x Action_ActionType) String() string {
	return proto.EnumName(Action_ActionType_name, int32(x))
}

func (Action_ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08adca4fdc7c089d, []int{3, 0}
}

// Listener level config for dlp filter
type FilterConfig struct {
	// The list of transformation, matcher pairs.
	// The first rule which matches will be applied.
	DlpRules             []*DlpRule `protobuf:"bytes,1,rep,name=dlp_rules,json=dlpRules,proto3" json:"dlp_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_08adca4fdc7c089d, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDlpRules() []*DlpRule {
	if m != nil {
		return m.DlpRules
	}
	return nil
}

// Rule which applies a given set of actions to a matching route.
// The route matching functions exactly the same as the envoy routes in the virtual host.
type DlpRule struct {
	// Matcher by which to determine if the given transformation should be applied
	// if omitted, will it match all (i.e., default to / prefix matcher)
	Matcher *matchers.Matcher `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// List of data loss prevention actions to be applied.
	// These actions will be applied in order, one at a time.
	Actions              []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DlpRule) Reset()         { *m = DlpRule{} }
func (m *DlpRule) String() string { return proto.CompactTextString(m) }
func (*DlpRule) ProtoMessage()    {}
func (*DlpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_08adca4fdc7c089d, []int{1}
}
func (m *DlpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DlpRule.Unmarshal(m, b)
}
func (m *DlpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DlpRule.Marshal(b, m, deterministic)
}
func (m *DlpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DlpRule.Merge(m, src)
}
func (m *DlpRule) XXX_Size() int {
	return xxx_messageInfo_DlpRule.Size(m)
}
func (m *DlpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_DlpRule.DiscardUnknown(m)
}

var xxx_messageInfo_DlpRule proto.InternalMessageInfo

func (m *DlpRule) GetMatcher() *matchers.Matcher {
	if m != nil {
		return m.Matcher
	}
	return nil
}

func (m *DlpRule) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

//
//Route/Vhost level config for dlp filter
//
//If a config is present on the route or vhost level it will completely overwrite the
//listener level config.
type Config struct {
	// List of data loss prevention actions to be applied.
	// These actions will be applied in order, one at a time.
	Actions              []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_08adca4fdc7c089d, []int{2}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

//
//A single action meant to mask sensitive data.
//The action type represents a set of pre configured actions,
//as well as the ability to create custom actions.
//These actions can also be shadowed, a shadowed action will be recorded
//in the statistics, and debug logs, but not actually committed in the response body.
//
//To use a pre-made action simply set the action type to anything other than `CUSTOM`
//
//``` yaml
//actionType: VISA
//```
//
//To create a custom action set the custom action field. The default enum value
//is custom, so that can be left empty.
//
//``` yaml
//customAction:
//name: test
//regex:
//- "hello"
//- "world"
//maskChar: Y
//percent: 60
//```
//
//
type Action struct {
	// The action type to implement.
	ActionType Action_ActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=dlp.options.gloo.solo.io.Action_ActionType" json:"action_type,omitempty"`
	// The custom user action to be applied.
	// This field will only be used if the custom action type is specified above.
	CustomAction *CustomAction `protobuf:"bytes,2,opt,name=custom_action,json=customAction,proto3" json:"custom_action,omitempty"`
	// Shadow represents whether the action should be taken, or just recorded.
	Shadow               bool     `protobuf:"varint,3,opt,name=shadow,proto3" json:"shadow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_08adca4fdc7c089d, []int{3}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetActionType() Action_ActionType {
	if m != nil {
		return m.ActionType
	}
	return Action_CUSTOM
}

func (m *Action) GetCustomAction() *CustomAction {
	if m != nil {
		return m.CustomAction
	}
	return nil
}

func (m *Action) GetShadow() bool {
	if m != nil {
		return m.Shadow
	}
	return false
}

//
//A user defined custom action to carry out on the response body.
//
//The list of regex strings are applied in order. So for instance, if there is a response body with the content:
//`hello world`
//
//And there is a custom action
//``` yaml
//customAction:
//name: test
//regex:
//- "hello"
//- "world"
//maskChar: Y
//percent: 60
//```
//
//the result would be:
//`YYYlo YYYld`
//
//If the mask_char, and percent were left to default, the result would be:
//`XXXXo XXXXd`
//
type CustomAction struct {
	// The name of the custom action.
	// This name is used for logging and debugging purposes.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of regex strings which will be applied in order.
	Regex []string `protobuf:"bytes,2,rep,name=regex,proto3" json:"regex,omitempty"`
	// The masking character for the sensitive data.
	// default value: X
	MaskChar string `protobuf:"bytes,3,opt,name=mask_char,json=maskChar,proto3" json:"mask_char,omitempty"`
	// The percent of the string which will be masked by the mask_char
	// default value: 75%
	// rounds ratio (percent/100) by std::round http://www.cplusplus.com/reference/cmath/round/
	Percent              *_type.Percent `protobuf:"bytes,4,opt,name=percent,proto3" json:"percent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CustomAction) Reset()         { *m = CustomAction{} }
func (m *CustomAction) String() string { return proto.CompactTextString(m) }
func (*CustomAction) ProtoMessage()    {}
func (*CustomAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_08adca4fdc7c089d, []int{4}
}
func (m *CustomAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomAction.Unmarshal(m, b)
}
func (m *CustomAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomAction.Marshal(b, m, deterministic)
}
func (m *CustomAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomAction.Merge(m, src)
}
func (m *CustomAction) XXX_Size() int {
	return xxx_messageInfo_CustomAction.Size(m)
}
func (m *CustomAction) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomAction.DiscardUnknown(m)
}

var xxx_messageInfo_CustomAction proto.InternalMessageInfo

func (m *CustomAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomAction) GetRegex() []string {
	if m != nil {
		return m.Regex
	}
	return nil
}

func (m *CustomAction) GetMaskChar() string {
	if m != nil {
		return m.MaskChar
	}
	return ""
}

func (m *CustomAction) GetPercent() *_type.Percent {
	if m != nil {
		return m.Percent
	}
	return nil
}

func init() {
	proto.RegisterEnum("dlp.options.gloo.solo.io.Action_ActionType", Action_ActionType_name, Action_ActionType_value)
	proto.RegisterType((*FilterConfig)(nil), "dlp.options.gloo.solo.io.FilterConfig")
	proto.RegisterType((*DlpRule)(nil), "dlp.options.gloo.solo.io.DlpRule")
	proto.RegisterType((*Config)(nil), "dlp.options.gloo.solo.io.Config")
	proto.RegisterType((*Action)(nil), "dlp.options.gloo.solo.io.Action")
	proto.RegisterType((*CustomAction)(nil), "dlp.options.gloo.solo.io.CustomAction")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/dlp/dlp.proto", fileDescriptor_08adca4fdc7c089d)
}

var fileDescriptor_08adca4fdc7c089d = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xea, 0x9f, 0x49, 0x28, 0xd6, 0x12, 0x21, 0xab, 0x48, 0x28, 0x04, 0x09, 0x55,
	0x42, 0xb5, 0x45, 0xb9, 0x21, 0x81, 0xe4, 0xd8, 0x06, 0x85, 0x26, 0x6d, 0xb5, 0x4e, 0x2b, 0xc4,
	0xc5, 0x72, 0x9d, 0xc5, 0x31, 0x75, 0xbc, 0xab, 0xb5, 0x53, 0xda, 0x3b, 0xaf, 0xc1, 0x9d, 0x77,
	0xe2, 0xc6, 0x53, 0x70, 0x44, 0xbb, 0x76, 0x4b, 0x38, 0x04, 0xe8, 0x21, 0xca, 0x37, 0x7f, 0xdf,
	0xcc, 0x37, 0xd9, 0x09, 0x4c, 0xd3, 0xac, 0x5a, 0xac, 0xce, 0xec, 0x84, 0x2e, 0x9d, 0x92, 0xe6,
	0x74, 0x2f, 0xa3, 0x4e, 0x9a, 0x53, 0xea, 0x30, 0x4e, 0x3f, 0x91, 0xa4, 0x2a, 0x6b, 0x2b, 0x66,
	0x99, 0x73, 0xf1, 0xdc, 0x21, 0x45, 0x45, 0x38, 0xe3, 0x59, 0x49, 0x1c, 0xca, 0xaa, 0x8c, 0x16,
	0xa5, 0x33, 0xcf, 0x99, 0xf8, 0xd8, 0x8c, 0xd3, 0x8a, 0x22, 0x4b, 0xc0, 0x26, 0x64, 0x8b, 0x4a,
	0x5b, 0x90, 0xda, 0x19, 0xdd, 0x19, 0xdf, 0xaa, 0x51, 0x42, 0x39, 0x71, 0x96, 0x71, 0x95, 0x2c,
	0x08, 0x2f, 0x6f, 0x40, 0xdd, 0x64, 0xc7, 0x22, 0xc5, 0x05, 0xbd, 0x72, 0xaa, 0x2b, 0x46, 0x1c,
	0x46, 0x78, 0x42, 0x8a, 0xaa, 0x89, 0xf4, 0x53, 0x9a, 0x52, 0x09, 0x1d, 0x81, 0x1a, 0x2f, 0x22,
	0x97, 0x55, 0xed, 0x24, 0x97, 0x4d, 0xe6, 0xf0, 0x10, 0x7a, 0x6f, 0xb2, 0xbc, 0x22, 0xdc, 0xa3,
	0xc5, 0xc7, 0x2c, 0x45, 0xaf, 0xc1, 0x98, 0xe7, 0x2c, 0xe2, 0xab, 0x9c, 0x94, 0x96, 0x32, 0x68,
	0xef, 0x76, 0xf7, 0x1f, 0xdb, 0x9b, 0xc4, 0xd8, 0x7e, 0xce, 0xf0, 0x2a, 0x27, 0x58, 0x9f, 0xd7,
	0xa0, 0x1c, 0x7e, 0x51, 0x40, 0x6b, 0xbc, 0xe8, 0x15, 0x68, 0xcd, 0xc4, 0x96, 0x32, 0x50, 0x76,
	0xbb, 0xfb, 0x4f, 0xec, 0x1b, 0x05, 0x42, 0xd8, 0x9f, 0x5c, 0xd3, 0x3a, 0x84, 0xaf, 0x6b, 0xd0,
	0x4b, 0xd0, 0xe2, 0x44, 0x36, 0xb5, 0x5a, 0x72, 0x90, 0xc1, 0xe6, 0x41, 0x5c, 0x99, 0x88, 0xaf,
	0x0b, 0x86, 0x3e, 0xa8, 0x8d, 0xa0, 0x35, 0x16, 0xe5, 0xb6, 0x2c, 0xdf, 0x5b, 0xa0, 0xd6, 0x3e,
	0x34, 0x81, 0x6e, 0xed, 0x8d, 0xc4, 0xba, 0xa5, 0x9e, 0xed, 0xfd, 0x67, 0xff, 0xa2, 0x6a, 0xbe,
	0x66, 0x57, 0x8c, 0x60, 0x88, 0x6f, 0x30, 0x3a, 0x80, 0xbb, 0xc9, 0xaa, 0xac, 0xe8, 0x32, 0xaa,
	0x9d, 0x56, 0x4b, 0xee, 0xe7, 0xe9, 0x66, 0x3e, 0x4f, 0xa6, 0x37, 0x03, 0xf6, 0x92, 0x35, 0x0b,
	0x3d, 0x00, 0xb5, 0x5c, 0xc4, 0x73, 0xfa, 0xd9, 0x6a, 0x0f, 0x94, 0x5d, 0x1d, 0x37, 0xd6, 0xf0,
	0xab, 0x02, 0xf0, 0xbb, 0x3f, 0x02, 0x50, 0xbd, 0x93, 0x70, 0x76, 0x34, 0x35, 0xef, 0x20, 0x0d,
	0xda, 0x61, 0x78, 0x68, 0x2a, 0x68, 0x1b, 0x60, 0xea, 0x86, 0xb3, 0x00, 0x7b, 0x2e, 0xf6, 0xcd,
	0x16, 0xd2, 0xa1, 0x73, 0x3a, 0x0e, 0x5d, 0xb3, 0x2d, 0x90, 0x3b, 0x0d, 0xde, 0x9b, 0x1d, 0xd4,
	0x03, 0xdd, 0x1f, 0x87, 0xde, 0xd1, 0x69, 0x80, 0xcd, 0x2d, 0x51, 0xfa, 0xce, 0x1b, 0x99, 0x2a,
	0xba, 0x07, 0x5d, 0x7f, 0x7c, 0x18, 0xe0, 0x30, 0xf2, 0x26, 0x27, 0x23, 0x53, 0x43, 0x16, 0xf4,
	0x3d, 0x1c, 0xf8, 0xe3, 0x59, 0x24, 0xc8, 0xa2, 0x19, 0x76, 0xbd, 0x83, 0x00, 0x87, 0xa6, 0x8e,
	0xfa, 0x60, 0xba, 0x93, 0x49, 0xb4, 0x16, 0x0d, 0x4d, 0x43, 0x3c, 0x95, 0xde, 0xba, 0x2c, 0x84,
	0xa0, 0x53, 0xc4, 0xcb, 0x7a, 0xb9, 0x06, 0x96, 0x18, 0xf5, 0x61, 0x8b, 0x93, 0x94, 0x5c, 0xca,
	0x27, 0x60, 0xe0, 0xda, 0x40, 0x0f, 0xc1, 0x58, 0xc6, 0xe5, 0x79, 0x94, 0x2c, 0x62, 0x2e, 0x55,
	0x1b, 0x58, 0x17, 0x0e, 0x6f, 0x11, 0x73, 0xb4, 0x07, 0x5a, 0x73, 0x0d, 0x56, 0x47, 0xae, 0xf5,
	0xbe, 0x2d, 0x0f, 0xc5, 0x16, 0xbf, 0x9c, 0x7d, 0x5c, 0x87, 0xf0, 0x75, 0xce, 0xe8, 0xf8, 0xe7,
	0x48, 0xf9, 0xf6, 0xe3, 0x91, 0xf2, 0xe1, 0xed, 0xff, 0x5d, 0x26, 0x3b, 0x4f, 0xff, 0xfe, 0x37,
	0x70, 0xa6, 0xca, 0xd3, 0x7a, 0xf1, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x29, 0x98, 0xd7, 0x54,
	0x04, 0x00, 0x00,
}

func (this *FilterConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FilterConfig)
	if !ok {
		that2, ok := that.(FilterConfig)
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
	if len(this.DlpRules) != len(that1.DlpRules) {
		return false
	}
	for i := range this.DlpRules {
		if !this.DlpRules[i].Equal(that1.DlpRules[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DlpRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DlpRule)
	if !ok {
		that2, ok := that.(DlpRule)
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
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	if len(this.Actions) != len(that1.Actions) {
		return false
	}
	for i := range this.Actions {
		if !this.Actions[i].Equal(that1.Actions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Config) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
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
	if len(this.Actions) != len(that1.Actions) {
		return false
	}
	for i := range this.Actions {
		if !this.Actions[i].Equal(that1.Actions[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Action) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Action)
	if !ok {
		that2, ok := that.(Action)
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
	if this.ActionType != that1.ActionType {
		return false
	}
	if !this.CustomAction.Equal(that1.CustomAction) {
		return false
	}
	if this.Shadow != that1.Shadow {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CustomAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CustomAction)
	if !ok {
		that2, ok := that.(CustomAction)
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
	if len(this.Regex) != len(that1.Regex) {
		return false
	}
	for i := range this.Regex {
		if this.Regex[i] != that1.Regex[i] {
			return false
		}
	}
	if this.MaskChar != that1.MaskChar {
		return false
	}
	if !this.Percent.Equal(that1.Percent) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
