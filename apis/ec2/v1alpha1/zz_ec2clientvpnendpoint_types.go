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

type AuthenticationOptionsObservation struct {
}

type AuthenticationOptionsParameters struct {
	ActiveDirectoryId *string `json:"activeDirectoryId,omitempty" tf:"active_directory_id"`

	RootCertificateChainArn *string `json:"rootCertificateChainArn,omitempty" tf:"root_certificate_chain_arn"`

	SamlProviderArn *string `json:"samlProviderArn,omitempty" tf:"saml_provider_arn"`

	Type string `json:"type" tf:"type"`
}

type ConnectionLogOptionsObservation struct {
}

type ConnectionLogOptionsParameters struct {
	CloudwatchLogGroup *string `json:"cloudwatchLogGroup,omitempty" tf:"cloudwatch_log_group"`

	CloudwatchLogStream *string `json:"cloudwatchLogStream,omitempty" tf:"cloudwatch_log_stream"`

	Enabled bool `json:"enabled" tf:"enabled"`
}

type Ec2ClientVpnEndpointObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DnsName string `json:"dnsName" tf:"dns_name"`

	Status string `json:"status" tf:"status"`
}

type Ec2ClientVpnEndpointParameters struct {
	AuthenticationOptions []AuthenticationOptionsParameters `json:"authenticationOptions" tf:"authentication_options"`

	ClientCidrBlock string `json:"clientCidrBlock" tf:"client_cidr_block"`

	ConnectionLogOptions []ConnectionLogOptionsParameters `json:"connectionLogOptions" tf:"connection_log_options"`

	Description *string `json:"description,omitempty" tf:"description"`

	DnsServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`

	Region string `json:"region" tf:"-"`

	ServerCertificateArn string `json:"serverCertificateArn" tf:"server_certificate_arn"`

	SplitTunnel *bool `json:"splitTunnel,omitempty" tf:"split_tunnel"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	TransportProtocol *string `json:"transportProtocol,omitempty" tf:"transport_protocol"`
}

// Ec2ClientVpnEndpointSpec defines the desired state of Ec2ClientVpnEndpoint
type Ec2ClientVpnEndpointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2ClientVpnEndpointParameters `json:"forProvider"`
}

// Ec2ClientVpnEndpointStatus defines the observed state of Ec2ClientVpnEndpoint.
type Ec2ClientVpnEndpointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2ClientVpnEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnEndpoint is the Schema for the Ec2ClientVpnEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2ClientVpnEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVpnEndpointSpec   `json:"spec"`
	Status            Ec2ClientVpnEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnEndpointList contains a list of Ec2ClientVpnEndpoints
type Ec2ClientVpnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnEndpoint `json:"items"`
}

// Repository type metadata.
var (
	Ec2ClientVpnEndpointKind             = "Ec2ClientVpnEndpoint"
	Ec2ClientVpnEndpointGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2ClientVpnEndpointKind}.String()
	Ec2ClientVpnEndpointKindAPIVersion   = Ec2ClientVpnEndpointKind + "." + GroupVersion.String()
	Ec2ClientVpnEndpointGroupVersionKind = GroupVersion.WithKind(Ec2ClientVpnEndpointKind)
)

func init() {
	SchemeBuilder.Register(&Ec2ClientVpnEndpoint{}, &Ec2ClientVpnEndpointList{})
}
