package awspinpoint


// Properties for defining a `CfnAPNSVoipChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAPNSVoipChannelProps := &CfnAPNSVoipChannelProps{
//   	ApplicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	BundleId: jsii.String("bundleId"),
//   	Certificate: jsii.String("certificate"),
//   	DefaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	Enabled: jsii.Boolean(false),
//   	PrivateKey: jsii.String("privateKey"),
//   	TeamId: jsii.String("teamId"),
//   	TokenKey: jsii.String("tokenKey"),
//   	TokenKeyId: jsii.String("tokenKeyId"),
//   }
//
type CfnAPNSVoipChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the APNs VoIP channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The bundle identifier that's assigned to your iOS app.
	//
	// This identifier is used for APNs tokens.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// The APNs client certificate that you received from Apple.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using an APNs certificate.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The default authentication method that you want Amazon Pinpoint to use when authenticating with APNs.
	//
	// Valid options are `key` or `certificate` .
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether to enable the APNs VoIP channel for the Amazon Pinpoint application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The private key for the APNs client certificate that you want Amazon Pinpoint to use to communicate with APNs.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The identifier that's assigned to your Apple Developer Account team.
	//
	// This identifier is used for APNs tokens.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// The authentication key to use for APNs tokens.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// The key identifier that's assigned to your APNs signing key.
	//
	// Specify this value if you want Amazon Pinpoint to communicate with APNs by using APNs tokens.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

