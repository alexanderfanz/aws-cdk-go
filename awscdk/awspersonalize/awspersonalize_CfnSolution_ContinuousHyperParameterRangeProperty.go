package awspersonalize


// Provides the name and range of a continuous hyperparameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   continuousHyperParameterRangeProperty := &continuousHyperParameterRangeProperty{
//   	maxValue: jsii.Number(123),
//   	minValue: jsii.Number(123),
//   	name: jsii.String("name"),
//   }
//
type CfnSolution_ContinuousHyperParameterRangeProperty struct {
	// The maximum allowable value for the hyperparameter.
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum allowable value for the hyperparameter.
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The name of the hyperparameter.
	Name *string `field:"optional" json:"name" yaml:"name"`
}
