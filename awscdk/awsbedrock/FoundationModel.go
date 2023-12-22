package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Bedrock base foundation model.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_ANTHROPIC_CLAUDE_V2())
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/models-supported.html
//
type FoundationModel interface {
	IModel
	// The foundation model ARN.
	ModelArn() *string
	// The foundation model ID.
	//
	// Example:
	//   "amazon.titan-text-express-v1"
	//
	ModelId() *string
}

// The jsii proxy struct for FoundationModel
type jsiiProxy_FoundationModel struct {
	jsiiProxy_IModel
}

func (j *jsiiProxy_FoundationModel) ModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FoundationModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}


// Construct a Bedrock base foundation model given the model identifier.
//
// Returns: A Bedrock base foundation model.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/model-ids-arns.html
//
func FoundationModel_FromFoundationModelId(scope constructs.Construct, _id *string, foundationModelId FoundationModelIdentifier) FoundationModel {
	_init_.Initialize()

	if err := validateFoundationModel_FromFoundationModelIdParameters(scope, _id, foundationModelId); err != nil {
		panic(err)
	}
	var returns FoundationModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.FoundationModel",
		"fromFoundationModelId",
		[]interface{}{scope, _id, foundationModelId},
		&returns,
	)

	return returns
}
