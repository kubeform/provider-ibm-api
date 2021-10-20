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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type VmInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VmInstanceSpec   `json:"spec,omitempty"`
	Status            VmInstanceStatus `json:"status,omitempty"`
}

type VmInstanceSpecBulkVms struct {
	Domain   *string `json:"domain" tf:"domain"`
	Hostname *string `json:"hostname" tf:"hostname"`
}

type VmInstanceSpec struct {
	State *VmInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource VmInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VmInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BlockStorageIDS []int64 `json:"blockStorageIDS,omitempty" tf:"block_storage_ids"`
	// +optional
	// +kubebuilder:validation:MinItems=2
	BulkVms []VmInstanceSpecBulkVms `json:"bulkVms,omitempty" tf:"bulk_vms"`
	// +optional
	Cores *int64 `json:"cores,omitempty" tf:"cores"`
	// +optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter"`
	// The user provided datacenter options
	// +optional
	// +optional
	DedicatedAcctHostOnly *bool `json:"dedicatedAcctHostOnly,omitempty" tf:"dedicated_acct_host_only"`
	// +optional
	DedicatedHostID *int64 `json:"dedicatedHostID,omitempty" tf:"dedicated_host_id"`
	// +optional
	DedicatedHostName *string `json:"dedicatedHostName,omitempty" tf:"dedicated_host_name"`
	// +optional
	Disks []int64 `json:"disks,omitempty" tf:"disks"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	Evault *int64 `json:"evault,omitempty" tf:"evault"`
	// +optional
	FileStorageIDS []int64 `json:"fileStorageIDS,omitempty" tf:"file_storage_ids"`
	// Flavor key name used to provision vm.
	// +optional
	FlavorKeyName *string `json:"flavorKeyName,omitempty" tf:"flavor_key_name"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	HourlyBilling *bool `json:"hourlyBilling,omitempty" tf:"hourly_billing"`
	// +optional
	ImageID *int64 `json:"imageID,omitempty" tf:"image_id"`
	// +optional
	IpAddressID *int64 `json:"ipAddressID,omitempty" tf:"ip_address_id"`
	// +optional
	IpAddressIDPrivate *int64 `json:"ipAddressIDPrivate,omitempty" tf:"ip_address_id_private"`
	// +optional
	Ipv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address"`
	// +optional
	Ipv4AddressPrivate *string `json:"ipv4AddressPrivate,omitempty" tf:"ipv4_address_private"`
	// +optional
	Ipv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address"`
	// +optional
	Ipv6AddressID *int64 `json:"ipv6AddressID,omitempty" tf:"ipv6_address_id"`
	// +optional
	Ipv6Enabled *bool `json:"ipv6Enabled,omitempty" tf:"ipv6_enabled"`
	// +optional
	Ipv6StaticEnabled *bool `json:"ipv6StaticEnabled,omitempty" tf:"ipv6_static_enabled"`
	// +optional
	LocalDisk *bool `json:"localDisk,omitempty" tf:"local_disk"`
	// +optional
	Memory *int64 `json:"memory,omitempty" tf:"memory"`
	// +optional
	NetworkSpeed *int64 `json:"networkSpeed,omitempty" tf:"network_speed"`
	// +optional
	Notes *string `json:"notes,omitempty" tf:"notes"`
	// +optional
	OsReferenceCode *string `json:"osReferenceCode,omitempty" tf:"os_reference_code"`
	// The placement group id
	// +optional
	PlacementGroupID *int64 `json:"placementGroupID,omitempty" tf:"placement_group_id"`
	// The placement group name
	// +optional
	PlacementGroupName *string `json:"placementGroupName,omitempty" tf:"placement_group_name"`
	// +optional
	PostInstallScriptURI *string `json:"postInstallScriptURI,omitempty" tf:"post_install_script_uri"`
	// +optional
	PrivateInterfaceID *int64 `json:"privateInterfaceID,omitempty" tf:"private_interface_id"`
	// +optional
	PrivateNetworkOnly *bool `json:"privateNetworkOnly,omitempty" tf:"private_network_only"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	PrivateSecurityGroupIDS []int64 `json:"privateSecurityGroupIDS,omitempty" tf:"private_security_group_ids"`
	// +optional
	PrivateSubnet *string `json:"privateSubnet,omitempty" tf:"private_subnet"`
	// +optional
	PrivateSubnetID *int64 `json:"privateSubnetID,omitempty" tf:"private_subnet_id"`
	// +optional
	PrivateVLANID *int64 `json:"privateVLANID,omitempty" tf:"private_vlan_id"`
	// +optional
	PublicBandwidthLimited *int64 `json:"publicBandwidthLimited,omitempty" tf:"public_bandwidth_limited"`
	// +optional
	PublicBandwidthUnlimited *bool `json:"publicBandwidthUnlimited,omitempty" tf:"public_bandwidth_unlimited"`
	// +optional
	PublicInterfaceID *int64 `json:"publicInterfaceID,omitempty" tf:"public_interface_id"`
	// +optional
	PublicIpv6Subnet *string `json:"publicIpv6Subnet,omitempty" tf:"public_ipv6_subnet"`
	// +optional
	PublicIpv6SubnetID *int64 `json:"publicIpv6SubnetID,omitempty" tf:"public_ipv6_subnet_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	PublicSecurityGroupIDS []int64 `json:"publicSecurityGroupIDS,omitempty" tf:"public_security_group_ids"`
	// +optional
	PublicSubnet *string `json:"publicSubnet,omitempty" tf:"public_subnet"`
	// +optional
	PublicSubnetID *int64 `json:"publicSubnetID,omitempty" tf:"public_subnet_id"`
	// +optional
	PublicVLANID *int64 `json:"publicVLANID,omitempty" tf:"public_vlan_id"`
	// Quote ID for Quote based provisioning
	// +optional
	QuoteID *int64 `json:"quoteID,omitempty" tf:"quote_id"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// The status of the resource
	// +optional
	ResourceStatus *string `json:"resourceStatus,omitempty" tf:"resource_status"`
	// +optional
	SecondaryIPAddresses []string `json:"secondaryIPAddresses,omitempty" tf:"secondary_ip_addresses"`
	// +optional
	SecondaryIPCount *int64 `json:"secondaryIPCount,omitempty" tf:"secondary_ip_count"`
	// +optional
	SshKeyIDS []int64 `json:"sshKeyIDS,omitempty" tf:"ssh_key_ids"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Transient *bool `json:"transient,omitempty" tf:"transient"`
	// +optional
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata"`
	// +optional
	// Deprecated
	WaitTimeMinutes *int64 `json:"waitTimeMinutes,omitempty" tf:"wait_time_minutes"`
}

type VmInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VmInstanceList is a list of VmInstances
type VmInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VmInstance CRD objects
	Items []VmInstance `json:"items,omitempty"`
}
