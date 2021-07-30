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
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/event/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StreamsTopicLister helps list StreamsTopics.
// All objects returned here must be treated as read-only.
type StreamsTopicLister interface {
	// List lists all StreamsTopics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamsTopic, err error)
	// StreamsTopics returns an object that can list and get StreamsTopics.
	StreamsTopics(namespace string) StreamsTopicNamespaceLister
	StreamsTopicListerExpansion
}

// streamsTopicLister implements the StreamsTopicLister interface.
type streamsTopicLister struct {
	indexer cache.Indexer
}

// NewStreamsTopicLister returns a new StreamsTopicLister.
func NewStreamsTopicLister(indexer cache.Indexer) StreamsTopicLister {
	return &streamsTopicLister{indexer: indexer}
}

// List lists all StreamsTopics in the indexer.
func (s *streamsTopicLister) List(selector labels.Selector) (ret []*v1alpha1.StreamsTopic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamsTopic))
	})
	return ret, err
}

// StreamsTopics returns an object that can list and get StreamsTopics.
func (s *streamsTopicLister) StreamsTopics(namespace string) StreamsTopicNamespaceLister {
	return streamsTopicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamsTopicNamespaceLister helps list and get StreamsTopics.
// All objects returned here must be treated as read-only.
type StreamsTopicNamespaceLister interface {
	// List lists all StreamsTopics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamsTopic, err error)
	// Get retrieves the StreamsTopic from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StreamsTopic, error)
	StreamsTopicNamespaceListerExpansion
}

// streamsTopicNamespaceLister implements the StreamsTopicNamespaceLister
// interface.
type streamsTopicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamsTopics in the indexer for a given namespace.
func (s streamsTopicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamsTopic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamsTopic))
	})
	return ret, err
}

// Get retrieves the StreamsTopic from the indexer for a given namespace and name.
func (s streamsTopicNamespaceLister) Get(name string) (*v1alpha1.StreamsTopic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamstopic"), name)
	}
	return obj.(*v1alpha1.StreamsTopic), nil
}
