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

// FakeWafRules implements WafRuleInterface
type FakeWafRules struct {
	Fake *FakeCisV1alpha1
	ns   string
}

var wafrulesResource = schema.GroupVersionResource{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Resource: "wafrules"}

var wafrulesKind = schema.GroupVersionKind{Group: "cis.ibm.kubeform.com", Version: "v1alpha1", Kind: "WafRule"}

// Get takes name of the wafRule, and returns the corresponding wafRule object, and an error if there is any.
func (c *FakeWafRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WafRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(wafrulesResource, c.ns, name), &v1alpha1.WafRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRule), err
}

// List takes label and field selectors, and returns the list of WafRules that match those selectors.
func (c *FakeWafRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WafRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(wafrulesResource, wafrulesKind, c.ns, opts), &v1alpha1.WafRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafRuleList{ListMeta: obj.(*v1alpha1.WafRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafRules.
func (c *FakeWafRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(wafrulesResource, c.ns, opts))

}

// Create takes the representation of a wafRule and creates it.  Returns the server's representation of the wafRule, and an error, if there is any.
func (c *FakeWafRules) Create(ctx context.Context, wafRule *v1alpha1.WafRule, opts v1.CreateOptions) (result *v1alpha1.WafRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(wafrulesResource, c.ns, wafRule), &v1alpha1.WafRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRule), err
}

// Update takes the representation of a wafRule and updates it. Returns the server's representation of the wafRule, and an error, if there is any.
func (c *FakeWafRules) Update(ctx context.Context, wafRule *v1alpha1.WafRule, opts v1.UpdateOptions) (result *v1alpha1.WafRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(wafrulesResource, c.ns, wafRule), &v1alpha1.WafRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafRules) UpdateStatus(ctx context.Context, wafRule *v1alpha1.WafRule, opts v1.UpdateOptions) (*v1alpha1.WafRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(wafrulesResource, "status", c.ns, wafRule), &v1alpha1.WafRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRule), err
}

// Delete takes name of the wafRule and deletes it. Returns an error if one occurs.
func (c *FakeWafRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(wafrulesResource, c.ns, name), &v1alpha1.WafRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(wafrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafRuleList{})
	return err
}

// Patch applies the patch and returns the patched wafRule.
func (c *FakeWafRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WafRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(wafrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.WafRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafRule), err
}
