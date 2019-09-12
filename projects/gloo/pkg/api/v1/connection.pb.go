// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/connection.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
	time "time"
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

// Fine tune the settings for connections to an upstream
type ConnectionConfig struct {
	// Maximum requests for a single upstream connection (unspecified or zero = no limit)
	MaxRequestsPerConnection uint32 `protobuf:"varint,1,opt,name=max_requests_per_connection,json=maxRequestsPerConnection,proto3" json:"max_requests_per_connection,omitempty"`
	// The timeout for new network connections to hosts in the cluster
	ConnectTimeout *time.Duration `protobuf:"bytes,2,opt,name=connect_timeout,json=connectTimeout,proto3,stdduration" json:"connect_timeout,omitempty"`
	// Configure OS-level tcp keepalive checks
	TcpKeepalive         *ConnectionConfig_TcpKeepAlive `protobuf:"bytes,3,opt,name=tcp_keepalive,json=tcpKeepalive,proto3" json:"tcp_keepalive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ConnectionConfig) Reset()         { *m = ConnectionConfig{} }
func (m *ConnectionConfig) String() string { return proto.CompactTextString(m) }
func (*ConnectionConfig) ProtoMessage()    {}
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_56610fe13cf10c84, []int{0}
}
func (m *ConnectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionConfig.Unmarshal(m, b)
}
func (m *ConnectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionConfig.Marshal(b, m, deterministic)
}
func (m *ConnectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionConfig.Merge(m, src)
}
func (m *ConnectionConfig) XXX_Size() int {
	return xxx_messageInfo_ConnectionConfig.Size(m)
}
func (m *ConnectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionConfig proto.InternalMessageInfo

func (m *ConnectionConfig) GetMaxRequestsPerConnection() uint32 {
	if m != nil {
		return m.MaxRequestsPerConnection
	}
	return 0
}

func (m *ConnectionConfig) GetConnectTimeout() *time.Duration {
	if m != nil {
		return m.ConnectTimeout
	}
	return nil
}

func (m *ConnectionConfig) GetTcpKeepalive() *ConnectionConfig_TcpKeepAlive {
	if m != nil {
		return m.TcpKeepalive
	}
	return nil
}

// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
// see more info here: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/address.proto#envoy-api-msg-core-tcpkeepalive
type ConnectionConfig_TcpKeepAlive struct {
	// Maximum number of keepalive probes to send without response before deciding the connection is dead.
	KeepaliveProbes uint32 `protobuf:"varint,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	// The number of seconds a connection needs to be idle before keep-alive probes start being sent. This is rounded up to the second.
	KeepaliveTime *time.Duration `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3,stdduration" json:"keepalive_time,omitempty"`
	// The number of seconds between keep-alive probes. This is rounded up to the second.
	KeepaliveInterval    *time.Duration `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3,stdduration" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConnectionConfig_TcpKeepAlive) Reset()         { *m = ConnectionConfig_TcpKeepAlive{} }
func (m *ConnectionConfig_TcpKeepAlive) String() string { return proto.CompactTextString(m) }
func (*ConnectionConfig_TcpKeepAlive) ProtoMessage()    {}
func (*ConnectionConfig_TcpKeepAlive) Descriptor() ([]byte, []int) {
	return fileDescriptor_56610fe13cf10c84, []int{0, 0}
}
func (m *ConnectionConfig_TcpKeepAlive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionConfig_TcpKeepAlive.Unmarshal(m, b)
}
func (m *ConnectionConfig_TcpKeepAlive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionConfig_TcpKeepAlive.Marshal(b, m, deterministic)
}
func (m *ConnectionConfig_TcpKeepAlive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionConfig_TcpKeepAlive.Merge(m, src)
}
func (m *ConnectionConfig_TcpKeepAlive) XXX_Size() int {
	return xxx_messageInfo_ConnectionConfig_TcpKeepAlive.Size(m)
}
func (m *ConnectionConfig_TcpKeepAlive) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionConfig_TcpKeepAlive.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionConfig_TcpKeepAlive proto.InternalMessageInfo

func (m *ConnectionConfig_TcpKeepAlive) GetKeepaliveProbes() uint32 {
	if m != nil {
		return m.KeepaliveProbes
	}
	return 0
}

func (m *ConnectionConfig_TcpKeepAlive) GetKeepaliveTime() *time.Duration {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *ConnectionConfig_TcpKeepAlive) GetKeepaliveInterval() *time.Duration {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionConfig)(nil), "gloo.solo.io.ConnectionConfig")
	proto.RegisterType((*ConnectionConfig_TcpKeepAlive)(nil), "gloo.solo.io.ConnectionConfig.TcpKeepAlive")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/connection.proto", fileDescriptor_56610fe13cf10c84)
}

var fileDescriptor_56610fe13cf10c84 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4e, 0xdb, 0x40,
	0x14, 0x96, 0x9b, 0xaa, 0x8b, 0x69, 0xfe, 0x3a, 0xea, 0xc2, 0x4d, 0xa5, 0x34, 0xea, 0x2a, 0x55,
	0xd5, 0x99, 0xb6, 0x48, 0xec, 0xb2, 0x20, 0x41, 0x08, 0x84, 0x84, 0x22, 0x2b, 0x2b, 0x36, 0x96,
	0x6d, 0x5e, 0x86, 0x21, 0xb6, 0xdf, 0x30, 0x1e, 0x87, 0x1c, 0x85, 0x23, 0x70, 0x0d, 0x8e, 0xc0,
	0x09, 0x90, 0x38, 0x09, 0xb2, 0x3d, 0xd8, 0x11, 0x2c, 0xc8, 0xce, 0x6f, 0xbe, 0x9f, 0xf7, 0x7d,
	0xd6, 0x23, 0x13, 0x21, 0xcd, 0x65, 0x1e, 0xb2, 0x08, 0x13, 0x9e, 0x61, 0x8c, 0x7f, 0x24, 0x72,
	0x11, 0x23, 0x72, 0xa5, 0xf1, 0x0a, 0x22, 0x93, 0x55, 0x53, 0xa0, 0x24, 0x5f, 0xff, 0xe3, 0x11,
	0xa6, 0x29, 0x44, 0x46, 0x62, 0xca, 0x94, 0x46, 0x83, 0xb4, 0x5d, 0xa0, 0xac, 0x10, 0x32, 0x89,
	0x83, 0xaf, 0x02, 0x05, 0x96, 0x00, 0x2f, 0xbe, 0x2a, 0xce, 0x60, 0x28, 0x10, 0x45, 0x0c, 0xbc,
	0x9c, 0xc2, 0x7c, 0xc9, 0x2f, 0x72, 0x1d, 0x34, 0x1e, 0x6f, 0xf1, 0x1b, 0x1d, 0x28, 0x05, 0x3a,
	0xab, 0xf0, 0x9f, 0xf7, 0x2d, 0xd2, 0x9f, 0xd5, 0x8b, 0x67, 0x98, 0x2e, 0xa5, 0xa0, 0x13, 0xf2,
	0x3d, 0x09, 0x36, 0xbe, 0x86, 0xeb, 0x1c, 0x32, 0x93, 0xf9, 0x0a, 0xb4, 0xdf, 0xa4, 0x73, 0x9d,
	0x91, 0x33, 0xee, 0x78, 0x6e, 0x12, 0x6c, 0x3c, 0xcb, 0x98, 0x83, 0x6e, 0x4c, 0xe8, 0x31, 0xe9,
	0x59, 0xb6, 0x6f, 0x64, 0x02, 0x98, 0x1b, 0xf7, 0xc3, 0xc8, 0x19, 0x7f, 0xfe, 0xff, 0x8d, 0x55,
	0x69, 0xd8, 0x4b, 0x1a, 0x76, 0x68, 0xd3, 0x4e, 0x3f, 0xde, 0x3e, 0xfe, 0x70, 0xbc, 0xae, 0xd5,
	0x2d, 0x2a, 0x19, 0x9d, 0x93, 0x8e, 0x89, 0x94, 0xbf, 0x02, 0x50, 0x41, 0x2c, 0xd7, 0xe0, 0xb6,
	0x4a, 0x9f, 0xdf, 0x6c, 0xfb, 0xcf, 0xb0, 0xd7, 0xf9, 0xd9, 0x22, 0x52, 0xa7, 0x00, 0xea, 0xa0,
	0x90, 0x78, 0x6d, 0x53, 0x4d, 0xa5, 0xc1, 0xe0, 0xc1, 0x21, 0xed, 0x6d, 0x98, 0xfe, 0x22, 0xfd,
	0xda, 0xde, 0x57, 0x1a, 0x43, 0xc8, 0x6c, 0xc1, 0x5e, 0xfd, 0x3e, 0x2f, 0x9f, 0xe9, 0x11, 0xe9,
	0x36, 0xd4, 0xa2, 0xd9, 0xae, 0xb5, 0x3a, 0xb5, 0xac, 0x28, 0x46, 0xcf, 0x08, 0x6d, 0x7c, 0x64,
	0x6a, 0x40, 0xaf, 0x83, 0xd8, 0x56, 0x7b, 0xd7, 0xeb, 0x4b, 0x2d, 0x3d, 0xb1, 0xca, 0xe9, 0xfe,
	0xdd, 0xd3, 0xd0, 0x39, 0xff, 0xbb, 0xdb, 0xb1, 0xa9, 0x95, 0xb0, 0x07, 0x17, 0x7e, 0x2a, 0x77,
	0xec, 0x3d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x84, 0x69, 0x71, 0xdb, 0xa7, 0x02, 0x00, 0x00,
}

func (this *ConnectionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConnectionConfig)
	if !ok {
		that2, ok := that.(ConnectionConfig)
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
	if this.MaxRequestsPerConnection != that1.MaxRequestsPerConnection {
		return false
	}
	if this.ConnectTimeout != nil && that1.ConnectTimeout != nil {
		if *this.ConnectTimeout != *that1.ConnectTimeout {
			return false
		}
	} else if this.ConnectTimeout != nil {
		return false
	} else if that1.ConnectTimeout != nil {
		return false
	}
	if !this.TcpKeepalive.Equal(that1.TcpKeepalive) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ConnectionConfig_TcpKeepAlive) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConnectionConfig_TcpKeepAlive)
	if !ok {
		that2, ok := that.(ConnectionConfig_TcpKeepAlive)
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
	if this.KeepaliveProbes != that1.KeepaliveProbes {
		return false
	}
	if this.KeepaliveTime != nil && that1.KeepaliveTime != nil {
		if *this.KeepaliveTime != *that1.KeepaliveTime {
			return false
		}
	} else if this.KeepaliveTime != nil {
		return false
	} else if that1.KeepaliveTime != nil {
		return false
	}
	if this.KeepaliveInterval != nil && that1.KeepaliveInterval != nil {
		if *this.KeepaliveInterval != *that1.KeepaliveInterval {
			return false
		}
	} else if this.KeepaliveInterval != nil {
		return false
	} else if that1.KeepaliveInterval != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
