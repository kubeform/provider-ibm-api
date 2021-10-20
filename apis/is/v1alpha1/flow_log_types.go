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

type FlowLog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec,omitempty"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

type FlowLogSpec struct {
	State *FlowLogSpecResource `json:"state,omitempty" tf:"-"`

	Resource FlowLogSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FlowLogSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Indicates whether this collector is active
	// +optional
	Active *bool `json:"active,omitempty" tf:"active"`
	// If set to true, this flow log collector will be automatically deleted when the target is deleted
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete"`
	// The date and time flow log was created
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The CRN for this flow log collector
	// +optional
	Crn *string `json:"crn,omitempty" tf:"crn"`
	// The URL for this flow log collector
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// The lifecycle state of the flow log collector
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state"`
	// Flow Log Collector name
	Name *string `json:"name" tf:"name"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The crn of the resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// The resource group of flow log
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// The resource group name in which resource is provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The name of the resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// The status of the resource
	// +optional
	ResourceStatus *string `json:"resourceStatus,omitempty" tf:"resource_status"`
	// The Cloud Object Storage bucket name where the collected flows will be logged
	StorageBucket *string `json:"storageBucket" tf:"storage_bucket"`
	// Tags for the VPC Flow logs
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The target id that the flow log collector is to collect flow logs
	Target *string `json:"target" tf:"target"`
	// The VPC this flow log collector is associated with
	// +optional
	Vpc *string `json:"vpc,omitempty" tf:"vpc"`
}

type FlowLogStatus struct {
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

// FlowLogList is a list of FlowLogs
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FlowLog CRD objects
	Items []FlowLog `json:"items,omitempty"`
}
