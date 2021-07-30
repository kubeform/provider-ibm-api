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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/hpcs/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HpcsesGetter has a method to return a HpcsInterface.
// A group's client should implement this interface.
type HpcsesGetter interface {
	Hpcses(namespace string) HpcsInterface
}

// HpcsInterface has methods to work with Hpcs resources.
type HpcsInterface interface {
	Create(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.CreateOptions) (*v1alpha1.Hpcs, error)
	Update(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.UpdateOptions) (*v1alpha1.Hpcs, error)
	UpdateStatus(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.UpdateOptions) (*v1alpha1.Hpcs, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Hpcs, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HpcsList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hpcs, err error)
	HpcsExpansion
}

// hpcses implements HpcsInterface
type hpcses struct {
	client rest.Interface
	ns     string
}

// newHpcses returns a Hpcses
func newHpcses(c *HpcsV1alpha1Client, namespace string) *hpcses {
	return &hpcses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hpcs, and returns the corresponding hpcs object, and an error if there is any.
func (c *hpcses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Hpcs, err error) {
	result = &v1alpha1.Hpcs{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hpcses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Hpcses that match those selectors.
func (c *hpcses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HpcsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HpcsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hpcses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hpcses.
func (c *hpcses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hpcses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hpcs and creates it.  Returns the server's representation of the hpcs, and an error, if there is any.
func (c *hpcses) Create(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.CreateOptions) (result *v1alpha1.Hpcs, err error) {
	result = &v1alpha1.Hpcs{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hpcses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hpcs).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hpcs and updates it. Returns the server's representation of the hpcs, and an error, if there is any.
func (c *hpcses) Update(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.UpdateOptions) (result *v1alpha1.Hpcs, err error) {
	result = &v1alpha1.Hpcs{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hpcses").
		Name(hpcs.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hpcs).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hpcses) UpdateStatus(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.UpdateOptions) (result *v1alpha1.Hpcs, err error) {
	result = &v1alpha1.Hpcs{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hpcses").
		Name(hpcs.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hpcs).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hpcs and deletes it. Returns an error if one occurs.
func (c *hpcses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hpcses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hpcses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hpcses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hpcs.
func (c *hpcses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hpcs, err error) {
	result = &v1alpha1.Hpcs{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hpcses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
