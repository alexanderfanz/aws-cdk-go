package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Options to integrate with various StepFunction API.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
//   	StateMachineType: stepfunctions.StateMachineType_EXPRESS,
//   	Definition: stepfunctions.Chain_Start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &RestApiProps{
//   	RestApiName: jsii.String("MyApi"),
//   })
//   api.Root.AddMethod(jsii.String("GET"), apigateway.StepFunctionsIntegration_StartExecution(stateMachine))
//
// Experimental.
type StepFunctionsIntegration interface {
}

// The jsii proxy struct for StepFunctionsIntegration
type jsiiProxy_StepFunctionsIntegration struct {
	_ byte // padding
}

// Experimental.
func NewStepFunctionsIntegration() StepFunctionsIntegration {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigateway.StepFunctionsIntegration",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionsIntegration_Override(s StepFunctionsIntegration) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.StepFunctionsIntegration",
		nil, // no parameters
		s,
	)
}

// Integrates a Synchronous Express State Machine from AWS Step Functions to an API Gateway method.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
//   	StateMachineType: stepfunctions.StateMachineType_EXPRESS,
//   	Definition: stepfunctions.Chain_Start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &RestApiProps{
//   	RestApiName: jsii.String("MyApi"),
//   })
//   api.Root.AddMethod(jsii.String("GET"), apigateway.StepFunctionsIntegration_StartExecution(stateMachine))
//
// Experimental.
func StepFunctionsIntegration_StartExecution(stateMachine awsstepfunctions.IStateMachine, options *StepFunctionsExecutionIntegrationOptions) AwsIntegration {
	_init_.Initialize()

	if err := validateStepFunctionsIntegration_StartExecutionParameters(stateMachine, options); err != nil {
		panic(err)
	}
	var returns AwsIntegration

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.StepFunctionsIntegration",
		"startExecution",
		[]interface{}{stateMachine, options},
		&returns,
	)

	return returns
}

