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

// FakeRangeApps implements RangeAppInterface
type FakeRangeApps struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var rangeappsResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "rangeapps"}

var rangeappsKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "RangeApp"}

// Get takes name of the rangeApp, and returns the corresponding rangeApp object, and an error if there is any.
func (c *FakeRangeApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RangeApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rangeappsResource, c.ns, name), &v1alpha1.RangeApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RangeApp), err
}

// List takes label and field selectors, and returns the list of RangeApps that match those selectors.
func (c *FakeRangeApps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RangeAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rangeappsResource, rangeappsKind, c.ns, opts), &v1alpha1.RangeAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RangeAppList{ListMeta: obj.(*v1alpha1.RangeAppList).ListMeta}
	for _, item := range obj.(*v1alpha1.RangeAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rangeApps.
func (c *FakeRangeApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rangeappsResource, c.ns, opts))

}

// Create takes the representation of a rangeApp and creates it.  Returns the server's representation of the rangeApp, and an error, if there is any.
func (c *FakeRangeApps) Create(ctx context.Context, rangeApp *v1alpha1.RangeApp, opts v1.CreateOptions) (result *v1alpha1.RangeApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rangeappsResource, c.ns, rangeApp), &v1alpha1.RangeApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RangeApp), err
}

// Update takes the representation of a rangeApp and updates it. Returns the server's representation of the rangeApp, and an error, if there is any.
func (c *FakeRangeApps) Update(ctx context.Context, rangeApp *v1alpha1.RangeApp, opts v1.UpdateOptions) (result *v1alpha1.RangeApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rangeappsResource, c.ns, rangeApp), &v1alpha1.RangeApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RangeApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRangeApps) UpdateStatus(ctx context.Context, rangeApp *v1alpha1.RangeApp, opts v1.UpdateOptions) (*v1alpha1.RangeApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rangeappsResource, "status", c.ns, rangeApp), &v1alpha1.RangeApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RangeApp), err
}

// Delete takes name of the rangeApp and deletes it. Returns an error if one occurs.
func (c *FakeRangeApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rangeappsResource, c.ns, name), &v1alpha1.RangeApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRangeApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rangeappsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RangeAppList{})
	return err
}

// Patch applies the patch and returns the patched rangeApp.
func (c *FakeRangeApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RangeApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rangeappsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RangeApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RangeApp), err
}