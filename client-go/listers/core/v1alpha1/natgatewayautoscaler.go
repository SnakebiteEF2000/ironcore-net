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

// NATGatewayAutoscalerLister helps list NATGatewayAutoscalers.
// All objects returned here must be treated as read-only.
type NATGatewayAutoscalerLister interface {
	// List lists all NATGatewayAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NATGatewayAutoscaler, err error)
	// NATGatewayAutoscalers returns an object that can list and get NATGatewayAutoscalers.
	NATGatewayAutoscalers(namespace string) NATGatewayAutoscalerNamespaceLister
	NATGatewayAutoscalerListerExpansion
}

// nATGatewayAutoscalerLister implements the NATGatewayAutoscalerLister interface.
type nATGatewayAutoscalerLister struct {
	listers.ResourceIndexer[*v1alpha1.NATGatewayAutoscaler]
}

// NewNATGatewayAutoscalerLister returns a new NATGatewayAutoscalerLister.
func NewNATGatewayAutoscalerLister(indexer cache.Indexer) NATGatewayAutoscalerLister {
	return &nATGatewayAutoscalerLister{listers.New[*v1alpha1.NATGatewayAutoscaler](indexer, v1alpha1.Resource("natgatewayautoscaler"))}
}

// NATGatewayAutoscalers returns an object that can list and get NATGatewayAutoscalers.
func (s *nATGatewayAutoscalerLister) NATGatewayAutoscalers(namespace string) NATGatewayAutoscalerNamespaceLister {
	return nATGatewayAutoscalerNamespaceLister{listers.NewNamespaced[*v1alpha1.NATGatewayAutoscaler](s.ResourceIndexer, namespace)}
}

// NATGatewayAutoscalerNamespaceLister helps list and get NATGatewayAutoscalers.
// All objects returned here must be treated as read-only.
type NATGatewayAutoscalerNamespaceLister interface {
	// List lists all NATGatewayAutoscalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NATGatewayAutoscaler, err error)
	// Get retrieves the NATGatewayAutoscaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NATGatewayAutoscaler, error)
	NATGatewayAutoscalerNamespaceListerExpansion
}

// nATGatewayAutoscalerNamespaceLister implements the NATGatewayAutoscalerNamespaceLister
// interface.
type nATGatewayAutoscalerNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.NATGatewayAutoscaler]
}
