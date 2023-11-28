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

// FakeLoadBalancers implements LoadBalancerInterface
type FakeLoadBalancers struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var loadbalancersResource = v1alpha1.SchemeGroupVersion.WithResource("loadbalancers")

var loadbalancersKind = v1alpha1.SchemeGroupVersion.WithKind("LoadBalancer")

// Get takes name of the loadBalancer, and returns the corresponding loadBalancer object, and an error if there is any.
func (c *FakeLoadBalancers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loadbalancersResource, c.ns, name), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// List takes label and field selectors, and returns the list of LoadBalancers that match those selectors.
func (c *FakeLoadBalancers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LoadBalancerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loadbalancersResource, loadbalancersKind, c.ns, opts), &v1alpha1.LoadBalancerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoadBalancerList{ListMeta: obj.(*v1alpha1.LoadBalancerList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoadBalancerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loadBalancers.
func (c *FakeLoadBalancers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loadbalancersResource, c.ns, opts))

}

// Create takes the representation of a loadBalancer and creates it.  Returns the server's representation of the loadBalancer, and an error, if there is any.
func (c *FakeLoadBalancers) Create(ctx context.Context, loadBalancer *v1alpha1.LoadBalancer, opts v1.CreateOptions) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loadbalancersResource, c.ns, loadBalancer), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// Update takes the representation of a loadBalancer and updates it. Returns the server's representation of the loadBalancer, and an error, if there is any.
func (c *FakeLoadBalancers) Update(ctx context.Context, loadBalancer *v1alpha1.LoadBalancer, opts v1.UpdateOptions) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loadbalancersResource, c.ns, loadBalancer), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoadBalancers) UpdateStatus(ctx context.Context, loadBalancer *v1alpha1.LoadBalancer, opts v1.UpdateOptions) (*v1alpha1.LoadBalancer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loadbalancersResource, "status", c.ns, loadBalancer), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// Delete takes name of the loadBalancer and deletes it. Returns an error if one occurs.
func (c *FakeLoadBalancers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(loadbalancersResource, c.ns, name, opts), &v1alpha1.LoadBalancer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoadBalancers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loadbalancersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoadBalancerList{})
	return err
}

// Patch applies the patch and returns the patched loadBalancer.
func (c *FakeLoadBalancers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loadbalancersResource, c.ns, name, pt, data, subresources...), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied loadBalancer.
func (c *FakeLoadBalancers) Apply(ctx context.Context, loadBalancer *corev1alpha1.LoadBalancerApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LoadBalancer, err error) {
	if loadBalancer == nil {
		return nil, fmt.Errorf("loadBalancer provided to Apply must not be nil")
	}
	data, err := json.Marshal(loadBalancer)
	if err != nil {
		return nil, err
	}
	name := loadBalancer.Name
	if name == nil {
		return nil, fmt.Errorf("loadBalancer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loadbalancersResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeLoadBalancers) ApplyStatus(ctx context.Context, loadBalancer *corev1alpha1.LoadBalancerApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LoadBalancer, err error) {
	if loadBalancer == nil {
		return nil, fmt.Errorf("loadBalancer provided to Apply must not be nil")
	}
	data, err := json.Marshal(loadBalancer)
	if err != nil {
		return nil, err
	}
	name := loadBalancer.Name
	if name == nil {
		return nil, fmt.Errorf("loadBalancer.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loadbalancersResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}
