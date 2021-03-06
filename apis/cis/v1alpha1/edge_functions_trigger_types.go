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

type EdgeFunctionsTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeFunctionsTriggerSpec   `json:"spec,omitempty"`
	Status            EdgeFunctionsTriggerStatus `json:"status,omitempty"`
}

type EdgeFunctionsTriggerSpec struct {
	State *EdgeFunctionsTriggerSpecResource `json:"state,omitempty" tf:"-"`

	Resource EdgeFunctionsTriggerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EdgeFunctionsTriggerSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Edge function trigger action name
	// +optional
	ActionName *string `json:"actionName,omitempty" tf:"action_name"`
	// CIS Intance CRN
	CisID *string `json:"cisID" tf:"cis_id"`
	// CIS Domain ID
	DomainID *string `json:"domainID" tf:"domain_id"`
	// Edge function trigger pattern
	PatternURL *string `json:"patternURL" tf:"pattern_url"`
	// Edge function trigger request limit fail open
	// +optional
	RequestLimitFailOpen *bool `json:"requestLimitFailOpen,omitempty" tf:"request_limit_fail_open"`
	// CIS Edge Functions trigger route ID
	// +optional
	TriggerID *string `json:"triggerID,omitempty" tf:"trigger_id"`
}

type EdgeFunctionsTriggerStatus struct {
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

// EdgeFunctionsTriggerList is a list of EdgeFunctionsTriggers
type EdgeFunctionsTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EdgeFunctionsTrigger CRD objects
	Items []EdgeFunctionsTrigger `json:"items,omitempty"`
}
