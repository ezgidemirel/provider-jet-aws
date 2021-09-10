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

type ApiGatewayMethodResponseObservation struct {
}

type ApiGatewayMethodResponseParameters struct {
	HttpMethod string `json:"httpMethod" tf:"http_method"`

	Region string `json:"region" tf:"-"`

	ResourceId string `json:"resourceId" tf:"resource_id"`

	ResponseModels map[string]string `json:"responseModels,omitempty" tf:"response_models"`

	ResponseParameters map[string]bool `json:"responseParameters,omitempty" tf:"response_parameters"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	StatusCode string `json:"statusCode" tf:"status_code"`
}

// ApiGatewayMethodResponseSpec defines the desired state of ApiGatewayMethodResponse
type ApiGatewayMethodResponseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayMethodResponseParameters `json:"forProvider"`
}

// ApiGatewayMethodResponseStatus defines the observed state of ApiGatewayMethodResponse.
type ApiGatewayMethodResponseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayMethodResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodResponse is the Schema for the ApiGatewayMethodResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayMethodResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodResponseSpec   `json:"spec"`
	Status            ApiGatewayMethodResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodResponseList contains a list of ApiGatewayMethodResponses
type ApiGatewayMethodResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayMethodResponse `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayMethodResponseKind             = "ApiGatewayMethodResponse"
	ApiGatewayMethodResponseGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayMethodResponseKind}.String()
	ApiGatewayMethodResponseKindAPIVersion   = ApiGatewayMethodResponseKind + "." + GroupVersion.String()
	ApiGatewayMethodResponseGroupVersionKind = GroupVersion.WithKind(ApiGatewayMethodResponseKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayMethodResponse{}, &ApiGatewayMethodResponseList{})
}
