// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tcp/tcp.proto

package tcp

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Contains various settings for Envoy's tcp proxy filter.
// See here for more information: https://www.envoyproxy.io/docs/envoy/v1.10.0/api-v2/config/filter/network/tcp_proxy/v2/tcp_proxy.proto#envoy-api-msg-config-filter-network-tcp-proxy-v2-tcpproxy
type TcpProxySettings struct {
	MaxConnectAttempts   *types.UInt32Value `protobuf:"bytes,1,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	IdleTimeout          *time.Duration     `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3,stdduration" json:"idle_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TcpProxySettings) Reset()         { *m = TcpProxySettings{} }
func (m *TcpProxySettings) String() string { return proto.CompactTextString(m) }
func (*TcpProxySettings) ProtoMessage()    {}
func (*TcpProxySettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eab2eea37fe83e7, []int{0}
}
func (m *TcpProxySettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxySettings.Unmarshal(m, b)
}
func (m *TcpProxySettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxySettings.Marshal(b, m, deterministic)
}
func (m *TcpProxySettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxySettings.Merge(m, src)
}
func (m *TcpProxySettings) XXX_Size() int {
	return xxx_messageInfo_TcpProxySettings.Size(m)
}
func (m *TcpProxySettings) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxySettings.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxySettings proto.InternalMessageInfo

func (m *TcpProxySettings) GetMaxConnectAttempts() *types.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

func (m *TcpProxySettings) GetIdleTimeout() *time.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProxySettings)(nil), "tcp.options.gloo.solo.io.TcpProxySettings")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/tcp/tcp.proto", fileDescriptor_7eab2eea37fe83e7)
}

var fileDescriptor_7eab2eea37fe83e7 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x59, 0x11, 0x8b, 0x3d, 0x0b, 0x59, 0xae, 0x58, 0x0f, 0x39, 0xc5, 0xca, 0xc6, 0x04,
	0xef, 0x5a, 0x1b, 0x57, 0x1b, 0x1b, 0x91, 0xf3, 0xb4, 0xb0, 0x59, 0xb2, 0xb9, 0x18, 0xa3, 0xbb,
	0x3b, 0x43, 0x32, 0xd1, 0xf5, 0x4d, 0x7c, 0x02, 0xf1, 0xad, 0x04, 0x9f, 0x44, 0xb2, 0x59, 0x2b,
	0x15, 0x2c, 0x02, 0x33, 0xfc, 0xff, 0xff, 0x65, 0xf8, 0xd3, 0x42, 0x1b, 0xba, 0xf7, 0x15, 0x93,
	0xd0, 0x70, 0x07, 0x35, 0x1c, 0x1a, 0xe0, 0xba, 0x06, 0xe0, 0x68, 0xe1, 0x41, 0x49, 0x72, 0x71,
	0x13, 0x68, 0xf8, 0xd3, 0x11, 0x07, 0x24, 0x03, 0xad, 0xe3, 0x24, 0x31, 0x3c, 0x86, 0x16, 0x08,
	0xb2, 0x3c, 0x8c, 0x83, 0xc4, 0x82, 0x9d, 0x05, 0x12, 0x33, 0x30, 0x19, 0x6b, 0xd0, 0xd0, 0x9b,
	0x78, 0x98, 0xa2, 0x7f, 0x32, 0xd5, 0x00, 0xba, 0x56, 0xbc, 0xdf, 0x2a, 0x7f, 0xc7, 0x9f, 0xad,
	0x40, 0x54, 0xd6, 0xfd, 0xa5, 0xaf, 0xbc, 0x15, 0x81, 0x1e, 0xf5, 0xfd, 0xb7, 0x24, 0xdd, 0x5a,
	0x4a, 0xbc, 0xb4, 0xd0, 0xbd, 0x5c, 0x29, 0x22, 0xd3, 0x6a, 0x97, 0x5d, 0xa4, 0xe3, 0x46, 0x74,
	0xa5, 0x84, 0xb6, 0x55, 0x92, 0x4a, 0x41, 0xa4, 0x1a, 0x24, 0x97, 0x27, 0x7b, 0xc9, 0xc1, 0x68,
	0xb6, 0xc3, 0x22, 0x93, 0x7d, 0x33, 0xd9, 0xf5, 0x79, 0x4b, 0xf3, 0xd9, 0x8d, 0xa8, 0xbd, 0x5a,
	0x64, 0x8d, 0xe8, 0x4e, 0x63, 0xf0, 0x64, 0xc8, 0x65, 0x45, 0xba, 0x69, 0x56, 0xb5, 0x2a, 0xc9,
	0x34, 0x0a, 0x3c, 0xe5, 0x6b, 0x3d, 0x67, 0xfb, 0x07, 0xe7, 0x6c, 0xb8, 0xad, 0x58, 0x7f, 0xfd,
	0xd8, 0x4d, 0x16, 0xa3, 0x10, 0x5a, 0xc6, 0x4c, 0x51, 0xbc, 0x7f, 0x4e, 0x93, 0xdb, 0xe3, 0xff,
	0x55, 0x8c, 0x8f, 0xfa, 0x97, 0x9a, 0xab, 0x8d, 0xfe, 0xa7, 0xf9, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0xc1, 0x2d, 0xb8, 0xa9, 0x01, 0x00, 0x00,
}

func (this *TcpProxySettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TcpProxySettings)
	if !ok {
		that2, ok := that.(TcpProxySettings)
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
	if !this.MaxConnectAttempts.Equal(that1.MaxConnectAttempts) {
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
