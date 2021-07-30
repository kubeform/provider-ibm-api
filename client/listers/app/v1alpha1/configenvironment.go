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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/app/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigEnvironmentLister helps list ConfigEnvironments.
// All objects returned here must be treated as read-only.
type ConfigEnvironmentLister interface {
	// List lists all ConfigEnvironments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigEnvironment, err error)
	// ConfigEnvironments returns an object that can list and get ConfigEnvironments.
	ConfigEnvironments(namespace string) ConfigEnvironmentNamespaceLister
	ConfigEnvironmentListerExpansion
}

// configEnvironmentLister implements the ConfigEnvironmentLister interface.
type configEnvironmentLister struct {
	indexer cache.Indexer
}

// NewConfigEnvironmentLister returns a new ConfigEnvironmentLister.
func NewConfigEnvironmentLister(indexer cache.Indexer) ConfigEnvironmentLister {
	return &configEnvironmentLister{indexer: indexer}
}

// List lists all ConfigEnvironments in the indexer.
func (s *configEnvironmentLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigEnvironment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigEnvironment))
	})
	return ret, err
}

// ConfigEnvironments returns an object that can list and get ConfigEnvironments.
func (s *configEnvironmentLister) ConfigEnvironments(namespace string) ConfigEnvironmentNamespaceLister {
	return configEnvironmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigEnvironmentNamespaceLister helps list and get ConfigEnvironments.
// All objects returned here must be treated as read-only.
type ConfigEnvironmentNamespaceLister interface {
	// List lists all ConfigEnvironments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigEnvironment, err error)
	// Get retrieves the ConfigEnvironment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConfigEnvironment, error)
	ConfigEnvironmentNamespaceListerExpansion
}

// configEnvironmentNamespaceLister implements the ConfigEnvironmentNamespaceLister
// interface.
type configEnvironmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigEnvironments in the indexer for a given namespace.
func (s configEnvironmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigEnvironment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigEnvironment))
	})
	return ret, err
}

// Get retrieves the ConfigEnvironment from the indexer for a given namespace and name.
func (s configEnvironmentNamespaceLister) Get(name string) (*v1alpha1.ConfigEnvironment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configenvironment"), name)
	}
	return obj.(*v1alpha1.ConfigEnvironment), nil
}
