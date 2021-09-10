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

type ApprunnerServiceObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ServiceId string `json:"serviceId" tf:"service_id"`

	ServiceUrl string `json:"serviceUrl" tf:"service_url"`

	Status string `json:"status" tf:"status"`
}

type ApprunnerServiceParameters struct {
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn,omitempty" tf:"auto_scaling_configuration_arn"`

	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration"`

	HealthCheckConfiguration []HealthCheckConfigurationParameters `json:"healthCheckConfiguration,omitempty" tf:"health_check_configuration"`

	InstanceConfiguration []InstanceConfigurationParameters `json:"instanceConfiguration,omitempty" tf:"instance_configuration"`

	Region string `json:"region" tf:"-"`

	ServiceName string `json:"serviceName" tf:"service_name"`

	SourceConfiguration []SourceConfigurationParameters `json:"sourceConfiguration" tf:"source_configuration"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AuthenticationConfigurationObservation struct {
}

type AuthenticationConfigurationParameters struct {
	AccessRoleArn *string `json:"accessRoleArn,omitempty" tf:"access_role_arn"`

	ConnectionArn *string `json:"connectionArn,omitempty" tf:"connection_arn"`
}

type CodeConfigurationObservation struct {
}

type CodeConfigurationParameters struct {
	CodeConfigurationValues []CodeConfigurationValuesParameters `json:"codeConfigurationValues,omitempty" tf:"code_configuration_values"`

	ConfigurationSource string `json:"configurationSource" tf:"configuration_source"`
}

type CodeConfigurationValuesObservation struct {
}

type CodeConfigurationValuesParameters struct {
	BuildCommand *string `json:"buildCommand,omitempty" tf:"build_command"`

	Port *string `json:"port,omitempty" tf:"port"`

	Runtime string `json:"runtime" tf:"runtime"`

	RuntimeEnvironmentVariables map[string]string `json:"runtimeEnvironmentVariables,omitempty" tf:"runtime_environment_variables"`

	StartCommand *string `json:"startCommand,omitempty" tf:"start_command"`
}

type CodeRepositoryObservation struct {
}

type CodeRepositoryParameters struct {
	CodeConfiguration []CodeConfigurationParameters `json:"codeConfiguration,omitempty" tf:"code_configuration"`

	RepositoryUrl string `json:"repositoryUrl" tf:"repository_url"`

	SourceCodeVersion []SourceCodeVersionParameters `json:"sourceCodeVersion" tf:"source_code_version"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {
	KmsKey string `json:"kmsKey" tf:"kms_key"`
}

type HealthCheckConfigurationObservation struct {
}

type HealthCheckConfigurationParameters struct {
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold"`

	Interval *int64 `json:"interval,omitempty" tf:"interval"`

	Path *string `json:"path,omitempty" tf:"path"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol"`

	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`

	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold"`
}

type ImageConfigurationObservation struct {
}

type ImageConfigurationParameters struct {
	Port *string `json:"port,omitempty" tf:"port"`

	RuntimeEnvironmentVariables map[string]string `json:"runtimeEnvironmentVariables,omitempty" tf:"runtime_environment_variables"`

	StartCommand *string `json:"startCommand,omitempty" tf:"start_command"`
}

type ImageRepositoryObservation struct {
}

type ImageRepositoryParameters struct {
	ImageConfiguration []ImageConfigurationParameters `json:"imageConfiguration,omitempty" tf:"image_configuration"`

	ImageIdentifier string `json:"imageIdentifier" tf:"image_identifier"`

	ImageRepositoryType string `json:"imageRepositoryType" tf:"image_repository_type"`
}

type InstanceConfigurationObservation struct {
}

type InstanceConfigurationParameters struct {
	Cpu *string `json:"cpu,omitempty" tf:"cpu"`

	InstanceRoleArn string `json:"instanceRoleArn" tf:"instance_role_arn"`

	Memory *string `json:"memory,omitempty" tf:"memory"`
}

type SourceCodeVersionObservation struct {
}

type SourceCodeVersionParameters struct {
	Type string `json:"type" tf:"type"`

	Value string `json:"value" tf:"value"`
}

type SourceConfigurationObservation struct {
}

type SourceConfigurationParameters struct {
	AuthenticationConfiguration []AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration"`

	AutoDeploymentsEnabled *bool `json:"autoDeploymentsEnabled,omitempty" tf:"auto_deployments_enabled"`

	CodeRepository []CodeRepositoryParameters `json:"codeRepository,omitempty" tf:"code_repository"`

	ImageRepository []ImageRepositoryParameters `json:"imageRepository,omitempty" tf:"image_repository"`
}

// ApprunnerServiceSpec defines the desired state of ApprunnerService
type ApprunnerServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApprunnerServiceParameters `json:"forProvider"`
}

// ApprunnerServiceStatus defines the observed state of ApprunnerService.
type ApprunnerServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApprunnerServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApprunnerService is the Schema for the ApprunnerServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApprunnerService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApprunnerServiceSpec   `json:"spec"`
	Status            ApprunnerServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApprunnerServiceList contains a list of ApprunnerServices
type ApprunnerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApprunnerService `json:"items"`
}

// Repository type metadata.
var (
	ApprunnerServiceKind             = "ApprunnerService"
	ApprunnerServiceGroupKind        = schema.GroupKind{Group: Group, Kind: ApprunnerServiceKind}.String()
	ApprunnerServiceKindAPIVersion   = ApprunnerServiceKind + "." + GroupVersion.String()
	ApprunnerServiceGroupVersionKind = GroupVersion.WithKind(ApprunnerServiceKind)
)

func init() {
	SchemeBuilder.Register(&ApprunnerService{}, &ApprunnerServiceList{})
}
