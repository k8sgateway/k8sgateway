// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Upstreams      UpstreamList
	Kubenamespaces github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceList
	Secrets        SecretList
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Upstreams:      s.Upstreams.Clone(),
		Kubenamespaces: s.Kubenamespaces.Clone(),
		Secrets:        s.Secrets.Clone(),
	}
}

func (s DiscoverySnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashKubenamespaces(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashSecrets(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s DiscoverySnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashKubenamespaces(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Kubenamespaces.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashSecrets(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Secrets.AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	KubenamespacesHash, err := s.hashKubenamespaces(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("kubenamespaces", KubenamespacesHash))
	SecretsHash, err := s.hashSecrets(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("secrets", SecretsHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

type DiscoverySnapshotStringer struct {
	Version        uint64
	Upstreams      []string
	Kubenamespaces []string
	Secrets        []string
}

func (ss DiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("DiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Kubenamespaces %v\n", len(ss.Kubenamespaces))
	for _, name := range ss.Kubenamespaces {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s DiscoverySnapshot) Stringer() DiscoverySnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	return DiscoverySnapshotStringer{
		Version:        snapshotHash,
		Upstreams:      s.Upstreams.NamespacesDotNames(),
		Kubenamespaces: s.Kubenamespaces.Names(),
		Secrets:        s.Secrets.NamespacesDotNames(),
	}
}
