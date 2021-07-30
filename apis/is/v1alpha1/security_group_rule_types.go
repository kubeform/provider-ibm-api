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

type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupRuleSpec   `json:"spec,omitempty"`
	Status            SecurityGroupRuleStatus `json:"status,omitempty"`
}

type SecurityGroupRuleSpecIcmp struct {
	// +optional
	Code *int64 `json:"code,omitempty" tf:"code"`
	// +optional
	Type *int64 `json:"type,omitempty" tf:"type"`
}

type SecurityGroupRuleSpecTcp struct {
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
}

type SecurityGroupRuleSpecUdp struct {
	// +optional
	PortMax *int64 `json:"portMax,omitempty" tf:"port_max"`
	// +optional
	PortMin *int64 `json:"portMin,omitempty" tf:"port_min"`
}

type SecurityGroupRuleSpec struct {
	State *SecurityGroupRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource SecurityGroupRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SecurityGroupRuleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Direction of traffic to enforce, either inbound or outbound
	Direction *string `json:"direction" tf:"direction"`
	// Security group id
	Group *string `json:"group" tf:"group"`
	// protocol=icmp
	// +optional
	Icmp *SecurityGroupRuleSpecIcmp `json:"icmp,omitempty" tf:"icmp"`
	// IP version: ipv4
	// +optional
	IpVersion *string `json:"ipVersion,omitempty" tf:"ip_version"`
	// The Security Group Rule Protocol
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// The crn of the Security Group
	// +optional
	RelatedCrn *string `json:"relatedCrn,omitempty" tf:"related_crn"`
	// Security group id: an IP address, a CIDR block, or a single security group identifier
	// +optional
	Remote *string `json:"remote,omitempty" tf:"remote"`
	// Rule id
	// +optional
	RuleID *string `json:"ruleID,omitempty" tf:"rule_id"`
	// protocol=tcp
	// +optional
	Tcp *SecurityGroupRuleSpecTcp `json:"tcp,omitempty" tf:"tcp"`
	// protocol=udp
	// +optional
	Udp *SecurityGroupRuleSpecUdp `json:"udp,omitempty" tf:"udp"`
}

type SecurityGroupRuleStatus struct {
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

// SecurityGroupRuleList is a list of SecurityGroupRules
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityGroupRule CRD objects
	Items []SecurityGroupRule `json:"items,omitempty"`
}
