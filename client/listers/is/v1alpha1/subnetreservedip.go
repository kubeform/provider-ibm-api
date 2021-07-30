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

// SubnetReservedIPLister helps list SubnetReservedIPs.
// All objects returned here must be treated as read-only.
type SubnetReservedIPLister interface {
	// List lists all SubnetReservedIPs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetReservedIP, err error)
	// SubnetReservedIPs returns an object that can list and get SubnetReservedIPs.
	SubnetReservedIPs(namespace string) SubnetReservedIPNamespaceLister
	SubnetReservedIPListerExpansion
}

// subnetReservedIPLister implements the SubnetReservedIPLister interface.
type subnetReservedIPLister struct {
	indexer cache.Indexer
}

// NewSubnetReservedIPLister returns a new SubnetReservedIPLister.
func NewSubnetReservedIPLister(indexer cache.Indexer) SubnetReservedIPLister {
	return &subnetReservedIPLister{indexer: indexer}
}

// List lists all SubnetReservedIPs in the indexer.
func (s *subnetReservedIPLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetReservedIP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetReservedIP))
	})
	return ret, err
}

// SubnetReservedIPs returns an object that can list and get SubnetReservedIPs.
func (s *subnetReservedIPLister) SubnetReservedIPs(namespace string) SubnetReservedIPNamespaceLister {
	return subnetReservedIPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubnetReservedIPNamespaceLister helps list and get SubnetReservedIPs.
// All objects returned here must be treated as read-only.
type SubnetReservedIPNamespaceLister interface {
	// List lists all SubnetReservedIPs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetReservedIP, err error)
	// Get retrieves the SubnetReservedIP from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SubnetReservedIP, error)
	SubnetReservedIPNamespaceListerExpansion
}

// subnetReservedIPNamespaceLister implements the SubnetReservedIPNamespaceLister
// interface.
type subnetReservedIPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SubnetReservedIPs in the indexer for a given namespace.
func (s subnetReservedIPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetReservedIP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetReservedIP))
	})
	return ret, err
}

// Get retrieves the SubnetReservedIP from the indexer for a given namespace and name.
func (s subnetReservedIPNamespaceLister) Get(name string) (*v1alpha1.SubnetReservedIP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subnetreservedip"), name)
	}
	return obj.(*v1alpha1.SubnetReservedIP), nil
}
