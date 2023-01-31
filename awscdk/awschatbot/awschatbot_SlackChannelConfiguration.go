package awschatbot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awschatbot/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v3"
)

// A new Slack channel configuration.
//
// Example:
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//
//   target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
//   	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := project.notifyOnBuildSucceeded(jsii.String("NotifyOnBuildSucceeded"), target)
//
// Experimental.
type SlackChannelConfiguration interface {
	awscdk.Resource
	ISlackChannelConfiguration
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
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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
	// The permission role of Slack channel configuration.
	// Experimental.
	Role() awsiam.IRole
	// The ARN of the Slack channel configuration In the form of arn:aws:chatbot:{region}:{account}:chat-configuration/slack-channel/{slackChannelName}.
	// Experimental.
	SlackChannelConfigurationArn() *string
	// The name of Slack channel configuration.
	// Experimental.
	SlackChannelConfigurationName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a SNS topic that deliver notifications to AWS Chatbot.
	// Experimental.
	AddNotificationTopic(notificationTopic awssns.ITopic)
	// Adds extra permission to iam-role of Slack channel configuration.
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
	// Returns a target configuration for notification rule.
	// Experimental.
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
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
	// Return the given named metric for this SlackChannelConfiguration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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

// The jsii proxy struct for SlackChannelConfiguration
type jsiiProxy_SlackChannelConfiguration struct {
	internal.Type__awscdkResource
	jsiiProxy_ISlackChannelConfiguration
}

func (j *jsiiProxy_SlackChannelConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) SlackChannelConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) SlackChannelConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSlackChannelConfiguration(scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) SlackChannelConfiguration {
	_init_.Initialize()

	if err := validateNewSlackChannelConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SlackChannelConfiguration{}

	_jsii_.Create(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSlackChannelConfiguration_Override(s SlackChannelConfiguration, scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an existing Slack channel configuration provided an ARN.
//
// Returns: a reference to the existing Slack channel configuration.
// Experimental.
func SlackChannelConfiguration_FromSlackChannelConfigurationArn(scope constructs.Construct, id *string, slackChannelConfigurationArn *string) ISlackChannelConfiguration {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_FromSlackChannelConfigurationArnParameters(scope, id, slackChannelConfigurationArn); err != nil {
		panic(err)
	}
	var returns ISlackChannelConfiguration

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"fromSlackChannelConfigurationArn",
		[]interface{}{scope, id, slackChannelConfigurationArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SlackChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func SlackChannelConfiguration_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for All SlackChannelConfigurations.
// Experimental.
func SlackChannelConfiguration_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_chatbot.SlackChannelConfiguration",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) AddNotificationTopic(notificationTopic awssns.ITopic) {
	if err := s.validateAddNotificationTopicParameters(notificationTopic); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addNotificationTopic",
		[]interface{}{notificationTopic},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := s.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	if err := s.validateBindAsNotificationRuleTargetParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		s,
		"bindAsNotificationRuleTarget",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

