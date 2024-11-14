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

	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUpstreamGroups implements UpstreamGroupInterface
type FakeUpstreamGroups struct {
	Fake *FakeGlooV1
	ns   string
}

var upstreamgroupsResource = v1.SchemeGroupVersion.WithResource("upstreamgroups")

var upstreamgroupsKind = v1.SchemeGroupVersion.WithKind("UpstreamGroup")

// Get takes name of the upstreamGroup, and returns the corresponding upstreamGroup object, and an error if there is any.
func (c *FakeUpstreamGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.UpstreamGroup, err error) {
	emptyResult := &v1.UpstreamGroup{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(upstreamgroupsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.UpstreamGroup), err
}

// List takes label and field selectors, and returns the list of UpstreamGroups that match those selectors.
func (c *FakeUpstreamGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UpstreamGroupList, err error) {
	emptyResult := &v1.UpstreamGroupList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(upstreamgroupsResource, upstreamgroupsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.UpstreamGroupList{ListMeta: obj.(*v1.UpstreamGroupList).ListMeta}
	for _, item := range obj.(*v1.UpstreamGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested upstreamGroups.
func (c *FakeUpstreamGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(upstreamgroupsResource, c.ns, opts))

}

// Create takes the representation of a upstreamGroup and creates it.  Returns the server's representation of the upstreamGroup, and an error, if there is any.
func (c *FakeUpstreamGroups) Create(ctx context.Context, upstreamGroup *v1.UpstreamGroup, opts metav1.CreateOptions) (result *v1.UpstreamGroup, err error) {
	emptyResult := &v1.UpstreamGroup{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(upstreamgroupsResource, c.ns, upstreamGroup, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.UpstreamGroup), err
}

// Update takes the representation of a upstreamGroup and updates it. Returns the server's representation of the upstreamGroup, and an error, if there is any.
func (c *FakeUpstreamGroups) Update(ctx context.Context, upstreamGroup *v1.UpstreamGroup, opts metav1.UpdateOptions) (result *v1.UpstreamGroup, err error) {
	emptyResult := &v1.UpstreamGroup{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(upstreamgroupsResource, c.ns, upstreamGroup, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.UpstreamGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUpstreamGroups) UpdateStatus(ctx context.Context, upstreamGroup *v1.UpstreamGroup, opts metav1.UpdateOptions) (result *v1.UpstreamGroup, err error) {
	emptyResult := &v1.UpstreamGroup{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(upstreamgroupsResource, "status", c.ns, upstreamGroup, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.UpstreamGroup), err
}

// Delete takes name of the upstreamGroup and deletes it. Returns an error if one occurs.
func (c *FakeUpstreamGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(upstreamgroupsResource, c.ns, name, opts), &v1.UpstreamGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUpstreamGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(upstreamgroupsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.UpstreamGroupList{})
	return err
}

// Patch applies the patch and returns the patched upstreamGroup.
func (c *FakeUpstreamGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UpstreamGroup, err error) {
	emptyResult := &v1.UpstreamGroup{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(upstreamgroupsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.UpstreamGroup), err
}
