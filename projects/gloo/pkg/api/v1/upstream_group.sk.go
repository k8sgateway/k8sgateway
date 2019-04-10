// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewUpstreamGroup(namespace, name string) *UpstreamGroup {
	upstreamgroup := &UpstreamGroup{}
	upstreamgroup.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return upstreamgroup
}

func (r *UpstreamGroup) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *UpstreamGroup) SetStatus(status core.Status) {
	r.Status = status
}

func (r *UpstreamGroup) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Destinations,
	)
}

type UpstreamGroupList []*UpstreamGroup
type UpstreamgroupsByNamespace map[string]UpstreamGroupList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list UpstreamGroupList) Find(namespace, name string) (*UpstreamGroup, error) {
	for _, upstreamGroup := range list {
		if upstreamGroup.GetMetadata().Name == name {
			if namespace == "" || upstreamGroup.GetMetadata().Namespace == namespace {
				return upstreamGroup, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find upstreamGroup %v.%v", namespace, name)
}

func (list UpstreamGroupList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, upstreamGroup := range list {
		ress = append(ress, upstreamGroup)
	}
	return ress
}

func (list UpstreamGroupList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, upstreamGroup := range list {
		ress = append(ress, upstreamGroup)
	}
	return ress
}

func (list UpstreamGroupList) Names() []string {
	var names []string
	for _, upstreamGroup := range list {
		names = append(names, upstreamGroup.GetMetadata().Name)
	}
	return names
}

func (list UpstreamGroupList) NamespacesDotNames() []string {
	var names []string
	for _, upstreamGroup := range list {
		names = append(names, upstreamGroup.GetMetadata().Namespace+"."+upstreamGroup.GetMetadata().Name)
	}
	return names
}

func (list UpstreamGroupList) Sort() UpstreamGroupList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list UpstreamGroupList) Clone() UpstreamGroupList {
	var upstreamGroupList UpstreamGroupList
	for _, upstreamGroup := range list {
		upstreamGroupList = append(upstreamGroupList, resources.Clone(upstreamGroup).(*UpstreamGroup))
	}
	return upstreamGroupList
}

func (list UpstreamGroupList) Each(f func(element *UpstreamGroup)) {
	for _, upstreamGroup := range list {
		f(upstreamGroup)
	}
}

func (list UpstreamGroupList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *UpstreamGroup) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace UpstreamgroupsByNamespace) Add(upstreamGroup ...*UpstreamGroup) {
	for _, item := range upstreamGroup {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace UpstreamgroupsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace UpstreamgroupsByNamespace) List() UpstreamGroupList {
	var list UpstreamGroupList
	for _, upstreamGroupList := range byNamespace {
		list = append(list, upstreamGroupList...)
	}
	return list.Sort()
}

func (byNamespace UpstreamgroupsByNamespace) Clone() UpstreamgroupsByNamespace {
	cloned := make(UpstreamgroupsByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &UpstreamGroup{}

// Kubernetes Adapter for UpstreamGroup

func (o *UpstreamGroup) GetObjectKind() schema.ObjectKind {
	t := UpstreamGroupCrd.TypeMeta()
	return &t
}

func (o *UpstreamGroup) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*UpstreamGroup)
}

var UpstreamGroupCrd = crd.NewCrd("gloo.solo.io",
	"upstreamgroups",
	"gloo.solo.io",
	"v1",
	"UpstreamGroup",
	"ug",
	false,
	&UpstreamGroup{})
