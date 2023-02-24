// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Username and password combination.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   var vpc vpc
//
//
//   awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   		MasterPassword: cdk.SecretValue_UnsafePlainText(jsii.String("tooshort")),
//   	},
//   	Vpc: Vpc,
//   	EnhancedVpcRouting: jsii.Boolean(true),
//   })
//
// Experimental.
type Login struct {
	// Username.
	// Experimental.
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	MasterPassword awscdk.SecretValue `field:"optional" json:"masterPassword" yaml:"masterPassword"`
}

