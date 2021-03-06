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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/enterprise/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEnterprises implements EnterpriseInterface
type FakeEnterprises struct {
	Fake *FakeEnterpriseV1alpha1
	ns   string
}

var enterprisesResource = schema.GroupVersionResource{Group: "enterprise.ibm.kubeform.com", Version: "v1alpha1", Resource: "enterprises"}

var enterprisesKind = schema.GroupVersionKind{Group: "enterprise.ibm.kubeform.com", Version: "v1alpha1", Kind: "Enterprise"}

// Get takes name of the enterprise, and returns the corresponding enterprise object, and an error if there is any.
func (c *FakeEnterprises) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Enterprise, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(enterprisesResource, c.ns, name), &v1alpha1.Enterprise{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Enterprise), err
}

// List takes label and field selectors, and returns the list of Enterprises that match those selectors.
func (c *FakeEnterprises) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EnterpriseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(enterprisesResource, enterprisesKind, c.ns, opts), &v1alpha1.EnterpriseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EnterpriseList{ListMeta: obj.(*v1alpha1.EnterpriseList).ListMeta}
	for _, item := range obj.(*v1alpha1.EnterpriseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested enterprises.
func (c *FakeEnterprises) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(enterprisesResource, c.ns, opts))

}

// Create takes the representation of a enterprise and creates it.  Returns the server's representation of the enterprise, and an error, if there is any.
func (c *FakeEnterprises) Create(ctx context.Context, enterprise *v1alpha1.Enterprise, opts v1.CreateOptions) (result *v1alpha1.Enterprise, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(enterprisesResource, c.ns, enterprise), &v1alpha1.Enterprise{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Enterprise), err
}

// Update takes the representation of a enterprise and updates it. Returns the server's representation of the enterprise, and an error, if there is any.
func (c *FakeEnterprises) Update(ctx context.Context, enterprise *v1alpha1.Enterprise, opts v1.UpdateOptions) (result *v1alpha1.Enterprise, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(enterprisesResource, c.ns, enterprise), &v1alpha1.Enterprise{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Enterprise), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEnterprises) UpdateStatus(ctx context.Context, enterprise *v1alpha1.Enterprise, opts v1.UpdateOptions) (*v1alpha1.Enterprise, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(enterprisesResource, "status", c.ns, enterprise), &v1alpha1.Enterprise{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Enterprise), err
}

// Delete takes name of the enterprise and deletes it. Returns an error if one occurs.
func (c *FakeEnterprises) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(enterprisesResource, c.ns, name), &v1alpha1.Enterprise{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEnterprises) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(enterprisesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EnterpriseList{})
	return err
}

// Patch applies the patch and returns the patched enterprise.
func (c *FakeEnterprises) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Enterprise, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(enterprisesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Enterprise{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Enterprise), err
}
