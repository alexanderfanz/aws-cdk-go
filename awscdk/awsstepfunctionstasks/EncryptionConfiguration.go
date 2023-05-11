package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Encryption Configuration of the S3 bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &AthenaStartQueryExecutionProps{
//   	QueryString: sfn.JsonPath_StringAt(jsii.String("$.queryString")),
//   	QueryExecutionContext: &QueryExecutionContext{
//   		DatabaseName: jsii.String("mydatabase"),
//   	},
//   	ResultConfiguration: &ResultConfiguration{
//   		EncryptionConfiguration: &EncryptionConfiguration{
//   			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
//   		},
//   		OutputLocation: &Location{
//   			BucketName: jsii.String("query-results-bucket"),
//   			ObjectKey: jsii.String("folder"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html
//
type EncryptionConfiguration struct {
	// Type of S3 server-side encryption enabled.
	EncryptionOption EncryptionOption `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// KMS key ARN or ID.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}
