// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// NetworkSpecApplyConfiguration represents a declarative configuration of the NetworkSpec type for use
// with apply.
type NetworkSpecApplyConfiguration struct {
	ID       *string                            `json:"id,omitempty"`
	Peerings []NetworkPeeringApplyConfiguration `json:"peerings,omitempty"`
}

// NetworkSpecApplyConfiguration constructs a declarative configuration of the NetworkSpec type for use with
// apply.
func NetworkSpec() *NetworkSpecApplyConfiguration {
	return &NetworkSpecApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *NetworkSpecApplyConfiguration) WithID(value string) *NetworkSpecApplyConfiguration {
	b.ID = &value
	return b
}

// WithPeerings adds the given value to the Peerings field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Peerings field.
func (b *NetworkSpecApplyConfiguration) WithPeerings(values ...*NetworkPeeringApplyConfiguration) *NetworkSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPeerings")
		}
		b.Peerings = append(b.Peerings, *values[i])
	}
	return b
}
