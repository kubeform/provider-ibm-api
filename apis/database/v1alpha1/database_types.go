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

type Database struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec,omitempty"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

type DatabaseSpecAutoScalingCpu struct {
	// Auto Scaling Rate: Increase Percent
	// +optional
	RateIncreasePercent *int64 `json:"rateIncreasePercent,omitempty" tf:"rate_increase_percent"`
	// Auto Scaling Rate: Limit count per number
	// +optional
	RateLimitCountPerMember *int64 `json:"rateLimitCountPerMember,omitempty" tf:"rate_limit_count_per_member"`
	// Auto Scaling Rate: Period Seconds
	// +optional
	RatePeriodSeconds *int64 `json:"ratePeriodSeconds,omitempty" tf:"rate_period_seconds"`
	// Auto Scaling Rate: Units
	// +optional
	RateUnits *string `json:"rateUnits,omitempty" tf:"rate_units"`
}

type DatabaseSpecAutoScalingDisk struct {
	// Auto Scaling Scalar: Capacity Enabled
	// +optional
	CapacityEnabled *bool `json:"capacityEnabled,omitempty" tf:"capacity_enabled"`
	// Auto Scaling Scalar: Capacity Free Space Less Than Percent
	// +optional
	FreeSpaceLessThanPercent *int64 `json:"freeSpaceLessThanPercent,omitempty" tf:"free_space_less_than_percent"`
	// Auto Scaling Scalar: IO Utilization Above Percent
	// +optional
	IoAbovePercent *int64 `json:"ioAbovePercent,omitempty" tf:"io_above_percent"`
	// Auto Scaling Scalar: IO Utilization Enabled
	// +optional
	IoEnabled *bool `json:"ioEnabled,omitempty" tf:"io_enabled"`
	// Auto Scaling Scalar: IO Utilization Over Period
	// +optional
	IoOverPeriod *string `json:"ioOverPeriod,omitempty" tf:"io_over_period"`
	// Auto Scaling Rate: Increase Percent
	// +optional
	RateIncreasePercent *int64 `json:"rateIncreasePercent,omitempty" tf:"rate_increase_percent"`
	// Auto Scaling Rate: Limit mb per member
	// +optional
	RateLimitMbPerMember *int64 `json:"rateLimitMbPerMember,omitempty" tf:"rate_limit_mb_per_member"`
	// Auto Scaling Rate: Period Seconds
	// +optional
	RatePeriodSeconds *int64 `json:"ratePeriodSeconds,omitempty" tf:"rate_period_seconds"`
	// Auto Scaling Rate: Units
	// +optional
	RateUnits *string `json:"rateUnits,omitempty" tf:"rate_units"`
}

type DatabaseSpecAutoScalingMemory struct {
	// Auto Scaling Scalar: IO Utilization Above Percent
	// +optional
	IoAbovePercent *int64 `json:"ioAbovePercent,omitempty" tf:"io_above_percent"`
	// Auto Scaling Scalar: IO Utilization Enabled
	// +optional
	IoEnabled *bool `json:"ioEnabled,omitempty" tf:"io_enabled"`
	// Auto Scaling Scalar: IO Utilization Over Period
	// +optional
	IoOverPeriod *string `json:"ioOverPeriod,omitempty" tf:"io_over_period"`
	// Auto Scaling Rate: Increase Percent
	// +optional
	RateIncreasePercent *int64 `json:"rateIncreasePercent,omitempty" tf:"rate_increase_percent"`
	// Auto Scaling Rate: Limit mb per member
	// +optional
	RateLimitMbPerMember *int64 `json:"rateLimitMbPerMember,omitempty" tf:"rate_limit_mb_per_member"`
	// Auto Scaling Rate: Period Seconds
	// +optional
	RatePeriodSeconds *int64 `json:"ratePeriodSeconds,omitempty" tf:"rate_period_seconds"`
	// Auto Scaling Rate: Units
	// +optional
	RateUnits *string `json:"rateUnits,omitempty" tf:"rate_units"`
}

type DatabaseSpecAutoScaling struct {
	// CPU Auto Scaling
	// +optional
	Cpu *DatabaseSpecAutoScalingCpu `json:"cpu,omitempty" tf:"cpu"`
	// Disk Auto Scaling
	// +optional
	Disk *DatabaseSpecAutoScalingDisk `json:"disk,omitempty" tf:"disk"`
	// Memory Auto Scaling
	// +optional
	Memory *DatabaseSpecAutoScalingMemory `json:"memory,omitempty" tf:"memory"`
}

type DatabaseSpecConnectionstringsHosts struct {
	// DB host name
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// DB port
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
}

type DatabaseSpecConnectionstrings struct {
	// Certificate in base64 encoding
	// +optional
	Certbase64 *string `json:"certbase64,omitempty" tf:"certbase64"`
	// Certificate Name
	// +optional
	Certname *string `json:"certname,omitempty" tf:"certname"`
	// Connection string
	// +optional
	Composed *string `json:"composed,omitempty" tf:"composed"`
	// DB name
	// +optional
	Database *string `json:"database,omitempty" tf:"database"`
	// +optional
	Hosts []DatabaseSpecConnectionstringsHosts `json:"hosts,omitempty" tf:"hosts"`
	// User name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Password
	// +optional
	Password *string `json:"password,omitempty" tf:"password"`
	// DB path
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// DB query options
	// +optional
	Queryoptions *string `json:"queryoptions,omitempty" tf:"queryoptions"`
	// DB scheme
	// +optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type DatabaseSpecGroupsCpu struct {
	// The current cpu allocation count
	// +optional
	AllocationCount *int64 `json:"allocationCount,omitempty" tf:"allocation_count"`
	// Can the number of CPUs be scaled down as well as up
	// +optional
	CanScaleDown *bool `json:"canScaleDown,omitempty" tf:"can_scale_down"`
	// Are the number of CPUs adjustable
	// +optional
	IsAdjustable *bool `json:"isAdjustable,omitempty" tf:"is_adjustable"`
	// The minimum number of cpus allowed
	// +optional
	MinimumCount *int64 `json:"minimumCount,omitempty" tf:"minimum_count"`
	// The number of CPUs allowed to step up or down by
	// +optional
	StepSizeCount *int64 `json:"stepSizeCount,omitempty" tf:"step_size_count"`
	// The .
	// +optional
	Units *string `json:"units,omitempty" tf:"units"`
}

type DatabaseSpecGroupsDisk struct {
	// The current disk allocation
	// +optional
	AllocationMb *int64 `json:"allocationMb,omitempty" tf:"allocation_mb"`
	// Can the disk size be scaled down as well as up
	// +optional
	CanScaleDown *bool `json:"canScaleDown,omitempty" tf:"can_scale_down"`
	// Is the disk size adjustable
	// +optional
	IsAdjustable *bool `json:"isAdjustable,omitempty" tf:"is_adjustable"`
	// The minimum disk size allowed
	// +optional
	MinimumMb *int64 `json:"minimumMb,omitempty" tf:"minimum_mb"`
	// The step size disk increases or decreases in
	// +optional
	StepSizeMb *int64 `json:"stepSizeMb,omitempty" tf:"step_size_mb"`
	// The units disk is allocated in
	// +optional
	Units *string `json:"units,omitempty" tf:"units"`
}

type DatabaseSpecGroupsMemory struct {
	// The current memory allocation for a group instance
	// +optional
	AllocationMb *int64 `json:"allocationMb,omitempty" tf:"allocation_mb"`
	// Can memory scale down as well as up.
	// +optional
	CanScaleDown *bool `json:"canScaleDown,omitempty" tf:"can_scale_down"`
	// Is the memory size adjustable.
	// +optional
	IsAdjustable *bool `json:"isAdjustable,omitempty" tf:"is_adjustable"`
	// The minimum memory size for a group instance
	// +optional
	MinimumMb *int64 `json:"minimumMb,omitempty" tf:"minimum_mb"`
	// The step size memory increases or decreases in.
	// +optional
	StepSizeMb *int64 `json:"stepSizeMb,omitempty" tf:"step_size_mb"`
	// The units memory is allocated in.
	// +optional
	Units *string `json:"units,omitempty" tf:"units"`
}

type DatabaseSpecGroups struct {
	// Count of scaling groups for the instance
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
	// +optional
	Cpu []DatabaseSpecGroupsCpu `json:"cpu,omitempty" tf:"cpu"`
	// +optional
	Disk []DatabaseSpecGroupsDisk `json:"disk,omitempty" tf:"disk"`
	// Scaling group name
	// +optional
	GroupID *string `json:"groupID,omitempty" tf:"group_id"`
	// +optional
	Memory []DatabaseSpecGroupsMemory `json:"memory,omitempty" tf:"memory"`
}

type DatabaseSpecUsers struct {
	// User name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// User password
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
}

type DatabaseSpecWhitelist struct {
	// Whitelist IP address in CIDR notation
	// +optional
	Address *string `json:"address,omitempty" tf:"address"`
	// Unique white list description
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
}

type DatabaseSpec struct {
	State *DatabaseSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatabaseSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DatabaseSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The admin user password for the instance
	// +optional
	Adminpassword *string `json:"-" sensitive:"true" tf:"adminpassword"`
	// The admin user id for the instance
	// +optional
	Adminuser *string `json:"adminuser,omitempty" tf:"adminuser"`
	// ICD Auto Scaling
	// +optional
	AutoScaling *DatabaseSpecAutoScaling `json:"autoScaling,omitempty" tf:"auto_scaling"`
	// The Backup Encryption Key CRN
	// +optional
	BackupEncryptionKeyCrn *string `json:"backupEncryptionKeyCrn,omitempty" tf:"backup_encryption_key_crn"`
	// The CRN of backup source database
	// +optional
	BackupID *string `json:"backupID,omitempty" tf:"backup_id"`
	// +optional
	Connectionstrings []DatabaseSpecConnectionstrings `json:"connectionstrings,omitempty" tf:"connectionstrings"`
	// +optional
	Groups []DatabaseSpecGroups `json:"groups,omitempty" tf:"groups"`
	// Unique identifier of resource instance
	// +optional
	Guid *string `json:"guid,omitempty" tf:"guid"`
	// The CRN of Key protect instance
	// +optional
	KeyProtectInstance *string `json:"keyProtectInstance,omitempty" tf:"key_protect_instance"`
	// The CRN of Key protect key
	// +optional
	KeyProtectKey *string `json:"keyProtectKey,omitempty" tf:"key_protect_key"`
	// The location or the region in which Database instance exists
	Location *string `json:"location" tf:"location"`
	// CPU allocation required for cluster
	// +optional
	MembersCPUAllocationCount *int64 `json:"membersCPUAllocationCount,omitempty" tf:"members_cpu_allocation_count"`
	// Disk allocation required for cluster
	// +optional
	MembersDiskAllocationMb *int64 `json:"membersDiskAllocationMb,omitempty" tf:"members_disk_allocation_mb"`
	// Memory allocation required for cluster
	// +optional
	MembersMemoryAllocationMb *int64 `json:"membersMemoryAllocationMb,omitempty" tf:"members_memory_allocation_mb"`
	// Resource instance name for example, my Database instance
	Name *string `json:"name" tf:"name"`
	// Total number of nodes in the cluster
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// CPU allocation per node
	// +optional
	NodeCPUAllocationCount *int64 `json:"nodeCPUAllocationCount,omitempty" tf:"node_cpu_allocation_count"`
	// Disk allocation per node
	// +optional
	NodeDiskAllocationMb *int64 `json:"nodeDiskAllocationMb,omitempty" tf:"node_disk_allocation_mb"`
	// Memory allocation per node
	// +optional
	NodeMemoryAllocationMb *int64 `json:"nodeMemoryAllocationMb,omitempty" tf:"node_memory_allocation_mb"`
	// The plan type of the Database instance
	Plan *string `json:"plan" tf:"plan"`
	// For elasticsearch and postgres perform database parameter validation during the plan phase. Otherwise, database parameter validation happens in apply phase.
	// +optional
	PlanValidation *bool `json:"planValidation,omitempty" tf:"plan_validation"`
	// The CRN of source instance
	// +optional
	PointInTimeRecoveryDeploymentID *string `json:"pointInTimeRecoveryDeploymentID,omitempty" tf:"point_in_time_recovery_deployment_id"`
	// The point in time recovery time stamp of the deployed instance
	// +optional
	PointInTimeRecoveryTime *string `json:"pointInTimeRecoveryTime,omitempty" tf:"point_in_time_recovery_time"`
	// The CRN of leader database
	// +optional
	RemoteLeaderID *string `json:"remoteLeaderID,omitempty" tf:"remote_leader_id"`
	// The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource
	// +optional
	ResourceControllerURL *string `json:"resourceControllerURL,omitempty" tf:"resource_controller_url"`
	// The??crn??of??the??resource
	// +optional
	ResourceCrn *string `json:"resourceCrn,omitempty" tf:"resource_crn"`
	// The id of the resource group in which the Database instance is present
	// +optional
	ResourceGroupID *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	// The??resource??group??name??in??which??resource??is??provisioned
	// +optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`
	// The??name??of??the??resource
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
	// The??status??of??the??resource
	// +optional
	ResourceStatus *string `json:"resourceStatus,omitempty" tf:"resource_status"`
	// The name of the Cloud Internet database service
	Service *string `json:"service" tf:"service"`
	// Types of the service endpoints. Possible values are 'public', 'private', 'public-and-private'.
	// +optional
	ServiceEndpoints *string `json:"serviceEndpoints,omitempty" tf:"service_endpoints"`
	// The resource instance status
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Users []DatabaseSpecUsers `json:"users,omitempty" tf:"users"`
	// The database version to provision if specified
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	Whitelist []DatabaseSpecWhitelist `json:"whitelist,omitempty" tf:"whitelist"`
}

type DatabaseStatus struct {
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

// DatabaseList is a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Database CRD objects
	Items []Database `json:"items,omitempty"`
}
