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

type TimestreamwriteDatabaseObservation struct {
	Arn string `json:"arn" tf:"arn"`

	TableCount int64 `json:"tableCount" tf:"table_count"`
}

type TimestreamwriteDatabaseParameters struct {
	DatabaseName string `json:"databaseName" tf:"database_name"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// TimestreamwriteDatabaseSpec defines the desired state of TimestreamwriteDatabase
type TimestreamwriteDatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TimestreamwriteDatabaseParameters `json:"forProvider"`
}

// TimestreamwriteDatabaseStatus defines the observed state of TimestreamwriteDatabase.
type TimestreamwriteDatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TimestreamwriteDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TimestreamwriteDatabase is the Schema for the TimestreamwriteDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TimestreamwriteDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TimestreamwriteDatabaseSpec   `json:"spec"`
	Status            TimestreamwriteDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TimestreamwriteDatabaseList contains a list of TimestreamwriteDatabases
type TimestreamwriteDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TimestreamwriteDatabase `json:"items"`
}

// Repository type metadata.
var (
	TimestreamwriteDatabaseKind             = "TimestreamwriteDatabase"
	TimestreamwriteDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: TimestreamwriteDatabaseKind}.String()
	TimestreamwriteDatabaseKindAPIVersion   = TimestreamwriteDatabaseKind + "." + GroupVersion.String()
	TimestreamwriteDatabaseGroupVersionKind = GroupVersion.WithKind(TimestreamwriteDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&TimestreamwriteDatabase{}, &TimestreamwriteDatabaseList{})
}
