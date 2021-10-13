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
// Experimental.
type NodeProxyAgentLayer interface {
	awslambda.LayerVersion
	CompatibleRuntimes() *[]awslambda.Runtime
	Env() *awscdk.ResourceEnvironment
	LayerVersionArn() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddPermission(id *string, permission *awslambda.LayerVersionPermission)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
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


// Experimental.
func NewNodeProxyAgentLayer(scope constructs.Construct, id *string) NodeProxyAgentLayer {
	_init_.Initialize()

	j := jsiiProxy_NodeProxyAgentLayer{}

	_jsii_.Create(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
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
// Experimental.
func NodeProxyAgentLayer_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) awslambda.ILayerVersion {
	_init_.Initialize()

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
// Experimental.
func NodeProxyAgentLayer_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) awslambda.ILayerVersion {
	_init_.Initialize()

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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NodeProxyAgentLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func NodeProxyAgentLayer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.lambda_layer_node_proxy_agent.NodeProxyAgentLayer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add permission for this layer version to specific entities.
//
// Usage within
// the same account where the layer is defined is always allowed and does not
// require calling this method. Note that the principal that creates the
// Lambda function using the layer (for example, a CloudFormation changeset
// execution role) also needs to have the ``lambda:GetLayerVersion``
// permission on the layer version.
// Experimental.
func (n *jsiiProxy_NodeProxyAgentLayer) AddPermission(id *string, permission *awslambda.LayerVersionPermission) {
	_jsii_.InvokeVoid(
		n,
		"addPermission",
		[]interface{}{id, permission},
	)
}

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
func (n *jsiiProxy_NodeProxyAgentLayer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
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

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (n *jsiiProxy_NodeProxyAgentLayer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (n *jsiiProxy_NodeProxyAgentLayer) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
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
