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

func NewIngress(namespace, name string) *Ingress {
	ingress := &Ingress{}
	ingress.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return ingress
}

func (r *Ingress) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Ingress) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.KubeIngressSpec,
	)
}

func (r *Ingress) GroupVersionKind() schema.GroupVersionKind {
	return IngressGVK
}

type IngressList []*Ingress

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list IngressList) Find(namespace, name string) (*Ingress, error) {
	for _, ingress := range list {
		if ingress.GetMetadata().Name == name {
			if namespace == "" || ingress.GetMetadata().Namespace == namespace {
				return ingress, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find ingress %v.%v", namespace, name)
}

func (list IngressList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, ingress := range list {
		ress = append(ress, ingress)
	}
	return ress
}

func (list IngressList) Names() []string {
	var names []string
	for _, ingress := range list {
		names = append(names, ingress.GetMetadata().Name)
	}
	return names
}

func (list IngressList) NamespacesDotNames() []string {
	var names []string
	for _, ingress := range list {
		names = append(names, ingress.GetMetadata().Namespace+"."+ingress.GetMetadata().Name)
	}
	return names
}

func (list IngressList) Sort() IngressList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list IngressList) Clone() IngressList {
	var ingressList IngressList
	for _, ingress := range list {
		ingressList = append(ingressList, resources.Clone(ingress).(*Ingress))
	}
	return ingressList
}

func (list IngressList) Each(f func(element *Ingress)) {
	for _, ingress := range list {
		f(ingress)
	}
}

func (list IngressList) EachResource(f func(element resources.Resource)) {
	for _, ingress := range list {
		f(ingress)
	}
}

func (list IngressList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Ingress) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for Ingress

func (o *Ingress) GetObjectKind() schema.ObjectKind {
	t := IngressCrd.TypeMeta()
	return &t
}

func (o *Ingress) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Ingress)
}

var (
	IngressCrd = crd.NewCrd(
		"ingresses",
		IngressGVK.Group,
		IngressGVK.Version,
		IngressGVK.Kind,
		"ig",
		false,
		&Ingress{})
)

func init() {
	if err := crd.AddCrd(IngressCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	IngressGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "ingress.solo.io",
		Kind:    "Ingress",
	}
)
