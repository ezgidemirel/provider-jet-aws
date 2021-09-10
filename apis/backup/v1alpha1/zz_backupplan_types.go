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

type AdvancedBackupSettingObservation struct {
}

type AdvancedBackupSettingParameters struct {
	BackupOptions map[string]string `json:"backupOptions" tf:"backup_options"`

	ResourceType string `json:"resourceType" tf:"resource_type"`
}

type BackupPlanObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Version string `json:"version" tf:"version"`
}

type BackupPlanParameters struct {
	AdvancedBackupSetting []AdvancedBackupSettingParameters `json:"advancedBackupSetting,omitempty" tf:"advanced_backup_setting"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	Rule []RuleParameters `json:"rule" tf:"rule"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type CopyActionObservation struct {
}

type CopyActionParameters struct {
	DestinationVaultArn string `json:"destinationVaultArn" tf:"destination_vault_arn"`

	Lifecycle []LifecycleParameters `json:"lifecycle,omitempty" tf:"lifecycle"`
}

type LifecycleObservation struct {
}

type LifecycleParameters struct {
	ColdStorageAfter *int64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after"`

	DeleteAfter *int64 `json:"deleteAfter,omitempty" tf:"delete_after"`
}

type RuleLifecycleObservation struct {
}

type RuleLifecycleParameters struct {
	ColdStorageAfter *int64 `json:"coldStorageAfter,omitempty" tf:"cold_storage_after"`

	DeleteAfter *int64 `json:"deleteAfter,omitempty" tf:"delete_after"`
}

type RuleObservation struct {
}

type RuleParameters struct {
	CompletionWindow *int64 `json:"completionWindow,omitempty" tf:"completion_window"`

	CopyAction []CopyActionParameters `json:"copyAction,omitempty" tf:"copy_action"`

	EnableContinuousBackup *bool `json:"enableContinuousBackup,omitempty" tf:"enable_continuous_backup"`

	Lifecycle []RuleLifecycleParameters `json:"lifecycle,omitempty" tf:"lifecycle"`

	RecoveryPointTags map[string]string `json:"recoveryPointTags,omitempty" tf:"recovery_point_tags"`

	RuleName string `json:"ruleName" tf:"rule_name"`

	Schedule *string `json:"schedule,omitempty" tf:"schedule"`

	StartWindow *int64 `json:"startWindow,omitempty" tf:"start_window"`

	TargetVaultName string `json:"targetVaultName" tf:"target_vault_name"`
}

// BackupPlanSpec defines the desired state of BackupPlan
type BackupPlanSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BackupPlanParameters `json:"forProvider"`
}

// BackupPlanStatus defines the observed state of BackupPlan.
type BackupPlanStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BackupPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPlan is the Schema for the BackupPlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type BackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPlanSpec   `json:"spec"`
	Status            BackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPlanList contains a list of BackupPlans
type BackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPlan `json:"items"`
}

// Repository type metadata.
var (
	BackupPlanKind             = "BackupPlan"
	BackupPlanGroupKind        = schema.GroupKind{Group: Group, Kind: BackupPlanKind}.String()
	BackupPlanKindAPIVersion   = BackupPlanKind + "." + GroupVersion.String()
	BackupPlanGroupVersionKind = GroupVersion.WithKind(BackupPlanKind)
)

func init() {
	SchemeBuilder.Register(&BackupPlan{}, &BackupPlanList{})
}
