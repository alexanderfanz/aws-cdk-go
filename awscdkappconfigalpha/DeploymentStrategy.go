package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS AppConfig deployment strategy.
//
// Example:
//   appconfig.NewDeploymentStrategy(this, jsii.String("MyDeploymentStrategy"), &DeploymentStrategyProps{
//   	RolloutStrategy: appconfig.RolloutStrategy_Linear(&RolloutStrategyProps{
//   		GrowthFactor: jsii.Number(20),
//   		DeploymentDuration: awscdk.Duration_Minutes(jsii.Number(30)),
//   		FinalBakeTime: awscdk.Duration_*Minutes(jsii.Number(30)),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-creating-deployment-strategy.html
//
// Experimental.
type DeploymentStrategy interface {
	awscdk.Resource
	IDeploymentStrategy
	// The deployment duration in minutes of the deployment strategy.
	// Experimental.
	DeploymentDurationInMinutes() *float64
	// The Amazon Resource Name (ARN) of the deployment strategy.
	// Experimental.
	DeploymentStrategyArn() *string
	// The ID of the deployment strategy.
	// Experimental.
	DeploymentStrategyId() *string
	// The description of the deployment strategy.
	// Experimental.
	Description() *string
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
	// The final bake time in minutes of the deployment strategy.
	// Experimental.
	FinalBakeTimeInMinutes() *float64
	// The growth factor of the deployment strategy.
	// Experimental.
	GrowthFactor() *float64
	// The growth type of the deployment strategy.
	// Experimental.
	GrowthType() GrowthType
	// The name of the deployment strategy.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeploymentStrategy
type jsiiProxy_DeploymentStrategy struct {
	internal.Type__awscdkResource
	jsiiProxy_IDeploymentStrategy
}

func (j *jsiiProxy_DeploymentStrategy) DeploymentDurationInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentDurationInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) DeploymentStrategyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStrategyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) DeploymentStrategyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStrategyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) FinalBakeTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalBakeTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) GrowthFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"growthFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) GrowthType() GrowthType {
	var returns GrowthType
	_jsii_.Get(
		j,
		"growthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentStrategy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDeploymentStrategy(scope constructs.Construct, id *string, props *DeploymentStrategyProps) DeploymentStrategy {
	_init_.Initialize()

	if err := validateNewDeploymentStrategyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentStrategy{}

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDeploymentStrategy_Override(d DeploymentStrategy, scope constructs.Construct, id *string, props *DeploymentStrategyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		[]interface{}{scope, id, props},
		d,
	)
}

// Imports a deployment strategy into the CDK using its Amazon Resource Name (ARN).
// Experimental.
func DeploymentStrategy_FromDeploymentStrategyArn(scope constructs.Construct, id *string, deploymentStrategyArn *string) IDeploymentStrategy {
	_init_.Initialize()

	if err := validateDeploymentStrategy_FromDeploymentStrategyArnParameters(scope, id, deploymentStrategyArn); err != nil {
		panic(err)
	}
	var returns IDeploymentStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		"fromDeploymentStrategyArn",
		[]interface{}{scope, id, deploymentStrategyArn},
		&returns,
	)

	return returns
}

// Imports a deployment strategy into the CDK using its ID.
// Experimental.
func DeploymentStrategy_FromDeploymentStrategyId(scope constructs.Construct, id *string, deploymentStrategyId *string) IDeploymentStrategy {
	_init_.Initialize()

	if err := validateDeploymentStrategy_FromDeploymentStrategyIdParameters(scope, id, deploymentStrategyId); err != nil {
		panic(err)
	}
	var returns IDeploymentStrategy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		"fromDeploymentStrategyId",
		[]interface{}{scope, id, deploymentStrategyId},
		&returns,
	)

	return returns
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
// Experimental.
func DeploymentStrategy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeploymentStrategy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func DeploymentStrategy_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDeploymentStrategy_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DeploymentStrategy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDeploymentStrategy_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.DeploymentStrategy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentStrategy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DeploymentStrategy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentStrategy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentStrategy) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentStrategy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
