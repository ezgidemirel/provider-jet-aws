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

type ApiGatewayRestApiEndpointConfigurationObservation struct {
}

type ApiGatewayRestApiEndpointConfigurationParameters struct {
	Types []string `json:"types" tf:"types"`

	VpcEndpointIds []string `json:"vpcEndpointIds,omitempty" tf:"vpc_endpoint_ids"`
}

type ApiGatewayRestApiObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	ExecutionArn string `json:"executionArn" tf:"execution_arn"`

	RootResourceId string `json:"rootResourceId" tf:"root_resource_id"`
}

type ApiGatewayRestApiParameters struct {
	ApiKeySource *string `json:"apiKeySource,omitempty" tf:"api_key_source"`

	BinaryMediaTypes []string `json:"binaryMediaTypes,omitempty" tf:"binary_media_types"`

	Body *string `json:"body,omitempty" tf:"body"`

	Description *string `json:"description,omitempty" tf:"description"`

	DisableExecuteApiEndpoint *bool `json:"disableExecuteApiEndpoint,omitempty" tf:"disable_execute_api_endpoint"`

	EndpointConfiguration []ApiGatewayRestApiEndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration"`

	MinimumCompressionSize *int64 `json:"minimumCompressionSize,omitempty" tf:"minimum_compression_size"`

	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	Policy *string `json:"policy,omitempty" tf:"policy"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ApiGatewayRestApiSpec defines the desired state of ApiGatewayRestApi
type ApiGatewayRestApiSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayRestApiParameters `json:"forProvider"`
}

// ApiGatewayRestApiStatus defines the observed state of ApiGatewayRestApi.
type ApiGatewayRestApiStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayRestApiObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayRestApi is the Schema for the ApiGatewayRestApis API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayRestApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayRestApiSpec   `json:"spec"`
	Status            ApiGatewayRestApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayRestApiList contains a list of ApiGatewayRestApis
type ApiGatewayRestApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayRestApi `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayRestApiKind             = "ApiGatewayRestApi"
	ApiGatewayRestApiGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayRestApiKind}.String()
	ApiGatewayRestApiKindAPIVersion   = ApiGatewayRestApiKind + "." + GroupVersion.String()
	ApiGatewayRestApiGroupVersionKind = GroupVersion.WithKind(ApiGatewayRestApiKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayRestApi{}, &ApiGatewayRestApiList{})
}
