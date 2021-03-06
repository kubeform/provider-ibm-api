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

type VpcRoutingTableRoute struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcRoutingTableRouteSpec   `json:"spec,omitempty"`
	Status            VpcRoutingTableRouteStatus `json:"status,omitempty"`
}

type VpcRoutingTableRouteSpec struct {
	State *VpcRoutingTableRouteSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpcRoutingTableRouteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VpcRoutingTableRouteSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The action to perform with a packet matching the route.
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// Routing table route Created At
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The destination of the route.
	Destination *string `json:"destination" tf:"destination"`
	// Routing table route Href
	// +optional
	Href *string `json:"href,omitempty" tf:"href"`
	// Routing table route Lifecycle State
	// +optional
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state"`
	// The user-defined name for this route.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// If action is deliver, the next hop that packets will be delivered to. For other action values, its address will be 0.0.0.0.
	NextHop *string `json:"nextHop" tf:"next_hop"`
	// The origin of this route.
	// +optional
	Origin *string `json:"origin,omitempty" tf:"origin"`
	// The routing table route identifier.
	// +optional
	RouteID *string `json:"routeID,omitempty" tf:"route_id"`
	// The routing table identifier.
	RoutingTable *string `json:"routingTable" tf:"routing_table"`
	// The VPC identifier.
	Vpc *string `json:"vpc" tf:"vpc"`
	// The zone to apply the route to. Traffic from subnets in this zone will be subject to this route.
	Zone *string `json:"zone" tf:"zone"`
}

type VpcRoutingTableRouteStatus struct {
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

// VpcRoutingTableRouteList is a list of VpcRoutingTableRoutes
type VpcRoutingTableRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcRoutingTableRoute CRD objects
	Items []VpcRoutingTableRoute `json:"items,omitempty"`
}
