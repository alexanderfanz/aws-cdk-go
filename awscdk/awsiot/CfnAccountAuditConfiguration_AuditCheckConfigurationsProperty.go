package awsiot


// The types of audit checks that can be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditCheckConfigurationsProperty := &AuditCheckConfigurationsProperty{
//   	AuthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	CaCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	CaCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	ConflictingClientIdsCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeviceCertificateExpiringCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeviceCertificateKeyQualityCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeviceCertificateSharedCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IntermediateCaRevokedForActiveDeviceCertificatesCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IotPolicyOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IoTPolicyPotentialMisConfigurationCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IotRoleAliasAllowsAccessToUnusedServicesCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	IotRoleAliasOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	LoggingDisabledCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	RevokedCaCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	RevokedDeviceCertificateStillActiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	UnauthenticatedCognitoRoleOverlyPermissiveCheck: &AuditCheckConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnAccountAuditConfiguration_AuditCheckConfigurationsProperty struct {
	// Checks the permissiveness of an authenticated Amazon Cognito identity pool role.
	//
	// For this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been used to connect to the AWS IoT message broker during the 31 days before the audit is performed.
	AuthenticatedCognitoRoleOverlyPermissiveCheck interface{} `field:"optional" json:"authenticatedCognitoRoleOverlyPermissiveCheck" yaml:"authenticatedCognitoRoleOverlyPermissiveCheck"`
	// Checks if a CA certificate is expiring.
	//
	// This check applies to CA certificates expiring within 30 days or that have expired.
	CaCertificateExpiringCheck interface{} `field:"optional" json:"caCertificateExpiringCheck" yaml:"caCertificateExpiringCheck"`
	// Checks the quality of the CA certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, and if the key meets a minimum required size. This check applies to CA certificates that are `ACTIVE` or `PENDING_TRANSFER` .
	CaCertificateKeyQualityCheck interface{} `field:"optional" json:"caCertificateKeyQualityCheck" yaml:"caCertificateKeyQualityCheck"`
	// Checks if multiple devices connect using the same client ID.
	ConflictingClientIdsCheck interface{} `field:"optional" json:"conflictingClientIdsCheck" yaml:"conflictingClientIdsCheck"`
	// Checks if a device certificate is expiring.
	//
	// This check applies to device certificates expiring within 30 days or that have expired.
	DeviceCertificateExpiringCheck interface{} `field:"optional" json:"deviceCertificateExpiringCheck" yaml:"deviceCertificateExpiringCheck"`
	// Checks the quality of the device certificate key.
	//
	// The quality checks if the key is in a valid format, not expired, signed by a registered certificate authority, and if the key meets a minimum required size.
	DeviceCertificateKeyQualityCheck interface{} `field:"optional" json:"deviceCertificateKeyQualityCheck" yaml:"deviceCertificateKeyQualityCheck"`
	// Checks if multiple concurrent connections use the same X.509 certificate to authenticate with AWS IoT .
	DeviceCertificateSharedCheck interface{} `field:"optional" json:"deviceCertificateSharedCheck" yaml:"deviceCertificateSharedCheck"`
	// `CfnAccountAuditConfiguration.AuditCheckConfigurationsProperty.IntermediateCaRevokedForActiveDeviceCertificatesCheck`.
	IntermediateCaRevokedForActiveDeviceCertificatesCheck interface{} `field:"optional" json:"intermediateCaRevokedForActiveDeviceCertificatesCheck" yaml:"intermediateCaRevokedForActiveDeviceCertificatesCheck"`
	// Checks the permissiveness of a policy attached to an authenticated Amazon Cognito identity pool role.
	IotPolicyOverlyPermissiveCheck interface{} `field:"optional" json:"iotPolicyOverlyPermissiveCheck" yaml:"iotPolicyOverlyPermissiveCheck"`
	// `CfnAccountAuditConfiguration.AuditCheckConfigurationsProperty.IoTPolicyPotentialMisConfigurationCheck`.
	IoTPolicyPotentialMisConfigurationCheck interface{} `field:"optional" json:"ioTPolicyPotentialMisConfigurationCheck" yaml:"ioTPolicyPotentialMisConfigurationCheck"`
	// Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.
	IotRoleAliasAllowsAccessToUnusedServicesCheck interface{} `field:"optional" json:"iotRoleAliasAllowsAccessToUnusedServicesCheck" yaml:"iotRoleAliasAllowsAccessToUnusedServicesCheck"`
	// Checks if the temporary credentials provided by AWS IoT role aliases are overly permissive.
	IotRoleAliasOverlyPermissiveCheck interface{} `field:"optional" json:"iotRoleAliasOverlyPermissiveCheck" yaml:"iotRoleAliasOverlyPermissiveCheck"`
	// Checks if AWS IoT logs are disabled.
	LoggingDisabledCheck interface{} `field:"optional" json:"loggingDisabledCheck" yaml:"loggingDisabledCheck"`
	// Checks if a revoked CA certificate is still active.
	RevokedCaCertificateStillActiveCheck interface{} `field:"optional" json:"revokedCaCertificateStillActiveCheck" yaml:"revokedCaCertificateStillActiveCheck"`
	// Checks if a revoked device certificate is still active.
	RevokedDeviceCertificateStillActiveCheck interface{} `field:"optional" json:"revokedDeviceCertificateStillActiveCheck" yaml:"revokedDeviceCertificateStillActiveCheck"`
	// Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too permissive.
	UnauthenticatedCognitoRoleOverlyPermissiveCheck interface{} `field:"optional" json:"unauthenticatedCognitoRoleOverlyPermissiveCheck" yaml:"unauthenticatedCognitoRoleOverlyPermissiveCheck"`
}

