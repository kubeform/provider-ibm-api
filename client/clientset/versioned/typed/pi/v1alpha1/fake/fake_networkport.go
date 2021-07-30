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

// FakeNetworkPorts implements NetworkPortInterface
type FakeNetworkPorts struct {
	Fake *FakePiV1alpha1
	ns   string
}

var networkportsResource = schema.GroupVersionResource{Group: "pi.ibm.kubeform.com", Version: "v1alpha1", Resource: "networkports"}

var networkportsKind = schema.GroupVersionKind{Group: "pi.ibm.kubeform.com", Version: "v1alpha1", Kind: "NetworkPort"}

// Get takes name of the networkPort, and returns the corresponding networkPort object, and an error if there is any.
func (c *FakeNetworkPorts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkPort, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkportsResource, c.ns, name), &v1alpha1.NetworkPort{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPort), err
}

// List takes label and field selectors, and returns the list of NetworkPorts that match those selectors.
func (c *FakeNetworkPorts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkPortList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkportsResource, networkportsKind, c.ns, opts), &v1alpha1.NetworkPortList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkPortList{ListMeta: obj.(*v1alpha1.NetworkPortList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkPortList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPorts.
func (c *FakeNetworkPorts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkportsResource, c.ns, opts))

}

// Create takes the representation of a networkPort and creates it.  Returns the server's representation of the networkPort, and an error, if there is any.
func (c *FakeNetworkPorts) Create(ctx context.Context, networkPort *v1alpha1.NetworkPort, opts v1.CreateOptions) (result *v1alpha1.NetworkPort, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkportsResource, c.ns, networkPort), &v1alpha1.NetworkPort{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPort), err
}

// Update takes the representation of a networkPort and updates it. Returns the server's representation of the networkPort, and an error, if there is any.
func (c *FakeNetworkPorts) Update(ctx context.Context, networkPort *v1alpha1.NetworkPort, opts v1.UpdateOptions) (result *v1alpha1.NetworkPort, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkportsResource, c.ns, networkPort), &v1alpha1.NetworkPort{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPort), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkPorts) UpdateStatus(ctx context.Context, networkPort *v1alpha1.NetworkPort, opts v1.UpdateOptions) (*v1alpha1.NetworkPort, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkportsResource, "status", c.ns, networkPort), &v1alpha1.NetworkPort{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPort), err
}

// Delete takes name of the networkPort and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPorts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkportsResource, c.ns, name), &v1alpha1.NetworkPort{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPorts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkPortList{})
	return err
}

// Patch applies the patch and returns the patched networkPort.
func (c *FakeNetworkPorts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkPort, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkPort{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkPort), err
}
