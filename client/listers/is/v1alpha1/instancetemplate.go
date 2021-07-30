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

// InstanceTemplateLister helps list InstanceTemplates.
// All objects returned here must be treated as read-only.
type InstanceTemplateLister interface {
	// List lists all InstanceTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceTemplate, err error)
	// InstanceTemplates returns an object that can list and get InstanceTemplates.
	InstanceTemplates(namespace string) InstanceTemplateNamespaceLister
	InstanceTemplateListerExpansion
}

// instanceTemplateLister implements the InstanceTemplateLister interface.
type instanceTemplateLister struct {
	indexer cache.Indexer
}

// NewInstanceTemplateLister returns a new InstanceTemplateLister.
func NewInstanceTemplateLister(indexer cache.Indexer) InstanceTemplateLister {
	return &instanceTemplateLister{indexer: indexer}
}

// List lists all InstanceTemplates in the indexer.
func (s *instanceTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceTemplate))
	})
	return ret, err
}

// InstanceTemplates returns an object that can list and get InstanceTemplates.
func (s *instanceTemplateLister) InstanceTemplates(namespace string) InstanceTemplateNamespaceLister {
	return instanceTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceTemplateNamespaceLister helps list and get InstanceTemplates.
// All objects returned here must be treated as read-only.
type InstanceTemplateNamespaceLister interface {
	// List lists all InstanceTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceTemplate, err error)
	// Get retrieves the InstanceTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceTemplate, error)
	InstanceTemplateNamespaceListerExpansion
}

// instanceTemplateNamespaceLister implements the InstanceTemplateNamespaceLister
// interface.
type instanceTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceTemplates in the indexer for a given namespace.
func (s instanceTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceTemplate))
	})
	return ret, err
}

// Get retrieves the InstanceTemplate from the indexer for a given namespace and name.
func (s instanceTemplateNamespaceLister) Get(name string) (*v1alpha1.InstanceTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instancetemplate"), name)
	}
	return obj.(*v1alpha1.InstanceTemplate), nil
}
