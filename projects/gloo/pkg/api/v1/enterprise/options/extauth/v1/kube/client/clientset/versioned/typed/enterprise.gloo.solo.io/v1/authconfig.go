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

package v1

import (
	"time"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1/kube/apis/enterprise.gloo.solo.io/v1"
	scheme "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1/kube/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthConfigsGetter has a method to return a AuthConfigInterface.
// A group's client should implement this interface.
type AuthConfigsGetter interface {
	AuthConfigs(namespace string) AuthConfigInterface
}

// AuthConfigInterface has methods to work with AuthConfig resources.
type AuthConfigInterface interface {
	Create(*v1.AuthConfig) (*v1.AuthConfig, error)
	Update(*v1.AuthConfig) (*v1.AuthConfig, error)
	UpdateStatus(*v1.AuthConfig) (*v1.AuthConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.AuthConfig, error)
	List(opts metav1.ListOptions) (*v1.AuthConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AuthConfig, err error)
	AuthConfigExpansion
}

// authConfigs implements AuthConfigInterface
type authConfigs struct {
	client rest.Interface
	ns     string
}

// newAuthConfigs returns a AuthConfigs
func newAuthConfigs(c *EnterpriseV1Client, namespace string) *authConfigs {
	return &authConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the authConfig, and returns the corresponding authConfig object, and an error if there is any.
func (c *authConfigs) Get(name string, options metav1.GetOptions) (result *v1.AuthConfig, err error) {
	result = &v1.AuthConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthConfigs that match those selectors.
func (c *authConfigs) List(opts metav1.ListOptions) (result *v1.AuthConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AuthConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authConfigs.
func (c *authConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("authconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a authConfig and creates it.  Returns the server's representation of the authConfig, and an error, if there is any.
func (c *authConfigs) Create(authConfig *v1.AuthConfig) (result *v1.AuthConfig, err error) {
	result = &v1.AuthConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("authconfigs").
		Body(authConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a authConfig and updates it. Returns the server's representation of the authConfig, and an error, if there is any.
func (c *authConfigs) Update(authConfig *v1.AuthConfig) (result *v1.AuthConfig, err error) {
	result = &v1.AuthConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authconfigs").
		Name(authConfig.Name).
		Body(authConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *authConfigs) UpdateStatus(authConfig *v1.AuthConfig) (result *v1.AuthConfig, err error) {
	result = &v1.AuthConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authconfigs").
		Name(authConfig.Name).
		SubResource("status").
		Body(authConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the authConfig and deletes it. Returns an error if one occurs.
func (c *authConfigs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched authConfig.
func (c *authConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AuthConfig, err error) {
	result = &v1.AuthConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("authconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
