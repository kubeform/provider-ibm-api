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

// FakeAddonses implements AddonsInterface
type FakeAddonses struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var addonsesResource = schema.GroupVersionResource{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Resource: "addonses"}

var addonsesKind = schema.GroupVersionKind{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Kind: "Addons"}

// Get takes name of the addons, and returns the corresponding addons object, and an error if there is any.
func (c *FakeAddonses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Addons, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(addonsesResource, c.ns, name), &v1alpha1.Addons{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Addons), err
}

// List takes label and field selectors, and returns the list of Addonses that match those selectors.
func (c *FakeAddonses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddonsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(addonsesResource, addonsesKind, c.ns, opts), &v1alpha1.AddonsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AddonsList{ListMeta: obj.(*v1alpha1.AddonsList).ListMeta}
	for _, item := range obj.(*v1alpha1.AddonsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addonses.
func (c *FakeAddonses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(addonsesResource, c.ns, opts))

}

// Create takes the representation of a addons and creates it.  Returns the server's representation of the addons, and an error, if there is any.
func (c *FakeAddonses) Create(ctx context.Context, addons *v1alpha1.Addons, opts v1.CreateOptions) (result *v1alpha1.Addons, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(addonsesResource, c.ns, addons), &v1alpha1.Addons{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Addons), err
}

// Update takes the representation of a addons and updates it. Returns the server's representation of the addons, and an error, if there is any.
func (c *FakeAddonses) Update(ctx context.Context, addons *v1alpha1.Addons, opts v1.UpdateOptions) (result *v1alpha1.Addons, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(addonsesResource, c.ns, addons), &v1alpha1.Addons{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Addons), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAddonses) UpdateStatus(ctx context.Context, addons *v1alpha1.Addons, opts v1.UpdateOptions) (*v1alpha1.Addons, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(addonsesResource, "status", c.ns, addons), &v1alpha1.Addons{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Addons), err
}

// Delete takes name of the addons and deletes it. Returns an error if one occurs.
func (c *FakeAddonses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(addonsesResource, c.ns, name), &v1alpha1.Addons{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddonses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(addonsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AddonsList{})
	return err
}

// Patch applies the patch and returns the patched addons.
func (c *FakeAddonses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Addons, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(addonsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Addons{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Addons), err
}
