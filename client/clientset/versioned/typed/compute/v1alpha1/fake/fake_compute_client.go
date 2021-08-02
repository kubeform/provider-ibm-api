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
	v1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/compute/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeComputeV1alpha1 struct {
	*testing.Fake
}

func (c *FakeComputeV1alpha1) AutoscaleGroups(namespace string) v1alpha1.AutoscaleGroupInterface {
	return &FakeAutoscaleGroups{c, namespace}
}

func (c *FakeComputeV1alpha1) AutoscalePolicies(namespace string) v1alpha1.AutoscalePolicyInterface {
	return &FakeAutoscalePolicies{c, namespace}
}

func (c *FakeComputeV1alpha1) BareMetals(namespace string) v1alpha1.BareMetalInterface {
	return &FakeBareMetals{c, namespace}
}

func (c *FakeComputeV1alpha1) DedicatedHosts(namespace string) v1alpha1.DedicatedHostInterface {
	return &FakeDedicatedHosts{c, namespace}
}

func (c *FakeComputeV1alpha1) Monitors(namespace string) v1alpha1.MonitorInterface {
	return &FakeMonitors{c, namespace}
}

func (c *FakeComputeV1alpha1) PlacementGroups(namespace string) v1alpha1.PlacementGroupInterface {
	return &FakePlacementGroups{c, namespace}
}

func (c *FakeComputeV1alpha1) ProvisioningHooks(namespace string) v1alpha1.ProvisioningHookInterface {
	return &FakeProvisioningHooks{c, namespace}
}

func (c *FakeComputeV1alpha1) SshKeys(namespace string) v1alpha1.SshKeyInterface {
	return &FakeSshKeys{c, namespace}
}

func (c *FakeComputeV1alpha1) SslCertificates(namespace string) v1alpha1.SslCertificateInterface {
	return &FakeSslCertificates{c, namespace}
}

func (c *FakeComputeV1alpha1) Users(namespace string) v1alpha1.UserInterface {
	return &FakeUsers{c, namespace}
}

func (c *FakeComputeV1alpha1) VmInstances(namespace string) v1alpha1.VmInstanceInterface {
	return &FakeVmInstances{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeComputeV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}