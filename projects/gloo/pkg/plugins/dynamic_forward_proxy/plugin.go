package dynamic_forward_proxy

import (
	"fmt"

	envoy_extensions_network_dns_resolver_apple_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/network/dns_resolver/apple/v3"
	envoy_extensions_network_dns_resolver_cares_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/network/dns_resolver/cares/v3"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"

	"github.com/rotisserie/eris"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_clusters_dynamic_forward_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/dynamic_forward_proxy/v3"
	envoy_extensions_common_dynamic_forward_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/dynamic_forward_proxy/v3"
	envoy_extensions_filters_http_dynamic_forward_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamic_forward_proxy/v3"
	"github.com/golang/protobuf/ptypes/duration"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/dynamic_forward_proxy"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/hashutils"
)

var (
	_ plugins.Plugin                  = new(plugin)
	_ plugins.RoutePlugin             = new(plugin)
	_ plugins.HttpFilterPlugin        = new(plugin)
	_ plugins.ResourceGeneratorPlugin = new(plugin)
)

const (
	ExtensionName = "dynamic-forward-proxy"
	FilterName    = "envoy.filters.http.dynamic_forward_proxy"
)

var (
	pluginStage = plugins.DuringStage(plugins.OutAuthStage)
)

type plugin struct {
	filterHashMap map[string]*dynamic_forward_proxy.FilterConfig
}

func (p *plugin) GeneratedResources(params plugins.Params, inClusters []*envoy_config_cluster_v3.Cluster, inEndpoints []*envoy_config_endpoint_v3.ClusterLoadAssignment, inRouteConfigurations []*envoy_config_route_v3.RouteConfiguration, inListeners []*envoy_config_listener_v3.Listener) ([]*envoy_config_cluster_v3.Cluster, []*envoy_config_endpoint_v3.ClusterLoadAssignment, []*envoy_config_route_v3.RouteConfiguration, []*envoy_config_listener_v3.Listener, error) {
	var generatedClusters []*envoy_config_cluster_v3.Cluster
	for _, lCfg := range p.filterHashMap {
		generatedClusters = append(generatedClusters, generateCustomDynamicForwardProxyCluster(lCfg))
	}
	return generatedClusters, nil, nil, nil, nil
}

// envoy is silly and thus dynamic forward proxy DNS config must be identical across HTTP filter and cluster config,
// https://github.com/envoyproxy/envoy/blob/v1.21.1/source/extensions/filters/http/dynamic_forward_proxy/proxy_filter.cc#L129-L132
//
// to be nice, we hide this behavior from the user and generate a cluster for each DNS cache config as provided
// in our http listener options.
//
// as a result of this, the generated cluster is very simple (e.g., no TLS config). this is intentional as the provided
// use case did not require it, and I wanted to keep the number of dangerous user configurations to a minimum. we could
// add a new upstream type in the future for dynamic forwarding and make other cluster fields configurable, but this
// would require very careful validation with all other features and require the extra user step of providing an
// upstream that in most cases the user does not want to customize at all.
func generateCustomDynamicForwardProxyCluster(lCfg *dynamic_forward_proxy.FilterConfig) *envoy_config_cluster_v3.Cluster {
	cc := &envoy_extensions_clusters_dynamic_forward_proxy_v3.ClusterConfig{
		DnsCacheConfig: convertDnsCacheConfig(lCfg.GetDnsCacheConfig()),
		// AllowInsecureClusterOptions is not needed to be configurable unless we make a
		// new upstream type so the cluster's upstream_http_protocol_options is configurable
		AllowInsecureClusterOptions: false,
		AllowCoalescedConnections:   false, // not-implemented in envoy yet
	}
	return &envoy_config_cluster_v3.Cluster{
		Name:           GetGeneratedClusterName(lCfg),
		ConnectTimeout: &duration.Duration{Seconds: 5},
		LbPolicy:       envoy_config_cluster_v3.Cluster_CLUSTER_PROVIDED,
		ClusterDiscoveryType: &envoy_config_cluster_v3.Cluster_ClusterType{
			ClusterType: &envoy_config_cluster_v3.Cluster_CustomClusterType{
				Name:        "envoy.clusters.dynamic_forward_proxy",
				TypedConfig: utils.MustMessageToAny(cc),
			},
		},
	}
}

func GetGeneratedClusterName(dfpListenerConf *dynamic_forward_proxy.FilterConfig) string {
	// should be safe, cluster names can get up to 60 characters long https://github.com/envoyproxy/envoy/pull/667/files
	return fmt.Sprintf("solo_io_generated_dfp:%s", getHashString(dfpListenerConf))
}

func getGeneratedCacheName(cc *dynamic_forward_proxy.DnsCacheConfig) string {
	// should be safe, cluster names can get up to 60 characters long https://github.com/envoyproxy/envoy/pull/667/files
	return fmt.Sprintf("solo_io_generated_dfp:%s", fmt.Sprintf("%v", hashutils.MustHash(cc)))
}

func getHashString(dfpListenerConf *dynamic_forward_proxy.FilterConfig) string {
	dfpListenerConf.GetDnsCacheConfig()
	return fmt.Sprintf("%v", hashutils.MustHash(dfpListenerConf))
}

func convertDnsCacheConfig(dfpListenerConf *dynamic_forward_proxy.DnsCacheConfig) *envoy_extensions_common_dynamic_forward_proxy_v3.DnsCacheConfig {
	return &envoy_extensions_common_dynamic_forward_proxy_v3.DnsCacheConfig{
		Name:                   getGeneratedCacheName(dfpListenerConf), // silly envoy behavior, MUST match other caches with exact same DNS config
		DnsLookupFamily:        convertDnsLookupFamily(dfpListenerConf.GetDnsLookupFamily()),
		DnsRefreshRate:         dfpListenerConf.GetDnsRefreshRate(),
		HostTtl:                dfpListenerConf.GetHostTtl(),
		MaxHosts:               dfpListenerConf.GetMaxHosts(),
		DnsFailureRefreshRate:  convertFailureRefreshRate(dfpListenerConf.GetDnsFailureRefreshRate()),
		DnsCacheCircuitBreaker: convertDnsCacheCircuitBreaker(dfpListenerConf.GetDnsCacheCircuitBreaker()),
		UseTcpForDnsLookups:    false, // deprecated, do not use. prefer TypedDnsResolverConfig
		DnsResolutionConfig:    nil,   // deprecated, do not use. prefer TypedDnsResolverConfig
		TypedDnsResolverConfig: convertTypedDnsResolverConfig(dfpListenerConf),
		PreresolveHostnames:    convertPreresolveHostnames(dfpListenerConf.GetPreresolveHostnames()),
		DnsQueryTimeout:        dfpListenerConf.GetDnsQueryTimeout(),
		KeyValueConfig:         nil, // not-implemented in envoy yet
	}
}

func convertAddress(addrs []*v3.Address) []*envoy_config_core_v3.Address {
	if len(addrs) == 0 {
		return nil
	}

	var addresses []*envoy_config_core_v3.Address
	for _, a := range addrs {
		newAddr := &envoy_config_core_v3.Address{
			Address: nil, // filled in later
		}
		switch ps := a.GetAddress().(type) {
		case *v3.Address_SocketAddress:
			newAddr.Address = &envoy_config_core_v3.Address_SocketAddress{
				SocketAddress: convertSocketAddress(ps.SocketAddress),
			}
		case *v3.Address_Pipe:
			newAddr.Address = &envoy_config_core_v3.Address_Pipe{
				Pipe: convertPipe(ps.Pipe),
			}
		}
		addresses = append(addresses, newAddr)
	}
	return addresses
}

func convertPipe(pipe *v3.Pipe) *envoy_config_core_v3.Pipe {
	if pipe == nil {
		return nil
	}
	return &envoy_config_core_v3.Pipe{
		Path: pipe.GetPath(),
		Mode: pipe.GetMode(),
	}
}

func convertDnsResolverOptions(opts *v3.DnsResolverOptions) *envoy_config_core_v3.DnsResolverOptions {
	if opts == nil {
		return nil
	}
	return &envoy_config_core_v3.DnsResolverOptions{
		UseTcpForDnsLookups:   opts.GetUseTcpForDnsLookups(),
		NoDefaultSearchDomain: opts.GetNoDefaultSearchDomain(),
	}
}

func convertTypedDnsResolverConfig(dfpListenerConf *dynamic_forward_proxy.DnsCacheConfig) *envoy_config_core_v3.TypedExtensionConfig {
	if dfpListenerConf.GetDnsCacheType() == nil {
		return nil
	}
	var typedConf *envoy_config_core_v3.TypedExtensionConfig
	switch cacheConf := dfpListenerConf.GetDnsCacheType().(type) {
	case *dynamic_forward_proxy.DnsCacheConfig_CaresDns:
		c := &envoy_extensions_network_dns_resolver_cares_v3.CaresDnsResolverConfig{
			Resolvers:          convertAddress(cacheConf.CaresDns.GetResolvers()),
			DnsResolverOptions: convertDnsResolverOptions(cacheConf.CaresDns.GetDnsResolverOptions()),
		}
		typedConf = &envoy_config_core_v3.TypedExtensionConfig{
			Name:        "envoy.network.dns_resolver.cares",
			TypedConfig: utils.MustMessageToAny(c),
		}
	case *dynamic_forward_proxy.DnsCacheConfig_AppleDns:
		c := &envoy_extensions_network_dns_resolver_apple_v3.AppleDnsResolverConfig{}
		typedConf = &envoy_config_core_v3.TypedExtensionConfig{
			Name:        "envoy.network.dns_resolver.apple",
			TypedConfig: utils.MustMessageToAny(c),
		}
	}
	return typedConf
}

func convertDnsCacheCircuitBreaker(breakers *dynamic_forward_proxy.DnsCacheCircuitBreakers) *envoy_extensions_common_dynamic_forward_proxy_v3.DnsCacheCircuitBreakers {
	if breakers == nil {
		return nil
	}
	return &envoy_extensions_common_dynamic_forward_proxy_v3.DnsCacheCircuitBreakers{
		MaxPendingRequests: breakers.GetMaxPendingRequests(),
	}
}

func convertDnsLookupFamily(family dynamic_forward_proxy.DnsLookupFamily) envoy_config_cluster_v3.Cluster_DnsLookupFamily {
	switch family {
	case dynamic_forward_proxy.DnsLookupFamily_AUTO:
		return envoy_config_cluster_v3.Cluster_AUTO
	case dynamic_forward_proxy.DnsLookupFamily_V6_ONLY:
		return envoy_config_cluster_v3.Cluster_V6_ONLY
	case dynamic_forward_proxy.DnsLookupFamily_V4_ONLY:
		return envoy_config_cluster_v3.Cluster_V4_ONLY
	case dynamic_forward_proxy.DnsLookupFamily_V4_PREFERRED:
		return envoy_config_cluster_v3.Cluster_V4_PREFERRED
	case dynamic_forward_proxy.DnsLookupFamily_ALL:
		return envoy_config_cluster_v3.Cluster_ALL
	}
	return envoy_config_cluster_v3.Cluster_V4_ONLY
}

func convertPreresolveHostnames(sas []*v3.SocketAddress) []*envoy_config_core_v3.SocketAddress {
	if len(sas) == 0 {
		return nil
	}
	var addresses []*envoy_config_core_v3.SocketAddress
	for _, a := range sas {
		addresses = append(addresses, convertSocketAddress(a))
	}
	return addresses
}

func convertSocketAddress(a *v3.SocketAddress) *envoy_config_core_v3.SocketAddress {
	var protocol envoy_config_core_v3.SocketAddress_Protocol
	switch a.GetProtocol() {
	case v3.SocketAddress_TCP:
		protocol = envoy_config_core_v3.SocketAddress_TCP
	case v3.SocketAddress_UDP:
		protocol = envoy_config_core_v3.SocketAddress_UDP
	}
	newAddr := &envoy_config_core_v3.SocketAddress{
		Protocol:      protocol,
		Address:       a.GetAddress(),
		PortSpecifier: nil, // set-later
		ResolverName:  a.GetResolverName(),
		Ipv4Compat:    a.GetIpv4Compat(),
	}
	switch ps := a.GetPortSpecifier().(type) {
	case *v3.SocketAddress_PortValue:
		newAddr.PortSpecifier = &envoy_config_core_v3.SocketAddress_PortValue{
			PortValue: ps.PortValue,
		}
	case *v3.SocketAddress_NamedPort:
		newAddr.PortSpecifier = &envoy_config_core_v3.SocketAddress_NamedPort{
			NamedPort: ps.NamedPort,
		}
	}
	return newAddr
}

func convertFailureRefreshRate(rate *dynamic_forward_proxy.RefreshRate) *envoy_config_cluster_v3.Cluster_RefreshRate {
	if rate == nil {
		return nil
	}
	return &envoy_config_cluster_v3.Cluster_RefreshRate{
		BaseInterval: rate.GetBaseInterval(),
		MaxInterval:  rate.GetMaxInterval(),
	}
}

func (p *plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	cpDfp := listener.GetOptions().GetDynamicForwardProxy()
	if cpDfp == nil {
		// TODO(kdorosh) add default dns config
		//return []plugins.StagedHttpFilter{}, nil
	}

	dfp := &envoy_extensions_filters_http_dynamic_forward_proxy_v3.FilterConfig{
		DnsCacheConfig:      convertDnsCacheConfig(cpDfp.GetDnsCacheConfig()),
		SaveUpstreamAddress: cpDfp.GetSaveUpstreamAddress(),
	}

	p.filterHashMap[getHashString(cpDfp)] = cpDfp

	c, err := plugins.NewStagedFilterWithConfig(FilterName, dfp, pluginStage)
	if err != nil {
		return []plugins.StagedHttpFilter{}, err
	}
	return []plugins.StagedHttpFilter{c}, nil
}

func (p *plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoy_config_route_v3.Route) error {
	dfpCfg := in.GetRouteAction().GetDynamicForwardProxy()
	if dfpCfg == nil {
		return nil
	}
	dfpRouteCfg := &envoy_extensions_filters_http_dynamic_forward_proxy_v3.PerRouteConfig{}
	switch d := dfpCfg.GetHostRewriteSpecifier().(type) {
	case *dynamic_forward_proxy.PerRouteConfig_HostRewrite:
		dfpRouteCfg.HostRewriteSpecifier = &envoy_extensions_filters_http_dynamic_forward_proxy_v3.PerRouteConfig_HostRewriteLiteral{
			HostRewriteLiteral: d.HostRewrite,
		}
	case *dynamic_forward_proxy.PerRouteConfig_AutoHostRewriteHeader:
		dfpRouteCfg.HostRewriteSpecifier = &envoy_extensions_filters_http_dynamic_forward_proxy_v3.PerRouteConfig_HostRewriteHeader{
			HostRewriteHeader: d.AutoHostRewriteHeader,
		}
	default:
		return eris.Errorf("unimplemented dynamic forward proxy route config type %T", d)
	}
	return pluginutils.SetRoutePerFilterConfig(out, FilterName, dfpRouteCfg)
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return ExtensionName
}

func (p *plugin) Init(_ plugins.InitParams) error {
	p.filterHashMap = map[string]*dynamic_forward_proxy.FilterConfig{}
	return nil
}
