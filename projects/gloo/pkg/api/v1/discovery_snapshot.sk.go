// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Upstreams UpstreamsByNamespace
	Secrets   SecretsByNamespace
	Proxies   ProxiesByNamespace
	Artifacts ArtifactsByNamespace
	Endpoints EndpointsByNamespace
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Upstreams: s.Upstreams.Clone(),
		Secrets:   s.Secrets.Clone(),
		Proxies:   s.Proxies.Clone(),
		Artifacts: s.Artifacts.Clone(),
		Endpoints: s.Endpoints.Clone(),
	}
}

func (s DiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashUpstreams(),
		s.hashSecrets(),
		s.hashProxies(),
		s.hashArtifacts(),
		s.hashEndpoints(),
	)
}

func (s DiscoverySnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) hashSecrets() uint64 {
	return hashutils.HashAll(s.Secrets.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) hashProxies() uint64 {
	return hashutils.HashAll(s.Proxies.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) hashArtifacts() uint64 {
	return hashutils.HashAll(s.Artifacts.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) hashEndpoints() uint64 {
	return hashutils.HashAll(s.Endpoints.List().AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))
	fields = append(fields, zap.Uint64("secrets", s.hashSecrets()))
	fields = append(fields, zap.Uint64("proxies", s.hashProxies()))
	fields = append(fields, zap.Uint64("artifacts", s.hashArtifacts()))
	fields = append(fields, zap.Uint64("endpoints", s.hashEndpoints()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type DiscoverySnapshotStringer struct {
	Version   uint64
	Upstreams []string
	Secrets   []string
	Proxies   []string
	Artifacts []string
	Endpoints []string
}

func (ss DiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("DiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Proxies %v\n", len(ss.Proxies))
	for _, name := range ss.Proxies {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Artifacts %v\n", len(ss.Artifacts))
	for _, name := range ss.Artifacts {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Endpoints %v\n", len(ss.Endpoints))
	for _, name := range ss.Endpoints {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s DiscoverySnapshot) Stringer() DiscoverySnapshotStringer {
	return DiscoverySnapshotStringer{
		Version:   s.Hash(),
		Upstreams: s.Upstreams.List().NamespacesDotNames(),
		Secrets:   s.Secrets.List().NamespacesDotNames(),
		Proxies:   s.Proxies.List().NamespacesDotNames(),
		Artifacts: s.Artifacts.List().NamespacesDotNames(),
		Endpoints: s.Endpoints.List().NamespacesDotNames(),
	}
}
