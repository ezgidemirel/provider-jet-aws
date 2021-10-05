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

type DefaultSubnetObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	AssignIPv6AddressOnCreation bool `json:"assignIpv6AddressOnCreation,omitempty" tf:"assign_ipv6_address_on_creation"`

	AvailabilityZoneID string `json:"availabilityZoneId,omitempty" tf:"availability_zone_id"`

	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block"`

	IPv6CidrBlock string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block"`

	IPv6CidrBlockAssociationID string `json:"ipv6CidrBlockAssociationId,omitempty" tf:"ipv6_cidr_block_association_id"`

	OwnerID string `json:"ownerId,omitempty" tf:"owner_id"`

	VpcID string `json:"vpcId,omitempty" tf:"vpc_id"`
}

type DefaultSubnetParameters struct {

	// +kubebuilder:validation:Required
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	// +kubebuilder:validation:Optional
	CustomerOwnedIPv4Pool *string `json:"customerOwnedIpv4Pool,omitempty" tf:"customer_owned_ipv4_pool"`

	// +kubebuilder:validation:Optional
	MapCustomerOwnedIPOnLaunch *bool `json:"mapCustomerOwnedIpOnLaunch,omitempty" tf:"map_customer_owned_ip_on_launch"`

	// +kubebuilder:validation:Optional
	MapPublicIPOnLaunch *bool `json:"mapPublicIpOnLaunch,omitempty" tf:"map_public_ip_on_launch"`

	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DefaultSubnetSpec defines the desired state of DefaultSubnet
type DefaultSubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultSubnetParameters `json:"forProvider"`
}

// DefaultSubnetStatus defines the observed state of DefaultSubnet.
type DefaultSubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultSubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSubnet is the Schema for the DefaultSubnets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DefaultSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSubnetSpec   `json:"spec"`
	Status            DefaultSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSubnetList contains a list of DefaultSubnets
type DefaultSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSubnet `json:"items"`
}

// Repository type metadata.
var (
	DefaultSubnetKind             = "DefaultSubnet"
	DefaultSubnetGroupKind        = schema.GroupKind{Group: Group, Kind: DefaultSubnetKind}.String()
	DefaultSubnetKindAPIVersion   = DefaultSubnetKind + "." + GroupVersion.String()
	DefaultSubnetGroupVersionKind = GroupVersion.WithKind(DefaultSubnetKind)
)

func init() {
	SchemeBuilder.Register(&DefaultSubnet{}, &DefaultSubnetList{})
}
