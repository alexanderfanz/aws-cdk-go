package awsappmesh


// An object that represents a Transport Layer Security (TLS) validation context trust.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsValidationContextTrustProperty := &TlsValidationContextTrustProperty{
//   	Acm: &TlsValidationContextAcmTrustProperty{
//   		CertificateAuthorityArns: []*string{
//   			jsii.String("certificateAuthorityArns"),
//   		},
//   	},
//   	File: &TlsValidationContextFileTrustProperty{
//   		CertificateChain: jsii.String("certificateChain"),
//   	},
//   	Sds: &TlsValidationContextSdsTrustProperty{
//   		SecretName: jsii.String("secretName"),
//   	},
//   }
//
type CfnVirtualNode_TlsValidationContextTrustProperty struct {
	// A reference to an object that represents a Transport Layer Security (TLS) validation context trust for an AWS Certificate Manager certificate.
	Acm interface{} `field:"optional" json:"acm" yaml:"acm"`
	// An object that represents a Transport Layer Security (TLS) validation context trust for a local file.
	File interface{} `field:"optional" json:"file" yaml:"file"`
	// A reference to an object that represents a Transport Layer Security (TLS) Secret Discovery Service validation context trust.
	Sds interface{} `field:"optional" json:"sds" yaml:"sds"`
}
