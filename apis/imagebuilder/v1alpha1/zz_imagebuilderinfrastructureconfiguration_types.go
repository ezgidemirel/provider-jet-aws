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

type ImagebuilderInfrastructureConfigurationObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DateCreated string `json:"dateCreated" tf:"date_created"`

	DateUpdated string `json:"dateUpdated" tf:"date_updated"`
}

type ImagebuilderInfrastructureConfigurationParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	InstanceProfileName string `json:"instanceProfileName" tf:"instance_profile_name"`

	InstanceTypes []string `json:"instanceTypes,omitempty" tf:"instance_types"`

	KeyPair *string `json:"keyPair,omitempty" tf:"key_pair"`

	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	ResourceTags map[string]string `json:"resourceTags,omitempty" tf:"resource_tags"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TerminateInstanceOnFailure *bool `json:"terminateInstanceOnFailure,omitempty" tf:"terminate_instance_on_failure"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {
	S3Logs []S3LogsParameters `json:"s3Logs" tf:"s3_logs"`
}

type S3LogsObservation struct {
}

type S3LogsParameters struct {
	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`

	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix"`
}

// ImagebuilderInfrastructureConfigurationSpec defines the desired state of ImagebuilderInfrastructureConfiguration
type ImagebuilderInfrastructureConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ImagebuilderInfrastructureConfigurationParameters `json:"forProvider"`
}

// ImagebuilderInfrastructureConfigurationStatus defines the observed state of ImagebuilderInfrastructureConfiguration.
type ImagebuilderInfrastructureConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ImagebuilderInfrastructureConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImagebuilderInfrastructureConfiguration is the Schema for the ImagebuilderInfrastructureConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ImagebuilderInfrastructureConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImagebuilderInfrastructureConfigurationSpec   `json:"spec"`
	Status            ImagebuilderInfrastructureConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImagebuilderInfrastructureConfigurationList contains a list of ImagebuilderInfrastructureConfigurations
type ImagebuilderInfrastructureConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagebuilderInfrastructureConfiguration `json:"items"`
}

// Repository type metadata.
var (
	ImagebuilderInfrastructureConfigurationKind             = "ImagebuilderInfrastructureConfiguration"
	ImagebuilderInfrastructureConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: ImagebuilderInfrastructureConfigurationKind}.String()
	ImagebuilderInfrastructureConfigurationKindAPIVersion   = ImagebuilderInfrastructureConfigurationKind + "." + GroupVersion.String()
	ImagebuilderInfrastructureConfigurationGroupVersionKind = GroupVersion.WithKind(ImagebuilderInfrastructureConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&ImagebuilderInfrastructureConfiguration{}, &ImagebuilderInfrastructureConfigurationList{})
}
