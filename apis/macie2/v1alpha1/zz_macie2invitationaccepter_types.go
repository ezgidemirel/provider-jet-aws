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

type Macie2InvitationAccepterObservation struct {
	InvitationId string `json:"invitationId" tf:"invitation_id"`
}

type Macie2InvitationAccepterParameters struct {
	AdministratorAccountId string `json:"administratorAccountId" tf:"administrator_account_id"`

	Region string `json:"region" tf:"-"`
}

// Macie2InvitationAccepterSpec defines the desired state of Macie2InvitationAccepter
type Macie2InvitationAccepterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Macie2InvitationAccepterParameters `json:"forProvider"`
}

// Macie2InvitationAccepterStatus defines the observed state of Macie2InvitationAccepter.
type Macie2InvitationAccepterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Macie2InvitationAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2InvitationAccepter is the Schema for the Macie2InvitationAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Macie2InvitationAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Macie2InvitationAccepterSpec   `json:"spec"`
	Status            Macie2InvitationAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2InvitationAccepterList contains a list of Macie2InvitationAccepters
type Macie2InvitationAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macie2InvitationAccepter `json:"items"`
}

// Repository type metadata.
var (
	Macie2InvitationAccepterKind             = "Macie2InvitationAccepter"
	Macie2InvitationAccepterGroupKind        = schema.GroupKind{Group: Group, Kind: Macie2InvitationAccepterKind}.String()
	Macie2InvitationAccepterKindAPIVersion   = Macie2InvitationAccepterKind + "." + GroupVersion.String()
	Macie2InvitationAccepterGroupVersionKind = GroupVersion.WithKind(Macie2InvitationAccepterKind)
)

func init() {
	SchemeBuilder.Register(&Macie2InvitationAccepter{}, &Macie2InvitationAccepterList{})
}
