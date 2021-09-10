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

type ServiceDiscoveryPublicDnsNamespaceObservation struct {
	Arn string `json:"arn" tf:"arn"`

	HostedZone string `json:"hostedZone" tf:"hosted_zone"`
}

type ServiceDiscoveryPublicDnsNamespaceParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ServiceDiscoveryPublicDnsNamespaceSpec defines the desired state of ServiceDiscoveryPublicDnsNamespace
type ServiceDiscoveryPublicDnsNamespaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceDiscoveryPublicDnsNamespaceParameters `json:"forProvider"`
}

// ServiceDiscoveryPublicDnsNamespaceStatus defines the observed state of ServiceDiscoveryPublicDnsNamespace.
type ServiceDiscoveryPublicDnsNamespaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceDiscoveryPublicDnsNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryPublicDnsNamespace is the Schema for the ServiceDiscoveryPublicDnsNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ServiceDiscoveryPublicDnsNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryPublicDnsNamespaceSpec   `json:"spec"`
	Status            ServiceDiscoveryPublicDnsNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryPublicDnsNamespaceList contains a list of ServiceDiscoveryPublicDnsNamespaces
type ServiceDiscoveryPublicDnsNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDiscoveryPublicDnsNamespace `json:"items"`
}

// Repository type metadata.
var (
	ServiceDiscoveryPublicDnsNamespaceKind             = "ServiceDiscoveryPublicDnsNamespace"
	ServiceDiscoveryPublicDnsNamespaceGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceDiscoveryPublicDnsNamespaceKind}.String()
	ServiceDiscoveryPublicDnsNamespaceKindAPIVersion   = ServiceDiscoveryPublicDnsNamespaceKind + "." + GroupVersion.String()
	ServiceDiscoveryPublicDnsNamespaceGroupVersionKind = GroupVersion.WithKind(ServiceDiscoveryPublicDnsNamespaceKind)
)

func init() {
	SchemeBuilder.Register(&ServiceDiscoveryPublicDnsNamespace{}, &ServiceDiscoveryPublicDnsNamespaceList{})
}
