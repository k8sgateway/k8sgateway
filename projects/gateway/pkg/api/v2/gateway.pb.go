// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v2/gateway.proto

package v2

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type Gateway struct {
	// if set to false, only use virtual services with no ssl configured.
	// if set to true, only use virtual services with ssl configured.
	Ssl bool `protobuf:"varint,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways in a namespace
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level plugin configuration for all routes on the gateway
	Plugins *v1.ListenerPlugins `protobuf:"bytes,5,opt,name=plugins,proto3" json:"plugins,omitempty"`
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
	GatewayType          isGateway_GatewayType `protobuf_oneof:"GatewayType"`
	GatewayProxyName     string                `protobuf:"bytes,11,opt,name=gateway_proxy_name,json=gatewayProxyName,proto3" json:"gateway_proxy_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ebb625f16d49f4, []int{0}
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
	HttpGateway *HttpGateway `protobuf:"bytes,9,opt,name=http_gateway,json=httpGateway,proto3,oneof"`
}
type Gateway_TcpGateway struct {
	TcpGateway *TcpGateway `protobuf:"bytes,10,opt,name=tcp_gateway,json=tcpGateway,proto3,oneof"`
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

func (m *Gateway) GetPlugins() *v1.ListenerPlugins {
	if m != nil {
		return m.Plugins
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

func (m *Gateway) GetGatewayProxyName() string {
	if m != nil {
		return m.GatewayProxyName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Gateway) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Gateway_OneofMarshaler, _Gateway_OneofUnmarshaler, _Gateway_OneofSizer, []interface{}{
		(*Gateway_HttpGateway)(nil),
		(*Gateway_TcpGateway)(nil),
	}
}

func _Gateway_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Gateway)
	// GatewayType
	switch x := m.GatewayType.(type) {
	case *Gateway_HttpGateway:
		_ = b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpGateway); err != nil {
			return err
		}
	case *Gateway_TcpGateway:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TcpGateway); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Gateway.GatewayType has unexpected type %T", x)
	}
	return nil
}

func _Gateway_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Gateway)
	switch tag {
	case 9: // GatewayType.http_gateway
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpGateway)
		err := b.DecodeMessage(msg)
		m.GatewayType = &Gateway_HttpGateway{msg}
		return true, err
	case 10: // GatewayType.tcp_gateway
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TcpGateway)
		err := b.DecodeMessage(msg)
		m.GatewayType = &Gateway_TcpGateway{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Gateway_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Gateway)
	// GatewayType
	switch x := m.GatewayType.(type) {
	case *Gateway_HttpGateway:
		s := proto.Size(x.HttpGateway)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Gateway_TcpGateway:
		s := proto.Size(x.TcpGateway)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HttpGateway struct {
	// names of the the virtual services, which contain the actual routes for the gateway
	// if the list is empty, all virtual services will apply to this gateway (with accordance to tls flag above).
	VirtualServices []core.ResourceRef `protobuf:"bytes,1,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services"`
	// http gateway configuration
	Plugins              *v1.HttpListenerPlugins `protobuf:"bytes,8,opt,name=plugins,proto3" json:"plugins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HttpGateway) Reset()         { *m = HttpGateway{} }
func (m *HttpGateway) String() string { return proto.CompactTextString(m) }
func (*HttpGateway) ProtoMessage()    {}
func (*HttpGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ebb625f16d49f4, []int{1}
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

func (m *HttpGateway) GetPlugins() *v1.HttpListenerPlugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

type TcpGateway struct {
	// Name of the destinations the gateway can route to
	Destinations []*v1.TcpHost `protobuf:"bytes,1,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// tcp gateway configuration
	Plugins              *v1.TcpListenerPlugins `protobuf:"bytes,8,opt,name=plugins,proto3" json:"plugins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TcpGateway) Reset()         { *m = TcpGateway{} }
func (m *TcpGateway) String() string { return proto.CompactTextString(m) }
func (*TcpGateway) ProtoMessage()    {}
func (*TcpGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ebb625f16d49f4, []int{2}
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

func (m *TcpGateway) GetPlugins() *v1.TcpListenerPlugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.v2.Gateway")
	proto.RegisterType((*HttpGateway)(nil), "gateway.solo.io.v2.HttpGateway")
	proto.RegisterType((*TcpGateway)(nil), "gateway.solo.io.v2.TcpGateway")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v2/gateway.proto", fileDescriptor_07ebb625f16d49f4)
}

var fileDescriptor_07ebb625f16d49f4 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x4e, 0xdb, 0x4c,
	0x14, 0xc6, 0x3f, 0xf9, 0x21, 0x8c, 0x41, 0xd0, 0x11, 0x45, 0x26, 0x55, 0x21, 0x64, 0x95, 0x05,
	0xb5, 0x45, 0x58, 0x14, 0x52, 0x75, 0x41, 0x54, 0x09, 0x54, 0xb5, 0x55, 0x64, 0x50, 0x17, 0xdd,
	0x44, 0x13, 0x7b, 0x62, 0xa6, 0x38, 0x7e, 0xd6, 0xcc, 0x73, 0x28, 0x5b, 0x7a, 0x87, 0x9e, 0xa1,
	0x57, 0xe8, 0x0d, 0x7a, 0x0a, 0x16, 0xbd, 0x01, 0x3d, 0x41, 0xe5, 0xf1, 0x38, 0x21, 0x05, 0x54,
	0x58, 0xc5, 0x6f, 0xde, 0xfb, 0xbe, 0xf9, 0xfc, 0xbd, 0x2f, 0x26, 0xaf, 0x23, 0x81, 0xa7, 0x59,
	0xdf, 0x0d, 0x60, 0xe8, 0x29, 0x88, 0xe1, 0x85, 0x00, 0x2f, 0x8a, 0x01, 0xbc, 0x54, 0xc2, 0x67,
	0x1e, 0xa0, 0xf2, 0x22, 0x86, 0xfc, 0x9c, 0x5d, 0x78, 0x2c, 0x15, 0xde, 0xa8, 0x55, 0x96, 0x6e,
	0x2a, 0x01, 0x81, 0xd2, 0xb2, 0xcc, 0xb1, 0xae, 0x00, 0x77, 0xd4, 0xaa, 0xad, 0x46, 0x10, 0x81,
	0x6e, 0x7b, 0xf9, 0x53, 0x31, 0x59, 0xdb, 0x88, 0x00, 0xa2, 0x98, 0x7b, 0xba, 0xea, 0x67, 0x03,
	0xef, 0x5c, 0xb2, 0x34, 0xe5, 0x52, 0x99, 0xfe, 0xce, 0x1d, 0x42, 0xf4, 0xef, 0x99, 0xc0, 0xe2,
	0xee, 0x1d, 0x6f, 0xc8, 0x91, 0x85, 0x0c, 0x99, 0x81, 0x78, 0x0f, 0x80, 0x28, 0x64, 0x98, 0x95,
	0x77, 0x6c, 0x3f, 0x00, 0x20, 0xf9, 0xe0, 0x11, 0x8a, 0xca, 0xda, 0x40, 0xf6, 0xfe, 0xed, 0x66,
	0x5e, 0x19, 0x70, 0x2a, 0xe1, 0x8b, 0x31, 0xb2, 0xd6, 0x7e, 0x1c, 0x32, 0xce, 0x22, 0x91, 0x98,
	0xd7, 0x6a, 0xfc, 0xa8, 0x90, 0xf9, 0xc3, 0x62, 0x0f, 0x74, 0x85, 0xcc, 0x2a, 0x15, 0x3b, 0x56,
	0xdd, 0x6a, 0x56, 0xfd, 0xfc, 0x91, 0x6e, 0x91, 0xc5, 0xbe, 0x48, 0xc2, 0x1e, 0x0b, 0x43, 0xc9,
	0x95, 0x72, 0x66, 0xeb, 0x56, 0x73, 0xc1, 0xb7, 0xf3, 0xb3, 0x83, 0xe2, 0x88, 0x3e, 0x23, 0x0b,
	0x7a, 0x24, 0x05, 0x89, 0x4e, 0xa5, 0x6e, 0x35, 0x97, 0xfc, 0x6a, 0x7e, 0xd0, 0x05, 0x89, 0xf4,
	0x25, 0x99, 0x37, 0xd7, 0x39, 0xff, 0xd7, 0xad, 0xa6, 0xdd, 0x7a, 0xee, 0xe6, 0x52, 0xc6, 0x1b,
	0x7f, 0x27, 0x14, 0xf2, 0x84, 0xcb, 0x6e, 0x31, 0xe4, 0x97, 0xd3, 0xf4, 0x90, 0xcc, 0x15, 0xee,
	0x3b, 0x73, 0x1a, 0xb7, 0xea, 0x06, 0x20, 0xf9, 0x18, 0x77, 0xac, 0x7b, 0x9d, 0xf5, 0x9f, 0x57,
	0x9b, 0x33, 0xbf, 0xaf, 0x36, 0x9f, 0x20, 0x57, 0x18, 0x8a, 0xc1, 0xa0, 0xdd, 0x10, 0x51, 0x02,
	0x92, 0x37, 0x7c, 0x03, 0xa7, 0x7b, 0xa4, 0x5a, 0x6e, 0xde, 0x99, 0xd7, 0x54, 0x6b, 0xd3, 0x54,
	0xef, 0x4d, 0xb7, 0x53, 0xc9, 0xc9, 0xfc, 0xf1, 0x34, 0xed, 0x90, 0xe5, 0x4c, 0xf1, 0x9e, 0x36,
	0xba, 0xa7, 0xcd, 0x72, 0xaa, 0x9a, 0xa0, 0xe6, 0x16, 0x71, 0x74, 0xcb, 0x38, 0xba, 0x1d, 0x80,
	0xf8, 0x23, 0x8b, 0x33, 0xee, 0x2f, 0x65, 0x8a, 0x77, 0x73, 0x44, 0x57, 0x47, 0xfc, 0x0d, 0x59,
	0x3c, 0x45, 0x4c, 0x7b, 0x26, 0xe9, 0xce, 0x82, 0x26, 0xd8, 0x74, 0x6f, 0x27, 0xdf, 0x3d, 0x42,
	0x4c, 0xcd, 0x22, 0x8e, 0x66, 0x7c, 0xfb, 0x74, 0x52, 0xd2, 0x03, 0x62, 0x63, 0x30, 0x21, 0x21,
	0x9a, 0x64, 0xe3, 0x2e, 0x92, 0x93, 0xe0, 0x06, 0x07, 0xc1, 0x71, 0x45, 0xb7, 0x49, 0xf9, 0x6f,
	0x33, 0x2f, 0x94, 0xb0, 0x21, 0x77, 0x6c, 0xbd, 0xce, 0x15, 0xd3, 0xd1, 0xba, 0x3f, 0xb0, 0x21,
	0x6f, 0xaf, 0x5d, 0x5e, 0x57, 0x2a, 0xe4, 0xbf, 0xe8, 0xfc, 0xf2, 0xba, 0x42, 0x68, 0xd5, 0xf4,
	0x55, 0x67, 0x89, 0xd8, 0x86, 0xf0, 0xe4, 0x22, 0xe5, 0x8d, 0x6f, 0x16, 0xb1, 0x6f, 0xc8, 0xa6,
	0x6f, 0xc9, 0xca, 0x48, 0x48, 0xcc, 0x58, 0xdc, 0x53, 0x5c, 0x8e, 0x44, 0xc0, 0x95, 0x63, 0xd5,
	0x67, 0x9b, 0x76, 0x6b, 0x7d, 0xda, 0x73, 0x9f, 0x2b, 0xc8, 0x64, 0xc0, 0x7d, 0x3e, 0x30, 0xb6,
	0x2f, 0x1b, 0xe0, 0xb1, 0xc1, 0xd1, 0x57, 0x93, 0xe4, 0x14, 0xae, 0x6f, 0x4d, 0x27, 0x27, 0xbf,
	0xf7, 0xbe, 0xf4, 0x34, 0xbe, 0x5a, 0x84, 0x4c, 0xac, 0xa0, 0xfb, 0x64, 0x31, 0xe4, 0x0a, 0x45,
	0xc2, 0x50, 0x40, 0x52, 0x6a, 0x7a, 0x3a, 0x4d, 0x78, 0x12, 0xa4, 0x47, 0xa0, 0xd0, 0x9f, 0x1a,
	0xa5, 0xed, 0xbf, 0x65, 0xd4, 0x6f, 0xa1, 0xee, 0x53, 0xd1, 0xd9, 0xff, 0xfe, 0x6b, 0xc3, 0xfa,
	0xb4, 0xfb, 0xe0, 0x8f, 0x64, 0x7a, 0x16, 0x99, 0x0f, 0x65, 0x7f, 0x4e, 0x47, 0x6b, 0xf7, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xd9, 0xdc, 0xda, 0x62, 0x05, 0x00, 0x00,
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
	if !this.Plugins.Equal(that1.Plugins) {
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
	if this.GatewayProxyName != that1.GatewayProxyName {
		return false
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
	if !this.Plugins.Equal(that1.Plugins) {
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
	if !this.Plugins.Equal(that1.Plugins) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
