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

type AutoscalingNotificationObservation struct {
}

type AutoscalingNotificationParameters struct {
	GroupNames []string `json:"groupNames" tf:"group_names"`

	Notifications []string `json:"notifications" tf:"notifications"`

	Region string `json:"region" tf:"-"`

	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

// AutoscalingNotificationSpec defines the desired state of AutoscalingNotification
type AutoscalingNotificationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AutoscalingNotificationParameters `json:"forProvider"`
}

// AutoscalingNotificationStatus defines the observed state of AutoscalingNotification.
type AutoscalingNotificationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AutoscalingNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingNotification is the Schema for the AutoscalingNotifications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AutoscalingNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingNotificationSpec   `json:"spec"`
	Status            AutoscalingNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingNotificationList contains a list of AutoscalingNotifications
type AutoscalingNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingNotification `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingNotificationKind             = "AutoscalingNotification"
	AutoscalingNotificationGroupKind        = schema.GroupKind{Group: Group, Kind: AutoscalingNotificationKind}.String()
	AutoscalingNotificationKindAPIVersion   = AutoscalingNotificationKind + "." + GroupVersion.String()
	AutoscalingNotificationGroupVersionKind = GroupVersion.WithKind(AutoscalingNotificationKind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingNotification{}, &AutoscalingNotificationList{})
}
