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

type DatasyncLocationSmbMountOptionsObservation struct {
}

type DatasyncLocationSmbMountOptionsParameters struct {
	Version *string `json:"version,omitempty" tf:"version"`
}

type DatasyncLocationSmbObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Uri string `json:"uri" tf:"uri"`
}

type DatasyncLocationSmbParameters struct {
	AgentArns []string `json:"agentArns" tf:"agent_arns"`

	Domain *string `json:"domain,omitempty" tf:"domain"`

	MountOptions []DatasyncLocationSmbMountOptionsParameters `json:"mountOptions,omitempty" tf:"mount_options"`

	Password string `json:"password" tf:"password"`

	Region string `json:"region" tf:"-"`

	ServerHostname string `json:"serverHostname" tf:"server_hostname"`

	Subdirectory string `json:"subdirectory" tf:"subdirectory"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	User string `json:"user" tf:"user"`
}

// DatasyncLocationSmbSpec defines the desired state of DatasyncLocationSmb
type DatasyncLocationSmbSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DatasyncLocationSmbParameters `json:"forProvider"`
}

// DatasyncLocationSmbStatus defines the observed state of DatasyncLocationSmb.
type DatasyncLocationSmbStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DatasyncLocationSmbObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncLocationSmb is the Schema for the DatasyncLocationSmbs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DatasyncLocationSmb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncLocationSmbSpec   `json:"spec"`
	Status            DatasyncLocationSmbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncLocationSmbList contains a list of DatasyncLocationSmbs
type DatasyncLocationSmbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncLocationSmb `json:"items"`
}

// Repository type metadata.
var (
	DatasyncLocationSmbKind             = "DatasyncLocationSmb"
	DatasyncLocationSmbGroupKind        = schema.GroupKind{Group: Group, Kind: DatasyncLocationSmbKind}.String()
	DatasyncLocationSmbKindAPIVersion   = DatasyncLocationSmbKind + "." + GroupVersion.String()
	DatasyncLocationSmbGroupVersionKind = GroupVersion.WithKind(DatasyncLocationSmbKind)
)

func init() {
	SchemeBuilder.Register(&DatasyncLocationSmb{}, &DatasyncLocationSmbList{})
}
