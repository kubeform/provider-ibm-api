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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/lbaas/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServerInstanceAttachments implements ServerInstanceAttachmentInterface
type FakeServerInstanceAttachments struct {
	Fake *FakeLbaasV1alpha1
	ns   string
}

var serverinstanceattachmentsResource = schema.GroupVersionResource{Group: "lbaas.ibm.kubeform.com", Version: "v1alpha1", Resource: "serverinstanceattachments"}

var serverinstanceattachmentsKind = schema.GroupVersionKind{Group: "lbaas.ibm.kubeform.com", Version: "v1alpha1", Kind: "ServerInstanceAttachment"}

// Get takes name of the serverInstanceAttachment, and returns the corresponding serverInstanceAttachment object, and an error if there is any.
func (c *FakeServerInstanceAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServerInstanceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serverinstanceattachmentsResource, c.ns, name), &v1alpha1.ServerInstanceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerInstanceAttachment), err
}

// List takes label and field selectors, and returns the list of ServerInstanceAttachments that match those selectors.
func (c *FakeServerInstanceAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServerInstanceAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serverinstanceattachmentsResource, serverinstanceattachmentsKind, c.ns, opts), &v1alpha1.ServerInstanceAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServerInstanceAttachmentList{ListMeta: obj.(*v1alpha1.ServerInstanceAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServerInstanceAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serverInstanceAttachments.
func (c *FakeServerInstanceAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serverinstanceattachmentsResource, c.ns, opts))

}

// Create takes the representation of a serverInstanceAttachment and creates it.  Returns the server's representation of the serverInstanceAttachment, and an error, if there is any.
func (c *FakeServerInstanceAttachments) Create(ctx context.Context, serverInstanceAttachment *v1alpha1.ServerInstanceAttachment, opts v1.CreateOptions) (result *v1alpha1.ServerInstanceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serverinstanceattachmentsResource, c.ns, serverInstanceAttachment), &v1alpha1.ServerInstanceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerInstanceAttachment), err
}

// Update takes the representation of a serverInstanceAttachment and updates it. Returns the server's representation of the serverInstanceAttachment, and an error, if there is any.
func (c *FakeServerInstanceAttachments) Update(ctx context.Context, serverInstanceAttachment *v1alpha1.ServerInstanceAttachment, opts v1.UpdateOptions) (result *v1alpha1.ServerInstanceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serverinstanceattachmentsResource, c.ns, serverInstanceAttachment), &v1alpha1.ServerInstanceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerInstanceAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServerInstanceAttachments) UpdateStatus(ctx context.Context, serverInstanceAttachment *v1alpha1.ServerInstanceAttachment, opts v1.UpdateOptions) (*v1alpha1.ServerInstanceAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serverinstanceattachmentsResource, "status", c.ns, serverInstanceAttachment), &v1alpha1.ServerInstanceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerInstanceAttachment), err
}

// Delete takes name of the serverInstanceAttachment and deletes it. Returns an error if one occurs.
func (c *FakeServerInstanceAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serverinstanceattachmentsResource, c.ns, name), &v1alpha1.ServerInstanceAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServerInstanceAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serverinstanceattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServerInstanceAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched serverInstanceAttachment.
func (c *FakeServerInstanceAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServerInstanceAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serverinstanceattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServerInstanceAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerInstanceAttachment), err
}
