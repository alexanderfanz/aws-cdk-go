package awslambdapython

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdapython/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A lambda layer version.
//
// Example:
//   lambda.NewPythonLayerVersion(this, jsii.String("MyLayer"), &pythonLayerVersionProps{
//   	entry: jsii.String("/path/to/my/layer"),
//   })
//
// Experimental.
type PythonLayerVersion interface {
	awslambda.LayerVersion
	// The runtimes compatible with this Layer.
	// Experimental.
	CompatibleRuntimes() *[]awslambda.Runtime
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
	// The ARN of the Lambda Layer version that this Layer defines.
	// Experimental.
	LayerVersionArn() *string
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
	// Add permission for this layer version to specific entities.
	//
	// Usage within
	// the same account where the layer is defined is always allowed and does not
	// require calling this method. Note that the principal that creates the
	// Lambda function using the layer (for example, a CloudFormation changeset
	// execution role) also needs to have the ``lambda:GetLayerVersion``
	// permission on the layer version.
	// Experimental.
	AddPermission(id *string, permission *awslambda.LayerVersionPermission)
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

// The jsii proxy struct for PythonLayerVersion
type jsiiProxy_PythonLayerVersion struct {
	internal.Type__awslambdaLayerVersion
}

func (j *jsiiProxy_PythonLayerVersion) CompatibleRuntimes() *[]awslambda.Runtime {
	var returns *[]awslambda.Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonLayerVersion(scope awscdk.Construct, id *string, props *PythonLayerVersionProps) PythonLayerVersion {
	_init_.Initialize()

	if err := validateNewPythonLayerVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PythonLayerVersion{}

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonLayerVersion_Override(p PythonLayerVersion, scope awscdk.Construct, id *string, props *PythonLayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a layer version by ARN.
//
// Assumes it is compatible with all Lambda runtimes.
// Experimental.
func PythonLayerVersion_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) awslambda.ILayerVersion {
	_init_.Initialize()

	if err := validatePythonLayerVersion_FromLayerVersionArnParameters(scope, id, layerVersionArn); err != nil {
		panic(err)
	}
	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"fromLayerVersionArn",
		[]interface{}{scope, id, layerVersionArn},
		&returns,
	)

	return returns
}

// Imports a Layer that has been defined externally.
// Experimental.
func PythonLayerVersion_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) awslambda.ILayerVersion {
	_init_.Initialize()

	if err := validatePythonLayerVersion_FromLayerVersionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"fromLayerVersionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func PythonLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePythonLayerVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func PythonLayerVersion_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePythonLayerVersion_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLayerVersion) AddPermission(id *string, permission *awslambda.LayerVersionPermission) {
	if err := p.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (p *jsiiProxy_PythonLayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PythonLayerVersion) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLayerVersion) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLayerVersion) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLayerVersion) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonLayerVersion) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PythonLayerVersion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLayerVersion) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonLayerVersion) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PythonLayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLayerVersion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

