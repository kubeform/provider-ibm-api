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

type CertificateUpload struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateUploadSpec   `json:"spec,omitempty"`
	Status            CertificateUploadStatus `json:"status,omitempty"`
}

type CertificateUploadSpec struct {
	State *CertificateUploadSpecResource `json:"state,omitempty" tf:"-"`

	Resource CertificateUploadSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CertificateUploadSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Certificate bundle method
	// +optional
	BundleMethod *string `json:"bundleMethod,omitempty" tf:"bundle_method"`
	// Certificate key
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// CIS instance crn
	CisID *string `json:"cisID" tf:"cis_id"`
	// +optional
	CustomCertID *string `json:"customCertID,omitempty" tf:"custom_cert_id"`
	// Associated CIS domain
	DomainID *string `json:"domainID" tf:"domain_id"`
	// certificate expires date
	// +optional
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on"`
	// hosts which the certificate uploaded to
	// +optional
	Hosts []string `json:"hosts,omitempty" tf:"hosts"`
	// certificate issuer
	// +optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer"`
	// certificate modified date
	// +optional
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on"`
	// Certificate priority
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// Certificate private key
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// certificate signature
	// +optional
	Signature *string `json:"signature,omitempty" tf:"signature"`
	// certificate status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// certificate uploaded date
	// +optional
	UploadedOn *string `json:"uploadedOn,omitempty" tf:"uploaded_on"`
}

type CertificateUploadStatus struct {
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

// CertificateUploadList is a list of CertificateUploads
type CertificateUploadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CertificateUpload CRD objects
	Items []CertificateUpload `json:"items,omitempty"`
}
