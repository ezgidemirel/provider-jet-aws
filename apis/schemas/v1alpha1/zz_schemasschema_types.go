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

type SchemasSchemaObservation struct {
	Arn string `json:"arn" tf:"arn"`

	LastModified string `json:"lastModified" tf:"last_modified"`

	Version string `json:"version" tf:"version"`

	VersionCreatedDate string `json:"versionCreatedDate" tf:"version_created_date"`
}

type SchemasSchemaParameters struct {
	Content string `json:"content" tf:"content"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	RegistryName string `json:"registryName" tf:"registry_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Type string `json:"type" tf:"type"`
}

// SchemasSchemaSpec defines the desired state of SchemasSchema
type SchemasSchemaSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SchemasSchemaParameters `json:"forProvider"`
}

// SchemasSchemaStatus defines the observed state of SchemasSchema.
type SchemasSchemaStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SchemasSchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SchemasSchema is the Schema for the SchemasSchemas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SchemasSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchemasSchemaSpec   `json:"spec"`
	Status            SchemasSchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemasSchemaList contains a list of SchemasSchemas
type SchemasSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SchemasSchema `json:"items"`
}

// Repository type metadata.
var (
	SchemasSchemaKind             = "SchemasSchema"
	SchemasSchemaGroupKind        = schema.GroupKind{Group: Group, Kind: SchemasSchemaKind}.String()
	SchemasSchemaKindAPIVersion   = SchemasSchemaKind + "." + GroupVersion.String()
	SchemasSchemaGroupVersionKind = GroupVersion.WithKind(SchemasSchemaKind)
)

func init() {
	SchemeBuilder.Register(&SchemasSchema{}, &SchemasSchemaList{})
}
