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

type Apigatewayv2RouteObservation struct {
}

type Apigatewayv2RouteParameters struct {
	ApiId string `json:"apiId" tf:"api_id"`

	ApiKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required"`

	AuthorizationScopes []string `json:"authorizationScopes,omitempty" tf:"authorization_scopes"`

	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type"`

	AuthorizerId *string `json:"authorizerId,omitempty" tf:"authorizer_id"`

	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression"`

	OperationName *string `json:"operationName,omitempty" tf:"operation_name"`

	Region string `json:"region" tf:"-"`

	RequestModels map[string]string `json:"requestModels,omitempty" tf:"request_models"`

	RequestParameter []RequestParameterParameters `json:"requestParameter,omitempty" tf:"request_parameter"`

	RouteKey string `json:"routeKey" tf:"route_key"`

	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression,omitempty" tf:"route_response_selection_expression"`

	Target *string `json:"target,omitempty" tf:"target"`
}

type RequestParameterObservation struct {
}

type RequestParameterParameters struct {
	RequestParameterKey string `json:"requestParameterKey" tf:"request_parameter_key"`

	Required bool `json:"required" tf:"required"`
}

// Apigatewayv2RouteSpec defines the desired state of Apigatewayv2Route
type Apigatewayv2RouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2RouteParameters `json:"forProvider"`
}

// Apigatewayv2RouteStatus defines the observed state of Apigatewayv2Route.
type Apigatewayv2RouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Route is the Schema for the Apigatewayv2Routes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Apigatewayv2Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2RouteSpec   `json:"spec"`
	Status            Apigatewayv2RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2RouteList contains a list of Apigatewayv2Routes
type Apigatewayv2RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Route `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2RouteKind             = "Apigatewayv2Route"
	Apigatewayv2RouteGroupKind        = schema.GroupKind{Group: Group, Kind: Apigatewayv2RouteKind}.String()
	Apigatewayv2RouteKindAPIVersion   = Apigatewayv2RouteKind + "." + GroupVersion.String()
	Apigatewayv2RouteGroupVersionKind = GroupVersion.WithKind(Apigatewayv2RouteKind)
)

func init() {
	SchemeBuilder.Register(&Apigatewayv2Route{}, &Apigatewayv2RouteList{})
}
