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
	v1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/container/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeContainerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeContainerV1alpha1) Addonses(namespace string) v1alpha1.AddonsInterface {
	return &FakeAddonses{c, namespace}
}

func (c *FakeContainerV1alpha1) Albs(namespace string) v1alpha1.AlbInterface {
	return &FakeAlbs{c, namespace}
}

func (c *FakeContainerV1alpha1) AlbCerts(namespace string) v1alpha1.AlbCertInterface {
	return &FakeAlbCerts{c, namespace}
}

func (c *FakeContainerV1alpha1) ApiKeyResets(namespace string) v1alpha1.ApiKeyResetInterface {
	return &FakeApiKeyResets{c, namespace}
}

func (c *FakeContainerV1alpha1) BindServices(namespace string) v1alpha1.BindServiceInterface {
	return &FakeBindServices{c, namespace}
}

func (c *FakeContainerV1alpha1) Clusters(namespace string) v1alpha1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeContainerV1alpha1) ClusterFeatures(namespace string) v1alpha1.ClusterFeatureInterface {
	return &FakeClusterFeatures{c, namespace}
}

func (c *FakeContainerV1alpha1) VpcAlbs(namespace string) v1alpha1.VpcAlbInterface {
	return &FakeVpcAlbs{c, namespace}
}

func (c *FakeContainerV1alpha1) VpcClusters(namespace string) v1alpha1.VpcClusterInterface {
	return &FakeVpcClusters{c, namespace}
}

func (c *FakeContainerV1alpha1) VpcWorkerPools(namespace string) v1alpha1.VpcWorkerPoolInterface {
	return &FakeVpcWorkerPools{c, namespace}
}

func (c *FakeContainerV1alpha1) WorkerPools(namespace string) v1alpha1.WorkerPoolInterface {
	return &FakeWorkerPools{c, namespace}
}

func (c *FakeContainerV1alpha1) WorkerPoolZoneAttachments(namespace string) v1alpha1.WorkerPoolZoneAttachmentInterface {
	return &FakeWorkerPoolZoneAttachments{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeContainerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}