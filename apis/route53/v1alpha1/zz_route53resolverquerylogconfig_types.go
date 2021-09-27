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

type Route53ResolverQueryLogConfigObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	OwnerID string `json:"ownerId,omitempty" tf:"owner_id"`

	ShareStatus string `json:"shareStatus,omitempty" tf:"share_status"`
}

type Route53ResolverQueryLogConfigParameters struct {

	// +kubebuilder:validation:Required
	DestinationArn string `json:"destinationArn" tf:"destination_arn"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// Route53ResolverQueryLogConfigSpec defines the desired state of Route53ResolverQueryLogConfig
type Route53ResolverQueryLogConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Route53ResolverQueryLogConfigParameters `json:"forProvider"`
}

// Route53ResolverQueryLogConfigStatus defines the observed state of Route53ResolverQueryLogConfig.
type Route53ResolverQueryLogConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Route53ResolverQueryLogConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverQueryLogConfig is the Schema for the Route53ResolverQueryLogConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Route53ResolverQueryLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverQueryLogConfigSpec   `json:"spec"`
	Status            Route53ResolverQueryLogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverQueryLogConfigList contains a list of Route53ResolverQueryLogConfigs
type Route53ResolverQueryLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverQueryLogConfig `json:"items"`
}

// Repository type metadata.
var (
	Route53ResolverQueryLogConfigKind             = "Route53ResolverQueryLogConfig"
	Route53ResolverQueryLogConfigGroupKind        = schema.GroupKind{Group: Group, Kind: Route53ResolverQueryLogConfigKind}.String()
	Route53ResolverQueryLogConfigKindAPIVersion   = Route53ResolverQueryLogConfigKind + "." + GroupVersion.String()
	Route53ResolverQueryLogConfigGroupVersionKind = GroupVersion.WithKind(Route53ResolverQueryLogConfigKind)
)

func init() {
	SchemeBuilder.Register(&Route53ResolverQueryLogConfig{}, &Route53ResolverQueryLogConfigList{})
}
