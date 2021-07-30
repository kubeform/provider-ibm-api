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

// ZoneLister helps list Zones.
// All objects returned here must be treated as read-only.
type ZoneLister interface {
	// List lists all Zones in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Zone, err error)
	// Zones returns an object that can list and get Zones.
	Zones(namespace string) ZoneNamespaceLister
	ZoneListerExpansion
}

// zoneLister implements the ZoneLister interface.
type zoneLister struct {
	indexer cache.Indexer
}

// NewZoneLister returns a new ZoneLister.
func NewZoneLister(indexer cache.Indexer) ZoneLister {
	return &zoneLister{indexer: indexer}
}

// List lists all Zones in the indexer.
func (s *zoneLister) List(selector labels.Selector) (ret []*v1alpha1.Zone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Zone))
	})
	return ret, err
}

// Zones returns an object that can list and get Zones.
func (s *zoneLister) Zones(namespace string) ZoneNamespaceLister {
	return zoneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ZoneNamespaceLister helps list and get Zones.
// All objects returned here must be treated as read-only.
type ZoneNamespaceLister interface {
	// List lists all Zones in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Zone, err error)
	// Get retrieves the Zone from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Zone, error)
	ZoneNamespaceListerExpansion
}

// zoneNamespaceLister implements the ZoneNamespaceLister
// interface.
type zoneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Zones in the indexer for a given namespace.
func (s zoneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Zone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Zone))
	})
	return ret, err
}

// Get retrieves the Zone from the indexer for a given namespace and name.
func (s zoneNamespaceLister) Get(name string) (*v1alpha1.Zone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("zone"), name)
	}
	return obj.(*v1alpha1.Zone), nil
}
