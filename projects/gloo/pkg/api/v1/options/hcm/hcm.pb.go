// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto

package hcm

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	protocol_upgrade "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"
	tracing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HttpConnectionManagerSettings_ForwardClientCertDetails int32

const (
	HttpConnectionManagerSettings_SANITIZE            HttpConnectionManagerSettings_ForwardClientCertDetails = 0
	HttpConnectionManagerSettings_FORWARD_ONLY        HttpConnectionManagerSettings_ForwardClientCertDetails = 1
	HttpConnectionManagerSettings_APPEND_FORWARD      HttpConnectionManagerSettings_ForwardClientCertDetails = 2
	HttpConnectionManagerSettings_SANITIZE_SET        HttpConnectionManagerSettings_ForwardClientCertDetails = 3
	HttpConnectionManagerSettings_ALWAYS_FORWARD_ONLY HttpConnectionManagerSettings_ForwardClientCertDetails = 4
)

var HttpConnectionManagerSettings_ForwardClientCertDetails_name = map[int32]string{
	0: "SANITIZE",
	1: "FORWARD_ONLY",
	2: "APPEND_FORWARD",
	3: "SANITIZE_SET",
	4: "ALWAYS_FORWARD_ONLY",
}

var HttpConnectionManagerSettings_ForwardClientCertDetails_value = map[string]int32{
	"SANITIZE":            0,
	"FORWARD_ONLY":        1,
	"APPEND_FORWARD":      2,
	"SANITIZE_SET":        3,
	"ALWAYS_FORWARD_ONLY": 4,
}

func (x HttpConnectionManagerSettings_ForwardClientCertDetails) String() string {
	return proto.EnumName(HttpConnectionManagerSettings_ForwardClientCertDetails_name, int32(x))
}

func (HttpConnectionManagerSettings_ForwardClientCertDetails) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08263ad65d35164d, []int{0, 0}
}

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
	AcceptHttp_10               bool                                                       `protobuf:"varint,15,opt,name=accept_http_10,json=acceptHttp10,proto3" json:"accept_http_10,omitempty"`
	DefaultHostForHttp_10       string                                                     `protobuf:"bytes,16,opt,name=default_host_for_http_10,json=defaultHostForHttp10,proto3" json:"default_host_for_http_10,omitempty"`
	Tracing                     *tracing.ListenerTracingSettings                           `protobuf:"bytes,17,opt,name=tracing,proto3" json:"tracing,omitempty"`
	ForwardClientCertDetails    HttpConnectionManagerSettings_ForwardClientCertDetails     `protobuf:"varint,18,opt,name=forward_client_cert_details,json=forwardClientCertDetails,proto3,enum=hcm.options.gloo.solo.io.HttpConnectionManagerSettings_ForwardClientCertDetails" json:"forward_client_cert_details,omitempty"`
	SetCurrentClientCertDetails *HttpConnectionManagerSettings_SetCurrentClientCertDetails `protobuf:"bytes,19,opt,name=set_current_client_cert_details,json=setCurrentClientCertDetails,proto3" json:"set_current_client_cert_details,omitempty"`
	PreserveExternalRequestId   bool                                                       `protobuf:"varint,20,opt,name=preserve_external_request_id,json=preserveExternalRequestId,proto3" json:"preserve_external_request_id,omitempty"`
	// HttpConnectionManager configuration for protocol upgrade requests.
	// Note: WebSocket upgrades are enabled by default on the HTTP Connection Manager and must be explicitly disabled.
	Upgrades             []*protocol_upgrade.ProtocolUpgradeConfig `protobuf:"bytes,21,rep,name=upgrades,proto3" json:"upgrades,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *HttpConnectionManagerSettings) Reset()         { *m = HttpConnectionManagerSettings{} }
func (m *HttpConnectionManagerSettings) String() string { return proto.CompactTextString(m) }
func (*HttpConnectionManagerSettings) ProtoMessage()    {}
func (*HttpConnectionManagerSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_08263ad65d35164d, []int{0}
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

func (m *HttpConnectionManagerSettings) GetTracing() *tracing.ListenerTracingSettings {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetForwardClientCertDetails() HttpConnectionManagerSettings_ForwardClientCertDetails {
	if m != nil {
		return m.ForwardClientCertDetails
	}
	return HttpConnectionManagerSettings_SANITIZE
}

func (m *HttpConnectionManagerSettings) GetSetCurrentClientCertDetails() *HttpConnectionManagerSettings_SetCurrentClientCertDetails {
	if m != nil {
		return m.SetCurrentClientCertDetails
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetPreserveExternalRequestId() bool {
	if m != nil {
		return m.PreserveExternalRequestId
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetUpgrades() []*protocol_upgrade.ProtocolUpgradeConfig {
	if m != nil {
		return m.Upgrades
	}
	return nil
}

type HttpConnectionManagerSettings_SetCurrentClientCertDetails struct {
	Subject              *types.BoolValue `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Cert                 bool             `protobuf:"varint,2,opt,name=cert,proto3" json:"cert,omitempty"`
	Chain                bool             `protobuf:"varint,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Dns                  bool             `protobuf:"varint,4,opt,name=dns,proto3" json:"dns,omitempty"`
	Uri                  bool             `protobuf:"varint,5,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Reset() {
	*m = HttpConnectionManagerSettings_SetCurrentClientCertDetails{}
}
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) String() string {
	return proto.CompactTextString(m)
}
func (*HttpConnectionManagerSettings_SetCurrentClientCertDetails) ProtoMessage() {}
func (*HttpConnectionManagerSettings_SetCurrentClientCertDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_08263ad65d35164d, []int{0, 0}
}
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConnectionManagerSettings_SetCurrentClientCertDetails.Unmarshal(m, b)
}
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConnectionManagerSettings_SetCurrentClientCertDetails.Marshal(b, m, deterministic)
}
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConnectionManagerSettings_SetCurrentClientCertDetails.Merge(m, src)
}
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) XXX_Size() int {
	return xxx_messageInfo_HttpConnectionManagerSettings_SetCurrentClientCertDetails.Size(m)
}
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpConnectionManagerSettings_SetCurrentClientCertDetails.DiscardUnknown(m)
}

var xxx_messageInfo_HttpConnectionManagerSettings_SetCurrentClientCertDetails proto.InternalMessageInfo

func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) GetSubject() *types.BoolValue {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) GetCert() bool {
	if m != nil {
		return m.Cert
	}
	return false
}

func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) GetChain() bool {
	if m != nil {
		return m.Chain
	}
	return false
}

func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) GetDns() bool {
	if m != nil {
		return m.Dns
	}
	return false
}

func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) GetUri() bool {
	if m != nil {
		return m.Uri
	}
	return false
}

func init() {
	proto.RegisterEnum("hcm.options.gloo.solo.io.HttpConnectionManagerSettings_ForwardClientCertDetails", HttpConnectionManagerSettings_ForwardClientCertDetails_name, HttpConnectionManagerSettings_ForwardClientCertDetails_value)
	proto.RegisterType((*HttpConnectionManagerSettings)(nil), "hcm.options.gloo.solo.io.HttpConnectionManagerSettings")
	proto.RegisterType((*HttpConnectionManagerSettings_SetCurrentClientCertDetails)(nil), "hcm.options.gloo.solo.io.HttpConnectionManagerSettings.SetCurrentClientCertDetails")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto", fileDescriptor_08263ad65d35164d)
}

var fileDescriptor_08263ad65d35164d = []byte{
	// 976 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdd, 0x52, 0xdb, 0x46,
	0x14, 0xc7, 0xab, 0x40, 0x82, 0xb3, 0x18, 0x30, 0x6b, 0xd2, 0x2a, 0x90, 0x06, 0x4f, 0xa6, 0xd3,
	0xf1, 0x45, 0x2b, 0x1b, 0xd2, 0x8f, 0x9b, 0xcc, 0x74, 0x8c, 0x81, 0xb1, 0x5b, 0x0a, 0x54, 0x26,
	0xcd, 0xc7, 0xcd, 0xce, 0x5a, 0x3a, 0x92, 0xb7, 0x48, 0x5a, 0x75, 0x77, 0x45, 0xcc, 0x53, 0xf4,
	0xb6, 0xbd, 0xec, 0x5d, 0x1f, 0xa1, 0xef, 0xd1, 0x07, 0xe8, 0x4c, 0xdf, 0xa1, 0xf7, 0x9d, 0xdd,
	0x95, 0x48, 0x32, 0x7c, 0x79, 0x7a, 0xc1, 0xb0, 0x3a, 0xe7, 0xfc, 0x7f, 0x3a, 0x7b, 0xb4, 0xfa,
	0x5b, 0x68, 0x27, 0x66, 0x6a, 0x52, 0x8c, 0xbd, 0x80, 0xa7, 0x1d, 0xc9, 0x13, 0xfe, 0x39, 0xe3,
	0x9d, 0x38, 0xe1, 0xbc, 0x93, 0x0b, 0xfe, 0x13, 0x04, 0x4a, 0xda, 0x2b, 0x9a, 0xb3, 0xce, 0xd9,
	0x56, 0x87, 0xe7, 0x8a, 0xf1, 0x4c, 0x76, 0x26, 0x41, 0xaa, 0xff, 0xbc, 0x5c, 0x70, 0xc5, 0xb1,
	0xab, 0x97, 0x65, 0xca, 0xd3, 0xe5, 0x9e, 0x26, 0x79, 0x8c, 0xaf, 0xaf, 0xc5, 0x3c, 0xe6, 0xa6,
	0xa8, 0xa3, 0x57, 0xb6, 0x7e, 0xfd, 0x71, 0xcc, 0x79, 0x9c, 0x40, 0xc7, 0x5c, 0x8d, 0x8b, 0xa8,
	0xf3, 0x46, 0xd0, 0x3c, 0x07, 0x21, 0xaf, 0xcb, 0x87, 0x85, 0xa0, 0x9a, 0x5e, 0xe6, 0xbf, 0xbe,
	0xbd, 0x41, 0x25, 0x68, 0xc0, 0xb2, 0xb8, 0xfa, 0x5f, 0x0a, 0x87, 0xb7, 0x0b, 0x4d, 0x61, 0xc0,
	0x13, 0x52, 0xe4, 0xb1, 0xa0, 0x21, 0x5c, 0x0a, 0x94, 0x28, 0x0c, 0x53, 0x65, 0x37, 0x06, 0x53,
	0x65, 0x63, 0x4f, 0xfe, 0x5a, 0x42, 0x1f, 0x0f, 0x94, 0xca, 0xfb, 0x3c, 0xcb, 0x20, 0xd0, 0xbc,
	0xef, 0x69, 0x46, 0x63, 0x10, 0x23, 0x50, 0x8a, 0x65, 0xb1, 0xc4, 0x9f, 0xa2, 0x15, 0x79, 0xca,
	0x72, 0x32, 0x8d, 0x22, 0xa2, 0xb7, 0x9c, 0x85, 0xae, 0xd3, 0x72, 0xda, 0x35, 0x7f, 0x49, 0x87,
	0x5f, 0x46, 0x51, 0xcf, 0x04, 0x71, 0x03, 0xcd, 0x9d, 0x31, 0xea, 0xde, 0x69, 0x39, 0xed, 0xfb,
	0xbe, 0x5e, 0xe2, 0x0e, 0x5a, 0xd3, 0xa2, 0xac, 0x48, 0x89, 0x12, 0x85, 0x54, 0x10, 0x92, 0x09,
	0xcf, 0xa5, 0x3b, 0xd7, 0x72, 0xda, 0x4b, 0xfe, 0xea, 0x34, 0x8a, 0x0e, 0x8b, 0xf4, 0xc4, 0x66,
	0x06, 0x3c, 0x97, 0x78, 0x80, 0x70, 0x21, 0x81, 0x08, 0x48, 0xb9, 0x02, 0x42, 0xc3, 0x50, 0x80,
	0x94, 0xee, 0x7c, 0xcb, 0x69, 0x2f, 0x6e, 0xaf, 0x7b, 0x76, 0xc2, 0x5e, 0x35, 0x61, 0x6f, 0x87,
	0xf3, 0xe4, 0x47, 0x9a, 0x14, 0xe0, 0x37, 0x0a, 0x09, 0xbe, 0x11, 0xf5, 0xac, 0x06, 0x7f, 0x8b,
	0x9a, 0x31, 0x64, 0x20, 0xa8, 0xd2, 0xb8, 0x9f, 0x0b, 0x90, 0x8a, 0xb0, 0xd0, 0xbd, 0x7b, 0x2b,
	0x6a, 0xb5, 0x92, 0xf9, 0x56, 0x35, 0x0c, 0xf1, 0x67, 0x08, 0xe7, 0x82, 0x4f, 0xcf, 0xc9, 0x56,
	0xb7, 0x4b, 0x02, 0x9e, 0x29, 0x96, 0x15, 0xe0, 0xde, 0x33, 0x33, 0x68, 0x98, 0xcc, 0x56, 0xb7,
	0xdb, 0x2f, 0xe3, 0xf8, 0x08, 0x35, 0xa5, 0x12, 0x40, 0x53, 0xc2, 0xc2, 0x04, 0x88, 0x62, 0x29,
	0xf0, 0x42, 0xb9, 0x0b, 0xe6, 0xce, 0x0f, 0x2f, 0xdd, 0x79, 0xb7, 0x3c, 0x26, 0x3b, 0xf3, 0xbf,
	0xfe, 0xbd, 0xe9, 0xf8, 0xab, 0x56, 0x3b, 0x0c, 0x13, 0x38, 0xb1, 0x4a, 0xbc, 0x83, 0xea, 0xef,
	0x91, 0x6a, 0xb3, 0x91, 0x16, 0xd9, 0x3b, 0x8c, 0x1f, 0xd0, 0x87, 0x29, 0x9d, 0x5e, 0x4c, 0x62,
	0x02, 0x34, 0x04, 0x21, 0xc9, 0xe9, 0xd8, 0xbd, 0x6f, 0x68, 0x8f, 0x2e, 0xd1, 0x9e, 0x0f, 0x33,
	0xf5, 0x74, 0xdb, 0xce, 0xa4, 0x99, 0xd2, 0x69, 0x39, 0x8e, 0x81, 0x55, 0x7e, 0x37, 0xc6, 0x03,
	0xb4, 0x52, 0xe1, 0xaa, 0xce, 0xd0, 0x6c, 0x9d, 0x2d, 0x97, 0xba, 0xaa, 0xb9, 0x5d, 0xb4, 0x14,
	0x0a, 0xca, 0xb2, 0x0b, 0x4e, 0x7d, 0x36, 0x4e, 0xdd, 0xa8, 0x2a, 0xca, 0x08, 0x3d, 0x08, 0x21,
	0xa1, 0xe7, 0x10, 0x92, 0x20, 0xe1, 0xf2, 0xed, 0xbc, 0x96, 0x66, 0xa3, 0x35, 0x4b, 0x75, 0x5f,
	0x8b, 0x2b, 0xe8, 0x26, 0x5a, 0x94, 0x20, 0xce, 0x40, 0x90, 0x8c, 0xa6, 0xe0, 0x2e, 0x9b, 0xb3,
	0x8d, 0x6c, 0xe8, 0x90, 0xa6, 0x80, 0x3f, 0x41, 0xcb, 0x34, 0x08, 0x20, 0x57, 0x64, 0xa2, 0x54,
	0x4e, 0xb6, 0xba, 0xee, 0x8a, 0x39, 0x17, 0x75, 0x1b, 0xd5, 0x6f, 0xd6, 0x56, 0x17, 0x7f, 0x85,
	0xdc, 0x10, 0x22, 0x5a, 0x24, 0x8a, 0x4c, 0xb8, 0x54, 0x24, 0xe2, 0xe2, 0xa2, 0xbe, 0x61, 0x98,
	0x6b, 0x65, 0x7e, 0xc0, 0xa5, 0xda, 0xe7, 0xa2, 0xd4, 0x1d, 0xa1, 0x85, 0xd2, 0x0c, 0xdc, 0x55,
	0xb3, 0x8b, 0x2f, 0xbd, 0xca, 0x1c, 0xae, 0xb2, 0x2e, 0xef, 0x80, 0x49, 0xa5, 0x8f, 0xef, 0x89,
	0x2d, 0xaa, 0x5e, 0x61, 0xbf, 0xa2, 0xe0, 0x5f, 0x1c, 0xb4, 0x11, 0x71, 0xf1, 0x86, 0x0a, 0x3d,
	0x25, 0x06, 0x99, 0x22, 0x01, 0x08, 0x45, 0x42, 0x50, 0x94, 0x25, 0xd2, 0xc5, 0x2d, 0xa7, 0xbd,
	0xbc, 0x7d, 0xec, 0x5d, 0x67, 0x8e, 0xde, 0x8d, 0x56, 0xe1, 0xed, 0x5b, 0x74, 0xdf, 0x90, 0xfb,
	0x20, 0xd4, 0xae, 0xe5, 0xfa, 0x6e, 0x74, 0x4d, 0x06, 0xff, 0xe6, 0xa0, 0x4d, 0x09, 0x8a, 0x04,
	0x85, 0x10, 0xa6, 0x9d, 0x2b, 0xba, 0x6a, 0x9a, 0xbd, 0x8f, 0xfe, 0x6f, 0x57, 0x23, 0x50, 0x7d,
	0x4b, 0xbf, 0xdc, 0xd8, 0x86, 0xbc, 0x3e, 0x89, 0xbf, 0x41, 0x8f, 0x72, 0x01, 0xe6, 0x69, 0x13,
	0x98, 0x2a, 0x10, 0x19, 0x4d, 0xde, 0x75, 0x93, 0x35, 0xf3, 0xa8, 0x1f, 0x56, 0x35, 0x7b, 0x65,
	0xc9, 0x5b, 0xe7, 0x78, 0x89, 0x6a, 0xa5, 0x03, 0x4b, 0xf7, 0x41, 0x6b, 0xae, 0xbd, 0xb8, 0xfd,
	0xcc, 0xbb, 0xe4, 0xcd, 0x57, 0xee, 0xe8, 0xb8, 0xac, 0x7a, 0x6e, 0x8b, 0xfa, 0x3c, 0x8b, 0x58,
	0xec, 0x5f, 0xd0, 0xd6, 0x7f, 0x77, 0xd0, 0xc6, 0x0d, 0xfb, 0xc2, 0x5f, 0xa0, 0x05, 0x59, 0x8c,
	0xf5, 0x4f, 0x86, 0x31, 0xeb, 0x9b, 0x3d, 0xaf, 0x2a, 0xc5, 0x18, 0xcd, 0xeb, 0xc1, 0x1b, 0x0f,
	0xaf, 0xf9, 0x66, 0x8d, 0xd7, 0xd0, 0xdd, 0x60, 0x42, 0x59, 0x66, 0x5c, 0xbb, 0xe6, 0xdb, 0x0b,
	0x6d, 0xf6, 0x61, 0x66, 0xad, 0xb9, 0xe6, 0xeb, 0xa5, 0x8e, 0x14, 0x82, 0x19, 0x87, 0xad, 0xf9,
	0x7a, 0xf9, 0xe4, 0x1c, 0xb9, 0xd7, 0x1d, 0x08, 0x5c, 0x47, 0xb5, 0x51, 0xef, 0x70, 0x78, 0x32,
	0x7c, 0xbd, 0xd7, 0xf8, 0x00, 0x37, 0x50, 0x7d, 0xff, 0xc8, 0x7f, 0xd1, 0xf3, 0x77, 0xc9, 0xd1,
	0xe1, 0xc1, 0xab, 0x86, 0x83, 0x31, 0x5a, 0xee, 0x1d, 0x1f, 0xef, 0x1d, 0xee, 0x92, 0x32, 0xd1,
	0xb8, 0xa3, 0xab, 0x2a, 0x0d, 0x19, 0xed, 0x9d, 0x34, 0xe6, 0xf0, 0x47, 0xa8, 0xd9, 0x3b, 0x78,
	0xd1, 0x7b, 0x35, 0x22, 0xef, 0xc9, 0xe7, 0x77, 0xf6, 0xff, 0xfc, 0x77, 0xde, 0xf9, 0xe3, 0x9f,
	0xc7, 0xce, 0xeb, 0x67, 0xb3, 0x7d, 0x2b, 0xe4, 0xa7, 0xf1, 0x15, 0xdf, 0x0b, 0xe3, 0x7b, 0x66,
	0x5a, 0x4f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x89, 0x92, 0x61, 0x72, 0x08, 0x00, 0x00,
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
	if this.ForwardClientCertDetails != that1.ForwardClientCertDetails {
		return false
	}
	if !this.SetCurrentClientCertDetails.Equal(that1.SetCurrentClientCertDetails) {
		return false
	}
	if this.PreserveExternalRequestId != that1.PreserveExternalRequestId {
		return false
	}
	if len(this.Upgrades) != len(that1.Upgrades) {
		return false
	}
	for i := range this.Upgrades {
		if !this.Upgrades[i].Equal(that1.Upgrades[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings_SetCurrentClientCertDetails)
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
	if !this.Subject.Equal(that1.Subject) {
		return false
	}
	if this.Cert != that1.Cert {
		return false
	}
	if this.Chain != that1.Chain {
		return false
	}
	if this.Dns != that1.Dns {
		return false
	}
	if this.Uri != that1.Uri {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
