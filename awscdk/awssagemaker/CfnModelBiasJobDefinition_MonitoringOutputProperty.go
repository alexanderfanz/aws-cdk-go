package awssagemaker


// The output object for a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringOutputProperty := &MonitoringOutputProperty{
//   	S3Output: &S3OutputProperty{
//   		LocalPath: jsii.String("localPath"),
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		S3UploadMode: jsii.String("s3UploadMode"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_MonitoringOutputProperty struct {
	// The Amazon S3 storage location where the results of a monitoring job are saved.
	S3Output interface{} `field:"required" json:"s3Output" yaml:"s3Output"`
}

