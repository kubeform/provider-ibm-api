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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/compute/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBareMetals implements BareMetalInterface
type FakeBareMetals struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var baremetalsResource = schema.GroupVersionResource{Group: "compute.ibm.kubeform.com", Version: "v1alpha1", Resource: "baremetals"}

var baremetalsKind = schema.GroupVersionKind{Group: "compute.ibm.kubeform.com", Version: "v1alpha1", Kind: "BareMetal"}

// Get takes name of the bareMetal, and returns the corresponding bareMetal object, and an error if there is any.
func (c *FakeBareMetals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BareMetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(baremetalsResource, c.ns, name), &v1alpha1.BareMetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetal), err
}

// List takes label and field selectors, and returns the list of BareMetals that match those selectors.
func (c *FakeBareMetals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BareMetalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(baremetalsResource, baremetalsKind, c.ns, opts), &v1alpha1.BareMetalList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BareMetalList{ListMeta: obj.(*v1alpha1.BareMetalList).ListMeta}
	for _, item := range obj.(*v1alpha1.BareMetalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bareMetals.
func (c *FakeBareMetals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(baremetalsResource, c.ns, opts))

}

// Create takes the representation of a bareMetal and creates it.  Returns the server's representation of the bareMetal, and an error, if there is any.
func (c *FakeBareMetals) Create(ctx context.Context, bareMetal *v1alpha1.BareMetal, opts v1.CreateOptions) (result *v1alpha1.BareMetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(baremetalsResource, c.ns, bareMetal), &v1alpha1.BareMetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetal), err
}

// Update takes the representation of a bareMetal and updates it. Returns the server's representation of the bareMetal, and an error, if there is any.
func (c *FakeBareMetals) Update(ctx context.Context, bareMetal *v1alpha1.BareMetal, opts v1.UpdateOptions) (result *v1alpha1.BareMetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(baremetalsResource, c.ns, bareMetal), &v1alpha1.BareMetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetal), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBareMetals) UpdateStatus(ctx context.Context, bareMetal *v1alpha1.BareMetal, opts v1.UpdateOptions) (*v1alpha1.BareMetal, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(baremetalsResource, "status", c.ns, bareMetal), &v1alpha1.BareMetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetal), err
}

// Delete takes name of the bareMetal and deletes it. Returns an error if one occurs.
func (c *FakeBareMetals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(baremetalsResource, c.ns, name), &v1alpha1.BareMetal{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBareMetals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(baremetalsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BareMetalList{})
	return err
}

// Patch applies the patch and returns the patched bareMetal.
func (c *FakeBareMetals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BareMetal, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(baremetalsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BareMetal{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetal), err
}