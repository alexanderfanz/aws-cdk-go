package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

// A new or imported clustered database.
// Experimental.
type DatabaseClusterBase interface {
	awscdk.Resource
	IDatabaseCluster
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// Access to the network connections.
	// Experimental.
	Connections() awsec2.Connections
	// The engine of this Cluster.
	//
	// May be not known for imported Clusters if it wasn't provided explicitly.
	// Experimental.
	Engine() IClusterEngine
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
	// Endpoints which address each individual replica.
	// Experimental.
	InstanceEndpoints() *[]Endpoint
	// Identifiers of the replicas.
	// Experimental.
	InstanceIdentifiers() *[]*string
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
	// Add a new db proxy to this cluster.
	// Experimental.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
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
	// Renders the secret attachment target specifications.
	// Experimental.
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
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
	// Return the given named metric for this DBCluster.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of deadlocks in the database per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of time that the instance has been running, in seconds.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory, in bytes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of local storage available, in bytes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput received from clients by each instance, in bytes per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput sent to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes for which you are billed.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of storage used by your Aurora DB instance, in bytes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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

// The jsii proxy struct for DatabaseClusterBase
type jsiiProxy_DatabaseClusterBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseCluster
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Engine() IClusterEngine {
	var returns IClusterEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseClusterBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseClusterBase_Override(d DatabaseClusterBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_rds.DatabaseClusterBase",
		[]interface{}{scope, id, props},
		d,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func DatabaseClusterBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseClusterBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseClusterBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseClusterBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseClusterBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.DatabaseClusterBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	if err := d.validateAddProxyParameters(id, options); err != nil {
		panic(err)
	}
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseClusterBase) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (d *jsiiProxy_DatabaseClusterBase) GetResourceNameAttribute(nameAttr *string) *string {
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

func (d *jsiiProxy_DatabaseClusterBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (d *jsiiProxy_DatabaseClusterBase) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricDatabaseConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricDeadlocksParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDeadlocks",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricEngineUptimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricEngineUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFreeableMemoryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFreeLocalStorageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeLocalStorage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricNetworkReceiveThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricNetworkThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricNetworkTransmitThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNetworkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricSnapshotStorageUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSnapshotStorageUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricTotalBackupStorageBilledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTotalBackupStorageBilled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricVolumeBytesUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeBytesUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricVolumeReadIOPsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeReadIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricVolumeWriteIOPsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricVolumeWriteIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseClusterBase) OnSynthesize(session constructs.ISynthesisSession) {
	if err := d.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseClusterBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseClusterBase) Synthesize(session awscdk.ISynthesisSession) {
	if err := d.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseClusterBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseClusterBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

