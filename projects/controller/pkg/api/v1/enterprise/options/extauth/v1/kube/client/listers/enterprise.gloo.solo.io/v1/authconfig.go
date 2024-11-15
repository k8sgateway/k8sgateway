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
	v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extauth/v1/kube/apis/enterprise.gloo.solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthConfigLister helps list AuthConfigs.
// All objects returned here must be treated as read-only.
type AuthConfigLister interface {
	// List lists all AuthConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AuthConfig, err error)
	// AuthConfigs returns an object that can list and get AuthConfigs.
	AuthConfigs(namespace string) AuthConfigNamespaceLister
	AuthConfigListerExpansion
}

// authConfigLister implements the AuthConfigLister interface.
type authConfigLister struct {
	indexer cache.Indexer
}

// NewAuthConfigLister returns a new AuthConfigLister.
func NewAuthConfigLister(indexer cache.Indexer) AuthConfigLister {
	return &authConfigLister{indexer: indexer}
}

// List lists all AuthConfigs in the indexer.
func (s *authConfigLister) List(selector labels.Selector) (ret []*v1.AuthConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AuthConfig))
	})
	return ret, err
}

// AuthConfigs returns an object that can list and get AuthConfigs.
func (s *authConfigLister) AuthConfigs(namespace string) AuthConfigNamespaceLister {
	return authConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuthConfigNamespaceLister helps list and get AuthConfigs.
// All objects returned here must be treated as read-only.
type AuthConfigNamespaceLister interface {
	// List lists all AuthConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AuthConfig, err error)
	// Get retrieves the AuthConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AuthConfig, error)
	AuthConfigNamespaceListerExpansion
}

// authConfigNamespaceLister implements the AuthConfigNamespaceLister
// interface.
type authConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuthConfigs in the indexer for a given namespace.
func (s authConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.AuthConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AuthConfig))
	})
	return ret, err
}

// Get retrieves the AuthConfig from the indexer for a given namespace and name.
func (s authConfigNamespaceLister) Get(name string) (*v1.AuthConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("authconfig"), name)
	}
	return obj.(*v1.AuthConfig), nil
}
