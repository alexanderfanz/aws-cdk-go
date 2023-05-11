package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisEventProperty := &KinesisEventProperty{
//   	StartingPosition: jsii.String("startingPosition"),
//   	Stream: jsii.String("stream"),
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	Enabled: jsii.Boolean(false),
//   	FunctionResponseTypes: []*string{
//   		jsii.String("functionResponseTypes"),
//   	},
//   }
//
type CfnFunction_KinesisEventProperty struct {
	// `CfnFunction.KinesisEventProperty.StartingPosition`.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// `CfnFunction.KinesisEventProperty.Stream`.
	Stream *string `field:"required" json:"stream" yaml:"stream"`
	// `CfnFunction.KinesisEventProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.KinesisEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnFunction.KinesisEventProperty.FunctionResponseTypes`.
	FunctionResponseTypes *[]*string `field:"optional" json:"functionResponseTypes" yaml:"functionResponseTypes"`
}
