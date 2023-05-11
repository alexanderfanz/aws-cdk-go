package awsgrafana

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgrafana/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Grafana::Workspace`.
//
// Specifies a *workspace* . In a workspace, you can create Grafana dashboards and visualizations to analyze your metrics, logs, and traces. You don't have to build, package, or deploy any hardware to run the Grafana server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspace := awscdk.Aws_grafana.NewCfnWorkspace(this, jsii.String("MyCfnWorkspace"), &CfnWorkspaceProps{
//   	AccountAccessType: jsii.String("accountAccessType"),
//   	AuthenticationProviders: []*string{
//   		jsii.String("authenticationProviders"),
//   	},
//   	PermissionType: jsii.String("permissionType"),
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	DataSources: []*string{
//   		jsii.String("dataSources"),
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	NetworkAccessControl: &NetworkAccessControlProperty{
//   		PrefixListIds: []*string{
//   			jsii.String("prefixListIds"),
//   		},
//   		VpceIds: []*string{
//   			jsii.String("vpceIds"),
//   		},
//   	},
//   	NotificationDestinations: []*string{
//   		jsii.String("notificationDestinations"),
//   	},
//   	OrganizationalUnits: []*string{
//   		jsii.String("organizationalUnits"),
//   	},
//   	OrganizationRoleName: jsii.String("organizationRoleName"),
//   	RoleArn: jsii.String("roleArn"),
//   	SamlConfiguration: &SamlConfigurationProperty{
//   		IdpMetadata: &IdpMetadataProperty{
//   			Url: jsii.String("url"),
//   			Xml: jsii.String("xml"),
//   		},
//
//   		// the properties below are optional
//   		AllowedOrganizations: []*string{
//   			jsii.String("allowedOrganizations"),
//   		},
//   		AssertionAttributes: &AssertionAttributesProperty{
//   			Email: jsii.String("email"),
//   			Groups: jsii.String("groups"),
//   			Login: jsii.String("login"),
//   			Name: jsii.String("name"),
//   			Org: jsii.String("org"),
//   			Role: jsii.String("role"),
//   		},
//   		LoginValidityDuration: jsii.Number(123),
//   		RoleValues: &RoleValuesProperty{
//   			Admin: []*string{
//   				jsii.String("admin"),
//   			},
//   			Editor: []*string{
//   				jsii.String("editor"),
//   			},
//   		},
//   	},
//   	StackSetName: jsii.String("stackSetName"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnWorkspace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in other accounts in the same organization.
	//
	// If this is `ORGANIZATION` , the `OrganizationalUnits` parameter specifies which organizational units the workspace can access.
	AccountAccessType() *string
	SetAccountAccessType(val *string)
	// The date that the workspace was created.
	//
	// Type: Timestamp.
	AttrCreationTimestamp() *string
	// The URL that users can use to access the Grafana console in the workspace.
	//
	// Type: String.
	AttrEndpoint() *string
	// The version of Grafana supported in this workspace.
	//
	// Type: String.
	AttrGrafanaVersion() *string
	// The unique ID of this workspace.
	//
	// Type: String.
	AttrId() *string
	// The most recent date that the workspace was modified.
	//
	// Type: Timestamp.
	AttrModificationTimestamp() *string
	// Specifies whether the workspace's SAML configuration is complete.
	//
	// Valid values: `CONFIGURED | NOT_CONFIGURED`
	//
	// Type: String.
	AttrSamlConfigurationStatus() *string
	// The ID of the IAM Identity Center-managed application that is created by Amazon Managed Grafana .
	//
	// Type: String.
	AttrSsoClientId() *string
	// The current status of the workspace.
	//
	// Valid values: `ACTIVE | CREATING | DELETING | FAILED | UPDATING | UPGRADING | DELETION_FAILED | CREATION_FAILED | UPDATE_FAILED | UPGRADE_FAILED | LICENSE_REMOVAL_FAILED`
	//
	// Type: String.
	AttrStatus() *string
	// Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center (successor to AWS Single Sign-On) , or both to authenticate users for using the Grafana console within a workspace. For more information, see [User authentication in Amazon Managed Grafana](https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html) .
	AuthenticationProviders() *[]*string
	SetAuthenticationProviders(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
	ClientToken() *string
	SetClientToken(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies the AWS data sources that have been configured to have IAM roles and permissions created to allow Amazon Managed Grafana to read data from these sources.
	//
	// This list is only used when the workspace was created through the AWS console, and the `permissionType` is `SERVICE_MANAGED` .
	DataSources() *[]*string
	SetDataSources(val *[]*string)
	// The user-defined description of the workspace.
	Description() *string
	SetDescription(val *string)
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
	// The name of the workspace.
	Name() *string
	SetName(val *string)
	// The configuration settings for network access to your workspace.
	NetworkAccessControl() interface{}
	SetNetworkAccessControl(val interface{})
	// The tree node.
	Node() constructs.Node
	// The AWS notification channels that Amazon Managed Grafana can automatically create IAM roles and permissions for, to allow Amazon Managed Grafana to use these channels.
	NotificationDestinations() *[]*string
	SetNotificationDestinations(val *[]*string)
	// Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.
	OrganizationalUnits() *[]*string
	SetOrganizationalUnits(val *[]*string)
	// The name of the IAM role that is used to access resources through Organizations .
	OrganizationRoleName() *string
	SetOrganizationRoleName(val *string)
	// If this is `SERVICE_MANAGED` , and the workplace was created through the Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the IAM roles and provisions the permissions that the workspace needs to use AWS data sources and notification channels.
	//
	// If this is `CUSTOMER_MANAGED` , you must manage those roles and permissions yourself.
	//
	// If you are working with a workspace in a member account of an organization and that account is not a delegated administrator account, and you want the workspace to access data sources in other AWS accounts in the organization, this parameter must be set to `CUSTOMER_MANAGED` .
	//
	// For more information about converting between customer and service managed, see [Managing permissions for data sources and notification channels](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-datasource-and-notification.html) . For more information about the roles and permissions that must be managed for customer managed workspaces, see [Amazon Managed Grafana permissions and policies for AWS data sources and notification channels](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	PermissionType() *string
	SetPermissionType(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The IAM role that grants permissions to the AWS resources that the workspace will view data from.
	//
	// This role must already exist.
	RoleArn() *string
	SetRoleArn(val *string)
	// If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the `Admin` and `Editor` roles in the workspace.
	SamlConfiguration() interface{}
	SetSamlConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the AWS CloudFormation stack set that is used to generate IAM roles to be used for this workspace.
	StackSetName() *string
	SetStackSetName(val *string)
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
	// The configuration for connecting to data sources in a private VPC ( Amazon Virtual Private Cloud ).
	VpcConfiguration() interface{}
	SetVpcConfiguration(val interface{})
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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

// The jsii proxy struct for CfnWorkspace
type jsiiProxy_CfnWorkspace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkspace) AccountAccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountAccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrCreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrGrafanaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGrafanaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrModificationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModificationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrSamlConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSamlConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrSsoClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSsoClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) AuthenticationProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) DataSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) NetworkAccessControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) NotificationDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) OrganizationalUnits() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) OrganizationRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) PermissionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) SamlConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"samlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) StackSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspace) VpcConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::Grafana::Workspace`.
func NewCfnWorkspace(scope constructs.Construct, id *string, props *CfnWorkspaceProps) CfnWorkspace {
	_init_.Initialize()

	if err := validateNewCfnWorkspaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_grafana.CfnWorkspace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Grafana::Workspace`.
func NewCfnWorkspace_Override(c CfnWorkspace, scope constructs.Construct, id *string, props *CfnWorkspaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_grafana.CfnWorkspace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetAccountAccessType(val *string) {
	if err := j.validateSetAccountAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountAccessType",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetAuthenticationProviders(val *[]*string) {
	if err := j.validateSetAuthenticationProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationProviders",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetClientToken(val *string) {
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetDataSources(val *[]*string) {
	_jsii_.Set(
		j,
		"dataSources",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetNetworkAccessControl(val interface{}) {
	if err := j.validateSetNetworkAccessControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAccessControl",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetNotificationDestinations(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationDestinations",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetOrganizationalUnits(val *[]*string) {
	_jsii_.Set(
		j,
		"organizationalUnits",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetOrganizationRoleName(val *string) {
	_jsii_.Set(
		j,
		"organizationRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetPermissionType(val *string) {
	if err := j.validateSetPermissionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionType",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetSamlConfiguration(val interface{}) {
	if err := j.validateSetSamlConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetStackSetName(val *string) {
	_jsii_.Set(
		j,
		"stackSetName",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspace)SetVpcConfiguration(val interface{}) {
	if err := j.validateSetVpcConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnWorkspace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspace_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_grafana.CfnWorkspace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnWorkspace_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspace_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_grafana.CfnWorkspace",
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
func CfnWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_grafana.CfnWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_grafana.CfnWorkspace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspace) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWorkspace) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWorkspace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWorkspace) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnWorkspace) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnWorkspace) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWorkspace) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkspace) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnWorkspace) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnWorkspace) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspace) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
