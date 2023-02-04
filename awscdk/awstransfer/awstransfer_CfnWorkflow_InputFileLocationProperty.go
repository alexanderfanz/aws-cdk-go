package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputFileLocationProperty := &inputFileLocationProperty{
//   	efsFileLocation: &efsInputFileLocationProperty{
//   		fileSystemId: jsii.String("fileSystemId"),
//   		path: jsii.String("path"),
//   	},
//   	s3FileLocation: &s3InputFileLocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnWorkflow_InputFileLocationProperty struct {
	// `CfnWorkflow.InputFileLocationProperty.EfsFileLocation`.
	EfsFileLocation interface{} `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// `CfnWorkflow.InputFileLocationProperty.S3FileLocation`.
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}
