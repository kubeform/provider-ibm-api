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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/iam/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserInvites implements UserInviteInterface
type FakeUserInvites struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var userinvitesResource = schema.GroupVersionResource{Group: "iam.ibm.kubeform.com", Version: "v1alpha1", Resource: "userinvites"}

var userinvitesKind = schema.GroupVersionKind{Group: "iam.ibm.kubeform.com", Version: "v1alpha1", Kind: "UserInvite"}

// Get takes name of the userInvite, and returns the corresponding userInvite object, and an error if there is any.
func (c *FakeUserInvites) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserInvite, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userinvitesResource, c.ns, name), &v1alpha1.UserInvite{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserInvite), err
}

// List takes label and field selectors, and returns the list of UserInvites that match those selectors.
func (c *FakeUserInvites) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserInviteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userinvitesResource, userinvitesKind, c.ns, opts), &v1alpha1.UserInviteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserInviteList{ListMeta: obj.(*v1alpha1.UserInviteList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserInviteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userInvites.
func (c *FakeUserInvites) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userinvitesResource, c.ns, opts))

}

// Create takes the representation of a userInvite and creates it.  Returns the server's representation of the userInvite, and an error, if there is any.
func (c *FakeUserInvites) Create(ctx context.Context, userInvite *v1alpha1.UserInvite, opts v1.CreateOptions) (result *v1alpha1.UserInvite, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userinvitesResource, c.ns, userInvite), &v1alpha1.UserInvite{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserInvite), err
}

// Update takes the representation of a userInvite and updates it. Returns the server's representation of the userInvite, and an error, if there is any.
func (c *FakeUserInvites) Update(ctx context.Context, userInvite *v1alpha1.UserInvite, opts v1.UpdateOptions) (result *v1alpha1.UserInvite, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userinvitesResource, c.ns, userInvite), &v1alpha1.UserInvite{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserInvite), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserInvites) UpdateStatus(ctx context.Context, userInvite *v1alpha1.UserInvite, opts v1.UpdateOptions) (*v1alpha1.UserInvite, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userinvitesResource, "status", c.ns, userInvite), &v1alpha1.UserInvite{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserInvite), err
}

// Delete takes name of the userInvite and deletes it. Returns an error if one occurs.
func (c *FakeUserInvites) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userinvitesResource, c.ns, name), &v1alpha1.UserInvite{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserInvites) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userinvitesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserInviteList{})
	return err
}

// Patch applies the patch and returns the patched userInvite.
func (c *FakeUserInvites) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserInvite, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userinvitesResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserInvite{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserInvite), err
}
