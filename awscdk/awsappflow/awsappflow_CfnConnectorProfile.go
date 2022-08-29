package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppFlow::ConnectorProfile`.
//
// The `AWS::AppFlow::ConnectorProfile` resource is an Amazon AppFlow resource type that specifies the configuration profile for an instance of a connector. This includes the provided name, credentials ARN, connection-mode, and so on. The fields that are common to all types of connector profiles are explicitly specified under the `Properties` field. The rest of the connector-specific properties are specified under `Properties/ConnectorProfileConfig` .
//
// > If you want to use AWS CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var credentialsMap interface{}
//   var oAuthCredentials interface{}
//   var profileProperties interface{}
//   var tokenUrlCustomProperties interface{}
//
//   cfnConnectorProfile := awscdk.Aws_appflow.NewCfnConnectorProfile(this, jsii.String("MyCfnConnectorProfile"), &cfnConnectorProfileProps{
//   	connectionMode: jsii.String("connectionMode"),
//   	connectorProfileName: jsii.String("connectorProfileName"),
//   	connectorType: jsii.String("connectorType"),
//
//   	// the properties below are optional
//   	connectorLabel: jsii.String("connectorLabel"),
//   	connectorProfileConfig: &connectorProfileConfigProperty{
//   		connectorProfileCredentials: &connectorProfileCredentialsProperty{
//   			amplitude: &amplitudeConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				secretKey: jsii.String("secretKey"),
//   			},
//   			customConnector: &customConnectorProfileCredentialsProperty{
//   				authenticationType: jsii.String("authenticationType"),
//
//   				// the properties below are optional
//   				apiKey: &apiKeyCredentialsProperty{
//   					apiKey: jsii.String("apiKey"),
//
//   					// the properties below are optional
//   					apiSecretKey: jsii.String("apiSecretKey"),
//   				},
//   				basic: &basicAuthCredentialsProperty{
//   					password: jsii.String("password"),
//   					username: jsii.String("username"),
//   				},
//   				custom: &customAuthCredentialsProperty{
//   					customAuthenticationType: jsii.String("customAuthenticationType"),
//
//   					// the properties below are optional
//   					credentialsMap: credentialsMap,
//   				},
//   				oauth2: &oAuth2CredentialsProperty{
//   					accessToken: jsii.String("accessToken"),
//   					clientId: jsii.String("clientId"),
//   					clientSecret: jsii.String("clientSecret"),
//   					oAuthRequest: &connectorOAuthRequestProperty{
//   						authCode: jsii.String("authCode"),
//   						redirectUri: jsii.String("redirectUri"),
//   					},
//   					refreshToken: jsii.String("refreshToken"),
//   				},
//   			},
//   			datadog: &datadogConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   				applicationKey: jsii.String("applicationKey"),
//   			},
//   			dynatrace: &dynatraceConnectorProfileCredentialsProperty{
//   				apiToken: jsii.String("apiToken"),
//   			},
//   			googleAnalytics: &googleAnalyticsConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			inforNexus: &inforNexusConnectorProfileCredentialsProperty{
//   				accessKeyId: jsii.String("accessKeyId"),
//   				datakey: jsii.String("datakey"),
//   				secretAccessKey: jsii.String("secretAccessKey"),
//   				userId: jsii.String("userId"),
//   			},
//   			marketo: &marketoConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			redshift: &redshiftConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			salesforce: &salesforceConnectorProfileCredentialsProperty{
//   				accessToken: jsii.String("accessToken"),
//   				clientCredentialsArn: jsii.String("clientCredentialsArn"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   				refreshToken: jsii.String("refreshToken"),
//   			},
//   			sapoData: &sAPODataConnectorProfileCredentialsProperty{
//   				basicAuthCredentials: &basicAuthCredentialsProperty{
//   					password: jsii.String("password"),
//   					username: jsii.String("username"),
//   				},
//   				oAuthCredentials: oAuthCredentials,
//   			},
//   			serviceNow: &serviceNowConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			singular: &singularConnectorProfileCredentialsProperty{
//   				apiKey: jsii.String("apiKey"),
//   			},
//   			slack: &slackConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   			snowflake: &snowflakeConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			trendmicro: &trendmicroConnectorProfileCredentialsProperty{
//   				apiSecretKey: jsii.String("apiSecretKey"),
//   			},
//   			veeva: &veevaConnectorProfileCredentialsProperty{
//   				password: jsii.String("password"),
//   				username: jsii.String("username"),
//   			},
//   			zendesk: &zendeskConnectorProfileCredentialsProperty{
//   				clientId: jsii.String("clientId"),
//   				clientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				accessToken: jsii.String("accessToken"),
//   				connectorOAuthRequest: &connectorOAuthRequestProperty{
//   					authCode: jsii.String("authCode"),
//   					redirectUri: jsii.String("redirectUri"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		connectorProfileProperties: &connectorProfilePropertiesProperty{
//   			customConnector: &customConnectorProfilePropertiesProperty{
//   				oAuth2Properties: &oAuth2PropertiesProperty{
//   					oAuth2GrantType: jsii.String("oAuth2GrantType"),
//   					tokenUrl: jsii.String("tokenUrl"),
//   					tokenUrlCustomProperties: tokenUrlCustomProperties,
//   				},
//   				profileProperties: profileProperties,
//   			},
//   			datadog: &datadogConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			dynatrace: &dynatraceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			inforNexus: &inforNexusConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			marketo: &marketoConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			redshift: &redshiftConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				databaseUrl: jsii.String("databaseUrl"),
//   				roleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   			},
//   			salesforce: &salesforceConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   				isSandboxEnvironment: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataConnectorProfilePropertiesProperty{
//   				applicationHostUrl: jsii.String("applicationHostUrl"),
//   				applicationServicePath: jsii.String("applicationServicePath"),
//   				clientNumber: jsii.String("clientNumber"),
//   				logonLanguage: jsii.String("logonLanguage"),
//   				oAuthProperties: &oAuthPropertiesProperty{
//   					authCodeUrl: jsii.String("authCodeUrl"),
//   					oAuthScopes: []*string{
//   						jsii.String("oAuthScopes"),
//   					},
//   					tokenUrl: jsii.String("tokenUrl"),
//   				},
//   				portNumber: jsii.Number(123),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   			},
//   			serviceNow: &serviceNowConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			slack: &slackConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			snowflake: &snowflakeConnectorProfilePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				stage: jsii.String("stage"),
//   				warehouse: jsii.String("warehouse"),
//
//   				// the properties below are optional
//   				accountName: jsii.String("accountName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//   				privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   				region: jsii.String("region"),
//   			},
//   			veeva: &veevaConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   			zendesk: &zendeskConnectorProfilePropertiesProperty{
//   				instanceUrl: jsii.String("instanceUrl"),
//   			},
//   		},
//   	},
//   	kmsArn: jsii.String("kmsArn"),
//   })
//
type CfnConnectorProfile interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the connector profile.
	AttrConnectorProfileArn() *string
	// The Amazon Resource Name (ARN) of the connector profile credentials.
	AttrCredentialsArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Indicates the connection mode and if it is public or private.
	ConnectionMode() *string
	SetConnectionMode(val *string)
	// `AWS::AppFlow::ConnectorProfile.ConnectorLabel`.
	ConnectorLabel() *string
	SetConnectorLabel(val *string)
	// Defines the connector-specific configuration and credentials.
	ConnectorProfileConfig() interface{}
	SetConnectorProfileConfig(val interface{})
	// The name of the connector profile.
	//
	// The name is unique for each `ConnectorProfile` in the AWS account .
	ConnectorProfileName() *string
	SetConnectorProfileName(val *string)
	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType() *string
	SetConnectorType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn() *string
	SetKmsArn(val *string)
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnConnectorProfile
type jsiiProxy_CfnConnectorProfile struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnectorProfile) AttrConnectorProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectorProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) AttrCredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCredentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorProfileConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectorProfileConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectorProfile) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppFlow::ConnectorProfile`.
func NewCfnConnectorProfile(scope constructs.Construct, id *string, props *CfnConnectorProfileProps) CfnConnectorProfile {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectorProfile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppFlow::ConnectorProfile`.
func NewCfnConnectorProfile_Override(c CfnConnectorProfile, scope constructs.Construct, id *string, props *CfnConnectorProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectionMode(val *string) {
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorLabel(val *string) {
	_jsii_.Set(
		j,
		"connectorLabel",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorProfileConfig(val interface{}) {
	_jsii_.Set(
		j,
		"connectorProfileConfig",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorProfileName(val *string) {
	_jsii_.Set(
		j,
		"connectorProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetConnectorType(val *string) {
	_jsii_.Set(
		j,
		"connectorType",
		val,
	)
}

func (j *jsiiProxy_CfnConnectorProfile) SetKmsArn(val *string) {
	_jsii_.Set(
		j,
		"kmsArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnectorProfile_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnectorProfile_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
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
func CfnConnectorProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectorProfile_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appflow.CfnConnectorProfile",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnectorProfile) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectorProfile) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
