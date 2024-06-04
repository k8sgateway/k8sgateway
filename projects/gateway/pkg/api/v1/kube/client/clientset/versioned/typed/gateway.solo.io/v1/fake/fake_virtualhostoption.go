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
	json "encoding/json"
	"fmt"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	gatewaysoloiov1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/client/applyconfiguration/gateway.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualHostOptions implements VirtualHostOptionInterface
type FakeVirtualHostOptions struct {
	Fake *FakeGatewayV1
	ns   string
}

var virtualhostoptionsResource = v1.SchemeGroupVersion.WithResource("virtualhostoptions")

var virtualhostoptionsKind = v1.SchemeGroupVersion.WithKind("VirtualHostOption")

// Get takes name of the virtualHostOption, and returns the corresponding virtualHostOption object, and an error if there is any.
func (c *FakeVirtualHostOptions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualHostOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualhostoptionsResource, c.ns, name), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}

// List takes label and field selectors, and returns the list of VirtualHostOptions that match those selectors.
func (c *FakeVirtualHostOptions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualHostOptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualhostoptionsResource, virtualhostoptionsKind, c.ns, opts), &v1.VirtualHostOptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.VirtualHostOptionList{ListMeta: obj.(*v1.VirtualHostOptionList).ListMeta}
	for _, item := range obj.(*v1.VirtualHostOptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualHostOptions.
func (c *FakeVirtualHostOptions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualhostoptionsResource, c.ns, opts))

}

// Create takes the representation of a virtualHostOption and creates it.  Returns the server's representation of the virtualHostOption, and an error, if there is any.
func (c *FakeVirtualHostOptions) Create(ctx context.Context, virtualHostOption *v1.VirtualHostOption, opts metav1.CreateOptions) (result *v1.VirtualHostOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualhostoptionsResource, c.ns, virtualHostOption), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}

// Update takes the representation of a virtualHostOption and updates it. Returns the server's representation of the virtualHostOption, and an error, if there is any.
func (c *FakeVirtualHostOptions) Update(ctx context.Context, virtualHostOption *v1.VirtualHostOption, opts metav1.UpdateOptions) (result *v1.VirtualHostOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualhostoptionsResource, c.ns, virtualHostOption), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualHostOptions) UpdateStatus(ctx context.Context, virtualHostOption *v1.VirtualHostOption, opts metav1.UpdateOptions) (*v1.VirtualHostOption, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualhostoptionsResource, "status", c.ns, virtualHostOption), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}

// Delete takes name of the virtualHostOption and deletes it. Returns an error if one occurs.
func (c *FakeVirtualHostOptions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualhostoptionsResource, c.ns, name, opts), &v1.VirtualHostOption{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualHostOptions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualhostoptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.VirtualHostOptionList{})
	return err
}

// Patch applies the patch and returns the patched virtualHostOption.
func (c *FakeVirtualHostOptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualHostOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostoptionsResource, c.ns, name, pt, data, subresources...), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied virtualHostOption.
func (c *FakeVirtualHostOptions) Apply(ctx context.Context, virtualHostOption *gatewaysoloiov1.VirtualHostOptionApplyConfiguration, opts metav1.ApplyOptions) (result *v1.VirtualHostOption, err error) {
	if virtualHostOption == nil {
		return nil, fmt.Errorf("virtualHostOption provided to Apply must not be nil")
	}
	data, err := json.Marshal(virtualHostOption)
	if err != nil {
		return nil, err
	}
	name := virtualHostOption.Name
	if name == nil {
		return nil, fmt.Errorf("virtualHostOption.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostoptionsResource, c.ns, *name, types.ApplyPatchType, data), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeVirtualHostOptions) ApplyStatus(ctx context.Context, virtualHostOption *gatewaysoloiov1.VirtualHostOptionApplyConfiguration, opts metav1.ApplyOptions) (result *v1.VirtualHostOption, err error) {
	if virtualHostOption == nil {
		return nil, fmt.Errorf("virtualHostOption provided to Apply must not be nil")
	}
	data, err := json.Marshal(virtualHostOption)
	if err != nil {
		return nil, err
	}
	name := virtualHostOption.Name
	if name == nil {
		return nil, fmt.Errorf("virtualHostOption.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostoptionsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.VirtualHostOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualHostOption), err
}
