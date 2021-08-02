/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/cis/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CacheSettingsLister helps list CacheSettingses.
// All objects returned here must be treated as read-only.
type CacheSettingsLister interface {
	// List lists all CacheSettingses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CacheSettings, err error)
	// CacheSettingses returns an object that can list and get CacheSettingses.
	CacheSettingses(namespace string) CacheSettingsNamespaceLister
	CacheSettingsListerExpansion
}

// cacheSettingsLister implements the CacheSettingsLister interface.
type cacheSettingsLister struct {
	indexer cache.Indexer
}

// NewCacheSettingsLister returns a new CacheSettingsLister.
func NewCacheSettingsLister(indexer cache.Indexer) CacheSettingsLister {
	return &cacheSettingsLister{indexer: indexer}
}

// List lists all CacheSettingses in the indexer.
func (s *cacheSettingsLister) List(selector labels.Selector) (ret []*v1alpha1.CacheSettings, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CacheSettings))
	})
	return ret, err
}

// CacheSettingses returns an object that can list and get CacheSettingses.
func (s *cacheSettingsLister) CacheSettingses(namespace string) CacheSettingsNamespaceLister {
	return cacheSettingsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CacheSettingsNamespaceLister helps list and get CacheSettingses.
// All objects returned here must be treated as read-only.
type CacheSettingsNamespaceLister interface {
	// List lists all CacheSettingses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CacheSettings, err error)
	// Get retrieves the CacheSettings from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CacheSettings, error)
	CacheSettingsNamespaceListerExpansion
}

// cacheSettingsNamespaceLister implements the CacheSettingsNamespaceLister
// interface.
type cacheSettingsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CacheSettingses in the indexer for a given namespace.
func (s cacheSettingsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CacheSettings, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CacheSettings))
	})
	return ret, err
}

// Get retrieves the CacheSettings from the indexer for a given namespace and name.
func (s cacheSettingsNamespaceLister) Get(name string) (*v1alpha1.CacheSettings, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cachesettings"), name)
	}
	return obj.(*v1alpha1.CacheSettings), nil
}