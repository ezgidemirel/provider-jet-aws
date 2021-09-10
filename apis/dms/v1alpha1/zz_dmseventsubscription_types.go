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

type DmsEventSubscriptionObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DmsEventSubscriptionParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	EventCategories []string `json:"eventCategories" tf:"event_categories"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	SnsTopicArn string `json:"snsTopicArn" tf:"sns_topic_arn"`

	SourceIds []string `json:"sourceIds,omitempty" tf:"source_ids"`

	SourceType *string `json:"sourceType,omitempty" tf:"source_type"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DmsEventSubscriptionSpec defines the desired state of DmsEventSubscription
type DmsEventSubscriptionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DmsEventSubscriptionParameters `json:"forProvider"`
}

// DmsEventSubscriptionStatus defines the observed state of DmsEventSubscription.
type DmsEventSubscriptionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DmsEventSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DmsEventSubscription is the Schema for the DmsEventSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DmsEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsEventSubscriptionSpec   `json:"spec"`
	Status            DmsEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsEventSubscriptionList contains a list of DmsEventSubscriptions
type DmsEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsEventSubscription `json:"items"`
}

// Repository type metadata.
var (
	DmsEventSubscriptionKind             = "DmsEventSubscription"
	DmsEventSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: DmsEventSubscriptionKind}.String()
	DmsEventSubscriptionKindAPIVersion   = DmsEventSubscriptionKind + "." + GroupVersion.String()
	DmsEventSubscriptionGroupVersionKind = GroupVersion.WithKind(DmsEventSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&DmsEventSubscription{}, &DmsEventSubscriptionList{})
}
