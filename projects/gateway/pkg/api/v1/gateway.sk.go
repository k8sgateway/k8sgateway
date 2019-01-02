// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewGateway(namespace, name string) *Gateway {
	return &Gateway{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Gateway) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Gateway) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Gateway) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.VirtualServices,
		r.BindAddress,
		r.BindPort,
		r.Plugins,
	)
}

type GatewayList []*Gateway
type GatewaysByNamespace map[string]GatewayList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list GatewayList) Find(namespace, name string) (*Gateway, error) {
	for _, gateway := range list {
		if gateway.Metadata.Name == name {
			if namespace == "" || gateway.Metadata.Namespace == namespace {
				return gateway, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find gateway %v.%v", namespace, name)
}

func (list GatewayList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, gateway := range list {
		ress = append(ress, gateway)
	}
	return ress
}

func (list GatewayList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, gateway := range list {
		ress = append(ress, gateway)
	}
	return ress
}

func (list GatewayList) Names() []string {
	var names []string
	for _, gateway := range list {
		names = append(names, gateway.Metadata.Name)
	}
	return names
}

func (list GatewayList) NamespacesDotNames() []string {
	var names []string
	for _, gateway := range list {
		names = append(names, gateway.Metadata.Namespace+"."+gateway.Metadata.Name)
	}
	return names
}

func (list GatewayList) Sort() GatewayList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list GatewayList) Clone() GatewayList {
	var gatewayList GatewayList
	for _, gateway := range list {
		gatewayList = append(gatewayList, proto.Clone(gateway).(*Gateway))
	}
	return gatewayList
}

func (list GatewayList) Each(f func(element *Gateway)) {
	for _, gateway := range list {
		f(gateway)
	}
}

func (list GatewayList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Gateway) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (list GatewayList) ByNamespace() GatewaysByNamespace {
	byNamespace := make(GatewaysByNamespace)
	for _, gateway := range list {
		byNamespace.Add(gateway)
	}
	return byNamespace
}

func (byNamespace GatewaysByNamespace) Add(gateway ...*Gateway) {
	for _, item := range gateway {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace GatewaysByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace GatewaysByNamespace) List() GatewayList {
	var list GatewayList
	for _, gatewayList := range byNamespace {
		list = append(list, gatewayList...)
	}
	return list.Sort()
}

func (byNamespace GatewaysByNamespace) Clone() GatewaysByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Gateway{}

// Kubernetes Adapter for Gateway

func (o *Gateway) GetObjectKind() schema.ObjectKind {
	t := GatewayCrd.TypeMeta()
	return &t
}

func (o *Gateway) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Gateway)
}

var GatewayCrd = crd.NewCrd("gateway.solo.io",
	"gateways",
	"gateway.solo.io",
	"v1",
	"Gateway",
	"gw",
	&Gateway{})
