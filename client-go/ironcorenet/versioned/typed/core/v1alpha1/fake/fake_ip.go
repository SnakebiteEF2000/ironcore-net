// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
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

// FakeIPs implements IPInterface
type FakeIPs struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var ipsResource = v1alpha1.SchemeGroupVersion.WithResource("ips")

var ipsKind = v1alpha1.SchemeGroupVersion.WithKind("IP")

// Get takes name of the iP, and returns the corresponding iP object, and an error if there is any.
func (c *FakeIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(ipsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// List takes label and field selectors, and returns the list of IPs that match those selectors.
func (c *FakeIPs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPList, err error) {
	emptyResult := &v1alpha1.IPList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(ipsResource, ipsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IPList{ListMeta: obj.(*v1alpha1.IPList).ListMeta}
	for _, item := range obj.(*v1alpha1.IPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPs.
func (c *FakeIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(ipsResource, c.ns, opts))

}

// Create takes the representation of a iP and creates it.  Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Create(ctx context.Context, iP *v1alpha1.IP, opts v1.CreateOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(ipsResource, c.ns, iP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// Update takes the representation of a iP and updates it. Returns the server's representation of the iP, and an error, if there is any.
func (c *FakeIPs) Update(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(ipsResource, c.ns, iP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPs) UpdateStatus(ctx context.Context, iP *v1alpha1.IP, opts v1.UpdateOptions) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(ipsResource, "status", c.ns, iP, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// Delete takes name of the iP and deletes it. Returns an error if one occurs.
func (c *FakeIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipsResource, c.ns, name, opts), &v1alpha1.IP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(ipsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IPList{})
	return err
}

// Patch applies the patch and returns the patched iP.
func (c *FakeIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IP, err error) {
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(ipsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied iP.
func (c *FakeIPs) Apply(ctx context.Context, iP *corev1alpha1.IPApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.IP, err error) {
	if iP == nil {
		return nil, fmt.Errorf("iP provided to Apply must not be nil")
	}
	data, err := json.Marshal(iP)
	if err != nil {
		return nil, err
	}
	name := iP.Name
	if name == nil {
		return nil, fmt.Errorf("iP.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(ipsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeIPs) ApplyStatus(ctx context.Context, iP *corev1alpha1.IPApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.IP, err error) {
	if iP == nil {
		return nil, fmt.Errorf("iP provided to Apply must not be nil")
	}
	data, err := json.Marshal(iP)
	if err != nil {
		return nil, err
	}
	name := iP.Name
	if name == nil {
		return nil, fmt.Errorf("iP.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.IP{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(ipsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.IP), err
}
