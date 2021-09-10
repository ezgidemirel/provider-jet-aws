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

type BackupVaultPolicyObservation struct {
	BackupVaultArn string `json:"backupVaultArn" tf:"backup_vault_arn"`
}

type BackupVaultPolicyParameters struct {
	BackupVaultName string `json:"backupVaultName" tf:"backup_vault_name"`

	Policy string `json:"policy" tf:"policy"`

	Region string `json:"region" tf:"-"`
}

// BackupVaultPolicySpec defines the desired state of BackupVaultPolicy
type BackupVaultPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BackupVaultPolicyParameters `json:"forProvider"`
}

// BackupVaultPolicyStatus defines the observed state of BackupVaultPolicy.
type BackupVaultPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BackupVaultPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultPolicy is the Schema for the BackupVaultPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type BackupVaultPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupVaultPolicySpec   `json:"spec"`
	Status            BackupVaultPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultPolicyList contains a list of BackupVaultPolicys
type BackupVaultPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVaultPolicy `json:"items"`
}

// Repository type metadata.
var (
	BackupVaultPolicyKind             = "BackupVaultPolicy"
	BackupVaultPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: BackupVaultPolicyKind}.String()
	BackupVaultPolicyKindAPIVersion   = BackupVaultPolicyKind + "." + GroupVersion.String()
	BackupVaultPolicyGroupVersionKind = GroupVersion.WithKind(BackupVaultPolicyKind)
)

func init() {
	SchemeBuilder.Register(&BackupVaultPolicy{}, &BackupVaultPolicyList{})
}
