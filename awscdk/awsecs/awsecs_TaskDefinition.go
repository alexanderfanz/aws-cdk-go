package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for all task definitions.
//
// Example:
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var vpc vpc
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &fargateServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   })
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//   listener := lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   service.registerLoadBalancerTargets(&ecsTarget{
//   	containerName: jsii.String("web"),
//   	containerPort: jsii.Number(80),
//   	newTargetGroupId: jsii.String("ECS"),
//   	listener: ecs.listenerConfig.applicationListener(listener, &addApplicationTargetsProps{
//   		protocol: elbv2.applicationProtocol_HTTPS,
//   	}),
//   })
//
type TaskDefinition interface {
	awscdk.Resource
	ITaskDefinition
	// The task launch type compatibility requirement.
	Compatibility() Compatibility
	// The container definitions.
	Containers() *[]ContainerDefinition
	// Default container for this task.
	//
	// Load balancers will send traffic to this container. The first
	// essential container that is added to this task will become the default
	// container.
	DefaultContainer() ContainerDefinition
	SetDefaultContainer(val ContainerDefinition)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB() *float64
	// Execution role for this task definition.
	ExecutionRole() awsiam.IRole
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family() *string
	// Public getter method to access list of inference accelerators attached to the instance.
	InferenceAccelerators() *[]*InferenceAccelerator
	// Return true if the task definition can be run on an EC2 cluster.
	IsEc2Compatible() *bool
	// Return true if the task definition can be run on a ECS anywhere cluster.
	IsExternalCompatible() *bool
	// Return true if the task definition can be run on a Fargate cluster.
	IsFargateCompatible() *bool
	// The networking mode to use for the containers in the task.
	NetworkMode() NetworkMode
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// Whether this task definition has at least a container that references a specific JSON field of a secret stored in Secrets Manager.
	ReferencesSecretJsonField() *bool
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The full Amazon Resource Name (ARN) of the task definition.
	TaskDefinitionArn() *string
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole() awsiam.IRole
	// Adds a new container to the task definition.
	AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition
	// Adds the specified extension to the task definition.
	//
	// Extension can be used to apply a packaged modification to
	// a task definition.
	AddExtension(extension ITaskDefinitionExtension)
	// Adds a firelens log router to the task definition.
	AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter
	// Adds an inference accelerator to the task definition.
	AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator)
	// Adds the specified placement constraint to the task definition.
	AddPlacementConstraint(constraint PlacementConstraint)
	// Adds a policy statement to the task execution IAM role.
	AddToExecutionRolePolicy(statement awsiam.PolicyStatement)
	// Adds a policy statement to the task IAM role.
	AddToTaskRolePolicy(statement awsiam.PolicyStatement)
	// Adds a volume to the task definition.
	AddVolume(volume *Volume)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns the container that match the provided containerName.
	FindContainer(containerName *string) ContainerDefinition
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants permissions to run this task definition.
	//
	// This will grant the following permissions:
	//
	//    - ecs:RunTask
	// - iam:PassRole.
	GrantRun(grantee awsiam.IGrantable) awsiam.Grant
	// Creates the task execution IAM role if it doesn't already exist.
	ObtainExecutionRole() awsiam.IRole
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TaskDefinition
type jsiiProxy_TaskDefinition struct {
	internal.Type__awscdkResource
	jsiiProxy_ITaskDefinition
}

func (j *jsiiProxy_TaskDefinition) Compatibility() Compatibility {
	var returns Compatibility
	_jsii_.Get(
		j,
		"compatibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Containers() *[]ContainerDefinition {
	var returns *[]ContainerDefinition
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) DefaultContainer() ContainerDefinition {
	var returns ContainerDefinition
	_jsii_.Get(
		j,
		"defaultContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) EphemeralStorageGiB() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageGiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) InferenceAccelerators() *[]*InferenceAccelerator {
	var returns *[]*InferenceAccelerator
	_jsii_.Get(
		j,
		"inferenceAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) IsEc2Compatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEc2Compatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) IsExternalCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isExternalCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) IsFargateCompatible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFargateCompatible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) NetworkMode() NetworkMode {
	var returns NetworkMode
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) ReferencesSecretJsonField() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"referencesSecretJsonField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskDefinition) TaskRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"taskRole",
		&returns,
	)
	return returns
}


// Constructs a new instance of the TaskDefinition class.
func NewTaskDefinition(scope constructs.Construct, id *string, props *TaskDefinitionProps) TaskDefinition {
	_init_.Initialize()

	j := jsiiProxy_TaskDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the TaskDefinition class.
func NewTaskDefinition_Override(t TaskDefinition, scope constructs.Construct, id *string, props *TaskDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TaskDefinition) SetDefaultContainer(val ContainerDefinition) {
	_jsii_.Set(
		j,
		"defaultContainer",
		val,
	)
}

// Imports a task definition from the specified task definition ARN.
//
// The task will have a compatibility of EC2+Fargate.
func TaskDefinition_FromTaskDefinitionArn(scope constructs.Construct, id *string, taskDefinitionArn *string) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"fromTaskDefinitionArn",
		[]interface{}{scope, id, taskDefinitionArn},
		&returns,
	)

	return returns
}

// Create a task definition from a task definition reference.
func TaskDefinition_FromTaskDefinitionAttributes(scope constructs.Construct, id *string, attrs *TaskDefinitionAttributes) ITaskDefinition {
	_init_.Initialize()

	var returns ITaskDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"fromTaskDefinitionAttributes",
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
func TaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func TaskDefinition_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func TaskDefinition_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TaskDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) AddContainer(id *string, props *ContainerDefinitionOptions) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		t,
		"addContainer",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) AddExtension(extension ITaskDefinitionExtension) {
	_jsii_.InvokeVoid(
		t,
		"addExtension",
		[]interface{}{extension},
	)
}

func (t *jsiiProxy_TaskDefinition) AddFirelensLogRouter(id *string, props *FirelensLogRouterDefinitionOptions) FirelensLogRouter {
	var returns FirelensLogRouter

	_jsii_.Invoke(
		t,
		"addFirelensLogRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) AddInferenceAccelerator(inferenceAccelerator *InferenceAccelerator) {
	_jsii_.InvokeVoid(
		t,
		"addInferenceAccelerator",
		[]interface{}{inferenceAccelerator},
	)
}

func (t *jsiiProxy_TaskDefinition) AddPlacementConstraint(constraint PlacementConstraint) {
	_jsii_.InvokeVoid(
		t,
		"addPlacementConstraint",
		[]interface{}{constraint},
	)
}

func (t *jsiiProxy_TaskDefinition) AddToExecutionRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		t,
		"addToExecutionRolePolicy",
		[]interface{}{statement},
	)
}

func (t *jsiiProxy_TaskDefinition) AddToTaskRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		t,
		"addToTaskRolePolicy",
		[]interface{}{statement},
	)
}

func (t *jsiiProxy_TaskDefinition) AddVolume(volume *Volume) {
	_jsii_.InvokeVoid(
		t,
		"addVolume",
		[]interface{}{volume},
	)
}

func (t *jsiiProxy_TaskDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TaskDefinition) FindContainer(containerName *string) ContainerDefinition {
	var returns ContainerDefinition

	_jsii_.Invoke(
		t,
		"findContainer",
		[]interface{}{containerName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) GrantRun(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantRun",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) ObtainExecutionRole() awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		t,
		"obtainExecutionRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
