package awscloudformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::CloudFormation::StackSet`.
//
// The `AWS::CloudFormation::StackSet` enables you to provision stacks into AWS accounts and across Regions by using a single CloudFormation template. In the stack set, you specify the template to use, in addition to any parameters and capabilities that the template requires.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var managedExecution interface{}
//
//   cfnStackSet := awscdk.Aws_cloudformation.NewCfnStackSet(this, jsii.String("MyCfnStackSet"), &CfnStackSetProps{
//   	PermissionModel: jsii.String("permissionModel"),
//   	StackSetName: jsii.String("stackSetName"),
//
//   	// the properties below are optional
//   	AdministrationRoleArn: jsii.String("administrationRoleArn"),
//   	AutoDeployment: &AutoDeploymentProperty{
//   		Enabled: jsii.Boolean(false),
//   		RetainStacksOnAccountRemoval: jsii.Boolean(false),
//   	},
//   	CallAs: jsii.String("callAs"),
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRoleName: jsii.String("executionRoleName"),
//   	ManagedExecution: managedExecution,
//   	OperationPreferences: &OperationPreferencesProperty{
//   		FailureToleranceCount: jsii.Number(123),
//   		FailureTolerancePercentage: jsii.Number(123),
//   		MaxConcurrentCount: jsii.Number(123),
//   		MaxConcurrentPercentage: jsii.Number(123),
//   		RegionConcurrencyType: jsii.String("regionConcurrencyType"),
//   		RegionOrder: []*string{
//   			jsii.String("regionOrder"),
//   		},
//   	},
//   	Parameters: []interface{}{
//   		&ParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	StackInstancesGroup: []interface{}{
//   		&StackInstancesProperty{
//   			DeploymentTargets: &DeploymentTargetsProperty{
//   				AccountFilterType: jsii.String("accountFilterType"),
//   				Accounts: []*string{
//   					jsii.String("accounts"),
//   				},
//   				OrganizationalUnitIds: []*string{
//   					jsii.String("organizationalUnitIds"),
//   				},
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//
//   			// the properties below are optional
//   			ParameterOverrides: []interface{}{
//   				&ParameterProperty{
//   					ParameterKey: jsii.String("parameterKey"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateUrl: jsii.String("templateUrl"),
//   })
//
type CfnStackSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack set.
	//
	// Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
	//
	// Use customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account. For more information, see [Prerequisites: Granting Permissions for Stack Set Operations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) in the *AWS CloudFormation User Guide* .
	//
	// *Minimum* : `20`
	//
	// *Maximum* : `2048`.
	AdministrationRoleArn() *string
	SetAdministrationRoleArn(val *string)
	// The ID of the stack that you're creating.
	AttrStackSetId() *string
	// [ `Service-managed` permissions] Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
	AutoDeployment() interface{}
	SetAutoDeployment(val interface{})
	// [Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.
	//
	// By default, `SELF` is specified. Use `SELF` for stack sets with self-managed permissions.
	//
	// - To create a stack set with service-managed permissions while signed in to the management account, specify `SELF` .
	// - To create a stack set with service-managed permissions while signed in to a delegated administrator account, specify `DELEGATED_ADMIN` .
	//
	// Your AWS account must be registered as a delegated admin in the management account. For more information, see [Register a delegated administrator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html) in the *AWS CloudFormation User Guide* .
	//
	// Stack sets with service-managed permissions are created in the management account, including stack sets that are created by delegated administrators.
	//
	// *Valid Values* : `SELF` | `DELEGATED_ADMIN`.
	CallAs() *string
	SetCallAs(val *string)
	// The capabilities that are allowed in the stack set.
	//
	// Some stack set templates might include resources that can affect permissions in your AWS account —for example, by creating new AWS Identity and Access Management ( IAM ) users. For more information, see [Acknowledging IAM Resources in AWS CloudFormation Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities) .
	Capabilities() *[]*string
	SetCapabilities(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the stack set.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	Description() *string
	SetDescription(val *string)
	// The name of the IAM execution role to use to create the stack set.
	//
	// If you don't specify an execution role, AWS CloudFormation uses the `AWSCloudFormationStackSetExecutionRole` role for the stack set operation.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
	//
	// *Pattern* : `[a-zA-Z_0-9+=,.@-]+`
	ExecutionRoleName() *string
	SetExecutionRoleName(val *string)
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
	// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// When active, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.
	//
	// > If there are already running or queued operations, StackSets queues all incoming operations even if they are non-conflicting.
	// >
	// > You can't modify your stack set's execution configuration while there are running or queued operations for that stack set.
	//
	// When inactive (default), StackSets performs one operation at a time in request order.
	ManagedExecution() interface{}
	SetManagedExecution(val interface{})
	// The tree node.
	Node() constructs.Node
	// The user-specified preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences() interface{}
	SetOperationPreferences(val interface{})
	// The input parameters for the stack set template.
	Parameters() interface{}
	SetParameters(val interface{})
	// Describes how the IAM roles required for stack set operations are created.
	//
	// - With `SELF_MANAGED` permissions, you must create the administrator and execution roles required to deploy to target accounts. For more information, see [Grant Self-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html) .
	// - With `SERVICE_MANAGED` permissions, StackSets automatically creates the IAM roles required to deploy to accounts managed by AWS Organizations . For more information, see [Grant Service-Managed Stack Set Permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html) .
	//
	// *Allowed Values* : `SERVICE_MANAGED` | `SELF_MANAGED`
	//
	// > The `PermissionModel` property is required.
	PermissionModel() *string
	SetPermissionModel(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A group of stack instances with parameters in some specific accounts and Regions.
	StackInstancesGroup() interface{}
	SetStackInstancesGroup(val interface{})
	// The name to associate with the stack set.
	//
	// The name must be unique in the Region where you create your stack set.
	//
	// *Maximum* : `128`
	//
	// *Pattern* : `^[a-zA-Z][a-zA-Z0-9-]{0,127}$`
	//
	// > The `StackSetName` property is required.
	StackSetName() *string
	SetStackSetName(val *string)
	// The key-value pairs to associate with this stack set and the stacks created from it.
	//
	// AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.
	Tags() awscdk.TagManager
	// The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both. Dynamic references in the `TemplateBody` may not work correctly in all cases. It's recommended to pass templates containing dynamic references through `TemplateUrl` instead.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `51200`.
	TemplateBody() *string
	SetTemplateBody(val *string)
	// Location of file containing the template body.
	//
	// The URL must point to a template (max size: 460,800 bytes) that's located in an Amazon S3 bucket.
	//
	// You must include either `TemplateURL` or `TemplateBody` in a StackSet, but you can't use both.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	TemplateUrl() *string
	SetTemplateUrl(val *string)
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
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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

// The jsii proxy struct for CfnStackSet
type jsiiProxy_CfnStackSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStackSet) AdministrationRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrationRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) AttrStackSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStackSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) AutoDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CallAs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callAs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Capabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) ExecutionRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) ManagedExecution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedExecution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) OperationPreferences() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) PermissionModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) StackInstancesGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stackInstancesGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) TemplateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStackSet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudFormation::StackSet`.
func NewCfnStackSet(scope constructs.Construct, id *string, props *CfnStackSetProps) CfnStackSet {
	_init_.Initialize()

	if err := validateNewCfnStackSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStackSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudformation.CfnStackSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudFormation::StackSet`.
func NewCfnStackSet_Override(c CfnStackSet, scope constructs.Construct, id *string, props *CfnStackSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudformation.CfnStackSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStackSet)SetAdministrationRoleArn(val *string) {
	_jsii_.Set(
		j,
		"administrationRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetAutoDeployment(val interface{}) {
	if err := j.validateSetAutoDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeployment",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetCallAs(val *string) {
	_jsii_.Set(
		j,
		"callAs",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetCapabilities(val *[]*string) {
	_jsii_.Set(
		j,
		"capabilities",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetExecutionRoleName(val *string) {
	_jsii_.Set(
		j,
		"executionRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetManagedExecution(val interface{}) {
	if err := j.validateSetManagedExecutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedExecution",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetOperationPreferences(val interface{}) {
	if err := j.validateSetOperationPreferencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationPreferences",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetParameters(val interface{}) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetPermissionModel(val *string) {
	if err := j.validateSetPermissionModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionModel",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetStackInstancesGroup(val interface{}) {
	if err := j.validateSetStackInstancesGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackInstancesGroup",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetStackSetName(val *string) {
	if err := j.validateSetStackSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetName",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetTemplateBody(val *string) {
	_jsii_.Set(
		j,
		"templateBody",
		val,
	)
}

func (j *jsiiProxy_CfnStackSet)SetTemplateUrl(val *string) {
	_jsii_.Set(
		j,
		"templateUrl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStackSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStackSet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudformation.CfnStackSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStackSet_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnStackSet_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudformation.CfnStackSet",
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
func CfnStackSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStackSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudformation.CfnStackSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStackSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudformation.CfnStackSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStackSet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStackSet) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStackSet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStackSet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStackSet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStackSet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStackSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStackSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStackSet) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStackSet) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStackSet) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStackSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnStackSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStackSet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

