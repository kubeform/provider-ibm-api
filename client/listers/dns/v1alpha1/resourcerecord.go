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

// ResourceRecordLister helps list ResourceRecords.
// All objects returned here must be treated as read-only.
type ResourceRecordLister interface {
	// List lists all ResourceRecords in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceRecord, err error)
	// ResourceRecords returns an object that can list and get ResourceRecords.
	ResourceRecords(namespace string) ResourceRecordNamespaceLister
	ResourceRecordListerExpansion
}

// resourceRecordLister implements the ResourceRecordLister interface.
type resourceRecordLister struct {
	indexer cache.Indexer
}

// NewResourceRecordLister returns a new ResourceRecordLister.
func NewResourceRecordLister(indexer cache.Indexer) ResourceRecordLister {
	return &resourceRecordLister{indexer: indexer}
}

// List lists all ResourceRecords in the indexer.
func (s *resourceRecordLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceRecord, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceRecord))
	})
	return ret, err
}

// ResourceRecords returns an object that can list and get ResourceRecords.
func (s *resourceRecordLister) ResourceRecords(namespace string) ResourceRecordNamespaceLister {
	return resourceRecordNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceRecordNamespaceLister helps list and get ResourceRecords.
// All objects returned here must be treated as read-only.
type ResourceRecordNamespaceLister interface {
	// List lists all ResourceRecords in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceRecord, err error)
	// Get retrieves the ResourceRecord from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceRecord, error)
	ResourceRecordNamespaceListerExpansion
}

// resourceRecordNamespaceLister implements the ResourceRecordNamespaceLister
// interface.
type resourceRecordNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceRecords in the indexer for a given namespace.
func (s resourceRecordNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceRecord, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceRecord))
	})
	return ret, err
}

// Get retrieves the ResourceRecord from the indexer for a given namespace and name.
func (s resourceRecordNamespaceLister) Get(name string) (*v1alpha1.ResourceRecord, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourcerecord"), name)
	}
	return obj.(*v1alpha1.ResourceRecord), nil
}
