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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/enterprise/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EnterpriseLister helps list Enterprises.
// All objects returned here must be treated as read-only.
type EnterpriseLister interface {
	// List lists all Enterprises in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Enterprise, err error)
	// Enterprises returns an object that can list and get Enterprises.
	Enterprises(namespace string) EnterpriseNamespaceLister
	EnterpriseListerExpansion
}

// enterpriseLister implements the EnterpriseLister interface.
type enterpriseLister struct {
	indexer cache.Indexer
}

// NewEnterpriseLister returns a new EnterpriseLister.
func NewEnterpriseLister(indexer cache.Indexer) EnterpriseLister {
	return &enterpriseLister{indexer: indexer}
}

// List lists all Enterprises in the indexer.
func (s *enterpriseLister) List(selector labels.Selector) (ret []*v1alpha1.Enterprise, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Enterprise))
	})
	return ret, err
}

// Enterprises returns an object that can list and get Enterprises.
func (s *enterpriseLister) Enterprises(namespace string) EnterpriseNamespaceLister {
	return enterpriseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EnterpriseNamespaceLister helps list and get Enterprises.
// All objects returned here must be treated as read-only.
type EnterpriseNamespaceLister interface {
	// List lists all Enterprises in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Enterprise, err error)
	// Get retrieves the Enterprise from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Enterprise, error)
	EnterpriseNamespaceListerExpansion
}

// enterpriseNamespaceLister implements the EnterpriseNamespaceLister
// interface.
type enterpriseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Enterprises in the indexer for a given namespace.
func (s enterpriseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Enterprise, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Enterprise))
	})
	return ret, err
}

// Get retrieves the Enterprise from the indexer for a given namespace and name.
func (s enterpriseNamespaceLister) Get(name string) (*v1alpha1.Enterprise, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("enterprise"), name)
	}
	return obj.(*v1alpha1.Enterprise), nil
}
