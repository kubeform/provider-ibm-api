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

type InstanceGroupManagerAction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceGroupManagerActionSpec   `json:"spec,omitempty"`
	Status            InstanceGroupManagerActionStatus `json:"status,omitempty"`
}

type InstanceGroupManagerActionSpec struct {
	State *InstanceGroupManagerActionSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceGroupManagerActionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceGroupManagerActionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance group manager action ID
	// +optional
	ActionID *string `json:"actionID,omitempty" tf:"action_id"`
	// The type of action for the instance group.
	// +optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type"`
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete"`
	// +optional
	AutoDeleteTimeout *int64 `json:"autoDeleteTimeout,omitempty" tf:"auto_delete_timeout"`
	// The date and time that the instance group manager action was modified.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.
	// +optional
	CronSpec *string `json:"cronSpec,omitempty" tf:"cron_spec"`
	// instance group ID
	InstanceGroup *string `json:"instanceGroup" tf:"instance_group"`
	// Instance group manager ID of type scheduled
	InstanceGroupManager *string `json:"instanceGroupManager" tf:"instance_group_manager"`
	// The date and time the scheduled action was last applied. If empty the action has never been applied.
	// +optional
	LastAppliedAt *string `json:"lastAppliedAt,omitempty" tf:"last_applied_at"`
	// The maximum number of members in a managed instance group
	// +optional
	MaxMembershipCount *int64 `json:"maxMembershipCount,omitempty" tf:"max_membership_count"`
	// The number of members the instance group should have at the scheduled time.
	// +optional
	MembershipCount *int64 `json:"membershipCount,omitempty" tf:"membership_count"`
	// The minimum number of members in a managed instance group
	// +optional
	MinMembershipCount *int64 `json:"minMembershipCount,omitempty" tf:"min_membership_count"`
	// instance group manager action name
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The date and time the scheduled action will next run. If empty the system is currently calculating the next run time.
	// +optional
	NextRunAt *string `json:"nextRunAt,omitempty" tf:"next_run_at"`
	// The resource type.
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// The date and time the scheduled action will run.
	// +optional
	RunAt *string `json:"runAt,omitempty" tf:"run_at"`
	// The status of the instance group action- `active`: Action is ready to be run- `completed`: Action was completed successfully- `failed`: Action could not be completed successfully- `incompatible`: Action parameters are not compatible with the group or manager- `omitted`: Action was not applied because this action's manager was disabled.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// The unique identifier for this instance group manager of type autoscale.
	// +optional
	TargetManager *string `json:"targetManager,omitempty" tf:"target_manager"`
	// Instance group manager name of type autoscale.
	// +optional
	TargetManagerName *string `json:"targetManagerName,omitempty" tf:"target_manager_name"`
	// The date and time that the instance group manager action was modified.
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
}

type InstanceGroupManagerActionStatus struct {
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

// InstanceGroupManagerActionList is a list of InstanceGroupManagerActions
type InstanceGroupManagerActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstanceGroupManagerAction CRD objects
	Items []InstanceGroupManagerAction `json:"items,omitempty"`
}
