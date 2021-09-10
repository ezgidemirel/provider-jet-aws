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

type IamPolicyAttachmentObservation struct {
}

type IamPolicyAttachmentParameters struct {
	Groups []string `json:"groups,omitempty" tf:"groups"`

	Name string `json:"name" tf:"name"`

	PolicyArn string `json:"policyArn" tf:"policy_arn"`

	Region string `json:"region" tf:"-"`

	Roles []string `json:"roles,omitempty" tf:"roles"`

	Users []string `json:"users,omitempty" tf:"users"`
}

// IamPolicyAttachmentSpec defines the desired state of IamPolicyAttachment
type IamPolicyAttachmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamPolicyAttachmentParameters `json:"forProvider"`
}

// IamPolicyAttachmentStatus defines the observed state of IamPolicyAttachment.
type IamPolicyAttachmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamPolicyAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamPolicyAttachment is the Schema for the IamPolicyAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamPolicyAttachmentSpec   `json:"spec"`
	Status            IamPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamPolicyAttachmentList contains a list of IamPolicyAttachments
type IamPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamPolicyAttachment `json:"items"`
}

// Repository type metadata.
var (
	IamPolicyAttachmentKind             = "IamPolicyAttachment"
	IamPolicyAttachmentGroupKind        = schema.GroupKind{Group: Group, Kind: IamPolicyAttachmentKind}.String()
	IamPolicyAttachmentKindAPIVersion   = IamPolicyAttachmentKind + "." + GroupVersion.String()
	IamPolicyAttachmentGroupVersionKind = GroupVersion.WithKind(IamPolicyAttachmentKind)
)

func init() {
	SchemeBuilder.Register(&IamPolicyAttachment{}, &IamPolicyAttachmentList{})
}
