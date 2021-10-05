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

type Route53VpcAssociationAuthorizationObservation struct {
}

type Route53VpcAssociationAuthorizationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	VpcID string `json:"vpcId" tf:"vpc_id"`

	// +kubebuilder:validation:Optional
	VpcRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region"`

	// +kubebuilder:validation:Required
	ZoneID string `json:"zoneId" tf:"zone_id"`
}

// Route53VpcAssociationAuthorizationSpec defines the desired state of Route53VpcAssociationAuthorization
type Route53VpcAssociationAuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Route53VpcAssociationAuthorizationParameters `json:"forProvider"`
}

// Route53VpcAssociationAuthorizationStatus defines the observed state of Route53VpcAssociationAuthorization.
type Route53VpcAssociationAuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Route53VpcAssociationAuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53VpcAssociationAuthorization is the Schema for the Route53VpcAssociationAuthorizations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53VpcAssociationAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53VpcAssociationAuthorizationSpec   `json:"spec"`
	Status            Route53VpcAssociationAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53VpcAssociationAuthorizationList contains a list of Route53VpcAssociationAuthorizations
type Route53VpcAssociationAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53VpcAssociationAuthorization `json:"items"`
}

// Repository type metadata.
var (
	Route53VpcAssociationAuthorizationKind             = "Route53VpcAssociationAuthorization"
	Route53VpcAssociationAuthorizationGroupKind        = schema.GroupKind{Group: Group, Kind: Route53VpcAssociationAuthorizationKind}.String()
	Route53VpcAssociationAuthorizationKindAPIVersion   = Route53VpcAssociationAuthorizationKind + "." + GroupVersion.String()
	Route53VpcAssociationAuthorizationGroupVersionKind = GroupVersion.WithKind(Route53VpcAssociationAuthorizationKind)
)

func init() {
	SchemeBuilder.Register(&Route53VpcAssociationAuthorization{}, &Route53VpcAssociationAuthorizationList{})
}
