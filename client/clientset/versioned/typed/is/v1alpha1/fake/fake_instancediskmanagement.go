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

// FakeInstanceDiskManagements implements InstanceDiskManagementInterface
type FakeInstanceDiskManagements struct {
	Fake *FakeIsV1alpha1
	ns   string
}

var instancediskmanagementsResource = schema.GroupVersionResource{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Resource: "instancediskmanagements"}

var instancediskmanagementsKind = schema.GroupVersionKind{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Kind: "InstanceDiskManagement"}

// Get takes name of the instanceDiskManagement, and returns the corresponding instanceDiskManagement object, and an error if there is any.
func (c *FakeInstanceDiskManagements) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstanceDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(instancediskmanagementsResource, c.ns, name), &v1alpha1.InstanceDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceDiskManagement), err
}

// List takes label and field selectors, and returns the list of InstanceDiskManagements that match those selectors.
func (c *FakeInstanceDiskManagements) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstanceDiskManagementList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(instancediskmanagementsResource, instancediskmanagementsKind, c.ns, opts), &v1alpha1.InstanceDiskManagementList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstanceDiskManagementList{ListMeta: obj.(*v1alpha1.InstanceDiskManagementList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstanceDiskManagementList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instanceDiskManagements.
func (c *FakeInstanceDiskManagements) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(instancediskmanagementsResource, c.ns, opts))

}

// Create takes the representation of a instanceDiskManagement and creates it.  Returns the server's representation of the instanceDiskManagement, and an error, if there is any.
func (c *FakeInstanceDiskManagements) Create(ctx context.Context, instanceDiskManagement *v1alpha1.InstanceDiskManagement, opts v1.CreateOptions) (result *v1alpha1.InstanceDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(instancediskmanagementsResource, c.ns, instanceDiskManagement), &v1alpha1.InstanceDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceDiskManagement), err
}

// Update takes the representation of a instanceDiskManagement and updates it. Returns the server's representation of the instanceDiskManagement, and an error, if there is any.
func (c *FakeInstanceDiskManagements) Update(ctx context.Context, instanceDiskManagement *v1alpha1.InstanceDiskManagement, opts v1.UpdateOptions) (result *v1alpha1.InstanceDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(instancediskmanagementsResource, c.ns, instanceDiskManagement), &v1alpha1.InstanceDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceDiskManagement), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstanceDiskManagements) UpdateStatus(ctx context.Context, instanceDiskManagement *v1alpha1.InstanceDiskManagement, opts v1.UpdateOptions) (*v1alpha1.InstanceDiskManagement, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(instancediskmanagementsResource, "status", c.ns, instanceDiskManagement), &v1alpha1.InstanceDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceDiskManagement), err
}

// Delete takes name of the instanceDiskManagement and deletes it. Returns an error if one occurs.
func (c *FakeInstanceDiskManagements) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(instancediskmanagementsResource, c.ns, name), &v1alpha1.InstanceDiskManagement{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstanceDiskManagements) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(instancediskmanagementsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstanceDiskManagementList{})
	return err
}

// Patch applies the patch and returns the patched instanceDiskManagement.
func (c *FakeInstanceDiskManagements) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstanceDiskManagement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(instancediskmanagementsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstanceDiskManagement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceDiskManagement), err
}