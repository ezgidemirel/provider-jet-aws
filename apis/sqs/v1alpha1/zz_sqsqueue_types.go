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

type SqsQueueObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Url string `json:"url" tf:"url"`
}

type SqsQueueParameters struct {
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication"`

	DeduplicationScope *string `json:"deduplicationScope,omitempty" tf:"deduplication_scope"`

	DelaySeconds *int64 `json:"delaySeconds,omitempty" tf:"delay_seconds"`

	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue"`

	FifoThroughputLimit *string `json:"fifoThroughputLimit,omitempty" tf:"fifo_throughput_limit"`

	KmsDataKeyReusePeriodSeconds *int64 `json:"kmsDataKeyReusePeriodSeconds,omitempty" tf:"kms_data_key_reuse_period_seconds"`

	KmsMasterKeyId *string `json:"kmsMasterKeyId,omitempty" tf:"kms_master_key_id"`

	MaxMessageSize *int64 `json:"maxMessageSize,omitempty" tf:"max_message_size"`

	MessageRetentionSeconds *int64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Policy *string `json:"policy,omitempty" tf:"policy"`

	ReceiveWaitTimeSeconds *int64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds"`

	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VisibilityTimeoutSeconds *int64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds"`
}

// SqsQueueSpec defines the desired state of SqsQueue
type SqsQueueSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SqsQueueParameters `json:"forProvider"`
}

// SqsQueueStatus defines the observed state of SqsQueue.
type SqsQueueStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SqsQueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqsQueue is the Schema for the SqsQueues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SqsQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqsQueueSpec   `json:"spec"`
	Status            SqsQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqsQueueList contains a list of SqsQueues
type SqsQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqsQueue `json:"items"`
}

// Repository type metadata.
var (
	SqsQueueKind             = "SqsQueue"
	SqsQueueGroupKind        = schema.GroupKind{Group: Group, Kind: SqsQueueKind}.String()
	SqsQueueKindAPIVersion   = SqsQueueKind + "." + GroupVersion.String()
	SqsQueueGroupVersionKind = GroupVersion.WithKind(SqsQueueKind)
)

func init() {
	SchemeBuilder.Register(&SqsQueue{}, &SqsQueueList{})
}
