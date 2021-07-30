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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/multi/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVlanFirewalls implements VlanFirewallInterface
type FakeVlanFirewalls struct {
	Fake *FakeMultiV1alpha1
	ns   string
}

var vlanfirewallsResource = schema.GroupVersionResource{Group: "multi.ibm.kubeform.com", Version: "v1alpha1", Resource: "vlanfirewalls"}

var vlanfirewallsKind = schema.GroupVersionKind{Group: "multi.ibm.kubeform.com", Version: "v1alpha1", Kind: "VlanFirewall"}

// Get takes name of the vlanFirewall, and returns the corresponding vlanFirewall object, and an error if there is any.
func (c *FakeVlanFirewalls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VlanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vlanfirewallsResource, c.ns, name), &v1alpha1.VlanFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VlanFirewall), err
}

// List takes label and field selectors, and returns the list of VlanFirewalls that match those selectors.
func (c *FakeVlanFirewalls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VlanFirewallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vlanfirewallsResource, vlanfirewallsKind, c.ns, opts), &v1alpha1.VlanFirewallList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VlanFirewallList{ListMeta: obj.(*v1alpha1.VlanFirewallList).ListMeta}
	for _, item := range obj.(*v1alpha1.VlanFirewallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vlanFirewalls.
func (c *FakeVlanFirewalls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vlanfirewallsResource, c.ns, opts))

}

// Create takes the representation of a vlanFirewall and creates it.  Returns the server's representation of the vlanFirewall, and an error, if there is any.
func (c *FakeVlanFirewalls) Create(ctx context.Context, vlanFirewall *v1alpha1.VlanFirewall, opts v1.CreateOptions) (result *v1alpha1.VlanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vlanfirewallsResource, c.ns, vlanFirewall), &v1alpha1.VlanFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VlanFirewall), err
}

// Update takes the representation of a vlanFirewall and updates it. Returns the server's representation of the vlanFirewall, and an error, if there is any.
func (c *FakeVlanFirewalls) Update(ctx context.Context, vlanFirewall *v1alpha1.VlanFirewall, opts v1.UpdateOptions) (result *v1alpha1.VlanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vlanfirewallsResource, c.ns, vlanFirewall), &v1alpha1.VlanFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VlanFirewall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVlanFirewalls) UpdateStatus(ctx context.Context, vlanFirewall *v1alpha1.VlanFirewall, opts v1.UpdateOptions) (*v1alpha1.VlanFirewall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vlanfirewallsResource, "status", c.ns, vlanFirewall), &v1alpha1.VlanFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VlanFirewall), err
}

// Delete takes name of the vlanFirewall and deletes it. Returns an error if one occurs.
func (c *FakeVlanFirewalls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vlanfirewallsResource, c.ns, name), &v1alpha1.VlanFirewall{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVlanFirewalls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vlanfirewallsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VlanFirewallList{})
	return err
}

// Patch applies the patch and returns the patched vlanFirewall.
func (c *FakeVlanFirewalls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VlanFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vlanfirewallsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VlanFirewall{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VlanFirewall), err
}
