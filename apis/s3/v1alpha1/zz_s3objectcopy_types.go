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

type S3ObjectCopyGrantObservation struct {
}

type S3ObjectCopyGrantParameters struct {
	Email *string `json:"email,omitempty" tf:"email"`

	Id *string `json:"id,omitempty" tf:"id"`

	Permissions []string `json:"permissions" tf:"permissions"`

	Type string `json:"type" tf:"type"`

	Uri *string `json:"uri,omitempty" tf:"uri"`
}

type S3ObjectCopyObservation struct {
	Etag string `json:"etag" tf:"etag"`

	Expiration string `json:"expiration" tf:"expiration"`

	LastModified string `json:"lastModified" tf:"last_modified"`

	RequestCharged bool `json:"requestCharged" tf:"request_charged"`

	SourceVersionId string `json:"sourceVersionId" tf:"source_version_id"`

	VersionId string `json:"versionId" tf:"version_id"`
}

type S3ObjectCopyParameters struct {
	Acl *string `json:"acl,omitempty" tf:"acl"`

	Bucket string `json:"bucket" tf:"bucket"`

	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled"`

	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control"`

	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition"`

	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding"`

	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language"`

	ContentType *string `json:"contentType,omitempty" tf:"content_type"`

	CopyIfMatch *string `json:"copyIfMatch,omitempty" tf:"copy_if_match"`

	CopyIfModifiedSince *string `json:"copyIfModifiedSince,omitempty" tf:"copy_if_modified_since"`

	CopyIfNoneMatch *string `json:"copyIfNoneMatch,omitempty" tf:"copy_if_none_match"`

	CopyIfUnmodifiedSince *string `json:"copyIfUnmodifiedSince,omitempty" tf:"copy_if_unmodified_since"`

	CustomerAlgorithm *string `json:"customerAlgorithm,omitempty" tf:"customer_algorithm"`

	CustomerKey *string `json:"customerKey,omitempty" tf:"customer_key"`

	CustomerKeyMd5 *string `json:"customerKeyMd5,omitempty" tf:"customer_key_md5"`

	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner"`

	ExpectedSourceBucketOwner *string `json:"expectedSourceBucketOwner,omitempty" tf:"expected_source_bucket_owner"`

	Expires *string `json:"expires,omitempty" tf:"expires"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	Grant []S3ObjectCopyGrantParameters `json:"grant,omitempty" tf:"grant"`

	Key string `json:"key" tf:"key"`

	KmsEncryptionContext *string `json:"kmsEncryptionContext,omitempty" tf:"kms_encryption_context"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata"`

	MetadataDirective *string `json:"metadataDirective,omitempty" tf:"metadata_directive"`

	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status"`

	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode"`

	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date"`

	Region string `json:"region" tf:"-"`

	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer"`

	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`

	Source string `json:"source" tf:"source"`

	SourceCustomerAlgorithm *string `json:"sourceCustomerAlgorithm,omitempty" tf:"source_customer_algorithm"`

	SourceCustomerKey *string `json:"sourceCustomerKey,omitempty" tf:"source_customer_key"`

	SourceCustomerKeyMd5 *string `json:"sourceCustomerKeyMd5,omitempty" tf:"source_customer_key_md5"`

	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class"`

	TaggingDirective *string `json:"taggingDirective,omitempty" tf:"tagging_directive"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect"`
}

// S3ObjectCopySpec defines the desired state of S3ObjectCopy
type S3ObjectCopySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3ObjectCopyParameters `json:"forProvider"`
}

// S3ObjectCopyStatus defines the observed state of S3ObjectCopy.
type S3ObjectCopyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3ObjectCopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3ObjectCopy is the Schema for the S3ObjectCopys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type S3ObjectCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3ObjectCopySpec   `json:"spec"`
	Status            S3ObjectCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3ObjectCopyList contains a list of S3ObjectCopys
type S3ObjectCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3ObjectCopy `json:"items"`
}

// Repository type metadata.
var (
	S3ObjectCopyKind             = "S3ObjectCopy"
	S3ObjectCopyGroupKind        = schema.GroupKind{Group: Group, Kind: S3ObjectCopyKind}.String()
	S3ObjectCopyKindAPIVersion   = S3ObjectCopyKind + "." + GroupVersion.String()
	S3ObjectCopyGroupVersionKind = GroupVersion.WithKind(S3ObjectCopyKind)
)

func init() {
	SchemeBuilder.Register(&S3ObjectCopy{}, &S3ObjectCopyList{})
}
