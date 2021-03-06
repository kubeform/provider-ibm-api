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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/object/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageAccounts implements StorageAccountInterface
type FakeStorageAccounts struct {
	Fake *FakeObjectV1alpha1
	ns   string
}

var storageaccountsResource = schema.GroupVersionResource{Group: "object.ibm.kubeform.com", Version: "v1alpha1", Resource: "storageaccounts"}

var storageaccountsKind = schema.GroupVersionKind{Group: "object.ibm.kubeform.com", Version: "v1alpha1", Kind: "StorageAccount"}

// Get takes name of the storageAccount, and returns the corresponding storageAccount object, and an error if there is any.
func (c *FakeStorageAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storageaccountsResource, c.ns, name), &v1alpha1.StorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccount), err
}

// List takes label and field selectors, and returns the list of StorageAccounts that match those selectors.
func (c *FakeStorageAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageAccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storageaccountsResource, storageaccountsKind, c.ns, opts), &v1alpha1.StorageAccountList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageAccountList{ListMeta: obj.(*v1alpha1.StorageAccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageAccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageAccounts.
func (c *FakeStorageAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storageaccountsResource, c.ns, opts))

}

// Create takes the representation of a storageAccount and creates it.  Returns the server's representation of the storageAccount, and an error, if there is any.
func (c *FakeStorageAccounts) Create(ctx context.Context, storageAccount *v1alpha1.StorageAccount, opts v1.CreateOptions) (result *v1alpha1.StorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storageaccountsResource, c.ns, storageAccount), &v1alpha1.StorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccount), err
}

// Update takes the representation of a storageAccount and updates it. Returns the server's representation of the storageAccount, and an error, if there is any.
func (c *FakeStorageAccounts) Update(ctx context.Context, storageAccount *v1alpha1.StorageAccount, opts v1.UpdateOptions) (result *v1alpha1.StorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storageaccountsResource, c.ns, storageAccount), &v1alpha1.StorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccount), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageAccounts) UpdateStatus(ctx context.Context, storageAccount *v1alpha1.StorageAccount, opts v1.UpdateOptions) (*v1alpha1.StorageAccount, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storageaccountsResource, "status", c.ns, storageAccount), &v1alpha1.StorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccount), err
}

// Delete takes name of the storageAccount and deletes it. Returns an error if one occurs.
func (c *FakeStorageAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storageaccountsResource, c.ns, name), &v1alpha1.StorageAccount{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storageaccountsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageAccountList{})
	return err
}

// Patch applies the patch and returns the patched storageAccount.
func (c *FakeStorageAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageAccount, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storageaccountsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageAccount{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageAccount), err
}
