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

type ActionObservation struct {
}

type ActionParameters struct {
	Target []TargetParameters `json:"target" tf:"target"`
}

type ActionTargetObservation struct {
}

type ActionTargetParameters struct {
	VirtualService []TargetVirtualServiceParameters `json:"virtualService" tf:"virtual_service"`
}

type ActionTargetVirtualServiceObservation struct {
}

type ActionTargetVirtualServiceParameters struct {
	VirtualServiceName string `json:"virtualServiceName" tf:"virtual_service_name"`
}

type AppmeshGatewayRouteObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	ResourceOwner string `json:"resourceOwner" tf:"resource_owner"`
}

type AppmeshGatewayRouteParameters struct {
	MeshName string `json:"meshName" tf:"mesh_name"`

	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`

	Name string `json:"name" tf:"name"`

	Region string `json:"region" tf:"-"`

	Spec []SpecParameters `json:"spec" tf:"spec"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VirtualGatewayName string `json:"virtualGatewayName" tf:"virtual_gateway_name"`
}

type GrpcRouteObservation struct {
}

type GrpcRouteParameters struct {
	Action []ActionParameters `json:"action" tf:"action"`

	Match []MatchParameters `json:"match" tf:"match"`
}

type Http2RouteActionObservation struct {
}

type Http2RouteActionParameters struct {
	Target []ActionTargetParameters `json:"target" tf:"target"`
}

type Http2RouteMatchObservation struct {
}

type Http2RouteMatchParameters struct {
	Prefix string `json:"prefix" tf:"prefix"`
}

type Http2RouteObservation struct {
}

type Http2RouteParameters struct {
	Action []Http2RouteActionParameters `json:"action" tf:"action"`

	Match []Http2RouteMatchParameters `json:"match" tf:"match"`
}

type HttpRouteActionObservation struct {
}

type HttpRouteActionParameters struct {
	Target []HttpRouteActionTargetParameters `json:"target" tf:"target"`
}

type HttpRouteActionTargetObservation struct {
}

type HttpRouteActionTargetParameters struct {
	VirtualService []ActionTargetVirtualServiceParameters `json:"virtualService" tf:"virtual_service"`
}

type HttpRouteMatchObservation struct {
}

type HttpRouteMatchParameters struct {
	Prefix string `json:"prefix" tf:"prefix"`
}

type HttpRouteObservation struct {
}

type HttpRouteParameters struct {
	Action []HttpRouteActionParameters `json:"action" tf:"action"`

	Match []HttpRouteMatchParameters `json:"match" tf:"match"`
}

type MatchObservation struct {
}

type MatchParameters struct {
	ServiceName string `json:"serviceName" tf:"service_name"`
}

type SpecObservation struct {
}

type SpecParameters struct {
	GrpcRoute []GrpcRouteParameters `json:"grpcRoute,omitempty" tf:"grpc_route"`

	Http2Route []Http2RouteParameters `json:"http2Route,omitempty" tf:"http2_route"`

	HttpRoute []HttpRouteParameters `json:"httpRoute,omitempty" tf:"http_route"`
}

type TargetObservation struct {
}

type TargetParameters struct {
	VirtualService []VirtualServiceParameters `json:"virtualService" tf:"virtual_service"`
}

type TargetVirtualServiceObservation struct {
}

type TargetVirtualServiceParameters struct {
	VirtualServiceName string `json:"virtualServiceName" tf:"virtual_service_name"`
}

type VirtualServiceObservation struct {
}

type VirtualServiceParameters struct {
	VirtualServiceName string `json:"virtualServiceName" tf:"virtual_service_name"`
}

// AppmeshGatewayRouteSpec defines the desired state of AppmeshGatewayRoute
type AppmeshGatewayRouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppmeshGatewayRouteParameters `json:"forProvider"`
}

// AppmeshGatewayRouteStatus defines the observed state of AppmeshGatewayRoute.
type AppmeshGatewayRouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppmeshGatewayRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshGatewayRoute is the Schema for the AppmeshGatewayRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppmeshGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshGatewayRouteSpec   `json:"spec"`
	Status            AppmeshGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshGatewayRouteList contains a list of AppmeshGatewayRoutes
type AppmeshGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshGatewayRoute `json:"items"`
}

// Repository type metadata.
var (
	AppmeshGatewayRouteKind             = "AppmeshGatewayRoute"
	AppmeshGatewayRouteGroupKind        = schema.GroupKind{Group: Group, Kind: AppmeshGatewayRouteKind}.String()
	AppmeshGatewayRouteKindAPIVersion   = AppmeshGatewayRouteKind + "." + GroupVersion.String()
	AppmeshGatewayRouteGroupVersionKind = GroupVersion.WithKind(AppmeshGatewayRouteKind)
)

func init() {
	SchemeBuilder.Register(&AppmeshGatewayRoute{}, &AppmeshGatewayRouteList{})
}
