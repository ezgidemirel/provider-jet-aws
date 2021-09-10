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

type SesDomainDkimObservation struct {
	DkimTokens []string `json:"dkimTokens" tf:"dkim_tokens"`
}

type SesDomainDkimParameters struct {
	Domain string `json:"domain" tf:"domain"`

	Region string `json:"region" tf:"-"`
}

// SesDomainDkimSpec defines the desired state of SesDomainDkim
type SesDomainDkimSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SesDomainDkimParameters `json:"forProvider"`
}

// SesDomainDkimStatus defines the observed state of SesDomainDkim.
type SesDomainDkimStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SesDomainDkimObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SesDomainDkim is the Schema for the SesDomainDkims API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SesDomainDkim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesDomainDkimSpec   `json:"spec"`
	Status            SesDomainDkimStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesDomainDkimList contains a list of SesDomainDkims
type SesDomainDkimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesDomainDkim `json:"items"`
}

// Repository type metadata.
var (
	SesDomainDkimKind             = "SesDomainDkim"
	SesDomainDkimGroupKind        = schema.GroupKind{Group: Group, Kind: SesDomainDkimKind}.String()
	SesDomainDkimKindAPIVersion   = SesDomainDkimKind + "." + GroupVersion.String()
	SesDomainDkimGroupVersionKind = GroupVersion.WithKind(SesDomainDkimKind)
)

func init() {
	SchemeBuilder.Register(&SesDomainDkim{}, &SesDomainDkimList{})
}
