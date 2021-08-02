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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/lb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpxServices implements VpxServiceInterface
type FakeVpxServices struct {
	Fake *FakeLbV1alpha1
	ns   string
}

var vpxservicesResource = schema.GroupVersionResource{Group: "lb.ibm.kubeform.com", Version: "v1alpha1", Resource: "vpxservices"}

var vpxservicesKind = schema.GroupVersionKind{Group: "lb.ibm.kubeform.com", Version: "v1alpha1", Kind: "VpxService"}

// Get takes name of the vpxService, and returns the corresponding vpxService object, and an error if there is any.
func (c *FakeVpxServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpxService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpxservicesResource, c.ns, name), &v1alpha1.VpxService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpxService), err
}

// List takes label and field selectors, and returns the list of VpxServices that match those selectors.
func (c *FakeVpxServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpxServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpxservicesResource, vpxservicesKind, c.ns, opts), &v1alpha1.VpxServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpxServiceList{ListMeta: obj.(*v1alpha1.VpxServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpxServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpxServices.
func (c *FakeVpxServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpxservicesResource, c.ns, opts))

}

// Create takes the representation of a vpxService and creates it.  Returns the server's representation of the vpxService, and an error, if there is any.
func (c *FakeVpxServices) Create(ctx context.Context, vpxService *v1alpha1.VpxService, opts v1.CreateOptions) (result *v1alpha1.VpxService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpxservicesResource, c.ns, vpxService), &v1alpha1.VpxService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpxService), err
}

// Update takes the representation of a vpxService and updates it. Returns the server's representation of the vpxService, and an error, if there is any.
func (c *FakeVpxServices) Update(ctx context.Context, vpxService *v1alpha1.VpxService, opts v1.UpdateOptions) (result *v1alpha1.VpxService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpxservicesResource, c.ns, vpxService), &v1alpha1.VpxService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpxService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpxServices) UpdateStatus(ctx context.Context, vpxService *v1alpha1.VpxService, opts v1.UpdateOptions) (*v1alpha1.VpxService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpxservicesResource, "status", c.ns, vpxService), &v1alpha1.VpxService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpxService), err
}

// Delete takes name of the vpxService and deletes it. Returns an error if one occurs.
func (c *FakeVpxServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpxservicesResource, c.ns, name), &v1alpha1.VpxService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpxServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpxservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpxServiceList{})
	return err
}

// Patch applies the patch and returns the patched vpxService.
func (c *FakeVpxServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpxService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpxservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpxService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpxService), err
}