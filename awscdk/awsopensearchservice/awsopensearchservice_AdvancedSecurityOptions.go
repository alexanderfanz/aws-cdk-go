package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies options for fine-grained access control.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	enforceHttps: jsii.Boolean(true),
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   	fineGrainedAccessControl: &advancedSecurityOptions{
//   		masterUserName: jsii.String("master-user"),
//   	},
//   	logging: &loggingOptions{
//   		auditLogEnabled: jsii.Boolean(true),
//   		slowSearchLogEnabled: jsii.Boolean(true),
//   		appLogEnabled: jsii.Boolean(true),
//   		slowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
type AdvancedSecurityOptions struct {
	// ARN for the master user.
	//
	// Only specify this or masterUserName, but not both.
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify this or masterUserArn, but not both.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Password for the master user.
	//
	// You can use `SecretValue.unsafePlainText` to specify a password in plain text or
	// use `secretsmanager.Secret.fromSecretAttributes` to reference a secret in
	// Secrets Manager.
	MasterUserPassword awscdk.SecretValue `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
}

