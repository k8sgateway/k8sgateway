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
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1/kube/apis/enterprise.gloo.solo.io/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1/kube/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type EnterpriseV1Interface interface {
	RESTClient() rest.Interface
	AuthConfigsGetter
}

// EnterpriseV1Client is used to interact with features provided by the enterprise.gloo.solo.io group.
type EnterpriseV1Client struct {
	restClient rest.Interface
}

func (c *EnterpriseV1Client) AuthConfigs(namespace string) AuthConfigInterface {
	return newAuthConfigs(c, namespace)
}

// NewForConfig creates a new EnterpriseV1Client for the given config.
func NewForConfig(c *rest.Config) (*EnterpriseV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &EnterpriseV1Client{client}, nil
}

// NewForConfigOrDie creates a new EnterpriseV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *EnterpriseV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new EnterpriseV1Client for the given RESTClient.
func New(c rest.Interface) *EnterpriseV1Client {
	return &EnterpriseV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *EnterpriseV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
