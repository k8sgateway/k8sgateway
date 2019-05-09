// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Artifacts      ArtifactList
	Endpoints      EndpointList
	Proxies        ProxyList
	Upstreamgroups UpstreamGroupList
	Secrets        SecretList
	Upstreams      UpstreamList
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Artifacts:      s.Artifacts.Clone(),
		Endpoints:      s.Endpoints.Clone(),
		Proxies:        s.Proxies.Clone(),
		Upstreamgroups: s.Upstreamgroups.Clone(),
		Secrets:        s.Secrets.Clone(),
		Upstreams:      s.Upstreams.Clone(),
	}
}

func (s ApiSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashArtifacts(),
		s.hashEndpoints(),
		s.hashProxies(),
		s.hashUpstreamgroups(),
		s.hashSecrets(),
		s.hashUpstreams(),
	)
}

func (s ApiSnapshot) hashArtifacts() uint64 {
	return hashutils.HashAll(s.Artifacts.AsInterfaces()...)
}

func (s ApiSnapshot) hashEndpoints() uint64 {
	return hashutils.HashAll(s.Endpoints.AsInterfaces()...)
}

func (s ApiSnapshot) hashProxies() uint64 {
	return hashutils.HashAll(s.Proxies.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreamgroups() uint64 {
	return hashutils.HashAll(s.Upstreamgroups.AsInterfaces()...)
}

func (s ApiSnapshot) hashSecrets() uint64 {
	return hashutils.HashAll(s.Secrets.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("artifacts", s.hashArtifacts()))
	fields = append(fields, zap.Uint64("endpoints", s.hashEndpoints()))
	fields = append(fields, zap.Uint64("proxies", s.hashProxies()))
	fields = append(fields, zap.Uint64("upstreamgroups", s.hashUpstreamgroups()))
	fields = append(fields, zap.Uint64("secrets", s.hashSecrets()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type ApiSnapshotStringer struct {
	Version        uint64
	Artifacts      []string
	Endpoints      []string
	Proxies        []string
	Upstreamgroups []string
	Secrets        []string
	Upstreams      []string
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

	s += fmt.Sprintf("  Upstreamgroups %v\n", len(ss.Upstreamgroups))
	for _, name := range ss.Upstreamgroups {
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

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	return ApiSnapshotStringer{
		Version:        s.Hash(),
		Artifacts:      s.Artifacts.NamespacesDotNames(),
		Endpoints:      s.Endpoints.NamespacesDotNames(),
		Proxies:        s.Proxies.NamespacesDotNames(),
		Upstreamgroups: s.Upstreamgroups.NamespacesDotNames(),
		Secrets:        s.Secrets.NamespacesDotNames(),
		Upstreams:      s.Upstreams.NamespacesDotNames(),
	}
}
