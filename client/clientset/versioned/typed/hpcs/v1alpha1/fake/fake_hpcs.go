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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/hpcs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHpcses implements HpcsInterface
type FakeHpcses struct {
	Fake *FakeHpcsV1alpha1
	ns   string
}

var hpcsesResource = schema.GroupVersionResource{Group: "hpcs.ibm.kubeform.com", Version: "v1alpha1", Resource: "hpcses"}

var hpcsesKind = schema.GroupVersionKind{Group: "hpcs.ibm.kubeform.com", Version: "v1alpha1", Kind: "Hpcs"}

// Get takes name of the hpcs, and returns the corresponding hpcs object, and an error if there is any.
func (c *FakeHpcses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Hpcs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hpcsesResource, c.ns, name), &v1alpha1.Hpcs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hpcs), err
}

// List takes label and field selectors, and returns the list of Hpcses that match those selectors.
func (c *FakeHpcses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HpcsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hpcsesResource, hpcsesKind, c.ns, opts), &v1alpha1.HpcsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HpcsList{ListMeta: obj.(*v1alpha1.HpcsList).ListMeta}
	for _, item := range obj.(*v1alpha1.HpcsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hpcses.
func (c *FakeHpcses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hpcsesResource, c.ns, opts))

}

// Create takes the representation of a hpcs and creates it.  Returns the server's representation of the hpcs, and an error, if there is any.
func (c *FakeHpcses) Create(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.CreateOptions) (result *v1alpha1.Hpcs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hpcsesResource, c.ns, hpcs), &v1alpha1.Hpcs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hpcs), err
}

// Update takes the representation of a hpcs and updates it. Returns the server's representation of the hpcs, and an error, if there is any.
func (c *FakeHpcses) Update(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.UpdateOptions) (result *v1alpha1.Hpcs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hpcsesResource, c.ns, hpcs), &v1alpha1.Hpcs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hpcs), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHpcses) UpdateStatus(ctx context.Context, hpcs *v1alpha1.Hpcs, opts v1.UpdateOptions) (*v1alpha1.Hpcs, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hpcsesResource, "status", c.ns, hpcs), &v1alpha1.Hpcs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hpcs), err
}

// Delete takes name of the hpcs and deletes it. Returns an error if one occurs.
func (c *FakeHpcses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hpcsesResource, c.ns, name), &v1alpha1.Hpcs{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHpcses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hpcsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HpcsList{})
	return err
}

// Patch applies the patch and returns the patched hpcs.
func (c *FakeHpcses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Hpcs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hpcsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Hpcs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Hpcs), err
}
