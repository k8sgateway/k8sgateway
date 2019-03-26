// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/aws"
	azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/azure"
	consul "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/consul"
	faultinjection "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/faultinjection"
	grpc "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/grpc"
	grpc_web "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/grpc_web"
	hcm "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/hcm"
	kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"
	rest "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/rest"
	retries "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/retries"
	static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/static"
	transformation "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/transformation"
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

// Plugin-specific configuration that lives on listeners
// Each ListenerPlugin object contains configuration for a specific plugin
// Note to developers: new Listener Plugins must be added to this struct
// to be usable by Gloo.
type ListenerPlugins struct {
	GrpcWeb                       *grpc_web.GrpcWeb                  `protobuf:"bytes,1,opt,name=grpc_web,json=grpcWeb,proto3" json:"grpc_web,omitempty"`
	HttpConnectionManagerSettings *hcm.HttpConnectionManagerSettings `protobuf:"bytes,2,opt,name=http_connection_manager_settings,json=httpConnectionManagerSettings,proto3" json:"http_connection_manager_settings,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                           `json:"-"`
	XXX_unrecognized              []byte                             `json:"-"`
	XXX_sizecache                 int32                              `json:"-"`
}

func (m *ListenerPlugins) Reset()         { *m = ListenerPlugins{} }
func (m *ListenerPlugins) String() string { return proto.CompactTextString(m) }
func (*ListenerPlugins) ProtoMessage()    {}
func (*ListenerPlugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{0}
}
func (m *ListenerPlugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerPlugins.Unmarshal(m, b)
}
func (m *ListenerPlugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerPlugins.Marshal(b, m, deterministic)
}
func (m *ListenerPlugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerPlugins.Merge(m, src)
}
func (m *ListenerPlugins) XXX_Size() int {
	return xxx_messageInfo_ListenerPlugins.Size(m)
}
func (m *ListenerPlugins) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerPlugins.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerPlugins proto.InternalMessageInfo

func (m *ListenerPlugins) GetGrpcWeb() *grpc_web.GrpcWeb {
	if m != nil {
		return m.GrpcWeb
	}
	return nil
}

func (m *ListenerPlugins) GetHttpConnectionManagerSettings() *hcm.HttpConnectionManagerSettings {
	if m != nil {
		return m.HttpConnectionManagerSettings
	}
	return nil
}

// Plugin-specific configuration that lives on virtual hosts
// Each VirtualHostPlugin object contains configuration for a specific plugin
// Note to developers: new Virtual Host Plugins must be added to this struct
// to be usable by Gloo.
type VirtualHostPlugins struct {
	Extensions           *Extensions `protobuf:"bytes,1,opt,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VirtualHostPlugins) Reset()         { *m = VirtualHostPlugins{} }
func (m *VirtualHostPlugins) String() string { return proto.CompactTextString(m) }
func (*VirtualHostPlugins) ProtoMessage()    {}
func (*VirtualHostPlugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{1}
}
func (m *VirtualHostPlugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHostPlugins.Unmarshal(m, b)
}
func (m *VirtualHostPlugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHostPlugins.Marshal(b, m, deterministic)
}
func (m *VirtualHostPlugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHostPlugins.Merge(m, src)
}
func (m *VirtualHostPlugins) XXX_Size() int {
	return xxx_messageInfo_VirtualHostPlugins.Size(m)
}
func (m *VirtualHostPlugins) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHostPlugins.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHostPlugins proto.InternalMessageInfo

func (m *VirtualHostPlugins) GetExtensions() *Extensions {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// Plugin-specific configuration that lives on routes
// Each RoutePlugin object contains configuration for a specific plugin
// Note to developers: new Route Plugins must be added to this struct
// to be usable by Gloo.
type RoutePlugins struct {
	Transformations      *transformation.RouteTransformations `protobuf:"bytes,1,opt,name=transformations,proto3" json:"transformations,omitempty"`
	Faults               *faultinjection.RouteFaults          `protobuf:"bytes,2,opt,name=faults,proto3" json:"faults,omitempty"`
	PrefixRewrite        *transformation.PrefixRewrite        `protobuf:"bytes,3,opt,name=prefix_rewrite,json=prefixRewrite,proto3" json:"prefix_rewrite,omitempty"`
	Timeout              *time.Duration                       `protobuf:"bytes,4,opt,name=timeout,proto3,stdduration" json:"timeout,omitempty"`
	Retries              *retries.RetryPolicy                 `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	Extensions           *Extensions                          `protobuf:"bytes,6,opt,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *RoutePlugins) Reset()         { *m = RoutePlugins{} }
func (m *RoutePlugins) String() string { return proto.CompactTextString(m) }
func (*RoutePlugins) ProtoMessage()    {}
func (*RoutePlugins) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{2}
}
func (m *RoutePlugins) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutePlugins.Unmarshal(m, b)
}
func (m *RoutePlugins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutePlugins.Marshal(b, m, deterministic)
}
func (m *RoutePlugins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutePlugins.Merge(m, src)
}
func (m *RoutePlugins) XXX_Size() int {
	return xxx_messageInfo_RoutePlugins.Size(m)
}
func (m *RoutePlugins) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutePlugins.DiscardUnknown(m)
}

var xxx_messageInfo_RoutePlugins proto.InternalMessageInfo

func (m *RoutePlugins) GetTransformations() *transformation.RouteTransformations {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *RoutePlugins) GetFaults() *faultinjection.RouteFaults {
	if m != nil {
		return m.Faults
	}
	return nil
}

func (m *RoutePlugins) GetPrefixRewrite() *transformation.PrefixRewrite {
	if m != nil {
		return m.PrefixRewrite
	}
	return nil
}

func (m *RoutePlugins) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RoutePlugins) GetRetries() *retries.RetryPolicy {
	if m != nil {
		return m.Retries
	}
	return nil
}

func (m *RoutePlugins) GetExtensions() *Extensions {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// Configuration for Destinations that are tied to the UpstreamSpec or ServiceSpec on that destination
type DestinationSpec struct {
	// Note to developers: new DestinationSpecs must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to DestinationType:
	//	*DestinationSpec_Aws
	//	*DestinationSpec_Azure
	//	*DestinationSpec_Rest
	//	*DestinationSpec_Grpc
	DestinationType      isDestinationSpec_DestinationType `protobuf_oneof:"destination_type"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DestinationSpec) Reset()         { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()    {}
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{3}
}
func (m *DestinationSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestinationSpec.Unmarshal(m, b)
}
func (m *DestinationSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestinationSpec.Marshal(b, m, deterministic)
}
func (m *DestinationSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestinationSpec.Merge(m, src)
}
func (m *DestinationSpec) XXX_Size() int {
	return xxx_messageInfo_DestinationSpec.Size(m)
}
func (m *DestinationSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DestinationSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DestinationSpec proto.InternalMessageInfo

type isDestinationSpec_DestinationType interface {
	isDestinationSpec_DestinationType()
	Equal(interface{}) bool
}

type DestinationSpec_Aws struct {
	Aws *aws.DestinationSpec `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}
type DestinationSpec_Azure struct {
	Azure *azure.DestinationSpec `protobuf:"bytes,2,opt,name=azure,proto3,oneof"`
}
type DestinationSpec_Rest struct {
	Rest *rest.DestinationSpec `protobuf:"bytes,3,opt,name=rest,proto3,oneof"`
}
type DestinationSpec_Grpc struct {
	Grpc *grpc.DestinationSpec `protobuf:"bytes,4,opt,name=grpc,proto3,oneof"`
}

func (*DestinationSpec_Aws) isDestinationSpec_DestinationType()   {}
func (*DestinationSpec_Azure) isDestinationSpec_DestinationType() {}
func (*DestinationSpec_Rest) isDestinationSpec_DestinationType()  {}
func (*DestinationSpec_Grpc) isDestinationSpec_DestinationType()  {}

func (m *DestinationSpec) GetDestinationType() isDestinationSpec_DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return nil
}

func (m *DestinationSpec) GetAws() *aws.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *DestinationSpec) GetAzure() *azure.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *DestinationSpec) GetRest() *rest.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Rest); ok {
		return x.Rest
	}
	return nil
}

func (m *DestinationSpec) GetGrpc() *grpc.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DestinationSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DestinationSpec_OneofMarshaler, _DestinationSpec_OneofUnmarshaler, _DestinationSpec_OneofSizer, []interface{}{
		(*DestinationSpec_Aws)(nil),
		(*DestinationSpec_Azure)(nil),
		(*DestinationSpec_Rest)(nil),
		(*DestinationSpec_Grpc)(nil),
	}
}

func _DestinationSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DestinationSpec)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *DestinationSpec_Aws:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *DestinationSpec_Azure:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *DestinationSpec_Rest:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rest); err != nil {
			return err
		}
	case *DestinationSpec_Grpc:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Grpc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DestinationSpec.DestinationType has unexpected type %T", x)
	}
	return nil
}

func _DestinationSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DestinationSpec)
	switch tag {
	case 1: // destination_type.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(aws.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Aws{msg}
		return true, err
	case 2: // destination_type.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(azure.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Azure{msg}
		return true, err
	case 3: // destination_type.rest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(rest.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Rest{msg}
		return true, err
	case 4: // destination_type.grpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(grpc.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Grpc{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DestinationSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DestinationSpec)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *DestinationSpec_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Rest:
		s := proto.Size(x.Rest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Grpc:
		s := proto.Size(x.Grpc)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin.
type UpstreamSpec struct {
	SslConfig *UpstreamSslConfig `protobuf:"bytes,6,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Circuite breakers for this upstream. if not set, the defaults ones from the Gloo settings will be used.
	// if those are not set,  [envoy's defaults](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
	// will be used.
	CircuitBreakers *CircuitBreakerConfig `protobuf:"bytes,7,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	// Note to developers: new Upstream Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to UpstreamType:
	//	*UpstreamSpec_Kube
	//	*UpstreamSpec_Static
	//	*UpstreamSpec_Aws
	//	*UpstreamSpec_Azure
	//	*UpstreamSpec_Consul
	UpstreamType         isUpstreamSpec_UpstreamType `protobuf_oneof:"upstream_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae47d2df5fad2a45, []int{4}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

type isUpstreamSpec_UpstreamType interface {
	isUpstreamSpec_UpstreamType()
	Equal(interface{}) bool
}

type UpstreamSpec_Kube struct {
	Kube *kubernetes.UpstreamSpec `protobuf:"bytes,1,opt,name=kube,proto3,oneof"`
}
type UpstreamSpec_Static struct {
	Static *static.UpstreamSpec `protobuf:"bytes,4,opt,name=static,proto3,oneof"`
}
type UpstreamSpec_Aws struct {
	Aws *aws.UpstreamSpec `protobuf:"bytes,2,opt,name=aws,proto3,oneof"`
}
type UpstreamSpec_Azure struct {
	Azure *azure.UpstreamSpec `protobuf:"bytes,3,opt,name=azure,proto3,oneof"`
}
type UpstreamSpec_Consul struct {
	Consul *consul.UpstreamSpec `protobuf:"bytes,5,opt,name=consul,proto3,oneof"`
}

func (*UpstreamSpec_Kube) isUpstreamSpec_UpstreamType()   {}
func (*UpstreamSpec_Static) isUpstreamSpec_UpstreamType() {}
func (*UpstreamSpec_Aws) isUpstreamSpec_UpstreamType()    {}
func (*UpstreamSpec_Azure) isUpstreamSpec_UpstreamType()  {}
func (*UpstreamSpec_Consul) isUpstreamSpec_UpstreamType() {}

func (m *UpstreamSpec) GetUpstreamType() isUpstreamSpec_UpstreamType {
	if m != nil {
		return m.UpstreamType
	}
	return nil
}

func (m *UpstreamSpec) GetSslConfig() *UpstreamSslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *UpstreamSpec) GetCircuitBreakers() *CircuitBreakerConfig {
	if m != nil {
		return m.CircuitBreakers
	}
	return nil
}

func (m *UpstreamSpec) GetKube() *kubernetes.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *UpstreamSpec) GetStatic() *static.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Static); ok {
		return x.Static
	}
	return nil
}

func (m *UpstreamSpec) GetAws() *aws.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *UpstreamSpec) GetAzure() *azure.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *UpstreamSpec) GetConsul() *consul.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Consul); ok {
		return x.Consul
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpstreamSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpstreamSpec_OneofMarshaler, _UpstreamSpec_OneofUnmarshaler, _UpstreamSpec_OneofSizer, []interface{}{
		(*UpstreamSpec_Kube)(nil),
		(*UpstreamSpec_Static)(nil),
		(*UpstreamSpec_Aws)(nil),
		(*UpstreamSpec_Azure)(nil),
		(*UpstreamSpec_Consul)(nil),
	}
}

func _UpstreamSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpstreamSpec)
	// upstream_type
	switch x := m.UpstreamType.(type) {
	case *UpstreamSpec_Kube:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Kube); err != nil {
			return err
		}
	case *UpstreamSpec_Static:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Static); err != nil {
			return err
		}
	case *UpstreamSpec_Aws:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *UpstreamSpec_Azure:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *UpstreamSpec_Consul:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Consul); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UpstreamSpec.UpstreamType has unexpected type %T", x)
	}
	return nil
}

func _UpstreamSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpstreamSpec)
	switch tag {
	case 1: // upstream_type.kube
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Kube{msg}
		return true, err
	case 4: // upstream_type.static
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(static.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Static{msg}
		return true, err
	case 2: // upstream_type.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(aws.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Aws{msg}
		return true, err
	case 3: // upstream_type.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(azure.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Azure{msg}
		return true, err
	case 5: // upstream_type.consul
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(consul.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Consul{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UpstreamSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpstreamSpec)
	// upstream_type
	switch x := m.UpstreamType.(type) {
	case *UpstreamSpec_Kube:
		s := proto.Size(x.Kube)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Static:
		s := proto.Size(x.Static)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Consul:
		s := proto.Size(x.Consul)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ListenerPlugins)(nil), "gloo.solo.io.ListenerPlugins")
	proto.RegisterType((*VirtualHostPlugins)(nil), "gloo.solo.io.VirtualHostPlugins")
	proto.RegisterType((*RoutePlugins)(nil), "gloo.solo.io.RoutePlugins")
	proto.RegisterType((*DestinationSpec)(nil), "gloo.solo.io.DestinationSpec")
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.solo.io.UpstreamSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins.proto", fileDescriptor_ae47d2df5fad2a45)
}

var fileDescriptor_ae47d2df5fad2a45 = []byte{
	// 867 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x49, 0x93, 0x4d, 0x60, 0xe8, 0x92, 0x95, 0xc5, 0x45, 0xa8, 0x60, 0x77, 0x95, 0x0b,
	0xd8, 0x2e, 0xda, 0x31, 0x2c, 0xd2, 0x42, 0x2b, 0x15, 0x4a, 0x52, 0x4a, 0x84, 0x5a, 0x88, 0x5c,
	0xbe, 0x6f, 0x22, 0xc7, 0x9d, 0x38, 0x43, 0x1d, 0x8f, 0x35, 0x73, 0x4c, 0x5a, 0xae, 0x78, 0x0c,
	0x1e, 0x81, 0x1b, 0x1e, 0x84, 0x27, 0xe0, 0x12, 0x89, 0xd7, 0xe0, 0x06, 0xcd, 0xcc, 0x19, 0x37,
	0x8e, 0xd2, 0x2a, 0x71, 0xb9, 0x88, 0x3d, 0xce, 0x9c, 0xff, 0xcf, 0x73, 0xe6, 0x9c, 0x39, 0xc7,
	0x64, 0x3f, 0xe6, 0x30, 0xcd, 0xc7, 0x34, 0x12, 0x33, 0x5f, 0x89, 0x44, 0x3c, 0xe3, 0xc2, 0x8f,
	0x13, 0x21, 0xfc, 0x4c, 0x8a, 0x9f, 0x58, 0x04, 0xca, 0x3e, 0x85, 0x19, 0xf7, 0x7f, 0x7e, 0xdf,
	0xcf, 0x92, 0x3c, 0xe6, 0xa9, 0xa2, 0x99, 0x14, 0x20, 0xbc, 0x6d, 0x3d, 0x45, 0xb5, 0x8a, 0x72,
	0xb1, 0xf3, 0x66, 0x2c, 0x44, 0x9c, 0x30, 0xdf, 0xcc, 0x8d, 0xf3, 0x89, 0xaf, 0x40, 0xe6, 0x11,
	0x58, 0xdb, 0x9d, 0xd7, 0x63, 0x11, 0x0b, 0x33, 0xf4, 0xf5, 0x08, 0xff, 0x7d, 0xb1, 0xd1, 0xdb,
	0x95, 0x4a, 0x50, 0x77, 0xb0, 0x91, 0x8e, 0x5d, 0x02, 0x4b, 0x15, 0x17, 0x6e, 0xe1, 0x3b, 0x47,
	0x1b, 0xc9, 0x23, 0x2e, 0xa3, 0x9c, 0x03, 0x1b, 0x8d, 0x25, 0x0b, 0x2f, 0x98, 0x74, 0x94, 0x5e,
	0x95, 0xad, 0xf3, 0xc3, 0xb9, 0xf9, 0x55, 0x5a, 0x89, 0x63, 0x48, 0xa6, 0xc0, 0x5c, 0xee, 0x44,
	0x89, 0x65, 0x16, 0x99, 0x0b, 0x52, 0x4e, 0x2a, 0x53, 0x46, 0x73, 0x36, 0x2e, 0x06, 0x77, 0xda,
	0x9d, 0x69, 0x34, 0xd3, 0x3f, 0x64, 0x1c, 0x57, 0xdb, 0xe1, 0x5f, 0x72, 0xc9, 0xec, 0x15, 0x39,
	0x83, 0x4a, 0x9c, 0x48, 0xa4, 0x2a, 0x4f, 0xf0, 0x86, 0xa4, 0x61, 0x25, 0xd2, 0x45, 0x3e, 0x66,
	0x32, 0x65, 0xc0, 0x16, 0x87, 0x48, 0xfc, 0xa2, 0x62, 0x06, 0x80, 0xe4, 0xac, 0xb8, 0xdf, 0xc9,
	0x4f, 0x05, 0x21, 0xf0, 0x08, 0x6f, 0x48, 0xfa, 0xbe, 0x12, 0x09, 0x64, 0x98, 0xaa, 0x89, 0x90,
	0xb3, 0x10, 0xb8, 0x48, 0xfd, 0x4c, 0xb2, 0x09, 0xbf, 0x1c, 0x49, 0x36, 0x97, 0x1c, 0xd8, 0xff,
	0x49, 0x2e, 0x3f, 0x22, 0xf9, 0xab, 0x4a, 0xe4, 0x49, 0x98, 0x27, 0xc0, 0x53, 0x6d, 0xa0, 0xc9,
	0xe6, 0x11, 0x81, 0x0f, 0x97, 0x2b, 0xda, 0x79, 0x2e, 0x17, 0x5e, 0xd8, 0xfd, 0xab, 0x46, 0xda,
	0x27, 0x5c, 0x01, 0x4b, 0x99, 0x1c, 0x5a, 0x9c, 0xf7, 0x29, 0x79, 0xd9, 0x1d, 0x84, 0x4e, 0xed,
	0x71, 0xed, 0xc9, 0xab, 0xcf, 0xdf, 0xa6, 0xd7, 0x27, 0x03, 0xcb, 0xe7, 0x62, 0xdd, 0xa4, 0x9f,
	0xcb, 0x2c, 0xfa, 0x8e, 0x8d, 0x83, 0x56, 0x6c, 0x07, 0xde, 0xaf, 0x35, 0xf2, 0x78, 0x0a, 0x90,
	0x8d, 0x22, 0x91, 0xa6, 0x76, 0x59, 0xa3, 0x59, 0x98, 0x86, 0x31, 0x93, 0x23, 0xc5, 0x00, 0x78,
	0x1a, 0xab, 0xce, 0x96, 0x61, 0x7f, 0x48, 0xcd, 0x61, 0x59, 0x85, 0x1d, 0x00, 0x64, 0xfd, 0x02,
	0x70, 0x6a, 0xf5, 0x67, 0x28, 0x0f, 0xde, 0x9a, 0xde, 0x36, 0xdd, 0xfd, 0x92, 0x78, 0xdf, 0x72,
	0x09, 0x79, 0x98, 0x0c, 0x84, 0x02, 0xe7, 0xdb, 0x47, 0x84, 0x5c, 0x97, 0x52, 0xf4, 0xae, 0x53,
	0x7e, 0xeb, 0x67, 0xc5, 0x7c, 0xb0, 0x60, 0xdb, 0xfd, 0xb3, 0x4e, 0xb6, 0x03, 0x91, 0x03, 0x73,
	0xa8, 0x88, 0xb4, 0xcb, 0x31, 0x74, 0xbc, 0x3d, 0xba, 0x1c, 0xdb, 0x55, 0xce, 0x19, 0xd6, 0xd7,
	0x65, 0x40, 0xb0, 0x4c, 0xf4, 0x3e, 0x21, 0x4d, 0x13, 0x4e, 0xb7, 0x5b, 0xef, 0x50, 0x8c, 0xee,
	0x8d, 0xc8, 0x63, 0x63, 0x1e, 0xa0, 0xcc, 0xfb, 0x81, 0xbc, 0x56, 0xce, 0xe1, 0x4e, 0xdd, 0x80,
	0x9e, 0xaf, 0xb5, 0xc8, 0xa1, 0x91, 0x06, 0x56, 0x19, 0xdc, 0xcf, 0x16, 0x1f, 0xbd, 0x3d, 0xd2,
	0x02, 0x3e, 0x63, 0x22, 0x87, 0x4e, 0xc3, 0x30, 0xdf, 0xa0, 0x36, 0xdb, 0xa8, 0xcb, 0x36, 0x7a,
	0x84, 0xd9, 0xd6, 0x6b, 0xfc, 0xf6, 0xf7, 0xa3, 0x5a, 0xe0, 0xec, 0xbd, 0x3e, 0x69, 0xe1, 0xb1,
	0xef, 0xdc, 0x33, 0xd2, 0x5d, 0x5a, 0x94, 0x81, 0x95, 0x9e, 0x31, 0x90, 0x57, 0x43, 0x91, 0xf0,
	0xe8, 0x2a, 0x70, 0xca, 0xa5, 0x58, 0x36, 0x37, 0x88, 0xe5, 0x1f, 0x5b, 0xa4, 0x7d, 0xc4, 0x14,
	0xf0, 0xd4, 0xac, 0xee, 0x2c, 0x63, 0x91, 0x77, 0x40, 0xea, 0xe1, 0xdc, 0x85, 0x70, 0x97, 0x9a,
	0xfe, 0xb6, 0x6a, 0x29, 0x4b, 0xba, 0xc1, 0x4b, 0x81, 0xd6, 0x79, 0x7d, 0x72, 0xcf, 0x94, 0x6b,
	0x8c, 0xd3, 0xbb, 0x14, 0x8b, 0xf7, 0x7a, 0x08, 0xab, 0xf5, 0x0e, 0x49, 0x43, 0xb7, 0x44, 0x0c,
	0xd1, 0x53, 0x6a, 0xfb, 0xe3, 0x7a, 0x08, 0xa3, 0xd4, 0x04, 0x7d, 0x06, 0x31, 0x20, 0x4f, 0xa9,
	0xed, 0x8d, 0x6b, 0x12, 0xb4, 0x71, 0xcf, 0x23, 0x0f, 0xce, 0xaf, 0xa7, 0x46, 0x70, 0x95, 0xb1,
	0xee, 0xbf, 0x75, 0xb2, 0xfd, 0x4d, 0xa6, 0x40, 0xb2, 0x70, 0x66, 0x36, 0xeb, 0x63, 0x42, 0x94,
	0x4a, 0xf4, 0xe9, 0x9e, 0xf0, 0x18, 0xb7, 0xfe, 0x51, 0x99, 0x5f, 0xd8, 0xab, 0xa4, 0x6f, 0xcc,
	0x82, 0x57, 0x94, 0x1b, 0x7a, 0xa7, 0xe4, 0x01, 0x7e, 0x92, 0x14, 0x5f, 0x24, 0x9d, 0x96, 0xa1,
	0x74, 0xcb, 0x94, 0xbe, 0xb5, 0xea, 0x59, 0x23, 0x04, 0xb5, 0xa3, 0xd2, 0xbf, 0x7a, 0xf3, 0x1b,
	0xba, 0x29, 0x61, 0xf0, 0x9e, 0xd1, 0xc5, 0x0e, 0xb5, 0xca, 0xf7, 0x45, 0x5f, 0xb4, 0xe3, 0xda,
	0xde, 0xeb, 0x93, 0xa6, 0xed, 0x1f, 0xb8, 0x79, 0xbb, 0xd4, 0xb5, 0x93, 0x35, 0x10, 0x28, 0xf5,
	0xf6, 0x6d, 0x16, 0x6d, 0x61, 0xd9, 0xbc, 0x31, 0x8b, 0x96, 0xe4, 0x26, 0x85, 0x0e, 0x5d, 0x0a,
	0xd9, 0xf0, 0x3f, 0xb9, 0x2d, 0x85, 0x96, 0xf4, 0x98, 0x3f, 0x7d, 0xd2, 0xb4, 0xad, 0xbe, 0x38,
	0x55, 0xae, 0xf3, 0xaf, 0xe3, 0x82, 0xb5, 0xed, 0xb5, 0xc9, 0xfd, 0x1c, 0x67, 0x4c, 0xf4, 0x7b,
	0x2f, 0x7e, 0xff, 0xe7, 0x61, 0xed, 0xc7, 0xf7, 0xd6, 0x6b, 0x4d, 0xd9, 0x45, 0x8c, 0xed, 0x69,
	0xdc, 0x34, 0x65, 0xe0, 0x83, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x1a, 0x14, 0x41, 0xa0,
	0x0b, 0x00, 0x00,
}

func (this *ListenerPlugins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerPlugins)
	if !ok {
		that2, ok := that.(ListenerPlugins)
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
	if !this.GrpcWeb.Equal(that1.GrpcWeb) {
		return false
	}
	if !this.HttpConnectionManagerSettings.Equal(that1.HttpConnectionManagerSettings) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualHostPlugins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHostPlugins)
	if !ok {
		that2, ok := that.(VirtualHostPlugins)
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
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RoutePlugins) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoutePlugins)
	if !ok {
		that2, ok := that.(RoutePlugins)
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
	if !this.Transformations.Equal(that1.Transformations) {
		return false
	}
	if !this.Faults.Equal(that1.Faults) {
		return false
	}
	if !this.PrefixRewrite.Equal(that1.PrefixRewrite) {
		return false
	}
	if this.Timeout != nil && that1.Timeout != nil {
		if *this.Timeout != *that1.Timeout {
			return false
		}
	} else if this.Timeout != nil {
		return false
	} else if that1.Timeout != nil {
		return false
	}
	if !this.Retries.Equal(that1.Retries) {
		return false
	}
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if that1.DestinationType == nil {
		if this.DestinationType != nil {
			return false
		}
	} else if this.DestinationType == nil {
		return false
	} else if !this.DestinationType.Equal(that1.DestinationType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DestinationSpec_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Aws)
	if !ok {
		that2, ok := that.(DestinationSpec_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *DestinationSpec_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Azure)
	if !ok {
		that2, ok := that.(DestinationSpec_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *DestinationSpec_Rest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Rest)
	if !ok {
		that2, ok := that.(DestinationSpec_Rest)
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
	if !this.Rest.Equal(that1.Rest) {
		return false
	}
	return true
}
func (this *DestinationSpec_Grpc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Grpc)
	if !ok {
		that2, ok := that.(DestinationSpec_Grpc)
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
	if !this.Grpc.Equal(that1.Grpc) {
		return false
	}
	return true
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.CircuitBreakers.Equal(that1.CircuitBreakers) {
		return false
	}
	if that1.UpstreamType == nil {
		if this.UpstreamType != nil {
			return false
		}
	} else if this.UpstreamType == nil {
		return false
	} else if !this.UpstreamType.Equal(that1.UpstreamType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Kube) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Kube)
	if !ok {
		that2, ok := that.(UpstreamSpec_Kube)
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
	if !this.Kube.Equal(that1.Kube) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Static) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Static)
	if !ok {
		that2, ok := that.(UpstreamSpec_Static)
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
	if !this.Static.Equal(that1.Static) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Aws)
	if !ok {
		that2, ok := that.(UpstreamSpec_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Azure)
	if !ok {
		that2, ok := that.(UpstreamSpec_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Consul)
	if !ok {
		that2, ok := that.(UpstreamSpec_Consul)
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
	if !this.Consul.Equal(that1.Consul) {
		return false
	}
	return true
}
