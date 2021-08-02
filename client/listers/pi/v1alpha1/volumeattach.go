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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/pi/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeAttachLister helps list VolumeAttaches.
// All objects returned here must be treated as read-only.
type VolumeAttachLister interface {
	// List lists all VolumeAttaches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeAttach, err error)
	// VolumeAttaches returns an object that can list and get VolumeAttaches.
	VolumeAttaches(namespace string) VolumeAttachNamespaceLister
	VolumeAttachListerExpansion
}

// volumeAttachLister implements the VolumeAttachLister interface.
type volumeAttachLister struct {
	indexer cache.Indexer
}

// NewVolumeAttachLister returns a new VolumeAttachLister.
func NewVolumeAttachLister(indexer cache.Indexer) VolumeAttachLister {
	return &volumeAttachLister{indexer: indexer}
}

// List lists all VolumeAttaches in the indexer.
func (s *volumeAttachLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeAttach, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeAttach))
	})
	return ret, err
}

// VolumeAttaches returns an object that can list and get VolumeAttaches.
func (s *volumeAttachLister) VolumeAttaches(namespace string) VolumeAttachNamespaceLister {
	return volumeAttachNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeAttachNamespaceLister helps list and get VolumeAttaches.
// All objects returned here must be treated as read-only.
type VolumeAttachNamespaceLister interface {
	// List lists all VolumeAttaches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeAttach, err error)
	// Get retrieves the VolumeAttach from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VolumeAttach, error)
	VolumeAttachNamespaceListerExpansion
}

// volumeAttachNamespaceLister implements the VolumeAttachNamespaceLister
// interface.
type volumeAttachNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeAttaches in the indexer for a given namespace.
func (s volumeAttachNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeAttach, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeAttach))
	})
	return ret, err
}

// Get retrieves the VolumeAttach from the indexer for a given namespace and name.
func (s volumeAttachNamespaceLister) Get(name string) (*v1alpha1.VolumeAttach, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("volumeattach"), name)
	}
	return obj.(*v1alpha1.VolumeAttach), nil
}