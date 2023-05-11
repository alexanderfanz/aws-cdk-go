package awsmediapackage


// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that provides encryption keys.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spekeKeyProviderProperty := &SpekeKeyProviderProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SystemIds: []*string{
//   		jsii.String("systemIds"),
//   	},
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   	},
//   }
//
type CfnPackagingConfiguration_SpekeKeyProviderProperty struct {
	// The ARN for the IAM role that's granted by the key provider to provide access to the key provider API.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.
	SystemIds *[]*string `field:"required" json:"systemIds" yaml:"systemIds"`
	// URL for the key provider's key retrieval API endpoint.
	//
	// Must start with https://.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Use `encryptionContractConfiguration` to configure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.
	EncryptionContractConfiguration interface{} `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
}
