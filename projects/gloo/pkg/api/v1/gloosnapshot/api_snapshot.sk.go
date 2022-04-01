// Code generated by solo-kit. DO NOT EDIT.

package gloosnapshot

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	gateway_solo_io "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit "github.com/solo-io/gloo/projects/gloo/pkg/api/external/solo/ratelimit"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	enterprise_gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	graphql_gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1beta1"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Artifacts          gloo_solo_io.ArtifactList
	Endpoints          gloo_solo_io.EndpointList
	Proxies            gloo_solo_io.ProxyList
	UpstreamGroups     gloo_solo_io.UpstreamGroupList
	Secrets            gloo_solo_io.SecretList
	Upstreams          gloo_solo_io.UpstreamList
	AuthConfigs        enterprise_gloo_solo_io.AuthConfigList
	Ratelimitconfigs   github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigList
	VirtualServices    gateway_solo_io.VirtualServiceList
	RouteTables        gateway_solo_io.RouteTableList
	Gateways           gateway_solo_io.GatewayList
	VirtualHostOptions gateway_solo_io.VirtualHostOptionList
	RouteOptions       gateway_solo_io.RouteOptionList
	GraphqlApis        graphql_gloo_solo_io.GraphQLApiList
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Artifacts:          s.Artifacts.Clone(),
		Endpoints:          s.Endpoints.Clone(),
		Proxies:            s.Proxies.Clone(),
		UpstreamGroups:     s.UpstreamGroups.Clone(),
		Secrets:            s.Secrets.Clone(),
		Upstreams:          s.Upstreams.Clone(),
		AuthConfigs:        s.AuthConfigs.Clone(),
		Ratelimitconfigs:   s.Ratelimitconfigs.Clone(),
		VirtualServices:    s.VirtualServices.Clone(),
		RouteTables:        s.RouteTables.Clone(),
		Gateways:           s.Gateways.Clone(),
		VirtualHostOptions: s.VirtualHostOptions.Clone(),
		RouteOptions:       s.RouteOptions.Clone(),
		GraphqlApis:        s.GraphqlApis.Clone(),
	}
}

func (s ApiSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashArtifacts(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashEndpoints(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashProxies(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashUpstreamGroups(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashSecrets(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashAuthConfigs(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRatelimitconfigs(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashVirtualServices(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRouteTables(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashGateways(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashVirtualHostOptions(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRouteOptions(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashGraphqlApis(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s ApiSnapshot) hashArtifacts(hasher hash.Hash64) (uint64, error) {
	clonedList := s.Artifacts.Clone()
	for _, v := range clonedList {
		v.Metadata.Annotations = nil
	}
	return hashutils.HashAllSafe(hasher, clonedList.AsInterfaces()...)
}

func (s ApiSnapshot) hashEndpoints(hasher hash.Hash64) (uint64, error) {
	clonedList := s.Endpoints.Clone()
	for _, v := range clonedList {
		v.Metadata.Annotations = nil
	}
	return hashutils.HashAllSafe(hasher, clonedList.AsInterfaces()...)
}

func (s ApiSnapshot) hashProxies(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Proxies.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreamGroups(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.UpstreamGroups.AsInterfaces()...)
}

func (s ApiSnapshot) hashSecrets(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Secrets.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s ApiSnapshot) hashAuthConfigs(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.AuthConfigs.AsInterfaces()...)
}

func (s ApiSnapshot) hashRatelimitconfigs(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ratelimitconfigs.AsInterfaces()...)
}

func (s ApiSnapshot) hashVirtualServices(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.VirtualServices.AsInterfaces()...)
}

func (s ApiSnapshot) hashRouteTables(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.RouteTables.AsInterfaces()...)
}

func (s ApiSnapshot) hashGateways(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Gateways.AsInterfaces()...)
}

func (s ApiSnapshot) hashVirtualHostOptions(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.VirtualHostOptions.AsInterfaces()...)
}

func (s ApiSnapshot) hashRouteOptions(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.RouteOptions.AsInterfaces()...)
}

func (s ApiSnapshot) hashGraphqlApis(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.GraphqlApis.AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	ArtifactsHash, err := s.hashArtifacts(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("artifacts", ArtifactsHash))
	EndpointsHash, err := s.hashEndpoints(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("endpoints", EndpointsHash))
	ProxiesHash, err := s.hashProxies(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("proxies", ProxiesHash))
	UpstreamGroupsHash, err := s.hashUpstreamGroups(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreamGroups", UpstreamGroupsHash))
	SecretsHash, err := s.hashSecrets(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("secrets", SecretsHash))
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	AuthConfigsHash, err := s.hashAuthConfigs(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("authConfigs", AuthConfigsHash))
	RatelimitconfigsHash, err := s.hashRatelimitconfigs(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("ratelimitconfigs", RatelimitconfigsHash))
	VirtualServicesHash, err := s.hashVirtualServices(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("virtualServices", VirtualServicesHash))
	RouteTablesHash, err := s.hashRouteTables(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("routeTables", RouteTablesHash))
	GatewaysHash, err := s.hashGateways(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("gateways", GatewaysHash))
	VirtualHostOptionsHash, err := s.hashVirtualHostOptions(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("virtualHostOptions", VirtualHostOptionsHash))
	RouteOptionsHash, err := s.hashRouteOptions(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("routeOptions", RouteOptionsHash))
	GraphqlApisHash, err := s.hashGraphqlApis(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("graphqlApis", GraphqlApisHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

type ApiSnapshotStringer struct {
	Version            uint64
	Artifacts          []string
	Endpoints          []string
	Proxies            []string
	UpstreamGroups     []string
	Secrets            []string
	Upstreams          []string
	AuthConfigs        []string
	Ratelimitconfigs   []string
	VirtualServices    []string
	RouteTables        []string
	Gateways           []string
	VirtualHostOptions []string
	RouteOptions       []string
	GraphqlApis        []string
}

func (ss ApiSnapshotStringer) String() string {
	s := fmt.Sprintf("ApiSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Artifacts %v\n", len(ss.Artifacts))
	for _, name := range ss.Artifacts {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Endpoints %v\n", len(ss.Endpoints))
	for _, name := range ss.Endpoints {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Proxies %v\n", len(ss.Proxies))
	for _, name := range ss.Proxies {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  UpstreamGroups %v\n", len(ss.UpstreamGroups))
	for _, name := range ss.UpstreamGroups {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  AuthConfigs %v\n", len(ss.AuthConfigs))
	for _, name := range ss.AuthConfigs {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Ratelimitconfigs %v\n", len(ss.Ratelimitconfigs))
	for _, name := range ss.Ratelimitconfigs {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  VirtualServices %v\n", len(ss.VirtualServices))
	for _, name := range ss.VirtualServices {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  RouteTables %v\n", len(ss.RouteTables))
	for _, name := range ss.RouteTables {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Gateways %v\n", len(ss.Gateways))
	for _, name := range ss.Gateways {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  VirtualHostOptions %v\n", len(ss.VirtualHostOptions))
	for _, name := range ss.VirtualHostOptions {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  RouteOptions %v\n", len(ss.RouteOptions))
	for _, name := range ss.RouteOptions {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  GraphqlApis %v\n", len(ss.GraphqlApis))
	for _, name := range ss.GraphqlApis {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return ApiSnapshotStringer{
		Version:            snapshotHash,
		Artifacts:          s.Artifacts.NamespacesDotNames(),
		Endpoints:          s.Endpoints.NamespacesDotNames(),
		Proxies:            s.Proxies.NamespacesDotNames(),
		UpstreamGroups:     s.UpstreamGroups.NamespacesDotNames(),
		Secrets:            s.Secrets.NamespacesDotNames(),
		Upstreams:          s.Upstreams.NamespacesDotNames(),
		AuthConfigs:        s.AuthConfigs.NamespacesDotNames(),
		Ratelimitconfigs:   s.Ratelimitconfigs.NamespacesDotNames(),
		VirtualServices:    s.VirtualServices.NamespacesDotNames(),
		RouteTables:        s.RouteTables.NamespacesDotNames(),
		Gateways:           s.Gateways.NamespacesDotNames(),
		VirtualHostOptions: s.VirtualHostOptions.NamespacesDotNames(),
		RouteOptions:       s.RouteOptions.NamespacesDotNames(),
		GraphqlApis:        s.GraphqlApis.NamespacesDotNames(),
	}
}
