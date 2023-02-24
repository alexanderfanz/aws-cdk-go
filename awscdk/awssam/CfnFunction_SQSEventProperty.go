package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sQSEventProperty := &SQSEventProperty{
//   	Queue: jsii.String("queue"),
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   }
//
type CfnFunction_SQSEventProperty struct {
	// `CfnFunction.SQSEventProperty.Queue`.
	Queue *string `field:"required" json:"queue" yaml:"queue"`
	// `CfnFunction.SQSEventProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.SQSEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

