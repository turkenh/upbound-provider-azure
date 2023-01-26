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

type ActionIncidentObservation struct {
}

type ActionIncidentParameters struct {

	// The classification of the incident, when closing it. Possible values are: BenignPositive_SuspiciousButExpected, FalsePositive_InaccurateData, FalsePositive_IncorrectAlertLogic, TruePositive_SuspiciousActivity and Undetermined.
	// +kubebuilder:validation:Optional
	Classification *string `json:"classification,omitempty" tf:"classification,omitempty"`

	// The comment why the incident is to be closed.
	// +kubebuilder:validation:Optional
	ClassificationComment *string `json:"classificationComment,omitempty" tf:"classification_comment,omitempty"`

	// Specifies a list of labels to add to the incident.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The execution order of this action.
	// +kubebuilder:validation:Required
	Order *float64 `json:"order" tf:"order,omitempty"`

	// The object ID of the entity this incident is assigned to.
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The severity to add to the incident. Possible values are High, Informational, Low and Medium.
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`

	// The status to set to the incident. Possible values are: Active, Closed, New.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ActionPlaybookObservation struct {
}

type ActionPlaybookParameters struct {

	// The ID of the Logic App that defines the playbook's logic.
	// +kubebuilder:validation:Required
	LogicAppID *string `json:"logicAppId" tf:"logic_app_id,omitempty"`

	// The execution order of this action.
	// +kubebuilder:validation:Required
	Order *float64 `json:"order" tf:"order,omitempty"`

	// The ID of the Tenant that owns the playbook.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// The operator to use for evaluate the condition. Possible values include: Equals, NotEquals, Contains, NotContains, StartsWith, NotStartsWith, EndsWith, NotEndsWith.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// The property to use for evaluate the condition. Possible values include: AccountAadTenantId, AccountAadUserId, AccountNTDomain, AccountName, AccountObjectGuid, AccountPUID, AccountSid, AccountUPNSuffix, AzureResourceResourceId, AzureResourceSubscriptionId, CloudApplicationAppId, CloudApplicationAppName, DNSDomainName, FileDirectory, FileHashValue, FileName, HostAzureID, HostNTDomain, HostName, HostNetBiosName, HostOSVersion, IPAddress, IncidentDescription, IncidentProviderName, IncidentRelatedAnalyticRuleIds, IncidentSeverity, IncidentStatus, IncidentTactics, IncidentTitle, IoTDeviceId, IoTDeviceModel, IoTDeviceName, IoTDeviceOperatingSystem, IoTDeviceType, IoTDeviceVendor, MailMessageDeliveryAction, MailMessageDeliveryLocation, MailMessageP1Sender, MailMessageP2Sender, MailMessageRecipient, MailMessageSenderIP, MailMessageSubject, MailboxDisplayName, MailboxPrimaryAddress, MailboxUPN, MalwareCategory, MalwareName, ProcessCommandLine, ProcessId, RegistryKey, RegistryValueData, Url.
	// +kubebuilder:validation:Required
	Property *string `json:"property" tf:"property,omitempty"`

	// Specifies a list of values to use for evaluate the condition.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type SentinelAutomationRuleObservation struct {

	// The ID of the Sentinel Automation Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SentinelAutomationRuleParameters struct {

	// One or more action_incident blocks as defined below.
	// +kubebuilder:validation:Optional
	ActionIncident []ActionIncidentParameters `json:"actionIncident,omitempty" tf:"action_incident,omitempty"`

	// One or more action_playbook blocks as defined below.
	// +kubebuilder:validation:Optional
	ActionPlaybook []ActionPlaybookParameters `json:"actionPlaybook,omitempty" tf:"action_playbook,omitempty"`

	// One or more condition blocks as defined below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// The display name which should be used for this Sentinel Automation Rule.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Whether this Sentinel Automation Rule is enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The time in RFC3339 format of kind UTC that determines when this Automation Rule should expire and be disabled.
	// +kubebuilder:validation:Optional
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The ID of the Log Analytics Workspace where this Sentinel applies to. Changing this forces a new Sentinel Automation Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationsmanagement/v1beta1.LogAnalyticsSolution
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("workspace_resource_id",false)
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a LogAnalyticsSolution in operationsmanagement to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a LogAnalyticsSolution in operationsmanagement to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The UUID which should be used for this Sentinel Automation Rule. Changing this forces a new Sentinel Automation Rule to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The order of this Sentinel Automation Rule. Possible values varies between 1 and 1000.
	// +kubebuilder:validation:Required
	Order *float64 `json:"order" tf:"order,omitempty"`
}

// SentinelAutomationRuleSpec defines the desired state of SentinelAutomationRule
type SentinelAutomationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAutomationRuleParameters `json:"forProvider"`
}

// SentinelAutomationRuleStatus defines the observed state of SentinelAutomationRule.
type SentinelAutomationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAutomationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAutomationRule is the Schema for the SentinelAutomationRules API. Manages a Sentinel Automation Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAutomationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelAutomationRuleSpec   `json:"spec"`
	Status            SentinelAutomationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAutomationRuleList contains a list of SentinelAutomationRules
type SentinelAutomationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAutomationRule `json:"items"`
}

// Repository type metadata.
var (
	SentinelAutomationRule_Kind             = "SentinelAutomationRule"
	SentinelAutomationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelAutomationRule_Kind}.String()
	SentinelAutomationRule_KindAPIVersion   = SentinelAutomationRule_Kind + "." + CRDGroupVersion.String()
	SentinelAutomationRule_GroupVersionKind = CRDGroupVersion.WithKind(SentinelAutomationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelAutomationRule{}, &SentinelAutomationRuleList{})
}
