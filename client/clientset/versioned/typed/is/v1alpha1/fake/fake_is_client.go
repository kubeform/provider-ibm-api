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
	v1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/is/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIsV1alpha1) DedicatedHosts(namespace string) v1alpha1.DedicatedHostInterface {
	return &FakeDedicatedHosts{c, namespace}
}

func (c *FakeIsV1alpha1) DedicatedHostDiskManagements(namespace string) v1alpha1.DedicatedHostDiskManagementInterface {
	return &FakeDedicatedHostDiskManagements{c, namespace}
}

func (c *FakeIsV1alpha1) DedicatedHostGroups(namespace string) v1alpha1.DedicatedHostGroupInterface {
	return &FakeDedicatedHostGroups{c, namespace}
}

func (c *FakeIsV1alpha1) FloatingIPs(namespace string) v1alpha1.FloatingIPInterface {
	return &FakeFloatingIPs{c, namespace}
}

func (c *FakeIsV1alpha1) FlowLogs(namespace string) v1alpha1.FlowLogInterface {
	return &FakeFlowLogs{c, namespace}
}

func (c *FakeIsV1alpha1) IkePolicies(namespace string) v1alpha1.IkePolicyInterface {
	return &FakeIkePolicies{c, namespace}
}

func (c *FakeIsV1alpha1) Images(namespace string) v1alpha1.ImageInterface {
	return &FakeImages{c, namespace}
}

func (c *FakeIsV1alpha1) Instances(namespace string) v1alpha1.InstanceInterface {
	return &FakeInstances{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceDiskManagements(namespace string) v1alpha1.InstanceDiskManagementInterface {
	return &FakeInstanceDiskManagements{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceGroups(namespace string) v1alpha1.InstanceGroupInterface {
	return &FakeInstanceGroups{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceGroupManagers(namespace string) v1alpha1.InstanceGroupManagerInterface {
	return &FakeInstanceGroupManagers{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceGroupManagerActions(namespace string) v1alpha1.InstanceGroupManagerActionInterface {
	return &FakeInstanceGroupManagerActions{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceGroupManagerPolicies(namespace string) v1alpha1.InstanceGroupManagerPolicyInterface {
	return &FakeInstanceGroupManagerPolicies{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceGroupMemberships(namespace string) v1alpha1.InstanceGroupMembershipInterface {
	return &FakeInstanceGroupMemberships{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceTemplates(namespace string) v1alpha1.InstanceTemplateInterface {
	return &FakeInstanceTemplates{c, namespace}
}

func (c *FakeIsV1alpha1) InstanceVolumeAttachments(namespace string) v1alpha1.InstanceVolumeAttachmentInterface {
	return &FakeInstanceVolumeAttachments{c, namespace}
}

func (c *FakeIsV1alpha1) IpsecPolicies(namespace string) v1alpha1.IpsecPolicyInterface {
	return &FakeIpsecPolicies{c, namespace}
}

func (c *FakeIsV1alpha1) Lbs(namespace string) v1alpha1.LbInterface {
	return &FakeLbs{c, namespace}
}

func (c *FakeIsV1alpha1) LbListeners(namespace string) v1alpha1.LbListenerInterface {
	return &FakeLbListeners{c, namespace}
}

func (c *FakeIsV1alpha1) LbListenerPolicies(namespace string) v1alpha1.LbListenerPolicyInterface {
	return &FakeLbListenerPolicies{c, namespace}
}

func (c *FakeIsV1alpha1) LbListenerPolicyRules(namespace string) v1alpha1.LbListenerPolicyRuleInterface {
	return &FakeLbListenerPolicyRules{c, namespace}
}

func (c *FakeIsV1alpha1) LbPools(namespace string) v1alpha1.LbPoolInterface {
	return &FakeLbPools{c, namespace}
}

func (c *FakeIsV1alpha1) LbPoolMembers(namespace string) v1alpha1.LbPoolMemberInterface {
	return &FakeLbPoolMembers{c, namespace}
}

func (c *FakeIsV1alpha1) NetworkACLs(namespace string) v1alpha1.NetworkACLInterface {
	return &FakeNetworkACLs{c, namespace}
}

func (c *FakeIsV1alpha1) NetworkACLRules(namespace string) v1alpha1.NetworkACLRuleInterface {
	return &FakeNetworkACLRules{c, namespace}
}

func (c *FakeIsV1alpha1) PublicGateways(namespace string) v1alpha1.PublicGatewayInterface {
	return &FakePublicGateways{c, namespace}
}

func (c *FakeIsV1alpha1) SecurityGroups(namespace string) v1alpha1.SecurityGroupInterface {
	return &FakeSecurityGroups{c, namespace}
}

func (c *FakeIsV1alpha1) SecurityGroupNetworkInterfaceAttachments(namespace string) v1alpha1.SecurityGroupNetworkInterfaceAttachmentInterface {
	return &FakeSecurityGroupNetworkInterfaceAttachments{c, namespace}
}

func (c *FakeIsV1alpha1) SecurityGroupRules(namespace string) v1alpha1.SecurityGroupRuleInterface {
	return &FakeSecurityGroupRules{c, namespace}
}

func (c *FakeIsV1alpha1) SecurityGroupTargets(namespace string) v1alpha1.SecurityGroupTargetInterface {
	return &FakeSecurityGroupTargets{c, namespace}
}

func (c *FakeIsV1alpha1) Snapshots(namespace string) v1alpha1.SnapshotInterface {
	return &FakeSnapshots{c, namespace}
}

func (c *FakeIsV1alpha1) SshKeys(namespace string) v1alpha1.SshKeyInterface {
	return &FakeSshKeys{c, namespace}
}

func (c *FakeIsV1alpha1) Subnets(namespace string) v1alpha1.SubnetInterface {
	return &FakeSubnets{c, namespace}
}

func (c *FakeIsV1alpha1) SubnetNetworkACLAttachments(namespace string) v1alpha1.SubnetNetworkACLAttachmentInterface {
	return &FakeSubnetNetworkACLAttachments{c, namespace}
}

func (c *FakeIsV1alpha1) SubnetReservedIPs(namespace string) v1alpha1.SubnetReservedIPInterface {
	return &FakeSubnetReservedIPs{c, namespace}
}

func (c *FakeIsV1alpha1) VirtualEndpointGateways(namespace string) v1alpha1.VirtualEndpointGatewayInterface {
	return &FakeVirtualEndpointGateways{c, namespace}
}

func (c *FakeIsV1alpha1) VirtualEndpointGatewayIPs(namespace string) v1alpha1.VirtualEndpointGatewayIPInterface {
	return &FakeVirtualEndpointGatewayIPs{c, namespace}
}

func (c *FakeIsV1alpha1) Volumes(namespace string) v1alpha1.VolumeInterface {
	return &FakeVolumes{c, namespace}
}

func (c *FakeIsV1alpha1) Vpcs(namespace string) v1alpha1.VpcInterface {
	return &FakeVpcs{c, namespace}
}

func (c *FakeIsV1alpha1) VpcAddressPrefixes(namespace string) v1alpha1.VpcAddressPrefixInterface {
	return &FakeVpcAddressPrefixes{c, namespace}
}

func (c *FakeIsV1alpha1) VpcRoutes(namespace string) v1alpha1.VpcRouteInterface {
	return &FakeVpcRoutes{c, namespace}
}

func (c *FakeIsV1alpha1) VpcRoutingTables(namespace string) v1alpha1.VpcRoutingTableInterface {
	return &FakeVpcRoutingTables{c, namespace}
}

func (c *FakeIsV1alpha1) VpcRoutingTableRoutes(namespace string) v1alpha1.VpcRoutingTableRouteInterface {
	return &FakeVpcRoutingTableRoutes{c, namespace}
}

func (c *FakeIsV1alpha1) VpnGateways(namespace string) v1alpha1.VpnGatewayInterface {
	return &FakeVpnGateways{c, namespace}
}

func (c *FakeIsV1alpha1) VpnGatewayConnections(namespace string) v1alpha1.VpnGatewayConnectionInterface {
	return &FakeVpnGatewayConnections{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
