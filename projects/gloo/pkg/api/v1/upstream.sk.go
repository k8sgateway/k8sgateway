// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewUpstream(namespace, name string) *Upstream {
	upstream := &Upstream{}
	upstream.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return upstream
}

func (r *Upstream) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Upstream) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Upstream) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	metaCopy.Generation = 0
	// investigate zeroing out owner refs as well
	return hashutils.HashAll(
		metaCopy,
		r.UpstreamSpec,
		r.DiscoveryMetadata,
	)
}

func (r *Upstream) GroupVersionKind() schema.GroupVersionKind {
	return UpstreamGVK
}

type UpstreamList []*Upstream

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list UpstreamList) Find(namespace, name string) (*Upstream, error) {
	for _, upstream := range list {
		if upstream.GetMetadata().Name == name {
			if namespace == "" || upstream.GetMetadata().Namespace == namespace {
				return upstream, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find upstream %v.%v", namespace, name)
}

func (list UpstreamList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, upstream := range list {
		ress = append(ress, upstream)
	}
	return ress
}

func (list UpstreamList) Names() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.GetMetadata().Name)
	}
	return names
}

func (list UpstreamList) NamespacesDotNames() []string {
	var names []string
	for _, upstream := range list {
		names = append(names, upstream.GetMetadata().Namespace+"."+upstream.GetMetadata().Name)
	}
	return names
}

func (list UpstreamList) Sort() UpstreamList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list UpstreamList) Clone() UpstreamList {
	var upstreamList UpstreamList
	for _, upstream := range list {
		upstreamList = append(upstreamList, resources.Clone(upstream).(*Upstream))
	}
	return upstreamList
}

func (list UpstreamList) Each(f func(element *Upstream)) {
	for _, upstream := range list {
		f(upstream)
	}
}

func (list UpstreamList) EachResource(f func(element resources.Resource)) {
	for _, upstream := range list {
		f(upstream)
	}
}

func (list UpstreamList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Upstream) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Upstream

func (o *Upstream) GetObjectKind() schema.ObjectKind {
	t := UpstreamCrd.TypeMeta()
	return &t
}

func (o *Upstream) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Upstream)
}

func (o *Upstream) DeepCopyInto(out *Upstream) {
	clone := resources.Clone(o).(*Upstream)
	*out = *clone
}

var (
	UpstreamCrd = crd.NewCrd(
		"upstreams",
		UpstreamGVK.Group,
		UpstreamGVK.Version,
		UpstreamGVK.Kind,
		"us",
		false,
		&Upstream{})
)

func init() {
	if err := crd.AddCrd(UpstreamCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	UpstreamGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gloo.solo.io",
		Kind:    "Upstream",
	}
)
