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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/dl/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProviderGateways implements ProviderGatewayInterface
type FakeProviderGateways struct {
	Fake *FakeDlV1alpha1
	ns   string
}

var providergatewaysResource = schema.GroupVersionResource{Group: "dl.ibm.kubeform.com", Version: "v1alpha1", Resource: "providergateways"}

var providergatewaysKind = schema.GroupVersionKind{Group: "dl.ibm.kubeform.com", Version: "v1alpha1", Kind: "ProviderGateway"}

// Get takes name of the providerGateway, and returns the corresponding providerGateway object, and an error if there is any.
func (c *FakeProviderGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProviderGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(providergatewaysResource, c.ns, name), &v1alpha1.ProviderGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProviderGateway), err
}

// List takes label and field selectors, and returns the list of ProviderGateways that match those selectors.
func (c *FakeProviderGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProviderGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(providergatewaysResource, providergatewaysKind, c.ns, opts), &v1alpha1.ProviderGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProviderGatewayList{ListMeta: obj.(*v1alpha1.ProviderGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProviderGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested providerGateways.
func (c *FakeProviderGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(providergatewaysResource, c.ns, opts))

}

// Create takes the representation of a providerGateway and creates it.  Returns the server's representation of the providerGateway, and an error, if there is any.
func (c *FakeProviderGateways) Create(ctx context.Context, providerGateway *v1alpha1.ProviderGateway, opts v1.CreateOptions) (result *v1alpha1.ProviderGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(providergatewaysResource, c.ns, providerGateway), &v1alpha1.ProviderGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProviderGateway), err
}

// Update takes the representation of a providerGateway and updates it. Returns the server's representation of the providerGateway, and an error, if there is any.
func (c *FakeProviderGateways) Update(ctx context.Context, providerGateway *v1alpha1.ProviderGateway, opts v1.UpdateOptions) (result *v1alpha1.ProviderGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(providergatewaysResource, c.ns, providerGateway), &v1alpha1.ProviderGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProviderGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProviderGateways) UpdateStatus(ctx context.Context, providerGateway *v1alpha1.ProviderGateway, opts v1.UpdateOptions) (*v1alpha1.ProviderGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(providergatewaysResource, "status", c.ns, providerGateway), &v1alpha1.ProviderGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProviderGateway), err
}

// Delete takes name of the providerGateway and deletes it. Returns an error if one occurs.
func (c *FakeProviderGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(providergatewaysResource, c.ns, name), &v1alpha1.ProviderGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProviderGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(providergatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProviderGatewayList{})
	return err
}

// Patch applies the patch and returns the patched providerGateway.
func (c *FakeProviderGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProviderGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(providergatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProviderGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProviderGateway), err
}
