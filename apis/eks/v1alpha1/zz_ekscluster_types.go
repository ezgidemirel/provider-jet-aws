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

type CertificateAuthorityObservation struct {
	Data string `json:"data,omitempty" tf:"data"`
}

type CertificateAuthorityParameters struct {
}

type EksClusterObservation struct {
	Arn string `json:"arn,omitempty" tf:"arn"`

	CertificateAuthority []CertificateAuthorityObservation `json:"certificateAuthority,omitempty" tf:"certificate_authority"`

	CreatedAt string `json:"createdAt,omitempty" tf:"created_at"`

	Endpoint string `json:"endpoint,omitempty" tf:"endpoint"`

	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity"`

	PlatformVersion string `json:"platformVersion,omitempty" tf:"platform_version"`

	Status string `json:"status,omitempty" tf:"status"`
}

type EksClusterParameters struct {

	// +kubebuilder:validation:Optional
	EnabledClusterLogTypes []string `json:"enabledClusterLogTypes,omitempty" tf:"enabled_cluster_log_types"`

	// +kubebuilder:validation:Optional
	EncryptionConfig []EncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config"`

	// +kubebuilder:validation:Optional
	KubernetesNetworkConfig []KubernetesNetworkConfigParameters `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RoleArn string `json:"roleArn" tf:"role_arn"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`

	// +kubebuilder:validation:Required
	VpcConfig []VpcConfigParameters `json:"vpcConfig" tf:"vpc_config"`
}

type EncryptionConfigObservation struct {
}

type EncryptionConfigParameters struct {

	// +kubebuilder:validation:Required
	Provider []ProviderParameters `json:"provider" tf:"provider"`

	// +kubebuilder:validation:Required
	Resources []string `json:"resources" tf:"resources"`
}

type IdentityObservation struct {
	Oidc []OidcObservation `json:"oidc,omitempty" tf:"oidc"`
}

type IdentityParameters struct {
}

type KubernetesNetworkConfigObservation struct {
}

type KubernetesNetworkConfigParameters struct {

	// +kubebuilder:validation:Optional
	ServiceIPv4Cidr *string `json:"serviceIpv4Cidr,omitempty" tf:"service_ipv4_cidr"`
}

type OidcObservation struct {
	Issuer string `json:"issuer,omitempty" tf:"issuer"`
}

type OidcParameters struct {
}

type ProviderObservation struct {
}

type ProviderParameters struct {

	// +kubebuilder:validation:Required
	KeyArn string `json:"keyArn" tf:"key_arn"`
}

type VpcConfigObservation struct {
	ClusterSecurityGroupID string `json:"clusterSecurityGroupId,omitempty" tf:"cluster_security_group_id"`

	VpcID string `json:"vpcId,omitempty" tf:"vpc_id"`
}

type VpcConfigParameters struct {

	// +kubebuilder:validation:Optional
	EndpointPrivateAccess *bool `json:"endpointPrivateAccess,omitempty" tf:"endpoint_private_access"`

	// +kubebuilder:validation:Optional
	EndpointPublicAccess *bool `json:"endpointPublicAccess,omitempty" tf:"endpoint_public_access"`

	// +kubebuilder:validation:Optional
	PublicAccessCidrs []string `json:"publicAccessCidrs,omitempty" tf:"public_access_cidrs"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	// +kubebuilder:validation:Required
	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`
}

// EksClusterSpec defines the desired state of EksCluster
type EksClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EksClusterParameters `json:"forProvider"`
}

// EksClusterStatus defines the observed state of EksCluster.
type EksClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EksClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EksCluster is the Schema for the EksClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EksCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksClusterSpec   `json:"spec"`
	Status            EksClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksClusterList contains a list of EksClusters
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksCluster `json:"items"`
}

// Repository type metadata.
var (
	EksClusterKind             = "EksCluster"
	EksClusterGroupKind        = schema.GroupKind{Group: Group, Kind: EksClusterKind}.String()
	EksClusterKindAPIVersion   = EksClusterKind + "." + GroupVersion.String()
	EksClusterGroupVersionKind = GroupVersion.WithKind(EksClusterKind)
)

func init() {
	SchemeBuilder.Register(&EksCluster{}, &EksClusterList{})
}
