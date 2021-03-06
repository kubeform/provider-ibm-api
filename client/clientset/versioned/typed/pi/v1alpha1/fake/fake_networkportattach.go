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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/pi/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkPortAttaches implements NetworkPortAttachInterface
type FakeNetworkPortAttaches struct {
	Fake *FakePiV1alpha1
	ns   string
}

var networkportattachesResource = schema.GroupVersionResource{Group: "pi.ibm.kubeform.com", Version: "v1alpha1", Resource: "networkportattaches"}

var networkportattachesKind = schema.GroupVersionKind{Group: "pi.ibm.kubeform.com", Version: "v1alpha1", Kind: "NetworkPortAttach"}

// Get takes name of the networkPortAttach, and returns the corresponding networkPortAttach object, and an error if there is any.
func (c *FakeNetworkPortAttaches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkPortAttach, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkportattachesResource, c.ns, name), &v1alpha1.NetworkPortAttach{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPortAttach), err
}

// List takes label and field selectors, and returns the list of NetworkPortAttaches that match those selectors.
func (c *FakeNetworkPortAttaches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkPortAttachList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkportattachesResource, networkportattachesKind, c.ns, opts), &v1alpha1.NetworkPortAttachList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkPortAttachList{ListMeta: obj.(*v1alpha1.NetworkPortAttachList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkPortAttachList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPortAttaches.
func (c *FakeNetworkPortAttaches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkportattachesResource, c.ns, opts))

}

// Create takes the representation of a networkPortAttach and creates it.  Returns the server's representation of the networkPortAttach, and an error, if there is any.
func (c *FakeNetworkPortAttaches) Create(ctx context.Context, networkPortAttach *v1alpha1.NetworkPortAttach, opts v1.CreateOptions) (result *v1alpha1.NetworkPortAttach, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkportattachesResource, c.ns, networkPortAttach), &v1alpha1.NetworkPortAttach{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPortAttach), err
}

// Update takes the representation of a networkPortAttach and updates it. Returns the server's representation of the networkPortAttach, and an error, if there is any.
func (c *FakeNetworkPortAttaches) Update(ctx context.Context, networkPortAttach *v1alpha1.NetworkPortAttach, opts v1.UpdateOptions) (result *v1alpha1.NetworkPortAttach, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkportattachesResource, c.ns, networkPortAttach), &v1alpha1.NetworkPortAttach{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPortAttach), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPortAttaches) UpdateStatus(ctx context.Context, networkPortAttach *v1alpha1.NetworkPortAttach, opts v1.UpdateOptions) (*v1alpha1.NetworkPortAttach, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkportattachesResource, "status", c.ns, networkPortAttach), &v1alpha1.NetworkPortAttach{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPortAttach), err
}

// Delete takes name of the networkPortAttach and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPortAttaches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkportattachesResource, c.ns, name), &v1alpha1.NetworkPortAttach{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPortAttaches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkportattachesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkPortAttachList{})
	return err
}

// Patch applies the patch and returns the patched networkPortAttach.
func (c *FakeNetworkPortAttaches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkPortAttach, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkportattachesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkPortAttach{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPortAttach), err
}
