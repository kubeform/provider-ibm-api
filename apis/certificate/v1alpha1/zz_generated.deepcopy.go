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

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerImport) DeepCopyInto(out *ManagerImport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerImport.
func (in *ManagerImport) DeepCopy() *ManagerImport {
	if in == nil {
		return nil
	}
	out := new(ManagerImport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerImport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerImportList) DeepCopyInto(out *ManagerImportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerImport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerImportList.
func (in *ManagerImportList) DeepCopy() *ManagerImportList {
	if in == nil {
		return nil
	}
	out := new(ManagerImportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerImportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerImportSpec) DeepCopyInto(out *ManagerImportSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ManagerImportSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerImportSpec.
func (in *ManagerImportSpec) DeepCopy() *ManagerImportSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerImportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerImportSpecResource) DeepCopyInto(out *ManagerImportSpecResource) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.BeginsOn != nil {
		in, out := &in.BeginsOn, &out.BeginsOn
		*out = new(int64)
		**out = **in
	}
	if in.CertificateManagerInstanceID != nil {
		in, out := &in.CertificateManagerInstanceID, &out.CertificateManagerInstanceID
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExpiresOn != nil {
		in, out := &in.ExpiresOn, &out.ExpiresOn
		*out = new(int64)
		**out = **in
	}
	if in.HasPrevious != nil {
		in, out := &in.HasPrevious, &out.HasPrevious
		*out = new(bool)
		**out = **in
	}
	if in.Imported != nil {
		in, out := &in.Imported, &out.Imported
		*out = new(bool)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerImportSpecResource.
func (in *ManagerImportSpecResource) DeepCopy() *ManagerImportSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagerImportSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerImportStatus) DeepCopyInto(out *ManagerImportStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerImportStatus.
func (in *ManagerImportStatus) DeepCopy() *ManagerImportStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerImportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerOrder) DeepCopyInto(out *ManagerOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerOrder.
func (in *ManagerOrder) DeepCopy() *ManagerOrder {
	if in == nil {
		return nil
	}
	out := new(ManagerOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerOrderList) DeepCopyInto(out *ManagerOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagerOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerOrderList.
func (in *ManagerOrderList) DeepCopy() *ManagerOrderList {
	if in == nil {
		return nil
	}
	out := new(ManagerOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagerOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerOrderSpec) DeepCopyInto(out *ManagerOrderSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ManagerOrderSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerOrderSpec.
func (in *ManagerOrderSpec) DeepCopy() *ManagerOrderSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerOrderSpecResource) DeepCopyInto(out *ManagerOrderSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.AutoRenewEnabled != nil {
		in, out := &in.AutoRenewEnabled, &out.AutoRenewEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BeginsOn != nil {
		in, out := &in.BeginsOn, &out.BeginsOn
		*out = new(int64)
		**out = **in
	}
	if in.CertificateManagerInstanceID != nil {
		in, out := &in.CertificateManagerInstanceID, &out.CertificateManagerInstanceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DnsProviderInstanceCrn != nil {
		in, out := &in.DnsProviderInstanceCrn, &out.DnsProviderInstanceCrn
		*out = new(string)
		**out = **in
	}
	if in.DomainValidationMethod != nil {
		in, out := &in.DomainValidationMethod, &out.DomainValidationMethod
		*out = new(string)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExpiresOn != nil {
		in, out := &in.ExpiresOn, &out.ExpiresOn
		*out = new(int64)
		**out = **in
	}
	if in.HasPrevious != nil {
		in, out := &in.HasPrevious, &out.HasPrevious
		*out = new(bool)
		**out = **in
	}
	if in.Imported != nil {
		in, out := &in.Imported, &out.Imported
		*out = new(bool)
		**out = **in
	}
	if in.IssuanceInfo != nil {
		in, out := &in.IssuanceInfo, &out.IssuanceInfo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.KeyAlgorithm != nil {
		in, out := &in.KeyAlgorithm, &out.KeyAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RenewCertificate != nil {
		in, out := &in.RenewCertificate, &out.RenewCertificate
		*out = new(bool)
		**out = **in
	}
	if in.RotateKeys != nil {
		in, out := &in.RotateKeys, &out.RotateKeys
		*out = new(bool)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerOrderSpecResource.
func (in *ManagerOrderSpecResource) DeepCopy() *ManagerOrderSpecResource {
	if in == nil {
		return nil
	}
	out := new(ManagerOrderSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerOrderStatus) DeepCopyInto(out *ManagerOrderStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerOrderStatus.
func (in *ManagerOrderStatus) DeepCopy() *ManagerOrderStatus {
	if in == nil {
		return nil
	}
	out := new(ManagerOrderStatus)
	in.DeepCopyInto(out)
	return out
}
