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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	securityv1alpha1 "kubeform.dev/provider-ibm-api/apis/security/v1alpha1"
	versioned "kubeform.dev/provider-ibm-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-ibm-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-ibm-api/client/listers/security/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GroupRuleInformer provides access to a shared informer and lister for
// GroupRules.
type GroupRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GroupRuleLister
}

type groupRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGroupRuleInformer constructs a new informer for GroupRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGroupRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGroupRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGroupRuleInformer constructs a new informer for GroupRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGroupRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1alpha1().GroupRules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1alpha1().GroupRules(namespace).Watch(context.TODO(), options)
			},
		},
		&securityv1alpha1.GroupRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *groupRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGroupRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *groupRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&securityv1alpha1.GroupRule{}, f.defaultInformer)
}

func (f *groupRuleInformer) Lister() v1alpha1.GroupRuleLister {
	return v1alpha1.NewGroupRuleLister(f.Informer().GetIndexer())
}