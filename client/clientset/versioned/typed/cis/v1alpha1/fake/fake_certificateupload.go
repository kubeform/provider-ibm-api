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

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/cis/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateUploads implements CertificateUploadInterface
type FakeCertificateUploads struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var certificateuploadsResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "certificateuploads"}

var certificateuploadsKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "CertificateUpload"}

// Get takes name of the certificateUpload, and returns the corresponding certificateUpload object, and an error if there is any.
func (c *FakeCertificateUploads) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateUpload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificateuploadsResource, c.ns, name), &v1alpha1.CertificateUpload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateUpload), err
}

// List takes label and field selectors, and returns the list of CertificateUploads that match those selectors.
func (c *FakeCertificateUploads) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateUploadList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificateuploadsResource, certificateuploadsKind, c.ns, opts), &v1alpha1.CertificateUploadList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateUploadList{ListMeta: obj.(*v1alpha1.CertificateUploadList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateUploadList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateUploads.
func (c *FakeCertificateUploads) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificateuploadsResource, c.ns, opts))

}

// Create takes the representation of a certificateUpload and creates it.  Returns the server's representation of the certificateUpload, and an error, if there is any.
func (c *FakeCertificateUploads) Create(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.CreateOptions) (result *v1alpha1.CertificateUpload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificateuploadsResource, c.ns, certificateUpload), &v1alpha1.CertificateUpload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateUpload), err
}

// Update takes the representation of a certificateUpload and updates it. Returns the server's representation of the certificateUpload, and an error, if there is any.
func (c *FakeCertificateUploads) Update(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.UpdateOptions) (result *v1alpha1.CertificateUpload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificateuploadsResource, c.ns, certificateUpload), &v1alpha1.CertificateUpload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateUpload), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateUploads) UpdateStatus(ctx context.Context, certificateUpload *v1alpha1.CertificateUpload, opts v1.UpdateOptions) (*v1alpha1.CertificateUpload, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificateuploadsResource, "status", c.ns, certificateUpload), &v1alpha1.CertificateUpload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateUpload), err
}

// Delete takes name of the certificateUpload and deletes it. Returns an error if one occurs.
func (c *FakeCertificateUploads) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificateuploadsResource, c.ns, name), &v1alpha1.CertificateUpload{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateUploads) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificateuploadsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateUploadList{})
	return err
}

// Patch applies the patch and returns the patched certificateUpload.
func (c *FakeCertificateUploads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateUpload, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificateuploadsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CertificateUpload{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateUpload), err
}
