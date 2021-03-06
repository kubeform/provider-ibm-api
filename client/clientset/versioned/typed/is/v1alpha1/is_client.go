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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-ibm-api/apis/is/v1alpha1"
	"kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type IsV1alpha1Interface interface {
	RESTClient() rest.Interface
	DedicatedHostsGetter
	DedicatedHostDiskManagementsGetter
	DedicatedHostGroupsGetter
	FloatingIPsGetter
	FlowLogsGetter
	IkePoliciesGetter
	ImagesGetter
	InstancesGetter
	InstanceDiskManagementsGetter
	InstanceGroupsGetter
	InstanceGroupManagersGetter
	InstanceGroupManagerActionsGetter
	InstanceGroupManagerPoliciesGetter
	InstanceGroupMembershipsGetter
	InstanceTemplatesGetter
	InstanceVolumeAttachmentsGetter
	IpsecPoliciesGetter
	LbsGetter
	LbListenersGetter
	LbListenerPoliciesGetter
	LbListenerPolicyRulesGetter
	LbPoolsGetter
	LbPoolMembersGetter
	NetworkACLsGetter
	NetworkACLRulesGetter
	PublicGatewaysGetter
	SecurityGroupsGetter
	SecurityGroupNetworkInterfaceAttachmentsGetter
	SecurityGroupRulesGetter
	SecurityGroupTargetsGetter
	SnapshotsGetter
	SshKeysGetter
	SubnetsGetter
	SubnetNetworkACLAttachmentsGetter
	SubnetReservedIPsGetter
	VirtualEndpointGatewaysGetter
	VirtualEndpointGatewayIPsGetter
	VolumesGetter
	VpcsGetter
	VpcAddressPrefixesGetter
	VpcRoutesGetter
	VpcRoutingTablesGetter
	VpcRoutingTableRoutesGetter
	VpnGatewaysGetter
	VpnGatewayConnectionsGetter
}

// IsV1alpha1Client is used to interact with features provided by the is.ibm.kubeform.com group.
type IsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *IsV1alpha1Client) DedicatedHosts(namespace string) DedicatedHostInterface {
	return newDedicatedHosts(c, namespace)
}

func (c *IsV1alpha1Client) DedicatedHostDiskManagements(namespace string) DedicatedHostDiskManagementInterface {
	return newDedicatedHostDiskManagements(c, namespace)
}

func (c *IsV1alpha1Client) DedicatedHostGroups(namespace string) DedicatedHostGroupInterface {
	return newDedicatedHostGroups(c, namespace)
}

func (c *IsV1alpha1Client) FloatingIPs(namespace string) FloatingIPInterface {
	return newFloatingIPs(c, namespace)
}

func (c *IsV1alpha1Client) FlowLogs(namespace string) FlowLogInterface {
	return newFlowLogs(c, namespace)
}

func (c *IsV1alpha1Client) IkePolicies(namespace string) IkePolicyInterface {
	return newIkePolicies(c, namespace)
}

func (c *IsV1alpha1Client) Images(namespace string) ImageInterface {
	return newImages(c, namespace)
}

func (c *IsV1alpha1Client) Instances(namespace string) InstanceInterface {
	return newInstances(c, namespace)
}

func (c *IsV1alpha1Client) InstanceDiskManagements(namespace string) InstanceDiskManagementInterface {
	return newInstanceDiskManagements(c, namespace)
}

func (c *IsV1alpha1Client) InstanceGroups(namespace string) InstanceGroupInterface {
	return newInstanceGroups(c, namespace)
}

func (c *IsV1alpha1Client) InstanceGroupManagers(namespace string) InstanceGroupManagerInterface {
	return newInstanceGroupManagers(c, namespace)
}

func (c *IsV1alpha1Client) InstanceGroupManagerActions(namespace string) InstanceGroupManagerActionInterface {
	return newInstanceGroupManagerActions(c, namespace)
}

func (c *IsV1alpha1Client) InstanceGroupManagerPolicies(namespace string) InstanceGroupManagerPolicyInterface {
	return newInstanceGroupManagerPolicies(c, namespace)
}

func (c *IsV1alpha1Client) InstanceGroupMemberships(namespace string) InstanceGroupMembershipInterface {
	return newInstanceGroupMemberships(c, namespace)
}

func (c *IsV1alpha1Client) InstanceTemplates(namespace string) InstanceTemplateInterface {
	return newInstanceTemplates(c, namespace)
}

func (c *IsV1alpha1Client) InstanceVolumeAttachments(namespace string) InstanceVolumeAttachmentInterface {
	return newInstanceVolumeAttachments(c, namespace)
}

func (c *IsV1alpha1Client) IpsecPolicies(namespace string) IpsecPolicyInterface {
	return newIpsecPolicies(c, namespace)
}

func (c *IsV1alpha1Client) Lbs(namespace string) LbInterface {
	return newLbs(c, namespace)
}

func (c *IsV1alpha1Client) LbListeners(namespace string) LbListenerInterface {
	return newLbListeners(c, namespace)
}

func (c *IsV1alpha1Client) LbListenerPolicies(namespace string) LbListenerPolicyInterface {
	return newLbListenerPolicies(c, namespace)
}

func (c *IsV1alpha1Client) LbListenerPolicyRules(namespace string) LbListenerPolicyRuleInterface {
	return newLbListenerPolicyRules(c, namespace)
}

func (c *IsV1alpha1Client) LbPools(namespace string) LbPoolInterface {
	return newLbPools(c, namespace)
}

func (c *IsV1alpha1Client) LbPoolMembers(namespace string) LbPoolMemberInterface {
	return newLbPoolMembers(c, namespace)
}

func (c *IsV1alpha1Client) NetworkACLs(namespace string) NetworkACLInterface {
	return newNetworkACLs(c, namespace)
}

func (c *IsV1alpha1Client) NetworkACLRules(namespace string) NetworkACLRuleInterface {
	return newNetworkACLRules(c, namespace)
}

func (c *IsV1alpha1Client) PublicGateways(namespace string) PublicGatewayInterface {
	return newPublicGateways(c, namespace)
}

func (c *IsV1alpha1Client) SecurityGroups(namespace string) SecurityGroupInterface {
	return newSecurityGroups(c, namespace)
}

func (c *IsV1alpha1Client) SecurityGroupNetworkInterfaceAttachments(namespace string) SecurityGroupNetworkInterfaceAttachmentInterface {
	return newSecurityGroupNetworkInterfaceAttachments(c, namespace)
}

func (c *IsV1alpha1Client) SecurityGroupRules(namespace string) SecurityGroupRuleInterface {
	return newSecurityGroupRules(c, namespace)
}

func (c *IsV1alpha1Client) SecurityGroupTargets(namespace string) SecurityGroupTargetInterface {
	return newSecurityGroupTargets(c, namespace)
}

func (c *IsV1alpha1Client) Snapshots(namespace string) SnapshotInterface {
	return newSnapshots(c, namespace)
}

func (c *IsV1alpha1Client) SshKeys(namespace string) SshKeyInterface {
	return newSshKeys(c, namespace)
}

func (c *IsV1alpha1Client) Subnets(namespace string) SubnetInterface {
	return newSubnets(c, namespace)
}

func (c *IsV1alpha1Client) SubnetNetworkACLAttachments(namespace string) SubnetNetworkACLAttachmentInterface {
	return newSubnetNetworkACLAttachments(c, namespace)
}

func (c *IsV1alpha1Client) SubnetReservedIPs(namespace string) SubnetReservedIPInterface {
	return newSubnetReservedIPs(c, namespace)
}

func (c *IsV1alpha1Client) VirtualEndpointGateways(namespace string) VirtualEndpointGatewayInterface {
	return newVirtualEndpointGateways(c, namespace)
}

func (c *IsV1alpha1Client) VirtualEndpointGatewayIPs(namespace string) VirtualEndpointGatewayIPInterface {
	return newVirtualEndpointGatewayIPs(c, namespace)
}

func (c *IsV1alpha1Client) Volumes(namespace string) VolumeInterface {
	return newVolumes(c, namespace)
}

func (c *IsV1alpha1Client) Vpcs(namespace string) VpcInterface {
	return newVpcs(c, namespace)
}

func (c *IsV1alpha1Client) VpcAddressPrefixes(namespace string) VpcAddressPrefixInterface {
	return newVpcAddressPrefixes(c, namespace)
}

func (c *IsV1alpha1Client) VpcRoutes(namespace string) VpcRouteInterface {
	return newVpcRoutes(c, namespace)
}

func (c *IsV1alpha1Client) VpcRoutingTables(namespace string) VpcRoutingTableInterface {
	return newVpcRoutingTables(c, namespace)
}

func (c *IsV1alpha1Client) VpcRoutingTableRoutes(namespace string) VpcRoutingTableRouteInterface {
	return newVpcRoutingTableRoutes(c, namespace)
}

func (c *IsV1alpha1Client) VpnGateways(namespace string) VpnGatewayInterface {
	return newVpnGateways(c, namespace)
}

func (c *IsV1alpha1Client) VpnGatewayConnections(namespace string) VpnGatewayConnectionInterface {
	return newVpnGatewayConnections(c, namespace)
}

// NewForConfig creates a new IsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*IsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new IsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *IsV1alpha1Client {
	return &IsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *IsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
