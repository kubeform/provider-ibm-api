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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/ob/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLoggings implements LoggingInterface
type FakeLoggings struct {
	Fake *FakeObV1alpha1
	ns   string
}

var loggingsResource = schema.GroupVersionResource{Group: "ob.ibm.kubeform.com", Version: "v1alpha1", Resource: "loggings"}

var loggingsKind = schema.GroupVersionKind{Group: "ob.ibm.kubeform.com", Version: "v1alpha1", Kind: "Logging"}

// Get takes name of the logging, and returns the corresponding logging object, and an error if there is any.
func (c *FakeLoggings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loggingsResource, c.ns, name), &v1alpha1.Logging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logging), err
}

// List takes label and field selectors, and returns the list of Loggings that match those selectors.
func (c *FakeLoggings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LoggingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loggingsResource, loggingsKind, c.ns, opts), &v1alpha1.LoggingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoggingList{ListMeta: obj.(*v1alpha1.LoggingList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoggingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggings.
func (c *FakeLoggings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loggingsResource, c.ns, opts))

}

// Create takes the representation of a logging and creates it.  Returns the server's representation of the logging, and an error, if there is any.
func (c *FakeLoggings) Create(ctx context.Context, logging *v1alpha1.Logging, opts v1.CreateOptions) (result *v1alpha1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loggingsResource, c.ns, logging), &v1alpha1.Logging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logging), err
}

// Update takes the representation of a logging and updates it. Returns the server's representation of the logging, and an error, if there is any.
func (c *FakeLoggings) Update(ctx context.Context, logging *v1alpha1.Logging, opts v1.UpdateOptions) (result *v1alpha1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loggingsResource, c.ns, logging), &v1alpha1.Logging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logging), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggings) UpdateStatus(ctx context.Context, logging *v1alpha1.Logging, opts v1.UpdateOptions) (*v1alpha1.Logging, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loggingsResource, "status", c.ns, logging), &v1alpha1.Logging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logging), err
}

// Delete takes name of the logging and deletes it. Returns an error if one occurs.
func (c *FakeLoggings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loggingsResource, c.ns, name), &v1alpha1.Logging{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loggingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoggingList{})
	return err
}

// Patch applies the patch and returns the patched logging.
func (c *FakeLoggings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loggingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Logging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logging), err
}
