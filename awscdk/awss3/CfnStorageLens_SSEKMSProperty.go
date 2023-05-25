package awss3


// Specifies the use of server-side encryption using an AWS Key Management Service key (SSE-KMS) to encrypt the delivered S3 Storage Lens metrics export file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSEKMSProperty := &SSEKMSProperty{
//   	KeyId: jsii.String("keyId"),
//   }
//
type CfnStorageLens_SSEKMSProperty struct {
	// Specifies the Amazon Resource Name (ARN) of the customer managed AWS KMS key to use for encrypting the S3 Storage Lens metrics export file.
	//
	// Amazon S3 only supports symmetric encryption keys. For more information, see [Special-purpose keys](https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html) in the *AWS Key Management Service Developer Guide* .
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

