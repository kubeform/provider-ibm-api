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

type DedicatedHostGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHostGroupSpec   `json:"spec,omitempty"`
	Status            DedicatedHostGroupStatus `json:"status,omitempty"`
}

type DedicatedHostGroupSpecDedicatedHostsDeleted struct {
	// Link to documentation about deleted resources.
	// +optional
	MoreInfo *string `json:"moreInfo,omitempty" tf:"more_info"`
}

type DedicatedHostGroupSpecDedicatedHosts struct {
	// The CRN for this dedicated host.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
	// +optional
	Deleted []DedicatedHostGroupSpecDedicatedHostsDeleted `json:"deleted,omitempty" tf:"deleted"`
	// The URL for this dedicated host.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique identifier for this dedicated host.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The type of resource referenced.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
}

type DedicatedHostGroupSpecSupportedInstanceProfiles struct {
	// The URL for this virtual server instance profile.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The globally unique name for this virtual server instance profile.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DedicatedHostGroupSpec struct {
	State *DedicatedHostGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource DedicatedHostGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DedicatedHostGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The dedicated host profile class for hosts in this group.
	Class *string `json:"class" tf:"class"`
	// The date and time that the dedicated host group was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The CRN for this dedicated host group.
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// The dedicated hosts that are in this dedicated host group.
	// +optional
	DedicatedHosts []DedicatedHostGroupSpecDedicatedHosts `json:"dedicatedHosts,omitempty" tf:"dedicated_hosts"`
	// The dedicated host profile family for hosts in this group.
	Family *string `json:"family" tf:"family"`
	// The URL for this dedicated host group.
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The unique identifier of the resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The type of resource referenced.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// Array of instance profiles that can be used by instances placed on this dedicated host group.
	// +optional
	SupportedInstanceProfiles []DedicatedHostGroupSpecSupportedInstanceProfiles `json:"supportedInstanceProfiles,omitempty" tf:"supported_instance_profiles"`
	// The globally unique name of the zone this dedicated host group will reside in.
	Zone *string `json:"zone" tf:"zone"`
}

type DedicatedHostGroupStatus struct {
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

// DedicatedHostGroupList is a list of DedicatedHostGroups
type DedicatedHostGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DedicatedHostGroup CRD objects
	Items []DedicatedHostGroup `json:"items,omitempty"`
}