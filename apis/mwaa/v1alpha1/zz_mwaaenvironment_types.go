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

type DagProcessingLogsObservation struct {
	CloudWatchLogGroupArn string `json:"cloudWatchLogGroupArn" tf:"cloud_watch_log_group_arn"`
}

type DagProcessingLogsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

type ErrorObservation struct {
	ErrorCode string `json:"errorCode" tf:"error_code"`

	ErrorMessage string `json:"errorMessage" tf:"error_message"`
}

type ErrorParameters struct {
}

type LastUpdatedObservation struct {
	CreatedAt string `json:"createdAt" tf:"created_at"`

	Error []ErrorObservation `json:"error" tf:"error"`

	Status string `json:"status" tf:"status"`
}

type LastUpdatedParameters struct {
}

type LoggingConfigurationObservation struct {
}

type LoggingConfigurationParameters struct {
	DagProcessingLogs []DagProcessingLogsParameters `json:"dagProcessingLogs,omitempty" tf:"dag_processing_logs"`

	SchedulerLogs []SchedulerLogsParameters `json:"schedulerLogs,omitempty" tf:"scheduler_logs"`

	TaskLogs []TaskLogsParameters `json:"taskLogs,omitempty" tf:"task_logs"`

	WebserverLogs []WebserverLogsParameters `json:"webserverLogs,omitempty" tf:"webserver_logs"`

	WorkerLogs []WorkerLogsParameters `json:"workerLogs,omitempty" tf:"worker_logs"`
}

type MwaaEnvironmentObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedAt string `json:"createdAt" tf:"created_at"`

	LastUpdated []LastUpdatedObservation `json:"lastUpdated" tf:"last_updated"`

	ServiceRoleArn string `json:"serviceRoleArn" tf:"service_role_arn"`

	Status string `json:"status" tf:"status"`

	WebserverUrl string `json:"webserverUrl" tf:"webserver_url"`
}

type MwaaEnvironmentParameters struct {
	AirflowConfigurationOptions map[string]string `json:"airflowConfigurationOptions,omitempty" tf:"airflow_configuration_options"`

	AirflowVersion *string `json:"airflowVersion,omitempty" tf:"airflow_version"`

	DagS3Path string `json:"dagS3Path" tf:"dag_s3_path"`

	EnvironmentClass *string `json:"environmentClass,omitempty" tf:"environment_class"`

	ExecutionRoleArn string `json:"executionRoleArn" tf:"execution_role_arn"`

	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key"`

	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration"`

	MaxWorkers *int64 `json:"maxWorkers,omitempty" tf:"max_workers"`

	MinWorkers *int64 `json:"minWorkers,omitempty" tf:"min_workers"`

	Name string `json:"name" tf:"name"`

	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration" tf:"network_configuration"`

	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion,omitempty" tf:"plugins_s3_object_version"`

	PluginsS3Path *string `json:"pluginsS3Path,omitempty" tf:"plugins_s3_path"`

	Region string `json:"region" tf:"-"`

	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion,omitempty" tf:"requirements_s3_object_version"`

	RequirementsS3Path *string `json:"requirementsS3Path,omitempty" tf:"requirements_s3_path"`

	SourceBucketArn string `json:"sourceBucketArn" tf:"source_bucket_arn"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	WebserverAccessMode *string `json:"webserverAccessMode,omitempty" tf:"webserver_access_mode"`

	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart,omitempty" tf:"weekly_maintenance_window_start"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`
}

type SchedulerLogsObservation struct {
	CloudWatchLogGroupArn string `json:"cloudWatchLogGroupArn" tf:"cloud_watch_log_group_arn"`
}

type SchedulerLogsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

type TaskLogsObservation struct {
	CloudWatchLogGroupArn string `json:"cloudWatchLogGroupArn" tf:"cloud_watch_log_group_arn"`
}

type TaskLogsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

type WebserverLogsObservation struct {
	CloudWatchLogGroupArn string `json:"cloudWatchLogGroupArn" tf:"cloud_watch_log_group_arn"`
}

type WebserverLogsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

type WorkerLogsObservation struct {
	CloudWatchLogGroupArn string `json:"cloudWatchLogGroupArn" tf:"cloud_watch_log_group_arn"`
}

type WorkerLogsParameters struct {
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

// MwaaEnvironmentSpec defines the desired state of MwaaEnvironment
type MwaaEnvironmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MwaaEnvironmentParameters `json:"forProvider"`
}

// MwaaEnvironmentStatus defines the observed state of MwaaEnvironment.
type MwaaEnvironmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MwaaEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MwaaEnvironment is the Schema for the MwaaEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type MwaaEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MwaaEnvironmentSpec   `json:"spec"`
	Status            MwaaEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MwaaEnvironmentList contains a list of MwaaEnvironments
type MwaaEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MwaaEnvironment `json:"items"`
}

// Repository type metadata.
var (
	MwaaEnvironmentKind             = "MwaaEnvironment"
	MwaaEnvironmentGroupKind        = schema.GroupKind{Group: Group, Kind: MwaaEnvironmentKind}.String()
	MwaaEnvironmentKindAPIVersion   = MwaaEnvironmentKind + "." + GroupVersion.String()
	MwaaEnvironmentGroupVersionKind = GroupVersion.WithKind(MwaaEnvironmentKind)
)

func init() {
	SchemeBuilder.Register(&MwaaEnvironment{}, &MwaaEnvironmentList{})
}
