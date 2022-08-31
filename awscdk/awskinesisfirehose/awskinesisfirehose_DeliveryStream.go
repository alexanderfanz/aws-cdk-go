package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Create a Kinesis Data Firehose delivery stream.
//
// Example:
//   var bucket bucket
//   // Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &dataProcessorProps{
//   	bufferInterval: awscdk.Duration.minutes(jsii.Number(5)),
//   	bufferSize: awscdk.Size.mebibytes(jsii.Number(5)),
//   	retries: jsii.Number(5),
//   })
//   s3Destination := destinations.NewS3Bucket(bucket, &s3BucketProps{
//   	processor: lambdaProcessor,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
// Experimental.
type DeliveryStream interface {
	awscdk.Resource
	IDeliveryStream
	// Network connections between Kinesis Data Firehose and other resources, i.e. Redshift cluster.
	// Experimental.
	Connections() awsec2.Connections
	// The ARN of the delivery stream.
	// Experimental.
	DeliveryStreamArn() *string
	// The name of the delivery stream.
	// Experimental.
	DeliveryStreamName() *string
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
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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
	// Grant the `grantee` identity permissions to perform `actions`.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the `grantee` identity permissions to perform `firehose:PutRecord` and `firehose:PutRecordBatch` actions on this delivery stream.
	// Experimental.
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this delivery stream.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of bytes delivered to Amazon S3 for backup over the specified time period.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the age (from getting into Kinesis Data Firehose to now) of the oldest record in Kinesis Data Firehose.
	//
	// Any record older than this age has been delivered to the Amazon S3 bucket for backup.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of records delivered to Amazon S3 for backup over the specified time period.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of bytes ingested successfully into the delivery stream over the specified time period after throttling.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of records ingested successfully into the delivery stream over the specified time period after throttling.
	//
	// By default, this metric will be calculated as an average over a period of 5 minutes.
	// Experimental.
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for DeliveryStream
type jsiiProxy_DeliveryStream struct {
	internal.Type__awscdkResource
	jsiiProxy_IDeliveryStream
}

func (j *jsiiProxy_DeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDeliveryStream(scope constructs.Construct, id *string, props *DeliveryStreamProps) DeliveryStream {
	_init_.Initialize()

	if err := validateNewDeliveryStreamParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeliveryStream{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDeliveryStream_Override(d DeliveryStream, scope constructs.Construct, id *string, props *DeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing delivery stream from its ARN.
// Experimental.
func DeliveryStream_FromDeliveryStreamArn(scope constructs.Construct, id *string, deliveryStreamArn *string) IDeliveryStream {
	_init_.Initialize()

	if err := validateDeliveryStream_FromDeliveryStreamArnParameters(scope, id, deliveryStreamArn); err != nil {
		panic(err)
	}
	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"fromDeliveryStreamArn",
		[]interface{}{scope, id, deliveryStreamArn},
		&returns,
	)

	return returns
}

// Import an existing delivery stream from its attributes.
// Experimental.
func DeliveryStream_FromDeliveryStreamAttributes(scope constructs.Construct, id *string, attrs *DeliveryStreamAttributes) IDeliveryStream {
	_init_.Initialize()

	if err := validateDeliveryStream_FromDeliveryStreamAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"fromDeliveryStreamAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing delivery stream from its name.
// Experimental.
func DeliveryStream_FromDeliveryStreamName(scope constructs.Construct, id *string, deliveryStreamName *string) IDeliveryStream {
	_init_.Initialize()

	if err := validateDeliveryStream_FromDeliveryStreamNameParameters(scope, id, deliveryStreamName); err != nil {
		panic(err)
	}
	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"fromDeliveryStreamName",
		[]interface{}{scope, id, deliveryStreamName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DeliveryStream_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDeliveryStream_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DeliveryStream) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := d.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantPutRecordsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricBackupToS3BytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricBackupToS3DataFreshnessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricBackupToS3RecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricIncomingBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricIncomingRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeliveryStream) OnSynthesize(session constructs.ISynthesisSession) {
	if err := d.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DeliveryStream) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeliveryStream) Synthesize(session awscdk.ISynthesisSession) {
	if err := d.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
