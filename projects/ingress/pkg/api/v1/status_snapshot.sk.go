// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type StatusSnapshot struct {
	Services  KubeServiceList
	Ingresses IngressList
}

func (s StatusSnapshot) Clone() StatusSnapshot {
	return StatusSnapshot{
		Services:  s.Services.Clone(),
		Ingresses: s.Ingresses.Clone(),
	}
}

func (s StatusSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashServices(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashIngresses(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s StatusSnapshot) hashServices(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Services.AsInterfaces()...)
}

func (s StatusSnapshot) hashIngresses(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ingresses.AsInterfaces()...)
}

func (s StatusSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
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

func (s *StatusSnapshot) GetResourcesList(resource resources.Resource) (resources.ResourceList, error) {
	switch resource.(type) {
	case *KubeService:
		return s.Services.AsResources(), nil
	case *Ingress:
		return s.Ingresses.AsResources(), nil
	default:
		return resources.ResourceList{}, eris.New("did not contain the input resource type returning empty list")
	}
}

func (s *StatusSnapshot) RemoveFromResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch resource.(type) {
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

func (s *StatusSnapshot) RemoveMatches(predicate core.Predicate) {
	var Services KubeServiceList
	for _, res := range s.Services {
		if matches := predicate(res.GetMetadata()); !matches {
			Services = append(Services, res)
		}
	}
	s.Services = Services
	var Ingresses IngressList
	for _, res := range s.Ingresses {
		if matches := predicate(res.GetMetadata()); !matches {
			Ingresses = append(Ingresses, res)
		}
	}
	s.Ingresses = Ingresses
}

func (s *StatusSnapshot) UpsertToResourceList(resource resources.Resource) error {
	refKey := resource.GetMetadata().Ref().Key()
	switch typed := resource.(type) {
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

type StatusSnapshotStringer struct {
	Version   uint64
	Services  []string
	Ingresses []string
}

func (ss StatusSnapshotStringer) String() string {
	s := fmt.Sprintf("StatusSnapshot %v\n", ss.Version)

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

func (s StatusSnapshot) Stringer() StatusSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return StatusSnapshotStringer{
		Version:   snapshotHash,
		Services:  s.Services.NamespacesDotNames(),
		Ingresses: s.Ingresses.NamespacesDotNames(),
	}
}

var StatusGvkToHashableResource = map[schema.GroupVersionKind]func() resources.HashableResource{
	KubeServiceGVK: NewKubeServiceHashableResource,
	IngressGVK:     NewIngressHashableResource,
}
