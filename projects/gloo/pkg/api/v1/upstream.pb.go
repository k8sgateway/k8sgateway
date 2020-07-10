// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
	UseHttp2 *types.BoolValue `protobuf:"bytes,10,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"`
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
	UpstreamType isUpstream_UpstreamType `protobuf_oneof:"upstream_type"`
	// Failover endpoints for this upstream. If omitted (the default) no failovers will be applied.
	Failover             *Failover `protobuf:"bytes,18,opt,name=failover,proto3" json:"failover,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *Upstream) GetUseHttp2() *types.BoolValue {
	if m != nil {
		return m.UseHttp2
	}
	return nil
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

func (m *Upstream) GetFailover() *Failover {
	if m != nil {
		return m.Failover
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
	// Labels inherited from the original upstream (e.g. Kubernetes labels)
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *DiscoveryMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.DiscoveryMetadata.LabelsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto", fileDescriptor_b74df493149f644d)
}

var fileDescriptor_b74df493149f644d = []byte{
	// 942 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xae, 0x13, 0x27, 0x8d, 0x27, 0x09, 0xb1, 0x47, 0x11, 0xac, 0x22, 0x48, 0x22, 0x23, 0xd1,
	0x50, 0x94, 0x59, 0xea, 0x0a, 0xb5, 0x18, 0x15, 0x21, 0x3b, 0x41, 0x91, 0x9a, 0x82, 0xb4, 0x11,
	0x5c, 0x70, 0xb3, 0x1a, 0x8f, 0xc7, 0xeb, 0xc1, 0x93, 0x9d, 0xd5, 0xce, 0xac, 0x13, 0x73, 0xd9,
	0x57, 0x40, 0xbc, 0x03, 0x8f, 0xc0, 0x23, 0xf0, 0x14, 0xbd, 0xe0, 0x0d, 0x8a, 0xc4, 0x3d, 0x9a,
	0x3f, 0xd7, 0x76, 0xea, 0x7a, 0x7b, 0x11, 0xef, 0xcc, 0x39, 0xdf, 0xf7, 0xed, 0xd9, 0xd9, 0x73,
	0xbe, 0x2c, 0xf8, 0x26, 0x61, 0x6a, 0x58, 0xf4, 0x10, 0x11, 0xd7, 0xa1, 0x14, 0x5c, 0x9c, 0x32,
	0x11, 0x26, 0x5c, 0x88, 0x30, 0xcb, 0xc5, 0xaf, 0x94, 0x28, 0x69, 0x77, 0x38, 0x63, 0xe1, 0xf8,
	0x51, 0x58, 0x64, 0x52, 0xe5, 0x14, 0x5f, 0xa3, 0x2c, 0x17, 0x4a, 0xc0, 0x1d, 0x9d, 0x43, 0x9a,
	0x86, 0x98, 0x38, 0xd8, 0x4f, 0x44, 0x22, 0x4c, 0x22, 0xd4, 0x2b, 0x8b, 0x39, 0x80, 0xf4, 0x56,
	0xd9, 0x20, 0xbd, 0x55, 0x2e, 0x76, 0x68, 0xee, 0x34, 0x62, 0xca, 0xeb, 0x5e, 0x53, 0x85, 0xfb,
	0x58, 0x61, 0x97, 0xff, 0x74, 0x79, 0x05, 0x52, 0x72, 0x07, 0x7a, 0x47, 0x99, 0x84, 0xe5, 0xa4,
	0x60, 0x2a, 0xee, 0xe5, 0x14, 0x8f, 0x68, 0xee, 0x08, 0xa7, 0xcb, 0x09, 0x5c, 0xe0, 0x7e, 0xdc,
	0xc3, 0x1c, 0xa7, 0x64, 0x0a, 0x7f, 0xf8, 0x0e, 0x7d, 0x91, 0xa6, 0x94, 0x28, 0x26, 0x52, 0x87,
	0x3d, 0x5b, 0x82, 0xa5, 0xb7, 0x8a, 0xe6, 0x29, 0xe6, 0x21, 0x4d, 0xc7, 0x62, 0x62, 0xe9, 0xad,
	0x90, 0x88, 0x9c, 0x86, 0x43, 0x8a, 0xb9, 0x1a, 0xc6, 0x64, 0x48, 0xc9, 0xc8, 0xa9, 0x7c, 0xbc,
	0x78, 0x2c, 0x52, 0x61, 0x55, 0x48, 0x97, 0xbd, 0x7c, 0xbf, 0x7b, 0xf0, 0x42, 0x2a, 0x9a, 0x87,
	0xa2, 0x50, 0x9c, 0xd1, 0x3c, 0xee, 0x53, 0x35, 0x57, 0xf1, 0x9d, 0x57, 0xe0, 0xf7, 0x2e, 0xff,
	0xd5, 0xf2, 0xa7, 0x17, 0x99, 0xd6, 0x91, 0xa6, 0x3a, 0x46, 0xdc, 0xc5, 0xd1, 0x1e, 0xad, 0xa6,
	0x65, 0x2c, 0xa3, 0xe6, 0xc7, 0x51, 0x9e, 0xad, 0xa6, 0x8c, 0x8a, 0x1e, 0xcd, 0x53, 0xaa, 0xe8,
	0xec, 0x72, 0x75, 0x1b, 0x78, 0x3a, 0xbe, 0x31, 0x7f, 0x8e, 0xf0, 0xb8, 0x04, 0xe1, 0xb7, 0x22,
	0xa7, 0xf6, 0xb7, 0xfc, 0x71, 0x10, 0x91, 0xca, 0x82, 0xbb, 0x8b, 0xa3, 0x3d, 0x29, 0x57, 0x1c,
	0x25, 0x2d, 0x7d, 0x8d, 0x29, 0x69, 0x39, 0xe2, 0x83, 0x95, 0x44, 0x07, 0x3c, 0x59, 0x0e, 0x1c,
	0x60, 0xc6, 0xc5, 0x78, 0xda, 0xcf, 0x87, 0x89, 0x10, 0x09, 0xa7, 0xa1, 0xd9, 0xf5, 0x8a, 0x41,
	0x78, 0x93, 0xe3, 0x2c, 0xa3, 0xb9, 0x53, 0x6a, 0xfe, 0x01, 0xc0, 0xd6, 0x4f, 0x6e, 0xbe, 0xe1,
	0x73, 0xb0, 0x69, 0x9b, 0x2f, 0xa8, 0x1c, 0x57, 0x4e, 0xb6, 0x5b, 0xfb, 0x48, 0x37, 0xad, 0x1f,
	0x75, 0x74, 0x65, 0x72, 0x9d, 0x4f, 0xfe, 0xfa, 0xaf, 0x5a, 0xf9, 0xfb, 0xd5, 0xd1, 0xbd, 0x7f,
	0x5f, 0x1d, 0x35, 0x14, 0x95, 0xaa, 0xcf, 0x06, 0x83, 0x76, 0x93, 0x25, 0xa9, 0xc8, 0x69, 0x33,
	0x72, 0x12, 0xf0, 0x29, 0xd8, 0xf2, 0x03, 0x1e, 0xac, 0x19, 0xb9, 0x0f, 0xe7, 0xe5, 0x5e, 0xb8,
	0x6c, 0xa7, 0xaa, 0xc5, 0xa2, 0x29, 0x1a, 0xfe, 0x00, 0x60, 0x9f, 0x49, 0xa2, 0x9f, 0x62, 0x12,
	0x4f, 0x35, 0xd6, 0x8d, 0xc6, 0x11, 0x9a, 0x75, 0x1f, 0x74, 0xe6, 0x71, 0x5e, 0x2c, 0x6a, 0xf4,
	0x17, 0x43, 0xf0, 0x5b, 0x00, 0xa4, 0xe4, 0x31, 0x11, 0xe9, 0x80, 0x25, 0x41, 0xf5, 0x6d, 0x3a,
	0xfe, 0x08, 0xae, 0x24, 0xef, 0x1a, 0x58, 0x54, 0x93, 0x7e, 0x09, 0x5f, 0x80, 0xfa, 0x82, 0xb7,
	0xc8, 0x60, 0xc3, 0xa8, 0x34, 0xe7, 0x55, 0xba, 0x16, 0xd5, 0xb1, 0x20, 0x27, 0xb4, 0x47, 0xe6,
	0xa2, 0x12, 0x46, 0x60, 0x7f, 0xce, 0x79, 0x7c, 0x61, 0x9b, 0x46, 0xf2, 0x78, 0x5e, 0xf2, 0x52,
	0xe0, 0x7e, 0xc7, 0x01, 0x9d, 0x20, 0xe4, 0x77, 0x62, 0xf0, 0x39, 0x68, 0xbc, 0xb1, 0x27, 0x2f,
	0x78, 0xdf, 0x08, 0x1e, 0x2e, 0xd4, 0x38, 0x85, 0x39, 0xb9, 0x3a, 0x59, 0x88, 0xc0, 0x2e, 0xd8,
	0x9d, 0xf5, 0x29, 0x19, 0x6c, 0x1d, 0xaf, 0x1b, 0x21, 0xe3, 0x35, 0x08, 0x67, 0x0c, 0x8d, 0x5b,
	0xf6, 0x5d, 0x5e, 0x18, 0x5c, 0x57, 0xc3, 0xa2, 0x9d, 0xe1, 0x9b, 0x8d, 0x84, 0x57, 0xa0, 0x71,
	0xc7, 0x85, 0x82, 0x9a, 0xa9, 0xe8, 0xb3, 0x05, 0x21, 0x6b, 0x5a, 0xe8, 0x47, 0x0b, 0x3f, 0xf3,
	0xe8, 0xa8, 0x2e, 0x16, 0x22, 0xf0, 0x09, 0xa8, 0x15, 0x92, 0xc6, 0x43, 0xa5, 0xb2, 0x56, 0x00,
	0x8c, 0xd8, 0x01, 0xb2, 0x1d, 0x8e, 0x7c, 0x87, 0xa3, 0x8e, 0x10, 0xfc, 0x67, 0xcc, 0x0b, 0x1a,
	0x6d, 0x15, 0x92, 0x5e, 0x68, 0x2c, 0xec, 0x82, 0xaa, 0xf6, 0x90, 0x60, 0xdb, 0x70, 0x4e, 0xd1,
	0x8c, 0xa1, 0xf8, 0xc9, 0x7a, 0x7b, 0x3f, 0x64, 0x94, 0x5c, 0xdc, 0x8b, 0x0c, 0x19, 0x76, 0xed,
	0x78, 0x30, 0x12, 0xec, 0x18, 0x99, 0xcf, 0x91, 0x73, 0xc1, 0x32, 0x12, 0x8e, 0x0a, 0x9f, 0x81,
	0xaa, 0xb6, 0xc1, 0x60, 0xd7, 0x48, 0x3c, 0x40, 0xc6, 0x13, 0x4b, 0xd5, 0xa0, 0x91, 0xb0, 0x0d,
	0xd6, 0xf1, 0x8d, 0x0c, 0x3e, 0x70, 0x07, 0xa9, 0x0d, 0xae, 0x0c, 0x59, 0x93, 0xe0, 0x77, 0x60,
	0xc3, 0xb8, 0x5b, 0xb0, 0x67, 0xd8, 0x27, 0xc8, 0x7a, 0x5d, 0x19, 0xbe, 0x25, 0xea, 0x13, 0xb0,
	0x4e, 0x17, 0xd4, 0xdd, 0x09, 0x38, 0xe3, 0x2b, 0x75, 0x02, 0x16, 0x0b, 0xcf, 0xc1, 0x7d, 0x67,
	0x7b, 0x41, 0xc3, 0xa8, 0x3c, 0x44, 0xde, 0x06, 0x4b, 0xc9, 0xe0, 0x1b, 0x79, 0x4e, 0x5a, 0xb0,
	0x05, 0xb6, 0xbc, 0xd7, 0x05, 0xd0, 0xf9, 0xcb, 0x1c, 0xef, 0x7b, 0x97, 0x8d, 0xa6, 0xb8, 0xf6,
	0x47, 0x2f, 0x5f, 0x57, 0xab, 0x60, 0xad, 0x90, 0x2f, 0x5f, 0x57, 0xb7, 0x61, 0xcd, 0x7f, 0xd8,
	0xc8, 0xce, 0x1e, 0xd8, 0xf5, 0x9b, 0x58, 0x4d, 0x32, 0xda, 0xfc, 0xbd, 0x02, 0x1a, 0x77, 0xcc,
	0x45, 0x3f, 0x3f, 0xc7, 0x3d, 0xca, 0xb5, 0x41, 0xea, 0x91, 0xf8, 0x62, 0x85, 0x1b, 0xa1, 0x4b,
	0x83, 0x3e, 0x4f, 0x55, 0x3e, 0x89, 0x1c, 0xf5, 0xe0, 0x6b, 0xb0, 0x3d, 0x13, 0x86, 0x75, 0xb0,
	0x3e, 0xa2, 0x13, 0xe3, 0xb8, 0xb5, 0x48, 0x2f, 0xe1, 0x3e, 0xd8, 0x18, 0xeb, 0xfe, 0x35, 0xb6,
	0x59, 0x8b, 0xec, 0xa6, 0xbd, 0xf6, 0xb4, 0xd2, 0x69, 0x6b, 0xeb, 0xfd, 0xf3, 0x9f, 0xc3, 0xca,
	0x2f, 0x5f, 0x96, 0xfb, 0x82, 0xcb, 0x46, 0x89, 0xfb, 0xc7, 0xd0, 0xdb, 0x34, 0x03, 0xf2, 0xf8,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xec, 0x34, 0x4e, 0xfc, 0x09, 0x00, 0x00,
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
	if !this.UseHttp2.Equal(that1.UseHttp2) {
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
	if !this.Failover.Equal(that1.Failover) {
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
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
