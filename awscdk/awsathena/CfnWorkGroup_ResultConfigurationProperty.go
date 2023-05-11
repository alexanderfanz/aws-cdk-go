package awsathena


// The location in Amazon S3 where query and calculation results are stored and the encryption option, if any, used for query and calculation results.
//
// These are known as "client-side settings". If workgroup settings override client-side settings, then the query uses the workgroup settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resultConfigurationProperty := &ResultConfigurationProperty{
//   	AclConfiguration: &AclConfigurationProperty{
//   		S3AclOption: jsii.String("s3AclOption"),
//   	},
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		EncryptionOption: jsii.String("encryptionOption"),
//
//   		// the properties below are optional
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	OutputLocation: jsii.String("outputLocation"),
//   }
//
type CfnWorkGroup_ResultConfigurationProperty struct {
	// `CfnWorkGroup.ResultConfigurationProperty.AclConfiguration`.
	AclConfiguration interface{} `field:"optional" json:"aclConfiguration" yaml:"aclConfiguration"`
	// If query results are encrypted in Amazon S3, indicates the encryption option used (for example, `SSE_KMS` or `CSE_KMS` ) and key information.
	//
	// This is a client-side setting. If workgroup settings override client-side settings, then the query uses the encryption configuration that is specified for the workgroup, and also uses the location for storing query results specified in the workgroup. See `EnforceWorkGroupConfiguration` and [Workgroup Settings Override Client-Side Settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html) .
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// `CfnWorkGroup.ResultConfigurationProperty.ExpectedBucketOwner`.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// The location in Amazon S3 where your query results are stored, such as `s3://path/to/query/bucket/` .
	//
	// To run a query, you must specify the query results location using either a client-side setting for individual queries or a location specified by the workgroup. If workgroup settings override client-side settings, then the query uses the location specified for the workgroup. If no query location is set, Athena issues an error. For more information, see [Working with Query Results, Output Files, and Query History](https://docs.aws.amazon.com/athena/latest/ug/querying.html) and `EnforceWorkGroupConfiguration` .
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
}
