package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sfn.NewStateMachine(stack, jsii.String("StateMachineFromString"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromString(jsii.String("{\"StartAt\":\"Pass\",\"States\":{\"Pass\":{\"Type\":\"Pass\",\"End\":true}}}")),
//   })
//
//   sfn.NewStateMachine(stack, jsii.String("StateMachineFromFile"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromFile(jsii.String("./asl.json")),
//   })
//
type DefinitionBody interface {
	Bind(scope constructs.Construct, sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps) *DefinitionConfig
}

// The jsii proxy struct for DefinitionBody
type jsiiProxy_DefinitionBody struct {
	_ byte // padding
}

func NewDefinitionBody_Override(d DefinitionBody) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.DefinitionBody",
		nil, // no parameters
		d,
	)
}

func DefinitionBody_FromChainable(chainable IChainable) DefinitionBody {
	_init_.Initialize()

	if err := validateDefinitionBody_FromChainableParameters(chainable); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DefinitionBody",
		"fromChainable",
		[]interface{}{chainable},
		&returns,
	)

	return returns
}

func DefinitionBody_FromFile(path *string, options *awss3assets.AssetOptions) DefinitionBody {
	_init_.Initialize()

	if err := validateDefinitionBody_FromFileParameters(path, options); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DefinitionBody",
		"fromFile",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

func DefinitionBody_FromString(definition *string) DefinitionBody {
	_init_.Initialize()

	if err := validateDefinitionBody_FromStringParameters(definition); err != nil {
		panic(err)
	}
	var returns DefinitionBody

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DefinitionBody",
		"fromString",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefinitionBody) Bind(scope constructs.Construct, sfnPrincipal awsiam.IPrincipal, sfnProps *StateMachineProps) *DefinitionConfig {
	if err := d.validateBindParameters(scope, sfnPrincipal, sfnProps); err != nil {
		panic(err)
	}
	var returns *DefinitionConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, sfnPrincipal, sfnProps},
		&returns,
	)

	return returns
}

