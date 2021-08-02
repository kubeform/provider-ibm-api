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

type LbPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbPoolSpec   `json:"spec,omitempty"`
	Status            LbPoolStatus `json:"status,omitempty"`
}

type LbPoolSpec struct {
	State *LbPoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource LbPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LbPoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Load Balancer Pool algorithm
	Algorithm *string `json:"algorithm" tf:"algorithm"`
	// Load Blancer health delay time period
	HealthDelay *int64 `json:"healthDelay" tf:"health_delay"`
	// Health monitor Port the LB Pool
	// +optional
	HealthMonitorPort *int64 `json:"healthMonitorPort,omitempty" tf:"health_monitor_port"`
	// Health monitor URL of LB Pool
	// +optional
	HealthMonitorURL *string `json:"healthMonitorURL,omitempty" tf:"health_monitor_url"`
	// Load Balancer health retry count
	HealthRetries *int64 `json:"healthRetries" tf:"health_retries"`
	// Load Balancer health timeout interval
	HealthTimeout *int64 `json:"healthTimeout" tf:"health_timeout"`
	// Load Balancer health type
	HealthType *string `json:"healthType" tf:"health_type"`
	// Load Balancer ID
	Lb *string `json:"lb" tf:"lb"`
	// Load Balancer Pool name
	Name *string `json:"name" tf:"name"`
	// The LB Pool id
	// +optional
	PoolID *string `json:"poolID,omitempty" tf:"pool_id"`
	// Load Balancer Protocol
	Protocol *string `json:"protocol" tf:"protocol"`
	// Status of the LB Pool
	// +optional
	ProvisioningStatus *string `json:"provisioningStatus,omitempty" tf:"provisioning_status"`
	// PROXY protocol setting for this pool
	// +optional
	ProxyProtocol *string `json:"proxyProtocol,omitempty" tf:"proxy_protocol"`
	// The crn of the LB resource
	// +optional
	RelatedCrn *string `json:"relatedCrn,omitempty" tf:"related_crn"`
	// Load Balancer Pool session persisence cookie name
	// +optional
	SessionPersistenceCookieName *string `json:"sessionPersistenceCookieName,omitempty" tf:"session_persistence_cookie_name"`
	// Load Balancer Pool session persisence type.
	// +optional
	SessionPersistenceType *string `json:"sessionPersistenceType,omitempty" tf:"session_persistence_type"`
}

type LbPoolStatus struct {
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

// LbPoolList is a list of LbPools
type LbPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbPool CRD objects
	Items []LbPool `json:"items,omitempty"`
}