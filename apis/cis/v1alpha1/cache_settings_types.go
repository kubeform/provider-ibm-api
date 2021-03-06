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

type CacheSettings struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CacheSettingsSpec   `json:"spec,omitempty"`
	Status            CacheSettingsStatus `json:"status,omitempty"`
}

type CacheSettingsSpec struct {
	State *CacheSettingsSpecResource `json:"state,omitempty" tf:"-"`

	Resource CacheSettingsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CacheSettingsSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Browser Expiration setting
	// +optional
	BrowserExpiration *int64 `json:"browserExpiration,omitempty" tf:"browser_expiration"`
	// Cache level setting
	// +optional
	CachingLevel *string `json:"cachingLevel,omitempty" tf:"caching_level"`
	// CIS instance crn
	CisID *string `json:"cisID" tf:"cis_id"`
	// Development mode setting
	// +optional
	DevelopmentMode *string `json:"developmentMode,omitempty" tf:"development_mode"`
	// Associated CIS domain
	DomainID *string `json:"domainID" tf:"domain_id"`
	// Purge all setting
	// +optional
	PurgeAll *bool `json:"purgeAll,omitempty" tf:"purge_all"`
	// Purge by hosts
	// +optional
	PurgeByHosts []string `json:"purgeByHosts,omitempty" tf:"purge_by_hosts"`
	// Purge by tags
	// +optional
	PurgeByTags []string `json:"purgeByTags,omitempty" tf:"purge_by_tags"`
	// Purge by URLs
	// +optional
	PurgeByUrls []string `json:"purgeByUrls,omitempty" tf:"purge_by_urls"`
	// Query String sort setting
	// +optional
	QueryStringSort *string `json:"queryStringSort,omitempty" tf:"query_string_sort"`
	// Serve Stale Content
	// +optional
	ServeStaleContent *string `json:"serveStaleContent,omitempty" tf:"serve_stale_content"`
}

type CacheSettingsStatus struct {
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

// CacheSettingsList is a list of CacheSettingss
type CacheSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CacheSettings CRD objects
	Items []CacheSettings `json:"items,omitempty"`
}
