//// Code generated by protoc-gen-gogo. DO NOT EDIT.
//// source: github.com/solo-io/gloo/projects/gloo/api/v1/extensions.proto
//
package v1
//
//import (
//	bytes "bytes"
//	fmt "fmt"
//	math "math"
//
//	_ "github.com/gogo/protobuf/gogoproto"
//	proto "github.com/gogo/protobuf/proto"
//	types "github.com/gogo/protobuf/types"
//)
//
//// Reference imports to suppress errors if they are not otherwise used.
//var _ = proto.Marshal
//var _ = fmt.Errorf
//var _ = math.Inf
//
//// This is a compile-time assertion to ensure that this generated file
//// is compatible with the proto package it is being compiled against.
//// A compilation error at this line likely means your copy of the
//// proto package needs to be updated.
//const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package
//
//type Extensions struct {
//	Configs              map[string]*types.Struct `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
//	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
//	XXX_unrecognized     []byte                   `json:"-"`
//	XXX_sizecache        int32                    `json:"-"`
//}
//
//func (m *Extensions) Reset()         { *m = Extensions{} }
//func (m *Extensions) String() string { return proto.CompactTextString(m) }
//func (*Extensions) ProtoMessage()    {}
//func (*Extensions) Descriptor() ([]byte, []int) {
//	return fileDescriptor_eb6aa7a8f802feeb, []int{0}
//}
//func (m *Extensions) XXX_Unmarshal(b []byte) error {
//	return xxx_messageInfo_Extensions.Unmarshal(m, b)
//}
//func (m *Extensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
//	return xxx_messageInfo_Extensions.Marshal(b, m, deterministic)
//}
//func (m *Extensions) XXX_Merge(src proto.Message) {
//	xxx_messageInfo_Extensions.Merge(m, src)
//}
//func (m *Extensions) XXX_Size() int {
//	return xxx_messageInfo_Extensions.Size(m)
//}
//func (m *Extensions) XXX_DiscardUnknown() {
//	xxx_messageInfo_Extensions.DiscardUnknown(m)
//}
//
//var xxx_messageInfo_Extensions proto.InternalMessageInfo
//
//func (m *Extensions) GetConfigs() map[string]*types.Struct {
//	if m != nil {
//		return m.Configs
//	}
//	return nil
//}
//
//type Extension struct {
//	Config               *types.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
//	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
//	XXX_unrecognized     []byte        `json:"-"`
//	XXX_sizecache        int32         `json:"-"`
//}
//
//func (m *Extension) Reset()         { *m = Extension{} }
//func (m *Extension) String() string { return proto.CompactTextString(m) }
//func (*Extension) ProtoMessage()    {}
//func (*Extension) Descriptor() ([]byte, []int) {
//	return fileDescriptor_eb6aa7a8f802feeb, []int{1}
//}
//func (m *Extension) XXX_Unmarshal(b []byte) error {
//	return xxx_messageInfo_Extension.Unmarshal(m, b)
//}
//func (m *Extension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
//	return xxx_messageInfo_Extension.Marshal(b, m, deterministic)
//}
//func (m *Extension) XXX_Merge(src proto.Message) {
//	xxx_messageInfo_Extension.Merge(m, src)
//}
//func (m *Extension) XXX_Size() int {
//	return xxx_messageInfo_Extension.Size(m)
//}
//func (m *Extension) XXX_DiscardUnknown() {
//	xxx_messageInfo_Extension.DiscardUnknown(m)
//}
//
//var xxx_messageInfo_Extension proto.InternalMessageInfo
//
//func (m *Extension) GetConfig() *types.Struct {
//	if m != nil {
//		return m.Config
//	}
//	return nil
//}
//
//func init() {
//	proto.RegisterType((*Extensions)(nil), "gloo.solo.io.Extensions")
//	proto.RegisterMapType((map[string]*types.Struct)(nil), "gloo.solo.io.Extensions.ConfigsEntry")
//	proto.RegisterType((*Extension)(nil), "gloo.solo.io.Extension")
//}
//
//func init() {
//	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/extensions.proto", fileDescriptor_eb6aa7a8f802feeb)
//}
//
//var fileDescriptor_eb6aa7a8f802feeb = []byte{
//	// 268 bytes of a gzipped FileDescriptorProto
//	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
//	0x10, 0x86, 0xd9, 0x16, 0x2b, 0x9d, 0xf6, 0x20, 0x41, 0x30, 0x04, 0x91, 0x50, 0x10, 0x72, 0xe9,
//	0x8c, 0x56, 0x10, 0x11, 0x45, 0x50, 0xfa, 0x02, 0xe9, 0xcd, 0x5b, 0x13, 0xb6, 0xeb, 0xda, 0x98,
//	0x09, 0xd9, 0x4d, 0xb1, 0xaf, 0xe3, 0xc9, 0xe7, 0xf2, 0x49, 0x24, 0xbb, 0x6d, 0xed, 0x49, 0x7a,
//	0x9b, 0xdd, 0xff, 0x9b, 0xff, 0xff, 0x19, 0x78, 0x54, 0xda, 0xbe, 0x35, 0x19, 0xe6, 0xfc, 0x41,
//	0x86, 0x0b, 0x1e, 0x6b, 0x26, 0x55, 0x30, 0x53, 0x55, 0xf3, 0xbb, 0xcc, 0xad, 0xf1, 0xaf, 0x79,
//	0xa5, 0x69, 0x75, 0x4d, 0xf2, 0xd3, 0xca, 0xd2, 0x68, 0x2e, 0x0d, 0x56, 0x35, 0x5b, 0x0e, 0x86,
//	0xad, 0x8a, 0xed, 0x22, 0x6a, 0x8e, 0xce, 0x15, 0xb3, 0x2a, 0x24, 0x39, 0x2d, 0x6b, 0x16, 0x64,
//	0x6c, 0xdd, 0xe4, 0xd6, 0xb3, 0xd1, 0xa9, 0x62, 0xc5, 0x6e, 0xa4, 0x76, 0xf2, 0xbf, 0xa3, 0x2f,
//	0x01, 0x30, 0xdd, 0xd9, 0x06, 0x4f, 0x70, 0x9c, 0x73, 0xb9, 0xd0, 0xca, 0x84, 0x22, 0xee, 0x26,
//	0x83, 0xc9, 0x25, 0xee, 0x47, 0xe0, 0x1f, 0x8a, 0x2f, 0x9e, 0x9b, 0x96, 0xb6, 0x5e, 0xa7, 0xdb,
//	0xad, 0x68, 0x06, 0xc3, 0x7d, 0x21, 0x38, 0x81, 0xee, 0x52, 0xae, 0x43, 0x11, 0x8b, 0xa4, 0x9f,
//	0xb6, 0x63, 0x30, 0x86, 0xa3, 0xd5, 0xbc, 0x68, 0x64, 0xd8, 0x89, 0x45, 0x32, 0x98, 0x9c, 0xa1,
//	0x6f, 0x8d, 0xdb, 0xd6, 0x38, 0x73, 0xad, 0x53, 0x4f, 0xdd, 0x77, 0xee, 0xc4, 0xe8, 0x01, 0xfa,
//	0xbb, 0xe0, 0x80, 0xa0, 0xe7, 0xc3, 0x9c, 0xe9, 0x3f, 0x06, 0x1b, 0xec, 0xf9, 0xf6, 0xfb, 0xe7,
//	0x42, 0xbc, 0x5e, 0x1d, 0x76, 0xe9, 0x6a, 0xa9, 0x36, 0xd7, 0xce, 0x7a, 0xce, 0xf0, 0xe6, 0x37,
//	0x00, 0x00, 0xff, 0xff, 0xf2, 0xeb, 0x0f, 0x57, 0xa4, 0x01, 0x00, 0x00,
//}
//
//func (this *Extensions) Equal(that interface{}) bool {
//	if that == nil {
//		return this == nil
//	}
//
//	that1, ok := that.(*Extensions)
//	if !ok {
//		that2, ok := that.(Extensions)
//		if ok {
//			that1 = &that2
//		} else {
//			return false
//		}
//	}
//	if that1 == nil {
//		return this == nil
//	} else if this == nil {
//		return false
//	}
//	if len(this.Configs) != len(that1.Configs) {
//		return false
//	}
//	for i := range this.Configs {
//		if !this.Configs[i].Equal(that1.Configs[i]) {
//			return false
//		}
//	}
//	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
//		return false
//	}
//	return true
//}
//func (this *Extension) Equal(that interface{}) bool {
//	if that == nil {
//		return this == nil
//	}
//
//	that1, ok := that.(*Extension)
//	if !ok {
//		that2, ok := that.(Extension)
//		if ok {
//			that1 = &that2
//		} else {
//			return false
//		}
//	}
//	if that1 == nil {
//		return this == nil
//	} else if this == nil {
//		return false
//	}
//	if !this.Config.Equal(that1.Config) {
//		return false
//	}
//	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
//		return false
//	}
//	return true
//}
