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

type AllowedPublishersObservation struct {
}

type AllowedPublishersParameters struct {
	SigningProfileVersionArns []string `json:"signingProfileVersionArns" tf:"signing_profile_version_arns"`
}

type LambdaCodeSigningConfigObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ConfigId string `json:"configId" tf:"config_id"`

	LastModified string `json:"lastModified" tf:"last_modified"`
}

type LambdaCodeSigningConfigParameters struct {
	AllowedPublishers []AllowedPublishersParameters `json:"allowedPublishers" tf:"allowed_publishers"`

	Description *string `json:"description,omitempty" tf:"description"`

	Policies []PoliciesParameters `json:"policies,omitempty" tf:"policies"`

	Region string `json:"region" tf:"-"`
}

type PoliciesObservation struct {
}

type PoliciesParameters struct {
	UntrustedArtifactOnDeployment string `json:"untrustedArtifactOnDeployment" tf:"untrusted_artifact_on_deployment"`
}

// LambdaCodeSigningConfigSpec defines the desired state of LambdaCodeSigningConfig
type LambdaCodeSigningConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LambdaCodeSigningConfigParameters `json:"forProvider"`
}

// LambdaCodeSigningConfigStatus defines the observed state of LambdaCodeSigningConfig.
type LambdaCodeSigningConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LambdaCodeSigningConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaCodeSigningConfig is the Schema for the LambdaCodeSigningConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LambdaCodeSigningConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaCodeSigningConfigSpec   `json:"spec"`
	Status            LambdaCodeSigningConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaCodeSigningConfigList contains a list of LambdaCodeSigningConfigs
type LambdaCodeSigningConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaCodeSigningConfig `json:"items"`
}

// Repository type metadata.
var (
	LambdaCodeSigningConfigKind             = "LambdaCodeSigningConfig"
	LambdaCodeSigningConfigGroupKind        = schema.GroupKind{Group: Group, Kind: LambdaCodeSigningConfigKind}.String()
	LambdaCodeSigningConfigKindAPIVersion   = LambdaCodeSigningConfigKind + "." + GroupVersion.String()
	LambdaCodeSigningConfigGroupVersionKind = GroupVersion.WithKind(LambdaCodeSigningConfigKind)
)

func init() {
	SchemeBuilder.Register(&LambdaCodeSigningConfig{}, &LambdaCodeSigningConfigList{})
}
