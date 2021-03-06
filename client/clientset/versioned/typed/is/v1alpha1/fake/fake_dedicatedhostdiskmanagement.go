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

// FakeDedicatedHostDiskManagements implements DedicatedHostDiskManagementInterface
type FakeDedicatedHostDiskManagements struct {
	Fake *FakeIsV1alpha1
	ns   string
}

var dedicatedhostdiskmanagementsResource = schema.GroupVersionResource{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Resource: "dedicatedhostdiskmanagements"}

var dedicatedhostdiskmanagementsKind = schema.GroupVersionKind{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Kind: "DedicatedHostDiskManagement"}

// Get takes name of the dedicatedHostDiskManagement, and returns the corresponding dedicatedHostDiskManagement object, and an error if there is any.
func (c *FakeDedicatedHostDiskManagements) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DedicatedHostDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dedicatedhostdiskmanagementsResource, c.ns, name), &v1alpha1.DedicatedHostDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostDiskManagement), err
}

// List takes label and field selectors, and returns the list of DedicatedHostDiskManagements that match those selectors.
func (c *FakeDedicatedHostDiskManagements) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DedicatedHostDiskManagementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dedicatedhostdiskmanagementsResource, dedicatedhostdiskmanagementsKind, c.ns, opts), &v1alpha1.DedicatedHostDiskManagementList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DedicatedHostDiskManagementList{ListMeta: obj.(*v1alpha1.DedicatedHostDiskManagementList).ListMeta}
	for _, item := range obj.(*v1alpha1.DedicatedHostDiskManagementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dedicatedHostDiskManagements.
func (c *FakeDedicatedHostDiskManagements) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dedicatedhostdiskmanagementsResource, c.ns, opts))

}

// Create takes the representation of a dedicatedHostDiskManagement and creates it.  Returns the server's representation of the dedicatedHostDiskManagement, and an error, if there is any.
func (c *FakeDedicatedHostDiskManagements) Create(ctx context.Context, dedicatedHostDiskManagement *v1alpha1.DedicatedHostDiskManagement, opts v1.CreateOptions) (result *v1alpha1.DedicatedHostDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dedicatedhostdiskmanagementsResource, c.ns, dedicatedHostDiskManagement), &v1alpha1.DedicatedHostDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostDiskManagement), err
}

// Update takes the representation of a dedicatedHostDiskManagement and updates it. Returns the server's representation of the dedicatedHostDiskManagement, and an error, if there is any.
func (c *FakeDedicatedHostDiskManagements) Update(ctx context.Context, dedicatedHostDiskManagement *v1alpha1.DedicatedHostDiskManagement, opts v1.UpdateOptions) (result *v1alpha1.DedicatedHostDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dedicatedhostdiskmanagementsResource, c.ns, dedicatedHostDiskManagement), &v1alpha1.DedicatedHostDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostDiskManagement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDedicatedHostDiskManagements) UpdateStatus(ctx context.Context, dedicatedHostDiskManagement *v1alpha1.DedicatedHostDiskManagement, opts v1.UpdateOptions) (*v1alpha1.DedicatedHostDiskManagement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dedicatedhostdiskmanagementsResource, "status", c.ns, dedicatedHostDiskManagement), &v1alpha1.DedicatedHostDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostDiskManagement), err
}

// Delete takes name of the dedicatedHostDiskManagement and deletes it. Returns an error if one occurs.
func (c *FakeDedicatedHostDiskManagements) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dedicatedhostdiskmanagementsResource, c.ns, name), &v1alpha1.DedicatedHostDiskManagement{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDedicatedHostDiskManagements) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dedicatedhostdiskmanagementsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DedicatedHostDiskManagementList{})
	return err
}

// Patch applies the patch and returns the patched dedicatedHostDiskManagement.
func (c *FakeDedicatedHostDiskManagements) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DedicatedHostDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dedicatedhostdiskmanagementsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DedicatedHostDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DedicatedHostDiskManagement), err
}
