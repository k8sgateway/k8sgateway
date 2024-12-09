package model

import (
	"context"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/solo-io/gloo/projects/controller/pkg/plugins"
	"google.golang.org/protobuf/types/known/anypb"
	"istio.io/istio/pkg/kube/krt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type AttachmentPoints int

const (
	HttpAttachmentPoint AttachmentPoints = iota
	HttpBackendRefAttachmentPoint
	ListenerAttachmentPoint
	UpstreamAttachmentPoint
)

type UpstreamInit struct {
	InitUpstream func(ctx context.Context, in Upstream, out *envoy_config_cluster_v3.Cluster)
}

type PolicyTargetRef struct {
	Group       string
	Kind        string
	Name        string
	SectionName string
}

type PolicyAtt struct {
	// original object. ideally with structural errors removed.
	// Opaque to us other than metadata.
	PolicyIr PolicyIR

	PolicyTargetRef PolicyTargetRef
}

func (c PolicyAtt) Obj() PolicyIR {
	return c.PolicyIr
}

func (c PolicyAtt) TargetRef() PolicyTargetRef {
	return c.PolicyTargetRef
}

type AttachedPolicies struct {
	Policies map[schema.GroupKind][]PolicyAtt
}

type Backend struct {
	ClusterName string
	Weight      uint32

	// upstream could be nil if not found or no ref grant
	Upstream *Upstream
	// if nil, error might say why
	Err error
}
type Route interface {
	GetGroupKind() schema.GroupKind
	// GetName returns the name of the route.
	GetName() string
	// GetNamespace returns the namespace of the route.
	GetNamespace() string

	GetParentRefs() []gwv1.ParentReference
	GetSourceObject() client.Object
}

// this is 1:1 with httproute, and is a krt type
// maybe move this to krtcollections package?
type HttpRouteIR struct {
	ObjectSource `json:",inline"`
	SourceObject client.Object
	ParentRefs   []gwv1.ParentReference

	Hostnames        []string
	AttachedPolicies AttachedPolicies
	Rules            []HttpRouteRuleIR
}

func (c *HttpRouteIR) GetParentRefs() []gwv1.ParentReference {
	return c.ParentRefs
}
func (c *HttpRouteIR) GetSourceObject() client.Object {
	return c.SourceObject
}

var _ Route = &HttpRouteIR{}

type TcpRouteIR struct {
	ObjectSource     `json:",inline"`
	SourceObject     client.Object
	ParentRefs       []gwv1.ParentReference
	AttachedPolicies AttachedPolicies
	Backends         []Backend
}

func (c *TcpRouteIR) GetParentRefs() []gwv1.ParentReference {
	return c.ParentRefs
}
func (c *TcpRouteIR) GetSourceObject() client.Object {
	return c.SourceObject
}

var _ Route = &TcpRouteIR{}

func (c HttpRouteIR) ResourceName() string {
	return c.ObjectSource.ResourceName()
}

func (c HttpRouteIR) Equals(in HttpRouteIR) bool {
	return c.ObjectSource == in.ObjectSource && versionEquals(c.SourceObject, in.SourceObject)
}

type HttpRouteRuleCommonIR struct {
	Parent           *HttpRouteIR
	SourceRule       *gwv1.HTTPRouteRule
	ExtensionRefs    AttachedPolicies
	AttachedPolicies AttachedPolicies
}

type HttpBackendOrDelegate struct {
	Backend  *Backend
	Delegate *ObjectSource
	AttachedPolicies
}

type HttpBackend struct {
	Backend Backend
	AttachedPolicies
}

type HttpRouteRuleIR struct {
	HttpRouteRuleCommonIR
	Backends []HttpBackendOrDelegate
	Matches  []gwv1.HTTPRouteMatch
	Name     string
}

type HttpRouteRuleMatchIR struct {
	HttpRouteRuleCommonIR
	Backends   []HttpBackend
	Match      gwv1.HTTPRouteMatch
	MatchIndex int
	Name       string
}

type ListenerIR struct {
	Name             string
	BindAddress      string
	BindPort         uint32
	AttachedPolicies AttachedPolicies

	HttpFilterChain []HttpFilterChainIR
	TcpFilterChain  []TcpIR
}

type VirtualHost struct {
	Name      string
	Hostnames []string
	Rules     []HttpRouteRuleMatchIR
}

type FilterChainMatch struct {
	SniDomains []string
}
type TlsBundle struct {
	//	CA            []byte
	PrivateKey    []byte
	CertChain     []byte
	AlpnProtocols []string
}

type FilterChainCommon struct {
	Matcher              FilterChainMatch
	FilterChainName      string
	ParentRef            gwv1.ParentReference
	CustomNetworkFilters []CustomEnvoyFilter
	TLS                  *TlsBundle
}
type CustomEnvoyFilter struct {
	// Determines filter ordering.
	FilterStage plugins.FilterStage[plugins.WellKnownFilterStage]
	// The name of the filter configuration.
	Name string
	// Filter specific configuration.
	Config *anypb.Any
}

type HttpFilterChainIR struct {
	FilterChainCommon
	Vhosts                  []*VirtualHost
	ParentRef               gwv1.ParentReference
	AttachedPolicies        AttachedPolicies
	AttachedNetworkPolicies AttachedPolicies
}

type TcpIR struct {
	FilterChainCommon
	BackendRefs []Backend
}

// this is 1:1 with envoy deployments
// not in a collection so doesn't need a krt interfaces.
type GatewayIR struct {
	Listeners    []ListenerIR
	SourceObject *gwv1.Gateway

	AttachedPolicies     AttachedPolicies
	AttachedHttpPolicies AttachedPolicies
}

type GatewayWithPoliciesIR struct {
	SourceObject *gwv1.Gateway

	AttachedPolicies     AttachedPolicies
	AttachedHttpPolicies AttachedPolicies
}

type Extension struct {
	ContributedUpstreams map[schema.GroupKind]krt.Collection[Upstream]
}
