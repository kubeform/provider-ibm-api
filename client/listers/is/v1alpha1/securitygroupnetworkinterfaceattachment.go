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

// SecurityGroupNetworkInterfaceAttachmentLister helps list SecurityGroupNetworkInterfaceAttachments.
// All objects returned here must be treated as read-only.
type SecurityGroupNetworkInterfaceAttachmentLister interface {
	// List lists all SecurityGroupNetworkInterfaceAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityGroupNetworkInterfaceAttachment, err error)
	// SecurityGroupNetworkInterfaceAttachments returns an object that can list and get SecurityGroupNetworkInterfaceAttachments.
	SecurityGroupNetworkInterfaceAttachments(namespace string) SecurityGroupNetworkInterfaceAttachmentNamespaceLister
	SecurityGroupNetworkInterfaceAttachmentListerExpansion
}

// securityGroupNetworkInterfaceAttachmentLister implements the SecurityGroupNetworkInterfaceAttachmentLister interface.
type securityGroupNetworkInterfaceAttachmentLister struct {
	indexer cache.Indexer
}

// NewSecurityGroupNetworkInterfaceAttachmentLister returns a new SecurityGroupNetworkInterfaceAttachmentLister.
func NewSecurityGroupNetworkInterfaceAttachmentLister(indexer cache.Indexer) SecurityGroupNetworkInterfaceAttachmentLister {
	return &securityGroupNetworkInterfaceAttachmentLister{indexer: indexer}
}

// List lists all SecurityGroupNetworkInterfaceAttachments in the indexer.
func (s *securityGroupNetworkInterfaceAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityGroupNetworkInterfaceAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityGroupNetworkInterfaceAttachment))
	})
	return ret, err
}

// SecurityGroupNetworkInterfaceAttachments returns an object that can list and get SecurityGroupNetworkInterfaceAttachments.
func (s *securityGroupNetworkInterfaceAttachmentLister) SecurityGroupNetworkInterfaceAttachments(namespace string) SecurityGroupNetworkInterfaceAttachmentNamespaceLister {
	return securityGroupNetworkInterfaceAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecurityGroupNetworkInterfaceAttachmentNamespaceLister helps list and get SecurityGroupNetworkInterfaceAttachments.
// All objects returned here must be treated as read-only.
type SecurityGroupNetworkInterfaceAttachmentNamespaceLister interface {
	// List lists all SecurityGroupNetworkInterfaceAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityGroupNetworkInterfaceAttachment, err error)
	// Get retrieves the SecurityGroupNetworkInterfaceAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SecurityGroupNetworkInterfaceAttachment, error)
	SecurityGroupNetworkInterfaceAttachmentNamespaceListerExpansion
}

// securityGroupNetworkInterfaceAttachmentNamespaceLister implements the SecurityGroupNetworkInterfaceAttachmentNamespaceLister
// interface.
type securityGroupNetworkInterfaceAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecurityGroupNetworkInterfaceAttachments in the indexer for a given namespace.
func (s securityGroupNetworkInterfaceAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityGroupNetworkInterfaceAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityGroupNetworkInterfaceAttachment))
	})
	return ret, err
}

// Get retrieves the SecurityGroupNetworkInterfaceAttachment from the indexer for a given namespace and name.
func (s securityGroupNetworkInterfaceAttachmentNamespaceLister) Get(name string) (*v1alpha1.SecurityGroupNetworkInterfaceAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("securitygroupnetworkinterfaceattachment"), name)
	}
	return obj.(*v1alpha1.SecurityGroupNetworkInterfaceAttachment), nil
}