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

type ConfigurationObservation struct {
}

type ConfigurationParameters struct {
	Id *string `json:"id,omitempty" tf:"id"`

	Revision *int64 `json:"revision,omitempty" tf:"revision"`
}

type EncryptionOptionsObservation struct {
}

type EncryptionOptionsParameters struct {
	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	UseAwsOwnedKey *bool `json:"useAwsOwnedKey,omitempty" tf:"use_aws_owned_key"`
}

type InstancesObservation struct {
	ConsoleUrl string `json:"consoleUrl" tf:"console_url"`

	Endpoints []string `json:"endpoints" tf:"endpoints"`

	IpAddress string `json:"ipAddress" tf:"ip_address"`
}

type InstancesParameters struct {
}

type LdapServerMetadataObservation struct {
}

type LdapServerMetadataParameters struct {
	Hosts []string `json:"hosts,omitempty" tf:"hosts"`

	RoleBase *string `json:"roleBase,omitempty" tf:"role_base"`

	RoleName *string `json:"roleName,omitempty" tf:"role_name"`

	RoleSearchMatching *string `json:"roleSearchMatching,omitempty" tf:"role_search_matching"`

	RoleSearchSubtree *bool `json:"roleSearchSubtree,omitempty" tf:"role_search_subtree"`

	ServiceAccountPassword *string `json:"serviceAccountPassword,omitempty" tf:"service_account_password"`

	ServiceAccountUsername *string `json:"serviceAccountUsername,omitempty" tf:"service_account_username"`

	UserBase *string `json:"userBase,omitempty" tf:"user_base"`

	UserRoleName *string `json:"userRoleName,omitempty" tf:"user_role_name"`

	UserSearchMatching *string `json:"userSearchMatching,omitempty" tf:"user_search_matching"`

	UserSearchSubtree *bool `json:"userSearchSubtree,omitempty" tf:"user_search_subtree"`
}

type LogsObservation struct {
}

type LogsParameters struct {
	Audit *string `json:"audit,omitempty" tf:"audit"`

	General *bool `json:"general,omitempty" tf:"general"`
}

type MaintenanceWindowStartTimeObservation struct {
}

type MaintenanceWindowStartTimeParameters struct {
	DayOfWeek string `json:"dayOfWeek" tf:"day_of_week"`

	TimeOfDay string `json:"timeOfDay" tf:"time_of_day"`

	TimeZone string `json:"timeZone" tf:"time_zone"`
}

type MqBrokerObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Instances []InstancesObservation `json:"instances" tf:"instances"`
}

type MqBrokerParameters struct {
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty" tf:"authentication_strategy"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`

	BrokerName string `json:"brokerName" tf:"broker_name"`

	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration"`

	DeploymentMode *string `json:"deploymentMode,omitempty" tf:"deployment_mode"`

	EncryptionOptions []EncryptionOptionsParameters `json:"encryptionOptions,omitempty" tf:"encryption_options"`

	EngineType string `json:"engineType" tf:"engine_type"`

	EngineVersion string `json:"engineVersion" tf:"engine_version"`

	HostInstanceType string `json:"hostInstanceType" tf:"host_instance_type"`

	LdapServerMetadata []LdapServerMetadataParameters `json:"ldapServerMetadata,omitempty" tf:"ldap_server_metadata"`

	Logs []LogsParameters `json:"logs,omitempty" tf:"logs"`

	MaintenanceWindowStartTime []MaintenanceWindowStartTimeParameters `json:"maintenanceWindowStartTime,omitempty" tf:"maintenance_window_start_time"`

	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible"`

	Region string `json:"region" tf:"-"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type"`

	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	User []UserParameters `json:"user" tf:"user"`
}

type UserObservation struct {
}

type UserParameters struct {
	ConsoleAccess *bool `json:"consoleAccess,omitempty" tf:"console_access"`

	Groups []string `json:"groups,omitempty" tf:"groups"`

	Password string `json:"password" tf:"password"`

	Username string `json:"username" tf:"username"`
}

// MqBrokerSpec defines the desired state of MqBroker
type MqBrokerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MqBrokerParameters `json:"forProvider"`
}

// MqBrokerStatus defines the observed state of MqBroker.
type MqBrokerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MqBrokerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MqBroker is the Schema for the MqBrokers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type MqBroker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MqBrokerSpec   `json:"spec"`
	Status            MqBrokerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MqBrokerList contains a list of MqBrokers
type MqBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MqBroker `json:"items"`
}

// Repository type metadata.
var (
	MqBrokerKind             = "MqBroker"
	MqBrokerGroupKind        = schema.GroupKind{Group: Group, Kind: MqBrokerKind}.String()
	MqBrokerKindAPIVersion   = MqBrokerKind + "." + GroupVersion.String()
	MqBrokerGroupVersionKind = GroupVersion.WithKind(MqBrokerKind)
)

func init() {
	SchemeBuilder.Register(&MqBroker{}, &MqBrokerList{})
}
