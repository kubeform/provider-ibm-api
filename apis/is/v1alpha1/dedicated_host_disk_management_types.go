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

type DedicatedHostDiskManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHostDiskManagementSpec   `json:"spec,omitempty"`
	Status            DedicatedHostDiskManagementStatus `json:"status,omitempty"`
}

type DedicatedHostDiskManagementSpecDisks struct {
	// The unique identifier for this disk.
	ID *string `json:"ID" tf:"id"`
	// The user-defined name for this disk. The disk will be updated with this new name
	Name *string `json:"name" tf:"name"`
}

type DedicatedHostDiskManagementSpec struct {
	State *DedicatedHostDiskManagementSpecResource `json:"state,omitempty" tf:"-"`

	Resource DedicatedHostDiskManagementSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DedicatedHostDiskManagementSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the dedicated host for which disks has to be managed
	DedicatedHost *string `json:"dedicatedHost" tf:"dedicated_host"`
	// Disk information that has to be updated.
	Disks []DedicatedHostDiskManagementSpecDisks `json:"disks" tf:"disks"`
}

type DedicatedHostDiskManagementStatus struct {
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

// DedicatedHostDiskManagementList is a list of DedicatedHostDiskManagements
type DedicatedHostDiskManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DedicatedHostDiskManagement CRD objects
	Items []DedicatedHostDiskManagement `json:"items,omitempty"`
}