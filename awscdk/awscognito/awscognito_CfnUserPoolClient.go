package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Cognito::UserPoolClient`.
//
// The `AWS::Cognito::UserPoolClient` resource specifies an Amazon Cognito user pool client.
//
// > If you don't specify a value for a parameter, Amazon Cognito sets it to a default value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolClient := awscdk.Aws_cognito.NewCfnUserPoolClient(this, jsii.String("MyCfnUserPoolClient"), &cfnUserPoolClientProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	accessTokenValidity: jsii.Number(123),
//   	allowedOAuthFlows: []*string{
//   		jsii.String("allowedOAuthFlows"),
//   	},
//   	allowedOAuthFlowsUserPoolClient: jsii.Boolean(false),
//   	allowedOAuthScopes: []*string{
//   		jsii.String("allowedOAuthScopes"),
//   	},
//   	analyticsConfiguration: &analyticsConfigurationProperty{
//   		applicationArn: jsii.String("applicationArn"),
//   		applicationId: jsii.String("applicationId"),
//   		externalId: jsii.String("externalId"),
//   		roleArn: jsii.String("roleArn"),
//   		userDataShared: jsii.Boolean(false),
//   	},
//   	callbackUrLs: []*string{
//   		jsii.String("callbackUrLs"),
//   	},
//   	clientName: jsii.String("clientName"),
//   	defaultRedirectUri: jsii.String("defaultRedirectUri"),
//   	enablePropagateAdditionalUserContextData: jsii.Boolean(false),
//   	enableTokenRevocation: jsii.Boolean(false),
//   	explicitAuthFlows: []*string{
//   		jsii.String("explicitAuthFlows"),
//   	},
//   	generateSecret: jsii.Boolean(false),
//   	idTokenValidity: jsii.Number(123),
//   	logoutUrLs: []*string{
//   		jsii.String("logoutUrLs"),
//   	},
//   	preventUserExistenceErrors: jsii.String("preventUserExistenceErrors"),
//   	readAttributes: []*string{
//   		jsii.String("readAttributes"),
//   	},
//   	refreshTokenValidity: jsii.Number(123),
//   	supportedIdentityProviders: []*string{
//   		jsii.String("supportedIdentityProviders"),
//   	},
//   	tokenValidityUnits: &tokenValidityUnitsProperty{
//   		accessToken: jsii.String("accessToken"),
//   		idToken: jsii.String("idToken"),
//   		refreshToken: jsii.String("refreshToken"),
//   	},
//   	writeAttributes: []*string{
//   		jsii.String("writeAttributes"),
//   	},
//   })
//
type CfnUserPoolClient interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The access token time limit.
	//
	// After this limit expires, your user can't use their access token. To specify the time unit for `AccessTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `AccessTokenValidity` to `10` and `TokenValidityUnits` to `hours` , your user can authorize access with their access token for 10 hours.
	//
	// The default time unit for `AccessTokenValidity` in an API request is hours.
	AccessTokenValidity() *float64
	SetAccessTokenValidity(val *float64)
	// The allowed OAuth flows.
	//
	// - **code** - Use a code grant flow, which provides an authorization code as the response. This code can be exchanged for access tokens with the `/oauth2/token` endpoint.
	// - **implicit** - Issue the access token (and, optionally, ID token, based on scopes) directly to your user.
	// - **client_credentials** - Issue the access token from the `/oauth2/token` endpoint directly to a non-person user using a combination of the client ID and client secret.
	AllowedOAuthFlows() *[]*string
	SetAllowedOAuthFlows(val *[]*string)
	// Set to true if the client is allowed to follow the OAuth protocol when interacting with Amazon Cognito user pools.
	AllowedOAuthFlowsUserPoolClient() interface{}
	SetAllowedOAuthFlowsUserPoolClient(val interface{})
	// The allowed OAuth scopes.
	//
	// Possible values provided by OAuth are `phone` , `email` , `openid` , and `profile` . Possible values provided by AWS are `aws.cognito.signin.user.admin` . Custom scopes created in Resource Servers are also supported.
	AllowedOAuthScopes() *[]*string
	SetAllowedOAuthScopes(val *[]*string)
	// The user pool analytics configuration for collecting metrics and sending them to your Amazon Pinpoint campaign.
	//
	// > In AWS Regions where Amazon Pinpoint isn't available, user pools only support sending events to Amazon Pinpoint projects in AWS Region us-east-1. In Regions where Amazon Pinpoint is available, user pools support sending events to Amazon Pinpoint projects within that same Region.
	AnalyticsConfiguration() interface{}
	SetAnalyticsConfiguration(val interface{})
	AttrClientSecret() *string
	AttrName() *string
	// A list of allowed redirect (callback) URLs for the IdPs.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackUrLs() *[]*string
	SetCallbackUrLs(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The client name for the user pool client you would like to create.
	ClientName() *string
	SetClientName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default redirect URI. Must be in the `CallbackURLs` list.
	//
	// A redirect URI must:
	//
	// - Be an absolute URI.
	// - Be registered with the authorization server.
	// - Not include a fragment component.
	//
	// See [OAuth 2.0 - Redirection Endpoint](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6749#section-3.1.2) .
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectUri() *string
	SetDefaultRedirectUri(val *string)
	// `AWS::Cognito::UserPoolClient.EnablePropagateAdditionalUserContextData`.
	EnablePropagateAdditionalUserContextData() interface{}
	SetEnablePropagateAdditionalUserContextData(val interface{})
	// Activates or deactivates token revocation. For more information about revoking tokens, see [RevokeToken](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_RevokeToken.html) .
	//
	// If you don't include this parameter, token revocation is automatically activated for the new user pool client.
	EnableTokenRevocation() interface{}
	SetEnableTokenRevocation(val interface{})
	// The authentication flows that are supported by the user pool clients.
	//
	// Flow names without the `ALLOW_` prefix are no longer supported, in favor of new names with the `ALLOW_` prefix.
	//
	// > Values with `ALLOW_` prefix must be used only along with the `ALLOW_` prefix.
	//
	// Valid values include:
	//
	// - `ALLOW_ADMIN_USER_PASSWORD_AUTH` : Enable admin based user password authentication flow `ADMIN_USER_PASSWORD_AUTH` . This setting replaces the `ADMIN_NO_SRP_AUTH` setting. With this authentication flow, Amazon Cognito receives the password in the request instead of using the Secure Remote Password (SRP) protocol to verify passwords.
	// - `ALLOW_CUSTOM_AUTH` : Enable AWS Lambda trigger based authentication.
	// - `ALLOW_USER_PASSWORD_AUTH` : Enable user password-based authentication. In this flow, Amazon Cognito receives the password in the request instead of using the SRP protocol to verify passwords.
	// - `ALLOW_USER_SRP_AUTH` : Enable SRP-based authentication.
	// - `ALLOW_REFRESH_TOKEN_AUTH` : Enable authflow to refresh tokens.
	//
	// If you don't specify a value for `ExplicitAuthFlows` , your app client activates the `ALLOW_USER_SRP_AUTH` and `ALLOW_CUSTOM_AUTH` authentication flows.
	ExplicitAuthFlows() *[]*string
	SetExplicitAuthFlows(val *[]*string)
	// Boolean to specify whether you want to generate a secret for the user pool client being created.
	GenerateSecret() interface{}
	SetGenerateSecret(val interface{})
	// The ID token time limit.
	//
	// After this limit expires, your user can't use their ID token. To specify the time unit for `IdTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `IdTokenValidity` as `10` and `TokenValidityUnits` as `hours` , your user can authenticate their session with their ID token for 10 hours.
	//
	// The default time unit for `AccessTokenValidity` in an API request is hours.
	IdTokenValidity() *float64
	SetIdTokenValidity(val *float64)
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
	// A list of allowed logout URLs for the IdPs.
	LogoutUrLs() *[]*string
	SetLogoutUrLs(val *[]*string)
	// The tree node.
	Node() constructs.Node
	// Use this setting to choose which errors and responses are returned by Cognito APIs during authentication, account confirmation, and password recovery when the user does not exist in the user pool.
	//
	// When set to `ENABLED` and the user does not exist, authentication returns an error indicating either the username or password was incorrect, and account confirmation and password recovery return a response indicating a code was sent to a simulated destination. When set to `LEGACY` , those APIs will return a `UserNotFoundException` exception if the user does not exist in the user pool.
	PreventUserExistenceErrors() *string
	SetPreventUserExistenceErrors(val *string)
	// The read attributes.
	ReadAttributes() *[]*string
	SetReadAttributes(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The refresh token time limit.
	//
	// After this limit expires, your user can't use their refresh token. To specify the time unit for `RefreshTokenValidity` as `seconds` , `minutes` , `hours` , or `days` , set a `TokenValidityUnits` value in your API request.
	//
	// For example, when you set `RefreshTokenValidity` as `10` and `TokenValidityUnits` as `days` , your user can refresh their session and retrieve new access and ID tokens for 10 days.
	//
	// The default time unit for `RefreshTokenValidity` in an API request is days. You can't set `RefreshTokenValidity` to 0. If you do, Amazon Cognito overrides the value with the default value of 30 days.
	RefreshTokenValidity() *float64
	SetRefreshTokenValidity(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of provider names for the IdPs that this client supports.
	//
	// The following are supported: `COGNITO` , `Facebook` , `Google` `LoginWithAmazon` , and the names of your own SAML and OIDC providers.
	SupportedIdentityProviders() *[]*string
	SetSupportedIdentityProviders(val *[]*string)
	// The units in which the validity times are represented.
	//
	// The default unit for RefreshToken is days, and default for ID and access tokens are hours.
	TokenValidityUnits() interface{}
	SetTokenValidityUnits(val interface{})
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
	// The user pool ID for the user pool where you want to create a user pool client.
	UserPoolId() *string
	SetUserPoolId(val *string)
	// The user pool attributes that the app client can write to.
	//
	// If your app client allows users to sign in through an IdP, this array must include all attributes that you have mapped to IdP attributes. Amazon Cognito updates mapped attributes when users sign in to your application through an IdP. If your app client does not have write access to a mapped attribute, Amazon Cognito throws an error when it tries to update the attribute. For more information, see [Specifying IdP Attribute Mappings for Your user pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	WriteAttributes() *[]*string
	SetWriteAttributes(val *[]*string)
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

// The jsii proxy struct for CfnUserPoolClient
type jsiiProxy_CfnUserPoolClient struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnUserPoolClient) AccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AllowedOAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AllowedOAuthFlowsUserPoolClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedOAuthFlowsUserPoolClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AllowedOAuthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOAuthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AnalyticsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AttrClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CallbackUrLs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"callbackUrLs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) ClientName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) DefaultRedirectUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRedirectUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) EnablePropagateAdditionalUserContextData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePropagateAdditionalUserContextData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) EnableTokenRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTokenRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) ExplicitAuthFlows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"explicitAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) GenerateSecret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) IdTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) LogoutUrLs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logoutUrLs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) PreventUserExistenceErrors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preventUserExistenceErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) ReadAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) RefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) SupportedIdentityProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedIdentityProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) TokenValidityUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenValidityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolClient) WriteAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"writeAttributes",
		&returns,
	)
	return returns
}


// Create a new `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClient(scope constructs.Construct, id *string, props *CfnUserPoolClientProps) CfnUserPoolClient {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolClient{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.CfnUserPoolClient",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Cognito::UserPoolClient`.
func NewCfnUserPoolClient_Override(c CfnUserPoolClient, scope constructs.Construct, id *string, props *CfnUserPoolClientProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.CfnUserPoolClient",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAccessTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"accessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAllowedOAuthFlows(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAllowedOAuthFlowsUserPoolClient(val interface{}) {
	_jsii_.Set(
		j,
		"allowedOAuthFlowsUserPoolClient",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAllowedOAuthScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOAuthScopes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetAnalyticsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"analyticsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetCallbackUrLs(val *[]*string) {
	_jsii_.Set(
		j,
		"callbackUrLs",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetClientName(val *string) {
	_jsii_.Set(
		j,
		"clientName",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetDefaultRedirectUri(val *string) {
	_jsii_.Set(
		j,
		"defaultRedirectUri",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetEnablePropagateAdditionalUserContextData(val interface{}) {
	_jsii_.Set(
		j,
		"enablePropagateAdditionalUserContextData",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetEnableTokenRevocation(val interface{}) {
	_jsii_.Set(
		j,
		"enableTokenRevocation",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetExplicitAuthFlows(val *[]*string) {
	_jsii_.Set(
		j,
		"explicitAuthFlows",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetGenerateSecret(val interface{}) {
	_jsii_.Set(
		j,
		"generateSecret",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetIdTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"idTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetLogoutUrLs(val *[]*string) {
	_jsii_.Set(
		j,
		"logoutUrLs",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetPreventUserExistenceErrors(val *string) {
	_jsii_.Set(
		j,
		"preventUserExistenceErrors",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetReadAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"readAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetRefreshTokenValidity(val *float64) {
	_jsii_.Set(
		j,
		"refreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetSupportedIdentityProviders(val *[]*string) {
	_jsii_.Set(
		j,
		"supportedIdentityProviders",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetTokenValidityUnits(val interface{}) {
	_jsii_.Set(
		j,
		"tokenValidityUnits",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CfnUserPoolClient) SetWriteAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"writeAttributes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnUserPoolClient_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.CfnUserPoolClient",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnUserPoolClient_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.CfnUserPoolClient",
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
func CfnUserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.CfnUserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolClient_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cognito.CfnUserPoolClient",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnUserPoolClient) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnUserPoolClient) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
