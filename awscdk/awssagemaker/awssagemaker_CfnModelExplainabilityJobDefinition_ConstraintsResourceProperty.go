package awssagemaker


// Input object for the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   constraintsResourceProperty := &constraintsResourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//   }
//
type CfnModelExplainabilityJobDefinition_ConstraintsResourceProperty struct {
	// The Amazon S3 URI for the constraints resource.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

