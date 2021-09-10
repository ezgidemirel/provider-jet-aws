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

type IngressObservation struct {
}

type IngressParameters struct {
	Cidr *string `json:"cidr,omitempty" tf:"cidr"`

	SecurityGroupName *string `json:"securityGroupName,omitempty" tf:"security_group_name"`

	SecurityGroupOwnerId *string `json:"securityGroupOwnerId,omitempty" tf:"security_group_owner_id"`
}

type RedshiftSecurityGroupObservation struct {
}

type RedshiftSecurityGroupParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Ingress []IngressParameters `json:"ingress" tf:"ingress"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`
}

// RedshiftSecurityGroupSpec defines the desired state of RedshiftSecurityGroup
type RedshiftSecurityGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RedshiftSecurityGroupParameters `json:"forProvider"`
}

// RedshiftSecurityGroupStatus defines the observed state of RedshiftSecurityGroup.
type RedshiftSecurityGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RedshiftSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftSecurityGroup is the Schema for the RedshiftSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RedshiftSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftSecurityGroupSpec   `json:"spec"`
	Status            RedshiftSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftSecurityGroupList contains a list of RedshiftSecurityGroups
type RedshiftSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	RedshiftSecurityGroupKind             = "RedshiftSecurityGroup"
	RedshiftSecurityGroupGroupKind        = schema.GroupKind{Group: Group, Kind: RedshiftSecurityGroupKind}.String()
	RedshiftSecurityGroupKindAPIVersion   = RedshiftSecurityGroupKind + "." + GroupVersion.String()
	RedshiftSecurityGroupGroupVersionKind = GroupVersion.WithKind(RedshiftSecurityGroupKind)
)

func init() {
	SchemeBuilder.Register(&RedshiftSecurityGroup{}, &RedshiftSecurityGroupList{})
}
