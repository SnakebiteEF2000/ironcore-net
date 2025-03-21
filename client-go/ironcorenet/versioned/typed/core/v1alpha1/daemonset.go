// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	corev1alpha1 "github.com/ironcore-dev/ironcore-net/client-go/applyconfigurations/core/v1alpha1"
	scheme "github.com/ironcore-dev/ironcore-net/client-go/ironcorenet/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// DaemonSetsGetter has a method to return a DaemonSetInterface.
// A group's client should implement this interface.
type DaemonSetsGetter interface {
	DaemonSets(namespace string) DaemonSetInterface
}

// DaemonSetInterface has methods to work with DaemonSet resources.
type DaemonSetInterface interface {
	Create(ctx context.Context, daemonSet *v1alpha1.DaemonSet, opts v1.CreateOptions) (*v1alpha1.DaemonSet, error)
	Update(ctx context.Context, daemonSet *v1alpha1.DaemonSet, opts v1.UpdateOptions) (*v1alpha1.DaemonSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, daemonSet *v1alpha1.DaemonSet, opts v1.UpdateOptions) (*v1alpha1.DaemonSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DaemonSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DaemonSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DaemonSet, err error)
	Apply(ctx context.Context, daemonSet *corev1alpha1.DaemonSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DaemonSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, daemonSet *corev1alpha1.DaemonSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DaemonSet, err error)
	DaemonSetExpansion
}

// daemonSets implements DaemonSetInterface
type daemonSets struct {
	*gentype.ClientWithListAndApply[*v1alpha1.DaemonSet, *v1alpha1.DaemonSetList, *corev1alpha1.DaemonSetApplyConfiguration]
}

// newDaemonSets returns a DaemonSets
func newDaemonSets(c *CoreV1alpha1Client, namespace string) *daemonSets {
	return &daemonSets{
		gentype.NewClientWithListAndApply[*v1alpha1.DaemonSet, *v1alpha1.DaemonSetList, *corev1alpha1.DaemonSetApplyConfiguration](
			"daemonsets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.DaemonSet { return &v1alpha1.DaemonSet{} },
			func() *v1alpha1.DaemonSetList { return &v1alpha1.DaemonSetList{} }),
	}
}
