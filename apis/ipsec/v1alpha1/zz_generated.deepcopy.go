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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vpn) DeepCopyInto(out *Vpn) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vpn.
func (in *Vpn) DeepCopy() *Vpn {
	if in == nil {
		return nil
	}
	out := new(Vpn)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vpn) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnList) DeepCopyInto(out *VpnList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vpn, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnList.
func (in *VpnList) DeepCopy() *VpnList {
	if in == nil {
		return nil
	}
	out := new(VpnList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnSpec) DeepCopyInto(out *VpnSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(VpnSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnSpec.
func (in *VpnSpec) DeepCopy() *VpnSpec {
	if in == nil {
		return nil
	}
	out := new(VpnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnSpecAddressTranslation) DeepCopyInto(out *VpnSpecAddressTranslation) {
	*out = *in
	if in.InternalIPAdress != nil {
		in, out := &in.InternalIPAdress, &out.InternalIPAdress
		*out = new(string)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.RemoteIPAdress != nil {
		in, out := &in.RemoteIPAdress, &out.RemoteIPAdress
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnSpecAddressTranslation.
func (in *VpnSpecAddressTranslation) DeepCopy() *VpnSpecAddressTranslation {
	if in == nil {
		return nil
	}
	out := new(VpnSpecAddressTranslation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnSpecPhaseOne) DeepCopyInto(out *VpnSpecPhaseOne) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(string)
		**out = **in
	}
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(int64)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(string)
		**out = **in
	}
	if in.Keylife != nil {
		in, out := &in.Keylife, &out.Keylife
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnSpecPhaseOne.
func (in *VpnSpecPhaseOne) DeepCopy() *VpnSpecPhaseOne {
	if in == nil {
		return nil
	}
	out := new(VpnSpecPhaseOne)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnSpecPhaseTwo) DeepCopyInto(out *VpnSpecPhaseTwo) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(string)
		**out = **in
	}
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(int64)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(string)
		**out = **in
	}
	if in.Keylife != nil {
		in, out := &in.Keylife, &out.Keylife
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnSpecPhaseTwo.
func (in *VpnSpecPhaseTwo) DeepCopy() *VpnSpecPhaseTwo {
	if in == nil {
		return nil
	}
	out := new(VpnSpecPhaseTwo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnSpecRemoteSubnet) DeepCopyInto(out *VpnSpecRemoteSubnet) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.RemoteIPAdress != nil {
		in, out := &in.RemoteIPAdress, &out.RemoteIPAdress
		*out = new(string)
		**out = **in
	}
	if in.RemoteIPCIDR != nil {
		in, out := &in.RemoteIPCIDR, &out.RemoteIPCIDR
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnSpecRemoteSubnet.
func (in *VpnSpecRemoteSubnet) DeepCopy() *VpnSpecRemoteSubnet {
	if in == nil {
		return nil
	}
	out := new(VpnSpecRemoteSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnSpecResource) DeepCopyInto(out *VpnSpecResource) {
	*out = *in
	if in.AddressTranslation != nil {
		in, out := &in.AddressTranslation, &out.AddressTranslation
		*out = new(VpnSpecAddressTranslation)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomerPeerIP != nil {
		in, out := &in.CustomerPeerIP, &out.CustomerPeerIP
		*out = new(string)
		**out = **in
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.InternalPeerIPAddress != nil {
		in, out := &in.InternalPeerIPAddress, &out.InternalPeerIPAddress
		*out = new(string)
		**out = **in
	}
	if in.InternalSubnetID != nil {
		in, out := &in.InternalSubnetID, &out.InternalSubnetID
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PhaseOne != nil {
		in, out := &in.PhaseOne, &out.PhaseOne
		*out = new(VpnSpecPhaseOne)
		(*in).DeepCopyInto(*out)
	}
	if in.PhaseTwo != nil {
		in, out := &in.PhaseTwo, &out.PhaseTwo
		*out = new(VpnSpecPhaseTwo)
		(*in).DeepCopyInto(*out)
	}
	if in.PresharedKey != nil {
		in, out := &in.PresharedKey, &out.PresharedKey
		*out = new(string)
		**out = **in
	}
	if in.RemoteSubnet != nil {
		in, out := &in.RemoteSubnet, &out.RemoteSubnet
		*out = new(VpnSpecRemoteSubnet)
		(*in).DeepCopyInto(*out)
	}
	if in.RemoteSubnetID != nil {
		in, out := &in.RemoteSubnetID, &out.RemoteSubnetID
		*out = new(int64)
		**out = **in
	}
	if in.ServiceSubnetID != nil {
		in, out := &in.ServiceSubnetID, &out.ServiceSubnetID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnSpecResource.
func (in *VpnSpecResource) DeepCopy() *VpnSpecResource {
	if in == nil {
		return nil
	}
	out := new(VpnSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnStatus) DeepCopyInto(out *VpnStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnStatus.
func (in *VpnStatus) DeepCopy() *VpnStatus {
	if in == nil {
		return nil
	}
	out := new(VpnStatus)
	in.DeepCopyInto(out)
	return out
}
