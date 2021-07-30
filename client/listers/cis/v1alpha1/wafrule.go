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

// WafRuleLister helps list WafRules.
// All objects returned here must be treated as read-only.
type WafRuleLister interface {
	// List lists all WafRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WafRule, err error)
	// WafRules returns an object that can list and get WafRules.
	WafRules(namespace string) WafRuleNamespaceLister
	WafRuleListerExpansion
}

// wafRuleLister implements the WafRuleLister interface.
type wafRuleLister struct {
	indexer cache.Indexer
}

// NewWafRuleLister returns a new WafRuleLister.
func NewWafRuleLister(indexer cache.Indexer) WafRuleLister {
	return &wafRuleLister{indexer: indexer}
}

// List lists all WafRules in the indexer.
func (s *wafRuleLister) List(selector labels.Selector) (ret []*v1alpha1.WafRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafRule))
	})
	return ret, err
}

// WafRules returns an object that can list and get WafRules.
func (s *wafRuleLister) WafRules(namespace string) WafRuleNamespaceLister {
	return wafRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafRuleNamespaceLister helps list and get WafRules.
// All objects returned here must be treated as read-only.
type WafRuleNamespaceLister interface {
	// List lists all WafRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WafRule, err error)
	// Get retrieves the WafRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WafRule, error)
	WafRuleNamespaceListerExpansion
}

// wafRuleNamespaceLister implements the WafRuleNamespaceLister
// interface.
type wafRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafRules in the indexer for a given namespace.
func (s wafRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafRule))
	})
	return ret, err
}

// Get retrieves the WafRule from the indexer for a given namespace and name.
func (s wafRuleNamespaceLister) Get(name string) (*v1alpha1.WafRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafrule"), name)
	}
	return obj.(*v1alpha1.WafRule), nil
}
