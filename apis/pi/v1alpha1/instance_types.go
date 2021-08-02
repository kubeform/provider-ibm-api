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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecAddresses struct {
	// +optional
	ExternalIP *string `json:"externalIP,omitempty" tf:"external_ip"`
	// +optional
	Ip *string `json:"ip,omitempty" tf:"ip"`
	// +optional
	Macaddress *string `json:"macaddress,omitempty" tf:"macaddress"`
	// +optional
	NetworkID *string `json:"networkID,omitempty" tf:"network_id"`
	// +optional
	NetworkName *string `json:"networkName,omitempty" tf:"network_name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Addresses []InstanceSpecAddresses `json:"addresses,omitempty" tf:"addresses"`
	// PI Instance health status
	// +optional
	HealthStatus *string `json:"healthStatus,omitempty" tf:"health_status"`
	// Instance ID
	// +optional
	InstanceID *string `json:"instanceID,omitempty" tf:"instance_id"`
	// Maximum memory size
	// +optional
	MaxMemory *float64 `json:"maxMemory,omitempty" tf:"max_memory"`
	// Maximum number of processors
	// +optional
	MaxProcessors *float64 `json:"maxProcessors,omitempty" tf:"max_processors"`
	// Maximum Virtual Cores Assigned to the PVMInstance
	// +optional
	MaxVirtualCores *int64 `json:"maxVirtualCores,omitempty" tf:"max_virtual_cores"`
	// set to true to enable migration of the PI instance
	// +optional
	// Deprecated
	Migratable *bool `json:"migratable,omitempty" tf:"migratable"`
	// Minimum memory
	// +optional
	MinMemory *float64 `json:"minMemory,omitempty" tf:"min_memory"`
	// Minimum number of the CPUs
	// +optional
	MinProcessors *float64 `json:"minProcessors,omitempty" tf:"min_processors"`
	// Minimum Virtual Cores Assigned to the PVMInstance
	// +optional
	MinVirtualCores *int64 `json:"minVirtualCores,omitempty" tf:"min_virtual_cores"`
	// Operating System
	// +optional
	OperatingSystem *string `json:"operatingSystem,omitempty" tf:"operating_system"`
	// OS Type
	// +optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`
	// This is the Power Instance id that is assigned to the account
	PiCloudInstanceID *string `json:"piCloudInstanceID" tf:"pi_cloud_instance_id"`
	// Allow the user to set the status of the lpar so that they can connect to it faster
	// +optional
	PiHealthStatus *string `json:"piHealthStatus,omitempty" tf:"pi_health_status"`
	// PI instance image name
	PiImageID *string `json:"piImageID" tf:"pi_image_id"`
	// PI Instance name
	PiInstanceName *string `json:"piInstanceName" tf:"pi_instance_name"`
	// SSH key name
	PiKeyPairName *string `json:"piKeyPairName" tf:"pi_key_pair_name"`
	// Memory size
	PiMemory *float64 `json:"piMemory" tf:"pi_memory"`
	// set to true to enable migration of the PI instance
	// +optional
	PiMigratable *bool `json:"piMigratable,omitempty" tf:"pi_migratable"`
	// List of Networks that have been configured for the account
	PiNetworkIDS []string `json:"piNetworkIDS" tf:"pi_network_ids"`
	// Pin Policy of the instance
	// +optional
	PiPinPolicy *string `json:"piPinPolicy,omitempty" tf:"pi_pin_policy"`
	// Instance processor type
	PiProcType *string `json:"piProcType" tf:"pi_proc_type"`
	// Processors count
	PiProcessors *float64 `json:"piProcessors" tf:"pi_processors"`
	// Progress of the operation
	// +optional
	PiProgress *float64 `json:"piProgress,omitempty" tf:"pi_progress"`
	// PI Instance replicas count
	// +optional
	PiReplicants *float64 `json:"piReplicants,omitempty" tf:"pi_replicants"`
	// Replication policy for the PI Instance
	// +optional
	PiReplicationPolicy *string `json:"piReplicationPolicy,omitempty" tf:"pi_replication_policy"`
	// Replication scheme
	// +optional
	PiReplicationScheme *string `json:"piReplicationScheme,omitempty" tf:"pi_replication_scheme"`
	// Storage type for server deployment
	// +optional
	PiStorageType *string `json:"piStorageType,omitempty" tf:"pi_storage_type"`
	// PI Instance system type
	PiSysType *string `json:"piSysType" tf:"pi_sys_type"`
	// Base64 encoded data to be passed in for invoking a cloud init script
	// +optional
	PiUserData *string `json:"piUserData,omitempty" tf:"pi_user_data"`
	// Virtual Cores Assigned to the PVMInstance
	// +optional
	PiVirtualCoresAssigned *int64 `json:"piVirtualCoresAssigned,omitempty" tf:"pi_virtual_cores_assigned"`
	// List of PI volumes
	// +optional
	PiVolumeIDS []string `json:"piVolumeIDS,omitempty" tf:"pi_volume_ids"`
	// PIN Policy of the Instance
	// +optional
	PinPolicy *string `json:"pinPolicy,omitempty" tf:"pin_policy"`
	// PI instance status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}