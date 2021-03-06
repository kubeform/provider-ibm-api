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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/network/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVlans implements VlanInterface
type FakeVlans struct {
	Fake *FakeNetworkV1alpha1
	ns   string
}

var vlansResource = schema.GroupVersionResource{Group: "network.ibm.kubeform.com", Version: "v1alpha1", Resource: "vlans"}

var vlansKind = schema.GroupVersionKind{Group: "network.ibm.kubeform.com", Version: "v1alpha1", Kind: "Vlan"}

// Get takes name of the vlan, and returns the corresponding vlan object, and an error if there is any.
func (c *FakeVlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vlansResource, c.ns, name), &v1alpha1.Vlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vlan), err
}

// List takes label and field selectors, and returns the list of Vlans that match those selectors.
func (c *FakeVlans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vlansResource, vlansKind, c.ns, opts), &v1alpha1.VlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VlanList{ListMeta: obj.(*v1alpha1.VlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.VlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vlans.
func (c *FakeVlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vlansResource, c.ns, opts))

}

// Create takes the representation of a vlan and creates it.  Returns the server's representation of the vlan, and an error, if there is any.
func (c *FakeVlans) Create(ctx context.Context, vlan *v1alpha1.Vlan, opts v1.CreateOptions) (result *v1alpha1.Vlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vlansResource, c.ns, vlan), &v1alpha1.Vlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vlan), err
}

// Update takes the representation of a vlan and updates it. Returns the server's representation of the vlan, and an error, if there is any.
func (c *FakeVlans) Update(ctx context.Context, vlan *v1alpha1.Vlan, opts v1.UpdateOptions) (result *v1alpha1.Vlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vlansResource, c.ns, vlan), &v1alpha1.Vlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVlans) UpdateStatus(ctx context.Context, vlan *v1alpha1.Vlan, opts v1.UpdateOptions) (*v1alpha1.Vlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vlansResource, "status", c.ns, vlan), &v1alpha1.Vlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vlan), err
}

// Delete takes name of the vlan and deletes it. Returns an error if one occurs.
func (c *FakeVlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vlansResource, c.ns, name), &v1alpha1.Vlan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vlansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VlanList{})
	return err
}

// Patch applies the patch and returns the patched vlan.
func (c *FakeVlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vlansResource, c.ns, name, pt, data, subresources...), &v1alpha1.Vlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vlan), err
}
