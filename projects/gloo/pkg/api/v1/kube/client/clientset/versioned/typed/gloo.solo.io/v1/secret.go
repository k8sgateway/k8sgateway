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
	"context"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	scheme "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SecretsGetter has a method to return a SecretInterface.
// A group's client should implement this interface.
type SecretsGetter interface {
	Secrets(namespace string) SecretInterface
}

// SecretInterface has methods to work with Secret resources.
type SecretInterface interface {
	Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) (*v1.Secret, error)
	Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) (*v1.Secret, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Secret, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SecretList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Secret, err error)
	SecretExpansion
}

// secrets implements SecretInterface
type secrets struct {
	*gentype.ClientWithList[*v1.Secret, *v1.SecretList]
}

// newSecrets returns a Secrets
func newSecrets(c *GlooV1Client, namespace string) *secrets {
	return &secrets{
		gentype.NewClientWithList[*v1.Secret, *v1.SecretList](
			"secrets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Secret { return &v1.Secret{} },
			func() *v1.SecretList { return &v1.SecretList{} }),
	}
}
