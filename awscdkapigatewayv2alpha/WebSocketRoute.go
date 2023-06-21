package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Route class that creates the Route for API Gateway WebSocket API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var webSocketApi webSocketApi
//   var webSocketRouteAuthorizer iWebSocketRouteAuthorizer
//   var webSocketRouteIntegration webSocketRouteIntegration
//
//   webSocketRoute := apigatewayv2_alpha.NewWebSocketRoute(this, jsii.String("MyWebSocketRoute"), &WebSocketRouteProps{
//   	Integration: webSocketRouteIntegration,
//   	RouteKey: jsii.String("routeKey"),
//   	WebSocketApi: webSocketApi,
//
//   	// the properties below are optional
//   	ApiKeyRequired: jsii.Boolean(false),
//   	Authorizer: webSocketRouteAuthorizer,
//   	ReturnResponse: jsii.Boolean(false),
//   })
//
// Experimental.
type WebSocketRoute interface {
	awscdk.Resource
	IWebSocketRoute
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
	// Integration response ID.
	// Experimental.
	IntegrationResponseId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Id of the Route.
	// Experimental.
	RouteId() *string
	// The key to this route.
	// Experimental.
	RouteKey() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The WebSocket API associated with this route.
	// Experimental.
	WebSocketApi() IWebSocketApi
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebSocketRoute
type jsiiProxy_WebSocketRoute struct {
	internal.Type__awscdkResource
	jsiiProxy_IWebSocketRoute
}

func (j *jsiiProxy_WebSocketRoute) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) IntegrationResponseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketRoute) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketRoute(scope constructs.Construct, id *string, props *WebSocketRouteProps) WebSocketRoute {
	_init_.Initialize()

	if err := validateNewWebSocketRouteParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketRoute{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketRoute_Override(w WebSocketRoute, scope constructs.Construct, id *string, props *WebSocketRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		[]interface{}{scope, id, props},
		w,
	)
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
func WebSocketRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebSocketRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func WebSocketRoute_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWebSocketRoute_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WebSocketRoute_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWebSocketRoute_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRoute",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := w.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (w *jsiiProxy_WebSocketRoute) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketRoute) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := w.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketRoute) GetResourceNameAttribute(nameAttr *string) *string {
	if err := w.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

