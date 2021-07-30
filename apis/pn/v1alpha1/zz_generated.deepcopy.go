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
func (in *ApplicationChrome) DeepCopyInto(out *ApplicationChrome) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationChrome.
func (in *ApplicationChrome) DeepCopy() *ApplicationChrome {
	if in == nil {
		return nil
	}
	out := new(ApplicationChrome)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationChrome) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationChromeList) DeepCopyInto(out *ApplicationChromeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationChrome, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationChromeList.
func (in *ApplicationChromeList) DeepCopy() *ApplicationChromeList {
	if in == nil {
		return nil
	}
	out := new(ApplicationChromeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationChromeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationChromeSpec) DeepCopyInto(out *ApplicationChromeSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ApplicationChromeSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationChromeSpec.
func (in *ApplicationChromeSpec) DeepCopy() *ApplicationChromeSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationChromeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationChromeSpecResource) DeepCopyInto(out *ApplicationChromeSpecResource) {
	*out = *in
	if in.Guid != nil {
		in, out := &in.Guid, &out.Guid
		*out = new(string)
		**out = **in
	}
	if in.ServerKey != nil {
		in, out := &in.ServerKey, &out.ServerKey
		*out = new(string)
		**out = **in
	}
	if in.WebSiteURL != nil {
		in, out := &in.WebSiteURL, &out.WebSiteURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationChromeSpecResource.
func (in *ApplicationChromeSpecResource) DeepCopy() *ApplicationChromeSpecResource {
	if in == nil {
		return nil
	}
	out := new(ApplicationChromeSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationChromeStatus) DeepCopyInto(out *ApplicationChromeStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationChromeStatus.
func (in *ApplicationChromeStatus) DeepCopy() *ApplicationChromeStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationChromeStatus)
	in.DeepCopyInto(out)
	return out
}
