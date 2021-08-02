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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/dns/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecondariesGetter has a method to return a SecondaryInterface.
// A group's client should implement this interface.
type SecondariesGetter interface {
	Secondaries(namespace string) SecondaryInterface
}

// SecondaryInterface has methods to work with Secondary resources.
type SecondaryInterface interface {
	Create(ctx context.Context, secondary *v1alpha1.Secondary, opts v1.CreateOptions) (*v1alpha1.Secondary, error)
	Update(ctx context.Context, secondary *v1alpha1.Secondary, opts v1.UpdateOptions) (*v1alpha1.Secondary, error)
	UpdateStatus(ctx context.Context, secondary *v1alpha1.Secondary, opts v1.UpdateOptions) (*v1alpha1.Secondary, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Secondary, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SecondaryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Secondary, err error)
	SecondaryExpansion
}

// secondaries implements SecondaryInterface
type secondaries struct {
	client rest.Interface
	ns     string
}

// newSecondaries returns a Secondaries
func newSecondaries(c *DnsV1alpha1Client, namespace string) *secondaries {
	return &secondaries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secondary, and returns the corresponding secondary object, and an error if there is any.
func (c *secondaries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Secondary, err error) {
	result = &v1alpha1.Secondary{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secondaries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Secondaries that match those selectors.
func (c *secondaries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecondaryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecondaryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secondaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secondaries.
func (c *secondaries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secondaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a secondary and creates it.  Returns the server's representation of the secondary, and an error, if there is any.
func (c *secondaries) Create(ctx context.Context, secondary *v1alpha1.Secondary, opts v1.CreateOptions) (result *v1alpha1.Secondary, err error) {
	result = &v1alpha1.Secondary{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secondaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secondary).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a secondary and updates it. Returns the server's representation of the secondary, and an error, if there is any.
func (c *secondaries) Update(ctx context.Context, secondary *v1alpha1.Secondary, opts v1.UpdateOptions) (result *v1alpha1.Secondary, err error) {
	result = &v1alpha1.Secondary{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secondaries").
		Name(secondary.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secondary).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *secondaries) UpdateStatus(ctx context.Context, secondary *v1alpha1.Secondary, opts v1.UpdateOptions) (result *v1alpha1.Secondary, err error) {
	result = &v1alpha1.Secondary{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secondaries").
		Name(secondary.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secondary).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the secondary and deletes it. Returns an error if one occurs.
func (c *secondaries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secondaries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secondaries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secondaries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched secondary.
func (c *secondaries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Secondary, err error) {
	result = &v1alpha1.Secondary{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secondaries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}