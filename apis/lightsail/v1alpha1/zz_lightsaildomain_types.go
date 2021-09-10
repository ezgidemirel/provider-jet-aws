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

type LightsailDomainObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type LightsailDomainParameters struct {
	DomainName string `json:"domainName" tf:"domain_name"`

	Region string `json:"region" tf:"-"`
}

// LightsailDomainSpec defines the desired state of LightsailDomain
type LightsailDomainSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LightsailDomainParameters `json:"forProvider"`
}

// LightsailDomainStatus defines the observed state of LightsailDomain.
type LightsailDomainStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LightsailDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailDomain is the Schema for the LightsailDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LightsailDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailDomainSpec   `json:"spec"`
	Status            LightsailDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailDomainList contains a list of LightsailDomains
type LightsailDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailDomain `json:"items"`
}

// Repository type metadata.
var (
	LightsailDomainKind             = "LightsailDomain"
	LightsailDomainGroupKind        = schema.GroupKind{Group: Group, Kind: LightsailDomainKind}.String()
	LightsailDomainKindAPIVersion   = LightsailDomainKind + "." + GroupVersion.String()
	LightsailDomainGroupVersionKind = GroupVersion.WithKind(LightsailDomainKind)
)

func init() {
	SchemeBuilder.Register(&LightsailDomain{}, &LightsailDomainList{})
}
