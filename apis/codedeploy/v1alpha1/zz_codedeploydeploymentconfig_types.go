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

type CodedeployDeploymentConfigObservation struct {
	DeploymentConfigId string `json:"deploymentConfigId" tf:"deployment_config_id"`
}

type CodedeployDeploymentConfigParameters struct {
	ComputePlatform *string `json:"computePlatform,omitempty" tf:"compute_platform"`

	DeploymentConfigName string `json:"deploymentConfigName" tf:"deployment_config_name"`

	MinimumHealthyHosts []MinimumHealthyHostsParameters `json:"minimumHealthyHosts,omitempty" tf:"minimum_healthy_hosts"`

	Region string `json:"region" tf:"-"`

	TrafficRoutingConfig []TrafficRoutingConfigParameters `json:"trafficRoutingConfig,omitempty" tf:"traffic_routing_config"`
}

type MinimumHealthyHostsObservation struct {
}

type MinimumHealthyHostsParameters struct {
	Type *string `json:"type,omitempty" tf:"type"`

	Value *int64 `json:"value,omitempty" tf:"value"`
}

type TimeBasedCanaryObservation struct {
}

type TimeBasedCanaryParameters struct {
	Interval *int64 `json:"interval,omitempty" tf:"interval"`

	Percentage *int64 `json:"percentage,omitempty" tf:"percentage"`
}

type TimeBasedLinearObservation struct {
}

type TimeBasedLinearParameters struct {
	Interval *int64 `json:"interval,omitempty" tf:"interval"`

	Percentage *int64 `json:"percentage,omitempty" tf:"percentage"`
}

type TrafficRoutingConfigObservation struct {
}

type TrafficRoutingConfigParameters struct {
	TimeBasedCanary []TimeBasedCanaryParameters `json:"timeBasedCanary,omitempty" tf:"time_based_canary"`

	TimeBasedLinear []TimeBasedLinearParameters `json:"timeBasedLinear,omitempty" tf:"time_based_linear"`

	Type *string `json:"type,omitempty" tf:"type"`
}

// CodedeployDeploymentConfigSpec defines the desired state of CodedeployDeploymentConfig
type CodedeployDeploymentConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodedeployDeploymentConfigParameters `json:"forProvider"`
}

// CodedeployDeploymentConfigStatus defines the observed state of CodedeployDeploymentConfig.
type CodedeployDeploymentConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodedeployDeploymentConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentConfig is the Schema for the CodedeployDeploymentConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodedeployDeploymentConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployDeploymentConfigSpec   `json:"spec"`
	Status            CodedeployDeploymentConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentConfigList contains a list of CodedeployDeploymentConfigs
type CodedeployDeploymentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodedeployDeploymentConfig `json:"items"`
}

// Repository type metadata.
var (
	CodedeployDeploymentConfigKind             = "CodedeployDeploymentConfig"
	CodedeployDeploymentConfigGroupKind        = schema.GroupKind{Group: Group, Kind: CodedeployDeploymentConfigKind}.String()
	CodedeployDeploymentConfigKindAPIVersion   = CodedeployDeploymentConfigKind + "." + GroupVersion.String()
	CodedeployDeploymentConfigGroupVersionKind = GroupVersion.WithKind(CodedeployDeploymentConfigKind)
)

func init() {
	SchemeBuilder.Register(&CodedeployDeploymentConfig{}, &CodedeployDeploymentConfigList{})
}
