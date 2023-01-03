package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define an ApplicationListener.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var alb applicationLoadBalancer
//
//   listener := alb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   targetGroup := listener.addTargets(jsii.String("Fleet"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
//   	loadBalancer: codedeploy.loadBalancer.application(targetGroup),
//   })
//
type ApplicationListener interface {
	BaseListener
	IApplicationListener
	// Manage connections to this ApplicationListener.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// ARN of the listener.
	ListenerArn() *string
	// Load balancer this listener is associated with.
	LoadBalancer() IApplicationLoadBalancer
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
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Perform the given default action on incoming requests.
	//
	// This allows full control of the default action of the load balancer,
	// including Action chaining, fixed responses and redirect responses. See
	// the `ListenerAction` class for all options.
	//
	// It's possible to add routing conditions to the Action added in this way.
	// At least one Action must be added without conditions (which becomes the
	// default Action).
	AddAction(id *string, props *AddApplicationActionProps)
	// Add one or more certificates to this listener.
	//
	// After the first certificate, this creates ApplicationListenerCertificates
	// resources since cloudformation requires the certificates array on the
	// listener resource to have a length of 1.
	AddCertificates(id *string, certificates *[]IListenerCertificate)
	// Load balance incoming requests to the given target groups.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use `addAction()`.
	//
	// It's possible to add routing conditions to the TargetGroups added in this
	// way. At least one TargetGroup must be added without conditions (which will
	// become the default Action for this listener).
	AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps)
	// Load balance incoming requests to the given load balancing targets.
	//
	// This method implicitly creates an ApplicationTargetGroup for the targets
	// involved, and a 'forward' action to route traffic to the given TargetGroup.
	//
	// If you want more control over the precise setup, create the TargetGroup
	// and use `addAction` yourself.
	//
	// It's possible to add conditions to the targets added in this way. At least
	// one set of targets must be added without conditions.
	//
	// Returns: The newly created target group.
	AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup
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
	// Register that a connectable that has been added to this load balancer.
	//
	// Don't call this directly. It is called by ApplicationTargetGroup.
	RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port)
	// Returns a string representation of this construct.
	ToString() *string
	// Validate this listener.
	ValidateListener() *[]*string
}

// The jsii proxy struct for ApplicationListener
type jsiiProxy_ApplicationListener struct {
	jsiiProxy_BaseListener
	jsiiProxy_IApplicationListener
}

func (j *jsiiProxy_ApplicationListener) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) LoadBalancer() IApplicationLoadBalancer {
	var returns IApplicationLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewApplicationListener(scope constructs.Construct, id *string, props *ApplicationListenerProps) ApplicationListener {
	_init_.Initialize()

	if err := validateNewApplicationListenerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationListener{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationListener_Override(a ApplicationListener, scope constructs.Construct, id *string, props *ApplicationListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing listener.
func ApplicationListener_FromApplicationListenerAttributes(scope constructs.Construct, id *string, attrs *ApplicationListenerAttributes) IApplicationListener {
	_init_.Initialize()

	if err := validateApplicationListener_FromApplicationListenerAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IApplicationListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"fromApplicationListenerAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Look up an ApplicationListener.
func ApplicationListener_FromLookup(scope constructs.Construct, id *string, options *ApplicationListenerLookupOptions) IApplicationListener {
	_init_.Initialize()

	if err := validateApplicationListener_FromLookupParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns IApplicationListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"fromLookup",
		[]interface{}{scope, id, options},
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
func ApplicationListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationListener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ApplicationListener_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplicationListener_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ApplicationListener_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplicationListener_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListener) AddAction(id *string, props *AddApplicationActionProps) {
	if err := a.validateAddActionParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAction",
		[]interface{}{id, props},
	)
}

func (a *jsiiProxy_ApplicationListener) AddCertificates(id *string, certificates *[]IListenerCertificate) {
	if err := a.validateAddCertificatesParameters(id, certificates); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addCertificates",
		[]interface{}{id, certificates},
	)
}

func (a *jsiiProxy_ApplicationListener) AddTargetGroups(id *string, props *AddApplicationTargetGroupsProps) {
	if err := a.validateAddTargetGroupsParameters(id, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addTargetGroups",
		[]interface{}{id, props},
	)
}

func (a *jsiiProxy_ApplicationListener) AddTargets(id *string, props *AddApplicationTargetsProps) ApplicationTargetGroup {
	if err := a.validateAddTargetsParameters(id, props); err != nil {
		panic(err)
	}
	var returns ApplicationTargetGroup

	_jsii_.Invoke(
		a,
		"addTargets",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ApplicationListener) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListener) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (a *jsiiProxy_ApplicationListener) GetResourceNameAttribute(nameAttr *string) *string {
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

func (a *jsiiProxy_ApplicationListener) RegisterConnectable(connectable awsec2.IConnectable, portRange awsec2.Port) {
	if err := a.validateRegisterConnectableParameters(connectable, portRange); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerConnectable",
		[]interface{}{connectable, portRange},
	)
}

func (a *jsiiProxy_ApplicationListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListener) ValidateListener() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateListener",
		nil, // no parameters
		&returns,
	)

	return returns
}

