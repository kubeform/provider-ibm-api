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

// FakeBindServices implements BindServiceInterface
type FakeBindServices struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var bindservicesResource = schema.GroupVersionResource{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Resource: "bindservices"}

var bindservicesKind = schema.GroupVersionKind{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Kind: "BindService"}

// Get takes name of the bindService, and returns the corresponding bindService object, and an error if there is any.
func (c *FakeBindServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BindService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bindservicesResource, c.ns, name), &v1alpha1.BindService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BindService), err
}

// List takes label and field selectors, and returns the list of BindServices that match those selectors.
func (c *FakeBindServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BindServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bindservicesResource, bindservicesKind, c.ns, opts), &v1alpha1.BindServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BindServiceList{ListMeta: obj.(*v1alpha1.BindServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.BindServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bindServices.
func (c *FakeBindServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bindservicesResource, c.ns, opts))

}

// Create takes the representation of a bindService and creates it.  Returns the server's representation of the bindService, and an error, if there is any.
func (c *FakeBindServices) Create(ctx context.Context, bindService *v1alpha1.BindService, opts v1.CreateOptions) (result *v1alpha1.BindService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bindservicesResource, c.ns, bindService), &v1alpha1.BindService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BindService), err
}

// Update takes the representation of a bindService and updates it. Returns the server's representation of the bindService, and an error, if there is any.
func (c *FakeBindServices) Update(ctx context.Context, bindService *v1alpha1.BindService, opts v1.UpdateOptions) (result *v1alpha1.BindService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bindservicesResource, c.ns, bindService), &v1alpha1.BindService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BindService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBindServices) UpdateStatus(ctx context.Context, bindService *v1alpha1.BindService, opts v1.UpdateOptions) (*v1alpha1.BindService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bindservicesResource, "status", c.ns, bindService), &v1alpha1.BindService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BindService), err
}

// Delete takes name of the bindService and deletes it. Returns an error if one occurs.
func (c *FakeBindServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bindservicesResource, c.ns, name), &v1alpha1.BindService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBindServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bindservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BindServiceList{})
	return err
}

// Patch applies the patch and returns the patched bindService.
func (c *FakeBindServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BindService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bindservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BindService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BindService), err
}
