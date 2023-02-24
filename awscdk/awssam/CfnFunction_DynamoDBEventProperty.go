package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoDBEventProperty := &DynamoDBEventProperty{
//   	StartingPosition: jsii.String("startingPosition"),
//   	Stream: jsii.String("stream"),
//
//   	// the properties below are optional
//   	BatchSize: jsii.Number(123),
//   	BisectBatchOnFunctionError: jsii.Boolean(false),
//   	DestinationConfig: &DestinationConfigProperty{
//   		OnFailure: &DestinationProperty{
//   			Destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   	MaximumBatchingWindowInSeconds: jsii.Number(123),
//   	MaximumRecordAgeInSeconds: jsii.Number(123),
//   	MaximumRetryAttempts: jsii.Number(123),
//   	ParallelizationFactor: jsii.Number(123),
//   }
//
type CfnFunction_DynamoDBEventProperty struct {
	// `CfnFunction.DynamoDBEventProperty.StartingPosition`.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// `CfnFunction.DynamoDBEventProperty.Stream`.
	Stream *string `field:"required" json:"stream" yaml:"stream"`
	// `CfnFunction.DynamoDBEventProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.DynamoDBEventProperty.BisectBatchOnFunctionError`.
	BisectBatchOnFunctionError interface{} `field:"optional" json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// `CfnFunction.DynamoDBEventProperty.DestinationConfig`.
	DestinationConfig interface{} `field:"optional" json:"destinationConfig" yaml:"destinationConfig"`
	// `CfnFunction.DynamoDBEventProperty.Enabled`.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// `CfnFunction.DynamoDBEventProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnFunction.DynamoDBEventProperty.MaximumRecordAgeInSeconds`.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// `CfnFunction.DynamoDBEventProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// `CfnFunction.DynamoDBEventProperty.ParallelizationFactor`.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
}

