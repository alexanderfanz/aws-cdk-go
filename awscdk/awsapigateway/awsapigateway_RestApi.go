package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a REST API in Amazon API Gateway.
//
// Use `addResource` and `addMethod` to configure the API model.
//
// By default, the API will automatically be deployed and accessible from a
// public endpoint.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
// Experimental.
type RestApi interface {
	RestApiBase
	// Experimental.
	CloudWatchAccount() CfnAccount
	// Experimental.
	SetCloudWatchAccount(val CfnAccount)
	// API Gateway stage that points to the latest deployment (if defined).
	//
	// If `deploy` is disabled, you will need to explicitly assign this value in order to
	// set up integrations.
	// Experimental.
	DeploymentStage() Stage
	// Experimental.
	SetDeploymentStage(val Stage)
	// The first domain name mapped to this API, if defined through the `domainName` configuration prop, or added via `addDomainName`.
	// Experimental.
	DomainName() DomainName
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
	// API Gateway deployment that represents the latest changes of the API.
	//
	// This resource will be automatically updated every time the REST API model changes.
	// This will be undefined if `deploy` is false.
	// Experimental.
	LatestDeployment() Deployment
	// The list of methods bound to this RestApi.
	// Experimental.
	Methods() *[]Method
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
	// The ID of this API Gateway RestApi.
	// Experimental.
	RestApiId() *string
	// A human friendly name for this Rest API.
	//
	// Note that this is different from `restApiId`.
	// Experimental.
	RestApiName() *string
	// The resource ID of the root resource.
	// Experimental.
	RestApiRootResourceId() *string
	// Represents the root resource of this API endpoint ('/').
	//
	// Resources and Methods are added to this resource.
	// Experimental.
	Root() IResource
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The deployed root URL of this REST API.
	// Experimental.
	Url() *string
	// Add an ApiKey.
	// Experimental.
	AddApiKey(id *string, options *ApiKeyOptions) IApiKey
	// Defines an API Gateway domain name and maps it to this API.
	// Experimental.
	AddDomainName(id *string, options *DomainNameOptions) DomainName
	// Adds a new gateway response.
	// Experimental.
	AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse
	// Adds a new model.
	// Experimental.
	AddModel(id *string, props *ModelOptions) Model
	// Adds a new request validator.
	// Experimental.
	AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator
	// Adds a usage plan.
	// Experimental.
	AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan
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
	// Gets the "execute-api" ARN.
	// Experimental.
	ArnForExecuteApi(method *string, path *string, stage *string) *string
	// Deprecated: This method will be made internal. No replacement
	ConfigureCloudWatchRole(apiResource CfnRestApi)
	// Deprecated: This method will be made internal. No replacement
	ConfigureDeployment(props *RestApiBaseProps)
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
	// Returns the given named metric for this API.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the API cache in a given period.
	//
	// Default: sum over 5 minutes.
	// Experimental.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of requests served from the backend in a given period, when API caching is enabled.
	//
	// Default: sum over 5 minutes.
	// Experimental.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	// Experimental.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	//
	// Default: sample count over 5 minutes.
	// Experimental.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	//
	// Default: average over 5 minutes.
	// Experimental.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	//
	// Default: average over 5 minutes.
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	//
	// Default: sum over 5 minutes.
	// Experimental.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Returns the URL for an HTTP path.
	//
	// Fails if `deploymentStage` is not set either by `deploy` or explicitly.
	// Experimental.
	UrlForPath(path *string) *string
	// Performs validation of the REST API.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for RestApi
type jsiiProxy_RestApi struct {
	jsiiProxy_RestApiBase
}

func (j *jsiiProxy_RestApi) CloudWatchAccount() CfnAccount {
	var returns CfnAccount
	_jsii_.Get(
		j,
		"cloudWatchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) DeploymentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"deploymentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) DomainName() DomainName {
	var returns DomainName
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) LatestDeployment() Deployment {
	var returns Deployment
	_jsii_.Get(
		j,
		"latestDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Methods() *[]Method {
	var returns *[]Method
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) RestApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) RestApiRootResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiRootResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Root() IResource {
	var returns IResource
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RestApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewRestApi(scope constructs.Construct, id *string, props *RestApiProps) RestApi {
	_init_.Initialize()

	if err := validateNewRestApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RestApi{}

	_jsii_.Create(
		"monocdk.aws_apigateway.RestApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRestApi_Override(r RestApi, scope constructs.Construct, id *string, props *RestApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.RestApi",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RestApi)SetCloudWatchAccount(val CfnAccount) {
	_jsii_.Set(
		j,
		"cloudWatchAccount",
		val,
	)
}

func (j *jsiiProxy_RestApi)SetDeploymentStage(val Stage) {
	if err := j.validateSetDeploymentStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentStage",
		val,
	)
}

// Import an existing RestApi that can be configured with additional Methods and Resources.
// Experimental.
func RestApi_FromRestApiAttributes(scope constructs.Construct, id *string, attrs *RestApiAttributes) IRestApi {
	_init_.Initialize()

	if err := validateRestApi_FromRestApiAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IRestApi

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.RestApi",
		"fromRestApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing RestApi.
// Experimental.
func RestApi_FromRestApiId(scope constructs.Construct, id *string, restApiId *string) IRestApi {
	_init_.Initialize()

	if err := validateRestApi_FromRestApiIdParameters(scope, id, restApiId); err != nil {
		panic(err)
	}
	var returns IRestApi

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.RestApi",
		"fromRestApiId",
		[]interface{}{scope, id, restApiId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func RestApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRestApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.RestApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func RestApi_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRestApi_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.RestApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddApiKey(id *string, options *ApiKeyOptions) IApiKey {
	if err := r.validateAddApiKeyParameters(id, options); err != nil {
		panic(err)
	}
	var returns IApiKey

	_jsii_.Invoke(
		r,
		"addApiKey",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddDomainName(id *string, options *DomainNameOptions) DomainName {
	if err := r.validateAddDomainNameParameters(id, options); err != nil {
		panic(err)
	}
	var returns DomainName

	_jsii_.Invoke(
		r,
		"addDomainName",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddGatewayResponse(id *string, options *GatewayResponseOptions) GatewayResponse {
	if err := r.validateAddGatewayResponseParameters(id, options); err != nil {
		panic(err)
	}
	var returns GatewayResponse

	_jsii_.Invoke(
		r,
		"addGatewayResponse",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddModel(id *string, props *ModelOptions) Model {
	if err := r.validateAddModelParameters(id, props); err != nil {
		panic(err)
	}
	var returns Model

	_jsii_.Invoke(
		r,
		"addModel",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddRequestValidator(id *string, props *RequestValidatorOptions) RequestValidator {
	if err := r.validateAddRequestValidatorParameters(id, props); err != nil {
		panic(err)
	}
	var returns RequestValidator

	_jsii_.Invoke(
		r,
		"addRequestValidator",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) AddUsagePlan(id *string, props *UsagePlanProps) UsagePlan {
	if err := r.validateAddUsagePlanParameters(id, props); err != nil {
		panic(err)
	}
	var returns UsagePlan

	_jsii_.Invoke(
		r,
		"addUsagePlan",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RestApi) ArnForExecuteApi(method *string, path *string, stage *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"arnForExecuteApi",
		[]interface{}{method, path, stage},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) ConfigureCloudWatchRole(apiResource CfnRestApi) {
	if err := r.validateConfigureCloudWatchRoleParameters(apiResource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"configureCloudWatchRole",
		[]interface{}{apiResource},
	)
}

func (r *jsiiProxy_RestApi) ConfigureDeployment(props *RestApiBaseProps) {
	if err := r.validateConfigureDeploymentParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"configureDeployment",
		[]interface{}{props},
	)
}

func (r *jsiiProxy_RestApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (r *jsiiProxy_RestApi) GetResourceNameAttribute(nameAttr *string) *string {
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

func (r *jsiiProxy_RestApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricCacheHitCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricCacheMissCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricClientErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricIntegrationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := r.validateMetricServerErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		r,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RestApi) OnSynthesize(session constructs.ISynthesisSession) {
	if err := r.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RestApi) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RestApi) Synthesize(session awscdk.ISynthesisSession) {
	if err := r.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

func (r *jsiiProxy_RestApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) UrlForPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"urlForPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RestApi) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

