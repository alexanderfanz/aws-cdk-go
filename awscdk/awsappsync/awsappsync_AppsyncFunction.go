package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// AppSync Functions are local functions that perform certain operations onto a backend data source.
//
// Developers can compose operations (Functions)
// and execute them in sequence with Pipeline Resolvers.
//
// Example:
//   var api graphqlApi
//
//
//   appsyncFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &appsyncFunctionProps{
//   	name: jsii.String("appsync_function"),
//   	api: api,
//   	dataSource: api.addNoneDataSource(jsii.String("none")),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type AppsyncFunction interface {
	awscdk.Resource
	IAppsyncFunction
	// the data source of this AppSync Function.
	// Experimental.
	DataSource() BaseDataSource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// the ARN of the AppSync function.
	// Experimental.
	FunctionArn() *string
	// the ID of the AppSync function.
	// Experimental.
	FunctionId() *string
	// the name of this AppSync Function.
	// Experimental.
	FunctionName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for AppsyncFunction
type jsiiProxy_AppsyncFunction struct {
	internal.Type__awscdkResource
	jsiiProxy_IAppsyncFunction
}

func (j *jsiiProxy_AppsyncFunction) DataSource() BaseDataSource {
	var returns BaseDataSource
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAppsyncFunction(scope constructs.Construct, id *string, props *AppsyncFunctionProps) AppsyncFunction {
	_init_.Initialize()

	if err := validateNewAppsyncFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncFunction{}

	_jsii_.Create(
		"monocdk.aws_appsync.AppsyncFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAppsyncFunction_Override(a AppsyncFunction, scope constructs.Construct, id *string, props *AppsyncFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.AppsyncFunction",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import Appsync Function from arn.
// Experimental.
func AppsyncFunction_FromAppsyncFunctionAttributes(scope constructs.Construct, id *string, attrs *AppsyncFunctionAttributes) IAppsyncFunction {
	_init_.Initialize()

	if err := validateAppsyncFunction_FromAppsyncFunctionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IAppsyncFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.AppsyncFunction",
		"fromAppsyncFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AppsyncFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppsyncFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.AppsyncFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AppsyncFunction_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAppsyncFunction_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.AppsyncFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AppsyncFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunction) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunction) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunction) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AppsyncFunction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunction) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunction) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AppsyncFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
