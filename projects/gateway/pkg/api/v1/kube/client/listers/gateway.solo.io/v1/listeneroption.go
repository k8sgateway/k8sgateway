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
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ListenerOptionLister helps list ListenerOptions.
// All objects returned here must be treated as read-only.
type ListenerOptionLister interface {
	// List lists all ListenerOptions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ListenerOption, err error)
	// ListenerOptions returns an object that can list and get ListenerOptions.
	ListenerOptions(namespace string) ListenerOptionNamespaceLister
	ListenerOptionListerExpansion
}

// listenerOptionLister implements the ListenerOptionLister interface.
type listenerOptionLister struct {
	indexer cache.Indexer
}

// NewListenerOptionLister returns a new ListenerOptionLister.
func NewListenerOptionLister(indexer cache.Indexer) ListenerOptionLister {
	return &listenerOptionLister{indexer: indexer}
}

// List lists all ListenerOptions in the indexer.
func (s *listenerOptionLister) List(selector labels.Selector) (ret []*v1.ListenerOption, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ListenerOption))
	})
	return ret, err
}

// ListenerOptions returns an object that can list and get ListenerOptions.
func (s *listenerOptionLister) ListenerOptions(namespace string) ListenerOptionNamespaceLister {
	return listenerOptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ListenerOptionNamespaceLister helps list and get ListenerOptions.
// All objects returned here must be treated as read-only.
type ListenerOptionNamespaceLister interface {
	// List lists all ListenerOptions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ListenerOption, err error)
	// Get retrieves the ListenerOption from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ListenerOption, error)
	ListenerOptionNamespaceListerExpansion
}

// listenerOptionNamespaceLister implements the ListenerOptionNamespaceLister
// interface.
type listenerOptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ListenerOptions in the indexer for a given namespace.
func (s listenerOptionNamespaceLister) List(selector labels.Selector) (ret []*v1.ListenerOption, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ListenerOption))
	})
	return ret, err
}

// Get retrieves the ListenerOption from the indexer for a given namespace and name.
func (s listenerOptionNamespaceLister) Get(name string) (*v1.ListenerOption, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("listeneroption"), name)
	}
	return obj.(*v1.ListenerOption), nil
}
