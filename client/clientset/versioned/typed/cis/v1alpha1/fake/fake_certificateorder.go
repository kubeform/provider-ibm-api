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

// FakeCertificateOrders implements CertificateOrderInterface
type FakeCertificateOrders struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var certificateordersResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "certificateorders"}

var certificateordersKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "CertificateOrder"}

// Get takes name of the certificateOrder, and returns the corresponding certificateOrder object, and an error if there is any.
func (c *FakeCertificateOrders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificateordersResource, c.ns, name), &v1alpha1.CertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateOrder), err
}

// List takes label and field selectors, and returns the list of CertificateOrders that match those selectors.
func (c *FakeCertificateOrders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateOrderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificateordersResource, certificateordersKind, c.ns, opts), &v1alpha1.CertificateOrderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateOrderList{ListMeta: obj.(*v1alpha1.CertificateOrderList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateOrderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateOrders.
func (c *FakeCertificateOrders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificateordersResource, c.ns, opts))

}

// Create takes the representation of a certificateOrder and creates it.  Returns the server's representation of the certificateOrder, and an error, if there is any.
func (c *FakeCertificateOrders) Create(ctx context.Context, certificateOrder *v1alpha1.CertificateOrder, opts v1.CreateOptions) (result *v1alpha1.CertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificateordersResource, c.ns, certificateOrder), &v1alpha1.CertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateOrder), err
}

// Update takes the representation of a certificateOrder and updates it. Returns the server's representation of the certificateOrder, and an error, if there is any.
func (c *FakeCertificateOrders) Update(ctx context.Context, certificateOrder *v1alpha1.CertificateOrder, opts v1.UpdateOptions) (result *v1alpha1.CertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificateordersResource, c.ns, certificateOrder), &v1alpha1.CertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateOrder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateOrders) UpdateStatus(ctx context.Context, certificateOrder *v1alpha1.CertificateOrder, opts v1.UpdateOptions) (*v1alpha1.CertificateOrder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificateordersResource, "status", c.ns, certificateOrder), &v1alpha1.CertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateOrder), err
}

// Delete takes name of the certificateOrder and deletes it. Returns an error if one occurs.
func (c *FakeCertificateOrders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificateordersResource, c.ns, name), &v1alpha1.CertificateOrder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateOrders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificateordersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateOrderList{})
	return err
}

// Patch applies the patch and returns the patched certificateOrder.
func (c *FakeCertificateOrders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateOrder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificateordersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CertificateOrder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateOrder), err
}
