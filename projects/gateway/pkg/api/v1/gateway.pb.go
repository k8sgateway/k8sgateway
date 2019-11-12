// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//A Gateway describes a single Listener (bind address:port)
//and the routing configuration to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type Gateway struct {
	// if set to false, only use virtual services without ssl configured.
	// if set to true, only use virtual services with ssl configured.
	Ssl bool `protobuf:"varint,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways for a single proxy
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level optional configuration for all routes on the gateway
	Options *v1.ListenerOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Enable ProxyProtocol support for this listener
	UseProxyProto *types.BoolValue `protobuf:"bytes,8,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	// The type of gateway being created
	// HttpGateway creates a listener with an http_connection_manager
	// TcpGateway creates a listener with a tcp proxy filter
	//
	// Types that are valid to be assigned to GatewayType:
	//	*Gateway_HttpGateway
	//	*Gateway_TcpGateway
	GatewayType isGateway_GatewayType `protobuf_oneof:"GatewayType"`
	//
	// Names of the [`Proxy`](https://gloo.solo.io/api/github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto.sk/)
	// resources to generate from this gateway. If other gateways exist which point to the same proxy,
	// Gloo will join them together.
	//
	// Proxies have a one-to-many relationship with Envoy bootstrap configuration.
	// In order to connect to Gloo, the Envoy bootstrap configuration sets a `role` in
	// the [node metadata](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/base.proto#envoy-api-msg-core-node)
	// Envoy instances announce their `role` to Gloo, which maps to the  `{{ .Namespace }}~{{ .Name }}`
	// of the Proxy resource.
	//
	// The template for this value can be seen in the [Gloo Helm chart](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/templates/9-gateway-proxy-configmap.yaml#L22)
	//
	// Note: this field also accepts fields written in camel-case. They will be converted
	// to kebab-case in the Proxy name. This allows use of the [Gateway Name Helm value](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/values-gateway-template.yaml#L47)
	// for this field
	//
	// Defaults to `["gateway-proxy"]`
	ProxyNames           []string `protobuf:"bytes,12,rep,name=proxy_names,json=proxyNames,proto3" json:"proxy_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{0}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (m *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(m, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

type isGateway_GatewayType interface {
	isGateway_GatewayType()
	Equal(interface{}) bool
}

type Gateway_HttpGateway struct {
	HttpGateway *HttpGateway `protobuf:"bytes,9,opt,name=http_gateway,json=httpGateway,proto3,oneof" json:"http_gateway,omitempty"`
}
type Gateway_TcpGateway struct {
	TcpGateway *TcpGateway `protobuf:"bytes,10,opt,name=tcp_gateway,json=tcpGateway,proto3,oneof" json:"tcp_gateway,omitempty"`
}

func (*Gateway_HttpGateway) isGateway_GatewayType() {}
func (*Gateway_TcpGateway) isGateway_GatewayType()  {}

func (m *Gateway) GetGatewayType() isGateway_GatewayType {
	if m != nil {
		return m.GatewayType
	}
	return nil
}

func (m *Gateway) GetSsl() bool {
	if m != nil {
		return m.Ssl
	}
	return false
}

func (m *Gateway) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *Gateway) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *Gateway) GetOptions() *v1.ListenerOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Gateway) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Gateway) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Gateway) GetUseProxyProto() *types.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *Gateway) GetHttpGateway() *HttpGateway {
	if x, ok := m.GetGatewayType().(*Gateway_HttpGateway); ok {
		return x.HttpGateway
	}
	return nil
}

func (m *Gateway) GetTcpGateway() *TcpGateway {
	if x, ok := m.GetGatewayType().(*Gateway_TcpGateway); ok {
		return x.TcpGateway
	}
	return nil
}

func (m *Gateway) GetProxyNames() []string {
	if m != nil {
		return m.ProxyNames
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Gateway) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Gateway_HttpGateway)(nil),
		(*Gateway_TcpGateway)(nil),
	}
}

type HttpGateway struct {
	// names of the the virtual services, which contain the actual routes for the gateway
	// if the list is empty, all virtual services will apply to this gateway (with accordance to tls flag above).
	VirtualServices []core.ResourceRef `protobuf:"bytes,1,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services"`
	// Select virtual services by their label. This will apply only to virtual services
	// in the same namespace as the gateway resource.
	// only one of `virtualServices` or `virtualServiceSelector` should be provided
	VirtualServiceSelector map[string]string `protobuf:"bytes,2,rep,name=virtual_service_selector,json=virtualServiceSelector,proto3" json:"virtual_service_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// http gateway configuration
	Options              *v1.HttpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HttpGateway) Reset()         { *m = HttpGateway{} }
func (m *HttpGateway) String() string { return proto.CompactTextString(m) }
func (*HttpGateway) ProtoMessage()    {}
func (*HttpGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{1}
}
func (m *HttpGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpGateway.Unmarshal(m, b)
}
func (m *HttpGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpGateway.Marshal(b, m, deterministic)
}
func (m *HttpGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpGateway.Merge(m, src)
}
func (m *HttpGateway) XXX_Size() int {
	return xxx_messageInfo_HttpGateway.Size(m)
}
func (m *HttpGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpGateway.DiscardUnknown(m)
}

var xxx_messageInfo_HttpGateway proto.InternalMessageInfo

func (m *HttpGateway) GetVirtualServices() []core.ResourceRef {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *HttpGateway) GetVirtualServiceSelector() map[string]string {
	if m != nil {
		return m.VirtualServiceSelector
	}
	return nil
}

func (m *HttpGateway) GetOptions() *v1.HttpListenerOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type TcpGateway struct {
	// Name of the destinations the gateway can route to
	Destinations []*v1.TcpHost `protobuf:"bytes,1,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// tcp gateway configuration
	Options              *v1.TcpListenerOptions `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TcpGateway) Reset()         { *m = TcpGateway{} }
func (m *TcpGateway) String() string { return proto.CompactTextString(m) }
func (*TcpGateway) ProtoMessage()    {}
func (*TcpGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{2}
}
func (m *TcpGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpGateway.Unmarshal(m, b)
}
func (m *TcpGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpGateway.Marshal(b, m, deterministic)
}
func (m *TcpGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpGateway.Merge(m, src)
}
func (m *TcpGateway) XXX_Size() int {
	return xxx_messageInfo_TcpGateway.Size(m)
}
func (m *TcpGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpGateway.DiscardUnknown(m)
}

var xxx_messageInfo_TcpGateway proto.InternalMessageInfo

func (m *TcpGateway) GetDestinations() []*v1.TcpHost {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *TcpGateway) GetOptions() *v1.TcpListenerOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.Gateway")
	proto.RegisterType((*HttpGateway)(nil), "gateway.solo.io.HttpGateway")
	proto.RegisterMapType((map[string]string)(nil), "gateway.solo.io.HttpGateway.VirtualServiceSelectorEntry")
	proto.RegisterType((*TcpGateway)(nil), "gateway.solo.io.TcpGateway")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto", fileDescriptor_30f7529f6633771c)
}

var fileDescriptor_30f7529f6633771c = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x49, 0x80, 0x64, 0x1c, 0x04, 0xdf, 0x88, 0x0f, 0x99, 0xd0, 0x82, 0xc9, 0x2a, 0x8b,
	0xd6, 0x16, 0xb0, 0x68, 0x48, 0xd5, 0x4a, 0x44, 0xaa, 0xa0, 0x55, 0x7f, 0xd0, 0x80, 0x58, 0x74,
	0x13, 0x39, 0xce, 0xc4, 0xb8, 0x18, 0x8f, 0x35, 0x73, 0x1d, 0x9a, 0x2d, 0x7d, 0x97, 0xaa, 0x8f,
	0xd2, 0x17, 0xe8, 0x96, 0x45, 0xdf, 0x80, 0x3e, 0x41, 0x35, 0xe3, 0x31, 0x21, 0xfc, 0x35, 0xac,
	0x3c, 0xf7, 0xce, 0x3d, 0x67, 0xee, 0x9c, 0x7b, 0xc6, 0xe8, 0x55, 0x10, 0xc2, 0x71, 0xda, 0x75,
	0x7c, 0x76, 0xea, 0x0a, 0x16, 0xb1, 0xe7, 0x21, 0x73, 0x83, 0x88, 0x31, 0x37, 0xe1, 0xec, 0x0b,
	0xf5, 0x41, 0xb8, 0x81, 0x07, 0xf4, 0xcc, 0x1b, 0xba, 0x5e, 0x12, 0xba, 0x83, 0x8d, 0x3c, 0x74,
	0x12, 0xce, 0x80, 0xe1, 0xf9, 0x3c, 0x94, 0x58, 0x27, 0x64, 0xb5, 0xc5, 0x80, 0x05, 0x4c, 0xed,
	0xb9, 0x72, 0x95, 0x95, 0xd5, 0x56, 0x03, 0xc6, 0x82, 0x88, 0xba, 0x2a, 0xea, 0xa6, 0x7d, 0xf7,
	0x8c, 0x7b, 0x49, 0x42, 0xb9, 0xd0, 0xfb, 0x1b, 0x77, 0x74, 0xa1, 0xbe, 0x27, 0x21, 0xe4, 0x07,
	0x9f, 0x52, 0xf0, 0x7a, 0x1e, 0x78, 0x1a, 0xe2, 0x4e, 0x00, 0x11, 0xe0, 0x41, 0x9a, 0x9f, 0xf1,
	0x6c, 0x02, 0x00, 0xa7, 0xfd, 0x47, 0x74, 0x94, 0xc7, 0x1a, 0xd2, 0xfc, 0xb7, 0x94, 0x32, 0xd2,
	0xe0, 0x84, 0xb3, 0xaf, 0x5a, 0xc5, 0x5a, 0xeb, 0x51, 0x48, 0x96, 0x40, 0xc8, 0x62, 0x7d, 0xad,
	0xfa, 0xf7, 0x12, 0x9a, 0xdd, 0xcd, 0x86, 0x80, 0x17, 0x50, 0x51, 0x88, 0xc8, 0x32, 0x6c, 0xa3,
	0x51, 0x26, 0x72, 0x89, 0xd7, 0x51, 0xb5, 0x1b, 0xc6, 0xbd, 0x8e, 0xd7, 0xeb, 0x71, 0x2a, 0x84,
	0x55, 0xb4, 0x8d, 0x46, 0x85, 0x98, 0x32, 0xb7, 0x93, 0xa5, 0xf0, 0x0a, 0xaa, 0xa8, 0x92, 0x84,
	0x71, 0xb0, 0x4a, 0xb6, 0xd1, 0x98, 0x23, 0x65, 0x99, 0xd8, 0x67, 0x1c, 0xf0, 0x0b, 0x34, 0xab,
	0x8f, 0xb3, 0xa6, 0x6d, 0xa3, 0x61, 0x6e, 0x3e, 0x75, 0x64, 0x2b, 0xf9, 0xb8, 0x9d, 0xf7, 0xa1,
	0x00, 0x1a, 0x53, 0xfe, 0x29, 0x2b, 0x22, 0x79, 0x35, 0xde, 0x45, 0x33, 0x99, 0xfa, 0xd6, 0x8c,
	0xc2, 0x2d, 0x3a, 0x3e, 0xe3, 0xf4, 0x0a, 0x77, 0xa0, 0xf6, 0xda, 0xcb, 0x3f, 0x2f, 0xd6, 0xa6,
	0xfe, 0x5c, 0xac, 0xfd, 0x07, 0x54, 0x40, 0x2f, 0xec, 0xf7, 0x5b, 0xf5, 0x30, 0x88, 0x19, 0xa7,
	0x75, 0xa2, 0xe1, 0xb8, 0x89, 0xca, 0xf9, 0xe4, 0xad, 0x59, 0x45, 0xb5, 0x34, 0x4e, 0xf5, 0x41,
	0xef, 0xb6, 0x4b, 0x92, 0x8c, 0x5c, 0x55, 0xe3, 0x36, 0x9a, 0x4f, 0x05, 0xed, 0x28, 0xa1, 0x3b,
	0x4a, 0x2c, 0xab, 0xac, 0x08, 0x6a, 0x4e, 0x66, 0x47, 0x27, 0xb7, 0xa3, 0xd3, 0x66, 0x2c, 0x3a,
	0xf2, 0xa2, 0x94, 0x92, 0xb9, 0x54, 0xd0, 0x7d, 0x89, 0xd8, 0x57, 0xfe, 0xde, 0x41, 0xd5, 0x63,
	0x80, 0xa4, 0xa3, 0x6d, 0x6e, 0x55, 0x14, 0xc1, 0x13, 0xe7, 0x86, 0xed, 0x9d, 0x3d, 0x80, 0x44,
	0x4f, 0x61, 0x6f, 0x8a, 0x98, 0xc7, 0xa3, 0x10, 0xbf, 0x46, 0x26, 0xf8, 0x23, 0x06, 0xa4, 0x18,
	0x56, 0x6e, 0x31, 0x1c, 0xfa, 0xd7, 0x08, 0x10, 0x5c, 0x45, 0x78, 0x0d, 0x99, 0xd9, 0x15, 0x62,
	0xef, 0x94, 0x0a, 0xab, 0x6a, 0x17, 0x1b, 0x15, 0x82, 0x54, 0xea, 0xa3, 0xcc, 0xb4, 0x96, 0xce,
	0x2f, 0x4b, 0x25, 0x54, 0x08, 0xce, 0xce, 0x2f, 0x4b, 0x08, 0x97, 0x35, 0xb1, 0x68, 0xcf, 0x21,
	0x53, 0x73, 0x1c, 0x0e, 0x13, 0x5a, 0xff, 0x55, 0x40, 0xe6, 0xb5, 0x36, 0xf1, 0x3b, 0xb4, 0x30,
	0x08, 0x39, 0xa4, 0x5e, 0xd4, 0x11, 0x94, 0x0f, 0x42, 0x9f, 0x0a, 0xcb, 0xb0, 0x8b, 0x0d, 0x73,
	0x73, 0x79, 0x5c, 0x60, 0x42, 0x05, 0x4b, 0xb9, 0x4f, 0x09, 0xed, 0x6b, 0x8d, 0xe7, 0x35, 0xf0,
	0x40, 0xe3, 0x30, 0x47, 0xd6, 0x0d, 0xae, 0x8e, 0xa0, 0x11, 0xf5, 0x81, 0x71, 0xab, 0xa0, 0x38,
	0x9b, 0x0f, 0x49, 0xe6, 0x1c, 0x8d, 0xf1, 0x1d, 0x68, 0xe8, 0x9b, 0x18, 0xf8, 0x90, 0x2c, 0x0d,
	0xee, 0xdc, 0xc4, 0x2f, 0x47, 0xd6, 0xcc, 0xc6, 0xba, 0x3e, 0x6e, 0x4d, 0xc9, 0x7f, 0x9f, 0x3d,
	0x6b, 0x6f, 0xd1, 0xca, 0x03, 0x67, 0xca, 0x87, 0x74, 0x42, 0x87, 0xea, 0x21, 0x55, 0x88, 0x5c,
	0xe2, 0x45, 0x34, 0x3d, 0x90, 0x06, 0xb1, 0x0a, 0x2a, 0x97, 0x05, 0xad, 0x42, 0xd3, 0xa8, 0x7f,
	0x33, 0x10, 0x1a, 0x0d, 0x0f, 0x6f, 0xa3, 0x6a, 0x8f, 0x0a, 0x08, 0x63, 0x2f, 0xeb, 0x2d, 0x93,
	0xf4, 0xff, 0xf1, 0xde, 0x0e, 0xfd, 0x64, 0x8f, 0x09, 0x20, 0x63, 0xa5, 0xb8, 0x75, 0xf3, 0x46,
	0xf6, 0x2d, 0xd4, 0x7d, 0x17, 0x6a, 0x6f, 0xff, 0xf8, 0xbd, 0x6a, 0x7c, 0xde, 0x9a, 0xf8, 0x6f,
	0x9e, 0x9c, 0x04, 0xfa, 0x7f, 0xd2, 0x9d, 0x51, 0xcf, 0x60, 0xeb, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x58, 0xa0, 0xfa, 0x0b, 0x06, 0x00, 0x00,
}

func (this *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
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
	if this.Ssl != that1.Ssl {
		return false
	}
	if this.BindAddress != that1.BindAddress {
		return false
	}
	if this.BindPort != that1.BindPort {
		return false
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.UseProxyProto.Equal(that1.UseProxyProto) {
		return false
	}
	if that1.GatewayType == nil {
		if this.GatewayType != nil {
			return false
		}
	} else if this.GatewayType == nil {
		return false
	} else if !this.GatewayType.Equal(that1.GatewayType) {
		return false
	}
	if len(this.ProxyNames) != len(that1.ProxyNames) {
		return false
	}
	for i := range this.ProxyNames {
		if this.ProxyNames[i] != that1.ProxyNames[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Gateway_HttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway_HttpGateway)
	if !ok {
		that2, ok := that.(Gateway_HttpGateway)
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
	if !this.HttpGateway.Equal(that1.HttpGateway) {
		return false
	}
	return true
}
func (this *Gateway_TcpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway_TcpGateway)
	if !ok {
		that2, ok := that.(Gateway_TcpGateway)
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
	if !this.TcpGateway.Equal(that1.TcpGateway) {
		return false
	}
	return true
}
func (this *HttpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpGateway)
	if !ok {
		that2, ok := that.(HttpGateway)
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
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(&that1.VirtualServices[i]) {
			return false
		}
	}
	if len(this.VirtualServiceSelector) != len(that1.VirtualServiceSelector) {
		return false
	}
	for i := range this.VirtualServiceSelector {
		if this.VirtualServiceSelector[i] != that1.VirtualServiceSelector[i] {
			return false
		}
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TcpGateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TcpGateway)
	if !ok {
		that2, ok := that.(TcpGateway)
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
	if len(this.Destinations) != len(that1.Destinations) {
		return false
	}
	for i := range this.Destinations {
		if !this.Destinations[i].Equal(that1.Destinations[i]) {
			return false
		}
	}
	if !this.Options.Equal(that1.Options) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
