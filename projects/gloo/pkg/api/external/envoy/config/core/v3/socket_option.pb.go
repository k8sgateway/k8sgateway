// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/socket_option.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type SocketOption_SocketState int32

const (
	// Socket options are applied after socket creation but before binding the socket to a port
	SocketOption_STATE_PREBIND SocketOption_SocketState = 0
	// Socket options are applied after binding the socket to a port but before calling listen()
	SocketOption_STATE_BOUND SocketOption_SocketState = 1
	// Socket options are applied after calling listen()
	SocketOption_STATE_LISTENING SocketOption_SocketState = 2
)

var SocketOption_SocketState_name = map[int32]string{
	0: "STATE_PREBIND",
	1: "STATE_BOUND",
	2: "STATE_LISTENING",
}

var SocketOption_SocketState_value = map[string]int32{
	"STATE_PREBIND":   0,
	"STATE_BOUND":     1,
	"STATE_LISTENING": 2,
}

func (x SocketOption_SocketState) String() string {
	return proto.EnumName(SocketOption_SocketState_name, int32(x))
}

func (SocketOption_SocketState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff80f4c8bc804c53, []int{0, 0}
}

// Generic socket option message. This would be used to set socket options that
// might not exist in upstream kernels or precompiled Envoy binaries.
// [#next-free-field: 7]
type SocketOption struct {
	// An optional name to give this socket option for debugging, etc.
	// Uniqueness is not required and no special meaning is assumed.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Corresponding to the level value passed to setsockopt, such as IPPROTO_TCP
	Level int64 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	// The numeric name as passed to setsockopt
	Name int64 `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*SocketOption_IntValue
	//	*SocketOption_BufValue
	Value isSocketOption_Value `protobuf_oneof:"value"`
	// The state in which the option will be applied. When used in BindConfig
	// STATE_PREBIND is currently the only valid value.
	State                SocketOption_SocketState `protobuf:"varint,6,opt,name=state,proto3,enum=envoy.config.core.v3.SocketOption_SocketState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SocketOption) Reset()         { *m = SocketOption{} }
func (m *SocketOption) String() string { return proto.CompactTextString(m) }
func (*SocketOption) ProtoMessage()    {}
func (*SocketOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff80f4c8bc804c53, []int{0}
}
func (m *SocketOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketOption.Unmarshal(m, b)
}
func (m *SocketOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketOption.Marshal(b, m, deterministic)
}
func (m *SocketOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketOption.Merge(m, src)
}
func (m *SocketOption) XXX_Size() int {
	return xxx_messageInfo_SocketOption.Size(m)
}
func (m *SocketOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketOption.DiscardUnknown(m)
}

var xxx_messageInfo_SocketOption proto.InternalMessageInfo

type isSocketOption_Value interface {
	isSocketOption_Value()
	Equal(interface{}) bool
}

type SocketOption_IntValue struct {
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof" json:"int_value,omitempty"`
}
type SocketOption_BufValue struct {
	BufValue []byte `protobuf:"bytes,5,opt,name=buf_value,json=bufValue,proto3,oneof" json:"buf_value,omitempty"`
}

func (*SocketOption_IntValue) isSocketOption_Value() {}
func (*SocketOption_BufValue) isSocketOption_Value() {}

func (m *SocketOption) GetValue() isSocketOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SocketOption) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SocketOption) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *SocketOption) GetName() int64 {
	if m != nil {
		return m.Name
	}
	return 0
}

func (m *SocketOption) GetIntValue() int64 {
	if x, ok := m.GetValue().(*SocketOption_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *SocketOption) GetBufValue() []byte {
	if x, ok := m.GetValue().(*SocketOption_BufValue); ok {
		return x.BufValue
	}
	return nil
}

func (m *SocketOption) GetState() SocketOption_SocketState {
	if m != nil {
		return m.State
	}
	return SocketOption_STATE_PREBIND
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SocketOption) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SocketOption_IntValue)(nil),
		(*SocketOption_BufValue)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.core.v3.SocketOption_SocketState", SocketOption_SocketState_name, SocketOption_SocketState_value)
	proto.RegisterType((*SocketOption)(nil), "envoy.config.core.v3.SocketOption")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/socket_option.proto", fileDescriptor_ff80f4c8bc804c53)
}

var fileDescriptor_ff80f4c8bc804c53 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x18, 0x85, 0xeb, 0x99, 0x49, 0x45, 0x3d, 0xa5, 0x4d, 0xcd, 0x48, 0x44, 0x95, 0x40, 0xd1, 0xac,
	0xb2, 0xc1, 0x96, 0x98, 0x13, 0x10, 0x75, 0x80, 0x11, 0x28, 0x8d, 0x92, 0x81, 0x05, 0x2c, 0x46,
	0x49, 0xea, 0x09, 0xa6, 0x69, 0xfe, 0x28, 0x71, 0xa2, 0x74, 0xcb, 0x01, 0x38, 0x07, 0x47, 0xa8,
	0x58, 0x55, 0x1c, 0x85, 0x5d, 0x0f, 0xc0, 0x1e, 0xd9, 0x1e, 0xa4, 0x2c, 0x2a, 0xd1, 0xdd, 0xff,
	0xbe, 0xf7, 0x12, 0x3f, 0xd9, 0x3f, 0xfe, 0x9c, 0x0b, 0xf9, 0xa5, 0x4d, 0x69, 0x06, 0x57, 0xac,
	0x81, 0x02, 0x5e, 0x08, 0x60, 0x79, 0x01, 0xc0, 0xaa, 0x1a, 0xbe, 0xf2, 0x4c, 0x36, 0x46, 0x25,
	0x95, 0x60, 0xbc, 0x97, 0xbc, 0x2e, 0x93, 0x82, 0xf1, 0xb2, 0x83, 0x6b, 0x96, 0x41, 0xb9, 0x15,
	0x39, 0xcb, 0xa0, 0xe6, 0xac, 0x5b, 0xb0, 0x06, 0xb2, 0x4b, 0x2e, 0x37, 0x50, 0x49, 0x01, 0x25,
	0xad, 0x6a, 0x90, 0x40, 0x66, 0x3a, 0x49, 0x4d, 0x92, 0xaa, 0x24, 0xed, 0x16, 0xa7, 0x4f, 0xbb,
	0xa4, 0x10, 0x17, 0x89, 0xe4, 0xec, 0xdf, 0x60, 0xe2, 0xa7, 0xb3, 0x1c, 0x72, 0xd0, 0x23, 0x53,
	0xd3, 0x8e, 0x12, 0xde, 0x4b, 0x03, 0x79, 0x2f, 0x0d, 0x9b, 0xff, 0x1a, 0xe1, 0xc3, 0x58, 0x1f,
	0x78, 0xae, 0xcf, 0x23, 0x2e, 0x9e, 0x5e, 0xf0, 0x26, 0xab, 0x85, 0x96, 0x0e, 0x72, 0x91, 0x77,
	0x10, 0x0d, 0x11, 0x99, 0x61, 0xab, 0xe0, 0x1d, 0x2f, 0x9c, 0x91, 0x8b, 0xbc, 0x71, 0x64, 0x04,
	0x21, 0x78, 0x52, 0x26, 0x57, 0xdc, 0x19, 0x6b, 0xa8, 0x67, 0xf2, 0x0c, 0x1f, 0x88, 0x52, 0x6e,
	0xba, 0xa4, 0x68, 0xb9, 0x33, 0x51, 0xc6, 0xdb, 0xbd, 0xe8, 0x91, 0x28, 0xe5, 0x47, 0x45, 0x94,
	0x9d, 0xb6, 0xdb, 0x9d, 0x6d, 0xb9, 0xc8, 0x3b, 0x54, 0x76, 0xda, 0x6e, 0x8d, 0x1d, 0x62, 0xab,
	0x91, 0x89, 0xe4, 0xce, 0xbe, 0x8b, 0xbc, 0xa3, 0x97, 0x94, 0xde, 0x77, 0x07, 0x74, 0x58, 0x7e,
	0x27, 0x62, 0xf5, 0x95, 0x8f, 0x7f, 0xde, 0xdd, 0x8e, 0xad, 0x6f, 0x68, 0x64, 0xa3, 0xc8, 0xfc,
	0x68, 0xfe, 0x1a, 0x4f, 0x07, 0x09, 0x72, 0x82, 0x1f, 0xc7, 0xeb, 0x57, 0xeb, 0xe5, 0x26, 0x8c,
	0x96, 0xfe, 0x2a, 0x38, 0xb3, 0xf7, 0xc8, 0x31, 0x9e, 0x1a, 0xe4, 0x9f, 0x7f, 0x08, 0xce, 0x6c,
	0x44, 0x9e, 0xe0, 0x63, 0x03, 0xde, 0xaf, 0xe2, 0xf5, 0x32, 0x58, 0x05, 0x6f, 0xec, 0x91, 0x7f,
	0x84, 0x2d, 0x5d, 0x9a, 0x58, 0x37, 0x77, 0xb7, 0x63, 0xe4, 0x7f, 0x47, 0x37, 0x7f, 0x26, 0xe8,
	0xc7, 0xef, 0xe7, 0x08, 0xcf, 0x05, 0x98, 0x9e, 0x55, 0x0d, 0xfd, 0xf5, 0xbd, 0x95, 0xfd, 0x93,
	0x61, 0xe7, 0x50, 0x3d, 0x43, 0x88, 0x3e, 0xbd, 0x7b, 0xd8, 0xfa, 0x54, 0x97, 0xf9, 0xff, 0x57,
	0x28, 0xdd, 0xd7, 0x8f, 0xbb, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x70, 0x83, 0x2c, 0xdc, 0x94,
	0x02, 0x00, 0x00,
}

func (this *SocketOption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketOption)
	if !ok {
		that2, ok := that.(SocketOption)
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
	if this.Description != that1.Description {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if that1.Value == nil {
		if this.Value != nil {
			return false
		}
	} else if this.Value == nil {
		return false
	} else if !this.Value.Equal(that1.Value) {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SocketOption_IntValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketOption_IntValue)
	if !ok {
		that2, ok := that.(SocketOption_IntValue)
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
	if this.IntValue != that1.IntValue {
		return false
	}
	return true
}
func (this *SocketOption_BufValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SocketOption_BufValue)
	if !ok {
		that2, ok := that.(SocketOption_BufValue)
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
	if !bytes.Equal(this.BufValue, that1.BufValue) {
		return false
	}
	return true
}
