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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lb/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLbV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLbV1alpha1) Lbs(namespace string) v1alpha1.LbInterface {
	return &FakeLbs{c, namespace}
}

func (c *FakeLbV1alpha1) Services(namespace string) v1alpha1.ServiceInterface {
	return &FakeServices{c, namespace}
}

func (c *FakeLbV1alpha1) ServiceGroups(namespace string) v1alpha1.ServiceGroupInterface {
	return &FakeServiceGroups{c, namespace}
}

func (c *FakeLbV1alpha1) Vpxes(namespace string) v1alpha1.VpxInterface {
	return &FakeVpxes{c, namespace}
}

func (c *FakeLbV1alpha1) VpxHas(namespace string) v1alpha1.VpxHaInterface {
	return &FakeVpxHas{c, namespace}
}

func (c *FakeLbV1alpha1) VpxServices(namespace string) v1alpha1.VpxServiceInterface {
	return &FakeVpxServices{c, namespace}
}

func (c *FakeLbV1alpha1) VpxVips(namespace string) v1alpha1.VpxVipInterface {
	return &FakeVpxVips{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLbV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}