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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/is/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIkePolicies implements IkePolicyInterface
type FakeIkePolicies struct {
	Fake *FakeIsV1alpha1
	ns   string
}

var ikepoliciesResource = schema.GroupVersionResource{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Resource: "ikepolicies"}

var ikepoliciesKind = schema.GroupVersionKind{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Kind: "IkePolicy"}

// Get takes name of the ikePolicy, and returns the corresponding ikePolicy object, and an error if there is any.
func (c *FakeIkePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IkePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ikepoliciesResource, c.ns, name), &v1alpha1.IkePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IkePolicy), err
}

// List takes label and field selectors, and returns the list of IkePolicies that match those selectors.
func (c *FakeIkePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IkePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ikepoliciesResource, ikepoliciesKind, c.ns, opts), &v1alpha1.IkePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IkePolicyList{ListMeta: obj.(*v1alpha1.IkePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IkePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ikePolicies.
func (c *FakeIkePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ikepoliciesResource, c.ns, opts))

}

// Create takes the representation of a ikePolicy and creates it.  Returns the server's representation of the ikePolicy, and an error, if there is any.
func (c *FakeIkePolicies) Create(ctx context.Context, ikePolicy *v1alpha1.IkePolicy, opts v1.CreateOptions) (result *v1alpha1.IkePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ikepoliciesResource, c.ns, ikePolicy), &v1alpha1.IkePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IkePolicy), err
}

// Update takes the representation of a ikePolicy and updates it. Returns the server's representation of the ikePolicy, and an error, if there is any.
func (c *FakeIkePolicies) Update(ctx context.Context, ikePolicy *v1alpha1.IkePolicy, opts v1.UpdateOptions) (result *v1alpha1.IkePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ikepoliciesResource, c.ns, ikePolicy), &v1alpha1.IkePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IkePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIkePolicies) UpdateStatus(ctx context.Context, ikePolicy *v1alpha1.IkePolicy, opts v1.UpdateOptions) (*v1alpha1.IkePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ikepoliciesResource, "status", c.ns, ikePolicy), &v1alpha1.IkePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IkePolicy), err
}

// Delete takes name of the ikePolicy and deletes it. Returns an error if one occurs.
func (c *FakeIkePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ikepoliciesResource, c.ns, name), &v1alpha1.IkePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIkePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ikepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IkePolicyList{})
	return err
}

// Patch applies the patch and returns the patched ikePolicy.
func (c *FakeIkePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IkePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ikepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IkePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IkePolicy), err
}
