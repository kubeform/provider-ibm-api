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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/iam/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServicePolicyLister helps list ServicePolicies.
// All objects returned here must be treated as read-only.
type ServicePolicyLister interface {
	// List lists all ServicePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicePolicy, err error)
	// ServicePolicies returns an object that can list and get ServicePolicies.
	ServicePolicies(namespace string) ServicePolicyNamespaceLister
	ServicePolicyListerExpansion
}

// servicePolicyLister implements the ServicePolicyLister interface.
type servicePolicyLister struct {
	indexer cache.Indexer
}

// NewServicePolicyLister returns a new ServicePolicyLister.
func NewServicePolicyLister(indexer cache.Indexer) ServicePolicyLister {
	return &servicePolicyLister{indexer: indexer}
}

// List lists all ServicePolicies in the indexer.
func (s *servicePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ServicePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicePolicy))
	})
	return ret, err
}

// ServicePolicies returns an object that can list and get ServicePolicies.
func (s *servicePolicyLister) ServicePolicies(namespace string) ServicePolicyNamespaceLister {
	return servicePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServicePolicyNamespaceLister helps list and get ServicePolicies.
// All objects returned here must be treated as read-only.
type ServicePolicyNamespaceLister interface {
	// List lists all ServicePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServicePolicy, err error)
	// Get retrieves the ServicePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServicePolicy, error)
	ServicePolicyNamespaceListerExpansion
}

// servicePolicyNamespaceLister implements the ServicePolicyNamespaceLister
// interface.
type servicePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServicePolicies in the indexer for a given namespace.
func (s servicePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServicePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServicePolicy))
	})
	return ret, err
}

// Get retrieves the ServicePolicy from the indexer for a given namespace and name.
func (s servicePolicyNamespaceLister) Get(name string) (*v1alpha1.ServicePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicepolicy"), name)
	}
	return obj.(*v1alpha1.ServicePolicy), nil
}