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

type CloudfrontDistributionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CallerReference string `json:"callerReference" tf:"caller_reference"`

	DomainName string `json:"domainName" tf:"domain_name"`

	Etag string `json:"etag" tf:"etag"`

	HostedZoneId string `json:"hostedZoneId" tf:"hosted_zone_id"`

	InProgressValidationBatches int64 `json:"inProgressValidationBatches" tf:"in_progress_validation_batches"`

	LastModifiedTime string `json:"lastModifiedTime" tf:"last_modified_time"`

	Status string `json:"status" tf:"status"`

	TrustedKeyGroups []TrustedKeyGroupsObservation `json:"trustedKeyGroups" tf:"trusted_key_groups"`

	TrustedSigners []TrustedSignersObservation `json:"trustedSigners" tf:"trusted_signers"`
}

type CloudfrontDistributionParameters struct {
	Aliases []string `json:"aliases,omitempty" tf:"aliases"`

	Comment *string `json:"comment,omitempty" tf:"comment"`

	CustomErrorResponse []CustomErrorResponseParameters `json:"customErrorResponse,omitempty" tf:"custom_error_response"`

	DefaultCacheBehavior []DefaultCacheBehaviorParameters `json:"defaultCacheBehavior" tf:"default_cache_behavior"`

	DefaultRootObject *string `json:"defaultRootObject,omitempty" tf:"default_root_object"`

	Enabled bool `json:"enabled" tf:"enabled"`

	HttpVersion *string `json:"httpVersion,omitempty" tf:"http_version"`

	IsIpv6Enabled *bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6_enabled"`

	LoggingConfig []LoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config"`

	OrderedCacheBehavior []OrderedCacheBehaviorParameters `json:"orderedCacheBehavior,omitempty" tf:"ordered_cache_behavior"`

	Origin []OriginParameters `json:"origin" tf:"origin"`

	OriginGroup []OriginGroupParameters `json:"originGroup,omitempty" tf:"origin_group"`

	PriceClass *string `json:"priceClass,omitempty" tf:"price_class"`

	Region string `json:"region" tf:"-"`

	Restrictions []RestrictionsParameters `json:"restrictions" tf:"restrictions"`

	RetainOnDelete *bool `json:"retainOnDelete,omitempty" tf:"retain_on_delete"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	ViewerCertificate []ViewerCertificateParameters `json:"viewerCertificate" tf:"viewer_certificate"`

	WaitForDeployment *bool `json:"waitForDeployment,omitempty" tf:"wait_for_deployment"`

	WebAclId *string `json:"webAclId,omitempty" tf:"web_acl_id"`
}

type CustomErrorResponseObservation struct {
}

type CustomErrorResponseParameters struct {
	ErrorCachingMinTtl *int64 `json:"errorCachingMinTtl,omitempty" tf:"error_caching_min_ttl"`

	ErrorCode int64 `json:"errorCode" tf:"error_code"`

	ResponseCode *int64 `json:"responseCode,omitempty" tf:"response_code"`

	ResponsePagePath *string `json:"responsePagePath,omitempty" tf:"response_page_path"`
}

type CustomHeaderObservation struct {
}

type CustomHeaderParameters struct {
	Name string `json:"name" tf:"name"`

	Value string `json:"value" tf:"value"`
}

type CustomOriginConfigObservation struct {
}

type CustomOriginConfigParameters struct {
	HttpPort int64 `json:"httpPort" tf:"http_port"`

	HttpsPort int64 `json:"httpsPort" tf:"https_port"`

	OriginKeepaliveTimeout *int64 `json:"originKeepaliveTimeout,omitempty" tf:"origin_keepalive_timeout"`

	OriginProtocolPolicy string `json:"originProtocolPolicy" tf:"origin_protocol_policy"`

	OriginReadTimeout *int64 `json:"originReadTimeout,omitempty" tf:"origin_read_timeout"`

	OriginSslProtocols []string `json:"originSslProtocols" tf:"origin_ssl_protocols"`
}

type DefaultCacheBehaviorObservation struct {
}

type DefaultCacheBehaviorParameters struct {
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`

	CachePolicyId *string `json:"cachePolicyId,omitempty" tf:"cache_policy_id"`

	CachedMethods []string `json:"cachedMethods" tf:"cached_methods"`

	Compress *bool `json:"compress,omitempty" tf:"compress"`

	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`

	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId,omitempty" tf:"field_level_encryption_id"`

	ForwardedValues []ForwardedValuesParameters `json:"forwardedValues,omitempty" tf:"forwarded_values"`

	FunctionAssociation []FunctionAssociationParameters `json:"functionAssociation,omitempty" tf:"function_association"`

	LambdaFunctionAssociation []LambdaFunctionAssociationParameters `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association"`

	MaxTtl *int64 `json:"maxTtl,omitempty" tf:"max_ttl"`

	MinTtl *int64 `json:"minTtl,omitempty" tf:"min_ttl"`

	OriginRequestPolicyId *string `json:"originRequestPolicyId,omitempty" tf:"origin_request_policy_id"`

	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn,omitempty" tf:"realtime_log_config_arn"`

	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`

	TargetOriginId string `json:"targetOriginId" tf:"target_origin_id"`

	TrustedKeyGroups []string `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups"`

	TrustedSigners []string `json:"trustedSigners,omitempty" tf:"trusted_signers"`

	ViewerProtocolPolicy string `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy"`
}

type FailoverCriteriaObservation struct {
}

type FailoverCriteriaParameters struct {
	StatusCodes []int64 `json:"statusCodes" tf:"status_codes"`
}

type ForwardedValuesCookiesObservation struct {
}

type ForwardedValuesCookiesParameters struct {
	Forward string `json:"forward" tf:"forward"`

	WhitelistedNames []string `json:"whitelistedNames,omitempty" tf:"whitelisted_names"`
}

type ForwardedValuesObservation struct {
}

type ForwardedValuesParameters struct {
	Cookies []ForwardedValuesCookiesParameters `json:"cookies" tf:"cookies"`

	Headers []string `json:"headers,omitempty" tf:"headers"`

	QueryString bool `json:"queryString" tf:"query_string"`

	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys"`
}

type FunctionAssociationObservation struct {
}

type FunctionAssociationParameters struct {
	EventType string `json:"eventType" tf:"event_type"`

	FunctionArn string `json:"functionArn" tf:"function_arn"`
}

type GeoRestrictionObservation struct {
}

type GeoRestrictionParameters struct {
	Locations []string `json:"locations,omitempty" tf:"locations"`

	RestrictionType string `json:"restrictionType" tf:"restriction_type"`
}

type ItemsObservation struct {
	KeyGroupId string `json:"keyGroupId" tf:"key_group_id"`

	KeyPairIds []string `json:"keyPairIds" tf:"key_pair_ids"`
}

type ItemsParameters struct {
}

type LambdaFunctionAssociationObservation struct {
}

type LambdaFunctionAssociationParameters struct {
	EventType string `json:"eventType" tf:"event_type"`

	IncludeBody *bool `json:"includeBody,omitempty" tf:"include_body"`

	LambdaArn string `json:"lambdaArn" tf:"lambda_arn"`
}

type LoggingConfigObservation struct {
}

type LoggingConfigParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	IncludeCookies *bool `json:"includeCookies,omitempty" tf:"include_cookies"`

	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type MemberObservation struct {
}

type MemberParameters struct {
	OriginId string `json:"originId" tf:"origin_id"`
}

type OrderedCacheBehaviorForwardedValuesCookiesObservation struct {
}

type OrderedCacheBehaviorForwardedValuesCookiesParameters struct {
	Forward string `json:"forward" tf:"forward"`

	WhitelistedNames []string `json:"whitelistedNames,omitempty" tf:"whitelisted_names"`
}

type OrderedCacheBehaviorForwardedValuesObservation struct {
}

type OrderedCacheBehaviorForwardedValuesParameters struct {
	Cookies []OrderedCacheBehaviorForwardedValuesCookiesParameters `json:"cookies" tf:"cookies"`

	Headers []string `json:"headers,omitempty" tf:"headers"`

	QueryString bool `json:"queryString" tf:"query_string"`

	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys"`
}

type OrderedCacheBehaviorFunctionAssociationObservation struct {
}

type OrderedCacheBehaviorFunctionAssociationParameters struct {
	EventType string `json:"eventType" tf:"event_type"`

	FunctionArn string `json:"functionArn" tf:"function_arn"`
}

type OrderedCacheBehaviorLambdaFunctionAssociationObservation struct {
}

type OrderedCacheBehaviorLambdaFunctionAssociationParameters struct {
	EventType string `json:"eventType" tf:"event_type"`

	IncludeBody *bool `json:"includeBody,omitempty" tf:"include_body"`

	LambdaArn string `json:"lambdaArn" tf:"lambda_arn"`
}

type OrderedCacheBehaviorObservation struct {
}

type OrderedCacheBehaviorParameters struct {
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`

	CachePolicyId *string `json:"cachePolicyId,omitempty" tf:"cache_policy_id"`

	CachedMethods []string `json:"cachedMethods" tf:"cached_methods"`

	Compress *bool `json:"compress,omitempty" tf:"compress"`

	DefaultTtl *int64 `json:"defaultTtl,omitempty" tf:"default_ttl"`

	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId,omitempty" tf:"field_level_encryption_id"`

	ForwardedValues []OrderedCacheBehaviorForwardedValuesParameters `json:"forwardedValues,omitempty" tf:"forwarded_values"`

	FunctionAssociation []OrderedCacheBehaviorFunctionAssociationParameters `json:"functionAssociation,omitempty" tf:"function_association"`

	LambdaFunctionAssociation []OrderedCacheBehaviorLambdaFunctionAssociationParameters `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association"`

	MaxTtl *int64 `json:"maxTtl,omitempty" tf:"max_ttl"`

	MinTtl *int64 `json:"minTtl,omitempty" tf:"min_ttl"`

	OriginRequestPolicyId *string `json:"originRequestPolicyId,omitempty" tf:"origin_request_policy_id"`

	PathPattern string `json:"pathPattern" tf:"path_pattern"`

	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn,omitempty" tf:"realtime_log_config_arn"`

	SmoothStreaming *bool `json:"smoothStreaming,omitempty" tf:"smooth_streaming"`

	TargetOriginId string `json:"targetOriginId" tf:"target_origin_id"`

	TrustedKeyGroups []string `json:"trustedKeyGroups,omitempty" tf:"trusted_key_groups"`

	TrustedSigners []string `json:"trustedSigners,omitempty" tf:"trusted_signers"`

	ViewerProtocolPolicy string `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy"`
}

type OriginGroupObservation struct {
}

type OriginGroupParameters struct {
	FailoverCriteria []FailoverCriteriaParameters `json:"failoverCriteria" tf:"failover_criteria"`

	Member []MemberParameters `json:"member" tf:"member"`

	OriginId string `json:"originId" tf:"origin_id"`
}

type OriginObservation struct {
}

type OriginParameters struct {
	ConnectionAttempts *int64 `json:"connectionAttempts,omitempty" tf:"connection_attempts"`

	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty" tf:"connection_timeout"`

	CustomHeader []CustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header"`

	CustomOriginConfig []CustomOriginConfigParameters `json:"customOriginConfig,omitempty" tf:"custom_origin_config"`

	DomainName string `json:"domainName" tf:"domain_name"`

	OriginId string `json:"originId" tf:"origin_id"`

	OriginPath *string `json:"originPath,omitempty" tf:"origin_path"`

	OriginShield []OriginShieldParameters `json:"originShield,omitempty" tf:"origin_shield"`

	S3OriginConfig []S3OriginConfigParameters `json:"s3OriginConfig,omitempty" tf:"s3_origin_config"`
}

type OriginShieldObservation struct {
}

type OriginShieldParameters struct {
	Enabled bool `json:"enabled" tf:"enabled"`

	OriginShieldRegion string `json:"originShieldRegion" tf:"origin_shield_region"`
}

type RestrictionsObservation struct {
}

type RestrictionsParameters struct {
	GeoRestriction []GeoRestrictionParameters `json:"geoRestriction" tf:"geo_restriction"`
}

type S3OriginConfigObservation struct {
}

type S3OriginConfigParameters struct {
	OriginAccessIdentity string `json:"originAccessIdentity" tf:"origin_access_identity"`
}

type TrustedKeyGroupsObservation struct {
	Enabled bool `json:"enabled" tf:"enabled"`

	Items []ItemsObservation `json:"items" tf:"items"`
}

type TrustedKeyGroupsParameters struct {
}

type TrustedSignersItemsObservation struct {
	AwsAccountNumber string `json:"awsAccountNumber" tf:"aws_account_number"`

	KeyPairIds []string `json:"keyPairIds" tf:"key_pair_ids"`
}

type TrustedSignersItemsParameters struct {
}

type TrustedSignersObservation struct {
	Enabled bool `json:"enabled" tf:"enabled"`

	Items []TrustedSignersItemsObservation `json:"items" tf:"items"`
}

type TrustedSignersParameters struct {
}

type ViewerCertificateObservation struct {
}

type ViewerCertificateParameters struct {
	AcmCertificateArn *string `json:"acmCertificateArn,omitempty" tf:"acm_certificate_arn"`

	CloudfrontDefaultCertificate *bool `json:"cloudfrontDefaultCertificate,omitempty" tf:"cloudfront_default_certificate"`

	IamCertificateId *string `json:"iamCertificateId,omitempty" tf:"iam_certificate_id"`

	MinimumProtocolVersion *string `json:"minimumProtocolVersion,omitempty" tf:"minimum_protocol_version"`

	SslSupportMethod *string `json:"sslSupportMethod,omitempty" tf:"ssl_support_method"`
}

// CloudfrontDistributionSpec defines the desired state of CloudfrontDistribution
type CloudfrontDistributionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudfrontDistributionParameters `json:"forProvider"`
}

// CloudfrontDistributionStatus defines the observed state of CloudfrontDistribution.
type CloudfrontDistributionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudfrontDistributionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontDistribution is the Schema for the CloudfrontDistributions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudfrontDistribution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontDistributionSpec   `json:"spec"`
	Status            CloudfrontDistributionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontDistributionList contains a list of CloudfrontDistributions
type CloudfrontDistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontDistribution `json:"items"`
}

// Repository type metadata.
var (
	CloudfrontDistributionKind             = "CloudfrontDistribution"
	CloudfrontDistributionGroupKind        = schema.GroupKind{Group: Group, Kind: CloudfrontDistributionKind}.String()
	CloudfrontDistributionKindAPIVersion   = CloudfrontDistributionKind + "." + GroupVersion.String()
	CloudfrontDistributionGroupVersionKind = GroupVersion.WithKind(CloudfrontDistributionKind)
)

func init() {
	SchemeBuilder.Register(&CloudfrontDistribution{}, &CloudfrontDistributionList{})
}
