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

type AccessLogSettingsObservation struct {
}

type AccessLogSettingsParameters struct {
	DestinationArn string `json:"destinationArn" tf:"destination_arn"`

	Format string `json:"format" tf:"format"`
}

type ApiGatewayStageObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ExecutionArn string `json:"executionArn" tf:"execution_arn"`

	InvokeUrl string `json:"invokeUrl" tf:"invoke_url"`
}

type ApiGatewayStageParameters struct {
	AccessLogSettings []AccessLogSettingsParameters `json:"accessLogSettings,omitempty" tf:"access_log_settings"`

	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled"`

	CacheClusterSize *string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size"`

	ClientCertificateId *string `json:"clientCertificateId,omitempty" tf:"client_certificate_id"`

	DeploymentId string `json:"deploymentId" tf:"deployment_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	DocumentationVersion *string `json:"documentationVersion,omitempty" tf:"documentation_version"`

	Region string `json:"region" tf:"-"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	StageName string `json:"stageName" tf:"stage_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Variables map[string]string `json:"variables,omitempty" tf:"variables"`

	XrayTracingEnabled *bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled"`
}

// ApiGatewayStageSpec defines the desired state of ApiGatewayStage
type ApiGatewayStageSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayStageParameters `json:"forProvider"`
}

// ApiGatewayStageStatus defines the observed state of ApiGatewayStage.
type ApiGatewayStageStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayStageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayStage is the Schema for the ApiGatewayStages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayStage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayStageSpec   `json:"spec"`
	Status            ApiGatewayStageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayStageList contains a list of ApiGatewayStages
type ApiGatewayStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayStage `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayStageKind             = "ApiGatewayStage"
	ApiGatewayStageGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayStageKind}.String()
	ApiGatewayStageKindAPIVersion   = ApiGatewayStageKind + "." + GroupVersion.String()
	ApiGatewayStageGroupVersionKind = GroupVersion.WithKind(ApiGatewayStageKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayStage{}, &ApiGatewayStageList{})
}
