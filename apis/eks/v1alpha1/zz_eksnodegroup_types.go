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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoscalingGroupsObservation struct {
	Name string `json:"name,omitempty" tf:"name"`
}

type AutoscalingGroupsParameters struct {
}

type EksNodeGroupObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources"`

	Status string `json:"status,omitempty" tf:"status"`
}

type EksNodeGroupParameters struct {

	// +kubebuilder:validation:Optional
	AmiType *string `json:"amiType,omitempty" tf:"ami_type"`

	// +kubebuilder:validation:Optional
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type"`

	// +kubebuilder:validation:Required
	ClusterName string `json:"clusterName" tf:"cluster_name"`

	// +kubebuilder:validation:Optional
	DiskSize *int64 `json:"diskSize,omitempty" tf:"disk_size"`

	// +kubebuilder:validation:Optional
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version"`

	// +kubebuilder:validation:Optional
	InstanceTypes []string `json:"instanceTypes,omitempty" tf:"instance_types"`

	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template"`

	// +kubebuilder:validation:Optional
	NodeGroupName *string `json:"nodeGroupName,omitempty" tf:"node_group_name"`

	// +kubebuilder:validation:Optional
	NodeGroupNamePrefix *string `json:"nodeGroupNamePrefix,omitempty" tf:"node_group_name_prefix"`

	// +kubebuilder:validation:Required
	NodeRoleArn string `json:"nodeRoleArn" tf:"node_role_arn"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version"`

	// +kubebuilder:validation:Optional
	RemoteAccess []RemoteAccessParameters `json:"remoteAccess,omitempty" tf:"remote_access"`

	// +kubebuilder:validation:Required
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig" tf:"scaling_config"`

	// +kubebuilder:validation:Required
	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	Taint []TaintParameters `json:"taint,omitempty" tf:"taint"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Required
	Version string `json:"version" tf:"version"`
}

type RemoteAccessObservation struct {
}

type RemoteAccessParameters struct {

	// +kubebuilder:validation:Optional
	Ec2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIds []string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids"`
}

type ResourcesObservation struct {
	AutoscalingGroups []AutoscalingGroupsObservation `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups"`

	RemoteAccessSecurityGroupID string `json:"remoteAccessSecurityGroupId,omitempty" tf:"remote_access_security_group_id"`
}

type ResourcesParameters struct {
}

type ScalingConfigObservation struct {
}

type ScalingConfigParameters struct {

	// +kubebuilder:validation:Required
	DesiredSize int64 `json:"desiredSize" tf:"desired_size"`

	// +kubebuilder:validation:Required
	MaxSize int64 `json:"maxSize" tf:"max_size"`

	// +kubebuilder:validation:Required
	MinSize int64 `json:"minSize" tf:"min_size"`
}

type TaintObservation struct {
}

type TaintParameters struct {

	// +kubebuilder:validation:Required
	Effect string `json:"effect" tf:"effect"`

	// +kubebuilder:validation:Required
	Key string `json:"key" tf:"key"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value"`
}

// EksNodeGroupSpec defines the desired state of EksNodeGroup
type EksNodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EksNodeGroupParameters `json:"forProvider"`
}

// EksNodeGroupStatus defines the observed state of EksNodeGroup.
type EksNodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EksNodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EksNodeGroup is the Schema for the EksNodeGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EksNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksNodeGroupSpec   `json:"spec"`
	Status            EksNodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksNodeGroupList contains a list of EksNodeGroups
type EksNodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksNodeGroup `json:"items"`
}

// Repository type metadata.
var (
	EksNodeGroupKind             = "EksNodeGroup"
	EksNodeGroupGroupKind        = schema.GroupKind{Group: Group, Kind: EksNodeGroupKind}.String()
	EksNodeGroupKindAPIVersion   = EksNodeGroupKind + "." + GroupVersion.String()
	EksNodeGroupGroupVersionKind = GroupVersion.WithKind(EksNodeGroupKind)
)

func init() {
	SchemeBuilder.Register(&EksNodeGroup{}, &EksNodeGroupList{})
}
