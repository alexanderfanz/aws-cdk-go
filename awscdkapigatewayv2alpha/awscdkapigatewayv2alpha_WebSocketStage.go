// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a stage where an instance of the API is deployed.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var messageHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &webSocketStageProps{
//   	webSocketApi: webSocketApi,
//   	stageName: jsii.String("dev"),
//   	autoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.addRoute(jsii.String("sendmessage"), &webSocketRouteOptions{
//   	integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
// Experimental.
type WebSocketStage interface {
	awscdk.Resource
	IStage
	IWebSocketStage
	// The API this stage is associated to.
	// Experimental.
	Api() IWebSocketApi
	// Experimental.
	BaseApi() IApi
	// The callback URL to this stage.
	// Experimental.
	CallbackUrl() *string
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
	// The name of the stage;
	//
	// its primary identifier.
	// Experimental.
	StageName() *string
	// The websocket URL to this stage.
	// Experimental.
	Url() *string
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
	// Grant access to the API Gateway management API for this WebSocket API Stage to an IAM principal (Role/Group/User).
	// Experimental.
	GrantManagementApiAccess(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this HTTP Api Gateway Stage.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebSocketStage
type jsiiProxy_WebSocketStage struct {
	internal.Type__awscdkResource
	jsiiProxy_IStage
	jsiiProxy_IWebSocketStage
}

func (j *jsiiProxy_WebSocketStage) Api() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) BaseApi() IApi {
	var returns IApi
	_jsii_.Get(
		j,
		"baseApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebSocketStage) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketStage(scope constructs.Construct, id *string, props *WebSocketStageProps) WebSocketStage {
	_init_.Initialize()

	if err := validateNewWebSocketStageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketStage{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketStage_Override(w WebSocketStage, scope constructs.Construct, id *string, props *WebSocketStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		[]interface{}{scope, id, props},
		w,
	)
}

// Import an existing stage into this CDK app.
// Experimental.
func WebSocketStage_FromWebSocketStageAttributes(scope constructs.Construct, id *string, attrs *WebSocketStageAttributes) IWebSocketStage {
	_init_.Initialize()

	if err := validateWebSocketStage_FromWebSocketStageAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IWebSocketStage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"fromWebSocketStageAttributes",
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
// Experimental.
func WebSocketStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebSocketStage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func WebSocketStage_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWebSocketStage_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WebSocketStage_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWebSocketStage_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketStage",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := w.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (w *jsiiProxy_WebSocketStage) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketStage) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (w *jsiiProxy_WebSocketStage) GetResourceNameAttribute(nameAttr *string) *string {
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

func (w *jsiiProxy_WebSocketStage) GrantManagementApiAccess(identity awsiam.IGrantable) awsiam.Grant {
	if err := w.validateGrantManagementApiAccessParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		w,
		"grantManagementApiAccess",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketStage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := w.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		w,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebSocketStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

