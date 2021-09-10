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

type StoragegatewayWorkingStorageObservation struct {
}

type StoragegatewayWorkingStorageParameters struct {
	DiskId string `json:"diskId" tf:"disk_id"`

	GatewayArn string `json:"gatewayArn" tf:"gateway_arn"`

	Region string `json:"region" tf:"-"`
}

// StoragegatewayWorkingStorageSpec defines the desired state of StoragegatewayWorkingStorage
type StoragegatewayWorkingStorageSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StoragegatewayWorkingStorageParameters `json:"forProvider"`
}

// StoragegatewayWorkingStorageStatus defines the observed state of StoragegatewayWorkingStorage.
type StoragegatewayWorkingStorageStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StoragegatewayWorkingStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayWorkingStorage is the Schema for the StoragegatewayWorkingStorages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type StoragegatewayWorkingStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayWorkingStorageSpec   `json:"spec"`
	Status            StoragegatewayWorkingStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayWorkingStorageList contains a list of StoragegatewayWorkingStorages
type StoragegatewayWorkingStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayWorkingStorage `json:"items"`
}

// Repository type metadata.
var (
	StoragegatewayWorkingStorageKind             = "StoragegatewayWorkingStorage"
	StoragegatewayWorkingStorageGroupKind        = schema.GroupKind{Group: Group, Kind: StoragegatewayWorkingStorageKind}.String()
	StoragegatewayWorkingStorageKindAPIVersion   = StoragegatewayWorkingStorageKind + "." + GroupVersion.String()
	StoragegatewayWorkingStorageGroupVersionKind = GroupVersion.WithKind(StoragegatewayWorkingStorageKind)
)

func init() {
	SchemeBuilder.Register(&StoragegatewayWorkingStorage{}, &StoragegatewayWorkingStorageList{})
}
