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

type CriterionObservation struct {
}

type CriterionParameters struct {
	Eq []string `json:"eq,omitempty" tf:"eq"`

	EqExactMatch []string `json:"eqExactMatch,omitempty" tf:"eq_exact_match"`

	Field string `json:"field" tf:"field"`

	Gt *string `json:"gt,omitempty" tf:"gt"`

	Gte *string `json:"gte,omitempty" tf:"gte"`

	Lt *string `json:"lt,omitempty" tf:"lt"`

	Lte *string `json:"lte,omitempty" tf:"lte"`

	Neq []string `json:"neq,omitempty" tf:"neq"`
}

type FindingCriteriaObservation struct {
}

type FindingCriteriaParameters struct {
	Criterion []CriterionParameters `json:"criterion,omitempty" tf:"criterion"`
}

type Macie2FindingsFilterObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type Macie2FindingsFilterParameters struct {
	Action string `json:"action" tf:"action"`

	Description *string `json:"description,omitempty" tf:"description"`

	FindingCriteria []FindingCriteriaParameters `json:"findingCriteria" tf:"finding_criteria"`

	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Position *int64 `json:"position,omitempty" tf:"position"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Macie2FindingsFilterSpec defines the desired state of Macie2FindingsFilter
type Macie2FindingsFilterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Macie2FindingsFilterParameters `json:"forProvider"`
}

// Macie2FindingsFilterStatus defines the observed state of Macie2FindingsFilter.
type Macie2FindingsFilterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Macie2FindingsFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2FindingsFilter is the Schema for the Macie2FindingsFilters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Macie2FindingsFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Macie2FindingsFilterSpec   `json:"spec"`
	Status            Macie2FindingsFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Macie2FindingsFilterList contains a list of Macie2FindingsFilters
type Macie2FindingsFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Macie2FindingsFilter `json:"items"`
}

// Repository type metadata.
var (
	Macie2FindingsFilterKind             = "Macie2FindingsFilter"
	Macie2FindingsFilterGroupKind        = schema.GroupKind{Group: Group, Kind: Macie2FindingsFilterKind}.String()
	Macie2FindingsFilterKindAPIVersion   = Macie2FindingsFilterKind + "." + GroupVersion.String()
	Macie2FindingsFilterGroupVersionKind = GroupVersion.WithKind(Macie2FindingsFilterKind)
)

func init() {
	SchemeBuilder.Register(&Macie2FindingsFilter{}, &Macie2FindingsFilterList{})
}
