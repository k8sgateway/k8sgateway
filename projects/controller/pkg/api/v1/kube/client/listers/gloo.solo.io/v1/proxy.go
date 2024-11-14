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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ProxyLister helps list Proxies.
// All objects returned here must be treated as read-only.
type ProxyLister interface {
	// List lists all Proxies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Proxy, err error)
	// Proxies returns an object that can list and get Proxies.
	Proxies(namespace string) ProxyNamespaceLister
	ProxyListerExpansion
}

// proxyLister implements the ProxyLister interface.
type proxyLister struct {
	listers.ResourceIndexer[*v1.Proxy]
}

// NewProxyLister returns a new ProxyLister.
func NewProxyLister(indexer cache.Indexer) ProxyLister {
	return &proxyLister{listers.New[*v1.Proxy](indexer, v1.Resource("proxy"))}
}

// Proxies returns an object that can list and get Proxies.
func (s *proxyLister) Proxies(namespace string) ProxyNamespaceLister {
	return proxyNamespaceLister{listers.NewNamespaced[*v1.Proxy](s.ResourceIndexer, namespace)}
}

// ProxyNamespaceLister helps list and get Proxies.
// All objects returned here must be treated as read-only.
type ProxyNamespaceLister interface {
	// List lists all Proxies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Proxy, err error)
	// Get retrieves the Proxy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Proxy, error)
	ProxyNamespaceListerExpansion
}

// proxyNamespaceLister implements the ProxyNamespaceLister
// interface.
type proxyNamespaceLister struct {
	listers.ResourceIndexer[*v1.Proxy]
}
