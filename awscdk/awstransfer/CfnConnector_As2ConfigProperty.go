package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   as2ConfigProperty := &As2ConfigProperty{
//   	Compression: jsii.String("compression"),
//   	EncryptionAlgorithm: jsii.String("encryptionAlgorithm"),
//   	LocalProfileId: jsii.String("localProfileId"),
//   	MdnResponse: jsii.String("mdnResponse"),
//   	MdnSigningAlgorithm: jsii.String("mdnSigningAlgorithm"),
//   	MessageSubject: jsii.String("messageSubject"),
//   	PartnerProfileId: jsii.String("partnerProfileId"),
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   }
//
type CfnConnector_As2ConfigProperty struct {
	// `CfnConnector.As2ConfigProperty.Compression`.
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// `CfnConnector.As2ConfigProperty.EncryptionAlgorithm`.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// `CfnConnector.As2ConfigProperty.LocalProfileId`.
	LocalProfileId *string `field:"optional" json:"localProfileId" yaml:"localProfileId"`
	// `CfnConnector.As2ConfigProperty.MdnResponse`.
	MdnResponse *string `field:"optional" json:"mdnResponse" yaml:"mdnResponse"`
	// `CfnConnector.As2ConfigProperty.MdnSigningAlgorithm`.
	MdnSigningAlgorithm *string `field:"optional" json:"mdnSigningAlgorithm" yaml:"mdnSigningAlgorithm"`
	// `CfnConnector.As2ConfigProperty.MessageSubject`.
	MessageSubject *string `field:"optional" json:"messageSubject" yaml:"messageSubject"`
	// `CfnConnector.As2ConfigProperty.PartnerProfileId`.
	PartnerProfileId *string `field:"optional" json:"partnerProfileId" yaml:"partnerProfileId"`
	// `CfnConnector.As2ConfigProperty.SigningAlgorithm`.
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
}

