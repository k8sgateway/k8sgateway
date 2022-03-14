// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewGraphQLApi(namespace, name string) *GraphQLApi {
	graphqlapi := &GraphQLApi{}
	graphqlapi.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return graphqlapi
}

func (r *GraphQLApi) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

// Deprecated
func (r *GraphQLApi) SetStatus(status *core.Status) {
	statusutils.SetSingleStatusInNamespacedStatuses(r, status)
}

// Deprecated
func (r *GraphQLApi) GetStatus() *core.Status {
	if r != nil {
		return statusutils.GetSingleStatusInNamespacedStatuses(r)
	}
	return nil
}

func (r *GraphQLApi) SetNamespacedStatuses(namespacedStatuses *core.NamespacedStatuses) {
	r.NamespacedStatuses = namespacedStatuses
}

func (r *GraphQLApi) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *GraphQLApi) GroupVersionKind() schema.GroupVersionKind {
	return GraphQLApiGVK
}

type GraphQLApiList []*GraphQLApi

func (list GraphQLApiList) Find(namespace, name string) (*GraphQLApi, error) {
	for _, graphQLApi := range list {
		if graphQLApi.GetMetadata().Name == name && graphQLApi.GetMetadata().Namespace == namespace {
			return graphQLApi, nil
		}
	}
	return nil, errors.Errorf("list did not find graphQLApi %v.%v", namespace, name)
}

func (list GraphQLApiList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, graphQLApi := range list {
		ress = append(ress, graphQLApi)
	}
	return ress
}

func (list GraphQLApiList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, graphQLApi := range list {
		ress = append(ress, graphQLApi)
	}
	return ress
}

func (list GraphQLApiList) Names() []string {
	var names []string
	for _, graphQLApi := range list {
		names = append(names, graphQLApi.GetMetadata().Name)
	}
	return names
}

func (list GraphQLApiList) NamespacesDotNames() []string {
	var names []string
	for _, graphQLApi := range list {
		names = append(names, graphQLApi.GetMetadata().Namespace+"."+graphQLApi.GetMetadata().Name)
	}
	return names
}

func (list GraphQLApiList) Sort() GraphQLApiList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list GraphQLApiList) Clone() GraphQLApiList {
	var graphQLApiList GraphQLApiList
	for _, graphQLApi := range list {
		graphQLApiList = append(graphQLApiList, resources.Clone(graphQLApi).(*GraphQLApi))
	}
	return graphQLApiList
}

func (list GraphQLApiList) Each(f func(element *GraphQLApi)) {
	for _, graphQLApi := range list {
		f(graphQLApi)
	}
}

func (list GraphQLApiList) EachResource(f func(element resources.Resource)) {
	for _, graphQLApi := range list {
		f(graphQLApi)
	}
}

func (list GraphQLApiList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *GraphQLApi) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for GraphQLApi

func (o *GraphQLApi) GetObjectKind() schema.ObjectKind {
	t := GraphQLApiCrd.TypeMeta()
	return &t
}

func (o *GraphQLApi) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*GraphQLApi)
}

func (o *GraphQLApi) DeepCopyInto(out *GraphQLApi) {
	clone := resources.Clone(o).(*GraphQLApi)
	*out = *clone
}

var (
	GraphQLApiCrd = crd.NewCrd(
		"graphqlapis",
		GraphQLApiGVK.Group,
		GraphQLApiGVK.Version,
		GraphQLApiGVK.Kind,
		"gqls",
		false,
		&GraphQLApi{})
)

var (
	GraphQLApiGVK = schema.GroupVersionKind{
		Version: "v1alpha1",
		Group:   "graphql.gloo.solo.io",
		Kind:    "GraphQLApi",
	}
)
