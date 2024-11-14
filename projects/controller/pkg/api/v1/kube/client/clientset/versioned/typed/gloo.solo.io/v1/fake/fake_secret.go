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

// FakeSecrets implements SecretInterface
type FakeSecrets struct {
	Fake *FakeGlooV1
	ns   string
}

var secretsResource = v1.SchemeGroupVersion.WithResource("secrets")

var secretsKind = v1.SchemeGroupVersion.WithKind("Secret")

// Get takes name of the secret, and returns the corresponding secret object, and an error if there is any.
func (c *FakeSecrets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Secret, err error) {
	emptyResult := &v1.Secret{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(secretsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Secret), err
}

// List takes label and field selectors, and returns the list of Secrets that match those selectors.
func (c *FakeSecrets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SecretList, err error) {
	emptyResult := &v1.SecretList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(secretsResource, secretsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SecretList{ListMeta: obj.(*v1.SecretList).ListMeta}
	for _, item := range obj.(*v1.SecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secrets.
func (c *FakeSecrets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(secretsResource, c.ns, opts))

}

// Create takes the representation of a secret and creates it.  Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) (result *v1.Secret, err error) {
	emptyResult := &v1.Secret{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(secretsResource, c.ns, secret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Secret), err
}

// Update takes the representation of a secret and updates it. Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) (result *v1.Secret, err error) {
	emptyResult := &v1.Secret{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(secretsResource, c.ns, secret, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Secret), err
}

// Delete takes name of the secret and deletes it. Returns an error if one occurs.
func (c *FakeSecrets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(secretsResource, c.ns, name, opts), &v1.Secret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecrets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(secretsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SecretList{})
	return err
}

// Patch applies the patch and returns the patched secret.
func (c *FakeSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Secret, err error) {
	emptyResult := &v1.Secret{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(secretsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Secret), err
}
