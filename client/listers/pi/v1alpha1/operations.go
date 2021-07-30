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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/pi/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OperationsLister helps list Operationses.
// All objects returned here must be treated as read-only.
type OperationsLister interface {
	// List lists all Operationses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Operations, err error)
	// Operationses returns an object that can list and get Operationses.
	Operationses(namespace string) OperationsNamespaceLister
	OperationsListerExpansion
}

// operationsLister implements the OperationsLister interface.
type operationsLister struct {
	indexer cache.Indexer
}

// NewOperationsLister returns a new OperationsLister.
func NewOperationsLister(indexer cache.Indexer) OperationsLister {
	return &operationsLister{indexer: indexer}
}

// List lists all Operationses in the indexer.
func (s *operationsLister) List(selector labels.Selector) (ret []*v1alpha1.Operations, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Operations))
	})
	return ret, err
}

// Operationses returns an object that can list and get Operationses.
func (s *operationsLister) Operationses(namespace string) OperationsNamespaceLister {
	return operationsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OperationsNamespaceLister helps list and get Operationses.
// All objects returned here must be treated as read-only.
type OperationsNamespaceLister interface {
	// List lists all Operationses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Operations, err error)
	// Get retrieves the Operations from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Operations, error)
	OperationsNamespaceListerExpansion
}

// operationsNamespaceLister implements the OperationsNamespaceLister
// interface.
type operationsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Operationses in the indexer for a given namespace.
func (s operationsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Operations, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Operations))
	})
	return ret, err
}

// Get retrieves the Operations from the indexer for a given namespace and name.
func (s operationsNamespaceLister) Get(name string) (*v1alpha1.Operations, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("operations"), name)
	}
	return obj.(*v1alpha1.Operations), nil
}
