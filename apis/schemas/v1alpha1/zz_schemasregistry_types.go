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

type SchemasRegistryObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SchemasRegistryParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SchemasRegistrySpec defines the desired state of SchemasRegistry
type SchemasRegistrySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SchemasRegistryParameters `json:"forProvider"`
}

// SchemasRegistryStatus defines the observed state of SchemasRegistry.
type SchemasRegistryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SchemasRegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SchemasRegistry is the Schema for the SchemasRegistrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SchemasRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemasRegistrySpec   `json:"spec"`
	Status            SchemasRegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemasRegistryList contains a list of SchemasRegistrys
type SchemasRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchemasRegistry `json:"items"`
}

// Repository type metadata.
var (
	SchemasRegistryKind             = "SchemasRegistry"
	SchemasRegistryGroupKind        = schema.GroupKind{Group: Group, Kind: SchemasRegistryKind}.String()
	SchemasRegistryKindAPIVersion   = SchemasRegistryKind + "." + GroupVersion.String()
	SchemasRegistryGroupVersionKind = GroupVersion.WithKind(SchemasRegistryKind)
)

func init() {
	SchemeBuilder.Register(&SchemasRegistry{}, &SchemasRegistryList{})
}
