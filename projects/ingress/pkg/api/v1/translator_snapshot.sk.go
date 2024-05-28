// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ json.Marshaler = new(TranslatorSnapshot)

type TranslatorSnapshot struct {
	Upstreams gloo_solo_io.UpstreamList `json:"upstreams"`
	Services  KubeServiceList           `json:"services"`
	Ingresses IngressList               `json:"ingresses"`
}

func (s TranslatorSnapshot) Clone() TranslatorSnapshot {
	return TranslatorSnapshot{
		Upstreams: s.Upstreams.Clone(),
		Services:  s.Services.Clone(),
		Ingresses: s.Ingresses.Clone(),
	}
}

func (s TranslatorSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashServices(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashIngresses(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s TranslatorSnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s TranslatorSnapshot) hashServices(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Services.AsInterfaces()...)
}

func (s TranslatorSnapshot) hashIngresses(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ingresses.AsInterfaces()...)
}

func (s TranslatorSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	ServicesHash, err := s.hashServices(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("services", ServicesHash))
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

func (s TranslatorSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(&s)

}

func (s *TranslatorSnapshot) GetResourcesList(resource resources.Resource) (resources.ResourceList, error) {
	switch resource.(type) {
	case *gloo_solo_io.Upstream:
		return s.Upstreams.AsResources(), nil
	case *KubeService:
		return s.Services.AsResources(), nil
	case *Ingress:
		return s.Ingresses.AsResources(), nil
	default:
		return resources.ResourceList{}, eris.New("did not contain the input resource type returning empty list")
	}
}

func (s *TranslatorSnapshot) RemoveFromResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch resource.(type) {
	case *gloo_solo_io.Upstream:

		for i, res := range s.Upstreams {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Upstreams = append(s.Upstreams[:i], s.Upstreams[i+1:]...)
				break
			}
		}
		return nil
	case *KubeService:

		for i, res := range s.Services {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Services = append(s.Services[:i], s.Services[i+1:]...)
				break
			}
		}
		return nil
	case *Ingress:

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

func (s *TranslatorSnapshot) UpsertToResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch typed := resource.(type) {
	case *gloo_solo_io.Upstream:
		updated := false
		for i, res := range s.Upstreams {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Upstreams[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Upstreams = append(s.Upstreams, typed)
		}
		s.Upstreams.Sort()
		return nil
	case *KubeService:
		updated := false
		for i, res := range s.Services {
			if refKey == res.GetMetadata().Ref().Key() {
				s.Services[i] = typed
				updated = true
			}
		}
		if !updated {
			s.Services = append(s.Services, typed)
		}
		s.Services.Sort()
		return nil
	case *Ingress:
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
	Upstreams []string
	Services  []string
	Ingresses []string
}

func (ss TranslatorSnapshotStringer) String() string {
	s := fmt.Sprintf("TranslatorSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Services %v\n", len(ss.Services))
	for _, name := range ss.Services {
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
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return TranslatorSnapshotStringer{
		Version:   snapshotHash,
		Upstreams: s.Upstreams.NamespacesDotNames(),
		Services:  s.Services.NamespacesDotNames(),
		Ingresses: s.Ingresses.NamespacesDotNames(),
	}
}

var TranslatorGvkToHashableResource = map[schema.GroupVersionKind]func() resources.HashableResource{
	gloo_solo_io.UpstreamGVK: gloo_solo_io.NewUpstreamHashableResource,
	KubeServiceGVK:           NewKubeServiceHashableResource,
	IngressGVK:               NewIngressHashableResource,
}
