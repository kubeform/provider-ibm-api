//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Block) DeepCopyInto(out *Block) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Block.
func (in *Block) DeepCopy() *Block {
	if in == nil {
		return nil
	}
	out := new(Block)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Block) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockList) DeepCopyInto(out *BlockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Block, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockList.
func (in *BlockList) DeepCopy() *BlockList {
	if in == nil {
		return nil
	}
	out := new(BlockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpec) DeepCopyInto(out *BlockSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BlockSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpec.
func (in *BlockSpec) DeepCopy() *BlockSpec {
	if in == nil {
		return nil
	}
	out := new(BlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpecAllowedHardwareInfo) DeepCopyInto(out *BlockSpecAllowedHardwareInfo) {
	*out = *in
	if in.HostIqn != nil {
		in, out := &in.HostIqn, &out.HostIqn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpecAllowedHardwareInfo.
func (in *BlockSpecAllowedHardwareInfo) DeepCopy() *BlockSpecAllowedHardwareInfo {
	if in == nil {
		return nil
	}
	out := new(BlockSpecAllowedHardwareInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpecAllowedHostInfo) DeepCopyInto(out *BlockSpecAllowedHostInfo) {
	*out = *in
	if in.HostIqn != nil {
		in, out := &in.HostIqn, &out.HostIqn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpecAllowedHostInfo.
func (in *BlockSpecAllowedHostInfo) DeepCopy() *BlockSpecAllowedHostInfo {
	if in == nil {
		return nil
	}
	out := new(BlockSpecAllowedHostInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpecAllowedVirtualGuestInfo) DeepCopyInto(out *BlockSpecAllowedVirtualGuestInfo) {
	*out = *in
	if in.HostIqn != nil {
		in, out := &in.HostIqn, &out.HostIqn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpecAllowedVirtualGuestInfo.
func (in *BlockSpecAllowedVirtualGuestInfo) DeepCopy() *BlockSpecAllowedVirtualGuestInfo {
	if in == nil {
		return nil
	}
	out := new(BlockSpecAllowedVirtualGuestInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpecResource) DeepCopyInto(out *BlockSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedHardwareIDS != nil {
		in, out := &in.AllowedHardwareIDS, &out.AllowedHardwareIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.AllowedHardwareInfo != nil {
		in, out := &in.AllowedHardwareInfo, &out.AllowedHardwareInfo
		*out = make([]BlockSpecAllowedHardwareInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedHostInfo != nil {
		in, out := &in.AllowedHostInfo, &out.AllowedHostInfo
		*out = make([]BlockSpecAllowedHostInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedIPAddresses != nil {
		in, out := &in.AllowedIPAddresses, &out.AllowedIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedVirtualGuestIDS != nil {
		in, out := &in.AllowedVirtualGuestIDS, &out.AllowedVirtualGuestIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.AllowedVirtualGuestInfo != nil {
		in, out := &in.AllowedVirtualGuestInfo, &out.AllowedVirtualGuestInfo
		*out = make([]BlockSpecAllowedVirtualGuestInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.HourlyBilling != nil {
		in, out := &in.HourlyBilling, &out.HourlyBilling
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.Lunid != nil {
		in, out := &in.Lunid, &out.Lunid
		*out = new(string)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.OsFormatType != nil {
		in, out := &in.OsFormatType, &out.OsFormatType
		*out = new(string)
		**out = **in
	}
	if in.ResourceControllerURL != nil {
		in, out := &in.ResourceControllerURL, &out.ResourceControllerURL
		*out = new(string)
		**out = **in
	}
	if in.ResourceName != nil {
		in, out := &in.ResourceName, &out.ResourceName
		*out = new(string)
		**out = **in
	}
	if in.SnapshotCapacity != nil {
		in, out := &in.SnapshotCapacity, &out.SnapshotCapacity
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TargetAddress != nil {
		in, out := &in.TargetAddress, &out.TargetAddress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Volumename != nil {
		in, out := &in.Volumename, &out.Volumename
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpecResource.
func (in *BlockSpecResource) DeepCopy() *BlockSpecResource {
	if in == nil {
		return nil
	}
	out := new(BlockSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockStatus) DeepCopyInto(out *BlockStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockStatus.
func (in *BlockStatus) DeepCopy() *BlockStatus {
	if in == nil {
		return nil
	}
	out := new(BlockStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Evault) DeepCopyInto(out *Evault) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Evault.
func (in *Evault) DeepCopy() *Evault {
	if in == nil {
		return nil
	}
	out := new(Evault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Evault) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaultList) DeepCopyInto(out *EvaultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Evault, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaultList.
func (in *EvaultList) DeepCopy() *EvaultList {
	if in == nil {
		return nil
	}
	out := new(EvaultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EvaultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaultSpec) DeepCopyInto(out *EvaultSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(EvaultSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaultSpec.
func (in *EvaultSpec) DeepCopy() *EvaultSpec {
	if in == nil {
		return nil
	}
	out := new(EvaultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaultSpecResource) DeepCopyInto(out *EvaultSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.HardwareInstanceID != nil {
		in, out := &in.HardwareInstanceID, &out.HardwareInstanceID
		*out = new(int64)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.ServiceResourceName != nil {
		in, out := &in.ServiceResourceName, &out.ServiceResourceName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VirtualInstanceID != nil {
		in, out := &in.VirtualInstanceID, &out.VirtualInstanceID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaultSpecResource.
func (in *EvaultSpecResource) DeepCopy() *EvaultSpecResource {
	if in == nil {
		return nil
	}
	out := new(EvaultSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaultStatus) DeepCopyInto(out *EvaultStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaultStatus.
func (in *EvaultStatus) DeepCopy() *EvaultStatus {
	if in == nil {
		return nil
	}
	out := new(EvaultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *File) DeepCopyInto(out *File) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new File.
func (in *File) DeepCopy() *File {
	if in == nil {
		return nil
	}
	out := new(File)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *File) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileList) DeepCopyInto(out *FileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]File, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileList.
func (in *FileList) DeepCopy() *FileList {
	if in == nil {
		return nil
	}
	out := new(FileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpec) DeepCopyInto(out *FileSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FileSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpec.
func (in *FileSpec) DeepCopy() *FileSpec {
	if in == nil {
		return nil
	}
	out := new(FileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpecResource) DeepCopyInto(out *FileSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedHardwareIDS != nil {
		in, out := &in.AllowedHardwareIDS, &out.AllowedHardwareIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.AllowedIPAddresses != nil {
		in, out := &in.AllowedIPAddresses, &out.AllowedIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedSubnets != nil {
		in, out := &in.AllowedSubnets, &out.AllowedSubnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedVirtualGuestIDS != nil {
		in, out := &in.AllowedVirtualGuestIDS, &out.AllowedVirtualGuestIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(int64)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.HourlyBilling != nil {
		in, out := &in.HourlyBilling, &out.HourlyBilling
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(float64)
		**out = **in
	}
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = new(string)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.ResourceControllerURL != nil {
		in, out := &in.ResourceControllerURL, &out.ResourceControllerURL
		*out = new(string)
		**out = **in
	}
	if in.ResourceName != nil {
		in, out := &in.ResourceName, &out.ResourceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceStatus != nil {
		in, out := &in.ResourceStatus, &out.ResourceStatus
		*out = new(string)
		**out = **in
	}
	if in.SnapshotCapacity != nil {
		in, out := &in.SnapshotCapacity, &out.SnapshotCapacity
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotSchedule != nil {
		in, out := &in.SnapshotSchedule, &out.SnapshotSchedule
		*out = make([]FileSpecSnapshotSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Volumename != nil {
		in, out := &in.Volumename, &out.Volumename
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpecResource.
func (in *FileSpecResource) DeepCopy() *FileSpecResource {
	if in == nil {
		return nil
	}
	out := new(FileSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSpecSnapshotSchedule) DeepCopyInto(out *FileSpecSnapshotSchedule) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.Hour != nil {
		in, out := &in.Hour, &out.Hour
		*out = new(int64)
		**out = **in
	}
	if in.Minute != nil {
		in, out := &in.Minute, &out.Minute
		*out = new(int64)
		**out = **in
	}
	if in.RetentionCount != nil {
		in, out := &in.RetentionCount, &out.RetentionCount
		*out = new(int64)
		**out = **in
	}
	if in.ScheduleType != nil {
		in, out := &in.ScheduleType, &out.ScheduleType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSpecSnapshotSchedule.
func (in *FileSpecSnapshotSchedule) DeepCopy() *FileSpecSnapshotSchedule {
	if in == nil {
		return nil
	}
	out := new(FileSpecSnapshotSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileStatus) DeepCopyInto(out *FileStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileStatus.
func (in *FileStatus) DeepCopy() *FileStatus {
	if in == nil {
		return nil
	}
	out := new(FileStatus)
	in.DeepCopyInto(out)
	return out
}
