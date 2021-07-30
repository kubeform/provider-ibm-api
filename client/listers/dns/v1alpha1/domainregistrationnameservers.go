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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/dns/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainRegistrationNameserversLister helps list DomainRegistrationNameserverses.
// All objects returned here must be treated as read-only.
type DomainRegistrationNameserversLister interface {
	// List lists all DomainRegistrationNameserverses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainRegistrationNameservers, err error)
	// DomainRegistrationNameserverses returns an object that can list and get DomainRegistrationNameserverses.
	DomainRegistrationNameserverses(namespace string) DomainRegistrationNameserversNamespaceLister
	DomainRegistrationNameserversListerExpansion
}

// domainRegistrationNameserversLister implements the DomainRegistrationNameserversLister interface.
type domainRegistrationNameserversLister struct {
	indexer cache.Indexer
}

// NewDomainRegistrationNameserversLister returns a new DomainRegistrationNameserversLister.
func NewDomainRegistrationNameserversLister(indexer cache.Indexer) DomainRegistrationNameserversLister {
	return &domainRegistrationNameserversLister{indexer: indexer}
}

// List lists all DomainRegistrationNameserverses in the indexer.
func (s *domainRegistrationNameserversLister) List(selector labels.Selector) (ret []*v1alpha1.DomainRegistrationNameservers, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainRegistrationNameservers))
	})
	return ret, err
}

// DomainRegistrationNameserverses returns an object that can list and get DomainRegistrationNameserverses.
func (s *domainRegistrationNameserversLister) DomainRegistrationNameserverses(namespace string) DomainRegistrationNameserversNamespaceLister {
	return domainRegistrationNameserversNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainRegistrationNameserversNamespaceLister helps list and get DomainRegistrationNameserverses.
// All objects returned here must be treated as read-only.
type DomainRegistrationNameserversNamespaceLister interface {
	// List lists all DomainRegistrationNameserverses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainRegistrationNameservers, err error)
	// Get retrieves the DomainRegistrationNameservers from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainRegistrationNameservers, error)
	DomainRegistrationNameserversNamespaceListerExpansion
}

// domainRegistrationNameserversNamespaceLister implements the DomainRegistrationNameserversNamespaceLister
// interface.
type domainRegistrationNameserversNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainRegistrationNameserverses in the indexer for a given namespace.
func (s domainRegistrationNameserversNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainRegistrationNameservers, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainRegistrationNameservers))
	})
	return ret, err
}

// Get retrieves the DomainRegistrationNameservers from the indexer for a given namespace and name.
func (s domainRegistrationNameserversNamespaceLister) Get(name string) (*v1alpha1.DomainRegistrationNameservers, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainregistrationnameservers"), name)
	}
	return obj.(*v1alpha1.DomainRegistrationNameservers), nil
}
