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

type EksFargateProfileObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Status string `json:"status" tf:"status"`
}

type EksFargateProfileParameters struct {
	ClusterName string `json:"clusterName" tf:"cluster_name"`

	FargateProfileName string `json:"fargateProfileName" tf:"fargate_profile_name"`

	PodExecutionRoleArn string `json:"podExecutionRoleArn" tf:"pod_execution_role_arn"`

	Region string `json:"region" tf:"-"`

	Selector []SelectorParameters `json:"selector" tf:"selector"`

	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type SelectorObservation struct {
}

type SelectorParameters struct {
	Labels map[string]string `json:"labels,omitempty" tf:"labels"`

	Namespace string `json:"namespace" tf:"namespace"`
}

// EksFargateProfileSpec defines the desired state of EksFargateProfile
type EksFargateProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EksFargateProfileParameters `json:"forProvider"`
}

// EksFargateProfileStatus defines the observed state of EksFargateProfile.
type EksFargateProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EksFargateProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EksFargateProfile is the Schema for the EksFargateProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EksFargateProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksFargateProfileSpec   `json:"spec"`
	Status            EksFargateProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksFargateProfileList contains a list of EksFargateProfiles
type EksFargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksFargateProfile `json:"items"`
}

// Repository type metadata.
var (
	EksFargateProfileKind             = "EksFargateProfile"
	EksFargateProfileGroupKind        = schema.GroupKind{Group: Group, Kind: EksFargateProfileKind}.String()
	EksFargateProfileKindAPIVersion   = EksFargateProfileKind + "." + GroupVersion.String()
	EksFargateProfileGroupVersionKind = GroupVersion.WithKind(EksFargateProfileKind)
)

func init() {
	SchemeBuilder.Register(&EksFargateProfile{}, &EksFargateProfileList{})
}
