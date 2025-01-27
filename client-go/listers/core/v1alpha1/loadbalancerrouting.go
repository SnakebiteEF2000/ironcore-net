// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ironcore-net/api/core/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// LoadBalancerRoutingLister helps list LoadBalancerRoutings.
// All objects returned here must be treated as read-only.
type LoadBalancerRoutingLister interface {
	// List lists all LoadBalancerRoutings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LoadBalancerRouting, err error)
	// LoadBalancerRoutings returns an object that can list and get LoadBalancerRoutings.
	LoadBalancerRoutings(namespace string) LoadBalancerRoutingNamespaceLister
	LoadBalancerRoutingListerExpansion
}

// loadBalancerRoutingLister implements the LoadBalancerRoutingLister interface.
type loadBalancerRoutingLister struct {
	listers.ResourceIndexer[*v1alpha1.LoadBalancerRouting]
}

// NewLoadBalancerRoutingLister returns a new LoadBalancerRoutingLister.
func NewLoadBalancerRoutingLister(indexer cache.Indexer) LoadBalancerRoutingLister {
	return &loadBalancerRoutingLister{listers.New[*v1alpha1.LoadBalancerRouting](indexer, v1alpha1.Resource("loadbalancerrouting"))}
}

// LoadBalancerRoutings returns an object that can list and get LoadBalancerRoutings.
func (s *loadBalancerRoutingLister) LoadBalancerRoutings(namespace string) LoadBalancerRoutingNamespaceLister {
	return loadBalancerRoutingNamespaceLister{listers.NewNamespaced[*v1alpha1.LoadBalancerRouting](s.ResourceIndexer, namespace)}
}

// LoadBalancerRoutingNamespaceLister helps list and get LoadBalancerRoutings.
// All objects returned here must be treated as read-only.
type LoadBalancerRoutingNamespaceLister interface {
	// List lists all LoadBalancerRoutings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LoadBalancerRouting, err error)
	// Get retrieves the LoadBalancerRouting from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LoadBalancerRouting, error)
	LoadBalancerRoutingNamespaceListerExpansion
}

// loadBalancerRoutingNamespaceLister implements the LoadBalancerRoutingNamespaceLister
// interface.
type loadBalancerRoutingNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.LoadBalancerRouting]
}
