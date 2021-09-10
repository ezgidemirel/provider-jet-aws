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

type Wafv2IpSetObservation struct {
	Arn string `json:"arn" tf:"arn"`

	LockToken string `json:"lockToken" tf:"lock_token"`
}

type Wafv2IpSetParameters struct {
	Addresses []string `json:"addresses,omitempty" tf:"addresses"`

	Description *string `json:"description,omitempty" tf:"description"`

	IpAddressVersion string `json:"ipAddressVersion" tf:"ip_address_version"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	Scope string `json:"scope" tf:"scope"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Wafv2IpSetSpec defines the desired state of Wafv2IpSet
type Wafv2IpSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Wafv2IpSetParameters `json:"forProvider"`
}

// Wafv2IpSetStatus defines the observed state of Wafv2IpSet.
type Wafv2IpSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Wafv2IpSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2IpSet is the Schema for the Wafv2IpSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Wafv2IpSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Wafv2IpSetSpec   `json:"spec"`
	Status            Wafv2IpSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2IpSetList contains a list of Wafv2IpSets
type Wafv2IpSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2IpSet `json:"items"`
}

// Repository type metadata.
var (
	Wafv2IpSetKind             = "Wafv2IpSet"
	Wafv2IpSetGroupKind        = schema.GroupKind{Group: Group, Kind: Wafv2IpSetKind}.String()
	Wafv2IpSetKindAPIVersion   = Wafv2IpSetKind + "." + GroupVersion.String()
	Wafv2IpSetGroupVersionKind = GroupVersion.WithKind(Wafv2IpSetKind)
)

func init() {
	SchemeBuilder.Register(&Wafv2IpSet{}, &Wafv2IpSetList{})
}
