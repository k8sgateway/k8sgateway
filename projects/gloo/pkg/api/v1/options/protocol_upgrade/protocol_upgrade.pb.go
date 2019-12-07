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
func (m *ProtocolUpgradeConfig) String() string { return proto.CompactTextString(m) }
func (*ProtocolUpgradeConfig) ProtoMessage()    {}
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
func (m *ProtocolUpgradeConfig_ProtocolUpgradeSpec) String() string { return proto.CompactTextString(m) }
func (*ProtocolUpgradeConfig_ProtocolUpgradeSpec) ProtoMessage()    {}
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
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xcd, 0x45, 0x71, 0x15, 0x0f, 0x55, 0x41, 0x72, 0x28, 0x22, 0x08, 0x5e, 0x9c, 0xc5,
	0x3f, 0x4f, 0x10, 0x2f, 0x8a, 0x07, 0x4b, 0x45, 0x0f, 0x22, 0x94, 0x24, 0x9d, 0xae, 0x6b, 0xd7,
	0x7c, 0x43, 0xb2, 0xb1, 0xf8, 0x46, 0x3e, 0x97, 0x07, 0x9f, 0x43, 0x92, 0x6c, 0x11, 0x6c, 0xc1,
	0xde, 0x66, 0x7f, 0x3b, 0xfc, 0xe6, 0x1b, 0x46, 0x3d, 0x1b, 0xeb, 0x5f, 0xea, 0x8c, 0x72, 0xbc,
	0xe9, 0x0a, 0x0e, 0xa7, 0x16, 0xda, 0x38, 0x40, 0x4b, 0x89, 0x57, 0xce, 0x7d, 0xd5, 0xbd, 0x52,
	0xb1, 0xfa, 0xfd, 0x4c, 0x43, 0xbc, 0x45, 0x51, 0x35, 0x9f, 0x1e, 0x39, 0xdc, 0xa8, 0x16, 0x53,
	0xa6, 0x63, 0x5e, 0x00, 0xd4, 0x82, 0xde, 0xf1, 0x02, 0x0f, 0x06, 0x6a, 0xac, 0xd4, 0x0c, 0x24,
	0x8b, 0xb8, 0x6f, 0x00, 0xe3, 0x82, 0x25, 0xab, 0x27, 0x7a, 0x56, 0xa6, 0x22, 0x5c, 0x56, 0x9d,
	0x26, 0xde, 0x33, 0x30, 0x68, 0x4b, 0xdd, 0x54, 0x1d, 0x3d, 0xfa, 0x8e, 0xd4, 0xfe, 0x20, 0xf8,
	0x1f, 0x3a, 0xfd, 0x15, 0x8a, 0x89, 0x35, 0x3d, 0x51, 0x9b, 0x33, 0xce, 0x2a, 0xe4, 0x53, 0xf6,
	0x07, 0xd1, 0x61, 0x74, 0xb2, 0x75, 0x3e, 0xa0, 0x95, 0xa2, 0xd0, 0x52, 0xe1, 0x5f, 0x7a, 0x2f,
	0x9c, 0x5f, 0xaf, 0x0d, 0x7f, 0x87, 0xc4, 0xb7, 0x6a, 0x77, 0x49, 0x4f, 0xef, 0x52, 0x6d, 0x70,
	0x91, 0x66, 0x8e, 0xc7, 0x21, 0x46, 0x4c, 0xdd, 0xaa, 0x34, 0x5f, 0x95, 0x12, 0xc0, 0x3d, 0xa6,
	0xae, 0xe6, 0xe1, 0xbc, 0x35, 0xd9, 0x51, 0xdb, 0x21, 0xe3, 0xc8, 0x7f, 0x08, 0x27, 0x77, 0x9f,
	0x5f, 0xfd, 0xe8, 0xe9, 0x66, 0xb5, 0x4b, 0xc9, 0xd4, 0xfc, 0x77, 0xad, 0x6c, 0xbd, 0x25, 0x17,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xa6, 0x3e, 0xd3, 0xfd, 0x01, 0x00, 0x00,
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
