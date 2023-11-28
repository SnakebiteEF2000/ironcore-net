// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// AffinityApplyConfiguration represents an declarative configuration of the Affinity type for use
// with apply.
type AffinityApplyConfiguration struct {
	NodeAffinity         *NodeAffinityApplyConfiguration         `json:"nodeAffinity,omitempty"`
	InstanceAntiAffinity *InstanceAntiAffinityApplyConfiguration `json:"instanceAntiAffinity,omitempty"`
}

// AffinityApplyConfiguration constructs an declarative configuration of the Affinity type for use with
// apply.
func Affinity() *AffinityApplyConfiguration {
	return &AffinityApplyConfiguration{}
}

// WithNodeAffinity sets the NodeAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeAffinity field is set to the value of the last call.
func (b *AffinityApplyConfiguration) WithNodeAffinity(value *NodeAffinityApplyConfiguration) *AffinityApplyConfiguration {
	b.NodeAffinity = value
	return b
}

// WithInstanceAntiAffinity sets the InstanceAntiAffinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InstanceAntiAffinity field is set to the value of the last call.
func (b *AffinityApplyConfiguration) WithInstanceAntiAffinity(value *InstanceAntiAffinityApplyConfiguration) *AffinityApplyConfiguration {
	b.InstanceAntiAffinity = value
	return b
}
