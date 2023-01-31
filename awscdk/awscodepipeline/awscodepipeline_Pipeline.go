package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// An AWS CodePipeline pipeline with its associated IAM role and S3 bucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // create a pipeline
//   import codecommit "github.com/aws-samples/dummy/awscdkawscodecommit"
//
//   // add a source action to the stage
//   var repo repository
//   var sourceArtifact artifact
//
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("Pipeline"))
//
//   // add a stage
//   sourceStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   })
//   sourceStage.addAction(codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	output: sourceArtifact,
//   	repository: repo,
//   }))
//
// Experimental.
type Pipeline interface {
	awscdk.Resource
	IPipeline
	// Bucket used to store output artifacts.
	// Experimental.
	ArtifactBucket() awss3.IBucket
	// Returns all of the {@link CrossRegionSupportStack}s that were generated automatically when dealing with Actions that reside in a different region than the Pipeline itself.
	// Experimental.
	CrossRegionSupport() *map[string]*CrossRegionSupport
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
	// ARN of this pipeline.
	// Experimental.
	PipelineArn() *string
	// The name of the pipeline.
	// Experimental.
	PipelineName() *string
	// The version of the pipeline.
	// Experimental.
	PipelineVersion() *string
	// The IAM role AWS CodePipeline will use to perform actions or assume roles for actions with a more specific IAM role.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Get the number of Stages in this Pipeline.
	// Experimental.
	StageCount() *float64
	// Returns the stages that comprise the pipeline.
	//
	// **Note**: the returned array is a defensive copy,
	// so adding elements to it has no effect.
	// Instead, use the {@link addStage} method if you want to add more stages
	// to the pipeline.
	// Experimental.
	Stages() *[]IStage
	// Creates a new Stage, and adds it to this Pipeline.
	//
	// Returns: the newly created Stage.
	// Experimental.
	AddStage(props *StageOptions) IStage
	// Adds a statement to the pipeline role.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
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
	// Returns a source configuration for notification rule.
	// Experimental.
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
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
	// Defines a CodeStar notification rule triggered when the pipeline events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnExecutionStateChange`, `notifyOnAnyStageStateChange`,
	// `notifyOnAnyActionStateChange` and `notifyOnAnyManualApprovalStateChange`
	// to define rules for these specific event emitted.
	// Experimental.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Action execution" events emitted from this pipeline.
	// Experimental.
	NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Manual approval" events emitted from this pipeline.
	// Experimental.
	NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Stage execution" events emitted from this pipeline.
	// Experimental.
	NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Define an notification rule triggered by the set of the "Pipeline execution" events emitted from this pipeline.
	// Experimental.
	NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines an event rule triggered by this CodePipeline.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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
	// Defines an event rule triggered by the "CodePipeline Pipeline Execution State Change" event emitted from this pipeline.
	// Experimental.
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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
	// Access one of the pipeline's stages by stage name.
	// Experimental.
	Stage(stageName *string) IStage
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the pipeline structure.
	//
	// Validation happens according to the rules documented at
	//
	// https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#pipeline-requirements
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for Pipeline
type jsiiProxy_Pipeline struct {
	internal.Type__awscdkResource
	jsiiProxy_IPipeline
}

func (j *jsiiProxy_Pipeline) ArtifactBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"artifactBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) CrossRegionSupport() *map[string]*CrossRegionSupport {
	var returns *map[string]*CrossRegionSupport
	_jsii_.Get(
		j,
		"crossRegionSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) PipelineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) StageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pipeline) Stages() *[]IStage {
	var returns *[]IStage
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}


// Experimental.
func NewPipeline(scope constructs.Construct, id *string, props *PipelineProps) Pipeline {
	_init_.Initialize()

	if err := validateNewPipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pipeline{}

	_jsii_.Create(
		"monocdk.aws_codepipeline.Pipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPipeline_Override(p Pipeline, scope constructs.Construct, id *string, props *PipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codepipeline.Pipeline",
		[]interface{}{scope, id, props},
		p,
	)
}

// Import a pipeline into this app.
// Experimental.
func Pipeline_FromPipelineArn(scope constructs.Construct, id *string, pipelineArn *string) IPipeline {
	_init_.Initialize()

	if err := validatePipeline_FromPipelineArnParameters(scope, id, pipelineArn); err != nil {
		panic(err)
	}
	var returns IPipeline

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Pipeline",
		"fromPipelineArn",
		[]interface{}{scope, id, pipelineArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Pipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Pipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Pipeline_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePipeline_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codepipeline.Pipeline",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) AddStage(props *StageOptions) IStage {
	if err := p.validateAddStageParameters(props); err != nil {
		panic(err)
	}
	var returns IStage

	_jsii_.Invoke(
		p,
		"addStage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := p.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (p *jsiiProxy_Pipeline) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_Pipeline) BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	if err := p.validateBindAsNotificationRuleSourceParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		p,
		"bindAsNotificationRuleSource",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (p *jsiiProxy_Pipeline) GetResourceNameAttribute(nameAttr *string) *string {
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

func (p *jsiiProxy_Pipeline) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) NotifyOnAnyActionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnAnyActionStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnAnyActionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) NotifyOnAnyManualApprovalStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnAnyManualApprovalStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnAnyManualApprovalStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) NotifyOnAnyStageStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnAnyStageStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnAnyStageStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) NotifyOnExecutionStateChange(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnExecutionStateChangeParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnExecutionStateChange",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnStateChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Pipeline) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pipeline) Stage(stageName *string) IStage {
	if err := p.validateStageParameters(stageName); err != nil {
		panic(err)
	}
	var returns IStage

	_jsii_.Invoke(
		p,
		"stage",
		[]interface{}{stageName},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Pipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

