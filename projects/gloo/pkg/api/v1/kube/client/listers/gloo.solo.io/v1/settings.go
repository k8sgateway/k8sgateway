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
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SettingsLister helps list Settingses.
// All objects returned here must be treated as read-only.
type SettingsLister interface {
	// List lists all Settingses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Settings, err error)
	// Settingses returns an object that can list and get Settingses.
	Settingses(namespace string) SettingsNamespaceLister
	SettingsListerExpansion
}

// settingsLister implements the SettingsLister interface.
type settingsLister struct {
	indexer cache.Indexer
}

// NewSettingsLister returns a new SettingsLister.
func NewSettingsLister(indexer cache.Indexer) SettingsLister {
	return &settingsLister{indexer: indexer}
}

// List lists all Settingses in the indexer.
func (s *settingsLister) List(selector labels.Selector) (ret []*v1.Settings, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Settings))
	})
	return ret, err
}

// Settingses returns an object that can list and get Settingses.
func (s *settingsLister) Settingses(namespace string) SettingsNamespaceLister {
	return settingsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SettingsNamespaceLister helps list and get Settingses.
// All objects returned here must be treated as read-only.
type SettingsNamespaceLister interface {
	// List lists all Settingses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Settings, err error)
	// Get retrieves the Settings from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Settings, error)
	SettingsNamespaceListerExpansion
}

// settingsNamespaceLister implements the SettingsNamespaceLister
// interface.
type settingsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Settingses in the indexer for a given namespace.
func (s settingsNamespaceLister) List(selector labels.Selector) (ret []*v1.Settings, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Settings))
	})
	return ret, err
}

// Get retrieves the Settings from the indexer for a given namespace and name.
func (s settingsNamespaceLister) Get(name string) (*v1.Settings, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("settings"), name)
	}
	return obj.(*v1.Settings), nil
}
