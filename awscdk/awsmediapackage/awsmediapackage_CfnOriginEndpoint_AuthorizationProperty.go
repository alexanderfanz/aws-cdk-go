package awsmediapackage


// Parameters for enabling CDN authorization on the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationProperty := &authorizationProperty{
//   	cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   	secretsRoleArn: jsii.String("secretsRoleArn"),
//   }
//
type CfnOriginEndpoint_AuthorizationProperty struct {
	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that your Content Distribution Network (CDN) uses for authorization to access your endpoint.
	CdnIdentifierSecret *string `field:"required" json:"cdnIdentifierSecret" yaml:"cdnIdentifierSecret"`
	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager .
	SecretsRoleArn *string `field:"required" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}
