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

type AccepterObservation struct {
}

type AccepterParameters struct {
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc"`

	AllowRemoteVpcDnsResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution"`

	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link"`
}

type RequesterObservation struct {
}

type RequesterParameters struct {
	AllowClassicLinkToRemoteVpc *bool `json:"allowClassicLinkToRemoteVpc,omitempty" tf:"allow_classic_link_to_remote_vpc"`

	AllowRemoteVpcDnsResolution *bool `json:"allowRemoteVpcDnsResolution,omitempty" tf:"allow_remote_vpc_dns_resolution"`

	AllowVpcToRemoteClassicLink *bool `json:"allowVpcToRemoteClassicLink,omitempty" tf:"allow_vpc_to_remote_classic_link"`
}

type VpcPeeringConnectionObservation struct {
	AcceptStatus string `json:"acceptStatus" tf:"accept_status"`
}

type VpcPeeringConnectionParameters struct {
	Accepter []AccepterParameters `json:"accepter,omitempty" tf:"accepter"`

	AutoAccept *bool `json:"autoAccept,omitempty" tf:"auto_accept"`

	PeerOwnerId *string `json:"peerOwnerId,omitempty" tf:"peer_owner_id"`

	PeerRegion *string `json:"peerRegion,omitempty" tf:"peer_region"`

	PeerVpcId string `json:"peerVpcId" tf:"peer_vpc_id"`

	Region string `json:"region" tf:"-"`

	Requester []RequesterParameters `json:"requester,omitempty" tf:"requester"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

// VpcPeeringConnectionSpec defines the desired state of VpcPeeringConnection
type VpcPeeringConnectionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpcPeeringConnectionParameters `json:"forProvider"`
}

// VpcPeeringConnectionStatus defines the observed state of VpcPeeringConnection.
type VpcPeeringConnectionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpcPeeringConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcPeeringConnection is the Schema for the VpcPeeringConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type VpcPeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcPeeringConnectionSpec   `json:"spec"`
	Status            VpcPeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcPeeringConnectionList contains a list of VpcPeeringConnections
type VpcPeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcPeeringConnection `json:"items"`
}

// Repository type metadata.
var (
	VpcPeeringConnectionKind             = "VpcPeeringConnection"
	VpcPeeringConnectionGroupKind        = schema.GroupKind{Group: Group, Kind: VpcPeeringConnectionKind}.String()
	VpcPeeringConnectionKindAPIVersion   = VpcPeeringConnectionKind + "." + GroupVersion.String()
	VpcPeeringConnectionGroupVersionKind = GroupVersion.WithKind(VpcPeeringConnectionKind)
)

func init() {
	SchemeBuilder.Register(&VpcPeeringConnection{}, &VpcPeeringConnectionList{})
}
