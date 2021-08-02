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

// DnsRecordsImportLister helps list DnsRecordsImports.
// All objects returned here must be treated as read-only.
type DnsRecordsImportLister interface {
	// List lists all DnsRecordsImports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsRecordsImport, err error)
	// DnsRecordsImports returns an object that can list and get DnsRecordsImports.
	DnsRecordsImports(namespace string) DnsRecordsImportNamespaceLister
	DnsRecordsImportListerExpansion
}

// dnsRecordsImportLister implements the DnsRecordsImportLister interface.
type dnsRecordsImportLister struct {
	indexer cache.Indexer
}

// NewDnsRecordsImportLister returns a new DnsRecordsImportLister.
func NewDnsRecordsImportLister(indexer cache.Indexer) DnsRecordsImportLister {
	return &dnsRecordsImportLister{indexer: indexer}
}

// List lists all DnsRecordsImports in the indexer.
func (s *dnsRecordsImportLister) List(selector labels.Selector) (ret []*v1alpha1.DnsRecordsImport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsRecordsImport))
	})
	return ret, err
}

// DnsRecordsImports returns an object that can list and get DnsRecordsImports.
func (s *dnsRecordsImportLister) DnsRecordsImports(namespace string) DnsRecordsImportNamespaceLister {
	return dnsRecordsImportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DnsRecordsImportNamespaceLister helps list and get DnsRecordsImports.
// All objects returned here must be treated as read-only.
type DnsRecordsImportNamespaceLister interface {
	// List lists all DnsRecordsImports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DnsRecordsImport, err error)
	// Get retrieves the DnsRecordsImport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DnsRecordsImport, error)
	DnsRecordsImportNamespaceListerExpansion
}

// dnsRecordsImportNamespaceLister implements the DnsRecordsImportNamespaceLister
// interface.
type dnsRecordsImportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DnsRecordsImports in the indexer for a given namespace.
func (s dnsRecordsImportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DnsRecordsImport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DnsRecordsImport))
	})
	return ret, err
}

// Get retrieves the DnsRecordsImport from the indexer for a given namespace and name.
func (s dnsRecordsImportNamespaceLister) Get(name string) (*v1alpha1.DnsRecordsImport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dnsrecordsimport"), name)
	}
	return obj.(*v1alpha1.DnsRecordsImport), nil
}