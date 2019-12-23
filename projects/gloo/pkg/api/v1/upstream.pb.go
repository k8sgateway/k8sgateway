// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	cluster "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/cluster"
	core1 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"
	aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"
	ec2 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws/ec2"
	azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"
	consul "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/consul"
	kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes"
	pipe "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/pipe"
	static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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
// Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto) in Envoy terminology.
// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin. (plugins currently need to be compiled into Gloo)
type Upstream struct {
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// Upstreams and their configuration can be automatically by Gloo Discovery
	// if this upstream is created or modified by Discovery, metadata about the operation will be placed here.
	DiscoveryMetadata *DiscoveryMetadata `protobuf:"bytes,3,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	SslConfig         *UpstreamSslConfig `protobuf:"bytes,4,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Circuit breakers for this upstream. if not set, the defaults ones from the Gloo settings will be used.
	// if those are not set, [envoy's defaults](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
	// will be used.
	CircuitBreakers    *CircuitBreakerConfig     `protobuf:"bytes,5,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	LoadBalancerConfig *LoadBalancerConfig       `protobuf:"bytes,6,opt,name=load_balancer_config,json=loadBalancerConfig,proto3" json:"load_balancer_config,omitempty"`
	ConnectionConfig   *ConnectionConfig         `protobuf:"bytes,7,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	HealthChecks       []*core1.HealthCheck      `protobuf:"bytes,8,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	OutlierDetection   *cluster.OutlierDetection `protobuf:"bytes,9,opt,name=outlier_detection,json=outlierDetection,proto3" json:"outlier_detection,omitempty"`
	// Use http2 when communicating with this upstream
	// this field is evaluated `true` for upstreams
	// with a grpc service spec. otherwise defaults to `false`
	UseHttp2 bool `protobuf:"varint,10,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"`
	// Note to developers: new Upstream plugins must be added to this oneof field
	// to be usable by Gloo. (plugins currently need to be compiled into Gloo)
	//
	// Types that are valid to be assigned to UpstreamType:
	//	*Upstream_Kube
	//	*Upstream_Static
	//	*Upstream_Pipe
	//	*Upstream_Aws
	//	*Upstream_Azure
	//	*Upstream_Consul
	//	*Upstream_AwsEc2
	UpstreamType         isUpstream_UpstreamType `protobuf_oneof:"upstream_type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Upstream) Reset()         { *m = Upstream{} }
func (m *Upstream) String() string { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()    {}
func (*Upstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{0}
}
func (m *Upstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upstream.Unmarshal(m, b)
}
func (m *Upstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upstream.Marshal(b, m, deterministic)
}
func (m *Upstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upstream.Merge(m, src)
}
func (m *Upstream) XXX_Size() int {
	return xxx_messageInfo_Upstream.Size(m)
}
func (m *Upstream) XXX_DiscardUnknown() {
	xxx_messageInfo_Upstream.DiscardUnknown(m)
}

var xxx_messageInfo_Upstream proto.InternalMessageInfo

type isUpstream_UpstreamType interface {
	isUpstream_UpstreamType()
	Equal(interface{}) bool
}

type Upstream_Kube struct {
	Kube *kubernetes.UpstreamSpec `protobuf:"bytes,11,opt,name=kube,proto3,oneof" json:"kube,omitempty"`
}
type Upstream_Static struct {
	Static *static.UpstreamSpec `protobuf:"bytes,12,opt,name=static,proto3,oneof" json:"static,omitempty"`
}
type Upstream_Pipe struct {
	Pipe *pipe.UpstreamSpec `protobuf:"bytes,13,opt,name=pipe,proto3,oneof" json:"pipe,omitempty"`
}
type Upstream_Aws struct {
	Aws *aws.UpstreamSpec `protobuf:"bytes,14,opt,name=aws,proto3,oneof" json:"aws,omitempty"`
}
type Upstream_Azure struct {
	Azure *azure.UpstreamSpec `protobuf:"bytes,15,opt,name=azure,proto3,oneof" json:"azure,omitempty"`
}
type Upstream_Consul struct {
	Consul *consul.UpstreamSpec `protobuf:"bytes,16,opt,name=consul,proto3,oneof" json:"consul,omitempty"`
}
type Upstream_AwsEc2 struct {
	AwsEc2 *ec2.UpstreamSpec `protobuf:"bytes,17,opt,name=aws_ec2,json=awsEc2,proto3,oneof" json:"aws_ec2,omitempty"`
}

func (*Upstream_Kube) isUpstream_UpstreamType()   {}
func (*Upstream_Static) isUpstream_UpstreamType() {}
func (*Upstream_Pipe) isUpstream_UpstreamType()   {}
func (*Upstream_Aws) isUpstream_UpstreamType()    {}
func (*Upstream_Azure) isUpstream_UpstreamType()  {}
func (*Upstream_Consul) isUpstream_UpstreamType() {}
func (*Upstream_AwsEc2) isUpstream_UpstreamType() {}

func (m *Upstream) GetUpstreamType() isUpstream_UpstreamType {
	if m != nil {
		return m.UpstreamType
	}
	return nil
}

func (m *Upstream) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Upstream) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Upstream) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

func (m *Upstream) GetSslConfig() *UpstreamSslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *Upstream) GetCircuitBreakers() *CircuitBreakerConfig {
	if m != nil {
		return m.CircuitBreakers
	}
	return nil
}

func (m *Upstream) GetLoadBalancerConfig() *LoadBalancerConfig {
	if m != nil {
		return m.LoadBalancerConfig
	}
	return nil
}

func (m *Upstream) GetConnectionConfig() *ConnectionConfig {
	if m != nil {
		return m.ConnectionConfig
	}
	return nil
}

func (m *Upstream) GetHealthChecks() []*core1.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *Upstream) GetOutlierDetection() *cluster.OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *Upstream) GetUseHttp2() bool {
	if m != nil {
		return m.UseHttp2
	}
	return false
}

func (m *Upstream) GetKube() *kubernetes.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *Upstream) GetStatic() *static.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Static); ok {
		return x.Static
	}
	return nil
}

func (m *Upstream) GetPipe() *pipe.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Pipe); ok {
		return x.Pipe
	}
	return nil
}

func (m *Upstream) GetAws() *aws.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Upstream) GetAzure() *azure.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Upstream) GetConsul() *consul.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Consul); ok {
		return x.Consul
	}
	return nil
}

func (m *Upstream) GetAwsEc2() *ec2.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_AwsEc2); ok {
		return x.AwsEc2
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Upstream) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Upstream_Kube)(nil),
		(*Upstream_Static)(nil),
		(*Upstream_Pipe)(nil),
		(*Upstream_Aws)(nil),
		(*Upstream_Azure)(nil),
		(*Upstream_Consul)(nil),
		(*Upstream_AwsEc2)(nil),
	}
}

// created by discovery services
type DiscoveryMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto", fileDescriptor_b74df493149f644d)
}

var fileDescriptor_b74df493149f644d = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdd, 0x4e, 0x33, 0x45,
	0x1c, 0xc6, 0xdf, 0xbe, 0x2c, 0xd0, 0x0e, 0x20, 0xed, 0x48, 0x74, 0x83, 0x11, 0x9a, 0x1e, 0x28,
	0x9a, 0x30, 0x2b, 0x35, 0x31, 0xa6, 0x06, 0x43, 0x5a, 0x30, 0x8d, 0x88, 0x9a, 0x25, 0xc6, 0xc4,
	0x93, 0xcd, 0x74, 0x3a, 0xb4, 0x63, 0x97, 0x9d, 0xcd, 0xce, 0x6c, 0x01, 0x0f, 0xb9, 0x1a, 0x2f,
	0xc1, 0x1b, 0x30, 0xf1, 0x2a, 0x38, 0xf0, 0x0e, 0x30, 0xf1, 0xfc, 0xcd, 0x7c, 0x95, 0xb6, 0x70,
	0xb0, 0x2d, 0x07, 0xdd, 0xd9, 0x99, 0x79, 0x9e, 0xdf, 0xcc, 0xfe, 0x77, 0xf6, 0x49, 0xc1, 0x37,
	0x03, 0x26, 0x87, 0x79, 0x0f, 0x11, 0x7e, 0x1d, 0x08, 0x1e, 0xf3, 0x43, 0xc6, 0x83, 0x41, 0xcc,
	0x79, 0x90, 0x66, 0xfc, 0x77, 0x4a, 0xa4, 0x30, 0x3d, 0x9c, 0xb2, 0x60, 0x7c, 0x14, 0xe4, 0xa9,
	0x90, 0x19, 0xc5, 0xd7, 0x28, 0xcd, 0xb8, 0xe4, 0x70, 0x53, 0xcd, 0x21, 0x65, 0x43, 0x8c, 0xef,
	0xee, 0x0c, 0xf8, 0x80, 0xeb, 0x89, 0x40, 0xdd, 0x19, 0xcd, 0x2e, 0xa4, 0xb7, 0xd2, 0x0c, 0xd2,
	0x5b, 0x69, 0xc7, 0x8e, 0x5e, 0x58, 0x54, 0xb7, 0x23, 0x26, 0xdd, 0x52, 0xd7, 0x54, 0xe2, 0x3e,
	0x96, 0xd8, 0x5a, 0xbe, 0x5a, 0x68, 0x9f, 0x42, 0xc4, 0xd6, 0xd7, 0x5e, 0xc8, 0x47, 0x58, 0x46,
	0x72, 0x26, 0xa3, 0x5e, 0x46, 0xf1, 0x88, 0x66, 0x96, 0x71, 0xb2, 0x10, 0x23, 0xe6, 0xb8, 0x1f,
	0xf5, 0x70, 0x8c, 0x13, 0x32, 0x21, 0x1c, 0x2f, 0xb6, 0x0b, 0x9e, 0x24, 0x94, 0x48, 0xc6, 0x13,
	0x6b, 0xff, 0xb5, 0xb8, 0x9d, 0xde, 0x4a, 0x9a, 0x25, 0x38, 0x0e, 0x68, 0x32, 0xe6, 0x77, 0x86,
	0xd8, 0x0c, 0x08, 0xcf, 0x68, 0x30, 0xa4, 0x38, 0x96, 0xc3, 0x88, 0x0c, 0x29, 0x19, 0x59, 0x70,
	0x50, 0xe0, 0x45, 0x08, 0x89, 0x65, 0x2e, 0xac, 0x21, 0x7a, 0xf5, 0x4e, 0xe2, 0x5c, 0x48, 0x9a,
	0x05, 0x3c, 0x97, 0x31, 0xa3, 0x59, 0xd4, 0xa7, 0x72, 0xe6, 0x51, 0x8b, 0x1c, 0x0d, 0xd7, 0xb7,
	0x96, 0xee, 0x42, 0xc5, 0xe5, 0xa9, 0x5a, 0x4d, 0xe8, 0xc7, 0x62, 0xc4, 0x36, 0x96, 0x74, 0xba,
	0x14, 0x29, 0x65, 0x29, 0xd5, 0x17, 0x4b, 0xf9, 0x79, 0x29, 0xca, 0x28, 0xef, 0xd1, 0x2c, 0xa1,
	0x92, 0x4e, 0xdf, 0x2e, 0x75, 0x88, 0x1d, 0x11, 0xdf, 0xe8, 0x9f, 0x65, 0x7c, 0xb7, 0x1c, 0xe3,
	0x8f, 0x3c, 0xa3, 0xe6, 0xfa, 0xaa, 0x6a, 0x13, 0x9e, 0x88, 0x3c, 0xb6, 0x8d, 0x25, 0x7d, 0xbf,
	0xf4, 0x53, 0x51, 0xd2, 0x54, 0x6d, 0x44, 0x49, 0xd3, 0xb2, 0x5a, 0xcb, 0xb0, 0x8c, 0xb7, 0xf1,
	0x77, 0x05, 0x94, 0x7f, 0xb1, 0xc1, 0x06, 0xcf, 0xc1, 0x9a, 0x39, 0xf0, 0x7e, 0xa9, 0x5e, 0x3a,
	0xd8, 0x68, 0xee, 0x20, 0xf5, 0xed, 0xb8, 0x8c, 0x43, 0x97, 0x7a, 0xae, 0xfd, 0xf1, 0x5f, 0xff,
	0x7b, 0xa5, 0x7f, 0x1e, 0xf6, 0xdf, 0xfc, 0xf7, 0xb0, 0x5f, 0x93, 0x54, 0xc8, 0x3e, 0xbb, 0xba,
	0x6a, 0x35, 0xd8, 0x20, 0xe1, 0x19, 0x6d, 0x84, 0x16, 0x01, 0xbf, 0x06, 0x65, 0x17, 0x63, 0xfe,
	0x5b, 0x8d, 0xfb, 0x60, 0x16, 0x77, 0x61, 0x67, 0xdb, 0x9e, 0x82, 0x85, 0x13, 0x35, 0xfc, 0x11,
	0xc0, 0x3e, 0x13, 0x84, 0x8f, 0x69, 0x76, 0x17, 0x4d, 0x18, 0x2b, 0x9a, 0xb1, 0x8f, 0xa6, 0x63,
	0x17, 0x9d, 0x3a, 0x9d, 0x83, 0x85, 0xb5, 0xfe, 0xfc, 0x10, 0xfc, 0x16, 0x00, 0x21, 0xe2, 0x88,
	0xf0, 0xe4, 0x8a, 0x0d, 0x7c, 0xef, 0x25, 0x8e, 0x2b, 0xc1, 0xa5, 0x88, 0x3b, 0x5a, 0x16, 0x56,
	0x84, 0xbb, 0x85, 0x17, 0xa0, 0x3a, 0x97, 0x8d, 0xc2, 0x5f, 0xd5, 0x94, 0xc6, 0x2c, 0xa5, 0x63,
	0x54, 0x6d, 0x23, 0xb2, 0xa0, 0x6d, 0x32, 0x33, 0x2a, 0x60, 0x08, 0x76, 0x66, 0x62, 0xd2, 0x6d,
	0x6c, 0x4d, 0x23, 0xeb, 0xb3, 0xc8, 0x1f, 0x38, 0xee, 0xb7, 0xad, 0xd0, 0x02, 0x61, 0xfc, 0x6c,
	0x0c, 0x9e, 0x83, 0xda, 0x53, 0x70, 0x3a, 0xe0, 0xba, 0x06, 0xee, 0xcd, 0xed, 0x71, 0x22, 0xb3,
	0xb8, 0x2a, 0x99, 0x1b, 0x81, 0x1d, 0xb0, 0x35, 0x1d, 0x97, 0xc2, 0x2f, 0xd7, 0x57, 0x34, 0x48,
	0x87, 0x19, 0xc2, 0x29, 0x43, 0xe3, 0xa6, 0x79, 0x97, 0x5d, 0xad, 0xeb, 0x28, 0x59, 0xb8, 0x39,
	0x7c, 0xea, 0x08, 0x78, 0x09, 0x6a, 0xcf, 0x62, 0xce, 0xaf, 0xe8, 0x1d, 0x7d, 0x32, 0x07, 0x32,
	0xa9, 0x88, 0x7e, 0x32, 0xf2, 0x53, 0xa7, 0x0e, 0xab, 0x7c, 0x6e, 0x04, 0x7e, 0x04, 0x2a, 0xb9,
	0xa0, 0xd1, 0x50, 0xca, 0xb4, 0xe9, 0x83, 0x7a, 0xe9, 0xa0, 0x1c, 0x96, 0x73, 0x41, 0xbb, 0xaa,
	0x0f, 0x3b, 0xc0, 0x53, 0xe1, 0xe1, 0x6f, 0xe8, 0x45, 0x0e, 0xd1, 0x54, 0x92, 0xb8, 0x33, 0xff,
	0xf2, 0x3b, 0x4f, 0x29, 0xe9, 0xbe, 0x09, 0xb5, 0x19, 0x76, 0xcc, 0x27, 0xc0, 0x88, 0xbf, 0xa9,
	0x31, 0x9f, 0x21, 0x1b, 0x92, 0x45, 0x10, 0xd6, 0x0a, 0x8f, 0x81, 0xa7, 0x22, 0xd1, 0xdf, 0xd2,
	0x88, 0x4f, 0x91, 0xce, 0xc7, 0x42, 0x7b, 0x50, 0x4a, 0xd8, 0x02, 0x2b, 0xf8, 0x46, 0xf8, 0xef,
	0xd9, 0x62, 0xa9, 0x18, 0x2b, 0x62, 0x56, 0x26, 0x78, 0x02, 0x56, 0x75, 0x60, 0xf9, 0xdb, 0xda,
	0x7d, 0x80, 0x4c, 0x7c, 0x15, 0xf1, 0x1b, 0xa3, 0xaa, 0x80, 0x49, 0x2a, 0xbf, 0x6a, 0x2b, 0x60,
	0x83, 0xab, 0x50, 0x05, 0x8c, 0x16, 0x9e, 0x81, 0x75, 0x9b, 0x51, 0x7e, 0x4d, 0x53, 0x3e, 0x47,
	0x2e, 0xb3, 0x0a, 0x61, 0xf0, 0x8d, 0x38, 0x23, 0xcd, 0xd6, 0x87, 0xf7, 0x8f, 0x9e, 0x07, 0xde,
	0xe6, 0xe2, 0xfe, 0xd1, 0xdb, 0x80, 0x15, 0xf7, 0x0f, 0x4c, 0xb4, 0xb7, 0xc1, 0x96, 0xeb, 0x44,
	0xf2, 0x2e, 0xa5, 0x8d, 0xf7, 0x41, 0xed, 0x59, 0x16, 0xb4, 0x5b, 0x2a, 0xa9, 0xfe, 0xfc, 0x77,
	0xaf, 0xf4, 0xdb, 0x17, 0xc5, 0x22, 0x32, 0x1d, 0x0d, 0x6c, 0x4c, 0xf6, 0xd6, 0x74, 0x3e, 0x7e,
	0xf9, 0x2e, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x96, 0x33, 0x34, 0x24, 0x0a, 0x00, 0x00,
}

func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.CircuitBreakers.Equal(that1.CircuitBreakers) {
		return false
	}
	if !this.LoadBalancerConfig.Equal(that1.LoadBalancerConfig) {
		return false
	}
	if !this.ConnectionConfig.Equal(that1.ConnectionConfig) {
		return false
	}
	if len(this.HealthChecks) != len(that1.HealthChecks) {
		return false
	}
	for i := range this.HealthChecks {
		if !this.HealthChecks[i].Equal(that1.HealthChecks[i]) {
			return false
		}
	}
	if !this.OutlierDetection.Equal(that1.OutlierDetection) {
		return false
	}
	if this.UseHttp2 != that1.UseHttp2 {
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
func (this *Upstream_Kube) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Kube)
	if !ok {
		that2, ok := that.(Upstream_Kube)
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
func (this *Upstream_Static) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Static)
	if !ok {
		that2, ok := that.(Upstream_Static)
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
func (this *Upstream_Pipe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Pipe)
	if !ok {
		that2, ok := that.(Upstream_Pipe)
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
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	return true
}
func (this *Upstream_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Aws)
	if !ok {
		that2, ok := that.(Upstream_Aws)
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
func (this *Upstream_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Azure)
	if !ok {
		that2, ok := that.(Upstream_Azure)
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
func (this *Upstream_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Consul)
	if !ok {
		that2, ok := that.(Upstream_Consul)
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
func (this *Upstream_AwsEc2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_AwsEc2)
	if !ok {
		that2, ok := that.(Upstream_AwsEc2)
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
	if !this.AwsEc2.Equal(that1.AwsEc2) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
