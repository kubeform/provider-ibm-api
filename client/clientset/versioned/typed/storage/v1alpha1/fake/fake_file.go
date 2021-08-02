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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/storage/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFiles implements FileInterface
type FakeFiles struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var filesResource = schema.GroupVersionResource{Group: "storage.ibm.kubeform.com", Version: "v1alpha1", Resource: "files"}

var filesKind = schema.GroupVersionKind{Group: "storage.ibm.kubeform.com", Version: "v1alpha1", Kind: "File"}

// Get takes name of the file, and returns the corresponding file object, and an error if there is any.
func (c *FakeFiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.File, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(filesResource, c.ns, name), &v1alpha1.File{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.File), err
}

// List takes label and field selectors, and returns the list of Files that match those selectors.
func (c *FakeFiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(filesResource, filesKind, c.ns, opts), &v1alpha1.FileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FileList{ListMeta: obj.(*v1alpha1.FileList).ListMeta}
	for _, item := range obj.(*v1alpha1.FileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested files.
func (c *FakeFiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(filesResource, c.ns, opts))

}

// Create takes the representation of a file and creates it.  Returns the server's representation of the file, and an error, if there is any.
func (c *FakeFiles) Create(ctx context.Context, file *v1alpha1.File, opts v1.CreateOptions) (result *v1alpha1.File, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(filesResource, c.ns, file), &v1alpha1.File{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.File), err
}

// Update takes the representation of a file and updates it. Returns the server's representation of the file, and an error, if there is any.
func (c *FakeFiles) Update(ctx context.Context, file *v1alpha1.File, opts v1.UpdateOptions) (result *v1alpha1.File, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(filesResource, c.ns, file), &v1alpha1.File{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.File), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFiles) UpdateStatus(ctx context.Context, file *v1alpha1.File, opts v1.UpdateOptions) (*v1alpha1.File, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(filesResource, "status", c.ns, file), &v1alpha1.File{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.File), err
}

// Delete takes name of the file and deletes it. Returns an error if one occurs.
func (c *FakeFiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(filesResource, c.ns, name), &v1alpha1.File{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(filesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FileList{})
	return err
}

// Patch applies the patch and returns the patched file.
func (c *FakeFiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.File, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(filesResource, c.ns, name, pt, data, subresources...), &v1alpha1.File{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.File), err
}