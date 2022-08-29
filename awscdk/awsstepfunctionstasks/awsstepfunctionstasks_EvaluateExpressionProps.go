package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for EvaluateExpression.
//
// Example:
//   convertToSeconds := tasks.NewEvaluateExpression(this, jsii.String("Convert to seconds"), &evaluateExpressionProps{
//   	expression: jsii.String("$.waitMilliseconds / 1000"),
//   	resultPath: jsii.String("$.waitSeconds"),
//   })
//
//   createMessage := tasks.NewEvaluateExpression(this, jsii.String("Create message"), &evaluateExpressionProps{
//   	// Note: this is a string inside a string.
//   	expression: jsii.String("`Now waiting ${$.waitSeconds} seconds...`"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	resultPath: jsii.String("$.message"),
//   })
//
//   publishMessage := tasks.NewSnsPublish(this, jsii.String("Publish message"), &snsPublishProps{
//   	topic: sns.NewTopic(this, jsii.String("cool-topic")),
//   	message: sfn.taskInput.fromJsonPathAt(jsii.String("$.message")),
//   	resultPath: jsii.String("$.sns"),
//   })
//
//   wait := sfn.NewWait(this, jsii.String("Wait"), &waitProps{
//   	time: sfn.waitTime.secondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: convertToSeconds.next(createMessage).next(publishMessage).next(wait),
//   })
//
type EvaluateExpressionProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The expression to evaluate. The expression may contain state paths.
	//
	// Example value: `'$.a + $.b'`
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The runtime language to use to evaluate the expression.
	Runtime awslambda.Runtime `field:"optional" json:"runtime" yaml:"runtime"`
}
