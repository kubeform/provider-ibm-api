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

type ServerInstanceAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerInstanceAttachmentSpec   `json:"spec,omitempty"`
	Status            ServerInstanceAttachmentStatus `json:"status,omitempty"`
}

type ServerInstanceAttachmentSpec struct {
	State *ServerInstanceAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServerInstanceAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServerInstanceAttachmentSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The UUID of a load balancer
	LbaasID *string `json:"lbaasID" tf:"lbaas_id"`
	// The Private IP address of a load balancer member.
	PrivateIPAddress *string `json:"privateIPAddress" tf:"private_ip_address"`
	// The UUID of a load balancer member
	// +optional
	Uuid *string `json:"uuid,omitempty" tf:"uuid"`
	// The weight of a load balancer member.
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type ServerInstanceAttachmentStatus struct {
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

// ServerInstanceAttachmentList is a list of ServerInstanceAttachments
type ServerInstanceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServerInstanceAttachment CRD objects
	Items []ServerInstanceAttachment `json:"items,omitempty"`
}
