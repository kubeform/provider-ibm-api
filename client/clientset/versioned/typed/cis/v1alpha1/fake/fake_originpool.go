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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/cis/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOriginPools implements OriginPoolInterface
type FakeOriginPools struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var originpoolsResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "originpools"}

var originpoolsKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "OriginPool"}

// Get takes name of the originPool, and returns the corresponding originPool object, and an error if there is any.
func (c *FakeOriginPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OriginPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(originpoolsResource, c.ns, name), &v1alpha1.OriginPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OriginPool), err
}

// List takes label and field selectors, and returns the list of OriginPools that match those selectors.
func (c *FakeOriginPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OriginPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(originpoolsResource, originpoolsKind, c.ns, opts), &v1alpha1.OriginPoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OriginPoolList{ListMeta: obj.(*v1alpha1.OriginPoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.OriginPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested originPools.
func (c *FakeOriginPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(originpoolsResource, c.ns, opts))

}

// Create takes the representation of a originPool and creates it.  Returns the server's representation of the originPool, and an error, if there is any.
func (c *FakeOriginPools) Create(ctx context.Context, originPool *v1alpha1.OriginPool, opts v1.CreateOptions) (result *v1alpha1.OriginPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(originpoolsResource, c.ns, originPool), &v1alpha1.OriginPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OriginPool), err
}

// Update takes the representation of a originPool and updates it. Returns the server's representation of the originPool, and an error, if there is any.
func (c *FakeOriginPools) Update(ctx context.Context, originPool *v1alpha1.OriginPool, opts v1.UpdateOptions) (result *v1alpha1.OriginPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(originpoolsResource, c.ns, originPool), &v1alpha1.OriginPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OriginPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOriginPools) UpdateStatus(ctx context.Context, originPool *v1alpha1.OriginPool, opts v1.UpdateOptions) (*v1alpha1.OriginPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(originpoolsResource, "status", c.ns, originPool), &v1alpha1.OriginPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OriginPool), err
}

// Delete takes name of the originPool and deletes it. Returns an error if one occurs.
func (c *FakeOriginPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(originpoolsResource, c.ns, name), &v1alpha1.OriginPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOriginPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(originpoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OriginPoolList{})
	return err
}

// Patch applies the patch and returns the patched originPool.
func (c *FakeOriginPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OriginPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(originpoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OriginPool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OriginPool), err
}
