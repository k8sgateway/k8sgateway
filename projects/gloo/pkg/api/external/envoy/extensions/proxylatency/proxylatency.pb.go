// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxylatency/proxylatency.proto

package proxylatency

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ProxyLatency_Start int32

const (
	ProxyLatency_FIRST_BYTE ProxyLatency_Start = 0
	ProxyLatency_LAST_BYTE  ProxyLatency_Start = 1
)

var ProxyLatency_Start_name = map[int32]string{
	0: "FIRST_BYTE",
	1: "LAST_BYTE",
}

var ProxyLatency_Start_value = map[string]int32{
	"FIRST_BYTE": 0,
	"LAST_BYTE":  1,
}

func (x ProxyLatency_Start) String() string {
	return proto.EnumName(ProxyLatency_Start_name, int32(x))
}

func (ProxyLatency_Start) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b72148e2a3523d34, []int{0, 0}
}

// Configure the proxy latency fitler. This filter measures the latency
// incurred by the filter chain in a histogram.
type ProxyLatency struct {
	// When to start measuring - the time of the last byte received, or the first one.
	Start ProxyLatency_Start `protobuf:"varint,1,opt,name=start,proto3,enum=envoy.config.filter.http.proxylatency.v2.ProxyLatency_Start" json:"start,omitempty"`
	// Charge a stat per upstream cluster.
	ChargeClusterStat bool `protobuf:"varint,2,opt,name=charge_cluster_stat,json=chargeClusterStat,proto3" json:"charge_cluster_stat,omitempty"`
	// Charge a stat per listener.
	ChargeListenerStat   bool     `protobuf:"varint,3,opt,name=charge_listener_stat,json=chargeListenerStat,proto3" json:"charge_listener_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyLatency) Reset()         { *m = ProxyLatency{} }
func (m *ProxyLatency) String() string { return proto.CompactTextString(m) }
func (*ProxyLatency) ProtoMessage()    {}
func (*ProxyLatency) Descriptor() ([]byte, []int) {
	return fileDescriptor_b72148e2a3523d34, []int{0}
}
func (m *ProxyLatency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyLatency.Unmarshal(m, b)
}
func (m *ProxyLatency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyLatency.Marshal(b, m, deterministic)
}
func (m *ProxyLatency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyLatency.Merge(m, src)
}
func (m *ProxyLatency) XXX_Size() int {
	return xxx_messageInfo_ProxyLatency.Size(m)
}
func (m *ProxyLatency) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyLatency.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyLatency proto.InternalMessageInfo

func (m *ProxyLatency) GetStart() ProxyLatency_Start {
	if m != nil {
		return m.Start
	}
	return ProxyLatency_FIRST_BYTE
}

func (m *ProxyLatency) GetChargeClusterStat() bool {
	if m != nil {
		return m.ChargeClusterStat
	}
	return false
}

func (m *ProxyLatency) GetChargeListenerStat() bool {
	if m != nil {
		return m.ChargeListenerStat
	}
	return false
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.proxylatency.v2.ProxyLatency_Start", ProxyLatency_Start_name, ProxyLatency_Start_value)
	proto.RegisterType((*ProxyLatency)(nil), "envoy.config.filter.http.proxylatency.v2.ProxyLatency")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxylatency/proxylatency.proto", fileDescriptor_b72148e2a3523d34)
}

var fileDescriptor_b72148e2a3523d34 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0x42, 0x31,
	0x14, 0x86, 0xad, 0x06, 0xa3, 0x8d, 0x12, 0xad, 0x0c, 0xc4, 0xc1, 0x10, 0x06, 0xc3, 0x62, 0x6b,
	0x70, 0x75, 0x11, 0xa3, 0x89, 0x09, 0x83, 0xb9, 0xb0, 0xc8, 0x42, 0x4a, 0x53, 0x4a, 0xb5, 0xf6,
	0xdc, 0xb4, 0x07, 0x02, 0x6f, 0xe4, 0x73, 0x39, 0xfb, 0x10, 0xe6, 0xb6, 0x90, 0xc8, 0x86, 0x5b,
	0xff, 0xfc, 0xdf, 0x77, 0x4e, 0xd2, 0x43, 0x95, 0xb1, 0x38, 0x9b, 0x4f, 0xb8, 0x82, 0x4f, 0x11,
	0xc1, 0xc1, 0x8d, 0x05, 0x61, 0x1c, 0x80, 0x28, 0x03, 0xbc, 0x6b, 0x85, 0x31, 0x27, 0x59, 0x5a,
	0xa1, 0x97, 0xa8, 0x83, 0x97, 0x4e, 0x68, 0xbf, 0x80, 0x55, 0x8a, 0x3e, 0x5a, 0xf0, 0xb1, 0x62,
	0x97, 0x2b, 0x27, 0x51, 0x7b, 0xb5, 0xda, 0x0a, 0xbc, 0x0c, 0x80, 0xc0, 0x3a, 0xc9, 0xe0, 0x0a,
	0xfc, 0xd4, 0x1a, 0x3e, 0xb5, 0x0e, 0x75, 0xe0, 0x33, 0xc4, 0x92, 0x6f, 0xc1, 0x8b, 0xee, 0x65,
	0xc3, 0x80, 0x81, 0x24, 0x89, 0xea, 0x95, 0xfd, 0xf6, 0x0f, 0xa1, 0x27, 0xaf, 0x15, 0xd9, 0xcf,
	0x24, 0x2b, 0x68, 0x2d, 0xa2, 0x0c, 0xd8, 0x24, 0x2d, 0xd2, 0xa9, 0x77, 0xef, 0xf9, 0xae, 0x0b,
	0xf8, 0xdf, 0x31, 0x7c, 0x50, 0xcd, 0x28, 0xf2, 0x28, 0xc6, 0xe9, 0x85, 0x9a, 0xc9, 0x60, 0xf4,
	0x58, 0xb9, 0x79, 0x44, 0x1d, 0xc6, 0x11, 0x25, 0x36, 0xf7, 0x5b, 0xa4, 0x73, 0x54, 0x9c, 0xe7,
	0xea, 0x31, 0x37, 0x03, 0x94, 0xc8, 0x6e, 0x69, 0x63, 0xcd, 0x3b, 0x1b, 0x51, 0xfb, 0x8d, 0x70,
	0x90, 0x04, 0x96, 0xbb, 0xfe, 0xba, 0xaa, 0x8c, 0xf6, 0x35, 0xad, 0xa5, 0x8d, 0xac, 0x4e, 0xe9,
	0xf3, 0x4b, 0x31, 0x18, 0x8e, 0x7b, 0x6f, 0xc3, 0xa7, 0xb3, 0x3d, 0x76, 0x4a, 0x8f, 0xfb, 0x0f,
	0x9b, 0x48, 0x7a, 0xa3, 0xaf, 0xef, 0x2b, 0x32, 0x1a, 0xee, 0x76, 0x99, 0xf2, 0xc3, 0xfc, 0xe3,
	0x3a, 0x93, 0xc3, 0xf4, 0xa3, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x83, 0x58, 0xf9,
	0xf8, 0x01, 0x00, 0x00,
}

func (this *ProxyLatency) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProxyLatency)
	if !ok {
		that2, ok := that.(ProxyLatency)
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
	if this.Start != that1.Start {
		return false
	}
	if this.ChargeClusterStat != that1.ChargeClusterStat {
		return false
	}
	if this.ChargeListenerStat != that1.ChargeListenerStat {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
