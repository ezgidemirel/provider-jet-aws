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

type EksIdentityProviderConfigObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Status string `json:"status" tf:"status"`
}

type EksIdentityProviderConfigOidcObservation struct {
}

type EksIdentityProviderConfigOidcParameters struct {
	ClientId string `json:"clientId" tf:"client_id"`

	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim"`

	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix"`

	IdentityProviderConfigName string `json:"identityProviderConfigName" tf:"identity_provider_config_name"`

	IssuerUrl string `json:"issuerUrl" tf:"issuer_url"`

	RequiredClaims map[string]string `json:"requiredClaims,omitempty" tf:"required_claims"`

	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim"`

	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix"`
}

type EksIdentityProviderConfigParameters struct {
	ClusterName string `json:"clusterName" tf:"cluster_name"`

	Oidc []EksIdentityProviderConfigOidcParameters `json:"oidc" tf:"oidc"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// EksIdentityProviderConfigSpec defines the desired state of EksIdentityProviderConfig
type EksIdentityProviderConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EksIdentityProviderConfigParameters `json:"forProvider"`
}

// EksIdentityProviderConfigStatus defines the observed state of EksIdentityProviderConfig.
type EksIdentityProviderConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EksIdentityProviderConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EksIdentityProviderConfig is the Schema for the EksIdentityProviderConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EksIdentityProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksIdentityProviderConfigSpec   `json:"spec"`
	Status            EksIdentityProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksIdentityProviderConfigList contains a list of EksIdentityProviderConfigs
type EksIdentityProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksIdentityProviderConfig `json:"items"`
}

// Repository type metadata.
var (
	EksIdentityProviderConfigKind             = "EksIdentityProviderConfig"
	EksIdentityProviderConfigGroupKind        = schema.GroupKind{Group: Group, Kind: EksIdentityProviderConfigKind}.String()
	EksIdentityProviderConfigKindAPIVersion   = EksIdentityProviderConfigKind + "." + GroupVersion.String()
	EksIdentityProviderConfigGroupVersionKind = GroupVersion.WithKind(EksIdentityProviderConfigKind)
)

func init() {
	SchemeBuilder.Register(&EksIdentityProviderConfig{}, &EksIdentityProviderConfigList{})
}
