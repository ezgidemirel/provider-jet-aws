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

type ApplicationSourceObservation struct {
}

type ApplicationSourceParameters struct {
	CloudformationStackArn *string `json:"cloudformationStackArn,omitempty" tf:"cloudformation_stack_arn"`

	TagFilter []TagFilterParameters `json:"tagFilter,omitempty" tf:"tag_filter"`
}

type AutoscalingplansScalingPlanObservation struct {
	ScalingPlanVersion int64 `json:"scalingPlanVersion" tf:"scaling_plan_version"`
}

type AutoscalingplansScalingPlanParameters struct {
	ApplicationSource []ApplicationSourceParameters `json:"applicationSource" tf:"application_source"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	ScalingInstruction []ScalingInstructionParameters `json:"scalingInstruction" tf:"scaling_instruction"`
}

type CustomizedLoadMetricSpecificationObservation struct {
}

type CustomizedLoadMetricSpecificationParameters struct {
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions"`

	MetricName string `json:"metricName" tf:"metric_name"`

	Namespace string `json:"namespace" tf:"namespace"`

	Statistic string `json:"statistic" tf:"statistic"`

	Unit *string `json:"unit,omitempty" tf:"unit"`
}

type CustomizedScalingMetricSpecificationObservation struct {
}

type CustomizedScalingMetricSpecificationParameters struct {
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions"`

	MetricName string `json:"metricName" tf:"metric_name"`

	Namespace string `json:"namespace" tf:"namespace"`

	Statistic string `json:"statistic" tf:"statistic"`

	Unit *string `json:"unit,omitempty" tf:"unit"`
}

type PredefinedLoadMetricSpecificationObservation struct {
}

type PredefinedLoadMetricSpecificationParameters struct {
	PredefinedLoadMetricType string `json:"predefinedLoadMetricType" tf:"predefined_load_metric_type"`

	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label"`
}

type PredefinedScalingMetricSpecificationObservation struct {
}

type PredefinedScalingMetricSpecificationParameters struct {
	PredefinedScalingMetricType string `json:"predefinedScalingMetricType" tf:"predefined_scaling_metric_type"`

	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label"`
}

type ScalingInstructionObservation struct {
}

type ScalingInstructionParameters struct {
	CustomizedLoadMetricSpecification []CustomizedLoadMetricSpecificationParameters `json:"customizedLoadMetricSpecification,omitempty" tf:"customized_load_metric_specification"`

	DisableDynamicScaling *bool `json:"disableDynamicScaling,omitempty" tf:"disable_dynamic_scaling"`

	MaxCapacity int64 `json:"maxCapacity" tf:"max_capacity"`

	MinCapacity int64 `json:"minCapacity" tf:"min_capacity"`

	PredefinedLoadMetricSpecification []PredefinedLoadMetricSpecificationParameters `json:"predefinedLoadMetricSpecification,omitempty" tf:"predefined_load_metric_specification"`

	PredictiveScalingMaxCapacityBehavior *string `json:"predictiveScalingMaxCapacityBehavior,omitempty" tf:"predictive_scaling_max_capacity_behavior"`

	PredictiveScalingMaxCapacityBuffer *int64 `json:"predictiveScalingMaxCapacityBuffer,omitempty" tf:"predictive_scaling_max_capacity_buffer"`

	PredictiveScalingMode *string `json:"predictiveScalingMode,omitempty" tf:"predictive_scaling_mode"`

	ResourceId string `json:"resourceId" tf:"resource_id"`

	ScalableDimension string `json:"scalableDimension" tf:"scalable_dimension"`

	ScalingPolicyUpdateBehavior *string `json:"scalingPolicyUpdateBehavior,omitempty" tf:"scaling_policy_update_behavior"`

	ScheduledActionBufferTime *int64 `json:"scheduledActionBufferTime,omitempty" tf:"scheduled_action_buffer_time"`

	ServiceNamespace string `json:"serviceNamespace" tf:"service_namespace"`

	TargetTrackingConfiguration []TargetTrackingConfigurationParameters `json:"targetTrackingConfiguration" tf:"target_tracking_configuration"`
}

type TagFilterObservation struct {
}

type TagFilterParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values,omitempty" tf:"values"`
}

type TargetTrackingConfigurationObservation struct {
}

type TargetTrackingConfigurationParameters struct {
	CustomizedScalingMetricSpecification []CustomizedScalingMetricSpecificationParameters `json:"customizedScalingMetricSpecification,omitempty" tf:"customized_scaling_metric_specification"`

	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in"`

	EstimatedInstanceWarmup *int64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup"`

	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationParameters `json:"predefinedScalingMetricSpecification,omitempty" tf:"predefined_scaling_metric_specification"`

	ScaleInCooldown *int64 `json:"scaleInCooldown,omitempty" tf:"scale_in_cooldown"`

	ScaleOutCooldown *int64 `json:"scaleOutCooldown,omitempty" tf:"scale_out_cooldown"`

	TargetValue float64 `json:"targetValue" tf:"target_value"`
}

// AutoscalingplansScalingPlanSpec defines the desired state of AutoscalingplansScalingPlan
type AutoscalingplansScalingPlanSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AutoscalingplansScalingPlanParameters `json:"forProvider"`
}

// AutoscalingplansScalingPlanStatus defines the observed state of AutoscalingplansScalingPlan.
type AutoscalingplansScalingPlanStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AutoscalingplansScalingPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingplansScalingPlan is the Schema for the AutoscalingplansScalingPlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AutoscalingplansScalingPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingplansScalingPlanSpec   `json:"spec"`
	Status            AutoscalingplansScalingPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingplansScalingPlanList contains a list of AutoscalingplansScalingPlans
type AutoscalingplansScalingPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingplansScalingPlan `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingplansScalingPlanKind             = "AutoscalingplansScalingPlan"
	AutoscalingplansScalingPlanGroupKind        = schema.GroupKind{Group: Group, Kind: AutoscalingplansScalingPlanKind}.String()
	AutoscalingplansScalingPlanKindAPIVersion   = AutoscalingplansScalingPlanKind + "." + GroupVersion.String()
	AutoscalingplansScalingPlanGroupVersionKind = GroupVersion.WithKind(AutoscalingplansScalingPlanKind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingplansScalingPlan{}, &AutoscalingplansScalingPlanList{})
}
