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

// WafGroupLister helps list WafGroups.
// All objects returned here must be treated as read-only.
type WafGroupLister interface {
	// List lists all WafGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WafGroup, err error)
	// WafGroups returns an object that can list and get WafGroups.
	WafGroups(namespace string) WafGroupNamespaceLister
	WafGroupListerExpansion
}

// wafGroupLister implements the WafGroupLister interface.
type wafGroupLister struct {
	indexer cache.Indexer
}

// NewWafGroupLister returns a new WafGroupLister.
func NewWafGroupLister(indexer cache.Indexer) WafGroupLister {
	return &wafGroupLister{indexer: indexer}
}

// List lists all WafGroups in the indexer.
func (s *wafGroupLister) List(selector labels.Selector) (ret []*v1alpha1.WafGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafGroup))
	})
	return ret, err
}

// WafGroups returns an object that can list and get WafGroups.
func (s *wafGroupLister) WafGroups(namespace string) WafGroupNamespaceLister {
	return wafGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafGroupNamespaceLister helps list and get WafGroups.
// All objects returned here must be treated as read-only.
type WafGroupNamespaceLister interface {
	// List lists all WafGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WafGroup, err error)
	// Get retrieves the WafGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WafGroup, error)
	WafGroupNamespaceListerExpansion
}

// wafGroupNamespaceLister implements the WafGroupNamespaceLister
// interface.
type wafGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafGroups in the indexer for a given namespace.
func (s wafGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafGroup))
	})
	return ret, err
}

// Get retrieves the WafGroup from the indexer for a given namespace and name.
func (s wafGroupNamespaceLister) Get(name string) (*v1alpha1.WafGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafgroup"), name)
	}
	return obj.(*v1alpha1.WafGroup), nil
}