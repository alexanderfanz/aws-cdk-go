package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::IAM::Policy resource associates an IAM policy with IAM users, roles, or groups.
//
// For more information about IAM policies, see [Overview of IAM
// Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html)
// in the IAM User Guide guide.
//
// Example:
//   var postAuthFn function
//
//
//   userpool := cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	lambdaTriggers: &userPoolTriggers{
//   		postAuthentication: postAuthFn,
//   	},
//   })
//
//   // provide permissions to describe the user pool scoped to the ARN the user pool
//   postAuthFn.role.attachInlinePolicy(iam.NewPolicy(this, jsii.String("userpool-policy"), &policyProps{
//   	statements: []policyStatement{
//   		iam.NewPolicyStatement(&policyStatementProps{
//   			actions: []*string{
//   				jsii.String("cognito-idp:DescribeUserPool"),
//   			},
//   			resources: []*string{
//   				userpool.userPoolArn,
//   			},
//   		}),
//   	},
//   }))
//
type Policy interface {
	awscdk.Resource
	IPolicy
	// The policy document.
	Document() PolicyDocument
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	// The name of this policy.
	PolicyName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a statement to the policy document.
	AddStatements(statement ...PolicyStatement)
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
	// Attaches this policy to a group.
	AttachToGroup(group IGroup)
	// Attaches this policy to a role.
	AttachToRole(role IRole)
	// Attaches this policy to a user.
	AttachToUser(user IUser)
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Policy
type jsiiProxy_Policy struct {
	internal.Type__awscdkResource
	jsiiProxy_IPolicy
}

func (j *jsiiProxy_Policy) Document() PolicyDocument {
	var returns PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Policy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Policy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Policy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Policy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Policy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewPolicy(scope constructs.Construct, id *string, props *PolicyProps) Policy {
	_init_.Initialize()

	j := jsiiProxy_Policy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.Policy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPolicy_Override(p Policy, scope constructs.Construct, id *string, props *PolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.Policy",
		[]interface{}{scope, id, props},
		p,
	)
}

// Import a policy in this app based on its name.
func Policy_FromPolicyName(scope constructs.Construct, id *string, policyName *string) IPolicy {
	_init_.Initialize()

	var returns IPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Policy",
		"fromPolicyName",
		[]interface{}{scope, id, policyName},
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
func Policy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Policy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Policy_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Policy",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Policy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.Policy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Policy) AddStatements(statement ...PolicyStatement) {
	args := []interface{}{}
	for _, a := range statement {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addStatements",
		args,
	)
}

func (p *jsiiProxy_Policy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_Policy) AttachToGroup(group IGroup) {
	_jsii_.InvokeVoid(
		p,
		"attachToGroup",
		[]interface{}{group},
	)
}

func (p *jsiiProxy_Policy) AttachToRole(role IRole) {
	_jsii_.InvokeVoid(
		p,
		"attachToRole",
		[]interface{}{role},
	)
}

func (p *jsiiProxy_Policy) AttachToUser(user IUser) {
	_jsii_.InvokeVoid(
		p,
		"attachToUser",
		[]interface{}{user},
	)
}

func (p *jsiiProxy_Policy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Policy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Policy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Policy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
