// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/v3/semantic_version.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
<<<<<<< HEAD
	_ "github.com/solo-io/protoc-gen-ext/extproto"
=======
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
>>>>>>> master
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

// Envoy uses SemVer (https://semver.org/). Major/minor versions indicate
// expected behaviors and APIs, the patch version field is used only
// for security fixes and can be generally ignored.
type SemanticVersion struct {
	MajorNumber          uint32   `protobuf:"varint,1,opt,name=major_number,json=majorNumber,proto3" json:"major_number,omitempty"`
	MinorNumber          uint32   `protobuf:"varint,2,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Patch                uint32   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SemanticVersion) Reset()         { *m = SemanticVersion{} }
func (m *SemanticVersion) String() string { return proto.CompactTextString(m) }
func (*SemanticVersion) ProtoMessage()    {}
func (*SemanticVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_522de03841b5881e, []int{0}
}
func (m *SemanticVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SemanticVersion.Unmarshal(m, b)
}
func (m *SemanticVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SemanticVersion.Marshal(b, m, deterministic)
}
func (m *SemanticVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SemanticVersion.Merge(m, src)
}
func (m *SemanticVersion) XXX_Size() int {
	return xxx_messageInfo_SemanticVersion.Size(m)
}
func (m *SemanticVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SemanticVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SemanticVersion proto.InternalMessageInfo

func (m *SemanticVersion) GetMajorNumber() uint32 {
	if m != nil {
		return m.MajorNumber
	}
	return 0
}

func (m *SemanticVersion) GetMinorNumber() uint32 {
	if m != nil {
		return m.MinorNumber
	}
	return 0
}

func (m *SemanticVersion) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func init() {
	proto.RegisterType((*SemanticVersion)(nil), "envoy.type.v3.SemanticVersion")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/type/v3/semantic_version.proto", fileDescriptor_522de03841b5881e)
}

var fileDescriptor_522de03841b5881e = []byte{
<<<<<<< HEAD
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0x87, 0x59, 0xff, 0x15, 0xd1, 0x43, 0x08, 0x29, 0x8e, 0x13, 0x44, 0xad, 0x6c, 0xdc, 0x29,
	0xf2, 0x06, 0x87, 0xb5, 0x1c, 0x0a, 0x0a, 0x36, 0xc7, 0x26, 0x0c, 0x7b, 0x7b, 0x66, 0x77, 0x96,
	0xcd, 0x5c, 0x48, 0x6a, 0x5f, 0xc6, 0x47, 0xf0, 0x79, 0x7c, 0x07, 0x7b, 0xc9, 0x6e, 0x40, 0xb4,
	0x10, 0xbb, 0x99, 0x6f, 0x3e, 0x86, 0xdf, 0x4c, 0xf6, 0xa4, 0x0d, 0x6f, 0x76, 0x95, 0xac, 0xc9,
	0x42, 0x4b, 0x0d, 0xdd, 0x18, 0x02, 0xdd, 0x10, 0x81, 0x0f, 0xb4, 0xc5, 0x9a, 0xdb, 0xd4, 0x29,
	0x6f, 0x00, 0x7b, 0xc6, 0xe0, 0x54, 0x03, 0xe8, 0x3a, 0x1a, 0x80, 0x07, 0x8f, 0xd0, 0x95, 0xd0,
	0xa2, 0x55, 0x8e, 0x4d, 0xbd, 0xee, 0x30, 0xb4, 0x86, 0x9c, 0xf4, 0x81, 0x98, 0xf2, 0x59, 0xb4,
	0xe4, 0x68, 0xc9, 0xae, 0x5c, 0x14, 0x9a, 0x34, 0xc5, 0x09, 0x8c, 0x55, 0x92, 0x16, 0x39, 0xf6,
	0x9c, 0x20, 0xf6, 0x9c, 0xd8, 0x95, 0xcd, 0x4e, 0x1f, 0xa6, 0x95, 0x8f, 0x69, 0x63, 0x7e, 0x99,
	0x9d, 0x58, 0xb5, 0xa5, 0xb0, 0x76, 0x3b, 0x5b, 0x61, 0x98, 0x8b, 0x0b, 0x71, 0x3d, 0xbb, 0x3f,
	0x8e, 0xec, 0x2e, 0xa2, 0xa8, 0x18, 0xf7, 0xad, 0xec, 0x4d, 0xca, 0xc8, 0x26, 0xa5, 0xc8, 0x0e,
	0xbd, 0xe2, 0x7a, 0x33, 0xdf, 0x8f, 0xb3, 0xd4, 0x2c, 0x5f, 0xc5, 0xfb, 0xe7, 0x81, 0x78, 0xfb,
	0x38, 0x17, 0xd9, 0x99, 0x21, 0x19, 0x53, 0xfb, 0x40, 0xfd, 0x20, 0x7f, 0x1c, 0xb0, 0x2c, 0x7e,
	0x85, 0x5a, 0x8d, 0x61, 0x57, 0xe2, 0xf9, 0xf6, 0x7f, 0x0f, 0xf4, 0x2f, 0xfa, 0x8f, 0x27, 0x56,
	0x47, 0xf1, 0xf6, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xad, 0x28, 0x22, 0xd7, 0x8f, 0x01, 0x00,
	0x00,
=======
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0xe5, 0x7e, 0xfa, 0x3a, 0x18, 0x2a, 0x50, 0xd5, 0xa1, 0x0a, 0x50, 0x11, 0x26, 0x16,
	0xec, 0x21, 0x1b, 0x63, 0xc5, 0x8c, 0x2a, 0x90, 0x40, 0x62, 0xa9, 0x9c, 0x60, 0xb9, 0x2e, 0x89,
	0xcf, 0xb2, 0x2f, 0x51, 0xb3, 0xb1, 0xc1, 0xc8, 0xcc, 0x13, 0xf0, 0x0c, 0xec, 0x48, 0xac, 0xbc,
	0x02, 0x4f, 0x82, 0xe2, 0x44, 0x02, 0x8a, 0x84, 0xd8, 0xee, 0xfe, 0xf7, 0x3b, 0xfb, 0xee, 0xfe,
	0xf4, 0x52, 0x69, 0x5c, 0x94, 0x29, 0xcb, 0xa0, 0xe0, 0x1e, 0x72, 0x38, 0xd2, 0xc0, 0x55, 0x0e,
	0xc0, 0xad, 0x83, 0xa5, 0xcc, 0xd0, 0xb7, 0x99, 0xb0, 0x9a, 0xcb, 0x15, 0x4a, 0x67, 0x44, 0xce,
	0xa5, 0xa9, 0xa0, 0xe6, 0x58, 0x5b, 0xc9, 0xab, 0x84, 0x7b, 0x59, 0x08, 0x83, 0x3a, 0x9b, 0x57,
	0xd2, 0x79, 0x0d, 0x86, 0x59, 0x07, 0x08, 0xc3, 0x41, 0xa0, 0x58, 0x43, 0xb1, 0x2a, 0x89, 0xf6,
	0xca, 0x6b, 0x2b, 0xb8, 0x30, 0x06, 0x50, 0xa0, 0x06, 0xe3, 0xb9, 0x47, 0x81, 0xa5, 0x6f, 0xe9,
	0x28, 0xfe, 0x51, 0xee, 0x5e, 0xd3, 0x46, 0x75, 0xc8, 0x48, 0x81, 0x82, 0x10, 0xf2, 0x26, 0x6a,
	0xd5, 0x83, 0x07, 0x42, 0xb7, 0xce, 0xbb, 0x09, 0x2e, 0xda, 0x96, 0x61, 0x4c, 0x37, 0x0b, 0xb1,
	0x04, 0x37, 0x37, 0x65, 0x91, 0x4a, 0x37, 0x26, 0xfb, 0xe4, 0x70, 0x70, 0xb6, 0x11, 0xb4, 0xd3,
	0x20, 0x05, 0x44, 0x9b, 0x4f, 0xa4, 0xd7, 0x21, 0x8d, 0xd6, 0x21, 0x23, 0xfa, 0xdf, 0x0a, 0xcc,
	0x16, 0xe3, 0x7f, 0xa1, 0xd6, 0x26, 0xc7, 0xf1, 0xe3, 0xcb, 0xfd, 0x64, 0x97, 0x46, 0x5f, 0xb6,
	0x5b, 0xfb, 0x7e, 0x7a, 0x47, 0x9e, 0xde, 0x27, 0xe4, 0xf9, 0xf6, 0xf5, 0xad, 0xdf, 0xdb, 0xee,
	0xd1, 0x1d, 0x0d, 0x2c, 0xd0, 0xd6, 0xc1, 0xaa, 0x66, 0xdf, 0xce, 0x32, 0x1d, 0xad, 0x35, 0xcf,
	0x9a, 0xa5, 0x66, 0xe4, 0xea, 0xe4, 0x6f, 0xb6, 0xd8, 0x1b, 0xf5, 0x8b, 0x35, 0x69, 0x3f, 0xdc,
	0x28, 0xf9, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x30, 0x3d, 0x1a, 0xe5, 0x01, 0x00, 0x00,
>>>>>>> master
}

func (this *SemanticVersion) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SemanticVersion)
	if !ok {
		that2, ok := that.(SemanticVersion)
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
	if this.MajorNumber != that1.MajorNumber {
		return false
	}
	if this.MinorNumber != that1.MinorNumber {
		return false
	}
	if this.Patch != that1.Patch {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
