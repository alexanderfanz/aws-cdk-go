package awstimestream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   magneticStoreRejectedDataLocationProperty := &MagneticStoreRejectedDataLocationProperty{
//   	S3Configuration: &S3ConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		EncryptionOption: jsii.String("encryptionOption"),
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	},
//   }
//
type CfnTable_MagneticStoreRejectedDataLocationProperty struct {
	// `CfnTable.MagneticStoreRejectedDataLocationProperty.S3Configuration`.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}
