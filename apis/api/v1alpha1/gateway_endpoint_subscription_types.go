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

type GatewayEndpointSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayEndpointSubscriptionSpec   `json:"spec,omitempty"`
	Status            GatewayEndpointSubscriptionStatus `json:"status,omitempty"`
}

type GatewayEndpointSubscriptionSpec struct {
	State *GatewayEndpointSubscriptionSpecResource `json:"state,omitempty" tf:"-"`

	Resource GatewayEndpointSubscriptionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type GatewayEndpointSubscriptionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Endpoint ID
	ArtifactID *string `json:"artifactID" tf:"artifact_id"`
	// Subscription Id, API key that is used to create subscription
	// +optional
	ClientID *string `json:"clientID,omitempty" tf:"client_id"`
	// Client Sercret of a Subscription
	// +optional
	ClientSecret *string `json:"-" sensitive:"true" tf:"client_secret"`
	// Indicates if Client Sercret has to be autogenerated
	// +optional
	GenerateSecret *bool `json:"generateSecret,omitempty" tf:"generate_secret"`
	// Subscription name
	Name *string `json:"name" tf:"name"`
	// Indicates if client secret is provided to subscription or not
	// +optional
	SecretProvided *bool `json:"secretProvided,omitempty" tf:"secret_provided"`
	// Subscription type. Allowable values are external, internal
	Type *string `json:"type" tf:"type"`
}

type GatewayEndpointSubscriptionStatus struct {
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

// GatewayEndpointSubscriptionList is a list of GatewayEndpointSubscriptions
type GatewayEndpointSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GatewayEndpointSubscription CRD objects
	Items []GatewayEndpointSubscription `json:"items,omitempty"`
}
