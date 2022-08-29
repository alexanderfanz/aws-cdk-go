package awsssm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::SSM::MaintenanceWindowTask`.
//
// The `AWS::SSM::MaintenanceWindowTask` resource defines information about a task for an AWS Systems Manager maintenance window. For more information, see [RegisterTaskWithMaintenanceWindow](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_RegisterTaskWithMaintenanceWindow.html) in the *AWS Systems Manager API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var taskParameters interface{}
//
//   cfnMaintenanceWindowTask := awscdk.Aws_ssm.NewCfnMaintenanceWindowTask(this, jsii.String("MyCfnMaintenanceWindowTask"), &cfnMaintenanceWindowTaskProps{
//   	priority: jsii.Number(123),
//   	taskArn: jsii.String("taskArn"),
//   	taskType: jsii.String("taskType"),
//   	windowId: jsii.String("windowId"),
//
//   	// the properties below are optional
//   	cutoffBehavior: jsii.String("cutoffBehavior"),
//   	description: jsii.String("description"),
//   	loggingInfo: &loggingInfoProperty{
//   		region: jsii.String("region"),
//   		s3Bucket: jsii.String("s3Bucket"),
//
//   		// the properties below are optional
//   		s3Prefix: jsii.String("s3Prefix"),
//   	},
//   	maxConcurrency: jsii.String("maxConcurrency"),
//   	maxErrors: jsii.String("maxErrors"),
//   	name: jsii.String("name"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			key: jsii.String("key"),
//   			values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	taskInvocationParameters: &taskInvocationParametersProperty{
//   		maintenanceWindowAutomationParameters: &maintenanceWindowAutomationParametersProperty{
//   			documentVersion: jsii.String("documentVersion"),
//   			parameters: parameters,
//   		},
//   		maintenanceWindowLambdaParameters: &maintenanceWindowLambdaParametersProperty{
//   			clientContext: jsii.String("clientContext"),
//   			payload: jsii.String("payload"),
//   			qualifier: jsii.String("qualifier"),
//   		},
//   		maintenanceWindowRunCommandParameters: &maintenanceWindowRunCommandParametersProperty{
//   			cloudWatchOutputConfig: &cloudWatchOutputConfigProperty{
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				cloudWatchOutputEnabled: jsii.Boolean(false),
//   			},
//   			comment: jsii.String("comment"),
//   			documentHash: jsii.String("documentHash"),
//   			documentHashType: jsii.String("documentHashType"),
//   			documentVersion: jsii.String("documentVersion"),
//   			notificationConfig: &notificationConfigProperty{
//   				notificationArn: jsii.String("notificationArn"),
//
//   				// the properties below are optional
//   				notificationEvents: []*string{
//   					jsii.String("notificationEvents"),
//   				},
//   				notificationType: jsii.String("notificationType"),
//   			},
//   			outputS3BucketName: jsii.String("outputS3BucketName"),
//   			outputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			parameters: parameters,
//   			serviceRoleArn: jsii.String("serviceRoleArn"),
//   			timeoutSeconds: jsii.Number(123),
//   		},
//   		maintenanceWindowStepFunctionsParameters: &maintenanceWindowStepFunctionsParametersProperty{
//   			input: jsii.String("input"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	taskParameters: taskParameters,
//   })
//
type CfnMaintenanceWindowTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The specification for whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached.
	CutoffBehavior() *string
	SetCutoffBehavior(val *string)
	// A description of the task.
	Description() *string
	SetDescription(val *string)
	// Information about an Amazon S3 bucket to write task-level logs to.
	//
	// > `LoggingInfo` has been deprecated. To specify an Amazon S3 bucket to contain logs, instead use the `OutputS3BucketName` and `OutputS3KeyPrefix` options in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [AWS Systems Manager MaintenanceWindowTask TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) .
	LoggingInfo() interface{}
	SetLoggingInfo(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The maximum number of targets this task can be run for, in parallel.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxConcurrency() *string
	SetMaxConcurrency(val *string)
	// The maximum number of errors allowed before this task stops being scheduled.
	//
	// > Although this element is listed as "Required: No", a value can be omitted only when you are registering or updating a [targetless task](https://docs.aws.amazon.com/systems-manager/latest/userguide/maintenance-windows-targetless-tasks.html) You must provide a value in all other cases.
	// >
	// > For maintenance window tasks without a target specified, you can't supply a value for this option. Instead, the system inserts a placeholder value of `1` . This value doesn't affect the running of your task.
	MaxErrors() *string
	SetMaxErrors(val *string)
	// The task name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The priority of the task in the maintenance window.
	//
	// The lower the number, the higher the priority. Tasks that have the same priority are scheduled in parallel.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The targets, either instances or window target IDs.
	//
	// - Specify instances using `Key=InstanceIds,Values= *instanceid1* , *instanceid2*` .
	// - Specify window target IDs using `Key=WindowTargetIds,Values= *window-target-id-1* , *window-target-id-2*` .
	Targets() interface{}
	SetTargets(val interface{})
	// The resource that the task uses during execution.
	//
	// For `RUN_COMMAND` and `AUTOMATION` task types, `TaskArn` is the SSM document name or Amazon Resource Name (ARN).
	//
	// For `LAMBDA` tasks, `TaskArn` is the function name or ARN.
	//
	// For `STEP_FUNCTIONS` tasks, `TaskArn` is the state machine ARN.
	TaskArn() *string
	SetTaskArn(val *string)
	// The parameters to pass to the task when it runs.
	//
	// Populate only the fields that match the task type. All other fields should be empty.
	//
	// > When you update a maintenance window task that has options specified in `TaskInvocationParameters` , you must provide again all the `TaskInvocationParameters` values that you want to retain. The values you do not specify again are removed. For example, suppose that when you registered a Run Command task, you specified `TaskInvocationParameters` values for `Comment` , `NotificationConfig` , and `OutputS3BucketName` . If you update the maintenance window task and specify only a different `OutputS3BucketName` value, the values for `Comment` and `NotificationConfig` are removed.
	TaskInvocationParameters() interface{}
	SetTaskInvocationParameters(val interface{})
	// The parameters to pass to the task when it runs.
	//
	// > `TaskParameters` has been deprecated. To specify parameters to pass to a task when it runs, instead use the `Parameters` option in the `TaskInvocationParameters` structure. For information about how Systems Manager handles these options for the supported maintenance window task types, see [MaintenanceWindowTaskInvocationParameters](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_MaintenanceWindowTaskInvocationParameters.html) .
	TaskParameters() interface{}
	SetTaskParameters(val interface{})
	// The type of task.
	//
	// Valid values: `RUN_COMMAND` , `AUTOMATION` , `LAMBDA` , `STEP_FUNCTIONS` .
	TaskType() *string
	SetTaskType(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the maintenance window where the task is registered.
	WindowId() *string
	SetWindowId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMaintenanceWindowTask
type jsiiProxy_CfnMaintenanceWindowTask struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) CutoffBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cutoffBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) LoggingInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) MaxConcurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) MaxErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskInvocationParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskInvocationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) TaskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) WindowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowId",
		&returns,
	)
	return returns
}


// Create a new `AWS::SSM::MaintenanceWindowTask`.
func NewCfnMaintenanceWindowTask(scope constructs.Construct, id *string, props *CfnMaintenanceWindowTaskProps) CfnMaintenanceWindowTask {
	_init_.Initialize()

	j := jsiiProxy_CfnMaintenanceWindowTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ssm.CfnMaintenanceWindowTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SSM::MaintenanceWindowTask`.
func NewCfnMaintenanceWindowTask_Override(c CfnMaintenanceWindowTask, scope constructs.Construct, id *string, props *CfnMaintenanceWindowTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ssm.CfnMaintenanceWindowTask",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetCutoffBehavior(val *string) {
	_jsii_.Set(
		j,
		"cutoffBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetLoggingInfo(val interface{}) {
	_jsii_.Set(
		j,
		"loggingInfo",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetMaxConcurrency(val *string) {
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetMaxErrors(val *string) {
	_jsii_.Set(
		j,
		"maxErrors",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskArn(val *string) {
	_jsii_.Set(
		j,
		"taskArn",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskInvocationParameters(val interface{}) {
	_jsii_.Set(
		j,
		"taskInvocationParameters",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskParameters(val interface{}) {
	_jsii_.Set(
		j,
		"taskParameters",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetTaskType(val *string) {
	_jsii_.Set(
		j,
		"taskType",
		val,
	)
}

func (j *jsiiProxy_CfnMaintenanceWindowTask) SetWindowId(val *string) {
	_jsii_.Set(
		j,
		"windowId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMaintenanceWindowTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.CfnMaintenanceWindowTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMaintenanceWindowTask_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.CfnMaintenanceWindowTask",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnMaintenanceWindowTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ssm.CfnMaintenanceWindowTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMaintenanceWindowTask_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ssm.CfnMaintenanceWindowTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMaintenanceWindowTask) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
