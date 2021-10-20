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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecPlanHistory struct {
	// +optional
	ResourcePlanID *string `json:"resourcePlanID,omitempty" tf:"resource_plan_id"`
	// +optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// An alpha-numeric value identifying the account ID.
	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
	// +optional
	AllowCleanup *bool `json:"allowCleanup,omitempty" tf:"allow_cleanup"`
	// The date when the instance was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The subject who created the instance.
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// CRN of resource instance
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// Dashboard URL to access resource.
	// +optional
	DashboardURL *string `json:"dashboardURL,omitempty" tf:"dashboard_url"`
	// The date when the instance was deleted.
	// +optional
	DeletedAt *string `json:"deletedAt,omitempty" tf:"deleted_at"`
	// The subject who deleted the instance.
	// +optional
	DeletedBy *string `json:"deletedBy,omitempty" tf:"deleted_by"`
	// The extended metadata as a map associated with the resource instance.
	// +optional
	Extensions map[string]string `json:"extensions,omitempty" tf:"extensions"`
	// Guid of resource instance
	// +optional
	Guid *string `json:"guid,omitempty" tf:"guid"`
	// The status of the last operation requested on the instance
	// +optional
	LastOperation map[string]string `json:"lastOperation,omitempty" tf:"last_operation"`
	// The location where the instance available
	Location *string `json:"location" tf:"location"`
	// A boolean that dictates if the resource instance should be deleted (cleaned up) during the processing of a region instance delete call.
	// +optional
	Locked *bool `json:"locked,omitempty" tf:"locked"`
	// A name for the resource instance
	Name *string `json:"name" tf:"name"`
	// Arbitrary parameters to pass. Must be a JSON object
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// The plan type of the service
	Plan *string `json:"plan" tf:"plan"`
	// The plan history of the instance.
	// +optional
	PlanHistory []InstanceSpecPlanHistory `json:"planHistory,omitempty" tf:"plan_history"`
	// The relative path to the resource aliases for the instance.
	// +optional
	ResourceAliasesURL *string `json:"resourceAliasesURL,omitempty" tf:"resource_aliases_url"`
	// The relative path to the resource bindings for the instance.
	// +optional
	ResourceBindingsURL *string `json:"resourceBindingsURL,omitempty" tf:"resource_bindings_url"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// The long ID (full CRN) of the resource group
	// +optional
	ResourceGroupCrn *string `json:"resourceGroupCrn,omitempty" tf:"resource_group_crn"`
	// The resource group id
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// The resource group name in which resource is provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The unique ID of the offering
	// +optional
	ResourceID *string `json:"resourceID,omitempty" tf:"resource_id"`
	// The relative path to the resource keys for the instance.
	// +optional
	ResourceKeysURL *string `json:"resourceKeysURL,omitempty" tf:"resource_keys_url"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// The unique ID of the plan associated with the offering
	// +optional
	ResourcePlanID *string `json:"resourcePlanID,omitempty" tf:"resource_plan_id"`
	// The status of the resource
	// +optional
	ResourceStatus *string `json:"resourceStatus,omitempty" tf:"resource_status"`
	// The date when the instance under reclamation was restored.
	// +optional
	RestoredAt *string `json:"restoredAt,omitempty" tf:"restored_at"`
	// The subject who restored the instance back from reclamation.
	// +optional
	RestoredBy *string `json:"restoredBy,omitempty" tf:"restored_by"`
	// The date when the instance was scheduled for reclamation.
	// +optional
	ScheduledReclaimAt *string `json:"scheduledReclaimAt,omitempty" tf:"scheduled_reclaim_at"`
	// The subject who initiated the instance reclamation.
	// +optional
	ScheduledReclaimBy *string `json:"scheduledReclaimBy,omitempty" tf:"scheduled_reclaim_by"`
	// The name of the service offering like cloud-object-storage, kms etc
	Service *string `json:"service" tf:"service"`
	// Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'.
	// +optional
	ServiceEndpoints *string `json:"serviceEndpoints,omitempty" tf:"service_endpoints"`
	// The current state of the instance.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Status of resource instance
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// The sub-type of instance, e.g. cfaas .
	// +optional
	SubType *string `json:"subType,omitempty" tf:"sub_type"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The full deployment CRN as defined in the global catalog
	// +optional
	TargetCrn *string `json:"targetCrn,omitempty" tf:"target_crn"`
	// The type of the instance, e.g. service_instance.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The date when the instance was last updated.
	// +optional
	UpdateAt *string `json:"updateAt,omitempty" tf:"update_at"`
	// The subject who updated the instance.
	// +optional
	UpdateBy *string `json:"updateBy,omitempty" tf:"update_by"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
