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

type VpcObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DefaultNetworkAclId string `json:"defaultNetworkAclId" tf:"default_network_acl_id"`

	DefaultRouteTableId string `json:"defaultRouteTableId" tf:"default_route_table_id"`

	DefaultSecurityGroupId string `json:"defaultSecurityGroupId" tf:"default_security_group_id"`

	DhcpOptionsId string `json:"dhcpOptionsId" tf:"dhcp_options_id"`

	Ipv6AssociationId string `json:"ipv6AssociationId" tf:"ipv6_association_id"`

	Ipv6CidrBlock string `json:"ipv6CidrBlock" tf:"ipv6_cidr_block"`

	MainRouteTableId string `json:"mainRouteTableId" tf:"main_route_table_id"`

	OwnerId string `json:"ownerId" tf:"owner_id"`
}

type VpcParameters struct {
	AssignGeneratedIpv6CidrBlock *bool `json:"assignGeneratedIpv6CidrBlock,omitempty" tf:"assign_generated_ipv6_cidr_block"`

	CidrBlock string `json:"cidrBlock" tf:"cidr_block"`

	EnableClassiclink *bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink"`

	EnableClassiclinkDnsSupport *bool `json:"enableClassiclinkDnsSupport,omitempty" tf:"enable_classiclink_dns_support"`

	EnableDnsHostnames *bool `json:"enableDnsHostnames,omitempty" tf:"enable_dns_hostnames"`

	EnableDnsSupport *bool `json:"enableDnsSupport,omitempty" tf:"enable_dns_support"`

	InstanceTenancy *string `json:"instanceTenancy,omitempty" tf:"instance_tenancy"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// VpcSpec defines the desired state of Vpc
type VpcSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpcParameters `json:"forProvider"`
}

// VpcStatus defines the observed state of Vpc.
type VpcStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpcObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vpc is the Schema for the Vpcs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Vpc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcSpec   `json:"spec"`
	Status            VpcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcList contains a list of Vpcs
type VpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vpc `json:"items"`
}

// Repository type metadata.
var (
	VpcKind             = "Vpc"
	VpcGroupKind        = schema.GroupKind{Group: Group, Kind: VpcKind}.String()
	VpcKindAPIVersion   = VpcKind + "." + GroupVersion.String()
	VpcGroupVersionKind = GroupVersion.WithKind(VpcKind)
)

func init() {
	SchemeBuilder.Register(&Vpc{}, &VpcList{})
}
