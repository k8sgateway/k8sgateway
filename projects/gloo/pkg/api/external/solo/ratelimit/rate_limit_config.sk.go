// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
	"log"
	"sort"

	github_com_solo_io_gloo_projects_gloo_api_external_solo_ratelimit "github.com/solo-io/gloo/projects/gloo/api/external/solo/ratelimit"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewRateLimitConfig(namespace, name string) *RateLimitConfig {
	ratelimitconfig := &RateLimitConfig{}
	ratelimitconfig.RateLimitConfig.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return ratelimitconfig
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneableRateLimitConfig interface {
	resources.Resource
	Clone() *github_com_solo_io_gloo_projects_gloo_api_external_solo_ratelimit.RateLimitConfig
}

var _ CloneableRateLimitConfig = &github_com_solo_io_gloo_projects_gloo_api_external_solo_ratelimit.RateLimitConfig{}

type RateLimitConfig struct {
	github_com_solo_io_gloo_projects_gloo_api_external_solo_ratelimit.RateLimitConfig
}

func (r *RateLimitConfig) Clone() resources.Resource {
	return &RateLimitConfig{RateLimitConfig: *r.RateLimitConfig.Clone()}
}

func (r *RateLimitConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	clone := r.RateLimitConfig.Clone()
	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})
	err := binary.Write(hasher, binary.LittleEndian, hashutils.HashAll(clone))
	if err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (r *RateLimitConfig) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *RateLimitConfig) GroupVersionKind() schema.GroupVersionKind {
	return RateLimitConfigGVK
}

type RateLimitConfigList []*RateLimitConfig

func (list RateLimitConfigList) Find(namespace, name string) (*RateLimitConfig, error) {
	for _, rateLimitConfig := range list {
		if rateLimitConfig.GetMetadata().Name == name && rateLimitConfig.GetMetadata().Namespace == namespace {
			return rateLimitConfig, nil
		}
	}
	return nil, errors.Errorf("list did not find rateLimitConfig %v.%v", namespace, name)
}

func (list RateLimitConfigList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, rateLimitConfig := range list {
		ress = append(ress, rateLimitConfig)
	}
	return ress
}

func (list RateLimitConfigList) Names() []string {
	var names []string
	for _, rateLimitConfig := range list {
		names = append(names, rateLimitConfig.GetMetadata().Name)
	}
	return names
}

func (list RateLimitConfigList) NamespacesDotNames() []string {
	var names []string
	for _, rateLimitConfig := range list {
		names = append(names, rateLimitConfig.GetMetadata().Namespace+"."+rateLimitConfig.GetMetadata().Name)
	}
	return names
}

func (list RateLimitConfigList) Sort() RateLimitConfigList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list RateLimitConfigList) Clone() RateLimitConfigList {
	var rateLimitConfigList RateLimitConfigList
	for _, rateLimitConfig := range list {
		rateLimitConfigList = append(rateLimitConfigList, resources.Clone(rateLimitConfig).(*RateLimitConfig))
	}
	return rateLimitConfigList
}

func (list RateLimitConfigList) Each(f func(element *RateLimitConfig)) {
	for _, rateLimitConfig := range list {
		f(rateLimitConfig)
	}
}

func (list RateLimitConfigList) EachResource(f func(element resources.Resource)) {
	for _, rateLimitConfig := range list {
		f(rateLimitConfig)
	}
}

func (list RateLimitConfigList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *RateLimitConfig) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var (
	RateLimitConfigGVK = schema.GroupVersionKind{
		Version: "v1beta1",
		Group:   "ratelimit.solo.io",
		Kind:    "RateLimitConfig",
	}
)
