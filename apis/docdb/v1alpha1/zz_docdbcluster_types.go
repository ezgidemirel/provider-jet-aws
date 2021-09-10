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

type DocdbClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ClusterResourceId string `json:"clusterResourceId" tf:"cluster_resource_id"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	HostedZoneId string `json:"hostedZoneId" tf:"hosted_zone_id"`

	ReaderEndpoint string `json:"readerEndpoint" tf:"reader_endpoint"`
}

type DocdbClusterParameters struct {
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`

	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty" tf:"backup_retention_period"`

	ClusterIdentifier *string `json:"clusterIdentifier,omitempty" tf:"cluster_identifier"`

	ClusterIdentifierPrefix *string `json:"clusterIdentifierPrefix,omitempty" tf:"cluster_identifier_prefix"`

	ClusterMembers []string `json:"clusterMembers,omitempty" tf:"cluster_members"`

	DbClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty" tf:"db_cluster_parameter_group_name"`

	DbSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name"`

	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`

	EnabledCloudwatchLogsExports []string `json:"enabledCloudwatchLogsExports,omitempty" tf:"enabled_cloudwatch_logs_exports"`

	Engine *string `json:"engine,omitempty" tf:"engine"`

	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`

	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	MasterPassword *string `json:"masterPassword,omitempty" tf:"master_password"`

	MasterUsername *string `json:"masterUsername,omitempty" tf:"master_username"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`

	Region string `json:"region" tf:"-"`

	SkipFinalSnapshot *bool `json:"skipFinalSnapshot,omitempty" tf:"skip_final_snapshot"`

	SnapshotIdentifier *string `json:"snapshotIdentifier,omitempty" tf:"snapshot_identifier"`

	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`
}

// DocdbClusterSpec defines the desired state of DocdbCluster
type DocdbClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DocdbClusterParameters `json:"forProvider"`
}

// DocdbClusterStatus defines the observed state of DocdbCluster.
type DocdbClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DocdbClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbCluster is the Schema for the DocdbClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DocdbCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterSpec   `json:"spec"`
	Status            DocdbClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocdbClusterList contains a list of DocdbClusters
type DocdbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocdbCluster `json:"items"`
}

// Repository type metadata.
var (
	DocdbClusterKind             = "DocdbCluster"
	DocdbClusterGroupKind        = schema.GroupKind{Group: Group, Kind: DocdbClusterKind}.String()
	DocdbClusterKindAPIVersion   = DocdbClusterKind + "." + GroupVersion.String()
	DocdbClusterGroupVersionKind = GroupVersion.WithKind(DocdbClusterKind)
)

func init() {
	SchemeBuilder.Register(&DocdbCluster{}, &DocdbClusterList{})
}
