package awsrolesanywhere


// Object representing the TrustAnchor type and its related certificate data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	SourceData: &SourceDataProperty{
//   		AcmPcaArn: jsii.String("acmPcaArn"),
//   		X509CertificateData: jsii.String("x509CertificateData"),
//   	},
//   	SourceType: jsii.String("sourceType"),
//   }
//
type CfnTrustAnchor_SourceProperty struct {
	// A union object representing the data field of the TrustAnchor depending on its type.
	SourceData interface{} `field:"optional" json:"sourceData" yaml:"sourceData"`
	// The type of the TrustAnchor.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}
