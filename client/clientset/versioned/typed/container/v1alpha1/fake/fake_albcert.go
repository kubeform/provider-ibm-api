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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/container/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlbCerts implements AlbCertInterface
type FakeAlbCerts struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var albcertsResource = schema.GroupVersionResource{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Resource: "albcerts"}

var albcertsKind = schema.GroupVersionKind{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Kind: "AlbCert"}

// Get takes name of the albCert, and returns the corresponding albCert object, and an error if there is any.
func (c *FakeAlbCerts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlbCert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(albcertsResource, c.ns, name), &v1alpha1.AlbCert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbCert), err
}

// List takes label and field selectors, and returns the list of AlbCerts that match those selectors.
func (c *FakeAlbCerts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlbCertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(albcertsResource, albcertsKind, c.ns, opts), &v1alpha1.AlbCertList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlbCertList{ListMeta: obj.(*v1alpha1.AlbCertList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlbCertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested albCerts.
func (c *FakeAlbCerts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(albcertsResource, c.ns, opts))

}

// Create takes the representation of a albCert and creates it.  Returns the server's representation of the albCert, and an error, if there is any.
func (c *FakeAlbCerts) Create(ctx context.Context, albCert *v1alpha1.AlbCert, opts v1.CreateOptions) (result *v1alpha1.AlbCert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(albcertsResource, c.ns, albCert), &v1alpha1.AlbCert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbCert), err
}

// Update takes the representation of a albCert and updates it. Returns the server's representation of the albCert, and an error, if there is any.
func (c *FakeAlbCerts) Update(ctx context.Context, albCert *v1alpha1.AlbCert, opts v1.UpdateOptions) (result *v1alpha1.AlbCert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(albcertsResource, c.ns, albCert), &v1alpha1.AlbCert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbCert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlbCerts) UpdateStatus(ctx context.Context, albCert *v1alpha1.AlbCert, opts v1.UpdateOptions) (*v1alpha1.AlbCert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(albcertsResource, "status", c.ns, albCert), &v1alpha1.AlbCert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbCert), err
}

// Delete takes name of the albCert and deletes it. Returns an error if one occurs.
func (c *FakeAlbCerts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(albcertsResource, c.ns, name), &v1alpha1.AlbCert{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlbCerts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(albcertsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlbCertList{})
	return err
}

// Patch applies the patch and returns the patched albCert.
func (c *FakeAlbCerts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlbCert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(albcertsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlbCert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlbCert), err
}
