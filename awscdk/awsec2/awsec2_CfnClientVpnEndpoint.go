package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::EC2::ClientVpnEndpoint`.
//
// Specifies a Client VPN endpoint. A Client VPN endpoint is the resource you create and configure to enable and manage client VPN sessions. It is the destination endpoint at which all client VPN sessions are terminated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnEndpoint := awscdk.Aws_ec2.NewCfnClientVpnEndpoint(this, jsii.String("MyCfnClientVpnEndpoint"), &cfnClientVpnEndpointProps{
//   	authenticationOptions: []interface{}{
//   		&clientAuthenticationRequestProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			activeDirectory: &directoryServiceAuthenticationRequestProperty{
//   				directoryId: jsii.String("directoryId"),
//   			},
//   			federatedAuthentication: &federatedAuthenticationRequestProperty{
//   				samlProviderArn: jsii.String("samlProviderArn"),
//
//   				// the properties below are optional
//   				selfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   			},
//   			mutualAuthentication: &certificateAuthenticationRequestProperty{
//   				clientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   			},
//   		},
//   	},
//   	clientCidrBlock: jsii.String("clientCidrBlock"),
//   	connectionLogOptions: &connectionLogOptionsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		cloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   		cloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   	},
//   	serverCertificateArn: jsii.String("serverCertificateArn"),
//
//   	// the properties below are optional
//   	clientConnectOptions: &clientConnectOptionsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	clientLoginBannerOptions: &clientLoginBannerOptionsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		bannerText: jsii.String("bannerText"),
//   	},
//   	description: jsii.String("description"),
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	selfServicePortal: jsii.String("selfServicePortal"),
//   	sessionTimeoutHours: jsii.Number(123),
//   	splitTunnel: jsii.Boolean(false),
//   	tagSpecifications: []interface{}{
//   		&tagSpecificationProperty{
//   			resourceType: jsii.String("resourceType"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	transportProtocol: jsii.String("transportProtocol"),
//   	vpcId: jsii.String("vpcId"),
//   	vpnPort: jsii.Number(123),
//   })
//
type CfnClientVpnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions() interface{}
	SetAuthenticationOptions(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses.
	//
	// The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. Client CIDR range must have a size of at least /22 and must not be greater than /12.
	ClientCidrBlock() *string
	SetClientCidrBlock(val *string)
	// The options for managing connection authorization for new client connections.
	ClientConnectOptions() interface{}
	SetClientConnectOptions(val interface{})
	// Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.
	ClientLoginBannerOptions() interface{}
	SetClientLoginBannerOptions(val interface{})
	// Information about the client connection logging options.
	//
	// If you enable client connection logging, data about client connections is sent to a Cloudwatch Logs log stream. The following information is logged:
	//
	// - Client connection requests
	// - Client connection results (successful and unsuccessful)
	// - Reasons for unsuccessful client connection requests
	// - Client connection termination time.
	ConnectionLogOptions() interface{}
	SetConnectionLogOptions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A brief description of the Client VPN endpoint.
	Description() *string
	SetDescription(val *string)
	// Information about the DNS servers to be used for DNS resolution.
	//
	// A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address configured on the device is used for the DNS server.
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The IDs of one or more security groups to apply to the target network.
	//
	// You must also specify the ID of the VPC that contains the security groups.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	//
	// Default Value: `enabled`.
	SelfServicePortal() *string
	SetSelfServicePortal(val *string)
	// The ARN of the server certificate.
	//
	// For more information, see the [AWS Certificate Manager User Guide](https://docs.aws.amazon.com/acm/latest/userguide/) .
	ServerCertificateArn() *string
	SetServerCertificateArn(val *string)
	// The maximum VPN session duration time in hours.
	//
	// Valid values: `8 | 10 | 12 | 24`
	//
	// Default value: `24`.
	SessionTimeoutHours() *float64
	SetSessionTimeoutHours(val *float64)
	// Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.
	//
	// By default, split-tunnel on a VPN endpoint is disabled.
	//
	// For information about split-tunnel VPN endpoints, see [Split-tunnel AWS Client VPN endpoint](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html) in the *AWS Client VPN Administrator Guide* .
	SplitTunnel() interface{}
	SetSplitTunnel(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags to apply to the Client VPN endpoint during creation.
	TagSpecifications() interface{}
	SetTagSpecifications(val interface{})
	// The transport protocol to be used by the VPN session.
	//
	// Default value: `udp`.
	TransportProtocol() *string
	SetTransportProtocol(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// The ID of the VPC to associate with the Client VPN endpoint.
	//
	// If no security group IDs are specified in the request, the default security group for the VPC is applied.
	VpcId() *string
	SetVpcId(val *string)
	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	//
	// Valid Values: `443` | `1194`
	//
	// Default Value: `443`.
	VpnPort() *float64
	SetVpnPort(val *float64)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
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
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnClientVpnEndpoint
type jsiiProxy_CfnClientVpnEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClientVpnEndpoint) AuthenticationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ClientConnectOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ClientLoginBannerOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ConnectionLogOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) SplitTunnel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) TagSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::ClientVpnEndpoint`.
func NewCfnClientVpnEndpoint(scope awscdk.Construct, id *string, props *CfnClientVpnEndpointProps) CfnClientVpnEndpoint {
	_init_.Initialize()

	if err := validateNewCfnClientVpnEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClientVpnEndpoint{}

	_jsii_.Create(
		"monocdk.aws_ec2.CfnClientVpnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::ClientVpnEndpoint`.
func NewCfnClientVpnEndpoint_Override(c CfnClientVpnEndpoint, scope awscdk.Construct, id *string, props *CfnClientVpnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.CfnClientVpnEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetAuthenticationOptions(val interface{}) {
	if err := j.validateSetAuthenticationOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetClientCidrBlock(val *string) {
	if err := j.validateSetClientCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCidrBlock",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetClientConnectOptions(val interface{}) {
	if err := j.validateSetClientConnectOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientConnectOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetClientLoginBannerOptions(val interface{}) {
	if err := j.validateSetClientLoginBannerOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientLoginBannerOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetConnectionLogOptions(val interface{}) {
	if err := j.validateSetConnectionLogOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionLogOptions",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetDnsServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSelfServicePortal(val *string) {
	_jsii_.Set(
		j,
		"selfServicePortal",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetServerCertificateArn(val *string) {
	if err := j.validateSetServerCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSessionTimeoutHours(val *float64) {
	_jsii_.Set(
		j,
		"sessionTimeoutHours",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetSplitTunnel(val interface{}) {
	if err := j.validateSetSplitTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitTunnel",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetTagSpecifications(val interface{}) {
	if err := j.validateSetTagSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetTransportProtocol(val *string) {
	_jsii_.Set(
		j,
		"transportProtocol",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_CfnClientVpnEndpoint)SetVpnPort(val *float64) {
	_jsii_.Set(
		j,
		"vpnPort",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnClientVpnEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpoint_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnClientVpnEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnClientVpnEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpoint_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnClientVpnEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnClientVpnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.CfnClientVpnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClientVpnEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_ec2.CfnClientVpnEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnClientVpnEndpoint) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnClientVpnEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpoint) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

