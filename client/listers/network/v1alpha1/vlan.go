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

// VlanLister helps list Vlans.
// All objects returned here must be treated as read-only.
type VlanLister interface {
	// List lists all Vlans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Vlan, err error)
	// Vlans returns an object that can list and get Vlans.
	Vlans(namespace string) VlanNamespaceLister
	VlanListerExpansion
}

// vlanLister implements the VlanLister interface.
type vlanLister struct {
	indexer cache.Indexer
}

// NewVlanLister returns a new VlanLister.
func NewVlanLister(indexer cache.Indexer) VlanLister {
	return &vlanLister{indexer: indexer}
}

// List lists all Vlans in the indexer.
func (s *vlanLister) List(selector labels.Selector) (ret []*v1alpha1.Vlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Vlan))
	})
	return ret, err
}

// Vlans returns an object that can list and get Vlans.
func (s *vlanLister) Vlans(namespace string) VlanNamespaceLister {
	return vlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VlanNamespaceLister helps list and get Vlans.
// All objects returned here must be treated as read-only.
type VlanNamespaceLister interface {
	// List lists all Vlans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Vlan, err error)
	// Get retrieves the Vlan from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Vlan, error)
	VlanNamespaceListerExpansion
}

// vlanNamespaceLister implements the VlanNamespaceLister
// interface.
type vlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Vlans in the indexer for a given namespace.
func (s vlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Vlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Vlan))
	})
	return ret, err
}

// Get retrieves the Vlan from the indexer for a given namespace and name.
func (s vlanNamespaceLister) Get(name string) (*v1alpha1.Vlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vlan"), name)
	}
	return obj.(*v1alpha1.Vlan), nil
}