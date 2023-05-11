// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for configuring a Redshift user.
//
// Example:
//   user := awscdkredshiftalpha.NewUser(this, jsii.String("User"), &UserProps{
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   })
//   cluster.AddRotationMultiUser(jsii.String("MultiUserRotation"), &RotationMultiUserOptions{
//   	Secret: user.Secret,
//   })
//
// Experimental.
type UserProps struct {
	// The cluster containing the database.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The secret containing credentials to a Redshift user with administrator privileges.
	//
	// Secret JSON schema: `{ username: string; password: string }`.
	// Experimental.
	AdminUser awssecretsmanager.ISecret `field:"optional" json:"adminUser" yaml:"adminUser"`
	// KMS key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The policy to apply when this resource is removed from the application.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The name of the user.
	//
	// For valid values, see: https://docs.aws.amazon.com/redshift/latest/dg/r_names.html
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}
