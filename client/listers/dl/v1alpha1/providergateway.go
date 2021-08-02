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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/dl/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProviderGatewayLister helps list ProviderGateways.
// All objects returned here must be treated as read-only.
type ProviderGatewayLister interface {
	// List lists all ProviderGateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderGateway, err error)
	// ProviderGateways returns an object that can list and get ProviderGateways.
	ProviderGateways(namespace string) ProviderGatewayNamespaceLister
	ProviderGatewayListerExpansion
}

// providerGatewayLister implements the ProviderGatewayLister interface.
type providerGatewayLister struct {
	indexer cache.Indexer
}

// NewProviderGatewayLister returns a new ProviderGatewayLister.
func NewProviderGatewayLister(indexer cache.Indexer) ProviderGatewayLister {
	return &providerGatewayLister{indexer: indexer}
}

// List lists all ProviderGateways in the indexer.
func (s *providerGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderGateway))
	})
	return ret, err
}

// ProviderGateways returns an object that can list and get ProviderGateways.
func (s *providerGatewayLister) ProviderGateways(namespace string) ProviderGatewayNamespaceLister {
	return providerGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProviderGatewayNamespaceLister helps list and get ProviderGateways.
// All objects returned here must be treated as read-only.
type ProviderGatewayNamespaceLister interface {
	// List lists all ProviderGateways in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProviderGateway, err error)
	// Get retrieves the ProviderGateway from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProviderGateway, error)
	ProviderGatewayNamespaceListerExpansion
}

// providerGatewayNamespaceLister implements the ProviderGatewayNamespaceLister
// interface.
type providerGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProviderGateways in the indexer for a given namespace.
func (s providerGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProviderGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProviderGateway))
	})
	return ret, err
}

// Get retrieves the ProviderGateway from the indexer for a given namespace and name.
func (s providerGatewayNamespaceLister) Get(name string) (*v1alpha1.ProviderGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("providergateway"), name)
	}
	return obj.(*v1alpha1.ProviderGateway), nil
}