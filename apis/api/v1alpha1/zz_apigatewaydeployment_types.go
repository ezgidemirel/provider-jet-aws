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

type ApiGatewayDeploymentObservation struct {
	CreatedDate string `json:"createdDate" tf:"created_date"`

	ExecutionArn string `json:"executionArn" tf:"execution_arn"`

	InvokeUrl string `json:"invokeUrl" tf:"invoke_url"`
}

type ApiGatewayDeploymentParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Region string `json:"region" tf:"-"`

	RestApiId string `json:"restApiId" tf:"rest_api_id"`

	StageDescription *string `json:"stageDescription,omitempty" tf:"stage_description"`

	StageName *string `json:"stageName,omitempty" tf:"stage_name"`

	Triggers map[string]string `json:"triggers,omitempty" tf:"triggers"`

	Variables map[string]string `json:"variables,omitempty" tf:"variables"`
}

// ApiGatewayDeploymentSpec defines the desired state of ApiGatewayDeployment
type ApiGatewayDeploymentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayDeploymentParameters `json:"forProvider"`
}

// ApiGatewayDeploymentStatus defines the observed state of ApiGatewayDeployment.
type ApiGatewayDeploymentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayDeployment is the Schema for the ApiGatewayDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDeploymentSpec   `json:"spec"`
	Status            ApiGatewayDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayDeploymentList contains a list of ApiGatewayDeployments
type ApiGatewayDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayDeployment `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayDeploymentKind             = "ApiGatewayDeployment"
	ApiGatewayDeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayDeploymentKind}.String()
	ApiGatewayDeploymentKindAPIVersion   = ApiGatewayDeploymentKind + "." + GroupVersion.String()
	ApiGatewayDeploymentGroupVersionKind = GroupVersion.WithKind(ApiGatewayDeploymentKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayDeployment{}, &ApiGatewayDeploymentList{})
}
