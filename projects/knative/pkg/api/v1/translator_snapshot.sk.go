// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	github_com_solo_io_gloo_projects_knative_pkg_api_external_knative "github.com/solo-io/gloo/projects/knative/pkg/api/external/knative"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type TranslatorSnapshot struct {
	Ingresses github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressList
}

func (s TranslatorSnapshot) Clone() TranslatorSnapshot {
	return TranslatorSnapshot{
		Ingresses: s.Ingresses.Clone(),
	}
}

func (s TranslatorSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashIngresses(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s TranslatorSnapshot) hashIngresses(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ingresses.AsInterfaces()...)
}

func (s TranslatorSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	IngressesHash, err := s.hashIngresses(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("ingresses", IngressesHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

func (s *TranslatorSnapshot) GetResourcesList(resource resources.Resource) (resources.ResourceList, error) {
	switch resource.(type) {
	case *github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.Ingress:
		return s.Ingresses.AsResources(), nil
	default:
		return resources.ResourceList{}, eris.New("did not contain the input resource type returning empty list")
	}
}

func (s *TranslatorSnapshot) RemoveFromResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch resource.(type) {
	case *github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.Ingress:

		for i, res := range s.Ingresses {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Ingresses = append(s.Ingresses[:i], s.Ingresses[i+1:]...)
				break
			}
		}
		return nil
	default:
		return eris.Errorf("did not remove the resource because its type does not exist [%T]", resource)
	}
}

func (s *TranslatorSnapshot) RemoveAllResourcesInNamespace(namespace string) error {

	for i, res := range s.Ingresses {
		if namespace == res.GetMetadata().GetNamespace() {
			s.Ingresses = append(s.Ingresses[:i], s.Ingresses[i+1:]...)
			break
		}
	}
	return nil
}

func (s *TranslatorSnapshot) UpsertToResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch typed := resource.(type) {
	case *github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.Ingress:
		updated := false
		for i, res := range s.Ingresses {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Ingresses[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Ingresses = append(s.Ingresses, typed)
		}
		s.Ingresses.Sort()
		return nil
	default:
		return eris.Errorf("did not add/replace the resource type because it does not exist %T", resource)
	}
}

type TranslatorSnapshotStringer struct {
	Version   uint64
	Ingresses []string
}

func (ss TranslatorSnapshotStringer) String() string {
	s := fmt.Sprintf("TranslatorSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Ingresses %v\n", len(ss.Ingresses))
	for _, name := range ss.Ingresses {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s TranslatorSnapshot) Stringer() TranslatorSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return TranslatorSnapshotStringer{
		Version:   snapshotHash,
		Ingresses: s.Ingresses.NamespacesDotNames(),
	}
}

var TranslatorGvkToHashableResource = map[schema.GroupVersionKind]func() resources.HashableResource{
	github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.IngressGVK: github_com_solo_io_gloo_projects_knative_pkg_api_external_knative.NewIngressHashableResource,
}
