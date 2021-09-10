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

type ByteMatchTuplesObservation struct {
}

type ByteMatchTuplesParameters struct {
	FieldToMatch []FieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`

	PositionalConstraint string `json:"positionalConstraint" tf:"positional_constraint"`

	TargetString *string `json:"targetString,omitempty" tf:"target_string"`

	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

type FieldToMatchObservation struct {
}

type FieldToMatchParameters struct {
	Data *string `json:"data,omitempty" tf:"data"`

	Type string `json:"type" tf:"type"`
}

type WafregionalByteMatchSetObservation struct {
}

type WafregionalByteMatchSetParameters struct {
	ByteMatchTuples []ByteMatchTuplesParameters `json:"byteMatchTuples,omitempty" tf:"byte_match_tuples"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`
}

// WafregionalByteMatchSetSpec defines the desired state of WafregionalByteMatchSet
type WafregionalByteMatchSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalByteMatchSetParameters `json:"forProvider"`
}

// WafregionalByteMatchSetStatus defines the observed state of WafregionalByteMatchSet.
type WafregionalByteMatchSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalByteMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalByteMatchSet is the Schema for the WafregionalByteMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafregionalByteMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalByteMatchSetSpec   `json:"spec"`
	Status            WafregionalByteMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalByteMatchSetList contains a list of WafregionalByteMatchSets
type WafregionalByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalByteMatchSet `json:"items"`
}

// Repository type metadata.
var (
	WafregionalByteMatchSetKind             = "WafregionalByteMatchSet"
	WafregionalByteMatchSetGroupKind        = schema.GroupKind{Group: Group, Kind: WafregionalByteMatchSetKind}.String()
	WafregionalByteMatchSetKindAPIVersion   = WafregionalByteMatchSetKind + "." + GroupVersion.String()
	WafregionalByteMatchSetGroupVersionKind = GroupVersion.WithKind(WafregionalByteMatchSetKind)
)

func init() {
	SchemeBuilder.Register(&WafregionalByteMatchSet{}, &WafregionalByteMatchSetList{})
}
