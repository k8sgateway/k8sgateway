// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/hcm/hcm.proto

package hcm

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Contains various settings for Envoy's http connection manager.
// See here for more information: https://www.envoyproxy.io/docs/envoy/v1.9.0/configuration/http_conn_man/http_conn_man
type HttpConnectionManagerSettings struct {
	SkipXffAppend       bool               `protobuf:"varint,1,opt,name=skip_xff_append,json=skipXffAppend,proto3" json:"skip_xff_append,omitempty"`
	Via                 string             `protobuf:"bytes,2,opt,name=via,proto3" json:"via,omitempty"`
	XffNumTrustedHops   uint32             `protobuf:"varint,3,opt,name=xff_num_trusted_hops,json=xffNumTrustedHops,proto3" json:"xff_num_trusted_hops,omitempty"`
	UseRemoteAddress    *types.BoolValue   `protobuf:"bytes,4,opt,name=use_remote_address,json=useRemoteAddress,proto3" json:"use_remote_address,omitempty"`
	GenerateRequestId   *types.BoolValue   `protobuf:"bytes,5,opt,name=generate_request_id,json=generateRequestId,proto3" json:"generate_request_id,omitempty"`
	Proxy_100Continue   bool               `protobuf:"varint,6,opt,name=proxy_100_continue,json=proxy100Continue,proto3" json:"proxy_100_continue,omitempty"`
	StreamIdleTimeout   *time.Duration     `protobuf:"bytes,7,opt,name=stream_idle_timeout,json=streamIdleTimeout,proto3,stdduration" json:"stream_idle_timeout,omitempty"`
	IdleTimeout         *time.Duration     `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3,stdduration" json:"idle_timeout,omitempty"`
	MaxRequestHeadersKb *types.UInt32Value `protobuf:"bytes,9,opt,name=max_request_headers_kb,json=maxRequestHeadersKb,proto3" json:"max_request_headers_kb,omitempty"`
	RequestTimeout      *time.Duration     `protobuf:"bytes,10,opt,name=request_timeout,json=requestTimeout,proto3,stdduration" json:"request_timeout,omitempty"`
	DrainTimeout        *time.Duration     `protobuf:"bytes,12,opt,name=drain_timeout,json=drainTimeout,proto3,stdduration" json:"drain_timeout,omitempty"`
	DelayedCloseTimeout *time.Duration     `protobuf:"bytes,13,opt,name=delayed_close_timeout,json=delayedCloseTimeout,proto3,stdduration" json:"delayed_close_timeout,omitempty"`
	ServerName          string             `protobuf:"bytes,14,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// For explanation of these settings see: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/protocol.proto#envoy-api-msg-core-http1protocoloptions
	AcceptHttp_10         bool                                           `protobuf:"varint,15,opt,name=accept_http_10,json=acceptHttp10,proto3" json:"accept_http_10,omitempty"`
	DefaultHostForHttp_10 string                                         `protobuf:"bytes,16,opt,name=default_host_for_http_10,json=defaultHostForHttp10,proto3" json:"default_host_for_http_10,omitempty"`
	Tracing               *HttpConnectionManagerSettings_TracingSettings `protobuf:"bytes,17,opt,name=tracing,proto3" json:"tracing,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                       `json:"-"`
	XXX_unrecognized      []byte                                         `json:"-"`
	XXX_sizecache         int32                                          `json:"-"`
}

func (m *HttpConnectionManagerSettings) Reset()         { *m = HttpConnectionManagerSettings{} }
func (m *HttpConnectionManagerSettings) String() string { return proto.CompactTextString(m) }
func (*HttpConnectionManagerSettings) ProtoMessage()    {}
func (*HttpConnectionManagerSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9393403d6dbb8c, []int{0}
}
func (m *HttpConnectionManagerSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConnectionManagerSettings.Unmarshal(m, b)
}
func (m *HttpConnectionManagerSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConnectionManagerSettings.Marshal(b, m, deterministic)
}
func (m *HttpConnectionManagerSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConnectionManagerSettings.Merge(m, src)
}
func (m *HttpConnectionManagerSettings) XXX_Size() int {
	return xxx_messageInfo_HttpConnectionManagerSettings.Size(m)
}
func (m *HttpConnectionManagerSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpConnectionManagerSettings.DiscardUnknown(m)
}

var xxx_messageInfo_HttpConnectionManagerSettings proto.InternalMessageInfo

func (m *HttpConnectionManagerSettings) GetSkipXffAppend() bool {
	if m != nil {
		return m.SkipXffAppend
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

func (m *HttpConnectionManagerSettings) GetXffNumTrustedHops() uint32 {
	if m != nil {
		return m.XffNumTrustedHops
	}
	return 0
}

func (m *HttpConnectionManagerSettings) GetUseRemoteAddress() *types.BoolValue {
	if m != nil {
		return m.UseRemoteAddress
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetGenerateRequestId() *types.BoolValue {
	if m != nil {
		return m.GenerateRequestId
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetProxy_100Continue() bool {
	if m != nil {
		return m.Proxy_100Continue
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetStreamIdleTimeout() *time.Duration {
	if m != nil {
		return m.StreamIdleTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetIdleTimeout() *time.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetMaxRequestHeadersKb() *types.UInt32Value {
	if m != nil {
		return m.MaxRequestHeadersKb
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetDrainTimeout() *time.Duration {
	if m != nil {
		return m.DrainTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetDelayedCloseTimeout() *time.Duration {
	if m != nil {
		return m.DelayedCloseTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *HttpConnectionManagerSettings) GetAcceptHttp_10() bool {
	if m != nil {
		return m.AcceptHttp_10
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetDefaultHostForHttp_10() string {
	if m != nil {
		return m.DefaultHostForHttp_10
	}
	return ""
}

func (m *HttpConnectionManagerSettings) GetTracing() *HttpConnectionManagerSettings_TracingSettings {
	if m != nil {
		return m.Tracing
	}
	return nil
}

// Contains settings for configuring Envoy's tracing capabilities at the listener level.
type HttpConnectionManagerSettings_TracingSettings struct {
	// Optional. If specified, Envoy will include the headers and header values for any matching request headers.
	RequestHeadersForTags []string `protobuf:"bytes,1,rep,name=request_headers_for_tags,json=requestHeadersForTags,proto3" json:"request_headers_for_tags,omitempty"`
	// Optional. If true, Envoy will include logs for streaming events. Default: false.
	Verbose              bool     `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpConnectionManagerSettings_TracingSettings) Reset() {
	*m = HttpConnectionManagerSettings_TracingSettings{}
}
func (m *HttpConnectionManagerSettings_TracingSettings) String() string {
	return proto.CompactTextString(m)
}
func (*HttpConnectionManagerSettings_TracingSettings) ProtoMessage() {}
func (*HttpConnectionManagerSettings_TracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9393403d6dbb8c, []int{0, 0}
}
func (m *HttpConnectionManagerSettings_TracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConnectionManagerSettings_TracingSettings.Unmarshal(m, b)
}
func (m *HttpConnectionManagerSettings_TracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConnectionManagerSettings_TracingSettings.Marshal(b, m, deterministic)
}
func (m *HttpConnectionManagerSettings_TracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConnectionManagerSettings_TracingSettings.Merge(m, src)
}
func (m *HttpConnectionManagerSettings_TracingSettings) XXX_Size() int {
	return xxx_messageInfo_HttpConnectionManagerSettings_TracingSettings.Size(m)
}
func (m *HttpConnectionManagerSettings_TracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpConnectionManagerSettings_TracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_HttpConnectionManagerSettings_TracingSettings proto.InternalMessageInfo

func (m *HttpConnectionManagerSettings_TracingSettings) GetRequestHeadersForTags() []string {
	if m != nil {
		return m.RequestHeadersForTags
	}
	return nil
}

func (m *HttpConnectionManagerSettings_TracingSettings) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

// Contains settings for configuring Envoy's tracing capabilities at the route level.
// Note: must also specify ListenerTracingSettings for the associated listener.
type RouteTracingSettings struct {
	// Optional. If set, will be used to identify the route that produced the trace.
	// Note that this value will be overridden if the "x-envoy-decorator-operation" header is passed.
	RouteDescriptor      string   `protobuf:"bytes,1,opt,name=route_descriptor,json=routeDescriptor,proto3" json:"route_descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteTracingSettings) Reset()         { *m = RouteTracingSettings{} }
func (m *RouteTracingSettings) String() string { return proto.CompactTextString(m) }
func (*RouteTracingSettings) ProtoMessage()    {}
func (*RouteTracingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9393403d6dbb8c, []int{1}
}
func (m *RouteTracingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTracingSettings.Unmarshal(m, b)
}
func (m *RouteTracingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTracingSettings.Marshal(b, m, deterministic)
}
func (m *RouteTracingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTracingSettings.Merge(m, src)
}
func (m *RouteTracingSettings) XXX_Size() int {
	return xxx_messageInfo_RouteTracingSettings.Size(m)
}
func (m *RouteTracingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTracingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTracingSettings proto.InternalMessageInfo

func (m *RouteTracingSettings) GetRouteDescriptor() string {
	if m != nil {
		return m.RouteDescriptor
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpConnectionManagerSettings)(nil), "hcm.plugins.gloo.solo.io.HttpConnectionManagerSettings")
	proto.RegisterType((*HttpConnectionManagerSettings_TracingSettings)(nil), "hcm.plugins.gloo.solo.io.HttpConnectionManagerSettings.TracingSettings")
	proto.RegisterType((*RouteTracingSettings)(nil), "hcm.plugins.gloo.solo.io.RouteTracingSettings")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/hcm/hcm.proto", fileDescriptor_1c9393403d6dbb8c)
}

var fileDescriptor_1c9393403d6dbb8c = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x8e, 0xe3, 0x34,
	0x14, 0x86, 0x15, 0x66, 0xd9, 0x69, 0x3d, 0xed, 0xb4, 0x75, 0xbb, 0xc8, 0x54, 0xb0, 0x5b, 0xad,
	0x10, 0x2a, 0x12, 0x24, 0xed, 0xae, 0x04, 0x37, 0xdc, 0x4c, 0x67, 0xb4, 0x74, 0x40, 0x2c, 0x22,
	0x5b, 0x10, 0xe2, 0xc6, 0x72, 0xe3, 0x93, 0xd4, 0x4c, 0x92, 0x13, 0x6c, 0xa7, 0x74, 0xdf, 0x84,
	0x47, 0xe0, 0xad, 0x90, 0xe0, 0x45, 0x50, 0x9c, 0xa6, 0xb0, 0x33, 0x0c, 0xf4, 0x22, 0x92, 0x7d,
	0x7e, 0xff, 0x9f, 0xed, 0x93, 0xe3, 0x43, 0x16, 0x89, 0xb2, 0x9b, 0x72, 0xed, 0x47, 0x98, 0x05,
	0x06, 0x53, 0xfc, 0x44, 0x61, 0x90, 0xa4, 0x88, 0x41, 0xa1, 0xf1, 0x27, 0x88, 0xac, 0xa9, 0x67,
	0xa2, 0x50, 0xc1, 0x76, 0x1e, 0x14, 0x69, 0x99, 0xa8, 0xdc, 0x04, 0x9b, 0x28, 0xab, 0x3e, 0xbf,
	0xd0, 0x68, 0x91, 0x32, 0x37, 0xac, 0x25, 0xbf, 0x5a, 0xee, 0x57, 0x24, 0x5f, 0xe1, 0x78, 0x94,
	0x60, 0x82, 0x6e, 0x51, 0x50, 0x8d, 0xea, 0xf5, 0xe3, 0xc7, 0x09, 0x62, 0x92, 0x42, 0xe0, 0x66,
	0xeb, 0x32, 0x0e, 0x7e, 0xd1, 0xa2, 0x28, 0x40, 0x9b, 0xfb, 0x74, 0x59, 0x6a, 0x61, 0x15, 0xe6,
	0xb5, 0xfe, 0xf4, 0xcf, 0x16, 0x79, 0x7f, 0x69, 0x6d, 0x71, 0x89, 0x79, 0x0e, 0x51, 0x25, 0x7c,
	0x2d, 0x72, 0x91, 0x80, 0x7e, 0x05, 0xd6, 0xaa, 0x3c, 0x31, 0xf4, 0x43, 0xd2, 0x33, 0x37, 0xaa,
	0xe0, 0xbb, 0x38, 0xe6, 0x15, 0x3a, 0x97, 0xcc, 0x9b, 0x78, 0xd3, 0x56, 0xd8, 0xad, 0xc2, 0x3f,
	0xc4, 0xf1, 0x85, 0x0b, 0xd2, 0x3e, 0x39, 0xd9, 0x2a, 0xc1, 0xde, 0x9a, 0x78, 0xd3, 0x76, 0x58,
	0x0d, 0x69, 0x40, 0x46, 0x95, 0x29, 0x2f, 0x33, 0x6e, 0x75, 0x69, 0x2c, 0x48, 0xbe, 0xc1, 0xc2,
	0xb0, 0x93, 0x89, 0x37, 0xed, 0x86, 0x83, 0x5d, 0x1c, 0xbf, 0x2c, 0xb3, 0x55, 0xad, 0x2c, 0xb1,
	0x30, 0x74, 0x49, 0x68, 0x69, 0x80, 0x6b, 0xc8, 0xd0, 0x02, 0x17, 0x52, 0x6a, 0x30, 0x86, 0x3d,
	0x98, 0x78, 0xd3, 0xb3, 0x67, 0x63, 0xbf, 0xbe, 0x89, 0xdf, 0xdc, 0xc4, 0x5f, 0x20, 0xa6, 0xdf,
	0x8b, 0xb4, 0x84, 0xb0, 0x5f, 0x1a, 0x08, 0x9d, 0xe9, 0xa2, 0xf6, 0xd0, 0x2f, 0xc9, 0x30, 0x81,
	0x1c, 0xb4, 0xb0, 0x15, 0xee, 0xe7, 0x12, 0x8c, 0xe5, 0x4a, 0xb2, 0xb7, 0xff, 0x17, 0x35, 0x68,
	0x6c, 0x61, 0xed, 0xba, 0x96, 0xf4, 0x63, 0x42, 0x0b, 0x8d, 0xbb, 0xd7, 0x7c, 0x3e, 0x9b, 0xf1,
	0x08, 0x73, 0xab, 0xf2, 0x12, 0xd8, 0x43, 0x97, 0x83, 0xbe, 0x53, 0xe6, 0xb3, 0xd9, 0xe5, 0x3e,
	0x4e, 0xbf, 0x21, 0x43, 0x63, 0x35, 0x88, 0x8c, 0x2b, 0x99, 0x02, 0xb7, 0x2a, 0x03, 0x2c, 0x2d,
	0x3b, 0x75, 0x3b, 0xbf, 0x7b, 0x67, 0xe7, 0xab, 0xfd, 0xef, 0x58, 0x3c, 0xf8, 0xf5, 0xf7, 0x27,
	0x5e, 0x38, 0xa8, 0xbd, 0xd7, 0x32, 0x85, 0x55, 0xed, 0xa4, 0x0b, 0xd2, 0x79, 0x83, 0xd4, 0x3a,
	0x8e, 0x74, 0xa6, 0xfe, 0xc1, 0xf8, 0x96, 0xbc, 0x93, 0x89, 0xdd, 0x21, 0x13, 0x1b, 0x10, 0x12,
	0xb4, 0xe1, 0x37, 0x6b, 0xd6, 0x76, 0xb4, 0xf7, 0xee, 0xd0, 0xbe, 0xbb, 0xce, 0xed, 0xf3, 0x67,
	0x75, 0x4e, 0x86, 0x99, 0xd8, 0xed, 0xd3, 0xb1, 0xac, 0x9d, 0x5f, 0xad, 0xe9, 0x92, 0xf4, 0x1a,
	0x5c, 0x73, 0x32, 0x72, 0xdc, 0xc9, 0xce, 0xf7, 0xbe, 0xe6, 0x70, 0x57, 0xa4, 0x2b, 0xb5, 0x50,
	0xf9, 0x81, 0xd3, 0x39, 0x8e, 0xd3, 0x71, 0xae, 0x86, 0xf2, 0x8a, 0x3c, 0x92, 0x90, 0x8a, 0xd7,
	0x20, 0x79, 0x94, 0xa2, 0xf9, 0x3b, 0x5f, 0xdd, 0xe3, 0x68, 0xc3, 0xbd, 0xfb, 0xb2, 0x32, 0x37,
	0xd0, 0x27, 0xe4, 0xcc, 0x80, 0xde, 0x82, 0xe6, 0xb9, 0xc8, 0x80, 0x9d, 0xbb, 0xda, 0x26, 0x75,
	0xe8, 0xa5, 0xc8, 0x80, 0x7e, 0x40, 0xce, 0x45, 0x14, 0x41, 0x61, 0xf9, 0xc6, 0xda, 0x82, 0xcf,
	0x67, 0xac, 0xe7, 0xea, 0xa2, 0x53, 0x47, 0xab, 0x97, 0x35, 0x9f, 0xd1, 0x4f, 0x09, 0x93, 0x10,
	0x8b, 0x32, 0xb5, 0x7c, 0x83, 0xc6, 0xf2, 0x18, 0xf5, 0x61, 0x7d, 0xdf, 0x31, 0x47, 0x7b, 0x7d,
	0x89, 0xc6, 0xbe, 0x40, 0xbd, 0xf7, 0x09, 0x72, 0x6a, 0xb5, 0x88, 0x54, 0x9e, 0xb0, 0x81, 0xbb,
	0xc5, 0x17, 0xfe, 0x7d, 0xed, 0xc1, 0xff, 0xcf, 0x47, 0xec, 0xaf, 0x6a, 0x4c, 0x33, 0x0f, 0x1b,
	0xee, 0x58, 0x92, 0xde, 0x2d, 0x8d, 0x7e, 0x46, 0xd8, 0xed, 0x42, 0xa9, 0x0e, 0x6c, 0x45, 0x62,
	0x98, 0x37, 0x39, 0x99, 0xb6, 0xc3, 0x47, 0xfa, 0x8d, 0x6a, 0x78, 0x81, 0x7a, 0x25, 0x12, 0x43,
	0x19, 0x39, 0xdd, 0x82, 0x5e, 0xa3, 0x01, 0xd7, 0x05, 0x5a, 0x61, 0x33, 0x7d, 0x7a, 0x41, 0x46,
	0x21, 0x96, 0x16, 0x6e, 0x6f, 0xf5, 0x11, 0xe9, 0xeb, 0x2a, 0xce, 0x25, 0x98, 0x48, 0xab, 0xc2,
	0xa2, 0x76, 0xcd, 0xa5, 0x1d, 0xf6, 0x5c, 0xfc, 0xea, 0x10, 0x5e, 0x2c, 0x7e, 0xfb, 0xe3, 0xb1,
	0xf7, 0xe3, 0xe7, 0xc7, 0xb5, 0xd8, 0xe2, 0x26, 0xf9, 0x97, 0x36, 0xbb, 0x7e, 0xe8, 0x7e, 0xfe,
	0xf3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0xd7, 0xb1, 0x39, 0xa9, 0x05, 0x00, 0x00,
}

func (this *HttpConnectionManagerSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpConnectionManagerSettings)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings)
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
	if this.SkipXffAppend != that1.SkipXffAppend {
		return false
	}
	if this.Via != that1.Via {
		return false
	}
	if this.XffNumTrustedHops != that1.XffNumTrustedHops {
		return false
	}
	if !this.UseRemoteAddress.Equal(that1.UseRemoteAddress) {
		return false
	}
	if !this.GenerateRequestId.Equal(that1.GenerateRequestId) {
		return false
	}
	if this.Proxy_100Continue != that1.Proxy_100Continue {
		return false
	}
	if this.StreamIdleTimeout != nil && that1.StreamIdleTimeout != nil {
		if *this.StreamIdleTimeout != *that1.StreamIdleTimeout {
			return false
		}
	} else if this.StreamIdleTimeout != nil {
		return false
	} else if that1.StreamIdleTimeout != nil {
		return false
	}
	if this.IdleTimeout != nil && that1.IdleTimeout != nil {
		if *this.IdleTimeout != *that1.IdleTimeout {
			return false
		}
	} else if this.IdleTimeout != nil {
		return false
	} else if that1.IdleTimeout != nil {
		return false
	}
	if !this.MaxRequestHeadersKb.Equal(that1.MaxRequestHeadersKb) {
		return false
	}
	if this.RequestTimeout != nil && that1.RequestTimeout != nil {
		if *this.RequestTimeout != *that1.RequestTimeout {
			return false
		}
	} else if this.RequestTimeout != nil {
		return false
	} else if that1.RequestTimeout != nil {
		return false
	}
	if this.DrainTimeout != nil && that1.DrainTimeout != nil {
		if *this.DrainTimeout != *that1.DrainTimeout {
			return false
		}
	} else if this.DrainTimeout != nil {
		return false
	} else if that1.DrainTimeout != nil {
		return false
	}
	if this.DelayedCloseTimeout != nil && that1.DelayedCloseTimeout != nil {
		if *this.DelayedCloseTimeout != *that1.DelayedCloseTimeout {
			return false
		}
	} else if this.DelayedCloseTimeout != nil {
		return false
	} else if that1.DelayedCloseTimeout != nil {
		return false
	}
	if this.ServerName != that1.ServerName {
		return false
	}
	if this.AcceptHttp_10 != that1.AcceptHttp_10 {
		return false
	}
	if this.DefaultHostForHttp_10 != that1.DefaultHostForHttp_10 {
		return false
	}
	if !this.Tracing.Equal(that1.Tracing) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HttpConnectionManagerSettings_TracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpConnectionManagerSettings_TracingSettings)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings_TracingSettings)
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
	if len(this.RequestHeadersForTags) != len(that1.RequestHeadersForTags) {
		return false
	}
	for i := range this.RequestHeadersForTags {
		if this.RequestHeadersForTags[i] != that1.RequestHeadersForTags[i] {
			return false
		}
	}
	if this.Verbose != that1.Verbose {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTracingSettings)
	if !ok {
		that2, ok := that.(RouteTracingSettings)
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
	if this.RouteDescriptor != that1.RouteDescriptor {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
