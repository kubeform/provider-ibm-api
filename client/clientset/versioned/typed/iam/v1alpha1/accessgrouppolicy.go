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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/iam/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AccessGroupPoliciesGetter has a method to return a AccessGroupPolicyInterface.
// A group's client should implement this interface.
type AccessGroupPoliciesGetter interface {
	AccessGroupPolicies(namespace string) AccessGroupPolicyInterface
}

// AccessGroupPolicyInterface has methods to work with AccessGroupPolicy resources.
type AccessGroupPolicyInterface interface {
	Create(ctx context.Context, accessGroupPolicy *v1alpha1.AccessGroupPolicy, opts v1.CreateOptions) (*v1alpha1.AccessGroupPolicy, error)
	Update(ctx context.Context, accessGroupPolicy *v1alpha1.AccessGroupPolicy, opts v1.UpdateOptions) (*v1alpha1.AccessGroupPolicy, error)
	UpdateStatus(ctx context.Context, accessGroupPolicy *v1alpha1.AccessGroupPolicy, opts v1.UpdateOptions) (*v1alpha1.AccessGroupPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AccessGroupPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AccessGroupPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessGroupPolicy, err error)
	AccessGroupPolicyExpansion
}

// accessGroupPolicies implements AccessGroupPolicyInterface
type accessGroupPolicies struct {
	client rest.Interface
	ns     string
}

// newAccessGroupPolicies returns a AccessGroupPolicies
func newAccessGroupPolicies(c *IamV1alpha1Client, namespace string) *accessGroupPolicies {
	return &accessGroupPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the accessGroupPolicy, and returns the corresponding accessGroupPolicy object, and an error if there is any.
func (c *accessGroupPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessGroupPolicy, err error) {
	result = &v1alpha1.AccessGroupPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AccessGroupPolicies that match those selectors.
func (c *accessGroupPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessGroupPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AccessGroupPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested accessGroupPolicies.
func (c *accessGroupPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a accessGroupPolicy and creates it.  Returns the server's representation of the accessGroupPolicy, and an error, if there is any.
func (c *accessGroupPolicies) Create(ctx context.Context, accessGroupPolicy *v1alpha1.AccessGroupPolicy, opts v1.CreateOptions) (result *v1alpha1.AccessGroupPolicy, err error) {
	result = &v1alpha1.AccessGroupPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(accessGroupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a accessGroupPolicy and updates it. Returns the server's representation of the accessGroupPolicy, and an error, if there is any.
func (c *accessGroupPolicies) Update(ctx context.Context, accessGroupPolicy *v1alpha1.AccessGroupPolicy, opts v1.UpdateOptions) (result *v1alpha1.AccessGroupPolicy, err error) {
	result = &v1alpha1.AccessGroupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		Name(accessGroupPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(accessGroupPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *accessGroupPolicies) UpdateStatus(ctx context.Context, accessGroupPolicy *v1alpha1.AccessGroupPolicy, opts v1.UpdateOptions) (result *v1alpha1.AccessGroupPolicy, err error) {
	result = &v1alpha1.AccessGroupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		Name(accessGroupPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(accessGroupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the accessGroupPolicy and deletes it. Returns an error if one occurs.
func (c *accessGroupPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *accessGroupPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched accessGroupPolicy.
func (c *accessGroupPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessGroupPolicy, err error) {
	result = &v1alpha1.AccessGroupPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("accessgrouppolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}