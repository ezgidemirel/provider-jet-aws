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

type ConfigDeliveryChannelObservation struct {
}

type ConfigDeliveryChannelParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	Region string `json:"region" tf:"-"`

	S3BucketName string `json:"s3BucketName" tf:"s3_bucket_name"`

	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty" tf:"s3_key_prefix"`

	SnapshotDeliveryProperties []SnapshotDeliveryPropertiesParameters `json:"snapshotDeliveryProperties,omitempty" tf:"snapshot_delivery_properties"`

	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn"`
}

type SnapshotDeliveryPropertiesObservation struct {
}

type SnapshotDeliveryPropertiesParameters struct {
	DeliveryFrequency *string `json:"deliveryFrequency,omitempty" tf:"delivery_frequency"`
}

// ConfigDeliveryChannelSpec defines the desired state of ConfigDeliveryChannel
type ConfigDeliveryChannelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ConfigDeliveryChannelParameters `json:"forProvider"`
}

// ConfigDeliveryChannelStatus defines the observed state of ConfigDeliveryChannel.
type ConfigDeliveryChannelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ConfigDeliveryChannelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigDeliveryChannel is the Schema for the ConfigDeliveryChannels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ConfigDeliveryChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigDeliveryChannelSpec   `json:"spec"`
	Status            ConfigDeliveryChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigDeliveryChannelList contains a list of ConfigDeliveryChannels
type ConfigDeliveryChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDeliveryChannel `json:"items"`
}

// Repository type metadata.
var (
	ConfigDeliveryChannelKind             = "ConfigDeliveryChannel"
	ConfigDeliveryChannelGroupKind        = schema.GroupKind{Group: Group, Kind: ConfigDeliveryChannelKind}.String()
	ConfigDeliveryChannelKindAPIVersion   = ConfigDeliveryChannelKind + "." + GroupVersion.String()
	ConfigDeliveryChannelGroupVersionKind = GroupVersion.WithKind(ConfigDeliveryChannelKind)
)

func init() {
	SchemeBuilder.Register(&ConfigDeliveryChannel{}, &ConfigDeliveryChannelList{})
}
