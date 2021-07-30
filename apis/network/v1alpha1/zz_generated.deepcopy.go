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

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(GatewaySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpecAssociatedVlans) DeepCopyInto(out *GatewaySpecAssociatedVlans) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(bool)
		**out = **in
	}
	if in.NetworkVLANID != nil {
		in, out := &in.NetworkVLANID, &out.NetworkVLANID
		*out = new(int64)
		**out = **in
	}
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpecAssociatedVlans.
func (in *GatewaySpecAssociatedVlans) DeepCopy() *GatewaySpecAssociatedVlans {
	if in == nil {
		return nil
	}
	out := new(GatewaySpecAssociatedVlans)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpecMembers) DeepCopyInto(out *GatewaySpecMembers) {
	*out = *in
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.DiskKeyNames != nil {
		in, out := &in.DiskKeyNames, &out.DiskKeyNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Ipv6Address != nil {
		in, out := &in.Ipv6Address, &out.Ipv6Address
		*out = new(string)
		**out = **in
	}
	if in.Ipv6Enabled != nil {
		in, out := &in.Ipv6Enabled, &out.Ipv6Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MemberID != nil {
		in, out := &in.MemberID, &out.MemberID
		*out = new(int64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(int64)
		**out = **in
	}
	if in.NetworkSpeed != nil {
		in, out := &in.NetworkSpeed, &out.NetworkSpeed
		*out = new(int64)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.OsKeyName != nil {
		in, out := &in.OsKeyName, &out.OsKeyName
		*out = new(string)
		**out = **in
	}
	if in.PackageKeyName != nil {
		in, out := &in.PackageKeyName, &out.PackageKeyName
		*out = new(string)
		**out = **in
	}
	if in.PostInstallScriptURI != nil {
		in, out := &in.PostInstallScriptURI, &out.PostInstallScriptURI
		*out = new(string)
		**out = **in
	}
	if in.PrivateIpv4Address != nil {
		in, out := &in.PrivateIpv4Address, &out.PrivateIpv4Address
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworkOnly != nil {
		in, out := &in.PrivateNetworkOnly, &out.PrivateNetworkOnly
		*out = new(bool)
		**out = **in
	}
	if in.PrivateVLANID != nil {
		in, out := &in.PrivateVLANID, &out.PrivateVLANID
		*out = new(int64)
		**out = **in
	}
	if in.ProcessKeyName != nil {
		in, out := &in.ProcessKeyName, &out.ProcessKeyName
		*out = new(string)
		**out = **in
	}
	if in.PublicBandwidth != nil {
		in, out := &in.PublicBandwidth, &out.PublicBandwidth
		*out = new(int64)
		**out = **in
	}
	if in.PublicIpv4Address != nil {
		in, out := &in.PublicIpv4Address, &out.PublicIpv4Address
		*out = new(string)
		**out = **in
	}
	if in.PublicVLANID != nil {
		in, out := &in.PublicVLANID, &out.PublicVLANID
		*out = new(int64)
		**out = **in
	}
	if in.RedundantNetwork != nil {
		in, out := &in.RedundantNetwork, &out.RedundantNetwork
		*out = new(bool)
		**out = **in
	}
	if in.RedundantPowerSupply != nil {
		in, out := &in.RedundantPowerSupply, &out.RedundantPowerSupply
		*out = new(bool)
		**out = **in
	}
	if in.SshKeyIDS != nil {
		in, out := &in.SshKeyIDS, &out.SshKeyIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.StorageGroups != nil {
		in, out := &in.StorageGroups, &out.StorageGroups
		*out = make([]GatewaySpecMembersStorageGroups, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TcpMonitoring != nil {
		in, out := &in.TcpMonitoring, &out.TcpMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.UnbondedNetwork != nil {
		in, out := &in.UnbondedNetwork, &out.UnbondedNetwork
		*out = new(bool)
		**out = **in
	}
	if in.UserMetadata != nil {
		in, out := &in.UserMetadata, &out.UserMetadata
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpecMembers.
func (in *GatewaySpecMembers) DeepCopy() *GatewaySpecMembers {
	if in == nil {
		return nil
	}
	out := new(GatewaySpecMembers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpecMembersStorageGroups) DeepCopyInto(out *GatewaySpecMembersStorageGroups) {
	*out = *in
	if in.ArraySize != nil {
		in, out := &in.ArraySize, &out.ArraySize
		*out = new(int64)
		**out = **in
	}
	if in.ArrayTypeID != nil {
		in, out := &in.ArrayTypeID, &out.ArrayTypeID
		*out = new(int64)
		**out = **in
	}
	if in.HardDrives != nil {
		in, out := &in.HardDrives, &out.HardDrives
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.PartitionTemplateID != nil {
		in, out := &in.PartitionTemplateID, &out.PartitionTemplateID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpecMembersStorageGroups.
func (in *GatewaySpecMembersStorageGroups) DeepCopy() *GatewaySpecMembersStorageGroups {
	if in == nil {
		return nil
	}
	out := new(GatewaySpecMembersStorageGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpecResource) DeepCopyInto(out *GatewaySpecResource) {
	*out = *in
	if in.AssociatedVlans != nil {
		in, out := &in.AssociatedVlans, &out.AssociatedVlans
		*out = make([]GatewaySpecAssociatedVlans, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]GatewaySpecMembers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PostInstallScriptURI != nil {
		in, out := &in.PostInstallScriptURI, &out.PostInstallScriptURI
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPAddressID != nil {
		in, out := &in.PrivateIPAddressID, &out.PrivateIPAddressID
		*out = new(int64)
		**out = **in
	}
	if in.PrivateIpv4Address != nil {
		in, out := &in.PrivateIpv4Address, &out.PrivateIpv4Address
		*out = new(string)
		**out = **in
	}
	if in.PrivateVLANID != nil {
		in, out := &in.PrivateVLANID, &out.PrivateVLANID
		*out = new(int64)
		**out = **in
	}
	if in.PublicIPAddressID != nil {
		in, out := &in.PublicIPAddressID, &out.PublicIPAddressID
		*out = new(int64)
		**out = **in
	}
	if in.PublicIpv4Address != nil {
		in, out := &in.PublicIpv4Address, &out.PublicIpv4Address
		*out = new(string)
		**out = **in
	}
	if in.PublicIpv6AddressID != nil {
		in, out := &in.PublicIpv6AddressID, &out.PublicIpv6AddressID
		*out = new(int64)
		**out = **in
	}
	if in.PublicVLANID != nil {
		in, out := &in.PublicVLANID, &out.PublicVLANID
		*out = new(int64)
		**out = **in
	}
	if in.SshKeyIDS != nil {
		in, out := &in.SshKeyIDS, &out.SshKeyIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpecResource.
func (in *GatewaySpecResource) DeepCopy() *GatewaySpecResource {
	if in == nil {
		return nil
	}
	out := new(GatewaySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayStatus) DeepCopyInto(out *GatewayStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayStatus.
func (in *GatewayStatus) DeepCopy() *GatewayStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayVLANAssociation) DeepCopyInto(out *GatewayVLANAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayVLANAssociation.
func (in *GatewayVLANAssociation) DeepCopy() *GatewayVLANAssociation {
	if in == nil {
		return nil
	}
	out := new(GatewayVLANAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayVLANAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayVLANAssociationList) DeepCopyInto(out *GatewayVLANAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayVLANAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayVLANAssociationList.
func (in *GatewayVLANAssociationList) DeepCopy() *GatewayVLANAssociationList {
	if in == nil {
		return nil
	}
	out := new(GatewayVLANAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayVLANAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayVLANAssociationSpec) DeepCopyInto(out *GatewayVLANAssociationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(GatewayVLANAssociationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayVLANAssociationSpec.
func (in *GatewayVLANAssociationSpec) DeepCopy() *GatewayVLANAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(GatewayVLANAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayVLANAssociationSpecResource) DeepCopyInto(out *GatewayVLANAssociationSpecResource) {
	*out = *in
	if in.Bypass != nil {
		in, out := &in.Bypass, &out.Bypass
		*out = new(bool)
		**out = **in
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(int64)
		**out = **in
	}
	if in.NetworkVLANID != nil {
		in, out := &in.NetworkVLANID, &out.NetworkVLANID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayVLANAssociationSpecResource.
func (in *GatewayVLANAssociationSpecResource) DeepCopy() *GatewayVLANAssociationSpecResource {
	if in == nil {
		return nil
	}
	out := new(GatewayVLANAssociationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayVLANAssociationStatus) DeepCopyInto(out *GatewayVLANAssociationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayVLANAssociationStatus.
func (in *GatewayVLANAssociationStatus) DeepCopy() *GatewayVLANAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayVLANAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSgAttachment) DeepCopyInto(out *InterfaceSgAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSgAttachment.
func (in *InterfaceSgAttachment) DeepCopy() *InterfaceSgAttachment {
	if in == nil {
		return nil
	}
	out := new(InterfaceSgAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterfaceSgAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSgAttachmentList) DeepCopyInto(out *InterfaceSgAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InterfaceSgAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSgAttachmentList.
func (in *InterfaceSgAttachmentList) DeepCopy() *InterfaceSgAttachmentList {
	if in == nil {
		return nil
	}
	out := new(InterfaceSgAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterfaceSgAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSgAttachmentSpec) DeepCopyInto(out *InterfaceSgAttachmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(InterfaceSgAttachmentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSgAttachmentSpec.
func (in *InterfaceSgAttachmentSpec) DeepCopy() *InterfaceSgAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(InterfaceSgAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSgAttachmentSpecResource) DeepCopyInto(out *InterfaceSgAttachmentSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(int64)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(int64)
		**out = **in
	}
	if in.SoftReboot != nil {
		in, out := &in.SoftReboot, &out.SoftReboot
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSgAttachmentSpecResource.
func (in *InterfaceSgAttachmentSpecResource) DeepCopy() *InterfaceSgAttachmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(InterfaceSgAttachmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSgAttachmentStatus) DeepCopyInto(out *InterfaceSgAttachmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSgAttachmentStatus.
func (in *InterfaceSgAttachmentStatus) DeepCopy() *InterfaceSgAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(InterfaceSgAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIP) DeepCopyInto(out *PublicIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIP.
func (in *PublicIP) DeepCopy() *PublicIP {
	if in == nil {
		return nil
	}
	out := new(PublicIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPList) DeepCopyInto(out *PublicIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPList.
func (in *PublicIPList) DeepCopy() *PublicIPList {
	if in == nil {
		return nil
	}
	out := new(PublicIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPSpec) DeepCopyInto(out *PublicIPSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PublicIPSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPSpec.
func (in *PublicIPSpec) DeepCopy() *PublicIPSpec {
	if in == nil {
		return nil
	}
	out := new(PublicIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPSpecResource) DeepCopyInto(out *PublicIPSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.RoutesTo != nil {
		in, out := &in.RoutesTo, &out.RoutesTo
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPSpecResource.
func (in *PublicIPSpecResource) DeepCopy() *PublicIPSpecResource {
	if in == nil {
		return nil
	}
	out := new(PublicIPSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPStatus) DeepCopyInto(out *PublicIPStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPStatus.
func (in *PublicIPStatus) DeepCopy() *PublicIPStatus {
	if in == nil {
		return nil
	}
	out := new(PublicIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vlan) DeepCopyInto(out *Vlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vlan.
func (in *Vlan) DeepCopy() *Vlan {
	if in == nil {
		return nil
	}
	out := new(Vlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanList) DeepCopyInto(out *VlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanList.
func (in *VlanList) DeepCopy() *VlanList {
	if in == nil {
		return nil
	}
	out := new(VlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpanning) DeepCopyInto(out *VlanSpanning) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpanning.
func (in *VlanSpanning) DeepCopy() *VlanSpanning {
	if in == nil {
		return nil
	}
	out := new(VlanSpanning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VlanSpanning) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpanningList) DeepCopyInto(out *VlanSpanningList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VlanSpanning, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpanningList.
func (in *VlanSpanningList) DeepCopy() *VlanSpanningList {
	if in == nil {
		return nil
	}
	out := new(VlanSpanningList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VlanSpanningList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpanningSpec) DeepCopyInto(out *VlanSpanningSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VlanSpanningSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpanningSpec.
func (in *VlanSpanningSpec) DeepCopy() *VlanSpanningSpec {
	if in == nil {
		return nil
	}
	out := new(VlanSpanningSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpanningSpecResource) DeepCopyInto(out *VlanSpanningSpecResource) {
	*out = *in
	if in.VlanSpanning != nil {
		in, out := &in.VlanSpanning, &out.VlanSpanning
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpanningSpecResource.
func (in *VlanSpanningSpecResource) DeepCopy() *VlanSpanningSpecResource {
	if in == nil {
		return nil
	}
	out := new(VlanSpanningSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpanningStatus) DeepCopyInto(out *VlanSpanningStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpanningStatus.
func (in *VlanSpanningStatus) DeepCopy() *VlanSpanningStatus {
	if in == nil {
		return nil
	}
	out := new(VlanSpanningStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpec) DeepCopyInto(out *VlanSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VlanSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpec.
func (in *VlanSpec) DeepCopy() *VlanSpec {
	if in == nil {
		return nil
	}
	out := new(VlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpecResource) DeepCopyInto(out *VlanSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ChildResourceCount != nil {
		in, out := &in.ChildResourceCount, &out.ChildResourceCount
		*out = new(int64)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
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
	if in.RouterHostname != nil {
		in, out := &in.RouterHostname, &out.RouterHostname
		*out = new(string)
		**out = **in
	}
	if in.SoftlayerManaged != nil {
		in, out := &in.SoftlayerManaged, &out.SoftlayerManaged
		*out = new(bool)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]VlanSpecSubnets, len(*in))
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
	if in.VlanNumber != nil {
		in, out := &in.VlanNumber, &out.VlanNumber
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpecResource.
func (in *VlanSpecResource) DeepCopy() *VlanSpecResource {
	if in == nil {
		return nil
	}
	out := new(VlanSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanSpecSubnets) DeepCopyInto(out *VlanSpecSubnets) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(int64)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.SubnetSize != nil {
		in, out := &in.SubnetSize, &out.SubnetSize
		*out = new(int64)
		**out = **in
	}
	if in.SubnetType != nil {
		in, out := &in.SubnetType, &out.SubnetType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanSpecSubnets.
func (in *VlanSpecSubnets) DeepCopy() *VlanSpecSubnets {
	if in == nil {
		return nil
	}
	out := new(VlanSpecSubnets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VlanStatus) DeepCopyInto(out *VlanStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VlanStatus.
func (in *VlanStatus) DeepCopy() *VlanStatus {
	if in == nil {
		return nil
	}
	out := new(VlanStatus)
	in.DeepCopyInto(out)
	return out
}
