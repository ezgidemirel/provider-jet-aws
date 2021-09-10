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

type CapacityRebalanceObservation struct {
}

type CapacityRebalanceParameters struct {
	ReplacementStrategy *string `json:"replacementStrategy,omitempty" tf:"replacement_strategy"`
}

type EbsBlockDeviceObservation struct {
}

type EbsBlockDeviceParameters struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	DeviceName string `json:"deviceName" tf:"device_name"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	SnapshotId *string `json:"snapshotId,omitempty" tf:"snapshot_id"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type EphemeralBlockDeviceObservation struct {
}

type EphemeralBlockDeviceParameters struct {
	DeviceName string `json:"deviceName" tf:"device_name"`

	VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type LaunchSpecificationObservation struct {
}

type LaunchSpecificationParameters struct {
	Ami string `json:"ami" tf:"ami"`

	AssociatePublicIpAddress *bool `json:"associatePublicIpAddress,omitempty" tf:"associate_public_ip_address"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	EbsBlockDevice []EbsBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device"`

	EbsOptimized *bool `json:"ebsOptimized,omitempty" tf:"ebs_optimized"`

	EphemeralBlockDevice []EphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device"`

	IamInstanceProfile *string `json:"iamInstanceProfile,omitempty" tf:"iam_instance_profile"`

	IamInstanceProfileArn *string `json:"iamInstanceProfileArn,omitempty" tf:"iam_instance_profile_arn"`

	InstanceType string `json:"instanceType" tf:"instance_type"`

	KeyName *string `json:"keyName,omitempty" tf:"key_name"`

	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring"`

	PlacementGroup *string `json:"placementGroup,omitempty" tf:"placement_group"`

	PlacementTenancy *string `json:"placementTenancy,omitempty" tf:"placement_tenancy"`

	RootBlockDevice []RootBlockDeviceParameters `json:"rootBlockDevice,omitempty" tf:"root_block_device"`

	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	UserData *string `json:"userData,omitempty" tf:"user_data"`

	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`

	WeightedCapacity *string `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type LaunchTemplateConfigObservation struct {
}

type LaunchTemplateConfigParameters struct {
	LaunchTemplateSpecification []LaunchTemplateSpecificationParameters `json:"launchTemplateSpecification" tf:"launch_template_specification"`

	Overrides []OverridesParameters `json:"overrides,omitempty" tf:"overrides"`
}

type LaunchTemplateSpecificationObservation struct {
}

type LaunchTemplateSpecificationParameters struct {
	Id *string `json:"id,omitempty" tf:"id"`

	Name *string `json:"name,omitempty" tf:"name"`

	Version *string `json:"version,omitempty" tf:"version"`
}

type OverridesObservation struct {
}

type OverridesParameters struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price"`

	SubnetId *string `json:"subnetId,omitempty" tf:"subnet_id"`

	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type RootBlockDeviceObservation struct {
}

type RootBlockDeviceParameters struct {
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`

	VolumeSize *int64 `json:"volumeSize,omitempty" tf:"volume_size"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type SpotFleetRequestObservation struct {
	ClientToken string `json:"clientToken" tf:"client_token"`

	SpotRequestState string `json:"spotRequestState" tf:"spot_request_state"`
}

type SpotFleetRequestParameters struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy"`

	ExcessCapacityTerminationPolicy *string `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy"`

	FleetType *string `json:"fleetType,omitempty" tf:"fleet_type"`

	IamFleetRole string `json:"iamFleetRole" tf:"iam_fleet_role"`

	InstanceInterruptionBehaviour *string `json:"instanceInterruptionBehaviour,omitempty" tf:"instance_interruption_behaviour"`

	InstancePoolsToUseCount *int64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count"`

	LaunchSpecification []LaunchSpecificationParameters `json:"launchSpecification,omitempty" tf:"launch_specification"`

	LaunchTemplateConfig []LaunchTemplateConfigParameters `json:"launchTemplateConfig,omitempty" tf:"launch_template_config"`

	LoadBalancers []string `json:"loadBalancers,omitempty" tf:"load_balancers"`

	OnDemandAllocationStrategy *string `json:"onDemandAllocationStrategy,omitempty" tf:"on_demand_allocation_strategy"`

	OnDemandMaxTotalPrice *string `json:"onDemandMaxTotalPrice,omitempty" tf:"on_demand_max_total_price"`

	OnDemandTargetCapacity *int64 `json:"onDemandTargetCapacity,omitempty" tf:"on_demand_target_capacity"`

	Region string `json:"region" tf:"-"`

	ReplaceUnhealthyInstances *bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances"`

	SpotMaintenanceStrategies []SpotMaintenanceStrategiesParameters `json:"spotMaintenanceStrategies,omitempty" tf:"spot_maintenance_strategies"`

	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TargetCapacity int64 `json:"targetCapacity" tf:"target_capacity"`

	TargetGroupArns []string `json:"targetGroupArns,omitempty" tf:"target_group_arns"`

	TerminateInstancesWithExpiration *bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration"`

	ValidFrom *string `json:"validFrom,omitempty" tf:"valid_from"`

	ValidUntil *string `json:"validUntil,omitempty" tf:"valid_until"`

	WaitForFulfillment *bool `json:"waitForFulfillment,omitempty" tf:"wait_for_fulfillment"`
}

type SpotMaintenanceStrategiesObservation struct {
}

type SpotMaintenanceStrategiesParameters struct {
	CapacityRebalance []CapacityRebalanceParameters `json:"capacityRebalance,omitempty" tf:"capacity_rebalance"`
}

// SpotFleetRequestSpec defines the desired state of SpotFleetRequest
type SpotFleetRequestSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SpotFleetRequestParameters `json:"forProvider"`
}

// SpotFleetRequestStatus defines the observed state of SpotFleetRequest.
type SpotFleetRequestStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SpotFleetRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpotFleetRequest is the Schema for the SpotFleetRequests API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SpotFleetRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpotFleetRequestSpec   `json:"spec"`
	Status            SpotFleetRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotFleetRequestList contains a list of SpotFleetRequests
type SpotFleetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotFleetRequest `json:"items"`
}

// Repository type metadata.
var (
	SpotFleetRequestKind             = "SpotFleetRequest"
	SpotFleetRequestGroupKind        = schema.GroupKind{Group: Group, Kind: SpotFleetRequestKind}.String()
	SpotFleetRequestKindAPIVersion   = SpotFleetRequestKind + "." + GroupVersion.String()
	SpotFleetRequestGroupVersionKind = GroupVersion.WithKind(SpotFleetRequestKind)
)

func init() {
	SchemeBuilder.Register(&SpotFleetRequest{}, &SpotFleetRequestList{})
}
