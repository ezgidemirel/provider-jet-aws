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

type CodeartifactRepositoryObservation struct {
	AdministratorAccount string `json:"administratorAccount" tf:"administrator_account"`

	Arn string `json:"arn" tf:"arn"`
}

type CodeartifactRepositoryParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Domain string `json:"domain" tf:"domain"`

	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner"`

	ExternalConnections []ExternalConnectionsParameters `json:"externalConnections,omitempty" tf:"external_connections"`

	Region string `json:"region" tf:"-"`

	Repository string `json:"repository" tf:"repository"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Upstream []UpstreamParameters `json:"upstream,omitempty" tf:"upstream"`
}

type ExternalConnectionsObservation struct {
	PackageFormat string `json:"packageFormat" tf:"package_format"`

	Status string `json:"status" tf:"status"`
}

type ExternalConnectionsParameters struct {
	ExternalConnectionName string `json:"externalConnectionName" tf:"external_connection_name"`
}

type UpstreamObservation struct {
}

type UpstreamParameters struct {
	RepositoryName string `json:"repositoryName" tf:"repository_name"`
}

// CodeartifactRepositorySpec defines the desired state of CodeartifactRepository
type CodeartifactRepositorySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodeartifactRepositoryParameters `json:"forProvider"`
}

// CodeartifactRepositoryStatus defines the observed state of CodeartifactRepository.
type CodeartifactRepositoryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodeartifactRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactRepository is the Schema for the CodeartifactRepositorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodeartifactRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodeartifactRepositorySpec   `json:"spec"`
	Status            CodeartifactRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactRepositoryList contains a list of CodeartifactRepositorys
type CodeartifactRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeartifactRepository `json:"items"`
}

// Repository type metadata.
var (
	CodeartifactRepositoryKind             = "CodeartifactRepository"
	CodeartifactRepositoryGroupKind        = schema.GroupKind{Group: Group, Kind: CodeartifactRepositoryKind}.String()
	CodeartifactRepositoryKindAPIVersion   = CodeartifactRepositoryKind + "." + GroupVersion.String()
	CodeartifactRepositoryGroupVersionKind = GroupVersion.WithKind(CodeartifactRepositoryKind)
)

func init() {
	SchemeBuilder.Register(&CodeartifactRepository{}, &CodeartifactRepositoryList{})
}
