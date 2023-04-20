package awsstepfunctions


// Properties for defining a Choice state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   choiceProps := &ChoiceProps{
//   	Comment: jsii.String("comment"),
//   	InputPath: jsii.String("inputPath"),
//   	OutputPath: jsii.String("outputPath"),
//   }
//
// Experimental.
type ChoiceProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
}

