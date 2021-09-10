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

type AcmpcaCertificateObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Certificate string `json:"certificate" tf:"certificate"`

	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`
}

type AcmpcaCertificateParameters struct {
	CertificateAuthorityArn string `json:"certificateAuthorityArn" tf:"certificate_authority_arn"`

	CertificateSigningRequest string `json:"certificateSigningRequest" tf:"certificate_signing_request"`

	Region string `json:"region" tf:"-"`

	SigningAlgorithm string `json:"signingAlgorithm" tf:"signing_algorithm"`

	TemplateArn *string `json:"templateArn,omitempty" tf:"template_arn"`

	Validity []ValidityParameters `json:"validity" tf:"validity"`
}

type ValidityObservation struct {
}

type ValidityParameters struct {
	Type string `json:"type" tf:"type"`

	Value string `json:"value" tf:"value"`
}

// AcmpcaCertificateSpec defines the desired state of AcmpcaCertificate
type AcmpcaCertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AcmpcaCertificateParameters `json:"forProvider"`
}

// AcmpcaCertificateStatus defines the observed state of AcmpcaCertificate.
type AcmpcaCertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AcmpcaCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AcmpcaCertificate is the Schema for the AcmpcaCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AcmpcaCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AcmpcaCertificateSpec   `json:"spec"`
	Status            AcmpcaCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcmpcaCertificateList contains a list of AcmpcaCertificates
type AcmpcaCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcmpcaCertificate `json:"items"`
}

// Repository type metadata.
var (
	AcmpcaCertificateKind             = "AcmpcaCertificate"
	AcmpcaCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: AcmpcaCertificateKind}.String()
	AcmpcaCertificateKindAPIVersion   = AcmpcaCertificateKind + "." + GroupVersion.String()
	AcmpcaCertificateGroupVersionKind = GroupVersion.WithKind(AcmpcaCertificateKind)
)

func init() {
	SchemeBuilder.Register(&AcmpcaCertificate{}, &AcmpcaCertificateList{})
}
