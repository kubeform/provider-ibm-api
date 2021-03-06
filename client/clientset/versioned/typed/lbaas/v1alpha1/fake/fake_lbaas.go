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

// FakeLbaases implements LbaasInterface
type FakeLbaases struct {
	Fake *FakeLbaasV1alpha1
	ns   string
}

var lbaasesResource = schema.GroupVersionResource{Group: "lbaas.ibm.kubeform.com", Version: "v1alpha1", Resource: "lbaases"}

var lbaasesKind = schema.GroupVersionKind{Group: "lbaas.ibm.kubeform.com", Version: "v1alpha1", Kind: "Lbaas"}

// Get takes name of the lbaas, and returns the corresponding lbaas object, and an error if there is any.
func (c *FakeLbaases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Lbaas, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lbaasesResource, c.ns, name), &v1alpha1.Lbaas{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lbaas), err
}

// List takes label and field selectors, and returns the list of Lbaases that match those selectors.
func (c *FakeLbaases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LbaasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lbaasesResource, lbaasesKind, c.ns, opts), &v1alpha1.LbaasList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbaasList{ListMeta: obj.(*v1alpha1.LbaasList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbaasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbaases.
func (c *FakeLbaases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lbaasesResource, c.ns, opts))

}

// Create takes the representation of a lbaas and creates it.  Returns the server's representation of the lbaas, and an error, if there is any.
func (c *FakeLbaases) Create(ctx context.Context, lbaas *v1alpha1.Lbaas, opts v1.CreateOptions) (result *v1alpha1.Lbaas, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lbaasesResource, c.ns, lbaas), &v1alpha1.Lbaas{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lbaas), err
}

// Update takes the representation of a lbaas and updates it. Returns the server's representation of the lbaas, and an error, if there is any.
func (c *FakeLbaases) Update(ctx context.Context, lbaas *v1alpha1.Lbaas, opts v1.UpdateOptions) (result *v1alpha1.Lbaas, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lbaasesResource, c.ns, lbaas), &v1alpha1.Lbaas{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lbaas), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbaases) UpdateStatus(ctx context.Context, lbaas *v1alpha1.Lbaas, opts v1.UpdateOptions) (*v1alpha1.Lbaas, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lbaasesResource, "status", c.ns, lbaas), &v1alpha1.Lbaas{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lbaas), err
}

// Delete takes name of the lbaas and deletes it. Returns an error if one occurs.
func (c *FakeLbaases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lbaasesResource, c.ns, name), &v1alpha1.Lbaas{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbaases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lbaasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbaasList{})
	return err
}

// Patch applies the patch and returns the patched lbaas.
func (c *FakeLbaases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lbaas, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lbaasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Lbaas{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lbaas), err
}
