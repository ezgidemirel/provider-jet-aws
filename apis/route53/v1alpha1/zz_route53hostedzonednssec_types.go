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

type Route53HostedZoneDnssecObservation struct {
}

type Route53HostedZoneDnssecParameters struct {

	// +kubebuilder:validation:Required
	HostedZoneID string `json:"hostedZoneId" tf:"hosted_zone_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SigningStatus *string `json:"signingStatus,omitempty" tf:"signing_status"`
}

// Route53HostedZoneDnssecSpec defines the desired state of Route53HostedZoneDnssec
type Route53HostedZoneDnssecSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Route53HostedZoneDnssecParameters `json:"forProvider"`
}

// Route53HostedZoneDnssecStatus defines the observed state of Route53HostedZoneDnssec.
type Route53HostedZoneDnssecStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Route53HostedZoneDnssecObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53HostedZoneDnssec is the Schema for the Route53HostedZoneDnssecs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53HostedZoneDnssec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53HostedZoneDnssecSpec   `json:"spec"`
	Status            Route53HostedZoneDnssecStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53HostedZoneDnssecList contains a list of Route53HostedZoneDnssecs
type Route53HostedZoneDnssecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53HostedZoneDnssec `json:"items"`
}

// Repository type metadata.
var (
	Route53HostedZoneDnssecKind             = "Route53HostedZoneDnssec"
	Route53HostedZoneDnssecGroupKind        = schema.GroupKind{Group: Group, Kind: Route53HostedZoneDnssecKind}.String()
	Route53HostedZoneDnssecKindAPIVersion   = Route53HostedZoneDnssecKind + "." + GroupVersion.String()
	Route53HostedZoneDnssecGroupVersionKind = GroupVersion.WithKind(Route53HostedZoneDnssecKind)
)

func init() {
	SchemeBuilder.Register(&Route53HostedZoneDnssec{}, &Route53HostedZoneDnssecList{})
}
