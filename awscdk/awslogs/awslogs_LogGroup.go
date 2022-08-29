package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a CloudWatch Log Group.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("Log Group"))
//   logBucket := s3.NewBucket(this, jsii.String("S3 Bucket"))
//
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
//   	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
//   	jobDriver: &jobDriver{
//   		sparkSubmitJobDriver: &sparkSubmitJobDriver{
//   			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   			sparkSubmitParameters: jsii.String("--conf spark.executor.instances=2 --conf spark.executor.memory=2G --conf spark.executor.cores=2 --conf spark.driver.cores=1"),
//   		},
//   	},
//   	monitoring: &monitoring{
//   		logGroup: logGroup,
//   		logBucket: logBucket,
//   	},
//   })
//
type LogGroup interface {
	awscdk.Resource
	ILogGroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of this log group.
	LogGroupArn() *string
	// The name of this log group.
	LogGroupName() *string
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
	// Create a new Metric Filter on this Log Group.
	AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter
	// Create a new Log Stream for this Log Group.
	AddStream(id *string, props *StreamOptions) LogStream
	// Create a new Subscription Filter on this Log Group.
	AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter
	// Adds a statement to the resource policy associated with this log group.
	//
	// A resource policy will be automatically created upon the first call to `addToResourcePolicy`.
	//
	// Any ARN Principals inside of the statement will be converted into AWS Account ID strings
	// because CloudWatch Logs Resource Policies do not accept ARN principals.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// Extract a metric from structured log events in the LogGroup.
	//
	// Creates a MetricFilter on this LogGroup that will extract the value
	// of the indicated JSON field in all records where it occurs.
	//
	// The metric will be available in CloudWatch Metrics under the
	// indicated namespace and name.
	//
	// Returns: A Metric object representing the extracted metric.
	ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric
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
	// Give the indicated permissions on this log group and all streams.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Give permissions to create and write to streams in this log group.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Public method to get the physical name of this log group.
	//
	// Returns: Physical name of log group.
	LogGroupPhysicalName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for LogGroup
type jsiiProxy_LogGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_ILogGroup
}

func (j *jsiiProxy_LogGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewLogGroup(scope constructs.Construct, id *string, props *LogGroupProps) LogGroup {
	_init_.Initialize()

	j := jsiiProxy_LogGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.LogGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLogGroup_Override(l LogGroup, scope constructs.Construct, id *string, props *LogGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_logs.LogGroup",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an existing LogGroup given its ARN.
func LogGroup_FromLogGroupArn(scope constructs.Construct, id *string, logGroupArn *string) ILogGroup {
	_init_.Initialize()

	var returns ILogGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogGroup",
		"fromLogGroupArn",
		[]interface{}{scope, id, logGroupArn},
		&returns,
	)

	return returns
}

// Import an existing LogGroup given its name.
func LogGroup_FromLogGroupName(scope constructs.Construct, id *string, logGroupName *string) ILogGroup {
	_init_.Initialize()

	var returns ILogGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogGroup",
		"fromLogGroupName",
		[]interface{}{scope, id, logGroupName},
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
func LogGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func LogGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LogGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_logs.LogGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter {
	var returns MetricFilter

	_jsii_.Invoke(
		l,
		"addMetricFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddStream(id *string, props *StreamOptions) LogStream {
	var returns LogStream

	_jsii_.Invoke(
		l,
		"addStream",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter {
	var returns SubscriptionFilter

	_jsii_.Invoke(
		l,
		"addSubscriptionFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		l,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LogGroup) ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"extractMetric",
		[]interface{}{jsonField, metricNamespace, metricName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		l,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) LogGroupPhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"logGroupPhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
