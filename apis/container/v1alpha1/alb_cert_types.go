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

type AlbCert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbCertSpec   `json:"spec,omitempty"`
	Status            AlbCertStatus `json:"status,omitempty"`
}

type AlbCertSpec struct {
	State *AlbCertSpecResource `json:"state,omitempty" tf:"-"`

	Resource AlbCertSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AlbCertSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Certificate CRN id
	CertCrn *string `json:"certCrn" tf:"cert_crn"`
	// cloud cert instance ID
	// +optional
	CloudCertInstanceID *string `json:"cloudCertInstanceID,omitempty" tf:"cloud_cert_instance_id"`
	// Cluster ID
	ClusterID *string `json:"clusterID" tf:"cluster_id"`
	// Domain name
	// +optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name"`
	// Certificate expaire on date
	// +optional
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on"`
	// certificate issuer name
	// +optional
	// Deprecated
	IssuerName *string `json:"issuerName,omitempty" tf:"issuer_name"`
	// Namespace of the secret
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// Persistence of secret
	// +optional
	Persistence *bool `json:"persistence,omitempty" tf:"persistence"`
	// region name
	// +optional
	// Deprecated
	Region *string `json:"region,omitempty" tf:"region"`
	// Secret name
	SecretName *string `json:"secretName" tf:"secret_name"`
	// Secret Status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type AlbCertStatus struct {
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

// AlbCertList is a list of AlbCerts
type AlbCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlbCert CRD objects
	Items []AlbCert `json:"items,omitempty"`
}
