// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// InstanceAntiAffinityApplyConfiguration represents an declarative configuration of the InstanceAntiAffinity type for use
// with apply.
type InstanceAntiAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution []InstanceAffinityTermApplyConfiguration `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// InstanceAntiAffinityApplyConfiguration constructs an declarative configuration of the InstanceAntiAffinity type for use with
// apply.
func InstanceAntiAffinity() *InstanceAntiAffinityApplyConfiguration {
	return &InstanceAntiAffinityApplyConfiguration{}
}

// WithRequiredDuringSchedulingIgnoredDuringExecution adds the given value to the RequiredDuringSchedulingIgnoredDuringExecution field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (b *InstanceAntiAffinityApplyConfiguration) WithRequiredDuringSchedulingIgnoredDuringExecution(values ...*InstanceAffinityTermApplyConfiguration) *InstanceAntiAffinityApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRequiredDuringSchedulingIgnoredDuringExecution")
		}
		b.RequiredDuringSchedulingIgnoredDuringExecution = append(b.RequiredDuringSchedulingIgnoredDuringExecution, *values[i])
	}
	return b
}
