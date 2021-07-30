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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/cm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CatalogLister helps list Catalogs.
// All objects returned here must be treated as read-only.
type CatalogLister interface {
	// List lists all Catalogs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Catalog, err error)
	// Catalogs returns an object that can list and get Catalogs.
	Catalogs(namespace string) CatalogNamespaceLister
	CatalogListerExpansion
}

// catalogLister implements the CatalogLister interface.
type catalogLister struct {
	indexer cache.Indexer
}

// NewCatalogLister returns a new CatalogLister.
func NewCatalogLister(indexer cache.Indexer) CatalogLister {
	return &catalogLister{indexer: indexer}
}

// List lists all Catalogs in the indexer.
func (s *catalogLister) List(selector labels.Selector) (ret []*v1alpha1.Catalog, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Catalog))
	})
	return ret, err
}

// Catalogs returns an object that can list and get Catalogs.
func (s *catalogLister) Catalogs(namespace string) CatalogNamespaceLister {
	return catalogNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CatalogNamespaceLister helps list and get Catalogs.
// All objects returned here must be treated as read-only.
type CatalogNamespaceLister interface {
	// List lists all Catalogs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Catalog, err error)
	// Get retrieves the Catalog from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Catalog, error)
	CatalogNamespaceListerExpansion
}

// catalogNamespaceLister implements the CatalogNamespaceLister
// interface.
type catalogNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Catalogs in the indexer for a given namespace.
func (s catalogNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Catalog, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Catalog))
	})
	return ret, err
}

// Get retrieves the Catalog from the indexer for a given namespace and name.
func (s catalogNamespaceLister) Get(name string) (*v1alpha1.Catalog, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("catalog"), name)
	}
	return obj.(*v1alpha1.Catalog), nil
}
