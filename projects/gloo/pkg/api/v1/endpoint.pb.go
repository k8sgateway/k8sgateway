// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/endpoint.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//
//Endpoints represent dynamically discovered address/ports where an upstream service is listening
type Endpoint struct {
	// List of the upstreams the endpoint belongs to
	Upstreams []*core.ResourceRef `protobuf:"bytes,1,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
	// Address of the endpoint (ip or hostname)
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// listening port for the endpoint
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7969f9617648787, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetUpstreams() []*core.ResourceRef {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "gloo.solo.io.Endpoint")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/endpoint.proto", fileDescriptor_f7969f9617648787)
}

var fileDescriptor_f7969f9617648787 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xfd, 0xdc, 0x46, 0x5f, 0x1b, 0x07, 0x16, 0x0b, 0x41, 0xe8, 0x00, 0x11, 0x53, 0x06, 0xb0,
	0x69, 0x91, 0x00, 0xc1, 0x56, 0x89, 0x91, 0xc5, 0x23, 0x9b, 0x9b, 0xb8, 0xc1, 0xb4, 0xe9, 0xb5,
	0x6c, 0x87, 0x07, 0xe8, 0xd3, 0xf0, 0x20, 0x0c, 0x3c, 0x05, 0x03, 0x6f, 0xd0, 0x37, 0x40, 0xf9,
	0x31, 0x08, 0x89, 0xa1, 0x4c, 0xbe, 0x47, 0xe7, 0x1c, 0x1f, 0x9d, 0x7b, 0xf1, 0x6d, 0xa1, 0xdc,
	0x63, 0x35, 0xa3, 0x19, 0x94, 0xcc, 0xc2, 0x12, 0xce, 0x14, 0xb0, 0x62, 0x09, 0xc0, 0xb4, 0x81,
	0x27, 0x99, 0x39, 0xdb, 0x22, 0xa1, 0x15, 0x7b, 0x1e, 0x33, 0xb9, 0xca, 0x35, 0xa8, 0x95, 0xa3,
	0xda, 0x80, 0x03, 0xb2, 0x53, 0x73, 0xb4, 0xb6, 0x51, 0x05, 0xa3, 0xbd, 0x02, 0x0a, 0x68, 0x08,
	0x56, 0x4f, 0xad, 0x66, 0x34, 0xfe, 0x25, 0xa0, 0x79, 0x17, 0xca, 0xf9, 0x6f, 0x4b, 0xe9, 0x44,
	0x2e, 0x9c, 0xe8, 0x2c, 0xa7, 0x5b, 0x58, 0x8c, 0x9c, 0xff, 0x21, 0xc0, 0xe3, 0xd6, 0x72, 0xf2,
	0x8a, 0xf0, 0xf0, 0xae, 0xab, 0x42, 0xae, 0x70, 0x58, 0x69, 0xeb, 0x8c, 0x14, 0xa5, 0x8d, 0x51,
	0xd2, 0x4f, 0xa3, 0xc9, 0x21, 0xcd, 0xc0, 0x48, 0x5f, 0x8c, 0x72, 0x69, 0xa1, 0x32, 0x99, 0xe4,
	0x72, 0xce, 0xbf, 0xb5, 0x24, 0xc6, 0x03, 0x91, 0xe7, 0x46, 0x5a, 0x1b, 0xf7, 0x12, 0x94, 0x86,
	0xdc, 0x43, 0x42, 0x70, 0xa0, 0xc1, 0xb8, 0xb8, 0x9f, 0xa0, 0x74, 0x97, 0x37, 0x33, 0xb9, 0xc6,
	0x43, 0x5f, 0x33, 0x1e, 0x24, 0x28, 0x8d, 0x26, 0xfb, 0x3f, 0x53, 0xee, 0x3b, 0x76, 0x1a, 0xbc,
	0xbd, 0x1f, 0xff, 0xe3, 0x5f, 0xea, 0x9b, 0x83, 0xf5, 0x26, 0x08, 0x70, 0x4f, 0xea, 0xf5, 0x26,
	0x88, 0x48, 0xe8, 0x6f, 0x60, 0xa7, 0x97, 0x2f, 0x1f, 0x47, 0xe8, 0xe1, 0x7c, 0xbb, 0x0b, 0xea,
	0x45, 0xd1, 0x6d, 0x63, 0xf6, 0xbf, 0xd9, 0xc2, 0xc5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x04, 0xc5, 0x66, 0xfc, 0x01, 0x00, 0x00,
}

func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
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
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
