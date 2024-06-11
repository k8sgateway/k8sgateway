// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Compile-time assertion
	_ resources.Resource = new(HttpListenerOption)
)

func NewHttpListenerOptionHashableResource() resources.HashableResource {
	return new(HttpListenerOption)
}

func NewHttpListenerOption(namespace, name string) *HttpListenerOption {
	httplisteneroption := &HttpListenerOption{}
	httplisteneroption.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return httplisteneroption
}

func (r *HttpListenerOption) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

func (r *HttpListenerOption) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *HttpListenerOption) GroupVersionKind() schema.GroupVersionKind {
	return HttpListenerOptionGVK
}

type HttpListenerOptionList []*HttpListenerOption

func (list HttpListenerOptionList) Find(namespace, name string) (*HttpListenerOption, error) {
	for _, httpListenerOption := range list {
		if httpListenerOption.GetMetadata().Name == name && httpListenerOption.GetMetadata().Namespace == namespace {
			return httpListenerOption, nil
		}
	}
	return nil, errors.Errorf("list did not find httpListenerOption %v.%v", namespace, name)
}

func (list HttpListenerOptionList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, httpListenerOption := range list {
		ress = append(ress, httpListenerOption)
	}
	return ress
}

func (list HttpListenerOptionList) Names() []string {
	var names []string
	for _, httpListenerOption := range list {
		names = append(names, httpListenerOption.GetMetadata().Name)
	}
	return names
}

func (list HttpListenerOptionList) NamespacesDotNames() []string {
	var names []string
	for _, httpListenerOption := range list {
		names = append(names, httpListenerOption.GetMetadata().Namespace+"."+httpListenerOption.GetMetadata().Name)
	}
	return names
}

func (list HttpListenerOptionList) Sort() HttpListenerOptionList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list HttpListenerOptionList) Clone() HttpListenerOptionList {
	var httpListenerOptionList HttpListenerOptionList
	for _, httpListenerOption := range list {
		httpListenerOptionList = append(httpListenerOptionList, resources.Clone(httpListenerOption).(*HttpListenerOption))
	}
	return httpListenerOptionList
}

func (list HttpListenerOptionList) Each(f func(element *HttpListenerOption)) {
	for _, httpListenerOption := range list {
		f(httpListenerOption)
	}
}

func (list HttpListenerOptionList) EachResource(f func(element resources.Resource)) {
	for _, httpListenerOption := range list {
		f(httpListenerOption)
	}
}

func (list HttpListenerOptionList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *HttpListenerOption) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for HttpListenerOption

func (o *HttpListenerOption) GetObjectKind() schema.ObjectKind {
	t := HttpListenerOptionCrd.TypeMeta()
	return &t
}

func (o *HttpListenerOption) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*HttpListenerOption)
}

func (o *HttpListenerOption) DeepCopyInto(out *HttpListenerOption) {
	clone := resources.Clone(o).(*HttpListenerOption)
	*out = *clone
}

var (
	HttpListenerOptionCrd = crd.NewCrd(
		"httplisteneroptions",
		HttpListenerOptionGVK.Group,
		HttpListenerOptionGVK.Version,
		HttpListenerOptionGVK.Kind,
		"hlisopts",
		false,
		&HttpListenerOption{})
)

var (
	HttpListenerOptionGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "gateway.solo.io",
		Kind:    "HttpListenerOption",
	}
)