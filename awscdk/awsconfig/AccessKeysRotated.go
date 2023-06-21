package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Checks whether the active access keys are rotated within the number of days specified in `maxAge`.
//
// Example:
//   // compliant if access keys have been rotated within the last 90 days
//   // compliant if access keys have been rotated within the last 90 days
//   config.NewAccessKeysRotated(this, jsii.String("AccessKeyRotated"))
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//
type AccessKeysRotated interface {
	ManagedRule
	// The arn of the rule.
	ConfigRuleArn() *string
	// The compliance status of the rule.
	ConfigRuleComplianceType() *string
	// The id of the rule.
	ConfigRuleId() *string
	// The name of the rule.
	ConfigRuleName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	// The stack in which this resource is defined.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
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
	// Defines an EventBridge event rule which triggers for rule compliance events.
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AccessKeysRotated
type jsiiProxy_AccessKeysRotated struct {
	jsiiProxy_ManagedRule
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessKeysRotated) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewAccessKeysRotated(scope constructs.Construct, id *string, props *AccessKeysRotatedProps) AccessKeysRotated {
	_init_.Initialize()

	if err := validateNewAccessKeysRotatedParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessKeysRotated{}

	_jsii_.Create(
		"aws-cdk-lib.aws_config.AccessKeysRotated",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAccessKeysRotated_Override(a AccessKeysRotated, scope constructs.Construct, id *string, props *AccessKeysRotatedProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_config.AccessKeysRotated",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AccessKeysRotated)SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_AccessKeysRotated)SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_AccessKeysRotated)SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
func AccessKeysRotated_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	if err := validateAccessKeysRotated_FromConfigRuleNameParameters(scope, id, configRuleName); err != nil {
		panic(err)
	}
	var returns IRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.AccessKeysRotated",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
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
func AccessKeysRotated_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessKeysRotated_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.AccessKeysRotated",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func AccessKeysRotated_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAccessKeysRotated_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.AccessKeysRotated",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func AccessKeysRotated_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAccessKeysRotated_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.AccessKeysRotated",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessKeysRotated) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AccessKeysRotated) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessKeysRotated) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (a *jsiiProxy_AccessKeysRotated) GetResourceNameAttribute(nameAttr *string) *string {
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

func (a *jsiiProxy_AccessKeysRotated) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := a.validateOnComplianceChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessKeysRotated) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := a.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessKeysRotated) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := a.validateOnReEvaluationStatusParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessKeysRotated) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

