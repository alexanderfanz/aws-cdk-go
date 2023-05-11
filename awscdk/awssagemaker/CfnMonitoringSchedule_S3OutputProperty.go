package awssagemaker


// Information about where and how you want to store the results of a monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputProperty := &S3OutputProperty{
//   	LocalPath: jsii.String("localPath"),
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	S3UploadMode: jsii.String("s3UploadMode"),
//   }
//
type CfnMonitoringSchedule_S3OutputProperty struct {
	// The local path to the S3 storage location where SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the S3 storage location where SageMaker saves the results of a monitoring job.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}
