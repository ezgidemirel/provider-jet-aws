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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IPAddressObservation struct {
	IPID string `json:"ipId,omitempty" tf:"ip_id"`
}

type IPAddressParameters struct {

	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

type Route53ResolverEndpointObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	HostVpcID string `json:"hostVpcId,omitempty" tf:"host_vpc_id"`
}

type Route53ResolverEndpointParameters struct {

	// +kubebuilder:validation:Required
	Direction string `json:"direction" tf:"direction"`

	// +kubebuilder:validation:Required
	IPAddress []IPAddressParameters `json:"ipAddress" tf:"ip_address"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Route53ResolverEndpointSpec defines the desired state of Route53ResolverEndpoint
type Route53ResolverEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Route53ResolverEndpointParameters `json:"forProvider"`
}

// Route53ResolverEndpointStatus defines the observed state of Route53ResolverEndpoint.
type Route53ResolverEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Route53ResolverEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverEndpoint is the Schema for the Route53ResolverEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverEndpointSpec   `json:"spec"`
	Status            Route53ResolverEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverEndpointList contains a list of Route53ResolverEndpoints
type Route53ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverEndpoint `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverEndpointKind             = "Route53ResolverEndpoint"
	Route53ResolverEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: Route53ResolverEndpointKind}.String()
	Route53ResolverEndpointKindAPIVersion   = Route53ResolverEndpointKind + "." + GroupVersion.String()
	Route53ResolverEndpointGroupVersionKind = GroupVersion.WithKind(Route53ResolverEndpointKind)
)

func init() {
	SchemeBuilder.Register(&Route53ResolverEndpoint{}, &Route53ResolverEndpointList{})
}
