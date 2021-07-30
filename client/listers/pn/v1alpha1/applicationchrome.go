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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/pn/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationChromeLister helps list ApplicationChromes.
// All objects returned here must be treated as read-only.
type ApplicationChromeLister interface {
	// List lists all ApplicationChromes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationChrome, err error)
	// ApplicationChromes returns an object that can list and get ApplicationChromes.
	ApplicationChromes(namespace string) ApplicationChromeNamespaceLister
	ApplicationChromeListerExpansion
}

// applicationChromeLister implements the ApplicationChromeLister interface.
type applicationChromeLister struct {
	indexer cache.Indexer
}

// NewApplicationChromeLister returns a new ApplicationChromeLister.
func NewApplicationChromeLister(indexer cache.Indexer) ApplicationChromeLister {
	return &applicationChromeLister{indexer: indexer}
}

// List lists all ApplicationChromes in the indexer.
func (s *applicationChromeLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationChrome, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationChrome))
	})
	return ret, err
}

// ApplicationChromes returns an object that can list and get ApplicationChromes.
func (s *applicationChromeLister) ApplicationChromes(namespace string) ApplicationChromeNamespaceLister {
	return applicationChromeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationChromeNamespaceLister helps list and get ApplicationChromes.
// All objects returned here must be treated as read-only.
type ApplicationChromeNamespaceLister interface {
	// List lists all ApplicationChromes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationChrome, err error)
	// Get retrieves the ApplicationChrome from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApplicationChrome, error)
	ApplicationChromeNamespaceListerExpansion
}

// applicationChromeNamespaceLister implements the ApplicationChromeNamespaceLister
// interface.
type applicationChromeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationChromes in the indexer for a given namespace.
func (s applicationChromeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationChrome, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationChrome))
	})
	return ret, err
}

// Get retrieves the ApplicationChrome from the indexer for a given namespace and name.
func (s applicationChromeNamespaceLister) Get(name string) (*v1alpha1.ApplicationChrome, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationchrome"), name)
	}
	return obj.(*v1alpha1.ApplicationChrome), nil
}
