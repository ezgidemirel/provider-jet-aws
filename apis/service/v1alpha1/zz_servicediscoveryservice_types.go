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

type DnsConfigObservation struct {
}

type DnsConfigParameters struct {
	DnsRecords []DnsRecordsParameters `json:"dnsRecords" tf:"dns_records"`

	NamespaceId string `json:"namespaceId" tf:"namespace_id"`

	RoutingPolicy *string `json:"routingPolicy,omitempty" tf:"routing_policy"`
}

type DnsRecordsObservation struct {
}

type DnsRecordsParameters struct {
	Ttl int64 `json:"ttl" tf:"ttl"`

	Type string `json:"type" tf:"type"`
}

type HealthCheckConfigObservation struct {
}

type HealthCheckConfigParameters struct {
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`

	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path"`

	Type *string `json:"type,omitempty" tf:"type"`
}

type HealthCheckCustomConfigObservation struct {
}

type HealthCheckCustomConfigParameters struct {
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
}

type ServiceDiscoveryServiceObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type ServiceDiscoveryServiceParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DnsConfig []DnsConfigParameters `json:"dnsConfig,omitempty" tf:"dns_config"`

	HealthCheckConfig []HealthCheckConfigParameters `json:"healthCheckConfig,omitempty" tf:"health_check_config"`

	HealthCheckCustomConfig []HealthCheckCustomConfigParameters `json:"healthCheckCustomConfig,omitempty" tf:"health_check_custom_config"`

	Name string `json:"name" tf:"name"`

	NamespaceId *string `json:"namespaceId,omitempty" tf:"namespace_id"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ServiceDiscoveryServiceSpec defines the desired state of ServiceDiscoveryService
type ServiceDiscoveryServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceDiscoveryServiceParameters `json:"forProvider"`
}

// ServiceDiscoveryServiceStatus defines the observed state of ServiceDiscoveryService.
type ServiceDiscoveryServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceDiscoveryServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryService is the Schema for the ServiceDiscoveryServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ServiceDiscoveryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryServiceSpec   `json:"spec"`
	Status            ServiceDiscoveryServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceDiscoveryServiceList contains a list of ServiceDiscoveryServices
type ServiceDiscoveryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDiscoveryService `json:"items"`
}

// Repository type metadata.
var (
	ServiceDiscoveryServiceKind             = "ServiceDiscoveryService"
	ServiceDiscoveryServiceGroupKind        = schema.GroupKind{Group: Group, Kind: ServiceDiscoveryServiceKind}.String()
	ServiceDiscoveryServiceKindAPIVersion   = ServiceDiscoveryServiceKind + "." + GroupVersion.String()
	ServiceDiscoveryServiceGroupVersionKind = GroupVersion.WithKind(ServiceDiscoveryServiceKind)
)

func init() {
	SchemeBuilder.Register(&ServiceDiscoveryService{}, &ServiceDiscoveryServiceList{})
}
