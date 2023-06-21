package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a GameLift FleetIQ game server group for managing game hosting on a collection of Amazon EC2 instances for game hosting.
//
// This operation creates the game server group, creates an Auto Scaling group in your AWS account, and establishes a link between the two groups.
// You can view the status of your game server groups in the GameLift console.
// Game server group metrics and events are emitted to Amazon CloudWatch.
// Before creating a new game server group, you must have the following:
//  - An Amazon EC2 launch template that specifies how to launch Amazon EC2 instances with your game server build.
//  - An IAM role that extends limited access to your AWS account to allow GameLift FleetIQ to create and interact with the Auto Scaling group.
//
// To create a new game server group, specify a unique group name, IAM role and Amazon EC2 launch template, and provide a list of instance types that can be used in the group.
// You must also set initial maximum and minimum limits on the group's instance count.
// You can optionally set an Auto Scaling policy with target tracking based on a GameLift FleetIQ metric.
//
// Once the game server group and corresponding Auto Scaling group are created, you have full access to change the Auto Scaling group's configuration as needed.
// Several properties that are set when creating a game server group, including maximum/minimum size and auto-scaling policy settings, must be updated directly in the Auto Scaling group.
// Keep in mind that some Auto Scaling group properties are periodically updated by GameLift FleetIQ as part of its balancing activities to optimize for availability and cost.
//
// Example:
//   var launchTemplate iLaunchTemplate
//   var vpc iVpc
//
//
//   gamelift.NewGameServerGroup(this, jsii.String("GameServerGroup"), &GameServerGroupProps{
//   	GameServerGroupName: jsii.String("sample-gameservergroup-name"),
//   	InstanceDefinitions: []instanceDefinition{
//   		&instanceDefinition{
//   			InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C5, ec2.InstanceSize_LARGE),
//   		},
//   		&instanceDefinition{
//   			InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_C4, ec2.InstanceSize_LARGE),
//   		},
//   	},
//   	LaunchTemplate: launchTemplate,
//   	Vpc: vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
//
// Experimental.
type GameServerGroup interface {
	GameServerGroupBase
	// The ARN of the generated AutoScaling group.
	// Experimental.
	AutoScalingGroupArn() *string
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
	// The ARN of the game server group.
	// Experimental.
	GameServerGroupArn() *string
	// The name of the game server group.
	// Experimental.
	GameServerGroupName() *string
	// The principal this GameLift game server group is using.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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
	// The IAM role that allows Amazon GameLift to access your Amazon EC2 Auto Scaling groups.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The VPC network to place the game server group in.
	// Experimental.
	Vpc() awsec2.IVpc
	// The game server group's subnets.
	// Experimental.
	VpcSubnets() *awsec2.SubnetSelection
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
	// Grant the `grantee` identity permissions to perform `actions`.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Return the given named metric for this fleet.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Experimental.
	ParseAutoScalingPolicy(props *GameServerGroupProps) *awsgamelift.CfnGameServerGroup_AutoScalingPolicyProperty
	// Experimental.
	ParseInstanceDefinitions(props *GameServerGroupProps) *[]*awsgamelift.CfnGameServerGroup_InstanceDefinitionProperty
	// Experimental.
	ParseLaunchTemplate(props *GameServerGroupProps) *awsgamelift.CfnGameServerGroup_LaunchTemplateProperty
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GameServerGroup
type jsiiProxy_GameServerGroup struct {
	jsiiProxy_GameServerGroupBase
}

func (j *jsiiProxy_GameServerGroup) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) GameServerGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) GameServerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GameServerGroup) VpcSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}


// Experimental.
func NewGameServerGroup(scope constructs.Construct, id *string, props *GameServerGroupProps) GameServerGroup {
	_init_.Initialize()

	if err := validateNewGameServerGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GameServerGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGameServerGroup_Override(g GameServerGroup, scope constructs.Construct, id *string, props *GameServerGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		[]interface{}{scope, id, props},
		g,
	)
}

// Import an existing game server group from its attributes.
// Experimental.
func GameServerGroup_FromGameServerGroupAttributes(scope constructs.Construct, id *string, attrs *GameServerGroupAttributes) IGameServerGroup {
	_init_.Initialize()

	if err := validateGameServerGroup_FromGameServerGroupAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IGameServerGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		"fromGameServerGroupAttributes",
		[]interface{}{scope, id, attrs},
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
func GameServerGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGameServerGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func GameServerGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGameServerGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GameServerGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGameServerGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.GameServerGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := g.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GameServerGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := g.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := g.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := g.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) ParseAutoScalingPolicy(props *GameServerGroupProps) *awsgamelift.CfnGameServerGroup_AutoScalingPolicyProperty {
	if err := g.validateParseAutoScalingPolicyParameters(props); err != nil {
		panic(err)
	}
	var returns *awsgamelift.CfnGameServerGroup_AutoScalingPolicyProperty

	_jsii_.Invoke(
		g,
		"parseAutoScalingPolicy",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) ParseInstanceDefinitions(props *GameServerGroupProps) *[]*awsgamelift.CfnGameServerGroup_InstanceDefinitionProperty {
	if err := g.validateParseInstanceDefinitionsParameters(props); err != nil {
		panic(err)
	}
	var returns *[]*awsgamelift.CfnGameServerGroup_InstanceDefinitionProperty

	_jsii_.Invoke(
		g,
		"parseInstanceDefinitions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) ParseLaunchTemplate(props *GameServerGroupProps) *awsgamelift.CfnGameServerGroup_LaunchTemplateProperty {
	if err := g.validateParseLaunchTemplateParameters(props); err != nil {
		panic(err)
	}
	var returns *awsgamelift.CfnGameServerGroup_LaunchTemplateProperty

	_jsii_.Invoke(
		g,
		"parseLaunchTemplate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GameServerGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

