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
	"net/http"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	"github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type GatewayV1Interface interface {
	RESTClient() rest.Interface
	GatewaysGetter
	HttpListenerOptionsGetter
	ListenerOptionsGetter
	MatchableHttpGatewaysGetter
	MatchableTcpGatewaysGetter
	RouteOptionsGetter
	RouteTablesGetter
	VirtualHostOptionsGetter
	VirtualServicesGetter
}

// GatewayV1Client is used to interact with features provided by the gateway.solo.io group.
type GatewayV1Client struct {
	restClient rest.Interface
}

func (c *GatewayV1Client) Gateways(namespace string) GatewayInterface {
	return newGateways(c, namespace)
}

func (c *GatewayV1Client) HttpListenerOptions(namespace string) HttpListenerOptionInterface {
	return newHttpListenerOptions(c, namespace)
}

func (c *GatewayV1Client) ListenerOptions(namespace string) ListenerOptionInterface {
	return newListenerOptions(c, namespace)
}

func (c *GatewayV1Client) MatchableHttpGateways(namespace string) MatchableHttpGatewayInterface {
	return newMatchableHttpGateways(c, namespace)
}

func (c *GatewayV1Client) MatchableTcpGateways(namespace string) MatchableTcpGatewayInterface {
	return newMatchableTcpGateways(c, namespace)
}

func (c *GatewayV1Client) RouteOptions(namespace string) RouteOptionInterface {
	return newRouteOptions(c, namespace)
}

func (c *GatewayV1Client) RouteTables(namespace string) RouteTableInterface {
	return newRouteTables(c, namespace)
}

func (c *GatewayV1Client) VirtualHostOptions(namespace string) VirtualHostOptionInterface {
	return newVirtualHostOptions(c, namespace)
}

func (c *GatewayV1Client) VirtualServices(namespace string) VirtualServiceInterface {
	return newVirtualServices(c, namespace)
}

// NewForConfig creates a new GatewayV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*GatewayV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new GatewayV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*GatewayV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &GatewayV1Client{client}, nil
}

// NewForConfigOrDie creates a new GatewayV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GatewayV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GatewayV1Client for the given RESTClient.
func New(c rest.Interface) *GatewayV1Client {
	return &GatewayV1Client{c}
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
func (c *GatewayV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
