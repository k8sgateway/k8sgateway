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

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMatchableTcpGateways implements MatchableTcpGatewayInterface
type FakeMatchableTcpGateways struct {
	Fake *FakeGatewayV1
	ns   string
}

var matchabletcpgatewaysResource = v1.SchemeGroupVersion.WithResource("tcpgateways")

var matchabletcpgatewaysKind = v1.SchemeGroupVersion.WithKind("MatchableTcpGateway")

// Get takes name of the matchableTcpGateway, and returns the corresponding matchableTcpGateway object, and an error if there is any.
func (c *FakeMatchableTcpGateways) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MatchableTcpGateway, err error) {
	emptyResult := &v1.MatchableTcpGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(matchabletcpgatewaysResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MatchableTcpGateway), err
}

// List takes label and field selectors, and returns the list of MatchableTcpGateways that match those selectors.
func (c *FakeMatchableTcpGateways) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MatchableTcpGatewayList, err error) {
	emptyResult := &v1.MatchableTcpGatewayList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(matchabletcpgatewaysResource, matchabletcpgatewaysKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.MatchableTcpGatewayList{ListMeta: obj.(*v1.MatchableTcpGatewayList).ListMeta}
	for _, item := range obj.(*v1.MatchableTcpGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested matchableTcpGateways.
func (c *FakeMatchableTcpGateways) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(matchabletcpgatewaysResource, c.ns, opts))

}

// Create takes the representation of a matchableTcpGateway and creates it.  Returns the server's representation of the matchableTcpGateway, and an error, if there is any.
func (c *FakeMatchableTcpGateways) Create(ctx context.Context, matchableTcpGateway *v1.MatchableTcpGateway, opts metav1.CreateOptions) (result *v1.MatchableTcpGateway, err error) {
	emptyResult := &v1.MatchableTcpGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(matchabletcpgatewaysResource, c.ns, matchableTcpGateway, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MatchableTcpGateway), err
}

// Update takes the representation of a matchableTcpGateway and updates it. Returns the server's representation of the matchableTcpGateway, and an error, if there is any.
func (c *FakeMatchableTcpGateways) Update(ctx context.Context, matchableTcpGateway *v1.MatchableTcpGateway, opts metav1.UpdateOptions) (result *v1.MatchableTcpGateway, err error) {
	emptyResult := &v1.MatchableTcpGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(matchabletcpgatewaysResource, c.ns, matchableTcpGateway, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MatchableTcpGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMatchableTcpGateways) UpdateStatus(ctx context.Context, matchableTcpGateway *v1.MatchableTcpGateway, opts metav1.UpdateOptions) (result *v1.MatchableTcpGateway, err error) {
	emptyResult := &v1.MatchableTcpGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(matchabletcpgatewaysResource, "status", c.ns, matchableTcpGateway, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MatchableTcpGateway), err
}

// Delete takes name of the matchableTcpGateway and deletes it. Returns an error if one occurs.
func (c *FakeMatchableTcpGateways) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(matchabletcpgatewaysResource, c.ns, name, opts), &v1.MatchableTcpGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMatchableTcpGateways) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(matchabletcpgatewaysResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.MatchableTcpGatewayList{})
	return err
}

// Patch applies the patch and returns the patched matchableTcpGateway.
func (c *FakeMatchableTcpGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MatchableTcpGateway, err error) {
	emptyResult := &v1.MatchableTcpGateway{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(matchabletcpgatewaysResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.MatchableTcpGateway), err
}
