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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/storage/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EvaultsGetter has a method to return a EvaultInterface.
// A group's client should implement this interface.
type EvaultsGetter interface {
	Evaults(namespace string) EvaultInterface
}

// EvaultInterface has methods to work with Evault resources.
type EvaultInterface interface {
	Create(ctx context.Context, evault *v1alpha1.Evault, opts v1.CreateOptions) (*v1alpha1.Evault, error)
	Update(ctx context.Context, evault *v1alpha1.Evault, opts v1.UpdateOptions) (*v1alpha1.Evault, error)
	UpdateStatus(ctx context.Context, evault *v1alpha1.Evault, opts v1.UpdateOptions) (*v1alpha1.Evault, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Evault, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EvaultList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Evault, err error)
	EvaultExpansion
}

// evaults implements EvaultInterface
type evaults struct {
	client rest.Interface
	ns     string
}

// newEvaults returns a Evaults
func newEvaults(c *StorageV1alpha1Client, namespace string) *evaults {
	return &evaults{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the evault, and returns the corresponding evault object, and an error if there is any.
func (c *evaults) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Evault, err error) {
	result = &v1alpha1.Evault{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("evaults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Evaults that match those selectors.
func (c *evaults) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EvaultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EvaultList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("evaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested evaults.
func (c *evaults) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("evaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a evault and creates it.  Returns the server's representation of the evault, and an error, if there is any.
func (c *evaults) Create(ctx context.Context, evault *v1alpha1.Evault, opts v1.CreateOptions) (result *v1alpha1.Evault, err error) {
	result = &v1alpha1.Evault{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("evaults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(evault).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a evault and updates it. Returns the server's representation of the evault, and an error, if there is any.
func (c *evaults) Update(ctx context.Context, evault *v1alpha1.Evault, opts v1.UpdateOptions) (result *v1alpha1.Evault, err error) {
	result = &v1alpha1.Evault{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("evaults").
		Name(evault.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(evault).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *evaults) UpdateStatus(ctx context.Context, evault *v1alpha1.Evault, opts v1.UpdateOptions) (result *v1alpha1.Evault, err error) {
	result = &v1alpha1.Evault{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("evaults").
		Name(evault.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(evault).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the evault and deletes it. Returns an error if one occurs.
func (c *evaults) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("evaults").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *evaults) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("evaults").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched evault.
func (c *evaults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Evault, err error) {
	result = &v1alpha1.Evault{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("evaults").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
