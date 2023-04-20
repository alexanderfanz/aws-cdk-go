package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for defining a State Machine.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   inputArtifact := codepipeline.NewArtifact()
//   startState := stepfunctions.NewPass(this, jsii.String("StartState"))
//   simpleStateMachine := stepfunctions.NewStateMachine(this, jsii.String("SimpleStateMachine"), &StateMachineProps{
//   	Definition: startState,
//   })
//   stepFunctionAction := codepipeline_actions.NewStepFunctionInvokeAction(&StepFunctionsInvokeActionProps{
//   	ActionName: jsii.String("Invoke"),
//   	StateMachine: simpleStateMachine,
//   	StateMachineInput: codepipeline_actions.StateMachineInput_FilePath(inputArtifact.AtPath(jsii.String("assets/input.json"))),
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("StepFunctions"),
//   	Actions: []iAction{
//   		stepFunctionAction,
//   	},
//   })
//
// Experimental.
type StateMachineProps struct {
	// Definition for this state machine.
	// Experimental.
	Definition IChainable `field:"required" json:"definition" yaml:"definition"`
	// Defines what execution history events are logged and where they are logged.
	// Experimental.
	Logs *LogOptions `field:"optional" json:"logs" yaml:"logs"`
	// The execution role for the state machine service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// A name for the state machine.
	// Experimental.
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
	// Type of the state machine.
	// Experimental.
	StateMachineType StateMachineType `field:"optional" json:"stateMachineType" yaml:"stateMachineType"`
	// Maximum run time for this state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Specifies whether Amazon X-Ray tracing is enabled for this state machine.
	// Experimental.
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
}

