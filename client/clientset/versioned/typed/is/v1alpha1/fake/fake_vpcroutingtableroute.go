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

// FakeVpcRoutingTableRoutes implements VpcRoutingTableRouteInterface
type FakeVpcRoutingTableRoutes struct {
	Fake *FakeIsV1alpha1
	ns   string
}

var vpcroutingtableroutesResource = schema.GroupVersionResource{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Resource: "vpcroutingtableroutes"}

var vpcroutingtableroutesKind = schema.GroupVersionKind{Group: "is.ibm.kubeform.com", Version: "v1alpha1", Kind: "VpcRoutingTableRoute"}

// Get takes name of the vpcRoutingTableRoute, and returns the corresponding vpcRoutingTableRoute object, and an error if there is any.
func (c *FakeVpcRoutingTableRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpcRoutingTableRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcroutingtableroutesResource, c.ns, name), &v1alpha1.VpcRoutingTableRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcRoutingTableRoute), err
}

// List takes label and field selectors, and returns the list of VpcRoutingTableRoutes that match those selectors.
func (c *FakeVpcRoutingTableRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcRoutingTableRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcroutingtableroutesResource, vpcroutingtableroutesKind, c.ns, opts), &v1alpha1.VpcRoutingTableRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcRoutingTableRouteList{ListMeta: obj.(*v1alpha1.VpcRoutingTableRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcRoutingTableRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcRoutingTableRoutes.
func (c *FakeVpcRoutingTableRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcroutingtableroutesResource, c.ns, opts))

}

// Create takes the representation of a vpcRoutingTableRoute and creates it.  Returns the server's representation of the vpcRoutingTableRoute, and an error, if there is any.
func (c *FakeVpcRoutingTableRoutes) Create(ctx context.Context, vpcRoutingTableRoute *v1alpha1.VpcRoutingTableRoute, opts v1.CreateOptions) (result *v1alpha1.VpcRoutingTableRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcroutingtableroutesResource, c.ns, vpcRoutingTableRoute), &v1alpha1.VpcRoutingTableRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcRoutingTableRoute), err
}

// Update takes the representation of a vpcRoutingTableRoute and updates it. Returns the server's representation of the vpcRoutingTableRoute, and an error, if there is any.
func (c *FakeVpcRoutingTableRoutes) Update(ctx context.Context, vpcRoutingTableRoute *v1alpha1.VpcRoutingTableRoute, opts v1.UpdateOptions) (result *v1alpha1.VpcRoutingTableRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcroutingtableroutesResource, c.ns, vpcRoutingTableRoute), &v1alpha1.VpcRoutingTableRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcRoutingTableRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcRoutingTableRoutes) UpdateStatus(ctx context.Context, vpcRoutingTableRoute *v1alpha1.VpcRoutingTableRoute, opts v1.UpdateOptions) (*v1alpha1.VpcRoutingTableRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcroutingtableroutesResource, "status", c.ns, vpcRoutingTableRoute), &v1alpha1.VpcRoutingTableRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcRoutingTableRoute), err
}

// Delete takes name of the vpcRoutingTableRoute and deletes it. Returns an error if one occurs.
func (c *FakeVpcRoutingTableRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcroutingtableroutesResource, c.ns, name), &v1alpha1.VpcRoutingTableRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcRoutingTableRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcroutingtableroutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcRoutingTableRouteList{})
	return err
}

// Patch applies the patch and returns the patched vpcRoutingTableRoute.
func (c *FakeVpcRoutingTableRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcRoutingTableRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcroutingtableroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcRoutingTableRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcRoutingTableRoute), err
}
