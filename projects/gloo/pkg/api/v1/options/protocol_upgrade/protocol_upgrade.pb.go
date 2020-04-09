// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/protocol_upgrade/protocol_upgrade.proto

package protocol_upgrade

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type ProtocolUpgradeConfig struct {
	// Types that are valid to be assigned to UpgradeType:
	//	*ProtocolUpgradeConfig_Websocket
	UpgradeType          isProtocolUpgradeConfig_UpgradeType `protobuf_oneof:"upgrade_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ProtocolUpgradeConfig) Reset()         { *m = ProtocolUpgradeConfig{} }
func (m *ProtocolUpgradeConfig) String() string {
	return proto.CompactTextString(m)
}
func (*ProtocolUpgradeConfig) ProtoMessage() {}
func (*ProtocolUpgradeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_384550b21127c365, []int{0}
}
func (m *ProtocolUpgradeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolUpgradeConfig.Unmarshal(m, b)
}
func (m *ProtocolUpgradeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolUpgradeConfig.Marshal(b, m, deterministic)
}
func (m *ProtocolUpgradeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolUpgradeConfig.Merge(m, src)
}
func (m *ProtocolUpgradeConfig) XXX_Size() int {
	return xxx_messageInfo_ProtocolUpgradeConfig.Size(m)
}
func (m *ProtocolUpgradeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolUpgradeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolUpgradeConfig proto.InternalMessageInfo

type isProtocolUpgradeConfig_UpgradeType interface {
	isProtocolUpgradeConfig_UpgradeType()
	Equal(interface{}) bool
}

type ProtocolUpgradeConfig_Websocket struct {
	Websocket *ProtocolUpgradeConfig_ProtocolUpgradeSpec `protobuf:"bytes,1,opt,name=websocket,proto3,oneof" json:"websocket,omitempty"`
}

func (*ProtocolUpgradeConfig_Websocket) isProtocolUpgradeConfig_UpgradeType() {}

func (m *ProtocolUpgradeConfig) GetUpgradeType() isProtocolUpgradeConfig_UpgradeType {
	if m != nil {
		return m.UpgradeType
	}
	return nil
}

func (m *ProtocolUpgradeConfig) GetWebsocket() *ProtocolUpgradeConfig_ProtocolUpgradeSpec {
	if x, ok := m.GetUpgradeType().(*ProtocolUpgradeConfig_Websocket); ok {
		return x.Websocket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProtocolUpgradeConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProtocolUpgradeConfig_Websocket)(nil),
	}
}

type ProtocolUpgradeConfig_ProtocolUpgradeSpec struct {
	// Whether the upgrade should be enabled. If left unset, Envoy will enable the protocol upgrade.
	Enabled              *types.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) Reset() {
	*m = ProtocolUpgradeConfig_ProtocolUpgradeSpec{}
}
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) String() string {
	return proto.CompactTextString(m)
}
func (*ProtocolUpgradeConfig_ProtocolUpgradeSpec) ProtoMessage() {}
func (*ProtocolUpgradeConfig_ProtocolUpgradeSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_384550b21127c365, []int{0, 0}
}
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolUpgradeConfig_ProtocolUpgradeSpec.Unmarshal(m, b)
}
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolUpgradeConfig_ProtocolUpgradeSpec.Marshal(b, m, deterministic)
}
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolUpgradeConfig_ProtocolUpgradeSpec.Merge(m, src)
}
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) XXX_Size() int {
	return xxx_messageInfo_ProtocolUpgradeConfig_ProtocolUpgradeSpec.Size(m)
}
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolUpgradeConfig_ProtocolUpgradeSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolUpgradeConfig_ProtocolUpgradeSpec proto.InternalMessageInfo

func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) GetEnabled() *types.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtocolUpgradeConfig)(nil), "protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig")
	proto.RegisterType((*ProtocolUpgradeConfig_ProtocolUpgradeSpec)(nil), "protocol_upgrade.options.gloo.solo.io.ProtocolUpgradeConfig.ProtocolUpgradeSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/protocol_upgrade/protocol_upgrade.proto", fileDescriptor_384550b21127c365)
}

var fileDescriptor_384550b21127c365 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x5d, 0x10, 0xc5, 0x28, 0x1e, 0x56, 0x05, 0xd9, 0x43, 0x11, 0x41, 0xf0, 0x62, 0x82,
	0x7f, 0x9e, 0x60, 0xbd, 0x28, 0x5e, 0x4a, 0x8b, 0x1e, 0x44, 0x28, 0xbb, 0xdb, 0x69, 0x8c, 0x8d,
	0xfb, 0x85, 0x4d, 0xd6, 0xd6, 0x37, 0xf2, 0x11, 0x7c, 0x1e, 0x0f, 0xbe, 0x81, 0x77, 0xd9, 0x4d,
	0x8a, 0x60, 0x0b, 0xed, 0x6d, 0xf2, 0xcb, 0xf0, 0xfb, 0x66, 0x18, 0xf6, 0x24, 0x95, 0x7b, 0xae,
	0x73, 0x5e, 0xe0, 0x55, 0x58, 0x68, 0x9c, 0x29, 0x08, 0xa9, 0x01, 0x61, 0x2a, 0xbc, 0x50, 0xe1,
	0xac, 0x7f, 0x65, 0x46, 0x89, 0xb7, 0x73, 0x01, 0xe3, 0x14, 0x4a, 0xdb, 0x7c, 0x3a, 0x14, 0xd0,
	0x83, 0xda, 0xc8, 0x2a, 0x1b, 0xd2, 0x1c, 0xe0, 0x2d, 0x88, 0x4f, 0xe6, 0x78, 0x30, 0xf0, 0xc6,
	0xca, 0x9b, 0x40, 0xae, 0x90, 0x74, 0x24, 0x20, 0x75, 0xb0, 0xe4, 0xf5, 0x48, 0x4c, 0xaa, 0xcc,
	0x18, 0xaa, 0xac, 0xd7, 0x24, 0xfb, 0x12, 0x12, 0x6d, 0x29, 0x9a, 0x2a, 0xd0, 0x98, 0xa6, 0xce,
	0x43, 0x9a, 0x3a, 0xcf, 0x8e, 0xbf, 0x23, 0x76, 0xd0, 0x0d, 0x99, 0xf7, 0x3e, 0xf2, 0x1a, 0xe5,
	0x48, 0xc9, 0xd8, 0xb0, 0xad, 0x09, 0xe5, 0x16, 0xc5, 0x98, 0xdc, 0x61, 0x74, 0x14, 0x9d, 0x6e,
	0x5f, 0x74, 0xf9, 0x4a, 0xe3, 0xf1, 0x85, 0xc2, 0xff, 0xb4, 0x6f, 0xa8, 0xb8, 0x59, 0xeb, 0xfd,
	0x85, 0x24, 0x77, 0x6c, 0x6f, 0x41, 0x4f, 0x7c, 0xc5, 0x36, 0xa9, 0xcc, 0x72, 0x4d, 0xc3, 0x30,
	0x46, 0xc2, 0xfd, 0xfa, 0x7c, 0xb6, 0x3e, 0x4f, 0x01, 0xfd, 0x90, 0xe9, 0x9a, 0x7a, 0xb3, 0xd6,
	0x74, 0x97, 0xed, 0x84, 0x19, 0x07, 0xee, 0xdd, 0x50, 0xda, 0xff, 0xfc, 0x59, 0x8f, 0x3e, 0xbe,
	0x3a, 0xd1, 0xe3, 0xed, 0x6a, 0x17, 0x34, 0x63, 0xb9, 0xec, 0x8a, 0xf9, 0x46, 0x4b, 0x2e, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x09, 0xbb, 0x4b, 0x15, 0x02, 0x00, 0x00,
}

func (this *ProtocolUpgradeConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProtocolUpgradeConfig)
	if !ok {
		that2, ok := that.(ProtocolUpgradeConfig)
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
func (this *ProtocolUpgradeConfig_Websocket) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProtocolUpgradeConfig_Websocket)
	if !ok {
		that2, ok := that.(ProtocolUpgradeConfig_Websocket)
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
func (this *ProtocolUpgradeConfig_ProtocolUpgradeSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProtocolUpgradeConfig_ProtocolUpgradeSpec)
	if !ok {
		that2, ok := that.(ProtocolUpgradeConfig_ProtocolUpgradeSpec)
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
