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

type DbClusterSnapshotObservation struct {
	AllocatedStorage int64 `json:"allocatedStorage" tf:"allocated_storage"`

	AvailabilityZones []string `json:"availabilityZones" tf:"availability_zones"`

	DbClusterSnapshotArn string `json:"dbClusterSnapshotArn" tf:"db_cluster_snapshot_arn"`

	Engine string `json:"engine" tf:"engine"`

	EngineVersion string `json:"engineVersion" tf:"engine_version"`

	KmsKeyId string `json:"kmsKeyId" tf:"kms_key_id"`

	LicenseModel string `json:"licenseModel" tf:"license_model"`

	Port int64 `json:"port" tf:"port"`

	SnapshotType string `json:"snapshotType" tf:"snapshot_type"`

	SourceDbClusterSnapshotArn string `json:"sourceDbClusterSnapshotArn" tf:"source_db_cluster_snapshot_arn"`

	Status string `json:"status" tf:"status"`

	StorageEncrypted bool `json:"storageEncrypted" tf:"storage_encrypted"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type DbClusterSnapshotParameters struct {
	DbClusterIdentifier string `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`

	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`

	Region string `json:"region" tf:"-"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DbClusterSnapshotSpec defines the desired state of DbClusterSnapshot
type DbClusterSnapshotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DbClusterSnapshotParameters `json:"forProvider"`
}

// DbClusterSnapshotStatus defines the observed state of DbClusterSnapshot.
type DbClusterSnapshotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DbClusterSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DbClusterSnapshot is the Schema for the DbClusterSnapshots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbClusterSnapshotSpec   `json:"spec"`
	Status            DbClusterSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbClusterSnapshotList contains a list of DbClusterSnapshots
type DbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbClusterSnapshot `json:"items"`
}

// Repository type metadata.
var (
	DbClusterSnapshotKind             = "DbClusterSnapshot"
	DbClusterSnapshotGroupKind        = schema.GroupKind{Group: Group, Kind: DbClusterSnapshotKind}.String()
	DbClusterSnapshotKindAPIVersion   = DbClusterSnapshotKind + "." + GroupVersion.String()
	DbClusterSnapshotGroupVersionKind = GroupVersion.WithKind(DbClusterSnapshotKind)
)

func init() {
	SchemeBuilder.Register(&DbClusterSnapshot{}, &DbClusterSnapshotList{})
}
