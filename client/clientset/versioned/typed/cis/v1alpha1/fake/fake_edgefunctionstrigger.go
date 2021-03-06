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

// FakeEdgeFunctionsTriggers implements EdgeFunctionsTriggerInterface
type FakeEdgeFunctionsTriggers struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var edgefunctionstriggersResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "edgefunctionstriggers"}

var edgefunctionstriggersKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "EdgeFunctionsTrigger"}

// Get takes name of the edgeFunctionsTrigger, and returns the corresponding edgeFunctionsTrigger object, and an error if there is any.
func (c *FakeEdgeFunctionsTriggers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EdgeFunctionsTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(edgefunctionstriggersResource, c.ns, name), &v1alpha1.EdgeFunctionsTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeFunctionsTrigger), err
}

// List takes label and field selectors, and returns the list of EdgeFunctionsTriggers that match those selectors.
func (c *FakeEdgeFunctionsTriggers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EdgeFunctionsTriggerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(edgefunctionstriggersResource, edgefunctionstriggersKind, c.ns, opts), &v1alpha1.EdgeFunctionsTriggerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EdgeFunctionsTriggerList{ListMeta: obj.(*v1alpha1.EdgeFunctionsTriggerList).ListMeta}
	for _, item := range obj.(*v1alpha1.EdgeFunctionsTriggerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested edgeFunctionsTriggers.
func (c *FakeEdgeFunctionsTriggers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(edgefunctionstriggersResource, c.ns, opts))

}

// Create takes the representation of a edgeFunctionsTrigger and creates it.  Returns the server's representation of the edgeFunctionsTrigger, and an error, if there is any.
func (c *FakeEdgeFunctionsTriggers) Create(ctx context.Context, edgeFunctionsTrigger *v1alpha1.EdgeFunctionsTrigger, opts v1.CreateOptions) (result *v1alpha1.EdgeFunctionsTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(edgefunctionstriggersResource, c.ns, edgeFunctionsTrigger), &v1alpha1.EdgeFunctionsTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeFunctionsTrigger), err
}

// Update takes the representation of a edgeFunctionsTrigger and updates it. Returns the server's representation of the edgeFunctionsTrigger, and an error, if there is any.
func (c *FakeEdgeFunctionsTriggers) Update(ctx context.Context, edgeFunctionsTrigger *v1alpha1.EdgeFunctionsTrigger, opts v1.UpdateOptions) (result *v1alpha1.EdgeFunctionsTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(edgefunctionstriggersResource, c.ns, edgeFunctionsTrigger), &v1alpha1.EdgeFunctionsTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeFunctionsTrigger), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEdgeFunctionsTriggers) UpdateStatus(ctx context.Context, edgeFunctionsTrigger *v1alpha1.EdgeFunctionsTrigger, opts v1.UpdateOptions) (*v1alpha1.EdgeFunctionsTrigger, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(edgefunctionstriggersResource, "status", c.ns, edgeFunctionsTrigger), &v1alpha1.EdgeFunctionsTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeFunctionsTrigger), err
}

// Delete takes name of the edgeFunctionsTrigger and deletes it. Returns an error if one occurs.
func (c *FakeEdgeFunctionsTriggers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(edgefunctionstriggersResource, c.ns, name), &v1alpha1.EdgeFunctionsTrigger{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEdgeFunctionsTriggers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(edgefunctionstriggersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EdgeFunctionsTriggerList{})
	return err
}

// Patch applies the patch and returns the patched edgeFunctionsTrigger.
func (c *FakeEdgeFunctionsTriggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EdgeFunctionsTrigger, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(edgefunctionstriggersResource, c.ns, name, pt, data, subresources...), &v1alpha1.EdgeFunctionsTrigger{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeFunctionsTrigger), err
}
