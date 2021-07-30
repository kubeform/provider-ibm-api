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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/container/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkerPoolZoneAttachmentLister helps list WorkerPoolZoneAttachments.
// All objects returned here must be treated as read-only.
type WorkerPoolZoneAttachmentLister interface {
	// List lists all WorkerPoolZoneAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkerPoolZoneAttachment, err error)
	// WorkerPoolZoneAttachments returns an object that can list and get WorkerPoolZoneAttachments.
	WorkerPoolZoneAttachments(namespace string) WorkerPoolZoneAttachmentNamespaceLister
	WorkerPoolZoneAttachmentListerExpansion
}

// workerPoolZoneAttachmentLister implements the WorkerPoolZoneAttachmentLister interface.
type workerPoolZoneAttachmentLister struct {
	indexer cache.Indexer
}

// NewWorkerPoolZoneAttachmentLister returns a new WorkerPoolZoneAttachmentLister.
func NewWorkerPoolZoneAttachmentLister(indexer cache.Indexer) WorkerPoolZoneAttachmentLister {
	return &workerPoolZoneAttachmentLister{indexer: indexer}
}

// List lists all WorkerPoolZoneAttachments in the indexer.
func (s *workerPoolZoneAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.WorkerPoolZoneAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkerPoolZoneAttachment))
	})
	return ret, err
}

// WorkerPoolZoneAttachments returns an object that can list and get WorkerPoolZoneAttachments.
func (s *workerPoolZoneAttachmentLister) WorkerPoolZoneAttachments(namespace string) WorkerPoolZoneAttachmentNamespaceLister {
	return workerPoolZoneAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkerPoolZoneAttachmentNamespaceLister helps list and get WorkerPoolZoneAttachments.
// All objects returned here must be treated as read-only.
type WorkerPoolZoneAttachmentNamespaceLister interface {
	// List lists all WorkerPoolZoneAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorkerPoolZoneAttachment, err error)
	// Get retrieves the WorkerPoolZoneAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorkerPoolZoneAttachment, error)
	WorkerPoolZoneAttachmentNamespaceListerExpansion
}

// workerPoolZoneAttachmentNamespaceLister implements the WorkerPoolZoneAttachmentNamespaceLister
// interface.
type workerPoolZoneAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkerPoolZoneAttachments in the indexer for a given namespace.
func (s workerPoolZoneAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorkerPoolZoneAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorkerPoolZoneAttachment))
	})
	return ret, err
}

// Get retrieves the WorkerPoolZoneAttachment from the indexer for a given namespace and name.
func (s workerPoolZoneAttachmentNamespaceLister) Get(name string) (*v1alpha1.WorkerPoolZoneAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workerpoolzoneattachment"), name)
	}
	return obj.(*v1alpha1.WorkerPoolZoneAttachment), nil
}
