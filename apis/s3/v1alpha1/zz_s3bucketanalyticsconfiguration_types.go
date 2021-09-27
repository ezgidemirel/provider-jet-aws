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

type DataExportDestinationObservation struct {
}

type DataExportDestinationParameters struct {

	// +kubebuilder:validation:Required
	S3BucketDestination []S3BucketDestinationParameters `json:"s3BucketDestination" tf:"s3_bucket_destination"`
}

type DataExportObservation struct {
}

type DataExportParameters struct {

	// +kubebuilder:validation:Required
	Destination []DataExportDestinationParameters `json:"destination" tf:"destination"`

	// +kubebuilder:validation:Optional
	OutputSchemaVersion *string `json:"outputSchemaVersion,omitempty" tf:"output_schema_version"`
}

type S3BucketAnalyticsConfigurationFilterObservation struct {
}

type S3BucketAnalyticsConfigurationFilterParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type S3BucketAnalyticsConfigurationObservation struct {
}

type S3BucketAnalyticsConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Bucket string `json:"bucket" tf:"bucket"`

	// +kubebuilder:validation:Optional
	Filter []S3BucketAnalyticsConfigurationFilterParameters `json:"filter,omitempty" tf:"filter"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageClassAnalysis []StorageClassAnalysisParameters `json:"storageClassAnalysis,omitempty" tf:"storage_class_analysis"`
}

type S3BucketDestinationObservation struct {
}

type S3BucketDestinationParameters struct {

	// +kubebuilder:validation:Optional
	BucketAccountID *string `json:"bucketAccountId,omitempty" tf:"bucket_account_id"`

	// +kubebuilder:validation:Required
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type StorageClassAnalysisObservation struct {
}

type StorageClassAnalysisParameters struct {

	// +kubebuilder:validation:Required
	DataExport []DataExportParameters `json:"dataExport" tf:"data_export"`
}

// S3BucketAnalyticsConfigurationSpec defines the desired state of S3BucketAnalyticsConfiguration
type S3BucketAnalyticsConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3BucketAnalyticsConfigurationParameters `json:"forProvider"`
}

// S3BucketAnalyticsConfigurationStatus defines the observed state of S3BucketAnalyticsConfiguration.
type S3BucketAnalyticsConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3BucketAnalyticsConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketAnalyticsConfiguration is the Schema for the S3BucketAnalyticsConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type S3BucketAnalyticsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketAnalyticsConfigurationSpec   `json:"spec"`
	Status            S3BucketAnalyticsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketAnalyticsConfigurationList contains a list of S3BucketAnalyticsConfigurations
type S3BucketAnalyticsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketAnalyticsConfiguration `json:"items"`
}

// Repository type metadata.
var (
	S3BucketAnalyticsConfigurationKind             = "S3BucketAnalyticsConfiguration"
	S3BucketAnalyticsConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: S3BucketAnalyticsConfigurationKind}.String()
	S3BucketAnalyticsConfigurationKindAPIVersion   = S3BucketAnalyticsConfigurationKind + "." + GroupVersion.String()
	S3BucketAnalyticsConfigurationGroupVersionKind = GroupVersion.WithKind(S3BucketAnalyticsConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&S3BucketAnalyticsConfiguration{}, &S3BucketAnalyticsConfigurationList{})
}
