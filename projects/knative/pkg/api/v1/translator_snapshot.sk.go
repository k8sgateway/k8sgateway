// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	github_com_solo_io_gloo_projects_knative_pkg_api_external_knative "github.com/solo-io/gloo/projects/knative/pkg/api/external/knative"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type TranslatorSnapshot struct {
	Secrets   gloo_solo_io.SecretList
	Ingresses github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList
}

func (s TranslatorSnapshot) Clone() TranslatorSnapshot {
	return TranslatorSnapshot{
		Secrets:   s.Secrets.Clone(),
		Ingresses: s.Ingresses.Clone(),
	}
}

func (s TranslatorSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashSecrets(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashIngresses(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s TranslatorSnapshot) hashSecrets(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Secrets.AsInterfaces()...)
}

func (s TranslatorSnapshot) hashIngresses(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ingresses.AsInterfaces()...)
}

func (s TranslatorSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	SecretsHash, err := s.hashSecrets(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("secrets", SecretsHash))
	IngressesHash, err := s.hashIngresses(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("ingresses", IngressesHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

type TranslatorSnapshotStringer struct {
	Version   uint64
	Secrets   []string
	Ingresses []string
}

func (ss TranslatorSnapshotStringer) String() string {
	s := fmt.Sprintf("TranslatorSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Ingresses %v\n", len(ss.Ingresses))
	for _, name := range ss.Ingresses {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s TranslatorSnapshot) Stringer() TranslatorSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(errors.Wrapf(err, "error hashing, this should never happen"))
	}
	return TranslatorSnapshotStringer{
		Version:   snapshotHash,
		Secrets:   s.Secrets.NamespacesDotNames(),
		Ingresses: s.Ingresses.NamespacesDotNames(),
	}
}
