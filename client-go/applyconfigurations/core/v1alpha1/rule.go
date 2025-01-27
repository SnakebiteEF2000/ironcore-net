// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// RuleApplyConfiguration represents a declarative configuration of the Rule type for use
// with apply.
type RuleApplyConfiguration struct {
	CIDRBlock          []IPBlockApplyConfiguration           `json:"ipBlock,omitempty"`
	ObjectIPs          []ObjectIPApplyConfiguration          `json:"ips,omitempty"`
	NetworkPolicyPorts []NetworkPolicyPortApplyConfiguration `json:"networkPolicyPorts,omitempty"`
}

// RuleApplyConfiguration constructs a declarative configuration of the Rule type for use with
// apply.
func Rule() *RuleApplyConfiguration {
	return &RuleApplyConfiguration{}
}

// WithCIDRBlock adds the given value to the CIDRBlock field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CIDRBlock field.
func (b *RuleApplyConfiguration) WithCIDRBlock(values ...*IPBlockApplyConfiguration) *RuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCIDRBlock")
		}
		b.CIDRBlock = append(b.CIDRBlock, *values[i])
	}
	return b
}

// WithObjectIPs adds the given value to the ObjectIPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ObjectIPs field.
func (b *RuleApplyConfiguration) WithObjectIPs(values ...*ObjectIPApplyConfiguration) *RuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithObjectIPs")
		}
		b.ObjectIPs = append(b.ObjectIPs, *values[i])
	}
	return b
}

// WithNetworkPolicyPorts adds the given value to the NetworkPolicyPorts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NetworkPolicyPorts field.
func (b *RuleApplyConfiguration) WithNetworkPolicyPorts(values ...*NetworkPolicyPortApplyConfiguration) *RuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNetworkPolicyPorts")
		}
		b.NetworkPolicyPorts = append(b.NetworkPolicyPorts, *values[i])
	}
	return b
}
