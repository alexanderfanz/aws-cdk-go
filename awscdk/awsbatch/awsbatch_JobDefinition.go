package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Batch Job Definition.
//
// Defines a batch job definition to execute a specific batch job.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repo := ecr.repository.fromRepositoryName(this, jsii.String("batch-job-repo"), jsii.String("todo-list"))
//
//   batch.NewJobDefinition(this, jsii.String("batch-job-def-from-ecr"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.NewEcrImage(repo, jsii.String("latest")),
//   	},
//   })
//
// Experimental.
type JobDefinition interface {
	awscdk.Resource
	IJobDefinition
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
	// The ARN of this batch job definition.
	// Experimental.
	JobDefinitionArn() *string
	// The name of the batch job definition.
	// Experimental.
	JobDefinitionName() *string
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

// The jsii proxy struct for JobDefinition
type jsiiProxy_JobDefinition struct {
	internal.Type__awscdkResource
	jsiiProxy_IJobDefinition
}

func (j *jsiiProxy_JobDefinition) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewJobDefinition(scope constructs.Construct, id *string, props *JobDefinitionProps) JobDefinition {
	_init_.Initialize()

	if err := validateNewJobDefinitionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobDefinition{}

	_jsii_.Create(
		"monocdk.aws_batch.JobDefinition",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJobDefinition_Override(j JobDefinition, scope constructs.Construct, id *string, props *JobDefinitionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_batch.JobDefinition",
		[]interface{}{scope, id, props},
		j,
	)
}

// Imports an existing batch job definition by its amazon resource name.
// Experimental.
func JobDefinition_FromJobDefinitionArn(scope constructs.Construct, id *string, jobDefinitionArn *string) IJobDefinition {
	_init_.Initialize()

	if err := validateJobDefinition_FromJobDefinitionArnParameters(scope, id, jobDefinitionArn); err != nil {
		panic(err)
	}
	var returns IJobDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"fromJobDefinitionArn",
		[]interface{}{scope, id, jobDefinitionArn},
		&returns,
	)

	return returns
}

// Imports an existing batch job definition by its name.
//
// If name is specified without a revision then the latest active revision is used.
// Experimental.
func JobDefinition_FromJobDefinitionName(scope constructs.Construct, id *string, jobDefinitionName *string) IJobDefinition {
	_init_.Initialize()

	if err := validateJobDefinition_FromJobDefinitionNameParameters(scope, id, jobDefinitionName); err != nil {
		panic(err)
	}
	var returns IJobDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"fromJobDefinitionName",
		[]interface{}{scope, id, jobDefinitionName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func JobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJobDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func JobDefinition_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJobDefinition_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_batch.JobDefinition",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := j.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_JobDefinition) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobDefinition) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := j.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobDefinition) GetResourceNameAttribute(nameAttr *string) *string {
	if err := j.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobDefinition) OnPrepare() {
	_jsii_.InvokeVoid(
		j,
		"onPrepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobDefinition) OnSynthesize(session constructs.ISynthesisSession) {
	if err := j.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (j *jsiiProxy_JobDefinition) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobDefinition) Prepare() {
	_jsii_.InvokeVoid(
		j,
		"prepare",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobDefinition) Synthesize(session awscdk.ISynthesisSession) {
	if err := j.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		[]interface{}{session},
	)
}

func (j *jsiiProxy_JobDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobDefinition) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

