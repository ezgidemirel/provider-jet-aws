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

type ServicecatalogProvisioningArtifactObservation struct {
	CreatedTime string `json:"createdTime" tf:"created_time"`
}

type ServicecatalogProvisioningArtifactParameters struct {
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language"`

	Active *bool `json:"active,omitempty" tf:"active"`

	Description *string `json:"description,omitempty" tf:"description"`

	DisableTemplateValidation *bool `json:"disableTemplateValidation,omitempty" tf:"disable_template_validation"`

	Guidance *string `json:"guidance,omitempty" tf:"guidance"`

	Name *string `json:"name,omitempty" tf:"name"`

	ProductId string `json:"productId" tf:"product_id"`

	Region string `json:"region" tf:"-"`

	TemplatePhysicalId *string `json:"templatePhysicalId,omitempty" tf:"template_physical_id"`

	TemplateUrl *string `json:"templateUrl,omitempty" tf:"template_url"`

	Type *string `json:"type,omitempty" tf:"type"`
}

// ServicecatalogProvisioningArtifactSpec defines the desired state of ServicecatalogProvisioningArtifact
type ServicecatalogProvisioningArtifactSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServicecatalogProvisioningArtifactParameters `json:"forProvider"`
}

// ServicecatalogProvisioningArtifactStatus defines the observed state of ServicecatalogProvisioningArtifact.
type ServicecatalogProvisioningArtifactStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServicecatalogProvisioningArtifactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogProvisioningArtifact is the Schema for the ServicecatalogProvisioningArtifacts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ServicecatalogProvisioningArtifact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogProvisioningArtifactSpec   `json:"spec"`
	Status            ServicecatalogProvisioningArtifactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicecatalogProvisioningArtifactList contains a list of ServicecatalogProvisioningArtifacts
type ServicecatalogProvisioningArtifactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicecatalogProvisioningArtifact `json:"items"`
}

// Repository type metadata.
var (
	ServicecatalogProvisioningArtifactKind             = "ServicecatalogProvisioningArtifact"
	ServicecatalogProvisioningArtifactGroupKind        = schema.GroupKind{Group: Group, Kind: ServicecatalogProvisioningArtifactKind}.String()
	ServicecatalogProvisioningArtifactKindAPIVersion   = ServicecatalogProvisioningArtifactKind + "." + GroupVersion.String()
	ServicecatalogProvisioningArtifactGroupVersionKind = GroupVersion.WithKind(ServicecatalogProvisioningArtifactKind)
)

func init() {
	SchemeBuilder.Register(&ServicecatalogProvisioningArtifact{}, &ServicecatalogProvisioningArtifactList{})
}
