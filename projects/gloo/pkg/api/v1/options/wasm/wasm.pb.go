// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/wasm/wasm.proto

package wasm

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

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

type PluginSource_VmType int32

const (
	PluginSource_V8   PluginSource_VmType = 0
	PluginSource_WAVM PluginSource_VmType = 1
)

var PluginSource_VmType_name = map[int32]string{
	0: "V8",
	1: "WAVM",
}

var PluginSource_VmType_value = map[string]int32{
	"V8":   0,
	"WAVM": 1,
}

func (x PluginSource_VmType) String() string {
	return proto.EnumName(PluginSource_VmType_name, int32(x))
}

func (PluginSource_VmType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_31d0a37a89c26012, []int{0, 0}
}

type PluginSource struct {
	// name of image which houses the compiled wasm filter
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// string of the config sent to the wasm filter
	// Currently has to be json or will crash
	Config               string              `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	Name                 string              `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	RootId               string              `protobuf:"bytes,6,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	VmType               PluginSource_VmType `protobuf:"varint,7,opt,name=vm_type,json=vmType,proto3,enum=wasm.options.gloo.solo.io.PluginSource_VmType" json:"vm_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PluginSource) Reset()         { *m = PluginSource{} }
func (m *PluginSource) String() string { return proto.CompactTextString(m) }
func (*PluginSource) ProtoMessage()    {}
func (*PluginSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_31d0a37a89c26012, []int{0}
}
func (m *PluginSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginSource.Unmarshal(m, b)
}
func (m *PluginSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginSource.Marshal(b, m, deterministic)
}
func (m *PluginSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginSource.Merge(m, src)
}
func (m *PluginSource) XXX_Size() int {
	return xxx_messageInfo_PluginSource.Size(m)
}
func (m *PluginSource) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginSource.DiscardUnknown(m)
}

var xxx_messageInfo_PluginSource proto.InternalMessageInfo

func (m *PluginSource) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *PluginSource) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *PluginSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginSource) GetRootId() string {
	if m != nil {
		return m.RootId
	}
	return ""
}

func (m *PluginSource) GetVmType() PluginSource_VmType {
	if m != nil {
		return m.VmType
	}
	return PluginSource_V8
}

func init() {
	proto.RegisterEnum("wasm.options.gloo.solo.io.PluginSource_VmType", PluginSource_VmType_name, PluginSource_VmType_value)
	proto.RegisterType((*PluginSource)(nil), "wasm.options.gloo.solo.io.PluginSource")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/wasm/wasm.proto", fileDescriptor_31d0a37a89c26012)
}

var fileDescriptor_31d0a37a89c26012 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xcd, 0xec, 0x52, 0x0d, 0x22, 0x23, 0x0c, 0xad, 0x3d, 0xc8, 0xd8, 0x69, 0x17, 0x13,
	0xd4, 0x8b, 0x17, 0x0f, 0x8a, 0x30, 0x3c, 0x08, 0x32, 0xa5, 0x82, 0x97, 0xd1, 0x75, 0x31, 0x46,
	0x9b, 0xfe, 0x43, 0x9b, 0xd6, 0xed, 0x8d, 0x7c, 0x04, 0x9f, 0xc2, 0x87, 0xf0, 0x1d, 0xbc, 0x4b,
	0x93, 0x1d, 0x3c, 0x28, 0xec, 0x92, 0x7c, 0xdf, 0x2f, 0xff, 0x2f, 0x24, 0x1f, 0xb9, 0x92, 0xca,
	0x3e, 0xd7, 0x33, 0x96, 0x81, 0xe6, 0x15, 0xe4, 0x70, 0xa4, 0x80, 0xcb, 0x1c, 0x80, 0x9b, 0x12,
	0x5e, 0x44, 0x66, 0x2b, 0xef, 0x52, 0xa3, 0x78, 0x73, 0xcc, 0xc1, 0x58, 0x05, 0x45, 0xc5, 0xdf,
	0xd2, 0x4a, 0xbb, 0x85, 0x99, 0x12, 0x2c, 0xd0, 0x03, 0xa7, 0x57, 0xa7, 0xac, 0x4d, 0xb0, 0xf6,
	0x32, 0xa6, 0x20, 0xee, 0x4b, 0x90, 0xe0, 0xa6, 0x78, 0xab, 0x7c, 0x20, 0xa6, 0x62, 0x61, 0x3d,
	0x14, 0x0b, 0xeb, 0xd9, 0xf0, 0x13, 0x91, 0x9d, 0xdb, 0xbc, 0x96, 0xaa, 0xb8, 0x83, 0xba, 0xcc,
	0x04, 0xed, 0x93, 0xae, 0xd2, 0xa9, 0x14, 0x51, 0x67, 0x80, 0x46, 0xdb, 0x13, 0x6f, 0xe8, 0x1e,
	0xc1, 0x19, 0x14, 0x4f, 0x4a, 0x46, 0x9b, 0x0e, 0xaf, 0x1c, 0xa5, 0x24, 0x28, 0x52, 0x2d, 0xa2,
	0xae, 0xa3, 0x4e, 0xd3, 0x7d, 0x12, 0x96, 0x00, 0x76, 0xaa, 0xe6, 0x11, 0xf6, 0xc3, 0xad, 0xbd,
	0x9e, 0xd3, 0x31, 0x09, 0x1b, 0x3d, 0xb5, 0x4b, 0x23, 0xa2, 0x70, 0x80, 0x46, 0xbb, 0x27, 0x8c,
	0xfd, 0xfb, 0x05, 0xf6, 0xfb, 0x51, 0x2c, 0xd1, 0xf7, 0x4b, 0x23, 0x26, 0xb8, 0x71, 0xfb, 0x30,
	0x26, 0xd8, 0x13, 0x8a, 0x49, 0x27, 0x39, 0xeb, 0x6d, 0xd0, 0x2d, 0x12, 0x3c, 0x5c, 0x24, 0x37,
	0x3d, 0x74, 0x39, 0xfe, 0xf8, 0x0e, 0xd0, 0xfb, 0xd7, 0x21, 0x7a, 0x3c, 0x5f, 0xaf, 0x65, 0xf3,
	0x2a, 0xff, 0x6a, 0x7a, 0x86, 0x5d, 0x41, 0xa7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0xe6,
	0x85, 0x50, 0xad, 0x01, 0x00, 0x00,
}

func (this *PluginSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PluginSource)
	if !ok {
		that2, ok := that.(PluginSource)
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
	if this.Image != that1.Image {
		return false
	}
	if this.Config != that1.Config {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.RootId != that1.RootId {
		return false
	}
	if this.VmType != that1.VmType {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
