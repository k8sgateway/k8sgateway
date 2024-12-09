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

type HttpBackendPolicy interface {
	markAdBackendPolicy()
}

type AttachableToHttpBackendPolicy struct{}

func (AttachableToHttpBackendPolicy) markAdBackendPolicy() {}

//type AttachedPolicies map[string]Policies

// type NetworkPolicy AttachedPolicy
// type HttpPolicy AttachedPolicy
// type UpstreamPolicy AttachedPolicy

// type ListenerPolicy AttachedPolicy

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

/*
(aws) upstream plugin:

	ContributesPolicies map[GroupKind:"kgw/Parameters"]struct {
		AttachmentPoints          []{BackendAttachmentPoint}
		NewGatewayTranslationPass func(ctx context.Context, tctx GwTranslationCtx) ProxyTranslationPass{

		ProcessBackend: func(ctx context.Context, Backend, RefPolicy) ProxyTranslationPass{
			// check backend upstream to be aws
			// check ref policy to be aws
		}
		Policies                  krt.Collection[model.Policy]
		PoliciesFetch(name, namespace) Policy {return RefPolicy{...}}
	}

	ContributesUpstreams map[GroupKind:"kgw/Upstream"]struct {
		ProcessUpstream: func(ctx context.Context, in model.Upstream, out *envoy_config_cluster_v3.Cluster){
			ourUs, ok := in.Obj.(*kgw.Upstream)
			if !ok {
				// log - should never happen
				return
			}
			if ourUs.aws != nil {
				do stuff and update the cluster
			}
		}
		Upstreams       krt.Collection[model.Upstream]
		Endpoints       []krt.Collection[krtcollections.EndpointsForUpstream]
	}
	ContributesGwClasses map[string]translator.K8sGwTranslator
*/
type HttpBackend struct {
	Backend
	AttachedPolicies
}

type HttpRouteIR struct {
	ObjectSource     `json:",inline"`
	SourceObject     client.Object
	ParentRefs       []gwv1.ParentReference
	Hostnames        []string
	AttachedPolicies AttachedPolicies
	Rules            []HttpRouteRuleIR
}

// type HttpRouteRuleIR struct {
// 	gwv1.HTTPRouteRule
// 	Parent           HttpRouteIR
// 	ExtensionRefs    AttachedPolicies
// 	AttachedPolicies AttachedPolicies
//
// 	Backends []HttpBackend
// }

// TODO: this is the structure we probably want,
// and maybe change name -- it's not a Rule, it's a Match
type HttpRouteRuleIR struct {
	Match            gwv1.HTTPRouteMatch
	MatchIndex       int
	Parent           *HttpRouteIR
	SourceRule       *gwv1.HTTPRouteRule
	Name             string
	ExtensionRefs    AttachedPolicies
	AttachedPolicies AttachedPolicies
	Backends         []HttpBackend
}

// TODO: temporary structure to represent an individual Match (equiv. to gloov1.Route)
// need to remove in favor of commented out HttpRouteRuleIR above
//type HttpRouteRuleMatchIR struct {
//	Match    gwv1.HTTPRouteMatch
//	Name     string // not sure if we actually need this anymore
//	Parent   HttpRouteRuleIR
//	Backends []HttpBackend
//}

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
	Rules     []HttpRouteRuleIR
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
