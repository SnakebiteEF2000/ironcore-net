// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	corev1alpha1 "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodes implements NodeInterface
type FakeNodes struct {
	Fake *FakeCoreV1alpha1
}

var nodesResource = v1alpha1.SchemeGroupVersion.WithResource("nodes")

var nodesKind = v1alpha1.SchemeGroupVersion.WithKind("Node")

// Get takes name of the node, and returns the corresponding node object, and an error if there is any.
func (c *FakeNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodesResource, name), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}

// List takes label and field selectors, and returns the list of Nodes that match those selectors.
func (c *FakeNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodesResource, nodesKind, opts), &v1alpha1.NodeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeList{ListMeta: obj.(*v1alpha1.NodeList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodes.
func (c *FakeNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodesResource, opts))
}

// Create takes the representation of a node and creates it.  Returns the server's representation of the node, and an error, if there is any.
func (c *FakeNodes) Create(ctx context.Context, node *v1alpha1.Node, opts v1.CreateOptions) (result *v1alpha1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodesResource, node), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}

// Update takes the representation of a node and updates it. Returns the server's representation of the node, and an error, if there is any.
func (c *FakeNodes) Update(ctx context.Context, node *v1alpha1.Node, opts v1.UpdateOptions) (result *v1alpha1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodesResource, node), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodes) UpdateStatus(ctx context.Context, node *v1alpha1.Node, opts v1.UpdateOptions) (*v1alpha1.Node, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodesResource, "status", node), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}

// Delete takes name of the node and deletes it. Returns an error if one occurs.
func (c *FakeNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nodesResource, name, opts), &v1alpha1.Node{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeList{})
	return err
}

// Patch applies the patch and returns the patched node.
func (c *FakeNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Node, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodesResource, name, pt, data, subresources...), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied node.
func (c *FakeNodes) Apply(ctx context.Context, node *corev1alpha1.NodeApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodesResource, *name, types.ApplyPatchType, data), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeNodes) ApplyStatus(ctx context.Context, node *corev1alpha1.NodeApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node provided to Apply must not be nil")
	}
	data, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	name := node.Name
	if name == nil {
		return nil, fmt.Errorf("node.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodesResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Node{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Node), err
}
