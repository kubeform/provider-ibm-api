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

// TlsSettingsLister helps list TlsSettingses.
// All objects returned here must be treated as read-only.
type TlsSettingsLister interface {
	// List lists all TlsSettingses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TlsSettings, err error)
	// TlsSettingses returns an object that can list and get TlsSettingses.
	TlsSettingses(namespace string) TlsSettingsNamespaceLister
	TlsSettingsListerExpansion
}

// tlsSettingsLister implements the TlsSettingsLister interface.
type tlsSettingsLister struct {
	indexer cache.Indexer
}

// NewTlsSettingsLister returns a new TlsSettingsLister.
func NewTlsSettingsLister(indexer cache.Indexer) TlsSettingsLister {
	return &tlsSettingsLister{indexer: indexer}
}

// List lists all TlsSettingses in the indexer.
func (s *tlsSettingsLister) List(selector labels.Selector) (ret []*v1alpha1.TlsSettings, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TlsSettings))
	})
	return ret, err
}

// TlsSettingses returns an object that can list and get TlsSettingses.
func (s *tlsSettingsLister) TlsSettingses(namespace string) TlsSettingsNamespaceLister {
	return tlsSettingsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TlsSettingsNamespaceLister helps list and get TlsSettingses.
// All objects returned here must be treated as read-only.
type TlsSettingsNamespaceLister interface {
	// List lists all TlsSettingses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TlsSettings, err error)
	// Get retrieves the TlsSettings from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TlsSettings, error)
	TlsSettingsNamespaceListerExpansion
}

// tlsSettingsNamespaceLister implements the TlsSettingsNamespaceLister
// interface.
type tlsSettingsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TlsSettingses in the indexer for a given namespace.
func (s tlsSettingsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TlsSettings, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TlsSettings))
	})
	return ret, err
}

// Get retrieves the TlsSettings from the indexer for a given namespace and name.
func (s tlsSettingsNamespaceLister) Get(name string) (*v1alpha1.TlsSettings, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tlssettings"), name)
	}
	return obj.(*v1alpha1.TlsSettings), nil
}
