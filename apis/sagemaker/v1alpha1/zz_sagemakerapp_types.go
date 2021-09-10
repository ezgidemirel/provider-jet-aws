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

type ResourceSpecObservation struct {
}

type ResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
}

type SagemakerAppObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SagemakerAppParameters struct {
	AppName string `json:"appName" tf:"app_name"`

	AppType string `json:"appType" tf:"app_type"`

	DomainId string `json:"domainId" tf:"domain_id"`

	Region string `json:"region" tf:"-"`

	ResourceSpec []ResourceSpecParameters `json:"resourceSpec,omitempty" tf:"resource_spec"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UserProfileName string `json:"userProfileName" tf:"user_profile_name"`
}

// SagemakerAppSpec defines the desired state of SagemakerApp
type SagemakerAppSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerAppParameters `json:"forProvider"`
}

// SagemakerAppStatus defines the observed state of SagemakerApp.
type SagemakerAppStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerApp is the Schema for the SagemakerApps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerAppSpec   `json:"spec"`
	Status            SagemakerAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerAppList contains a list of SagemakerApps
type SagemakerAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerApp `json:"items"`
}

// Repository type metadata.
var (
	SagemakerAppKind             = "SagemakerApp"
	SagemakerAppGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerAppKind}.String()
	SagemakerAppKindAPIVersion   = SagemakerAppKind + "." + GroupVersion.String()
	SagemakerAppGroupVersionKind = GroupVersion.WithKind(SagemakerAppKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerApp{}, &SagemakerAppList{})
}
