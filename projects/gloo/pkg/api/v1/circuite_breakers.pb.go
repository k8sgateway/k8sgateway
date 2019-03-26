// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/circuite_breakers.proto

package v1

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// SslConfig contains the options necessary to configure a virtual host or listener to use TLS
// See the [envoy docs](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
// for the meaning of these values.
type CircuitBreakerConfig struct {
	MaxConnections       *types.UInt32Value `protobuf:"bytes,1,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *types.UInt32Value `protobuf:"bytes,2,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *types.UInt32Value `protobuf:"bytes,3,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *types.UInt32Value `protobuf:"bytes,4,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CircuitBreakerConfig) Reset()         { *m = CircuitBreakerConfig{} }
func (m *CircuitBreakerConfig) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakerConfig) ProtoMessage()    {}
func (*CircuitBreakerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce2af6125bf26c85, []int{0}
}
func (m *CircuitBreakerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakerConfig.Unmarshal(m, b)
}
func (m *CircuitBreakerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakerConfig.Marshal(b, m, deterministic)
}
func (m *CircuitBreakerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakerConfig.Merge(m, src)
}
func (m *CircuitBreakerConfig) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakerConfig.Size(m)
}
func (m *CircuitBreakerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakerConfig proto.InternalMessageInfo

func (m *CircuitBreakerConfig) GetMaxConnections() *types.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakerConfig) GetMaxPendingRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakerConfig) GetMaxRequests() *types.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakerConfig) GetMaxRetries() *types.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakerConfig)(nil), "gloo.solo.io.CircuitBreakerConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/circuite_breakers.proto", fileDescriptor_ce2af6125bf26c85)
}

var fileDescriptor_ce2af6125bf26c85 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc0, 0x71, 0xa5, 0xdf, 0x27, 0x06, 0xa7, 0x02, 0x29, 0xca, 0x10, 0x55, 0xa8, 0x42, 0x4c,
	0x2c, 0xd8, 0xd0, 0x4a, 0x6c, 0x08, 0xa9, 0x81, 0x81, 0x05, 0xa1, 0x4a, 0x30, 0xb0, 0x44, 0x8e,
	0xb9, 0x1a, 0xd3, 0xc4, 0x67, 0x6c, 0x07, 0xf2, 0x48, 0x2c, 0xbc, 0x14, 0x4f, 0x82, 0x12, 0x97,
	0xb0, 0x66, 0xf3, 0xc9, 0xf7, 0xff, 0x2d, 0x47, 0xae, 0xa5, 0xf2, 0x2f, 0x4d, 0x49, 0x05, 0xd6,
	0xcc, 0x61, 0x85, 0xa7, 0x0a, 0x99, 0xac, 0x10, 0x99, 0xb1, 0xf8, 0x0a, 0xc2, 0xbb, 0x30, 0x71,
	0xa3, 0xd8, 0xfb, 0x39, 0x13, 0xca, 0x8a, 0x46, 0x79, 0x28, 0x4a, 0x0b, 0x7c, 0x0b, 0xd6, 0x51,
	0x63, 0xd1, 0x63, 0x32, 0xed, 0x96, 0x68, 0xd7, 0x53, 0x85, 0xb3, 0x54, 0xa2, 0xc4, 0xfe, 0x83,
	0x75, 0xaf, 0xb0, 0x33, 0x9b, 0x4b, 0x44, 0x59, 0x01, 0xeb, 0xa7, 0xb2, 0xd9, 0xb0, 0x0f, 0xcb,
	0x8d, 0x19, 0x8c, 0xe3, 0xaf, 0x09, 0x49, 0xf3, 0xe0, 0xaf, 0x82, 0x9e, 0xa3, 0xde, 0x28, 0x99,
	0xdc, 0x90, 0x83, 0x9a, 0xb7, 0x85, 0x40, 0xad, 0x41, 0x78, 0x85, 0xda, 0x65, 0xd1, 0x51, 0x74,
	0x12, 0x2f, 0x0e, 0x69, 0x20, 0xe9, 0x2f, 0x49, 0x1f, 0x6e, 0xb5, 0x5f, 0x2e, 0x1e, 0x79, 0xd5,
	0xc0, 0x7a, 0xbf, 0xe6, 0x6d, 0xfe, 0xd7, 0x24, 0x77, 0x24, 0xed, 0x18, 0x03, 0xfa, 0x59, 0x69,
	0x59, 0x58, 0x78, 0x6b, 0xc0, 0x79, 0x97, 0x4d, 0x46, 0x58, 0x49, 0xcd, 0xdb, 0xfb, 0x10, 0xae,
	0x77, 0x5d, 0x72, 0x45, 0xa6, 0x9d, 0x37, 0x38, 0xff, 0x46, 0x38, 0x71, 0xcd, 0xdb, 0x01, 0xb8,
	0x24, 0x71, 0x00, 0xbc, 0x55, 0xe0, 0xb2, 0xff, 0x23, 0x7a, 0xd2, 0xf7, 0xfd, 0xfe, 0xea, 0xe2,
	0xf3, 0x7b, 0x1e, 0x3d, 0x9d, 0x8d, 0xbb, 0x9f, 0xd9, 0xca, 0xdd, 0x0d, 0xcb, 0xbd, 0x5e, 0x5e,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x85, 0x12, 0xe5, 0xe1, 0xfa, 0x01, 0x00, 0x00,
}

func (this *CircuitBreakerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CircuitBreakerConfig)
	if !ok {
		that2, ok := that.(CircuitBreakerConfig)
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
	if !this.MaxConnections.Equal(that1.MaxConnections) {
		return false
	}
	if !this.MaxPendingRequests.Equal(that1.MaxPendingRequests) {
		return false
	}
	if !this.MaxRequests.Equal(that1.MaxRequests) {
		return false
	}
	if !this.MaxRetries.Equal(that1.MaxRetries) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
