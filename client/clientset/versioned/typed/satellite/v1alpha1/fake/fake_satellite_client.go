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
	v1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/satellite/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSatelliteV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSatelliteV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeSatelliteV1alpha1) ClusterWorkerPools(namespace string) v1alpha1.ClusterWorkerPoolInterface {
	return &FakeClusterWorkerPools{c, namespace}
}

func (c *FakeSatelliteV1alpha1) Hosts(namespace string) v1alpha1.HostInterface {
	return &FakeHosts{c, namespace}
}

func (c *FakeSatelliteV1alpha1) Locations(namespace string) v1alpha1.LocationInterface {
	return &FakeLocations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSatelliteV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
