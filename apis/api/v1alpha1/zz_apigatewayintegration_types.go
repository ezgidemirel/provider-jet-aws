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

type ApiGatewayIntegrationObservation struct {
}

type ApiGatewayIntegrationParameters struct {
	CacheKeyParameters []string `json:"cacheKeyParameters,omitempty" tf:"cache_key_parameters"`

	CacheNamespace *string `json:"cacheNamespace,omitempty" tf:"cache_namespace"`

	ConnectionId *string `json:"connectionId,omitempty" tf:"connection_id"`

	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type"`

	ContentHandling *string `json:"contentHandling,omitempty" tf:"content_handling"`

	Credentials *string `json:"credentials,omitempty" tf:"credentials"`

	HttpMethod string `json:"httpMethod" tf:"http_method"`

	IntegrationHttpMethod *string `json:"integrationHttpMethod,omitempty" tf:"integration_http_method"`

	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior"`

	Region string `json:"region" tf:"-"`

	RequestParameters map[string]string `json:"requestParameters,omitempty" tf:"request_parameters"`

	RequestTemplates map[string]string `json:"requestTemplates,omitempty" tf:"request_templates"`

	ResourceId string `json:"resourceId" tf:"resource_id"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	TimeoutMilliseconds *int64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds"`

	TlsConfig []TlsConfigParameters `json:"tlsConfig,omitempty" tf:"tls_config"`

	Type string `json:"type" tf:"type"`

	Uri *string `json:"uri,omitempty" tf:"uri"`
}

type TlsConfigObservation struct {
}

type TlsConfigParameters struct {
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty" tf:"insecure_skip_verification"`
}

// ApiGatewayIntegrationSpec defines the desired state of ApiGatewayIntegration
type ApiGatewayIntegrationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayIntegrationParameters `json:"forProvider"`
}

// ApiGatewayIntegrationStatus defines the observed state of ApiGatewayIntegration.
type ApiGatewayIntegrationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayIntegration is the Schema for the ApiGatewayIntegrations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayIntegrationSpec   `json:"spec"`
	Status            ApiGatewayIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayIntegrationList contains a list of ApiGatewayIntegrations
type ApiGatewayIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayIntegration `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayIntegrationKind             = "ApiGatewayIntegration"
	ApiGatewayIntegrationGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayIntegrationKind}.String()
	ApiGatewayIntegrationKindAPIVersion   = ApiGatewayIntegrationKind + "." + GroupVersion.String()
	ApiGatewayIntegrationGroupVersionKind = GroupVersion.WithKind(ApiGatewayIntegrationKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayIntegration{}, &ApiGatewayIntegrationList{})
}
