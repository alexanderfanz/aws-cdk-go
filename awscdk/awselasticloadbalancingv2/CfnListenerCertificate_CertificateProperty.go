package awselasticloadbalancingv2


// Specifies an SSL server certificate for the certificate list of a secure listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProperty := &CertificateProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   }
//
type CfnListenerCertificate_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the certificate.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
}

