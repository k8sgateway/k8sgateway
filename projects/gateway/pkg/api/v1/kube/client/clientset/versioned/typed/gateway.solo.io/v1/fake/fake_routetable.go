/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	gatewaysoloiov1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouteTables implements RouteTableInterface
type FakeRouteTables struct {
	Fake *FakeGatewayV1
	ns   string
}

var routetablesResource = schema.GroupVersionResource{Group: "gateway.solo.io", Version: "v1", Resource: "routetables"}

var routetablesKind = schema.GroupVersionKind{Group: "gateway.solo.io", Version: "v1", Kind: "RouteTable"}

// Get takes name of the routeTable, and returns the corresponding routeTable object, and an error if there is any.
func (c *FakeRouteTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *gatewaysoloiov1.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routetablesResource, c.ns, name), &gatewaysoloiov1.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.RouteTable), err
}

// List takes label and field selectors, and returns the list of RouteTables that match those selectors.
func (c *FakeRouteTables) List(ctx context.Context, opts v1.ListOptions) (result *gatewaysoloiov1.RouteTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routetablesResource, routetablesKind, c.ns, opts), &gatewaysoloiov1.RouteTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gatewaysoloiov1.RouteTableList{ListMeta: obj.(*gatewaysoloiov1.RouteTableList).ListMeta}
	for _, item := range obj.(*gatewaysoloiov1.RouteTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routeTables.
func (c *FakeRouteTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routetablesResource, c.ns, opts))

}

// Create takes the representation of a routeTable and creates it.  Returns the server's representation of the routeTable, and an error, if there is any.
func (c *FakeRouteTables) Create(ctx context.Context, routeTable *gatewaysoloiov1.RouteTable, opts v1.CreateOptions) (result *gatewaysoloiov1.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routetablesResource, c.ns, routeTable), &gatewaysoloiov1.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.RouteTable), err
}

// Update takes the representation of a routeTable and updates it. Returns the server's representation of the routeTable, and an error, if there is any.
func (c *FakeRouteTables) Update(ctx context.Context, routeTable *gatewaysoloiov1.RouteTable, opts v1.UpdateOptions) (result *gatewaysoloiov1.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routetablesResource, c.ns, routeTable), &gatewaysoloiov1.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.RouteTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouteTables) UpdateStatus(ctx context.Context, routeTable *gatewaysoloiov1.RouteTable, opts v1.UpdateOptions) (*gatewaysoloiov1.RouteTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routetablesResource, "status", c.ns, routeTable), &gatewaysoloiov1.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.RouteTable), err
}

// Delete takes name of the routeTable and deletes it. Returns an error if one occurs.
func (c *FakeRouteTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routetablesResource, c.ns, name, opts), &gatewaysoloiov1.RouteTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouteTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routetablesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &gatewaysoloiov1.RouteTableList{})
	return err
}

// Patch applies the patch and returns the patched routeTable.
func (c *FakeRouteTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *gatewaysoloiov1.RouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routetablesResource, c.ns, name, pt, data, subresources...), &gatewaysoloiov1.RouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gatewaysoloiov1.RouteTable), err
}
