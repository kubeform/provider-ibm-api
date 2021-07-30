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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/container/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkerPoolsGetter has a method to return a WorkerPoolInterface.
// A group's client should implement this interface.
type WorkerPoolsGetter interface {
	WorkerPools(namespace string) WorkerPoolInterface
}

// WorkerPoolInterface has methods to work with WorkerPool resources.
type WorkerPoolInterface interface {
	Create(ctx context.Context, workerPool *v1alpha1.WorkerPool, opts v1.CreateOptions) (*v1alpha1.WorkerPool, error)
	Update(ctx context.Context, workerPool *v1alpha1.WorkerPool, opts v1.UpdateOptions) (*v1alpha1.WorkerPool, error)
	UpdateStatus(ctx context.Context, workerPool *v1alpha1.WorkerPool, opts v1.UpdateOptions) (*v1alpha1.WorkerPool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.WorkerPool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.WorkerPoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkerPool, err error)
	WorkerPoolExpansion
}

// workerPools implements WorkerPoolInterface
type workerPools struct {
	client rest.Interface
	ns     string
}

// newWorkerPools returns a WorkerPools
func newWorkerPools(c *ContainerV1alpha1Client, namespace string) *workerPools {
	return &workerPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the workerPool, and returns the corresponding workerPool object, and an error if there is any.
func (c *workerPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkerPool, err error) {
	result = &v1alpha1.WorkerPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workerpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkerPools that match those selectors.
func (c *workerPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkerPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WorkerPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workerpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workerPools.
func (c *workerPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("workerpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workerPool and creates it.  Returns the server's representation of the workerPool, and an error, if there is any.
func (c *workerPools) Create(ctx context.Context, workerPool *v1alpha1.WorkerPool, opts v1.CreateOptions) (result *v1alpha1.WorkerPool, err error) {
	result = &v1alpha1.WorkerPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("workerpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workerPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workerPool and updates it. Returns the server's representation of the workerPool, and an error, if there is any.
func (c *workerPools) Update(ctx context.Context, workerPool *v1alpha1.WorkerPool, opts v1.UpdateOptions) (result *v1alpha1.WorkerPool, err error) {
	result = &v1alpha1.WorkerPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workerpools").
		Name(workerPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workerPool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *workerPools) UpdateStatus(ctx context.Context, workerPool *v1alpha1.WorkerPool, opts v1.UpdateOptions) (result *v1alpha1.WorkerPool, err error) {
	result = &v1alpha1.WorkerPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workerpools").
		Name(workerPool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workerPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workerPool and deletes it. Returns an error if one occurs.
func (c *workerPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workerpools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workerPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workerpools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workerPool.
func (c *workerPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkerPool, err error) {
	result = &v1alpha1.WorkerPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("workerpools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
