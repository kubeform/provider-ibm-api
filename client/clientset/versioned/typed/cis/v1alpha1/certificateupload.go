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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/cis/v1alpha1"
	scheme "kubeform.dev/provider-ibm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CertificateUploadsGetter has a method to return a CertificateUploadInterface.
// A group's client should implement this interface.
type CertificateUploadsGetter interface {
	CertificateUploads(namespace string) CertificateUploadInterface
}

// CertificateUploadInterface has methods to work with CertificateUpload resources.
type CertificateUploadInterface interface {
	Create(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.CreateOptions) (*v1alpha1.CertificateUpload, error)
	Update(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.UpdateOptions) (*v1alpha1.CertificateUpload, error)
	UpdateStatus(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.UpdateOptions) (*v1alpha1.CertificateUpload, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CertificateUpload, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CertificateUploadList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateUpload, err error)
	CertificateUploadExpansion
}

// certificateUploads implements CertificateUploadInterface
type certificateUploads struct {
	client rest.Interface
	ns     string
}

// newCertificateUploads returns a CertificateUploads
func newCertificateUploads(c *CisV1alpha1Client, namespace string) *certificateUploads {
	return &certificateUploads{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the certificateUpload, and returns the corresponding certificateUpload object, and an error if there is any.
func (c *certificateUploads) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateUpload, err error) {
	result = &v1alpha1.CertificateUpload{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificateuploads").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CertificateUploads that match those selectors.
func (c *certificateUploads) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateUploadList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CertificateUploadList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificateuploads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certificateUploads.
func (c *certificateUploads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("certificateuploads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a certificateUpload and creates it.  Returns the server's representation of the certificateUpload, and an error, if there is any.
func (c *certificateUploads) Create(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.CreateOptions) (result *v1alpha1.CertificateUpload, err error) {
	result = &v1alpha1.CertificateUpload{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("certificateuploads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateUpload).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a certificateUpload and updates it. Returns the server's representation of the certificateUpload, and an error, if there is any.
func (c *certificateUploads) Update(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.UpdateOptions) (result *v1alpha1.CertificateUpload, err error) {
	result = &v1alpha1.CertificateUpload{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificateuploads").
		Name(certificateUpload.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateUpload).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *certificateUploads) UpdateStatus(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.UpdateOptions) (result *v1alpha1.CertificateUpload, err error) {
	result = &v1alpha1.CertificateUpload{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificateuploads").
		Name(certificateUpload.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateUpload).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the certificateUpload and deletes it. Returns an error if one occurs.
func (c *certificateUploads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificateuploads").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certificateUploads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificateuploads").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched certificateUpload.
func (c *certificateUploads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateUpload, err error) {
	result = &v1alpha1.CertificateUpload{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("certificateuploads").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
