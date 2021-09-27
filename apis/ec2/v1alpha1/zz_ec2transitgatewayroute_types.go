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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type Ec2TransitGatewayRouteObservation struct {
}

type Ec2TransitGatewayRouteParameters struct {

	// +kubebuilder:validation:Optional
	Blackhole *bool `json:"blackhole,omitempty" tf:"blackhole"`

	// +kubebuilder:validation:Required
	DestinationCidrBlock string `json:"destinationCidrBlock" tf:"destination_cidr_block"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id"`

	// +kubebuilder:validation:Required
	TransitGatewayRouteTableID string `json:"transitGatewayRouteTableId" tf:"transit_gateway_route_table_id"`
}

// Ec2TransitGatewayRouteSpec defines the desired state of Ec2TransitGatewayRoute
type Ec2TransitGatewayRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Ec2TransitGatewayRouteParameters `json:"forProvider"`
}

// Ec2TransitGatewayRouteStatus defines the observed state of Ec2TransitGatewayRoute.
type Ec2TransitGatewayRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Ec2TransitGatewayRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRoute is the Schema for the Ec2TransitGatewayRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteSpec   `json:"spec"`
	Status            Ec2TransitGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteList contains a list of Ec2TransitGatewayRoutes
type Ec2TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayRoute `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayRouteKind             = "Ec2TransitGatewayRoute"
	Ec2TransitGatewayRouteGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2TransitGatewayRouteKind}.String()
	Ec2TransitGatewayRouteKindAPIVersion   = Ec2TransitGatewayRouteKind + "." + GroupVersion.String()
	Ec2TransitGatewayRouteGroupVersionKind = GroupVersion.WithKind(Ec2TransitGatewayRouteKind)
)

func init() {
	SchemeBuilder.Register(&Ec2TransitGatewayRoute{}, &Ec2TransitGatewayRouteList{})
}
