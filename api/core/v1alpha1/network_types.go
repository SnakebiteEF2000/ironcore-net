// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"github.com/ironcore-dev/ironcore-net/apimachinery/api/net"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkSpec struct {
	// ID is the ID of the network.
	ID string `json:"id,omitempty"`
	// Peerings are the network peerings with this network
	Peerings []NetworkPeering `json:"peerings,omitempty"`
}

// NetworkPeering defines a network peering with another network.
type NetworkPeering struct {
	// Name is the semantical name of the network peering.
	Name string `json:"name"`
	// ID is the ID of the network to peer with.
	ID string `json:"id"`
	// Prefixes is a list of prefixes that we want only to be exposed
	// to the peered network, if no prefixes are specified no filtering will be done.
	Prefixes []PeeringPrefix `json:"prefixes,omitempty"`
}

// PeeringPrefix defines prefixes to be exposed to the peered network
type PeeringPrefix struct {
	// Name is the semantical name of the peering prefixes
	Name string `json:"name"`
	// CIDR to be exposed to the peered network
	Prefix *net.IPPrefix `json:"prefix,omitempty"`
}

type NetworkStatus struct {
	// Peerings contains the states of the network peerings for the network.
	Peerings map[string][]NetworkPeeringStatus `json:"peerings,omitempty"`
}

// NetworkState is the state of a network.
// +enum
type NetworkState string

// NetworkPeeringState is the state a NetworkPeering can be in
type NetworkPeeringState string

const (
	// NetworkPeeringStatePending signals that the network peering is not applied.
	NetworkPeeringStatePending NetworkPeeringState = "Pending"
	// NetworkPeeringStateReady signals that the network peering is ready.
	NetworkPeeringStateReady NetworkPeeringState = "Ready"
	// NetworkPeeringStateError signals that the network peering is in error state.
	NetworkPeeringStateError NetworkPeeringState = "Error"
)

// NetworkPeeringStatus is the status of a network peering.
type NetworkPeeringStatus struct {
	// ID is the ID of network
	ID int32 `json:"id"`
	// State represents the network peering state
	State NetworkPeeringState `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient

// Network is the schema for the networks API.
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSpec   `json:"spec,omitempty"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList contains a list of Network.
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}
