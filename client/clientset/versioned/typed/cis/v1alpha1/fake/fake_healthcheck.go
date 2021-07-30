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

// FakeHealthchecks implements HealthcheckInterface
type FakeHealthchecks struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var healthchecksResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "healthchecks"}

var healthchecksKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "Healthcheck"}

// Get takes name of the healthcheck, and returns the corresponding healthcheck object, and an error if there is any.
func (c *FakeHealthchecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Healthcheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(healthchecksResource, c.ns, name), &v1alpha1.Healthcheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Healthcheck), err
}

// List takes label and field selectors, and returns the list of Healthchecks that match those selectors.
func (c *FakeHealthchecks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HealthcheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(healthchecksResource, healthchecksKind, c.ns, opts), &v1alpha1.HealthcheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HealthcheckList{ListMeta: obj.(*v1alpha1.HealthcheckList).ListMeta}
	for _, item := range obj.(*v1alpha1.HealthcheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested healthchecks.
func (c *FakeHealthchecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(healthchecksResource, c.ns, opts))

}

// Create takes the representation of a healthcheck and creates it.  Returns the server's representation of the healthcheck, and an error, if there is any.
func (c *FakeHealthchecks) Create(ctx context.Context, healthcheck *v1alpha1.Healthcheck, opts v1.CreateOptions) (result *v1alpha1.Healthcheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(healthchecksResource, c.ns, healthcheck), &v1alpha1.Healthcheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Healthcheck), err
}

// Update takes the representation of a healthcheck and updates it. Returns the server's representation of the healthcheck, and an error, if there is any.
func (c *FakeHealthchecks) Update(ctx context.Context, healthcheck *v1alpha1.Healthcheck, opts v1.UpdateOptions) (result *v1alpha1.Healthcheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(healthchecksResource, c.ns, healthcheck), &v1alpha1.Healthcheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Healthcheck), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHealthchecks) UpdateStatus(ctx context.Context, healthcheck *v1alpha1.Healthcheck, opts v1.UpdateOptions) (*v1alpha1.Healthcheck, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(healthchecksResource, "status", c.ns, healthcheck), &v1alpha1.Healthcheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Healthcheck), err
}

// Delete takes name of the healthcheck and deletes it. Returns an error if one occurs.
func (c *FakeHealthchecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(healthchecksResource, c.ns, name), &v1alpha1.Healthcheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHealthchecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(healthchecksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HealthcheckList{})
	return err
}

// Patch applies the patch and returns the patched healthcheck.
func (c *FakeHealthchecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Healthcheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(healthchecksResource, c.ns, name, pt, data, subresources...), &v1alpha1.Healthcheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Healthcheck), err
}
