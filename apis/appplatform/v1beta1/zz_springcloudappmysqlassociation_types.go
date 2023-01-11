/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SpringCloudAppMySQLAssociationObservation struct {

	// The ID of the Spring Cloud Application MySQL Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SpringCloudAppMySQLAssociationParameters struct {

	// Specifies the name of the MySQL Database which the Spring Cloud App should be associated with.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// Specifies the ID of the MySQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbformysql/v1beta1.Server
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MySQLServerID *string `json:"mysqlServerId,omitempty" tf:"mysql_server_id,omitempty"`

	// Reference to a Server in dbformysql to populate mysqlServerId.
	// +kubebuilder:validation:Optional
	MySQLServerIDRef *v1.Reference `json:"mysqlServerIdRef,omitempty" tf:"-"`

	// Selector for a Server in dbformysql to populate mysqlServerId.
	// +kubebuilder:validation:Optional
	MySQLServerIDSelector *v1.Selector `json:"mysqlServerIdSelector,omitempty" tf:"-"`

	// Specifies the password which should be used when connecting to the MySQL Database from the Spring Cloud App.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`

	// Specifies the username which should be used when connecting to the MySQL Database from the Spring Cloud App.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbformysql/v1beta1.Server
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("administrator_login",false)
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// Reference to a Server in dbformysql to populate username.
	// +kubebuilder:validation:Optional
	UsernameRef *v1.Reference `json:"usernameRef,omitempty" tf:"-"`

	// Selector for a Server in dbformysql to populate username.
	// +kubebuilder:validation:Optional
	UsernameSelector *v1.Selector `json:"usernameSelector,omitempty" tf:"-"`
}

// SpringCloudAppMySQLAssociationSpec defines the desired state of SpringCloudAppMySQLAssociation
type SpringCloudAppMySQLAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAppMySQLAssociationParameters `json:"forProvider"`
}

// SpringCloudAppMySQLAssociationStatus defines the observed state of SpringCloudAppMySQLAssociation.
type SpringCloudAppMySQLAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAppMySQLAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppMySQLAssociation is the Schema for the SpringCloudAppMySQLAssociations API. Associates a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAppMySQLAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudAppMySQLAssociationSpec   `json:"spec"`
	Status            SpringCloudAppMySQLAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppMySQLAssociationList contains a list of SpringCloudAppMySQLAssociations
type SpringCloudAppMySQLAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAppMySQLAssociation `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAppMySQLAssociation_Kind             = "SpringCloudAppMySQLAssociation"
	SpringCloudAppMySQLAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudAppMySQLAssociation_Kind}.String()
	SpringCloudAppMySQLAssociation_KindAPIVersion   = SpringCloudAppMySQLAssociation_Kind + "." + CRDGroupVersion.String()
	SpringCloudAppMySQLAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudAppMySQLAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAppMySQLAssociation{}, &SpringCloudAppMySQLAssociationList{})
}
