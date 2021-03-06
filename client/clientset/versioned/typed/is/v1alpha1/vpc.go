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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/is/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VpcsGetter has a method to return a VpcInterface.
// A group's client should implement this interface.
type VpcsGetter interface {
	Vpcs(namespace string) VpcInterface
}

// VpcInterface has methods to work with Vpc resources.
type VpcInterface interface {
	Create(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.CreateOptions) (*v1alpha1.Vpc, error)
	Update(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.UpdateOptions) (*v1alpha1.Vpc, error)
	UpdateStatus(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.UpdateOptions) (*v1alpha1.Vpc, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Vpc, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VpcList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vpc, err error)
	VpcExpansion
}

// vpcs implements VpcInterface
type vpcs struct {
	client rest.Interface
	ns     string
}

// newVpcs returns a Vpcs
func newVpcs(c *IsV1alpha1Client, namespace string) *vpcs {
	return &vpcs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vpc, and returns the corresponding vpc object, and an error if there is any.
func (c *vpcs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vpc, err error) {
	result = &v1alpha1.Vpc{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Vpcs that match those selectors.
func (c *vpcs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VpcList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpcs.
func (c *vpcs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vpc and creates it.  Returns the server's representation of the vpc, and an error, if there is any.
func (c *vpcs) Create(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.CreateOptions) (result *v1alpha1.Vpc, err error) {
	result = &v1alpha1.Vpc{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpcs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpc).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vpc and updates it. Returns the server's representation of the vpc, and an error, if there is any.
func (c *vpcs) Update(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.UpdateOptions) (result *v1alpha1.Vpc, err error) {
	result = &v1alpha1.Vpc{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcs").
		Name(vpc.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpc).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vpcs) UpdateStatus(ctx context.Context, vpc *v1alpha1.Vpc, opts v1.UpdateOptions) (result *v1alpha1.Vpc, err error) {
	result = &v1alpha1.Vpc{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpcs").
		Name(vpc.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpc).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vpc and deletes it. Returns an error if one occurs.
func (c *vpcs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpcs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpcs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vpc.
func (c *vpcs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vpc, err error) {
	result = &v1alpha1.Vpc{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpcs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
