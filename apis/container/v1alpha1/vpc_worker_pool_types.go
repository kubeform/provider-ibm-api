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

type VpcWorkerPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcWorkerPoolSpec   `json:"spec,omitempty"`
	Status            VpcWorkerPoolStatus `json:"status,omitempty"`
}

type VpcWorkerPoolSpecTaints struct {
	// Effect for taint. Accepted values are NoSchedule, PreferNoSchedule and NoExecute.
	Effect *string `json:"effect" tf:"effect"`
	// Key for taint
	Key *string `json:"key" tf:"key"`
	// Value for taint.
	Value *string `json:"value" tf:"value"`
}

type VpcWorkerPoolSpecZones struct {
	// zone name
	Name *string `json:"name" tf:"name"`
	// subnet ID
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
}

type VpcWorkerPoolSpec struct {
	State *VpcWorkerPoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpcWorkerPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VpcWorkerPoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Cluster name
	Cluster *string `json:"cluster" tf:"cluster"`
	// Entitlement option reduces additional OCP Licence cost in Openshift Clusters
	// +optional
	Entitlement *string `json:"entitlement,omitempty" tf:"entitlement"`
	// cluster node falvor
	Flavor *string `json:"flavor" tf:"flavor"`
	// Labels
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Resource Controller URL
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// ID of the resource group.
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// WorkerPool Taints
	// +optional
	Taints []VpcWorkerPoolSpecTaints `json:"taints,omitempty" tf:"taints"`
	// The vpc id where the cluster is
	VpcID *string `json:"vpcID" tf:"vpc_id"`
	// The number of workers
	WorkerCount *int64 `json:"workerCount" tf:"worker_count"`
	// worker pool name
	WorkerPoolName *string `json:"workerPoolName" tf:"worker_pool_name"`
	// Zones info
	Zones []VpcWorkerPoolSpecZones `json:"zones" tf:"zones"`
}

type VpcWorkerPoolStatus struct {
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

// VpcWorkerPoolList is a list of VpcWorkerPools
type VpcWorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcWorkerPool CRD objects
	Items []VpcWorkerPool `json:"items,omitempty"`
}
