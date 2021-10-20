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

type IkePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IkePolicySpec   `json:"spec,omitempty"`
	Status            IkePolicyStatus `json:"status,omitempty"`
}

type IkePolicySpecVpnConnections struct {
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type IkePolicySpec struct {
	State *IkePolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource IkePolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type IkePolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Authentication algorithm type
	AuthenticationAlgorithm *string `json:"authenticationAlgorithm" tf:"authentication_algorithm"`
	// IKE DH group
	DhGroup *int64 `json:"dhGroup" tf:"dh_group"`
	// Encryption alogorithm type
	EncryptionAlgorithm *string `json:"encryptionAlgorithm" tf:"encryption_algorithm"`
	// IKE href value
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// IKE version
	// +optional
	IkeVersion *int64 `json:"ikeVersion,omitempty" tf:"ike_version"`
	// IKE Key lifetime
	// +optional
	KeyLifetime *int64 `json:"keyLifetime,omitempty" tf:"key_lifetime"`
	// IKE name
	Name *string `json:"name" tf:"name"`
	// IKE negotiation mode
	// +optional
	NegotiationMode *string `json:"negotiationMode,omitempty" tf:"negotiation_mode"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// IKE resource group ID
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The resource group name in which resource is provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// +optional
	VpnConnections []IkePolicySpecVpnConnections `json:"vpnConnections,omitempty" tf:"vpn_connections"`
}

type IkePolicyStatus struct {
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

// IkePolicyList is a list of IkePolicys
type IkePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IkePolicy CRD objects
	Items []IkePolicy `json:"items,omitempty"`
}
