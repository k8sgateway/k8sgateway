// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/upgrade/upgrade.proto

package upgrade

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type UpgradeConfig struct {
	// Types that are valid to be assigned to UpgradeType:
	//	*UpgradeConfig_Websocket
	UpgradeType          isUpgradeConfig_UpgradeType `protobuf_oneof:"upgrade_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UpgradeConfig) Reset()         { *m = UpgradeConfig{} }
func (m *UpgradeConfig) String() string { return proto.CompactTextString(m) }
func (*UpgradeConfig) ProtoMessage()    {}
func (*UpgradeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c46a83ce4167c77, []int{0}
}
func (m *UpgradeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeConfig.Unmarshal(m, b)
}
func (m *UpgradeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeConfig.Marshal(b, m, deterministic)
}
func (m *UpgradeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeConfig.Merge(m, src)
}
func (m *UpgradeConfig) XXX_Size() int {
	return xxx_messageInfo_UpgradeConfig.Size(m)
}
func (m *UpgradeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeConfig proto.InternalMessageInfo

type isUpgradeConfig_UpgradeType interface {
	isUpgradeConfig_UpgradeType()
	Equal(interface{}) bool
}

type UpgradeConfig_Websocket struct {
	Websocket *UpgradeConfig_UpgradeSpec `protobuf:"bytes,1,opt,name=websocket,proto3,oneof" json:"websocket,omitempty"`
}

func (*UpgradeConfig_Websocket) isUpgradeConfig_UpgradeType() {}

func (m *UpgradeConfig) GetUpgradeType() isUpgradeConfig_UpgradeType {
	if m != nil {
		return m.UpgradeType
	}
	return nil
}

func (m *UpgradeConfig) GetWebsocket() *UpgradeConfig_UpgradeSpec {
	if x, ok := m.GetUpgradeType().(*UpgradeConfig_Websocket); ok {
		return x.Websocket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UpgradeConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UpgradeConfig_Websocket)(nil),
	}
}

type UpgradeConfig_UpgradeSpec struct {
	// Whether the upgrade should be enabled. `true` by default.
	Enabled              *types.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpgradeConfig_UpgradeSpec) Reset()         { *m = UpgradeConfig_UpgradeSpec{} }
func (m *UpgradeConfig_UpgradeSpec) String() string { return proto.CompactTextString(m) }
func (*UpgradeConfig_UpgradeSpec) ProtoMessage()    {}
func (*UpgradeConfig_UpgradeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c46a83ce4167c77, []int{0, 0}
}
func (m *UpgradeConfig_UpgradeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeConfig_UpgradeSpec.Unmarshal(m, b)
}
func (m *UpgradeConfig_UpgradeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeConfig_UpgradeSpec.Marshal(b, m, deterministic)
}
func (m *UpgradeConfig_UpgradeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeConfig_UpgradeSpec.Merge(m, src)
}
func (m *UpgradeConfig_UpgradeSpec) XXX_Size() int {
	return xxx_messageInfo_UpgradeConfig_UpgradeSpec.Size(m)
}
func (m *UpgradeConfig_UpgradeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeConfig_UpgradeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeConfig_UpgradeSpec proto.InternalMessageInfo

func (m *UpgradeConfig_UpgradeSpec) GetEnabled() *types.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func init() {
	proto.RegisterType((*UpgradeConfig)(nil), "upgrade.options.gloo.solo.io.UpgradeConfig")
	proto.RegisterType((*UpgradeConfig_UpgradeSpec)(nil), "upgrade.options.gloo.solo.io.UpgradeConfig.UpgradeSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/upgrade/upgrade.proto", fileDescriptor_1c46a83ce4167c77)
}

var fileDescriptor_1c46a83ce4167c77 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0xb5, 0x1b, 0xc5, 0x8c, 0xba, 0x28, 0x2e, 0xa4, 0xc8, 0x20, 0xae, 0xdc, 0x78, 0x83, 0x0f,
	0x70, 0x2b, 0x1d, 0x10, 0x71, 0x39, 0xa2, 0x82, 0x1b, 0x69, 0x3b, 0x77, 0x62, 0x9c, 0xd8, 0x73,
	0x69, 0x53, 0x07, 0xff, 0xc8, 0x9f, 0xf0, 0x67, 0xfc, 0x12, 0xe9, 0x23, 0xa8, 0x0b, 0xc5, 0x55,
	0xee, 0x49, 0xce, 0x23, 0xf7, 0xa8, 0x2b, 0x63, 0xfd, 0x63, 0x93, 0x53, 0x81, 0x67, 0x5d, 0xc3,
	0xe1, 0xd0, 0x42, 0x1b, 0x07, 0x68, 0xa9, 0xf0, 0xc4, 0x85, 0xaf, 0x7b, 0x94, 0x89, 0xd5, 0x2f,
	0x47, 0x1a, 0xe2, 0x2d, 0xca, 0x5a, 0x37, 0x62, 0xaa, 0x6c, 0xc6, 0xe1, 0x24, 0xa9, 0xe0, 0x11,
	0xef, 0x06, 0x38, 0xd0, 0xa8, 0x95, 0x52, 0xeb, 0x4a, 0x16, 0xc9, 0xd8, 0x00, 0xc6, 0xb1, 0xee,
	0xb8, 0x79, 0x33, 0xd7, 0xcb, 0x2a, 0x13, 0xe1, 0xaa, 0xee, 0xd5, 0xc9, 0xb6, 0x81, 0x41, 0x37,
	0xea, 0x76, 0xea, 0x6f, 0xf7, 0xdf, 0x23, 0xb5, 0x79, 0xd3, 0xdb, 0x4e, 0x50, 0xce, 0xad, 0x89,
	0xef, 0xd4, 0xfa, 0x92, 0xf3, 0x1a, 0xc5, 0x82, 0xfd, 0x4e, 0xb4, 0x17, 0x1d, 0x8c, 0x8e, 0xcf,
	0xe8, 0xaf, 0x64, 0xfa, 0xa1, 0x0f, 0xe8, 0x5a, 0xb8, 0xb8, 0x5c, 0x99, 0x7e, 0x79, 0x25, 0x13,
	0x35, 0xfa, 0xf6, 0x16, 0x9f, 0xaa, 0x35, 0x2e, 0xb3, 0xdc, 0xf1, 0x6c, 0x48, 0x49, 0xa8, 0xdf,
	0x80, 0xc2, 0x06, 0x94, 0x02, 0xee, 0x36, 0x73, 0x0d, 0x4f, 0x03, 0x35, 0xdd, 0x52, 0x1b, 0xc3,
	0x5f, 0x1e, 0xfc, 0xab, 0x70, 0x7a, 0xf1, 0xf6, 0x31, 0x8e, 0xee, 0xcf, 0xff, 0xd7, 0xb2, 0x2c,
	0xcc, 0x2f, 0x4d, 0xe7, 0xab, 0x5d, 0xe8, 0xc9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x3a,
	0x94, 0xbb, 0xb0, 0x01, 0x00, 0x00,
}

func (this *UpgradeConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeConfig)
	if !ok {
		that2, ok := that.(UpgradeConfig)
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
	if that1.UpgradeType == nil {
		if this.UpgradeType != nil {
			return false
		}
	} else if this.UpgradeType == nil {
		return false
	} else if !this.UpgradeType.Equal(that1.UpgradeType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpgradeConfig_Websocket) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeConfig_Websocket)
	if !ok {
		that2, ok := that.(UpgradeConfig_Websocket)
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
	if !this.Websocket.Equal(that1.Websocket) {
		return false
	}
	return true
}
func (this *UpgradeConfig_UpgradeSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpgradeConfig_UpgradeSpec)
	if !ok {
		that2, ok := that.(UpgradeConfig_UpgradeSpec)
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
	if !this.Enabled.Equal(that1.Enabled) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
