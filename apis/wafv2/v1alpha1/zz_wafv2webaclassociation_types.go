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

type Wafv2WebAclAssociationObservation struct {
}

type Wafv2WebAclAssociationParameters struct {
	Region string `json:"region" tf:"-"`

	ResourceArn string `json:"resourceArn" tf:"resource_arn"`

	WebAclArn string `json:"webAclArn" tf:"web_acl_arn"`
}

// Wafv2WebAclAssociationSpec defines the desired state of Wafv2WebAclAssociation
type Wafv2WebAclAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Wafv2WebAclAssociationParameters `json:"forProvider"`
}

// Wafv2WebAclAssociationStatus defines the observed state of Wafv2WebAclAssociation.
type Wafv2WebAclAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Wafv2WebAclAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclAssociation is the Schema for the Wafv2WebAclAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Wafv2WebAclAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Wafv2WebAclAssociationSpec   `json:"spec"`
	Status            Wafv2WebAclAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclAssociationList contains a list of Wafv2WebAclAssociations
type Wafv2WebAclAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2WebAclAssociation `json:"items"`
}

// Repository type metadata.
var (
	Wafv2WebAclAssociationKind             = "Wafv2WebAclAssociation"
	Wafv2WebAclAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: Wafv2WebAclAssociationKind}.String()
	Wafv2WebAclAssociationKindAPIVersion   = Wafv2WebAclAssociationKind + "." + GroupVersion.String()
	Wafv2WebAclAssociationGroupVersionKind = GroupVersion.WithKind(Wafv2WebAclAssociationKind)
)

func init() {
	SchemeBuilder.Register(&Wafv2WebAclAssociation{}, &Wafv2WebAclAssociationList{})
}
