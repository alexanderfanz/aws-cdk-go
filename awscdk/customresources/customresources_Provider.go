package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an AWS CloudFormation custom resource provider.
//
// Example:
//   var onEvent function
//   var isComplete function
//   var myRole role
//
//   myProvider := cr.NewProvider(this, jsii.String("MyProvider"), &providerProps{
//   	onEventHandler: onEvent,
//   	isCompleteHandler: isComplete,
//   	logRetention: logs.retentionDays_ONE_DAY,
//   	role: myRole,
//   	providerFunctionName: jsii.String("the-lambda-name"),
//   })
//
type Provider interface {
	constructs.Construct
	// The user-defined AWS Lambda function which is invoked asynchronously in order to determine if the operation is complete.
	IsCompleteHandler() awslambda.IFunction
	// The tree node.
	Node() constructs.Node
	// The user-defined AWS Lambda function which is invoked for all resource lifecycle operations (CREATE/UPDATE/DELETE).
	OnEventHandler() awslambda.IFunction
	// The service token to use in order to define custom resources that are backed by this provider.
	ServiceToken() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Provider
type jsiiProxy_Provider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Provider) IsCompleteHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"isCompleteHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) OnEventHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"onEventHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


func NewProvider(scope constructs.Construct, id *string, props *ProviderProps) Provider {
	_init_.Initialize()

	j := jsiiProxy_Provider{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.Provider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewProvider_Override(p Provider, scope constructs.Construct, id *string, props *ProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.Provider",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Provider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.Provider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
