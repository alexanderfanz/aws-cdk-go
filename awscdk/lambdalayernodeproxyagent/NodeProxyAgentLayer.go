package lambdalayernodeproxyagent

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/lambdalayernodeproxyagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Lambda layer that includes the NPM dependency `proxy-agent`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   fn.AddLayers(awscdk.NewNodeProxyAgentLayer(this, jsii.String("NodeProxyAgentLayer")))
//
type NodeProxyAgentLayer interface {
	awslambda.LayerVersion
	// The runtimes compatible with this Layer.
	CompatibleRuntimes() *[]awslambda.Runtime
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of the Lambda Layer version that this Layer defines.
	LayerVersionArn() *string
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
	// Add permission for this layer version to specific entities.
	//
	// Usage within
	// the same account where the layer is defined is always allowed and does not
	// require calling this method. Note that the principal that creates the
	// Lambda function using the layer (for example, a CloudFormation changeset
	// execution role) also needs to have the ``lambda:GetLayerVersion``
	// permission on the layer version.
	AddPermission(id *string, permission *awslambda.LayerVersionPermission)
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NodeProxyAgentLayer
type jsiiProxy_NodeProxyAgentLayer struct {
	internal.Type__awslambdaLayerVersion
}

func (j *jsiiProxy_NodeProxyAgentLayer) CompatibleRuntimes() *[]awslambda.Runtime {
	var returns *[]awslambda.Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProxyAgentLayer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProxyAgentLayer) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProxyAgentLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProxyAgentLayer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProxyAgentLayer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNodeProxyAgentLayer(scope constructs.Construct, id *string) NodeProxyAgentLayer {
	_init_.Initialize()

	if err := validateNewNodeProxyAgentLayerParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodeProxyAgentLayer{}

	_jsii_.Create(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

func NewNodeProxyAgentLayer_Override(n NodeProxyAgentLayer, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		[]interface{}{scope, id},
		n,
	)
}

// Imports a layer version by ARN.
//
// Assumes it is compatible with all Lambda runtimes.
func NodeProxyAgentLayer_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) awslambda.ILayerVersion {
	_init_.Initialize()

	if err := validateNodeProxyAgentLayer_FromLayerVersionArnParameters(scope, id, layerVersionArn); err != nil {
		panic(err)
	}
	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"fromLayerVersionArn",
		[]interface{}{scope, id, layerVersionArn},
		&returns,
	)

	return returns
}

// Imports a Layer that has been defined externally.
func NodeProxyAgentLayer_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) awslambda.ILayerVersion {
	_init_.Initialize()

	if err := validateNodeProxyAgentLayer_FromLayerVersionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"fromLayerVersionAttributes",
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
func NodeProxyAgentLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodeProxyAgentLayer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func NodeProxyAgentLayer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNodeProxyAgentLayer_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NodeProxyAgentLayer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNodeProxyAgentLayer_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProxyAgentLayer) AddPermission(id *string, permission *awslambda.LayerVersionPermission) {
	if err := n.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (n *jsiiProxy_NodeProxyAgentLayer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NodeProxyAgentLayer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProxyAgentLayer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := n.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProxyAgentLayer) GetResourceNameAttribute(nameAttr *string) *string {
	if err := n.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProxyAgentLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

