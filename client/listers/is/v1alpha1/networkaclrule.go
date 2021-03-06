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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/is/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkACLRuleLister helps list NetworkACLRules.
// All objects returned here must be treated as read-only.
type NetworkACLRuleLister interface {
	// List lists all NetworkACLRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkACLRule, err error)
	// NetworkACLRules returns an object that can list and get NetworkACLRules.
	NetworkACLRules(namespace string) NetworkACLRuleNamespaceLister
	NetworkACLRuleListerExpansion
}

// networkACLRuleLister implements the NetworkACLRuleLister interface.
type networkACLRuleLister struct {
	indexer cache.Indexer
}

// NewNetworkACLRuleLister returns a new NetworkACLRuleLister.
func NewNetworkACLRuleLister(indexer cache.Indexer) NetworkACLRuleLister {
	return &networkACLRuleLister{indexer: indexer}
}

// List lists all NetworkACLRules in the indexer.
func (s *networkACLRuleLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkACLRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkACLRule))
	})
	return ret, err
}

// NetworkACLRules returns an object that can list and get NetworkACLRules.
func (s *networkACLRuleLister) NetworkACLRules(namespace string) NetworkACLRuleNamespaceLister {
	return networkACLRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkACLRuleNamespaceLister helps list and get NetworkACLRules.
// All objects returned here must be treated as read-only.
type NetworkACLRuleNamespaceLister interface {
	// List lists all NetworkACLRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkACLRule, err error)
	// Get retrieves the NetworkACLRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NetworkACLRule, error)
	NetworkACLRuleNamespaceListerExpansion
}

// networkACLRuleNamespaceLister implements the NetworkACLRuleNamespaceLister
// interface.
type networkACLRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkACLRules in the indexer for a given namespace.
func (s networkACLRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkACLRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkACLRule))
	})
	return ret, err
}

// Get retrieves the NetworkACLRule from the indexer for a given namespace and name.
func (s networkACLRuleNamespaceLister) Get(name string) (*v1alpha1.NetworkACLRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkaclrule"), name)
	}
	return obj.(*v1alpha1.NetworkACLRule), nil
}
