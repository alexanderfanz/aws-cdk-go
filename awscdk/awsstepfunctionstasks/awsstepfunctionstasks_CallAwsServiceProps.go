package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for calling an AWS service's API action from your state machine.
//
// Example:
//   var myBucket bucket
//
//   getObject := tasks.NewCallAwsService(this, jsii.String("GetObject"), &CallAwsServiceProps{
//   	Service: jsii.String("s3"),
//   	Action: jsii.String("getObject"),
//   	Parameters: map[string]interface{}{
//   		"Bucket": myBucket.bucketName,
//   		"Key": sfn.JsonPath_stringAt(jsii.String("$.key")),
//   	},
//   	IamResources: []*string{
//   		myBucket.ArnForObjects(jsii.String("*")),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html
//
// Experimental.
type CallAwsServiceProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The API action to call.
	//
	// Use camelCase.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The resources for the IAM statement that will be added to the state machine role's policy to allow the state machine to make the API call.
	//
	// By default the action for this IAM statement will be `service:action`.
	// Experimental.
	IamResources *[]*string `field:"required" json:"iamResources" yaml:"iamResources"`
	// The AWS service to call.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html
	//
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The action for the IAM statement that will be added to the state machine role's policy to allow the state machine to make the API call.
	//
	// Use in the case where the IAM action name does not match with the
	// API service/action name, e.g. `s3:ListBuckets` requires `s3:ListAllMyBuckets`.
	// Experimental.
	IamAction *string `field:"optional" json:"iamAction" yaml:"iamAction"`
	// Parameters for the API action call.
	//
	// Use PascalCase for the parameter names.
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

