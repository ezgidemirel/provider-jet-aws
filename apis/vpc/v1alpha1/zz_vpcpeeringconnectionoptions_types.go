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

type VpcPeeringConnectionOptionsAccepterObservation struct {
}

type VpcPeeringConnectionOptionsAccepterParameters struct {
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc"`

	AllowRemoteVpcDnsResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution"`

	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link"`
}

type VpcPeeringConnectionOptionsObservation struct {
}

type VpcPeeringConnectionOptionsParameters struct {
	Accepter []VpcPeeringConnectionOptionsAccepterParameters `json:"accepter,omitempty" tf:"accepter"`

	Region string `json:"region" tf:"-"`

	Requester []VpcPeeringConnectionOptionsRequesterParameters `json:"requester,omitempty" tf:"requester"`

	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId" tf:"vpc_peering_connection_id"`
}

type VpcPeeringConnectionOptionsRequesterObservation struct {
}

type VpcPeeringConnectionOptionsRequesterParameters struct {
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc"`

	AllowRemoteVpcDnsResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution"`

	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link"`
}

// VpcPeeringConnectionOptionsSpec defines the desired state of VpcPeeringConnectionOptions
type VpcPeeringConnectionOptionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpcPeeringConnectionOptionsParameters `json:"forProvider"`
}

// VpcPeeringConnectionOptionsStatus defines the observed state of VpcPeeringConnectionOptions.
type VpcPeeringConnectionOptionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpcPeeringConnectionOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcPeeringConnectionOptions is the Schema for the VpcPeeringConnectionOptionss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type VpcPeeringConnectionOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionOptionsSpec   `json:"spec"`
	Status            VpcPeeringConnectionOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcPeeringConnectionOptionsList contains a list of VpcPeeringConnectionOptionss
type VpcPeeringConnectionOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcPeeringConnectionOptions `json:"items"`
}

// Repository type metadata.
var (
	VpcPeeringConnectionOptionsKind             = "VpcPeeringConnectionOptions"
	VpcPeeringConnectionOptionsGroupKind        = schema.GroupKind{Group: Group, Kind: VpcPeeringConnectionOptionsKind}.String()
	VpcPeeringConnectionOptionsKindAPIVersion   = VpcPeeringConnectionOptionsKind + "." + GroupVersion.String()
	VpcPeeringConnectionOptionsGroupVersionKind = GroupVersion.WithKind(VpcPeeringConnectionOptionsKind)
)

func init() {
	SchemeBuilder.Register(&VpcPeeringConnectionOptions{}, &VpcPeeringConnectionOptionsList{})
}
