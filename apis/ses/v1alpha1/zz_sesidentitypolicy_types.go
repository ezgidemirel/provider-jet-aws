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

type SesIdentityPolicyObservation struct {
}

type SesIdentityPolicyParameters struct {
	Identity string `json:"identity" tf:"identity"`

	Name string `json:"name" tf:"name"`

	Policy string `json:"policy" tf:"policy"`

	Region string `json:"region" tf:"-"`
}

// SesIdentityPolicySpec defines the desired state of SesIdentityPolicy
type SesIdentityPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SesIdentityPolicyParameters `json:"forProvider"`
}

// SesIdentityPolicyStatus defines the observed state of SesIdentityPolicy.
type SesIdentityPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SesIdentityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SesIdentityPolicy is the Schema for the SesIdentityPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SesIdentityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesIdentityPolicySpec   `json:"spec"`
	Status            SesIdentityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesIdentityPolicyList contains a list of SesIdentityPolicys
type SesIdentityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesIdentityPolicy `json:"items"`
}

// Repository type metadata.
var (
	SesIdentityPolicyKind             = "SesIdentityPolicy"
	SesIdentityPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: SesIdentityPolicyKind}.String()
	SesIdentityPolicyKindAPIVersion   = SesIdentityPolicyKind + "." + GroupVersion.String()
	SesIdentityPolicyGroupVersionKind = GroupVersion.WithKind(SesIdentityPolicyKind)
)

func init() {
	SchemeBuilder.Register(&SesIdentityPolicy{}, &SesIdentityPolicyList{})
}
