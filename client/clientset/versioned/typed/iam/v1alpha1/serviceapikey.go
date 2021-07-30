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

// ServiceAPIKeysGetter has a method to return a ServiceAPIKeyInterface.
// A group's client should implement this interface.
type ServiceAPIKeysGetter interface {
	ServiceAPIKeys(namespace string) ServiceAPIKeyInterface
}

// ServiceAPIKeyInterface has methods to work with ServiceAPIKey resources.
type ServiceAPIKeyInterface interface {
	Create(ctx context.Context, serviceAPIKey *v1alpha1.ServiceAPIKey, opts v1.CreateOptions) (*v1alpha1.ServiceAPIKey, error)
	Update(ctx context.Context, serviceAPIKey *v1alpha1.ServiceAPIKey, opts v1.UpdateOptions) (*v1alpha1.ServiceAPIKey, error)
	UpdateStatus(ctx context.Context, serviceAPIKey *v1alpha1.ServiceAPIKey, opts v1.UpdateOptions) (*v1alpha1.ServiceAPIKey, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceAPIKey, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceAPIKeyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceAPIKey, err error)
	ServiceAPIKeyExpansion
}

// serviceAPIKeys implements ServiceAPIKeyInterface
type serviceAPIKeys struct {
	client rest.Interface
	ns     string
}

// newServiceAPIKeys returns a ServiceAPIKeys
func newServiceAPIKeys(c *IamV1alpha1Client, namespace string) *serviceAPIKeys {
	return &serviceAPIKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceAPIKey, and returns the corresponding serviceAPIKey object, and an error if there is any.
func (c *serviceAPIKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceAPIKey, err error) {
	result = &v1alpha1.ServiceAPIKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceapikeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceAPIKeys that match those selectors.
func (c *serviceAPIKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceAPIKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceAPIKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceAPIKeys.
func (c *serviceAPIKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceAPIKey and creates it.  Returns the server's representation of the serviceAPIKey, and an error, if there is any.
func (c *serviceAPIKeys) Create(ctx context.Context, serviceAPIKey *v1alpha1.ServiceAPIKey, opts v1.CreateOptions) (result *v1alpha1.ServiceAPIKey, err error) {
	result = &v1alpha1.ServiceAPIKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceapikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceAPIKey).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceAPIKey and updates it. Returns the server's representation of the serviceAPIKey, and an error, if there is any.
func (c *serviceAPIKeys) Update(ctx context.Context, serviceAPIKey *v1alpha1.ServiceAPIKey, opts v1.UpdateOptions) (result *v1alpha1.ServiceAPIKey, err error) {
	result = &v1alpha1.ServiceAPIKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceapikeys").
		Name(serviceAPIKey.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceAPIKey).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceAPIKeys) UpdateStatus(ctx context.Context, serviceAPIKey *v1alpha1.ServiceAPIKey, opts v1.UpdateOptions) (result *v1alpha1.ServiceAPIKey, err error) {
	result = &v1alpha1.ServiceAPIKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceapikeys").
		Name(serviceAPIKey.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceAPIKey).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceAPIKey and deletes it. Returns an error if one occurs.
func (c *serviceAPIKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceapikeys").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceAPIKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceapikeys").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceAPIKey.
func (c *serviceAPIKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceAPIKey, err error) {
	result = &v1alpha1.ServiceAPIKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceapikeys").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
