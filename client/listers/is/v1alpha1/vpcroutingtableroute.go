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

// VpcRoutingTableRouteLister helps list VpcRoutingTableRoutes.
// All objects returned here must be treated as read-only.
type VpcRoutingTableRouteLister interface {
	// List lists all VpcRoutingTableRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcRoutingTableRoute, err error)
	// VpcRoutingTableRoutes returns an object that can list and get VpcRoutingTableRoutes.
	VpcRoutingTableRoutes(namespace string) VpcRoutingTableRouteNamespaceLister
	VpcRoutingTableRouteListerExpansion
}

// vpcRoutingTableRouteLister implements the VpcRoutingTableRouteLister interface.
type vpcRoutingTableRouteLister struct {
	indexer cache.Indexer
}

// NewVpcRoutingTableRouteLister returns a new VpcRoutingTableRouteLister.
func NewVpcRoutingTableRouteLister(indexer cache.Indexer) VpcRoutingTableRouteLister {
	return &vpcRoutingTableRouteLister{indexer: indexer}
}

// List lists all VpcRoutingTableRoutes in the indexer.
func (s *vpcRoutingTableRouteLister) List(selector labels.Selector) (ret []*v1alpha1.VpcRoutingTableRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcRoutingTableRoute))
	})
	return ret, err
}

// VpcRoutingTableRoutes returns an object that can list and get VpcRoutingTableRoutes.
func (s *vpcRoutingTableRouteLister) VpcRoutingTableRoutes(namespace string) VpcRoutingTableRouteNamespaceLister {
	return vpcRoutingTableRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcRoutingTableRouteNamespaceLister helps list and get VpcRoutingTableRoutes.
// All objects returned here must be treated as read-only.
type VpcRoutingTableRouteNamespaceLister interface {
	// List lists all VpcRoutingTableRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcRoutingTableRoute, err error)
	// Get retrieves the VpcRoutingTableRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VpcRoutingTableRoute, error)
	VpcRoutingTableRouteNamespaceListerExpansion
}

// vpcRoutingTableRouteNamespaceLister implements the VpcRoutingTableRouteNamespaceLister
// interface.
type vpcRoutingTableRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcRoutingTableRoutes in the indexer for a given namespace.
func (s vpcRoutingTableRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcRoutingTableRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcRoutingTableRoute))
	})
	return ret, err
}

// Get retrieves the VpcRoutingTableRoute from the indexer for a given namespace and name.
func (s vpcRoutingTableRouteNamespaceLister) Get(name string) (*v1alpha1.VpcRoutingTableRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcroutingtableroute"), name)
	}
	return obj.(*v1alpha1.VpcRoutingTableRoute), nil
}