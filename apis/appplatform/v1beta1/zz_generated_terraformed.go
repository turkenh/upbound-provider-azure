/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this SpringCloudActiveDeployment
func (mg *SpringCloudActiveDeployment) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_active_deployment"
}

// GetConnectionDetailsMapping for this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudActiveDeployment
func (tr *SpringCloudActiveDeployment) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudActiveDeployment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudActiveDeployment) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudActiveDeploymentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudActiveDeployment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudApp
func (mg *SpringCloudApp) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_app"
}

// GetConnectionDetailsMapping for this SpringCloudApp
func (tr *SpringCloudApp) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudApp
func (tr *SpringCloudApp) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudApp
func (tr *SpringCloudApp) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudApp
func (tr *SpringCloudApp) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudApp
func (tr *SpringCloudApp) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudApp
func (tr *SpringCloudApp) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudApp using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudApp) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudAppParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudApp) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudAppCosmosDBAssociation
func (mg *SpringCloudAppCosmosDBAssociation) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_app_cosmosdb_association"
}

// GetConnectionDetailsMapping for this SpringCloudAppCosmosDBAssociation
func (tr *SpringCloudAppCosmosDBAssociation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudAppCosmosDBAssociation
func (tr *SpringCloudAppCosmosDBAssociation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudAppCosmosDBAssociation
func (tr *SpringCloudAppCosmosDBAssociation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudAppCosmosDBAssociation
func (tr *SpringCloudAppCosmosDBAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudAppCosmosDBAssociation
func (tr *SpringCloudAppCosmosDBAssociation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudAppCosmosDBAssociation
func (tr *SpringCloudAppCosmosDBAssociation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudAppCosmosDBAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudAppCosmosDBAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudAppCosmosDBAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudAppCosmosDBAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudAppMySQLAssociation
func (mg *SpringCloudAppMySQLAssociation) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_app_mysql_association"
}

// GetConnectionDetailsMapping for this SpringCloudAppMySQLAssociation
func (tr *SpringCloudAppMySQLAssociation) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"password": "spec.forProvider.passwordSecretRef"}
}

// GetObservation of this SpringCloudAppMySQLAssociation
func (tr *SpringCloudAppMySQLAssociation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudAppMySQLAssociation
func (tr *SpringCloudAppMySQLAssociation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudAppMySQLAssociation
func (tr *SpringCloudAppMySQLAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudAppMySQLAssociation
func (tr *SpringCloudAppMySQLAssociation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudAppMySQLAssociation
func (tr *SpringCloudAppMySQLAssociation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudAppMySQLAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudAppMySQLAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudAppMySQLAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudAppMySQLAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudAppRedisAssociation
func (mg *SpringCloudAppRedisAssociation) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_app_redis_association"
}

// GetConnectionDetailsMapping for this SpringCloudAppRedisAssociation
func (tr *SpringCloudAppRedisAssociation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudAppRedisAssociation
func (tr *SpringCloudAppRedisAssociation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudAppRedisAssociation
func (tr *SpringCloudAppRedisAssociation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudAppRedisAssociation
func (tr *SpringCloudAppRedisAssociation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudAppRedisAssociation
func (tr *SpringCloudAppRedisAssociation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudAppRedisAssociation
func (tr *SpringCloudAppRedisAssociation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudAppRedisAssociation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudAppRedisAssociation) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudAppRedisAssociationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudAppRedisAssociation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudCertificate
func (mg *SpringCloudCertificate) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_certificate"
}

// GetConnectionDetailsMapping for this SpringCloudCertificate
func (tr *SpringCloudCertificate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudCertificate
func (tr *SpringCloudCertificate) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudCertificate
func (tr *SpringCloudCertificate) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudCertificate
func (tr *SpringCloudCertificate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudCertificate
func (tr *SpringCloudCertificate) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudCertificate
func (tr *SpringCloudCertificate) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudCertificate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudCertificate) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudCertificateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudCertificate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudCustomDomain
func (mg *SpringCloudCustomDomain) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_custom_domain"
}

// GetConnectionDetailsMapping for this SpringCloudCustomDomain
func (tr *SpringCloudCustomDomain) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudCustomDomain
func (tr *SpringCloudCustomDomain) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudCustomDomain
func (tr *SpringCloudCustomDomain) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudCustomDomain
func (tr *SpringCloudCustomDomain) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudCustomDomain
func (tr *SpringCloudCustomDomain) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudCustomDomain
func (tr *SpringCloudCustomDomain) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudCustomDomain using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudCustomDomain) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudCustomDomainParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudCustomDomain) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudJavaDeployment
func (mg *SpringCloudJavaDeployment) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_java_deployment"
}

// GetConnectionDetailsMapping for this SpringCloudJavaDeployment
func (tr *SpringCloudJavaDeployment) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudJavaDeployment
func (tr *SpringCloudJavaDeployment) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudJavaDeployment
func (tr *SpringCloudJavaDeployment) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudJavaDeployment
func (tr *SpringCloudJavaDeployment) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudJavaDeployment
func (tr *SpringCloudJavaDeployment) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudJavaDeployment
func (tr *SpringCloudJavaDeployment) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudJavaDeployment using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudJavaDeployment) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudJavaDeploymentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudJavaDeployment) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudService
func (mg *SpringCloudService) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_service"
}

// GetConnectionDetailsMapping for this SpringCloudService
func (tr *SpringCloudService) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"config_server_git_setting[*].http_basic_auth[*].password": "spec.forProvider.configServerGitSetting[*].httpBasicAuth[*].passwordSecretRef", "config_server_git_setting[*].repository[*].http_basic_auth[*].password": "spec.forProvider.configServerGitSetting[*].repository[*].httpBasicAuth[*].passwordSecretRef", "config_server_git_setting[*].repository[*].ssh_auth[*].host_key": "spec.forProvider.configServerGitSetting[*].repository[*].sshAuth[*].hostKeySecretRef", "config_server_git_setting[*].repository[*].ssh_auth[*].private_key": "spec.forProvider.configServerGitSetting[*].repository[*].sshAuth[*].privateKeySecretRef", "config_server_git_setting[*].ssh_auth[*].host_key": "spec.forProvider.configServerGitSetting[*].sshAuth[*].hostKeySecretRef", "config_server_git_setting[*].ssh_auth[*].private_key": "spec.forProvider.configServerGitSetting[*].sshAuth[*].privateKeySecretRef"}
}

// GetObservation of this SpringCloudService
func (tr *SpringCloudService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudService
func (tr *SpringCloudService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudService
func (tr *SpringCloudService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudService
func (tr *SpringCloudService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudService
func (tr *SpringCloudService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudService) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudService) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SpringCloudStorage
func (mg *SpringCloudStorage) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_storage"
}

// GetConnectionDetailsMapping for this SpringCloudStorage
func (tr *SpringCloudStorage) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SpringCloudStorage
func (tr *SpringCloudStorage) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudStorage
func (tr *SpringCloudStorage) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudStorage
func (tr *SpringCloudStorage) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudStorage
func (tr *SpringCloudStorage) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudStorage
func (tr *SpringCloudStorage) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SpringCloudStorage using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudStorage) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudStorageParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudStorage) GetTerraformSchemaVersion() int {
	return 0
}
