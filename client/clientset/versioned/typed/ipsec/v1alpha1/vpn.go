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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/ipsec/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VpnsGetter has a method to return a VpnInterface.
// A group's client should implement this interface.
type VpnsGetter interface {
	Vpns(namespace string) VpnInterface
}

// VpnInterface has methods to work with Vpn resources.
type VpnInterface interface {
	Create(ctx context.Context, vpn *v1alpha1.Vpn, opts v1.CreateOptions) (*v1alpha1.Vpn, error)
	Update(ctx context.Context, vpn *v1alpha1.Vpn, opts v1.UpdateOptions) (*v1alpha1.Vpn, error)
	UpdateStatus(ctx context.Context, vpn *v1alpha1.Vpn, opts v1.UpdateOptions) (*v1alpha1.Vpn, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Vpn, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VpnList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vpn, err error)
	VpnExpansion
}

// vpns implements VpnInterface
type vpns struct {
	client rest.Interface
	ns     string
}

// newVpns returns a Vpns
func newVpns(c *IpsecV1alpha1Client, namespace string) *vpns {
	return &vpns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpn, and returns the corresponding vpn object, and an error if there is any.
func (c *vpns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vpn, err error) {
	result = &v1alpha1.Vpn{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Vpns that match those selectors.
func (c *vpns) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpnList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpnList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpns.
func (c *vpns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vpn and creates it.  Returns the server's representation of the vpn, and an error, if there is any.
func (c *vpns) Create(ctx context.Context, vpn *v1alpha1.Vpn, opts v1.CreateOptions) (result *v1alpha1.Vpn, err error) {
	result = &v1alpha1.Vpn{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpn).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vpn and updates it. Returns the server's representation of the vpn, and an error, if there is any.
func (c *vpns) Update(ctx context.Context, vpn *v1alpha1.Vpn, opts v1.UpdateOptions) (result *v1alpha1.Vpn, err error) {
	result = &v1alpha1.Vpn{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpns").
		Name(vpn.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpn).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vpns) UpdateStatus(ctx context.Context, vpn *v1alpha1.Vpn, opts v1.UpdateOptions) (result *v1alpha1.Vpn, err error) {
	result = &v1alpha1.Vpn{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpns").
		Name(vpn.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpn).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vpn and deletes it. Returns an error if one occurs.
func (c *vpns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpns").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpns").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vpn.
func (c *vpns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vpn, err error) {
	result = &v1alpha1.Vpn{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpns").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}