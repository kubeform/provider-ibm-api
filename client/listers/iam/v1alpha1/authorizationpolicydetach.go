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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/iam/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthorizationPolicyDetachLister helps list AuthorizationPolicyDetaches.
// All objects returned here must be treated as read-only.
type AuthorizationPolicyDetachLister interface {
	// List lists all AuthorizationPolicyDetaches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AuthorizationPolicyDetach, err error)
	// AuthorizationPolicyDetaches returns an object that can list and get AuthorizationPolicyDetaches.
	AuthorizationPolicyDetaches(namespace string) AuthorizationPolicyDetachNamespaceLister
	AuthorizationPolicyDetachListerExpansion
}

// authorizationPolicyDetachLister implements the AuthorizationPolicyDetachLister interface.
type authorizationPolicyDetachLister struct {
	indexer cache.Indexer
}

// NewAuthorizationPolicyDetachLister returns a new AuthorizationPolicyDetachLister.
func NewAuthorizationPolicyDetachLister(indexer cache.Indexer) AuthorizationPolicyDetachLister {
	return &authorizationPolicyDetachLister{indexer: indexer}
}

// List lists all AuthorizationPolicyDetaches in the indexer.
func (s *authorizationPolicyDetachLister) List(selector labels.Selector) (ret []*v1alpha1.AuthorizationPolicyDetach, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthorizationPolicyDetach))
	})
	return ret, err
}

// AuthorizationPolicyDetaches returns an object that can list and get AuthorizationPolicyDetaches.
func (s *authorizationPolicyDetachLister) AuthorizationPolicyDetaches(namespace string) AuthorizationPolicyDetachNamespaceLister {
	return authorizationPolicyDetachNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuthorizationPolicyDetachNamespaceLister helps list and get AuthorizationPolicyDetaches.
// All objects returned here must be treated as read-only.
type AuthorizationPolicyDetachNamespaceLister interface {
	// List lists all AuthorizationPolicyDetaches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AuthorizationPolicyDetach, err error)
	// Get retrieves the AuthorizationPolicyDetach from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AuthorizationPolicyDetach, error)
	AuthorizationPolicyDetachNamespaceListerExpansion
}

// authorizationPolicyDetachNamespaceLister implements the AuthorizationPolicyDetachNamespaceLister
// interface.
type authorizationPolicyDetachNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuthorizationPolicyDetaches in the indexer for a given namespace.
func (s authorizationPolicyDetachNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AuthorizationPolicyDetach, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuthorizationPolicyDetach))
	})
	return ret, err
}

// Get retrieves the AuthorizationPolicyDetach from the indexer for a given namespace and name.
func (s authorizationPolicyDetachNamespaceLister) Get(name string) (*v1alpha1.AuthorizationPolicyDetach, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("authorizationpolicydetach"), name)
	}
	return obj.(*v1alpha1.AuthorizationPolicyDetach), nil
}
