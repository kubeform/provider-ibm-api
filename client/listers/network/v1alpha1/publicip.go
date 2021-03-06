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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/network/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PublicIPLister helps list PublicIPs.
// All objects returned here must be treated as read-only.
type PublicIPLister interface {
	// List lists all PublicIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PublicIP, err error)
	// PublicIPs returns an object that can list and get PublicIPs.
	PublicIPs(namespace string) PublicIPNamespaceLister
	PublicIPListerExpansion
}

// publicIPLister implements the PublicIPLister interface.
type publicIPLister struct {
	indexer cache.Indexer
}

// NewPublicIPLister returns a new PublicIPLister.
func NewPublicIPLister(indexer cache.Indexer) PublicIPLister {
	return &publicIPLister{indexer: indexer}
}

// List lists all PublicIPs in the indexer.
func (s *publicIPLister) List(selector labels.Selector) (ret []*v1alpha1.PublicIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PublicIP))
	})
	return ret, err
}

// PublicIPs returns an object that can list and get PublicIPs.
func (s *publicIPLister) PublicIPs(namespace string) PublicIPNamespaceLister {
	return publicIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PublicIPNamespaceLister helps list and get PublicIPs.
// All objects returned here must be treated as read-only.
type PublicIPNamespaceLister interface {
	// List lists all PublicIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PublicIP, err error)
	// Get retrieves the PublicIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PublicIP, error)
	PublicIPNamespaceListerExpansion
}

// publicIPNamespaceLister implements the PublicIPNamespaceLister
// interface.
type publicIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PublicIPs in the indexer for a given namespace.
func (s publicIPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PublicIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PublicIP))
	})
	return ret, err
}

// Get retrieves the PublicIP from the indexer for a given namespace and name.
func (s publicIPNamespaceLister) Get(name string) (*v1alpha1.PublicIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("publicip"), name)
	}
	return obj.(*v1alpha1.PublicIP), nil
}
