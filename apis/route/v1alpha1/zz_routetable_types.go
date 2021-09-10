/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type RouteTableObservation struct {
	Arn string `json:"arn" tf:"arn"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type RouteTableParameters struct {
	PropagatingVgws []string `json:"propagatingVgws,omitempty" tf:"propagating_vgws"`

	Region string `json:"region" tf:"-"`

	Route []RouteTableRouteParameters `json:"route,omitempty" tf:"route"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type RouteTableRouteObservation struct {
}

type RouteTableRouteParameters struct {
	CarrierGatewayId *string `json:"carrierGatewayId,omitempty" tf:"carrier_gateway_id"`

	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	DestinationPrefixListId *string `json:"destinationPrefixListId,omitempty" tf:"destination_prefix_list_id"`

	EgressOnlyGatewayId *string `json:"egressOnlyGatewayId,omitempty" tf:"egress_only_gateway_id"`

	GatewayId *string `json:"gatewayId,omitempty" tf:"gateway_id"`

	InstanceId *string `json:"instanceId,omitempty" tf:"instance_id"`

	Ipv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	LocalGatewayId *string `json:"localGatewayId,omitempty" tf:"local_gateway_id"`

	NatGatewayId *string `json:"natGatewayId,omitempty" tf:"nat_gateway_id"`

	NetworkInterfaceId *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id"`

	TransitGatewayId *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id"`

	VpcEndpointId *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	VpcPeeringConnectionId *string `json:"vpcPeeringConnectionId,omitempty" tf:"vpc_peering_connection_id"`
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RouteTableParameters `json:"forProvider"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable is the Schema for the RouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTableKind             = "RouteTable"
	RouteTableGroupKind        = schema.GroupKind{Group: Group, Kind: RouteTableKind}.String()
	RouteTableKindAPIVersion   = RouteTableKind + "." + GroupVersion.String()
	RouteTableGroupVersionKind = GroupVersion.WithKind(RouteTableKind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
