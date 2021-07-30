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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/schematics/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeActions implements ActionInterface
type FakeActions struct {
	Fake *FakeSchematicsV1alpha1
	ns   string
}

var actionsResource = schema.GroupVersionResource{Group: "schematics.ibm.kubeform.com", Version: "v1alpha1", Resource: "actions"}

var actionsKind = schema.GroupVersionKind{Group: "schematics.ibm.kubeform.com", Version: "v1alpha1", Kind: "Action"}

// Get takes name of the action, and returns the corresponding action object, and an error if there is any.
func (c *FakeActions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Action, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(actionsResource, c.ns, name), &v1alpha1.Action{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Action), err
}

// List takes label and field selectors, and returns the list of Actions that match those selectors.
func (c *FakeActions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ActionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(actionsResource, actionsKind, c.ns, opts), &v1alpha1.ActionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ActionList{ListMeta: obj.(*v1alpha1.ActionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ActionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested actions.
func (c *FakeActions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(actionsResource, c.ns, opts))

}

// Create takes the representation of a action and creates it.  Returns the server's representation of the action, and an error, if there is any.
func (c *FakeActions) Create(ctx context.Context, action *v1alpha1.Action, opts v1.CreateOptions) (result *v1alpha1.Action, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(actionsResource, c.ns, action), &v1alpha1.Action{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Action), err
}

// Update takes the representation of a action and updates it. Returns the server's representation of the action, and an error, if there is any.
func (c *FakeActions) Update(ctx context.Context, action *v1alpha1.Action, opts v1.UpdateOptions) (result *v1alpha1.Action, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(actionsResource, c.ns, action), &v1alpha1.Action{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Action), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeActions) UpdateStatus(ctx context.Context, action *v1alpha1.Action, opts v1.UpdateOptions) (*v1alpha1.Action, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(actionsResource, "status", c.ns, action), &v1alpha1.Action{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Action), err
}

// Delete takes name of the action and deletes it. Returns an error if one occurs.
func (c *FakeActions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(actionsResource, c.ns, name), &v1alpha1.Action{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeActions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(actionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ActionList{})
	return err
}

// Patch applies the patch and returns the patched action.
func (c *FakeActions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Action, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(actionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Action{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Action), err
}
