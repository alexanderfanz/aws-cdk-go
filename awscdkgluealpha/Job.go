package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkgluealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Glue Job.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonStreaming(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	Description: jsii.String("an example Python Streaming job"),
//   })
//
// Experimental.
type Job interface {
	awscdk.Resource
	IJob
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
	// The principal this Glue Job is running as.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The ARN of the job.
	// Experimental.
	JobArn() *string
	// The name of the job.
	// Experimental.
	JobName() *string
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
	// The IAM role Glue assumes to run this job.
	// Experimental.
	Role() awsiam.IRole
	// The Spark UI logs location if Spark UI monitoring and debugging is enabled.
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html
	//
	// Experimental.
	SparkUILoggingLocation() *SparkUILoggingLocation
	// The stack in which this resource is defined.
	// Experimental.
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
	// Create a CloudWatch metric.
	// See: https://docs.aws.amazon.com/glue/latest/dg/monitoring-awsglue-with-cloudwatch-metrics.html
	//
	// Experimental.
	Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a CloudWatch Metric indicating job failure.
	//
	// This metric is based on the Rule returned by no-args onFailure() call.
	// Experimental.
	MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a CloudWatch Metric indicating job success.
	//
	// This metric is based on the Rule returned by no-args onSuccess() call.
	// Experimental.
	MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a CloudWatch Metric indicating job timeout.
	//
	// This metric is based on the Rule returned by no-args onTimeout() call.
	// Experimental.
	MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Create a CloudWatch Event Rule for this Glue Job when it's in a given state.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#glue-event-types
	//
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Return a CloudWatch Event Rule matching FAILED state.
	// Experimental.
	OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Create a CloudWatch Event Rule for the transition into the input jobState.
	// Experimental.
	OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule
	// Create a CloudWatch Event Rule matching JobState.SUCCEEDED.
	// Experimental.
	OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Return a CloudWatch Event Rule matching TIMEOUT state.
	// Experimental.
	OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Job
type jsiiProxy_Job struct {
	internal.Type__awscdkResource
	jsiiProxy_IJob
}

func (j *jsiiProxy_Job) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) SparkUILoggingLocation() *SparkUILoggingLocation {
	var returns *SparkUILoggingLocation
	_jsii_.Get(
		j,
		"sparkUILoggingLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Job) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewJob(scope constructs.Construct, id *string, props *JobProps) Job {
	_init_.Initialize()

	if err := validateNewJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Job{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Job",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJob_Override(j Job, scope constructs.Construct, id *string, props *JobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.Job",
		[]interface{}{scope, id, props},
		j,
	)
}

// Creates a Glue Job.
// Experimental.
func Job_FromJobAttributes(scope constructs.Construct, id *string, attrs *JobAttributes) IJob {
	_init_.Initialize()

	if err := validateJob_FromJobAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IJob

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"fromJobAttributes",
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
func Job_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Job_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJob_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Job_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateJob_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Job",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := j.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_Job) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := j.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) GetResourceNameAttribute(nameAttr *string) *string {
	if err := j.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) Metric(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := j.validateMetricParameters(metricName, type_, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metric",
		[]interface{}{metricName, type_, props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MetricFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := j.validateMetricFailureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metricFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MetricSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := j.validateMetricSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metricSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) MetricTimeout(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := j.validateMetricTimeoutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		j,
		"metricTimeout",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := j.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnFailure(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := j.validateOnFailureParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onFailure",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnStateChange(id *string, jobState JobState, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := j.validateOnStateChangeParameters(id, jobState, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onStateChange",
		[]interface{}{id, jobState, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnSuccess(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := j.validateOnSuccessParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onSuccess",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) OnTimeout(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := j.validateOnTimeoutParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		j,
		"onTimeout",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Job) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

