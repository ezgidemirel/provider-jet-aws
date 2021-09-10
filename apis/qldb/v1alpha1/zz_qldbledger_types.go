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

type QldbLedgerObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type QldbLedgerParameters struct {
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`

	Name *string `json:"name,omitempty" tf:"name"`

	PermissionsMode string `json:"permissionsMode" tf:"permissions_mode"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// QldbLedgerSpec defines the desired state of QldbLedger
type QldbLedgerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       QldbLedgerParameters `json:"forProvider"`
}

// QldbLedgerStatus defines the observed state of QldbLedger.
type QldbLedgerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          QldbLedgerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QldbLedger is the Schema for the QldbLedgers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type QldbLedger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QldbLedgerSpec   `json:"spec"`
	Status            QldbLedgerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QldbLedgerList contains a list of QldbLedgers
type QldbLedgerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QldbLedger `json:"items"`
}

// Repository type metadata.
var (
	QldbLedgerKind             = "QldbLedger"
	QldbLedgerGroupKind        = schema.GroupKind{Group: Group, Kind: QldbLedgerKind}.String()
	QldbLedgerKindAPIVersion   = QldbLedgerKind + "." + GroupVersion.String()
	QldbLedgerGroupVersionKind = GroupVersion.WithKind(QldbLedgerKind)
)

func init() {
	SchemeBuilder.Register(&QldbLedger{}, &QldbLedgerList{})
}
