// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxylatency/proxylatency.proto

package proxylatency

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

// How to perform the latency measurement. Given an incoming request from downstream and
// outging request to upstream; or incoming response from upstream and outgoing repsonse to
// downstream, This outlines how to measure the latency used by the proxy.
type ProxyLatency_Measurement int32

const (
	// Count from the last byte of the incoming request\response to the first byte of the outgoing request\response.
	ProxyLatency_LAST_INCOMING_FIRST_OUTGOING ProxyLatency_Measurement = 0
	// Count from the first byte of the incoming request\response to the first byte of the outgoing request\response.
	ProxyLatency_FIRST_INCOMING_FIRST_OUTGOING ProxyLatency_Measurement = 1
	// Count from the last byte of the incoming request\response to the last byte of the outgoing request\response.
	ProxyLatency_LAST_INCOMING_LAST_OUTGOING ProxyLatency_Measurement = 2
	// Count from the first byte of the incoming request\response to the last byte of the outgoing request\response.
	ProxyLatency_FIRST_INCOMING_LAST_OUTGOING ProxyLatency_Measurement = 3
)

var ProxyLatency_Measurement_name = map[int32]string{
	0: "LAST_INCOMING_FIRST_OUTGOING",
	1: "FIRST_INCOMING_FIRST_OUTGOING",
	2: "LAST_INCOMING_LAST_OUTGOING",
	3: "FIRST_INCOMING_LAST_OUTGOING",
}

var ProxyLatency_Measurement_value = map[string]int32{
	"LAST_INCOMING_FIRST_OUTGOING":  0,
	"FIRST_INCOMING_FIRST_OUTGOING": 1,
	"LAST_INCOMING_LAST_OUTGOING":   2,
	"FIRST_INCOMING_LAST_OUTGOING":  3,
}

func (x ProxyLatency_Measurement) String() string {
	return proto.EnumName(ProxyLatency_Measurement_name, int32(x))
}

func (ProxyLatency_Measurement) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b72148e2a3523d34, []int{0, 0}
}

// Configure the proxy latency filter. This filter measures the latency
// incurred by the filter chain in a histogram.
type ProxyLatency struct {
	// How to measure the request.
	Request ProxyLatency_Measurement `protobuf:"varint,1,opt,name=request,proto3,enum=envoy.config.filter.http.proxylatency.v2.ProxyLatency_Measurement" json:"request,omitempty"`
	// When *_FIRST_OUTGOING is selected for a request, measure the time when decodeHeader for this filter is hit
	// When FIRST_OUTGOING (i.e. LAST_INCOMING_FIRST_OUTGOING or FIRST_INCOMING_FIRST_OUTGOING) is
	// instead of when the first byte is sent upstream. This has the advantage of not measuring the time
	// selected for request measurment, finish measuring proxy latency when decodeHeader for this
	// it takes a connection to form, which may skew the P99.
	// filter is hit instead of when the first byte is sent upstream. This has the advantage of not
	// for this to work the filter should be inserted last, just before the router filter.
	// measuring the time it takes a connection to form, which may skew the P99. For this to work
	// this filter should be inserted last, just before the router filter. This has no effect if
	// other measurement type is selected, and has no effect on how response is measured.
	MeasureRequestInternally bool `protobuf:"varint,5,opt,name=measure_request_internally,json=measureRequestInternally,proto3" json:"measure_request_internally,omitempty"`
	// How measure the response.
	Response ProxyLatency_Measurement `protobuf:"varint,2,opt,name=response,proto3,enum=envoy.config.filter.http.proxylatency.v2.ProxyLatency_Measurement" json:"response,omitempty"`
	// Charge a stat per upstream cluster. If not specified, defaults to true.
	ChargeClusterStat *types.BoolValue `protobuf:"bytes,3,opt,name=charge_cluster_stat,json=chargeClusterStat,proto3" json:"charge_cluster_stat,omitempty"`
	// Charge a stat per listener. If not specified, defaults to true.
	ChargeListenerStat   *types.BoolValue `protobuf:"bytes,4,opt,name=charge_listener_stat,json=chargeListenerStat,proto3" json:"charge_listener_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
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

func (m *ProxyLatency) GetRequest() ProxyLatency_Measurement {
	if m != nil {
		return m.Request
	}
	return ProxyLatency_LAST_INCOMING_FIRST_OUTGOING
}

func (m *ProxyLatency) GetMeasureRequestInternally() bool {
	if m != nil {
		return m.MeasureRequestInternally
	}
	return false
}

func (m *ProxyLatency) GetResponse() ProxyLatency_Measurement {
	if m != nil {
		return m.Response
	}
	return ProxyLatency_LAST_INCOMING_FIRST_OUTGOING
}

func (m *ProxyLatency) GetChargeClusterStat() *types.BoolValue {
	if m != nil {
		return m.ChargeClusterStat
	}
	return nil
}

func (m *ProxyLatency) GetChargeListenerStat() *types.BoolValue {
	if m != nil {
		return m.ChargeListenerStat
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.proxylatency.v2.ProxyLatency_Measurement", ProxyLatency_Measurement_name, ProxyLatency_Measurement_value)
	proto.RegisterType((*ProxyLatency)(nil), "envoy.config.filter.http.proxylatency.v2.ProxyLatency")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxylatency/proxylatency.proto", fileDescriptor_b72148e2a3523d34)
}

var fileDescriptor_b72148e2a3523d34 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x71, 0x5b, 0xa0, 0xda, 0x22, 0x04, 0x4b, 0x0f, 0x56, 0x80, 0x12, 0x7a, 0xca, 0x85,
	0xb5, 0x14, 0xae, 0x5c, 0x48, 0x25, 0x22, 0xa3, 0x34, 0x41, 0x4e, 0xe0, 0x50, 0x21, 0xac, 0x8d,
	0x35, 0xd9, 0x18, 0x36, 0x3b, 0xcb, 0xee, 0xb8, 0x34, 0x0f, 0xc2, 0x3b, 0xf0, 0x10, 0x3c, 0x0d,
	0x4f, 0x82, 0xec, 0x4d, 0xa2, 0x06, 0x09, 0x11, 0xa4, 0xde, 0x76, 0x7e, 0xff, 0xff, 0xe7, 0xd1,
	0x2f, 0x0d, 0x2b, 0x54, 0x49, 0xf3, 0x6a, 0x2a, 0x0a, 0x5c, 0x24, 0x1e, 0x35, 0xbe, 0x28, 0x31,
	0x51, 0x1a, 0x31, 0xb1, 0x0e, 0x3f, 0x43, 0x41, 0x3e, 0x4c, 0xd2, 0x96, 0x09, 0x5c, 0x11, 0x38,
	0x23, 0x75, 0x02, 0xe6, 0x12, 0x97, 0xcd, 0x68, 0x7c, 0x89, 0xc6, 0xd7, 0xde, 0xab, 0xa5, 0x96,
	0x04, 0xa6, 0x58, 0x6e, 0x0d, 0xc2, 0x3a, 0x24, 0xe4, 0x9d, 0x26, 0x21, 0x0a, 0x34, 0xb3, 0x52,
	0x89, 0x59, 0xa9, 0x09, 0x9c, 0x98, 0x13, 0x59, 0xb1, 0x65, 0xbe, 0xec, 0xb6, 0x4e, 0x14, 0xa2,
	0xd2, 0x90, 0x34, 0xb9, 0x69, 0x35, 0x4b, 0xbe, 0x39, 0x69, 0x2d, 0x38, 0x1f, 0x48, 0xad, 0x63,
	0x85, 0x0a, 0x9b, 0x67, 0x52, 0xbf, 0x82, 0x7a, 0xfa, 0xf3, 0x80, 0xdd, 0x7b, 0x57, 0x93, 0x06,
	0x81, 0xc4, 0x3f, 0xb2, 0xbb, 0x0e, 0xbe, 0x56, 0xe0, 0x29, 0x8e, 0xda, 0x51, 0xe7, 0x7e, 0xb7,
	0x27, 0x76, 0x5d, 0x41, 0x5c, 0x07, 0x89, 0x73, 0x90, 0xbe, 0x72, 0xb0, 0x00, 0x43, 0xd9, 0x1a,
	0xc9, 0x5f, 0xb1, 0xd6, 0x22, 0xe8, 0xf9, 0x4a, 0xca, 0x4b, 0x13, 0xba, 0xd1, 0xcb, 0xf8, 0x76,
	0x3b, 0xea, 0x1c, 0x66, 0xf1, 0xca, 0x91, 0x05, 0x43, 0xba, 0xf9, 0xce, 0x3f, 0xb1, 0x43, 0x07,
	0xde, 0xa2, 0xf1, 0x10, 0xef, 0xdd, 0xd8, 0x72, 0x1b, 0x26, 0x7f, 0xcb, 0x1e, 0x15, 0x73, 0xe9,
	0x14, 0xe4, 0x85, 0xae, 0x3c, 0x81, 0xcb, 0x3d, 0x49, 0x8a, 0xf7, 0xdb, 0x51, 0xe7, 0xa8, 0xdb,
	0x12, 0xa1, 0x60, 0xb1, 0x2e, 0x58, 0xf4, 0x10, 0xf5, 0x07, 0xa9, 0x2b, 0xc8, 0x1e, 0x86, 0xd8,
	0x59, 0x48, 0x8d, 0x49, 0x12, 0x1f, 0xb0, 0xe3, 0x15, 0x4b, 0x97, 0x9e, 0xc0, 0xac, 0x61, 0x07,
	0xff, 0x84, 0xf1, 0x90, 0x1b, 0xac, 0x62, 0x35, 0xed, 0xf4, 0x7b, 0xc4, 0x8e, 0xae, 0xed, 0xcc,
	0xdb, 0xec, 0xc9, 0xe0, 0xf5, 0x78, 0x92, 0xa7, 0xc3, 0xb3, 0xd1, 0x79, 0x3a, 0xec, 0xe7, 0x6f,
	0xd2, 0x6c, 0x3c, 0xc9, 0x47, 0xef, 0x27, 0xfd, 0x51, 0x3a, 0xec, 0x3f, 0xb8, 0xc5, 0x9f, 0xb3,
	0xa7, 0x41, 0xfb, 0x9b, 0x25, 0xe2, 0xcf, 0xd8, 0xe3, 0x6d, 0x48, 0x33, 0x6d, 0x0c, 0x7b, 0xf5,
	0x5f, 0xfe, 0x60, 0x6c, 0x3b, 0xf6, 0x7b, 0x17, 0x3f, 0x7e, 0x9d, 0x44, 0x17, 0x93, 0xdd, 0x2e,
	0xc1, 0x7e, 0x51, 0xff, 0x71, 0x0d, 0xd3, 0x3b, 0x4d, 0x37, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0x28, 0x48, 0x25, 0x68, 0x03, 0x00, 0x00,
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
	if this.Request != that1.Request {
		return false
	}
	if this.MeasureRequestInternally != that1.MeasureRequestInternally {
		return false
	}
	if this.Response != that1.Response {
		return false
	}
	if !this.ChargeClusterStat.Equal(that1.ChargeClusterStat) {
		return false
	}
	if !this.ChargeListenerStat.Equal(that1.ChargeListenerStat) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
