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

type GlobalacceleratorListenerObservation struct {
}

type GlobalacceleratorListenerParameters struct {
	AcceleratorArn string `json:"acceleratorArn" tf:"accelerator_arn"`

	ClientAffinity *string `json:"clientAffinity,omitempty" tf:"client_affinity"`

	PortRange []PortRangeParameters `json:"portRange" tf:"port_range"`

	Protocol string `json:"protocol" tf:"protocol"`

	Region string `json:"region" tf:"-"`
}

type PortRangeObservation struct {
}

type PortRangeParameters struct {
	FromPort *int64 `json:"fromPort,omitempty" tf:"from_port"`

	ToPort *int64 `json:"toPort,omitempty" tf:"to_port"`
}

// GlobalacceleratorListenerSpec defines the desired state of GlobalacceleratorListener
type GlobalacceleratorListenerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlobalacceleratorListenerParameters `json:"forProvider"`
}

// GlobalacceleratorListenerStatus defines the observed state of GlobalacceleratorListener.
type GlobalacceleratorListenerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlobalacceleratorListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalacceleratorListener is the Schema for the GlobalacceleratorListeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlobalacceleratorListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorListenerSpec   `json:"spec"`
	Status            GlobalacceleratorListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalacceleratorListenerList contains a list of GlobalacceleratorListeners
type GlobalacceleratorListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalacceleratorListener `json:"items"`
}

// Repository type metadata.
var (
	GlobalacceleratorListenerKind             = "GlobalacceleratorListener"
	GlobalacceleratorListenerGroupKind        = schema.GroupKind{Group: Group, Kind: GlobalacceleratorListenerKind}.String()
	GlobalacceleratorListenerKindAPIVersion   = GlobalacceleratorListenerKind + "." + GroupVersion.String()
	GlobalacceleratorListenerGroupVersionKind = GroupVersion.WithKind(GlobalacceleratorListenerKind)
)

func init() {
	SchemeBuilder.Register(&GlobalacceleratorListener{}, &GlobalacceleratorListenerList{})
}
