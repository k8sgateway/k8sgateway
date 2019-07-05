// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/service_spec.proto

package plugins

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/grpc"
	rest "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/rest"
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

// Describes APIs and application-level information for services
// Gloo routes to. ServiceSpec is contained within the UpstreamSpec for certain types
// of upstreams, including Kubernetes, Consul, and Static.
// ServiceSpec configuration is opaque to Gloo and handled by Service Plugins.
type ServiceSpec struct {
	// Note to developers: new Service Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to PluginType:
	//	*ServiceSpec_Rest
	//	*ServiceSpec_Grpc
	PluginType           isServiceSpec_PluginType `protobuf_oneof:"plugin_type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_354c0c4fb380b5cd, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(m, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

type isServiceSpec_PluginType interface {
	isServiceSpec_PluginType()
	Equal(interface{}) bool
}

type ServiceSpec_Rest struct {
	Rest *rest.ServiceSpec `protobuf:"bytes,1,opt,name=rest,proto3,oneof"`
}
type ServiceSpec_Grpc struct {
	Grpc *grpc.ServiceSpec `protobuf:"bytes,2,opt,name=grpc,proto3,oneof"`
}

func (*ServiceSpec_Rest) isServiceSpec_PluginType() {}
func (*ServiceSpec_Grpc) isServiceSpec_PluginType() {}

func (m *ServiceSpec) GetPluginType() isServiceSpec_PluginType {
	if m != nil {
		return m.PluginType
	}
	return nil
}

func (m *ServiceSpec) GetRest() *rest.ServiceSpec {
	if x, ok := m.GetPluginType().(*ServiceSpec_Rest); ok {
		return x.Rest
	}
	return nil
}

func (m *ServiceSpec) GetGrpc() *grpc.ServiceSpec {
	if x, ok := m.GetPluginType().(*ServiceSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServiceSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServiceSpec_OneofMarshaler, _ServiceSpec_OneofUnmarshaler, _ServiceSpec_OneofSizer, []interface{}{
		(*ServiceSpec_Rest)(nil),
		(*ServiceSpec_Grpc)(nil),
	}
}

func _ServiceSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServiceSpec)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *ServiceSpec_Rest:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rest); err != nil {
			return err
		}
	case *ServiceSpec_Grpc:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Grpc); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServiceSpec.PluginType has unexpected type %T", x)
	}
	return nil
}

func _ServiceSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServiceSpec)
	switch tag {
	case 1: // plugin_type.rest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(rest.ServiceSpec)
		err := b.DecodeMessage(msg)
		m.PluginType = &ServiceSpec_Rest{msg}
		return true, err
	case 2: // plugin_type.grpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(grpc.ServiceSpec)
		err := b.DecodeMessage(msg)
		m.PluginType = &ServiceSpec_Grpc{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServiceSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServiceSpec)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *ServiceSpec_Rest:
		s := proto.Size(x.Rest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServiceSpec_Grpc:
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

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "plugins.gloo.solo.io.ServiceSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/service_spec.proto", fileDescriptor_354c0c4fb380b5cd)
}

var fileDescriptor_354c0c4fb380b5cd = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0x0b, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x8a, 0xf5, 0x8b, 0x53, 0x8b,
	0xca, 0x32, 0x93, 0x53, 0xe3, 0x8b, 0x0b, 0x52, 0x93, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x44, 0xa0, 0x72, 0x7a, 0x20, 0xf5, 0x7a, 0x20, 0xa3, 0xf4, 0x32, 0xf3, 0xa5, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x0a, 0xf4, 0x41, 0x2c, 0x88, 0x5a, 0x29, 0x17, 0xb2, 0x2c, 0x2d, 0x4a, 0x2d,
	0x2e, 0x01, 0x13, 0x14, 0x99, 0x92, 0x5e, 0x54, 0x90, 0x0c, 0x26, 0x20, 0xa6, 0x28, 0xcd, 0x60,
	0xe4, 0xe2, 0x0e, 0x86, 0x78, 0x27, 0xb8, 0x20, 0x35, 0x59, 0xc8, 0x86, 0x8b, 0x05, 0x64, 0x87,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x9a, 0x1e, 0xc4, 0x42, 0x2c, 0x7e, 0xd3, 0x43, 0xd2,
	0xe5, 0xc1, 0x10, 0x04, 0xd6, 0x05, 0xd2, 0x0d, 0x32, 0x5b, 0x82, 0x09, 0xaa, 0x1b, 0x62, 0x11,
	0x11, 0xba, 0x41, 0x0a, 0x9d, 0x78, 0xb9, 0xb8, 0x21, 0x6a, 0xe3, 0x4b, 0x2a, 0x0b, 0x52, 0x9d,
	0xec, 0x56, 0x3c, 0x92, 0x63, 0x8c, 0xb2, 0x20, 0xce, 0x9b, 0x05, 0xd9, 0xe9, 0x68, 0x5e, 0x4d,
	0x62, 0x03, 0xfb, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xa0, 0xbd, 0x4a, 0xe4, 0x01,
	0x00, 0x00,
}

func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if that1.PluginType == nil {
		if this.PluginType != nil {
			return false
		}
	} else if this.PluginType == nil {
		return false
	} else if !this.PluginType.Equal(that1.PluginType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_Rest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_Rest)
	if !ok {
		that2, ok := that.(ServiceSpec_Rest)
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
func (this *ServiceSpec_Grpc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_Grpc)
	if !ok {
		that2, ok := that.(ServiceSpec_Grpc)
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
