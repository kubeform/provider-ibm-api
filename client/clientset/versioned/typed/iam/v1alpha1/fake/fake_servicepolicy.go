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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/iam/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServicePolicies implements ServicePolicyInterface
type FakeServicePolicies struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var servicepoliciesResource = schema.GroupVersionResource{Group: "iam.ibm.kubeform.com", Version: "v1alpha1", Resource: "servicepolicies"}

var servicepoliciesKind = schema.GroupVersionKind{Group: "iam.ibm.kubeform.com", Version: "v1alpha1", Kind: "ServicePolicy"}

// Get takes name of the servicePolicy, and returns the corresponding servicePolicy object, and an error if there is any.
func (c *FakeServicePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicepoliciesResource, c.ns, name), &v1alpha1.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicePolicy), err
}

// List takes label and field selectors, and returns the list of ServicePolicies that match those selectors.
func (c *FakeServicePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServicePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicepoliciesResource, servicepoliciesKind, c.ns, opts), &v1alpha1.ServicePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServicePolicyList{ListMeta: obj.(*v1alpha1.ServicePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServicePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servicePolicies.
func (c *FakeServicePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicepoliciesResource, c.ns, opts))

}

// Create takes the representation of a servicePolicy and creates it.  Returns the server's representation of the servicePolicy, and an error, if there is any.
func (c *FakeServicePolicies) Create(ctx context.Context, servicePolicy *v1alpha1.ServicePolicy, opts v1.CreateOptions) (result *v1alpha1.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicepoliciesResource, c.ns, servicePolicy), &v1alpha1.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicePolicy), err
}

// Update takes the representation of a servicePolicy and updates it. Returns the server's representation of the servicePolicy, and an error, if there is any.
func (c *FakeServicePolicies) Update(ctx context.Context, servicePolicy *v1alpha1.ServicePolicy, opts v1.UpdateOptions) (result *v1alpha1.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicepoliciesResource, c.ns, servicePolicy), &v1alpha1.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServicePolicies) UpdateStatus(ctx context.Context, servicePolicy *v1alpha1.ServicePolicy, opts v1.UpdateOptions) (*v1alpha1.ServicePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicepoliciesResource, "status", c.ns, servicePolicy), &v1alpha1.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicePolicy), err
}

// Delete takes name of the servicePolicy and deletes it. Returns an error if one occurs.
func (c *FakeServicePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicepoliciesResource, c.ns, name), &v1alpha1.ServicePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServicePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServicePolicyList{})
	return err
}

// Patch applies the patch and returns the patched servicePolicy.
func (c *FakeServicePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServicePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServicePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServicePolicy), err
}
