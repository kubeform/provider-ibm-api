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

type StreamsTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamsTopicSpec   `json:"spec,omitempty"`
	Status            StreamsTopicStatus `json:"status,omitempty"`
}

type StreamsTopicSpec struct {
	State *StreamsTopicSpecResource `json:"state,omitempty" tf:"-"`

	Resource StreamsTopicSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type StreamsTopicSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The configuration parameters of a topic
	// +optional
	Config map[string]string `json:"config,omitempty" tf:"config"`
	// Kafka brokers addresses for interacting with Kafka native API
	// +optional
	KafkaBrokersSasl []string `json:"kafkaBrokersSasl,omitempty" tf:"kafka_brokers_sasl"`
	// API endpoint for interacting with Event Streams REST API
	// +optional
	KafkaHTTPURL *string `json:"kafkaHTTPURL,omitempty" tf:"kafka_http_url"`
	// The name of the topic
	Name *string `json:"name" tf:"name"`
	// The number of partitions
	// +optional
	Partitions *int64 `json:"partitions,omitempty" tf:"partitions"`
	// The CRN of the Event Streams instance
	ResourceInstanceID *string `json:"resourceInstanceID" tf:"resource_instance_id"`
}

type StreamsTopicStatus struct {
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

// StreamsTopicList is a list of StreamsTopics
type StreamsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamsTopic CRD objects
	Items []StreamsTopic `json:"items,omitempty"`
}
