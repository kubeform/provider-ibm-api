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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/container/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BindServiceLister helps list BindServices.
// All objects returned here must be treated as read-only.
type BindServiceLister interface {
	// List lists all BindServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BindService, err error)
	// BindServices returns an object that can list and get BindServices.
	BindServices(namespace string) BindServiceNamespaceLister
	BindServiceListerExpansion
}

// bindServiceLister implements the BindServiceLister interface.
type bindServiceLister struct {
	indexer cache.Indexer
}

// NewBindServiceLister returns a new BindServiceLister.
func NewBindServiceLister(indexer cache.Indexer) BindServiceLister {
	return &bindServiceLister{indexer: indexer}
}

// List lists all BindServices in the indexer.
func (s *bindServiceLister) List(selector labels.Selector) (ret []*v1alpha1.BindService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BindService))
	})
	return ret, err
}

// BindServices returns an object that can list and get BindServices.
func (s *bindServiceLister) BindServices(namespace string) BindServiceNamespaceLister {
	return bindServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BindServiceNamespaceLister helps list and get BindServices.
// All objects returned here must be treated as read-only.
type BindServiceNamespaceLister interface {
	// List lists all BindServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BindService, err error)
	// Get retrieves the BindService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BindService, error)
	BindServiceNamespaceListerExpansion
}

// bindServiceNamespaceLister implements the BindServiceNamespaceLister
// interface.
type bindServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BindServices in the indexer for a given namespace.
func (s bindServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BindService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BindService))
	})
	return ret, err
}

// Get retrieves the BindService from the indexer for a given namespace and name.
func (s bindServiceNamespaceLister) Get(name string) (*v1alpha1.BindService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bindservice"), name)
	}
	return obj.(*v1alpha1.BindService), nil
}