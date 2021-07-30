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

// EdgeFunctionsTriggerLister helps list EdgeFunctionsTriggers.
// All objects returned here must be treated as read-only.
type EdgeFunctionsTriggerLister interface {
	// List lists all EdgeFunctionsTriggers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeFunctionsTrigger, err error)
	// EdgeFunctionsTriggers returns an object that can list and get EdgeFunctionsTriggers.
	EdgeFunctionsTriggers(namespace string) EdgeFunctionsTriggerNamespaceLister
	EdgeFunctionsTriggerListerExpansion
}

// edgeFunctionsTriggerLister implements the EdgeFunctionsTriggerLister interface.
type edgeFunctionsTriggerLister struct {
	indexer cache.Indexer
}

// NewEdgeFunctionsTriggerLister returns a new EdgeFunctionsTriggerLister.
func NewEdgeFunctionsTriggerLister(indexer cache.Indexer) EdgeFunctionsTriggerLister {
	return &edgeFunctionsTriggerLister{indexer: indexer}
}

// List lists all EdgeFunctionsTriggers in the indexer.
func (s *edgeFunctionsTriggerLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeFunctionsTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeFunctionsTrigger))
	})
	return ret, err
}

// EdgeFunctionsTriggers returns an object that can list and get EdgeFunctionsTriggers.
func (s *edgeFunctionsTriggerLister) EdgeFunctionsTriggers(namespace string) EdgeFunctionsTriggerNamespaceLister {
	return edgeFunctionsTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EdgeFunctionsTriggerNamespaceLister helps list and get EdgeFunctionsTriggers.
// All objects returned here must be treated as read-only.
type EdgeFunctionsTriggerNamespaceLister interface {
	// List lists all EdgeFunctionsTriggers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EdgeFunctionsTrigger, err error)
	// Get retrieves the EdgeFunctionsTrigger from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EdgeFunctionsTrigger, error)
	EdgeFunctionsTriggerNamespaceListerExpansion
}

// edgeFunctionsTriggerNamespaceLister implements the EdgeFunctionsTriggerNamespaceLister
// interface.
type edgeFunctionsTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EdgeFunctionsTriggers in the indexer for a given namespace.
func (s edgeFunctionsTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EdgeFunctionsTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EdgeFunctionsTrigger))
	})
	return ret, err
}

// Get retrieves the EdgeFunctionsTrigger from the indexer for a given namespace and name.
func (s edgeFunctionsTriggerNamespaceLister) Get(name string) (*v1alpha1.EdgeFunctionsTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("edgefunctionstrigger"), name)
	}
	return obj.(*v1alpha1.EdgeFunctionsTrigger), nil
}
