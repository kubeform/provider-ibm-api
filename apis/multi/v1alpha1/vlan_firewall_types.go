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

type VlanFirewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VlanFirewallSpec   `json:"spec,omitempty"`
	Status            VlanFirewallStatus `json:"status,omitempty"`
}

type VlanFirewallSpec struct {
	State *VlanFirewallSpecResource `json:"state,omitempty" tf:"-"`

	Resource VlanFirewallSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type VlanFirewallSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// High Availability - [Web Filtering Add-on, NGFW Add-on, AV Add-on] or [Web Filtering Add-on, NGFW Add-on, AV Add-on]
	// +optional
	AddonConfiguration []string `json:"addonConfiguration,omitempty" tf:"addon_configuration"`
	// Datacenter name
	Datacenter *string `json:"datacenter" tf:"datacenter"`
	// Firewall type
	FirewallType *string `json:"firewallType" tf:"firewall_type"`
	// name
	Name *string `json:"name" tf:"name"`
	// Password
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// POD name
	Pod *string `json:"pod" tf:"pod"`
	// Private IP Address
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// Private VLAN id
	// +optional
	PrivateVLANID *int64 `json:"privateVLANID,omitempty" tf:"private_vlan_id"`
	// Public IP Address
	// +optional
	PublicIP *string `json:"publicIP,omitempty" tf:"public_ip"`
	// Public IPV6 IP
	// +optional
	PublicIpv6 *string `json:"publicIpv6,omitempty" tf:"public_ipv6"`
	// Public VLAN id
	// +optional
	PublicVLANID *int64 `json:"publicVLANID,omitempty" tf:"public_vlan_id"`
	// User name
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type VlanFirewallStatus struct {
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

// VlanFirewallList is a list of VlanFirewalls
type VlanFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VlanFirewall CRD objects
	Items []VlanFirewall `json:"items,omitempty"`
}