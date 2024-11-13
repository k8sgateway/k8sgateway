package snapshotadapter

import (
	"fmt"

	gateway_solo_io "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	enterprise_gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	graphql_gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1beta1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"istio.io/istio/pkg/kube/krt"
	"k8s.io/apimachinery/pkg/types"
)

type GlooCollection[T any] interface {
	Find(namespace string, name string) (T, error)
	List() []T
}

type SliceCollection[T interface {
	GetMetadata() *core.Metadata
}] []T

func (t SliceCollection[T]) Find(namespace string, name string) (T, error) {
	for _, t := range t {
		if t.GetMetadata().GetName() == name && t.GetMetadata().GetNamespace() == namespace {
			return t, nil
		}
	}
	var zero T
	return zero, fmt.Errorf("list did not find %T %v.%v", zero, namespace, name)
}

func (t SliceCollection[T]) List() []T {
	return t
}

func ToSlice[S ~[]T, T interface {
	GetMetadata() *core.Metadata
}](s S) SliceCollection[T] {
	return SliceCollection[T](([]T)(s))
}

type KrtCollection[T any] struct {
	C    krt.Collection[T]
	Kctx krt.HandlerContext
}

func (t KrtCollection[T]) Find(namespace string, name string) (T, error) {
	ret := krt.Fetch(t.Kctx, t.C, krt.FilterObjectName(types.NamespacedName{Name: name, Namespace: namespace}))
	if len(ret) != 1 {
		var zero T
		return zero, fmt.Errorf("list did not find %T %v.%v", zero, namespace, name)
	}
	return ret[0], nil
}

func (t KrtCollection[T]) List() []T {
	return krt.Fetch(t.Kctx, t.C)
}

type SnapCollection[T interface {
	GetMetadata() *core.Metadata
}] struct {
	list func() []T
}

func (t SnapCollection[T]) Find(namespace string, name string) (T, error) {
	for _, t := range t.list() {
		if t.GetMetadata().GetName() == name && t.GetMetadata().GetNamespace() == namespace {
			return t, nil
		}
	}
	var zero T
	return zero, fmt.Errorf("list did not find %T %v.%v", zero, namespace, name)
}

func (t SnapCollection[T]) List() []T {
	return t.list()
}

func FromApiSnapshot(snapshot *v1snap.ApiSnapshot) ApiSnapshot {
	return ApiSnapshot{
		Artifacts:      SnapCollection[*gloo_solo_io.Artifact]{list: func() []*gloo_solo_io.Artifact { return snapshot.Artifacts }},
		Endpoints:      SnapCollection[*gloo_solo_io.Endpoint]{list: func() []*gloo_solo_io.Endpoint { return snapshot.Endpoints }},
		Proxies:        SnapCollection[*gloo_solo_io.Proxy]{list: func() []*gloo_solo_io.Proxy { return snapshot.Proxies }},
		UpstreamGroups: SnapCollection[*gloo_solo_io.UpstreamGroup]{list: func() []*gloo_solo_io.UpstreamGroup { return snapshot.UpstreamGroups }},
		Secrets:        SnapCollection[*gloo_solo_io.Secret]{list: func() []*gloo_solo_io.Secret { return snapshot.Secrets }},
		Upstreams:      SnapCollection[*gloo_solo_io.Upstream]{list: func() []*gloo_solo_io.Upstream { return snapshot.Upstreams }},
		AuthConfigs:    SnapCollection[*enterprise_gloo_solo_io.AuthConfig]{list: func() []*enterprise_gloo_solo_io.AuthConfig { return snapshot.AuthConfigs }},
		Ratelimitconfigs: SnapCollection[*github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfig]{list: func() []*github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfig {
			return snapshot.Ratelimitconfigs
		}},
		VirtualServices:    SnapCollection[*gateway_solo_io.VirtualService]{list: func() []*gateway_solo_io.VirtualService { return snapshot.VirtualServices }},
		RouteTables:        SnapCollection[*gateway_solo_io.RouteTable]{list: func() []*gateway_solo_io.RouteTable { return snapshot.RouteTables }},
		Gateways:           SnapCollection[*gateway_solo_io.Gateway]{list: func() []*gateway_solo_io.Gateway { return snapshot.Gateways }},
		VirtualHostOptions: SnapCollection[*gateway_solo_io.VirtualHostOption]{list: func() []*gateway_solo_io.VirtualHostOption { return snapshot.VirtualHostOptions }},
		RouteOptions:       SnapCollection[*gateway_solo_io.RouteOption]{list: func() []*gateway_solo_io.RouteOption { return snapshot.RouteOptions }},
		HttpGateways:       SnapCollection[*gateway_solo_io.MatchableHttpGateway]{list: func() []*gateway_solo_io.MatchableHttpGateway { return snapshot.HttpGateways }},
		TcpGateways:        SnapCollection[*gateway_solo_io.MatchableTcpGateway]{list: func() []*gateway_solo_io.MatchableTcpGateway { return snapshot.TcpGateways }},
		GraphqlApis:        SnapCollection[*graphql_gloo_solo_io.GraphQLApi]{list: func() []*graphql_gloo_solo_io.GraphQLApi { return snapshot.GraphqlApis }},
	}
}

type ArtifactList = GlooCollection[*gloo_solo_io.Artifact]
type EndpointList = GlooCollection[*gloo_solo_io.Endpoint]
type ProxieList = GlooCollection[*gloo_solo_io.Proxy]
type UpstreamGroupList = GlooCollection[*gloo_solo_io.UpstreamGroup]
type SecretList = GlooCollection[*gloo_solo_io.Secret]
type UpstreamList = GlooCollection[*gloo_solo_io.Upstream]
type AuthConfigList = GlooCollection[*enterprise_gloo_solo_io.AuthConfig]
type RatelimitconfigList = GlooCollection[*github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfig]
type VirtualServiceList = GlooCollection[*gateway_solo_io.VirtualService]
type RouteTableList = GlooCollection[*gateway_solo_io.RouteTable]
type GatewayList = GlooCollection[*gateway_solo_io.Gateway]
type VirtualHostOptionList = GlooCollection[*gateway_solo_io.VirtualHostOption]
type RouteOptionList = GlooCollection[*gateway_solo_io.RouteOption]
type HttpGatewayList = GlooCollection[*gateway_solo_io.MatchableHttpGateway]
type TcpGatewayList = GlooCollection[*gateway_solo_io.MatchableTcpGateway]
type GraphqlApiList = GlooCollection[*graphql_gloo_solo_io.GraphQLApi]

type ApiSnapshot struct {
	Artifacts          GlooCollection[*gloo_solo_io.Artifact]
	Endpoints          GlooCollection[*gloo_solo_io.Endpoint]
	Proxies            GlooCollection[*gloo_solo_io.Proxy]
	UpstreamGroups     GlooCollection[*gloo_solo_io.UpstreamGroup]
	Secrets            GlooCollection[*gloo_solo_io.Secret]
	Upstreams          GlooCollection[*gloo_solo_io.Upstream]
	AuthConfigs        GlooCollection[*enterprise_gloo_solo_io.AuthConfig]
	Ratelimitconfigs   GlooCollection[*github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfig]
	VirtualServices    GlooCollection[*gateway_solo_io.VirtualService]
	RouteTables        GlooCollection[*gateway_solo_io.RouteTable]
	Gateways           GlooCollection[*gateway_solo_io.Gateway]
	VirtualHostOptions GlooCollection[*gateway_solo_io.VirtualHostOption]
	RouteOptions       GlooCollection[*gateway_solo_io.RouteOption]
	HttpGateways       GlooCollection[*gateway_solo_io.MatchableHttpGateway]
	TcpGateways        GlooCollection[*gateway_solo_io.MatchableTcpGateway]
	GraphqlApis        GlooCollection[*graphql_gloo_solo_io.GraphQLApi]
}
