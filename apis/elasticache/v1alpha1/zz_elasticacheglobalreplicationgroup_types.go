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

type ElasticacheGlobalReplicationGroupObservation struct {
	ActualEngineVersion string `json:"actualEngineVersion" tf:"actual_engine_version"`

	Arn string `json:"arn" tf:"arn"`

	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled" tf:"at_rest_encryption_enabled"`

	AuthTokenEnabled bool `json:"authTokenEnabled" tf:"auth_token_enabled"`

	CacheNodeType string `json:"cacheNodeType" tf:"cache_node_type"`

	ClusterEnabled bool `json:"clusterEnabled" tf:"cluster_enabled"`

	Engine string `json:"engine" tf:"engine"`

	EngineVersionActual string `json:"engineVersionActual" tf:"engine_version_actual"`

	GlobalReplicationGroupId string `json:"globalReplicationGroupId" tf:"global_replication_group_id"`

	TransitEncryptionEnabled bool `json:"transitEncryptionEnabled" tf:"transit_encryption_enabled"`
}

type ElasticacheGlobalReplicationGroupParameters struct {
	GlobalReplicationGroupDescription *string `json:"globalReplicationGroupDescription,omitempty" tf:"global_replication_group_description"`

	GlobalReplicationGroupIdSuffix string `json:"globalReplicationGroupIdSuffix" tf:"global_replication_group_id_suffix"`

	PrimaryReplicationGroupId string `json:"primaryReplicationGroupId" tf:"primary_replication_group_id"`

	Region string `json:"region" tf:"-"`
}

// ElasticacheGlobalReplicationGroupSpec defines the desired state of ElasticacheGlobalReplicationGroup
type ElasticacheGlobalReplicationGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticacheGlobalReplicationGroupParameters `json:"forProvider"`
}

// ElasticacheGlobalReplicationGroupStatus defines the observed state of ElasticacheGlobalReplicationGroup.
type ElasticacheGlobalReplicationGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticacheGlobalReplicationGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheGlobalReplicationGroup is the Schema for the ElasticacheGlobalReplicationGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticacheGlobalReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheGlobalReplicationGroupSpec   `json:"spec"`
	Status            ElasticacheGlobalReplicationGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheGlobalReplicationGroupList contains a list of ElasticacheGlobalReplicationGroups
type ElasticacheGlobalReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheGlobalReplicationGroup `json:"items"`
}

// Repository type metadata.
var (
	ElasticacheGlobalReplicationGroupKind             = "ElasticacheGlobalReplicationGroup"
	ElasticacheGlobalReplicationGroupGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticacheGlobalReplicationGroupKind}.String()
	ElasticacheGlobalReplicationGroupKindAPIVersion   = ElasticacheGlobalReplicationGroupKind + "." + GroupVersion.String()
	ElasticacheGlobalReplicationGroupGroupVersionKind = GroupVersion.WithKind(ElasticacheGlobalReplicationGroupKind)
)

func init() {
	SchemeBuilder.Register(&ElasticacheGlobalReplicationGroup{}, &ElasticacheGlobalReplicationGroupList{})
}
