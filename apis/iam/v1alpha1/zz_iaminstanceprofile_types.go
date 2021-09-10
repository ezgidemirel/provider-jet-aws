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

type IamInstanceProfileObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreateDate string `json:"createDate" tf:"create_date"`

	UniqueId string `json:"uniqueId" tf:"unique_id"`
}

type IamInstanceProfileParameters struct {
	Name *string `json:"name,omitempty" tf:"name"`

	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	Path *string `json:"path,omitempty" tf:"path"`

	Region string `json:"region" tf:"-"`

	Role *string `json:"role,omitempty" tf:"role"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// IamInstanceProfileSpec defines the desired state of IamInstanceProfile
type IamInstanceProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamInstanceProfileParameters `json:"forProvider"`
}

// IamInstanceProfileStatus defines the observed state of IamInstanceProfile.
type IamInstanceProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamInstanceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamInstanceProfile is the Schema for the IamInstanceProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamInstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamInstanceProfileSpec   `json:"spec"`
	Status            IamInstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamInstanceProfileList contains a list of IamInstanceProfiles
type IamInstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamInstanceProfile `json:"items"`
}

// Repository type metadata.
var (
	IamInstanceProfileKind             = "IamInstanceProfile"
	IamInstanceProfileGroupKind        = schema.GroupKind{Group: Group, Kind: IamInstanceProfileKind}.String()
	IamInstanceProfileKindAPIVersion   = IamInstanceProfileKind + "." + GroupVersion.String()
	IamInstanceProfileGroupVersionKind = GroupVersion.WithKind(IamInstanceProfileKind)
)

func init() {
	SchemeBuilder.Register(&IamInstanceProfile{}, &IamInstanceProfileList{})
}
