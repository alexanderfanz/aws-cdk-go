package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsses/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A new receipt rule set.
//
// Example:
//   ruleSet := ses.NewReceiptRuleSet(this, jsii.String("RuleSet"))
//
//   awsRule := ruleSet.addRule(jsii.String("Aws"), &receiptRuleOptions{
//   	recipients: []*string{
//   		jsii.String("aws.com"),
//   	},
//   })
//
// Experimental.
type ReceiptRuleSet interface {
	awscdk.Resource
	IReceiptRuleSet
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
	// The receipt rule set name.
	// Experimental.
	ReceiptRuleSetName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a drop spam rule.
	// Experimental.
	AddDropSpamRule()
	// Adds a new receipt rule in this rule set.
	//
	// The new rule is added after
	// the last added rule unless `after` is specified.
	// Experimental.
	AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule
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

// The jsii proxy struct for ReceiptRuleSet
type jsiiProxy_ReceiptRuleSet struct {
	internal.Type__awscdkResource
	jsiiProxy_IReceiptRuleSet
}

func (j *jsiiProxy_ReceiptRuleSet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) ReceiptRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReceiptRuleSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewReceiptRuleSet(scope constructs.Construct, id *string, props *ReceiptRuleSetProps) ReceiptRuleSet {
	_init_.Initialize()

	if err := validateNewReceiptRuleSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReceiptRuleSet{}

	_jsii_.Create(
		"monocdk.aws_ses.ReceiptRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewReceiptRuleSet_Override(r ReceiptRuleSet, scope constructs.Construct, id *string, props *ReceiptRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses.ReceiptRuleSet",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an exported receipt rule set.
// Experimental.
func ReceiptRuleSet_FromReceiptRuleSetName(scope constructs.Construct, id *string, receiptRuleSetName *string) IReceiptRuleSet {
	_init_.Initialize()

	if err := validateReceiptRuleSet_FromReceiptRuleSetNameParameters(scope, id, receiptRuleSetName); err != nil {
		panic(err)
	}
	var returns IReceiptRuleSet

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.ReceiptRuleSet",
		"fromReceiptRuleSetName",
		[]interface{}{scope, id, receiptRuleSetName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ReceiptRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReceiptRuleSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.ReceiptRuleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ReceiptRuleSet_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateReceiptRuleSet_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ses.ReceiptRuleSet",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) AddDropSpamRule() {
	_jsii_.InvokeVoid(
		r,
		"addDropSpamRule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReceiptRuleSet) AddRule(id *string, options *ReceiptRuleOptions) ReceiptRule {
	if err := r.validateAddRuleParameters(id, options); err != nil {
		panic(err)
	}
	var returns ReceiptRule

	_jsii_.Invoke(
		r,
		"addRule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ReceiptRuleSet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReceiptRuleSet) OnSynthesize(session constructs.ISynthesisSession) {
	if err := r.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_ReceiptRuleSet) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReceiptRuleSet) Synthesize(session awscdk.ISynthesisSession) {
	if err := r.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_ReceiptRuleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReceiptRuleSet) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

