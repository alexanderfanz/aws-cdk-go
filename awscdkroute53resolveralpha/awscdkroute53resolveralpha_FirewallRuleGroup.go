// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Firewall Rule Group.
//
// Example:
//   var myBlockList firewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleGroup interface {
	awscdk.Resource
	IFirewallRuleGroup
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
	// The ARN (Amazon Resource Name) of the rule group.
	// Experimental.
	FirewallRuleGroupArn() *string
	// The date and time that the rule group was created.
	// Experimental.
	FirewallRuleGroupCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallRuleGroupCreatorRequestId() *string
	// The ID of the rule group.
	// Experimental.
	FirewallRuleGroupId() *string
	// The date and time that the rule group was last modified.
	// Experimental.
	FirewallRuleGroupModificationTime() *string
	// The AWS account ID for the account that created the rule group.
	// Experimental.
	FirewallRuleGroupOwnerId() *string
	// The number of rules in the rule group.
	// Experimental.
	FirewallRuleGroupRuleCount() *float64
	// Whether the rule group is shared with other AWS accounts, or was shared with the current account by another AWS account.
	// Experimental.
	FirewallRuleGroupShareStatus() *string
	// The status of the rule group.
	// Experimental.
	FirewallRuleGroupStatus() *string
	// Additional information about the status of the rule group.
	// Experimental.
	FirewallRuleGroupStatusMessage() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Adds a rule to this group.
	// Experimental.
	AddRule(rule *FirewallRule) FirewallRuleGroup
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
	// Associates this Firewall Rule Group with a VPC.
	// Experimental.
	Associate(id *string, props *FirewallRuleGroupAssociationOptions) FirewallRuleGroupAssociation
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

// The jsii proxy struct for FirewallRuleGroup
type jsiiProxy_FirewallRuleGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IFirewallRuleGroup
}

func (j *jsiiProxy_FirewallRuleGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupRuleCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallRuleGroupRuleCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) FirewallRuleGroupStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallRuleGroupStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleGroup(scope constructs.Construct, id *string, props *FirewallRuleGroupProps) FirewallRuleGroup {
	_init_.Initialize()

	if err := validateNewFirewallRuleGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallRuleGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallRuleGroup_Override(f FirewallRuleGroup, scope constructs.Construct, id *string, props *FirewallRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing Firewall Rule Group.
// Experimental.
func FirewallRuleGroup_FromFirewallRuleGroupId(scope constructs.Construct, id *string, firewallRuleGroupId *string) IFirewallRuleGroup {
	_init_.Initialize()

	if err := validateFirewallRuleGroup_FromFirewallRuleGroupIdParameters(scope, id, firewallRuleGroupId); err != nil {
		panic(err)
	}
	var returns IFirewallRuleGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"fromFirewallRuleGroupId",
		[]interface{}{scope, id, firewallRuleGroupId},
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
func FirewallRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFirewallRuleGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func FirewallRuleGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFirewallRuleGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallRuleGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFirewallRuleGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) AddRule(rule *FirewallRule) FirewallRuleGroup {
	if err := f.validateAddRuleParameters(rule); err != nil {
		panic(err)
	}
	var returns FirewallRuleGroup

	_jsii_.Invoke(
		f,
		"addRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallRuleGroup) Associate(id *string, props *FirewallRuleGroupAssociationOptions) FirewallRuleGroupAssociation {
	if err := f.validateAssociateParameters(id, props); err != nil {
		panic(err)
	}
	var returns FirewallRuleGroupAssociation

	_jsii_.Invoke(
		f,
		"associate",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

