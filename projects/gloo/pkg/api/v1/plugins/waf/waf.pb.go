// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	waf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/waf"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Settings struct {
	// disable waf on this listener
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Add owasp core rule set
	// if nil will not be added
	CoreRuleSet *CoreRuleSet `protobuf:"bytes,2,opt,name=core_rule_set,json=coreRuleSet,proto3" json:"core_rule_set,omitempty"`
	// custom rule sets rules to add
	RuleSets             []*waf.RuleSet `protobuf:"bytes,3,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b92862d396c16f, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Settings) GetCoreRuleSet() *CoreRuleSet {
	if m != nil {
		return m.CoreRuleSet
	}
	return nil
}

func (m *Settings) GetRuleSets() []*waf.RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

type CoreRuleSet struct {
	// Optional custom settings for the OWASP core rule set.
	// For an example on the configuration options see: https://github.com/SpiderLabs/owasp-modsecurity-crs/blob/v3.2/dev/crs-setup.conf.example
	// The same rules apply to these options as do to the `RuleSet`s. The file option is better if possible.
	//
	// Types that are valid to be assigned to CustomSettingsType:
	//	*CoreRuleSet_CustomSettingsString
	//	*CoreRuleSet_CustomSettingsFile
	CustomSettingsType   isCoreRuleSet_CustomSettingsType `protobuf_oneof:"CustomSettingsType"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CoreRuleSet) Reset()         { *m = CoreRuleSet{} }
func (m *CoreRuleSet) String() string { return proto.CompactTextString(m) }
func (*CoreRuleSet) ProtoMessage()    {}
func (*CoreRuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b92862d396c16f, []int{1}
}
func (m *CoreRuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreRuleSet.Unmarshal(m, b)
}
func (m *CoreRuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreRuleSet.Marshal(b, m, deterministic)
}
func (m *CoreRuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreRuleSet.Merge(m, src)
}
func (m *CoreRuleSet) XXX_Size() int {
	return xxx_messageInfo_CoreRuleSet.Size(m)
}
func (m *CoreRuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreRuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_CoreRuleSet proto.InternalMessageInfo

type isCoreRuleSet_CustomSettingsType interface {
	isCoreRuleSet_CustomSettingsType()
	Equal(interface{}) bool
}

type CoreRuleSet_CustomSettingsString struct {
	CustomSettingsString string `protobuf:"bytes,2,opt,name=custom_settings_string,json=customSettingsString,proto3,oneof"`
}
type CoreRuleSet_CustomSettingsFile struct {
	CustomSettingsFile string `protobuf:"bytes,3,opt,name=custom_settings_file,json=customSettingsFile,proto3,oneof"`
}

func (*CoreRuleSet_CustomSettingsString) isCoreRuleSet_CustomSettingsType() {}
func (*CoreRuleSet_CustomSettingsFile) isCoreRuleSet_CustomSettingsType()   {}

func (m *CoreRuleSet) GetCustomSettingsType() isCoreRuleSet_CustomSettingsType {
	if m != nil {
		return m.CustomSettingsType
	}
	return nil
}

func (m *CoreRuleSet) GetCustomSettingsString() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsString); ok {
		return x.CustomSettingsString
	}
	return ""
}

func (m *CoreRuleSet) GetCustomSettingsFile() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsFile); ok {
		return x.CustomSettingsFile
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CoreRuleSet) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CoreRuleSet_OneofMarshaler, _CoreRuleSet_OneofUnmarshaler, _CoreRuleSet_OneofSizer, []interface{}{
		(*CoreRuleSet_CustomSettingsString)(nil),
		(*CoreRuleSet_CustomSettingsFile)(nil),
	}
}

func _CoreRuleSet_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CoreRuleSet)
	// CustomSettingsType
	switch x := m.CustomSettingsType.(type) {
	case *CoreRuleSet_CustomSettingsString:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.CustomSettingsString)
	case *CoreRuleSet_CustomSettingsFile:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.CustomSettingsFile)
	case nil:
	default:
		return fmt.Errorf("CoreRuleSet.CustomSettingsType has unexpected type %T", x)
	}
	return nil
}

func _CoreRuleSet_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CoreRuleSet)
	switch tag {
	case 2: // CustomSettingsType.custom_settings_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.CustomSettingsType = &CoreRuleSet_CustomSettingsString{x}
		return true, err
	case 3: // CustomSettingsType.custom_settings_file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.CustomSettingsType = &CoreRuleSet_CustomSettingsFile{x}
		return true, err
	default:
		return false, nil
	}
}

func _CoreRuleSet_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CoreRuleSet)
	// CustomSettingsType
	switch x := m.CustomSettingsType.(type) {
	case *CoreRuleSet_CustomSettingsString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CustomSettingsString)))
		n += len(x.CustomSettingsString)
	case *CoreRuleSet_CustomSettingsFile:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CustomSettingsFile)))
		n += len(x.CustomSettingsFile)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type VhostSettings struct {
	// disable waf on this virtual host
	Disabled             bool      `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Settings             *Settings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VhostSettings) Reset()         { *m = VhostSettings{} }
func (m *VhostSettings) String() string { return proto.CompactTextString(m) }
func (*VhostSettings) ProtoMessage()    {}
func (*VhostSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b92862d396c16f, []int{2}
}
func (m *VhostSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VhostSettings.Unmarshal(m, b)
}
func (m *VhostSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VhostSettings.Marshal(b, m, deterministic)
}
func (m *VhostSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VhostSettings.Merge(m, src)
}
func (m *VhostSettings) XXX_Size() int {
	return xxx_messageInfo_VhostSettings.Size(m)
}
func (m *VhostSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_VhostSettings.DiscardUnknown(m)
}

var xxx_messageInfo_VhostSettings proto.InternalMessageInfo

func (m *VhostSettings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *VhostSettings) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type RouteSettings struct {
	// disable waf on this route
	Disabled             bool      `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Settings             *Settings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RouteSettings) Reset()         { *m = RouteSettings{} }
func (m *RouteSettings) String() string { return proto.CompactTextString(m) }
func (*RouteSettings) ProtoMessage()    {}
func (*RouteSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b92862d396c16f, []int{3}
}
func (m *RouteSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSettings.Unmarshal(m, b)
}
func (m *RouteSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSettings.Marshal(b, m, deterministic)
}
func (m *RouteSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSettings.Merge(m, src)
}
func (m *RouteSettings) XXX_Size() int {
	return xxx_messageInfo_RouteSettings.Size(m)
}
func (m *RouteSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSettings proto.InternalMessageInfo

func (m *RouteSettings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *RouteSettings) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*Settings)(nil), "waf.plugins.gloo.solo.io.Settings")
	proto.RegisterType((*CoreRuleSet)(nil), "waf.plugins.gloo.solo.io.CoreRuleSet")
	proto.RegisterType((*VhostSettings)(nil), "waf.plugins.gloo.solo.io.VhostSettings")
	proto.RegisterType((*RouteSettings)(nil), "waf.plugins.gloo.solo.io.RouteSettings")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/waf/waf.proto", fileDescriptor_91b92862d396c16f)
}

var fileDescriptor_91b92862d396c16f = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xc1, 0x0a, 0xd3, 0x40,
	0x10, 0x35, 0x16, 0x24, 0xdd, 0xd0, 0xcb, 0x12, 0x24, 0xf4, 0x20, 0x25, 0x20, 0xf4, 0xe2, 0xae,
	0x46, 0xf0, 0x24, 0x1e, 0x52, 0x28, 0x7a, 0xf0, 0x92, 0x8a, 0x07, 0x2f, 0x25, 0x4d, 0x27, 0xdb,
	0xb5, 0xdb, 0x4c, 0xd8, 0x9d, 0x54, 0xfb, 0x15, 0xfe, 0x86, 0x9f, 0xe1, 0xb7, 0xf8, 0x25, 0x92,
	0xa4, 0xa9, 0x56, 0x2c, 0xf4, 0xe4, 0x61, 0xe1, 0xcd, 0xcc, 0xbe, 0xf7, 0x66, 0x86, 0x61, 0xa9,
	0xd2, 0xb4, 0x6b, 0x36, 0xa2, 0xc0, 0x83, 0x74, 0x68, 0xf0, 0x99, 0x46, 0xa9, 0x0c, 0xa2, 0xac,
	0x2d, 0x7e, 0x86, 0x82, 0x5c, 0x1f, 0xe5, 0xb5, 0x96, 0xc7, 0x17, 0xb2, 0x36, 0x8d, 0xd2, 0x95,
	0x93, 0x5f, 0xf2, 0xb2, 0x7d, 0xa2, 0xb6, 0x48, 0xc8, 0xa3, 0x0e, 0xf6, 0x25, 0xd1, 0x7e, 0x17,
	0xad, 0x92, 0xd0, 0x38, 0x5d, 0xde, 0xaf, 0x0e, 0x5f, 0x09, 0x6c, 0x95, 0x1b, 0x09, 0xd5, 0x11,
	0x4f, 0xd7, 0x0e, 0xd3, 0x50, 0xa1, 0xc2, 0x0e, 0xca, 0x16, 0xf5, 0xd9, 0xf8, 0x87, 0xc7, 0xfc,
	0x15, 0x10, 0xe9, 0x4a, 0x39, 0x3e, 0x65, 0xfe, 0x56, 0xbb, 0x7c, 0x63, 0x60, 0x1b, 0x79, 0x33,
	0x6f, 0xee, 0x67, 0x97, 0x98, 0xbf, 0x63, 0x93, 0x02, 0x2d, 0xac, 0x6d, 0x63, 0x60, 0xed, 0x80,
	0xa2, 0x87, 0x33, 0x6f, 0x1e, 0x24, 0x4f, 0xc5, 0xad, 0xc6, 0xc5, 0x02, 0x2d, 0x64, 0x8d, 0x81,
	0x15, 0x50, 0x16, 0x14, 0xbf, 0x03, 0xfe, 0x9e, 0x8d, 0x07, 0x15, 0x17, 0x8d, 0x66, 0xa3, 0x79,
	0x90, 0x3c, 0x17, 0x5d, 0xcb, 0xa2, 0xc0, 0xaa, 0xd4, 0x4a, 0x94, 0xda, 0x10, 0x58, 0xb1, 0x23,
	0xaa, 0xc5, 0x01, 0xb7, 0x0e, 0x8a, 0xc6, 0x6a, 0x3a, 0x89, 0x63, 0x22, 0x06, 0x45, 0xdf, 0xf6,
	0xc0, 0xc5, 0xdf, 0x3c, 0x16, 0xfc, 0xe1, 0xc5, 0x5f, 0xb1, 0xc7, 0x45, 0xe3, 0x08, 0x0f, 0xad,
	0x41, 0x37, 0xd8, 0xda, 0x91, 0xd5, 0x95, 0xea, 0x5a, 0x1e, 0xbf, 0x7d, 0x90, 0x85, 0x7d, 0x7d,
	0x98, 0x7b, 0xd5, 0x55, 0x79, 0xc2, 0xc2, 0xbf, 0x79, 0xa5, 0x36, 0x10, 0x8d, 0xce, 0x2c, 0x7e,
	0xcd, 0x5a, 0x6a, 0x03, 0x69, 0xc8, 0xf8, 0xe2, 0x2a, 0xfb, 0xe1, 0x54, 0x43, 0xbc, 0x67, 0x93,
	0x8f, 0x3b, 0x74, 0x74, 0xd7, 0x62, 0xdf, 0x30, 0x7f, 0xf0, 0x3b, 0xef, 0x34, 0xbe, 0xbd, 0xd3,
	0x41, 0x31, 0xbb, 0x70, 0x5a, 0xb3, 0x0c, 0x1b, 0x82, 0xff, 0x61, 0x96, 0xa6, 0xdf, 0x7f, 0x3e,
	0xf1, 0x3e, 0xbd, 0xbe, 0xef, 0x24, 0xeb, 0xbd, 0xfa, 0xc7, 0xd1, 0x6f, 0x1e, 0x75, 0x97, 0xf7,
	0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xa0, 0x79, 0x0f, 0x37, 0x03, 0x00, 0x00,
}

func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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
	if !this.CoreRuleSet.Equal(that1.CoreRuleSet) {
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet)
	if !ok {
		that2, ok := that.(CoreRuleSet)
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
	if that1.CustomSettingsType == nil {
		if this.CustomSettingsType != nil {
			return false
		}
	} else if this.CustomSettingsType == nil {
		return false
	} else if !this.CustomSettingsType.Equal(that1.CustomSettingsType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsString) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsString)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsString)
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
	if this.CustomSettingsString != that1.CustomSettingsString {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsFile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsFile)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsFile)
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
	if this.CustomSettingsFile != that1.CustomSettingsFile {
		return false
	}
	return true
}
func (this *VhostSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VhostSettings)
	if !ok {
		that2, ok := that.(VhostSettings)
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
	if !this.Settings.Equal(that1.Settings) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteSettings)
	if !ok {
		that2, ok := that.(RouteSettings)
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
	if !this.Settings.Equal(that1.Settings) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
