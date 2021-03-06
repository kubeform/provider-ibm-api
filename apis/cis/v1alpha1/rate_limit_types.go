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

type RateLimit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RateLimitSpec   `json:"spec,omitempty"`
	Status            RateLimitStatus `json:"status,omitempty"`
}

type RateLimitSpecActionResponse struct {
	// The body to return. The content here must confirm to the 'content_type'
	Body *string `json:"body" tf:"body"`
	// Custom content-type and body to return. It must be one of following 'text/plain', 'text/xml', 'application/json'.
	ContentType *string `json:"contentType" tf:"content_type"`
}

type RateLimitSpecAction struct {
	// Type of action performed.Valid values are: 'simulate', 'ban', 'challenge', 'js_challenge'.
	Mode *string `json:"mode" tf:"mode"`
	// Rate Limiting Action Response
	// +optional
	Response *RateLimitSpecActionResponse `json:"response,omitempty" tf:"response"`
	// The time to perform the mitigation action. Timeout be the same or greater than the period.
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
}

type RateLimitSpecBypass struct {
	// bypass URL name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// bypass URL value
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type RateLimitSpecCorrelate struct {
	// Whether to enable NAT based rate limiting
	// +optional
	By *string `json:"by,omitempty" tf:"by"`
}

type RateLimitSpecMatchRequest struct {
	// HTTP Methos of matching request. It can be one or many. Example methods 'POST', 'PUT'
	// +optional
	Methods []string `json:"methods,omitempty" tf:"methods"`
	// HTTP Schemes of matching request. It can be one or many. Example schemes 'HTTP', 'HTTPS'.
	// +optional
	Schemes []string `json:"schemes,omitempty" tf:"schemes"`
	// URL pattern of matching request
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type RateLimitSpecMatchResponseHeaders struct {
	// The name of the response header to match.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The operator when matching. Valid values are 'eq' and 'ne'.
	// +optional
	Op *string `json:"op,omitempty" tf:"op"`
	// The value of the header, which is exactly matched.
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type RateLimitSpecMatchResponse struct {
	// +optional
	Headers []RateLimitSpecMatchResponseHeaders `json:"headers,omitempty" tf:"headers"`
	// Origin Traffic of matching response.
	// +optional
	OriginTraffic *bool `json:"originTraffic,omitempty" tf:"origin_traffic"`
	// HTTP Status Codes of matching response. It can be one or many. Example status codes '403', '401
	// +optional
	Status []int64 `json:"status,omitempty" tf:"status"`
}

type RateLimitSpecMatch struct {
	// Rate Limiting Match Request
	// +optional
	Request *RateLimitSpecMatchRequest `json:"request,omitempty" tf:"request"`
	// Rate Limiting Response
	// +optional
	Response *RateLimitSpecMatchResponse `json:"response,omitempty" tf:"response"`
}

type RateLimitSpec struct {
	State *RateLimitSpecResource `json:"state,omitempty" tf:"-"`

	Resource RateLimitSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RateLimitSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Rate Limiting Action
	Action *RateLimitSpecAction `json:"action" tf:"action"`
	// Bypass URL
	// +optional
	Bypass []RateLimitSpecBypass `json:"bypass,omitempty" tf:"bypass"`
	// CIS Intance CRN
	CisID *string `json:"cisID" tf:"cis_id"`
	// Ratelimiting Correlate
	// +optional
	Correlate *RateLimitSpecCorrelate `json:"correlate,omitempty" tf:"correlate"`
	// A note that you can use to describe the reason for a rate limiting rule.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Whether this rate limiting rule is currently disabled.
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// CIS Domain ID
	DomainID *string `json:"domainID" tf:"domain_id"`
	// Rate Limiting Match
	// +optional
	Match *RateLimitSpecMatch `json:"match,omitempty" tf:"match"`
	// Rate Limiting Period
	Period *int64 `json:"period" tf:"period"`
	// Rate Limit rule Id
	// +optional
	RuleID *string `json:"ruleID,omitempty" tf:"rule_id"`
	// Rate Limiting Threshold
	Threshold *int64 `json:"threshold" tf:"threshold"`
}

type RateLimitStatus struct {
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

// RateLimitList is a list of RateLimits
type RateLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RateLimit CRD objects
	Items []RateLimit `json:"items,omitempty"`
}
