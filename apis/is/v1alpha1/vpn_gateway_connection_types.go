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

type VpnGatewayConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewayConnectionSpec   `json:"spec,omitempty"`
	Status            VpnGatewayConnectionStatus `json:"status,omitempty"`
}

type VpnGatewayConnectionSpecTunnels struct {
	// The IP address of the VPN gateway member in which the tunnel resides
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// The status of the VPN Tunnel
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type VpnGatewayConnectionSpec struct {
	State *VpnGatewayConnectionSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpnGatewayConnectionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VpnGatewayConnectionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Action detection for dead peer detection action
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// VPN gateway connection admin state
	// +optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up"`
	// The authentication mode
	// +optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode"`
	// The date and time that this VPN gateway connection was created
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The unique identifier for this VPN gateway connection
	// +optional
	GatewayConnection *string `json:"gatewayConnection,omitempty" tf:"gateway_connection"`
	// VPN gateway connection IKE Policy
	// +optional
	IkePolicy *string `json:"ikePolicy,omitempty" tf:"ike_policy"`
	// Interval for dead peer detection interval
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
	// IP security policy for vpn gateway connection
	// +optional
	IpsecPolicy *string `json:"ipsecPolicy,omitempty" tf:"ipsec_policy"`
	// VPN gateway connection local CIDRs
	// +optional
	LocalCidrs []string `json:"localCidrs,omitempty" tf:"local_cidrs"`
	// The mode of the VPN gateway
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// VPN Gateway connection name
	Name *string `json:"name" tf:"name"`
	// VPN gateway connection peer address
	PeerAddress *string `json:"peerAddress" tf:"peer_address"`
	// VPN gateway connection peer CIDRs
	// +optional
	PeerCidrs []string `json:"peerCidrs,omitempty" tf:"peer_cidrs"`
	// vpn gateway
	PresharedKey *string `json:"presharedKey" tf:"preshared_key"`
	// The crn of the VPN Gateway resource
	// +optional
	RelatedCrn *string `json:"relatedCrn,omitempty" tf:"related_crn"`
	// The resource type
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// VPN gateway connection status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Timeout for dead peer detection
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// The VPN tunnel configuration for this VPN gateway connection (in static route mode)
	// +optional
	Tunnels []VpnGatewayConnectionSpecTunnels `json:"tunnels,omitempty" tf:"tunnels"`
	// VPN Gateway info
	VpnGateway *string `json:"vpnGateway" tf:"vpn_gateway"`
}

type VpnGatewayConnectionStatus struct {
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

// VpnGatewayConnectionList is a list of VpnGatewayConnections
type VpnGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnGatewayConnection CRD objects
	Items []VpnGatewayConnection `json:"items,omitempty"`
}
