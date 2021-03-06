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

	v1alpha1 "kubeform.dev/provider-ibm-api/apis/container/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterFeatures implements ClusterFeatureInterface
type FakeClusterFeatures struct {
	Fake *FakeContainerV1alpha1
	ns   string
}

var clusterfeaturesResource = schema.GroupVersionResource{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Resource: "clusterfeatures"}

var clusterfeaturesKind = schema.GroupVersionKind{Group: "container.ibm.kubeform.com", Version: "v1alpha1", Kind: "ClusterFeature"}

// Get takes name of the clusterFeature, and returns the corresponding clusterFeature object, and an error if there is any.
func (c *FakeClusterFeatures) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterfeaturesResource, c.ns, name), &v1alpha1.ClusterFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterFeature), err
}

// List takes label and field selectors, and returns the list of ClusterFeatures that match those selectors.
func (c *FakeClusterFeatures) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterFeatureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterfeaturesResource, clusterfeaturesKind, c.ns, opts), &v1alpha1.ClusterFeatureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterFeatureList{ListMeta: obj.(*v1alpha1.ClusterFeatureList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterFeatureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterFeatures.
func (c *FakeClusterFeatures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterfeaturesResource, c.ns, opts))

}

// Create takes the representation of a clusterFeature and creates it.  Returns the server's representation of the clusterFeature, and an error, if there is any.
func (c *FakeClusterFeatures) Create(ctx context.Context, clusterFeature *v1alpha1.ClusterFeature, opts v1.CreateOptions) (result *v1alpha1.ClusterFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterfeaturesResource, c.ns, clusterFeature), &v1alpha1.ClusterFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterFeature), err
}

// Update takes the representation of a clusterFeature and updates it. Returns the server's representation of the clusterFeature, and an error, if there is any.
func (c *FakeClusterFeatures) Update(ctx context.Context, clusterFeature *v1alpha1.ClusterFeature, opts v1.UpdateOptions) (result *v1alpha1.ClusterFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterfeaturesResource, c.ns, clusterFeature), &v1alpha1.ClusterFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterFeature), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterFeatures) UpdateStatus(ctx context.Context, clusterFeature *v1alpha1.ClusterFeature, opts v1.UpdateOptions) (*v1alpha1.ClusterFeature, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterfeaturesResource, "status", c.ns, clusterFeature), &v1alpha1.ClusterFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterFeature), err
}

// Delete takes name of the clusterFeature and deletes it. Returns an error if one occurs.
func (c *FakeClusterFeatures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterfeaturesResource, c.ns, name), &v1alpha1.ClusterFeature{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterFeatures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterfeaturesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterFeatureList{})
	return err
}

// Patch applies the patch and returns the patched clusterFeature.
func (c *FakeClusterFeatures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterFeature, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterfeaturesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterFeature{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterFeature), err
}
