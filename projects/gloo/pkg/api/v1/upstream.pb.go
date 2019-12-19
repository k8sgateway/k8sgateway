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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xcf, 0x4e, 0x23, 0x37,
	0x1c, 0xc7, 0x09, 0x04, 0x48, 0x0c, 0x94, 0xc4, 0x45, 0xed, 0x88, 0x56, 0x10, 0xe5, 0xd0, 0xd2,
	0x4a, 0x78, 0x4a, 0x2a, 0x55, 0x6d, 0x2a, 0x2a, 0x34, 0x81, 0x2a, 0xa2, 0xa5, 0xad, 0x06, 0x55,
	0x2b, 0xed, 0x65, 0xe4, 0x38, 0x26, 0xf1, 0x66, 0x18, 0x8f, 0xc6, 0x9e, 0x00, 0x7b, 0xe4, 0x69,
	0xf6, 0x41, 0xf6, 0xb0, 0x4f, 0xc1, 0x61, 0xdf, 0x80, 0x3d, 0x71, 0x5c, 0x8d, 0xff, 0x84, 0x24,
	0x70, 0x98, 0x84, 0x43, 0xc6, 0x63, 0xfb, 0xfb, 0xfd, 0xd8, 0xf3, 0x1b, 0xcf, 0x57, 0x01, 0xbf,
	0xf7, 0x98, 0xec, 0xa7, 0x1d, 0x44, 0xf8, 0xa5, 0x2b, 0x78, 0xc8, 0xf7, 0x19, 0x77, 0x7b, 0x21,
	0xe7, 0x6e, 0x9c, 0xf0, 0x37, 0x94, 0x48, 0xa1, 0x7b, 0x38, 0x66, 0xee, 0xf0, 0xc0, 0x4d, 0x63,
	0x21, 0x13, 0x8a, 0x2f, 0x51, 0x9c, 0x70, 0xc9, 0xe1, 0x7a, 0x36, 0x87, 0x32, 0x1b, 0x62, 0x7c,
	0x7b, 0xab, 0xc7, 0x7b, 0x5c, 0x4d, 0xb8, 0xd9, 0x9d, 0xd6, 0x6c, 0x43, 0x7a, 0x2d, 0xf5, 0x20,
	0xbd, 0x96, 0x66, 0xec, 0xe0, 0x99, 0x45, 0x55, 0x3b, 0x60, 0xd2, 0x2e, 0x75, 0x49, 0x25, 0xee,
	0x62, 0x89, 0x8d, 0xe5, 0x97, 0x99, 0xf6, 0x29, 0x44, 0x68, 0x7c, 0xde, 0x4c, 0x3e, 0xc2, 0x12,
	0x92, 0x32, 0x19, 0x74, 0x12, 0x8a, 0x07, 0x34, 0x31, 0x8c, 0xa3, 0x99, 0x18, 0x21, 0xc7, 0xdd,
	0xa0, 0x83, 0x43, 0x1c, 0x91, 0x11, 0xe1, 0x70, 0xb6, 0x5d, 0xf0, 0x28, 0xa2, 0x44, 0x32, 0x1e,
	0x19, 0xfb, 0xab, 0xfc, 0x76, 0x7a, 0x2d, 0x69, 0x12, 0xe1, 0xd0, 0xa5, 0xd1, 0x90, 0xdf, 0x68,
	0x62, 0xc3, 0x25, 0x3c, 0xa1, 0x6e, 0x9f, 0xe2, 0x50, 0xf6, 0x03, 0xd2, 0xa7, 0x64, 0x60, 0xc0,
	0x6e, 0x8e, 0x17, 0x21, 0x24, 0x96, 0xa9, 0x30, 0x86, 0xe0, 0xc5, 0x3b, 0x09, 0x53, 0x21, 0x69,
	0xe2, 0xf2, 0x54, 0x86, 0x8c, 0x26, 0x41, 0x97, 0xca, 0x89, 0x47, 0xcd, 0x73, 0x34, 0x6c, 0xdf,
	0x58, 0xda, 0x33, 0x15, 0x97, 0xc7, 0xd9, 0x6a, 0x42, 0x3d, 0x16, 0x23, 0xa6, 0x31, 0xa4, 0xe3,
	0xb9, 0x48, 0x31, 0x8b, 0xa9, 0xba, 0x18, 0xca, 0x7f, 0x73, 0x51, 0x06, 0x69, 0x87, 0x26, 0x11,
	0x95, 0x74, 0xfc, 0x76, 0xae, 0x43, 0x6c, 0x89, 0xf8, 0x4a, 0xfd, 0x0c, 0xe3, 0xcf, 0xf9, 0x18,
	0x6f, 0xd3, 0x84, 0xea, 0xeb, 0x8b, 0xaa, 0x4d, 0x78, 0x24, 0xd2, 0xd0, 0x34, 0x86, 0x74, 0x3a,
	0xf7, 0x53, 0x51, 0xd2, 0xc8, 0xda, 0x80, 0x92, 0x86, 0x61, 0x35, 0xe7, 0x61, 0x69, 0x6f, 0xfd,
	0x7d, 0x19, 0x94, 0xfe, 0x37, 0xc1, 0x06, 0x4f, 0xc1, 0x8a, 0x3e, 0xf0, 0x4e, 0xa1, 0x56, 0xd8,
	0x5b, 0x6b, 0x6c, 0xa1, 0xec, 0xdb, 0xb1, 0x19, 0x87, 0xce, 0xd5, 0x9c, 0xf7, 0xed, 0x83, 0x57,
	0xf8, 0x70, 0xb7, 0xbb, 0xf0, 0xe9, 0x6e, 0xb7, 0x2a, 0xa9, 0x90, 0x5d, 0x76, 0x71, 0xd1, 0xac,
	0xb3, 0x5e, 0xc4, 0x13, 0x5a, 0xf7, 0x0d, 0x01, 0xfe, 0x0a, 0x4a, 0x36, 0xc5, 0x9c, 0x45, 0x45,
	0xfb, 0x6a, 0x92, 0x76, 0x66, 0x66, 0xbd, 0x62, 0x06, 0xf3, 0x47, 0x6a, 0xf8, 0x0f, 0x80, 0x5d,
	0x26, 0x08, 0x1f, 0xd2, 0xe4, 0x26, 0x18, 0x31, 0x96, 0x14, 0x63, 0x17, 0x8d, 0xa7, 0x2e, 0x3a,
	0xb6, 0x3a, 0x0b, 0xf3, 0xab, 0xdd, 0xe9, 0x21, 0xf8, 0x07, 0x00, 0x42, 0x84, 0x01, 0xe1, 0xd1,
	0x05, 0xeb, 0x39, 0xc5, 0xe7, 0x38, 0xb6, 0x02, 0xe7, 0x22, 0x6c, 0x29, 0x99, 0x5f, 0x16, 0xf6,
	0x16, 0x9e, 0x81, 0xca, 0x54, 0x34, 0x0a, 0x67, 0x59, 0x51, 0xea, 0x93, 0x94, 0x96, 0x56, 0x79,
	0x5a, 0x64, 0x40, 0x9b, 0x64, 0x62, 0x54, 0x40, 0x1f, 0x6c, 0x4d, 0xa4, 0xa4, 0xdd, 0xd8, 0x8a,
	0x42, 0xd6, 0x26, 0x91, 0x7f, 0x73, 0xdc, 0xf5, 0x8c, 0xd0, 0x00, 0x61, 0xf8, 0x64, 0x0c, 0xfe,
	0x05, 0xaa, 0x8f, 0xb9, 0x69, 0x81, 0xab, 0x0a, 0xb8, 0x33, 0xb5, 0xc7, 0x91, 0xcc, 0xe0, 0x2a,
	0x64, 0x6a, 0x04, 0xb6, 0xc0, 0xc6, 0x78, 0x5a, 0x0a, 0xa7, 0x54, 0x5b, 0x52, 0x20, 0x95, 0x65,
	0x08, 0xc7, 0x0c, 0x0d, 0x1b, 0xfa, 0x5d, 0xb6, 0x95, 0xae, 0x95, 0xc9, 0xfc, 0xf5, 0xfe, 0x63,
	0x47, 0xc0, 0x73, 0x50, 0x7d, 0x92, 0x72, 0x4e, 0x59, 0xed, 0xe8, 0xbb, 0x29, 0x90, 0x0e, 0x45,
	0xf4, 0xaf, 0x96, 0x1f, 0x5b, 0xb5, 0x5f, 0xe1, 0x53, 0x23, 0xf0, 0x1b, 0x50, 0x4e, 0x05, 0x0d,
	0xfa, 0x52, 0xc6, 0x0d, 0x07, 0xd4, 0x0a, 0x7b, 0x25, 0xbf, 0x94, 0x0a, 0xda, 0xce, 0xfa, 0xb0,
	0x05, 0x8a, 0x59, 0x76, 0x38, 0x6b, 0x6a, 0x91, 0x7d, 0x34, 0x16, 0x24, 0xf6, 0xc8, 0x3f, 0xff,
	0xce, 0x63, 0x4a, 0xda, 0x0b, 0xbe, 0x32, 0xc3, 0x96, 0xfe, 0x02, 0x18, 0x71, 0xd6, 0x15, 0xe6,
	0x07, 0x64, 0x32, 0x32, 0x0f, 0xc2, 0x58, 0xe1, 0x21, 0x28, 0x66, 0x89, 0xe8, 0x6c, 0x28, 0xc4,
	0xf7, 0x48, 0xc5, 0x63, 0xae, 0x3d, 0x64, 0x4a, 0xd8, 0x04, 0x4b, 0xf8, 0x4a, 0x38, 0x5f, 0x98,
	0x62, 0x65, 0x29, 0x96, 0xc7, 0x9c, 0x99, 0xe0, 0x11, 0x58, 0x56, 0x79, 0xe5, 0x6c, 0x2a, 0xf7,
	0x1e, 0xd2, 0xe9, 0x95, 0xc7, 0xaf, 0x8d, 0x59, 0x05, 0x74, 0x50, 0x39, 0x15, 0x53, 0x01, 0x93,
	0x5b, 0xb9, 0x2a, 0xa0, 0xb5, 0xf0, 0x04, 0xac, 0x9a, 0x88, 0x72, 0xaa, 0x8a, 0xf2, 0x23, 0xb2,
	0x91, 0x95, 0x0b, 0x83, 0xaf, 0xc4, 0x09, 0x69, 0x34, 0xbf, 0xbe, 0xbd, 0x2f, 0x16, 0xc1, 0x62,
	0x2a, 0x6e, 0xef, 0x8b, 0x6b, 0xb0, 0x6c, 0xff, 0x80, 0x09, 0x6f, 0x13, 0x6c, 0xd8, 0x4e, 0x20,
	0x6f, 0x62, 0x5a, 0xff, 0x12, 0x54, 0x9f, 0x64, 0x81, 0xf7, 0xdb, 0x83, 0x57, 0x78, 0xf7, 0x71,
	0xa7, 0xf0, 0xfa, 0xa7, 0x7c, 0x01, 0x19, 0x0f, 0x7a, 0x26, 0x24, 0x3b, 0x2b, 0x2a, 0x1d, 0x7f,
	0xfe, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x62, 0xbd, 0x13, 0xda, 0x22, 0x0a, 0x00, 0x00,
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
