package awspinpoint


// Properties for defining a `CfnAPNSVoipSandboxChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSVoipSandboxChannelProps := &cfnAPNSVoipSandboxChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	bundleId: jsii.String("bundleId"),
//   	certificate: jsii.String("certificate"),
//   	defaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	enabled: jsii.Boolean(false),
//   	privateKey: jsii.String("privateKey"),
//   	teamId: jsii.String("teamId"),
//   	tokenKey: jsii.String("tokenKey"),
//   	tokenKeyId: jsii.String("tokenKeyId"),
//   }
//
type CfnAPNSVoipSandboxChannelProps struct {
	// The unique identifier for the application that the APNs VoIP sandbox channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with the APNs sandbox environment by using an APNs certificate.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether the APNs VoIP sandbox channel is enabled for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with the APNs sandbox environment.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple developer account team.
	//
	// This identifier is used for APNs tokens.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with the APNs sandbox environment by using APNs tokens.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

