package awsglobalaccelerator

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// The Accelerator construct.
// Experimental.
type Accelerator interface {
	awscdk.Resource
	IAccelerator
	AcceleratorArn() *string
	DnsName() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Accelerator
type jsiiProxy_Accelerator struct {
	internal.Type__awscdkResource
	jsiiProxy_IAccelerator
}

func (j *jsiiProxy_Accelerator) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Accelerator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAccelerator(scope constructs.Construct, id *string, props *AcceleratorProps) Accelerator {
	_init_.Initialize()

	j := jsiiProxy_Accelerator{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Accelerator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAccelerator_Override(a Accelerator, scope constructs.Construct, id *string, props *AcceleratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Accelerator",
		[]interface{}{scope, id, props},
		a,
	)
}

// import from attributes.
// Experimental.
func Accelerator_FromAcceleratorAttributes(scope constructs.Construct, id *string, attrs *AcceleratorAttributes) IAccelerator {
	_init_.Initialize()

	var returns IAccelerator

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Accelerator",
		"fromAcceleratorAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Accelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Accelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Accelerator_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Accelerator",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (a *jsiiProxy_Accelerator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_Accelerator) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (a *jsiiProxy_Accelerator) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (a *jsiiProxy_Accelerator) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_Accelerator) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_Accelerator) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_Accelerator) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_Accelerator) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_Accelerator) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_Accelerator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_Accelerator) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes required to import an existing accelerator to the stack.
// Experimental.
type AcceleratorAttributes struct {
	// The ARN of the accelerator.
	// Experimental.
	AcceleratorArn *string `json:"acceleratorArn"`
	// The DNS name of the accelerator.
	// Experimental.
	DnsName *string `json:"dnsName"`
}

// Construct properties of the Accelerator.
// Experimental.
type AcceleratorProps struct {
	// The name of the accelerator.
	// Experimental.
	AcceleratorName *string `json:"acceleratorName"`
	// Indicates whether the accelerator is enabled.
	// Experimental.
	Enabled *bool `json:"enabled"`
}

// The security group used by a Global Accelerator to send traffic to resources in a VPC.
// Experimental.
type AcceleratorSecurityGroup interface {
}

// The jsii proxy struct for AcceleratorSecurityGroup
type jsiiProxy_AcceleratorSecurityGroup struct {
	_ byte // padding
}

// Lookup the Global Accelerator security group at CloudFormation deployment time.
//
// As of this writing, Global Accelerators (AGA) create a single security group per VPC. AGA security groups are shared
// by all AGAs in an account. Additionally, there is no CloudFormation mechanism to reference the AGA security groups.
//
// This makes creating security group rules which allow traffic from an AGA complicated in CDK. This lookup will identify
// the AGA security group for a given VPC at CloudFormation deployment time, and lets you create rules for traffic from AGA
// to other resources created by CDK.
// Experimental.
func AcceleratorSecurityGroup_FromVpc(scope awscdk.Construct, id *string, vpc awsec2.IVpc, endpointGroup EndpointGroup) awsec2.ISecurityGroup {
	_init_.Initialize()

	var returns awsec2.ISecurityGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.AcceleratorSecurityGroup",
		"fromVpc",
		[]interface{}{scope, id, vpc, endpointGroup},
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::GlobalAccelerator::Accelerator`.
type CfnAccelerator interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrAcceleratorArn() *string
	AttrDnsName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	IpAddresses() *[]*string
	SetIpAddresses(val *[]*string)
	IpAddressType() *string
	SetIpAddressType(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAccelerator
type jsiiProxy_CfnAccelerator struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccelerator) AttrAcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAcceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) AttrDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccelerator) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GlobalAccelerator::Accelerator`.
func NewCfnAccelerator(scope awscdk.Construct, id *string, props *CfnAcceleratorProps) CfnAccelerator {
	_init_.Initialize()

	j := jsiiProxy_CfnAccelerator{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::Accelerator`.
func NewCfnAccelerator_Override(c CfnAccelerator, scope awscdk.Construct, id *string, props *CfnAcceleratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetIpAddresses(val *[]*string) {
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetIpAddressType(val *string) {
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_CfnAccelerator) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
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
func CfnAccelerator_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAccelerator_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnAccelerator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccelerator_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_globalaccelerator.CfnAccelerator",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnAccelerator) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnAccelerator) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnAccelerator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnAccelerator) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnAccelerator) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::GlobalAccelerator::Accelerator`.
type CfnAcceleratorProps struct {
	// `AWS::GlobalAccelerator::Accelerator.Name`.
	Name *string `json:"name"`
	// `AWS::GlobalAccelerator::Accelerator.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `AWS::GlobalAccelerator::Accelerator.IpAddresses`.
	IpAddresses *[]*string `json:"ipAddresses"`
	// `AWS::GlobalAccelerator::Accelerator.IpAddressType`.
	IpAddressType *string `json:"ipAddressType"`
	// `AWS::GlobalAccelerator::Accelerator.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::GlobalAccelerator::EndpointGroup`.
type CfnEndpointGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrEndpointGroupArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EndpointConfigurations() interface{}
	SetEndpointConfigurations(val interface{})
	EndpointGroupRegion() *string
	SetEndpointGroupRegion(val *string)
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPort() *float64
	SetHealthCheckPort(val *float64)
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	ListenerArn() *string
	SetListenerArn(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	PortOverrides() interface{}
	SetPortOverrides(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	ThresholdCount() *float64
	SetThresholdCount(val *float64)
	TrafficDialPercentage() *float64
	SetTrafficDialPercentage(val *float64)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEndpointGroup
type jsiiProxy_CfnEndpointGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpointGroup) AttrEndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) EndpointConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) EndpointGroupRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) PortOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) ThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) TrafficDialPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trafficDialPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroup(scope awscdk.Construct, id *string, props *CfnEndpointGroupProps) CfnEndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpointGroup{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::EndpointGroup`.
func NewCfnEndpointGroup_Override(c CfnEndpointGroup, scope awscdk.Construct, id *string, props *CfnEndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetEndpointConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"endpointConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetEndpointGroupRegion(val *string) {
	_jsii_.Set(
		j,
		"endpointGroupRegion",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckIntervalSeconds(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckPath(val *string) {
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckPort(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetHealthCheckProtocol(val *string) {
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetListenerArn(val *string) {
	_jsii_.Set(
		j,
		"listenerArn",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetPortOverrides(val interface{}) {
	_jsii_.Set(
		j,
		"portOverrides",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetThresholdCount(val *float64) {
	_jsii_.Set(
		j,
		"thresholdCount",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointGroup) SetTrafficDialPercentage(val *float64) {
	_jsii_.Set(
		j,
		"trafficDialPercentage",
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
func CfnEndpointGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEndpointGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnEndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_globalaccelerator.CfnEndpointGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnEndpointGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnEndpointGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnEndpointGroup_EndpointConfigurationProperty struct {
	// `CfnEndpointGroup.EndpointConfigurationProperty.EndpointId`.
	EndpointId *string `json:"endpointId"`
	// `CfnEndpointGroup.EndpointConfigurationProperty.ClientIPPreservationEnabled`.
	ClientIpPreservationEnabled interface{} `json:"clientIpPreservationEnabled"`
	// `CfnEndpointGroup.EndpointConfigurationProperty.Weight`.
	Weight *float64 `json:"weight"`
}

type CfnEndpointGroup_PortOverrideProperty struct {
	// `CfnEndpointGroup.PortOverrideProperty.EndpointPort`.
	EndpointPort *float64 `json:"endpointPort"`
	// `CfnEndpointGroup.PortOverrideProperty.ListenerPort`.
	ListenerPort *float64 `json:"listenerPort"`
}

// Properties for defining a `AWS::GlobalAccelerator::EndpointGroup`.
type CfnEndpointGroupProps struct {
	// `AWS::GlobalAccelerator::EndpointGroup.EndpointGroupRegion`.
	EndpointGroupRegion *string `json:"endpointGroupRegion"`
	// `AWS::GlobalAccelerator::EndpointGroup.ListenerArn`.
	ListenerArn *string `json:"listenerArn"`
	// `AWS::GlobalAccelerator::EndpointGroup.EndpointConfigurations`.
	EndpointConfigurations interface{} `json:"endpointConfigurations"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckIntervalSeconds`.
	HealthCheckIntervalSeconds *float64 `json:"healthCheckIntervalSeconds"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckPath`.
	HealthCheckPath *string `json:"healthCheckPath"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckPort`.
	HealthCheckPort *float64 `json:"healthCheckPort"`
	// `AWS::GlobalAccelerator::EndpointGroup.HealthCheckProtocol`.
	HealthCheckProtocol *string `json:"healthCheckProtocol"`
	// `AWS::GlobalAccelerator::EndpointGroup.PortOverrides`.
	PortOverrides interface{} `json:"portOverrides"`
	// `AWS::GlobalAccelerator::EndpointGroup.ThresholdCount`.
	ThresholdCount *float64 `json:"thresholdCount"`
	// `AWS::GlobalAccelerator::EndpointGroup.TrafficDialPercentage`.
	TrafficDialPercentage *float64 `json:"trafficDialPercentage"`
}

// A CloudFormation `AWS::GlobalAccelerator::Listener`.
type CfnListener interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AcceleratorArn() *string
	SetAcceleratorArn(val *string)
	AttrListenerArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClientAffinity() *string
	SetClientAffinity(val *string)
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	PortRanges() interface{}
	SetPortRanges(val interface{})
	Protocol() *string
	SetProtocol(val *string)
	Ref() *string
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnListener
type jsiiProxy_CfnListener struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnListener) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) AttrListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrListenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) ClientAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) PortRanges() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnListener) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GlobalAccelerator::Listener`.
func NewCfnListener(scope awscdk.Construct, id *string, props *CfnListenerProps) CfnListener {
	_init_.Initialize()

	j := jsiiProxy_CfnListener{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnListener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GlobalAccelerator::Listener`.
func NewCfnListener_Override(c CfnListener, scope awscdk.Construct, id *string, props *CfnListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.CfnListener",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnListener) SetAcceleratorArn(val *string) {
	_jsii_.Set(
		j,
		"acceleratorArn",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetClientAffinity(val *string) {
	_jsii_.Set(
		j,
		"clientAffinity",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetPortRanges(val interface{}) {
	_jsii_.Set(
		j,
		"portRanges",
		val,
	)
}

func (j *jsiiProxy_CfnListener) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
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
func CfnListener_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnListener",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnListener_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnListener",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.CfnListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnListener_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_globalaccelerator.CfnListener",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnListener) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnListener) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnListener) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnListener) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnListener) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnListener) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnListener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnListener) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnListener) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
// Experimental.
func (c *jsiiProxy_CfnListener) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnListener) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnListener) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnListener) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnListener) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnListener) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnListener) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnListener) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnListener) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnListener) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnListener) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnListener_PortRangeProperty struct {
	// `CfnListener.PortRangeProperty.FromPort`.
	FromPort *float64 `json:"fromPort"`
	// `CfnListener.PortRangeProperty.ToPort`.
	ToPort *float64 `json:"toPort"`
}

// Properties for defining a `AWS::GlobalAccelerator::Listener`.
type CfnListenerProps struct {
	// `AWS::GlobalAccelerator::Listener.AcceleratorArn`.
	AcceleratorArn *string `json:"acceleratorArn"`
	// `AWS::GlobalAccelerator::Listener.PortRanges`.
	PortRanges interface{} `json:"portRanges"`
	// `AWS::GlobalAccelerator::Listener.Protocol`.
	Protocol *string `json:"protocol"`
	// `AWS::GlobalAccelerator::Listener.ClientAffinity`.
	ClientAffinity *string `json:"clientAffinity"`
}

// Client affinity lets you direct all requests from a user to the same endpoint, if you have stateful applications, regardless of the port and protocol of the client request.
//
// Client affinity gives you control over whether to always
// route each client to the same specific endpoint. If you want a given client to always be routed to the same
// endpoint, set client affinity to SOURCE_IP.
// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-listeners.html#about-listeners-client-affinity
//
// Experimental.
type ClientAffinity string

const (
	ClientAffinity_NONE ClientAffinity = "NONE"
	ClientAffinity_SOURCE_IP ClientAffinity = "SOURCE_IP"
)

// The protocol for the connections from clients to the accelerator.
// Experimental.
type ConnectionProtocol string

const (
	ConnectionProtocol_TCP ConnectionProtocol = "TCP"
	ConnectionProtocol_UDP ConnectionProtocol = "UDP"
)

// EC2 Instance interface.
// Experimental.
type Ec2Instance struct {
	// The id of the instance resource.
	// Experimental.
	InstanceId *string `json:"instanceId"`
}

// EIP Interface.
// Experimental.
type ElasticIpAddress struct {
	// allocation ID of the EIP resoruce.
	// Experimental.
	AttrAllocationId *string `json:"attrAllocationId"`
}

// The class for endpoint configuration.
// Experimental.
type EndpointConfiguration interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	Props() *EndpointConfigurationProps
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderEndpointConfiguration() *CfnEndpointGroup_EndpointConfigurationProperty
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for EndpointConfiguration
type jsiiProxy_EndpointConfiguration struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_EndpointConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointConfiguration) Props() *EndpointConfigurationProps {
	var returns *EndpointConfigurationProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewEndpointConfiguration(scope constructs.Construct, id *string, props *EndpointConfigurationProps) EndpointConfiguration {
	_init_.Initialize()

	j := jsiiProxy_EndpointConfiguration{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.EndpointConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEndpointConfiguration_Override(e EndpointConfiguration, scope constructs.Construct, id *string, props *EndpointConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.EndpointConfiguration",
		[]interface{}{scope, id, props},
		e,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func EndpointConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// render the endpoint configuration for the endpoint group.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) RenderEndpointConfiguration() *CfnEndpointGroup_EndpointConfigurationProperty {
	var returns *CfnEndpointGroup_EndpointConfigurationProperty

	_jsii_.Invoke(
		e,
		"renderEndpointConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EndpointConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for `addLoadBalancer`, `addElasticIpAddress` and `addEc2Instance` to add endpoints into the endpoint group.
// Experimental.
type EndpointConfigurationOptions struct {
	// Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint.
	// Experimental.
	ClientIpReservation *bool `json:"clientIpReservation"`
	// The weight associated with the endpoint.
	//
	// When you add weights to endpoints, you configure AWS Global Accelerator
	// to route traffic based on proportions that you specify. For example, you might specify endpoint weights of 4, 5,
	// 5, and 6 (sum=20). The result is that 4/20 of your traffic, on average, is routed to the first endpoint, 5/20 is
	// routed both to the second and third endpoints, and 6/20 is routed to the last endpoint.
	// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html
	//
	// Experimental.
	Weight *float64 `json:"weight"`
}

// Properties to create EndpointConfiguration.
// Experimental.
type EndpointConfigurationProps struct {
	// Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint.
	// Experimental.
	ClientIpReservation *bool `json:"clientIpReservation"`
	// The weight associated with the endpoint.
	//
	// When you add weights to endpoints, you configure AWS Global Accelerator
	// to route traffic based on proportions that you specify. For example, you might specify endpoint weights of 4, 5,
	// 5, and 6 (sum=20). The result is that 4/20 of your traffic, on average, is routed to the first endpoint, 5/20 is
	// routed both to the second and third endpoints, and 6/20 is routed to the last endpoint.
	// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html
	//
	// Experimental.
	Weight *float64 `json:"weight"`
	// The endopoint group reesource.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	EndpointGroup EndpointGroup `json:"endpointGroup"`
	// An ID for the endpoint.
	//
	// If the endpoint is a Network Load Balancer or Application Load Balancer,
	// this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address,
	// this is the Elastic IP address allocation ID. For EC2 instances, this is the EC2 instance ID.
	// Experimental.
	EndpointId *string `json:"endpointId"`
}

// EndpointGroup construct.
// Experimental.
type EndpointGroup interface {
	awscdk.Resource
	IEndpointGroup
	EndpointGroupArn() *string
	EndpointGroupName() *string
	Endpoints() *[]EndpointConfiguration
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddEc2Instance(id *string, instance *Ec2Instance, props *EndpointConfigurationOptions) EndpointConfiguration
	AddElasticIpAddress(id *string, eip *ElasticIpAddress, props *EndpointConfigurationOptions) EndpointConfiguration
	AddEndpoint(id *string, endpointId *string, props *EndpointConfigurationOptions) EndpointConfiguration
	AddLoadBalancer(id *string, lb *LoadBalancer, props *EndpointConfigurationOptions) EndpointConfiguration
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for EndpointGroup
type jsiiProxy_EndpointGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IEndpointGroup
}

func (j *jsiiProxy_EndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) EndpointGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Endpoints() *[]EndpointConfiguration {
	var returns *[]EndpointConfiguration
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EndpointGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewEndpointGroup(scope constructs.Construct, id *string, props *EndpointGroupProps) EndpointGroup {
	_init_.Initialize()

	j := jsiiProxy_EndpointGroup{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEndpointGroup_Override(e EndpointGroup, scope constructs.Construct, id *string, props *EndpointGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from ARN.
// Experimental.
func EndpointGroup_FromEndpointGroupArn(scope constructs.Construct, id *string, endpointGroupArn *string) IEndpointGroup {
	_init_.Initialize()

	var returns IEndpointGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		"fromEndpointGroupArn",
		[]interface{}{scope, id, endpointGroupArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EndpointGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func EndpointGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.EndpointGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add an EC2 Instance as an endpoint in this endpoint group.
// Experimental.
func (e *jsiiProxy_EndpointGroup) AddEc2Instance(id *string, instance *Ec2Instance, props *EndpointConfigurationOptions) EndpointConfiguration {
	var returns EndpointConfiguration

	_jsii_.Invoke(
		e,
		"addEc2Instance",
		[]interface{}{id, instance, props},
		&returns,
	)

	return returns
}

// Add an EIP as an endpoint in this endpoint group.
// Experimental.
func (e *jsiiProxy_EndpointGroup) AddElasticIpAddress(id *string, eip *ElasticIpAddress, props *EndpointConfigurationOptions) EndpointConfiguration {
	var returns EndpointConfiguration

	_jsii_.Invoke(
		e,
		"addElasticIpAddress",
		[]interface{}{id, eip, props},
		&returns,
	)

	return returns
}

// Add an endpoint.
// Experimental.
func (e *jsiiProxy_EndpointGroup) AddEndpoint(id *string, endpointId *string, props *EndpointConfigurationOptions) EndpointConfiguration {
	var returns EndpointConfiguration

	_jsii_.Invoke(
		e,
		"addEndpoint",
		[]interface{}{id, endpointId, props},
		&returns,
	)

	return returns
}

// Add an Elastic Load Balancer as an endpoint in this endpoint group.
// Experimental.
func (e *jsiiProxy_EndpointGroup) AddLoadBalancer(id *string, lb *LoadBalancer, props *EndpointConfigurationOptions) EndpointConfiguration {
	var returns EndpointConfiguration

	_jsii_.Invoke(
		e,
		"addLoadBalancer",
		[]interface{}{id, lb, props},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (e *jsiiProxy_EndpointGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (e *jsiiProxy_EndpointGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (e *jsiiProxy_EndpointGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (e *jsiiProxy_EndpointGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EndpointGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EndpointGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EndpointGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EndpointGroup) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EndpointGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EndpointGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EndpointGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Property of the EndpointGroup.
// Experimental.
type EndpointGroupProps struct {
	// The Amazon Resource Name (ARN) of the listener.
	// Experimental.
	Listener IListener `json:"listener"`
	// Name of the endpoint group.
	// Experimental.
	EndpointGroupName *string `json:"endpointGroupName"`
	// The AWS Region where the endpoint group is located.
	// Experimental.
	Region *string `json:"region"`
}

// The interface of the Accelerator.
// Experimental.
type IAccelerator interface {
	awscdk.IResource
	// The ARN of the accelerator.
	// Experimental.
	AcceleratorArn() *string
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
	// Experimental.
	DnsName() *string
}

// The jsii proxy for IAccelerator
type jsiiProxy_IAccelerator struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccelerator) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

// The interface of the EndpointGroup.
// Experimental.
type IEndpointGroup interface {
	awscdk.IResource
	// EndpointGroup ARN.
	// Experimental.
	EndpointGroupArn() *string
}

// The jsii proxy for IEndpointGroup
type jsiiProxy_IEndpointGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

// Interface of the Listener.
// Experimental.
type IListener interface {
	awscdk.IResource
	// The ARN of the listener.
	// Experimental.
	ListenerArn() *string
}

// The jsii proxy for IListener
type jsiiProxy_IListener struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

// The construct for the Listener.
// Experimental.
type Listener interface {
	awscdk.Resource
	IListener
	Env() *awscdk.ResourceEnvironment
	ListenerArn() *string
	ListenerName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Listener
type jsiiProxy_Listener struct {
	internal.Type__awscdkResource
	jsiiProxy_IListener
}

func (j *jsiiProxy_Listener) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) ListenerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Listener) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewListener(scope constructs.Construct, id *string, props *ListenerProps) Listener {
	_init_.Initialize()

	j := jsiiProxy_Listener{}

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Listener",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewListener_Override(l Listener, scope constructs.Construct, id *string, props *ListenerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_globalaccelerator.Listener",
		[]interface{}{scope, id, props},
		l,
	)
}

// import from ARN.
// Experimental.
func Listener_FromListenerArn(scope constructs.Construct, id *string, listenerArn *string) IListener {
	_init_.Initialize()

	var returns IListener

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Listener",
		"fromListenerArn",
		[]interface{}{scope, id, listenerArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Listener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Listener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Listener_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_globalaccelerator.Listener",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (l *jsiiProxy_Listener) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (l *jsiiProxy_Listener) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (l *jsiiProxy_Listener) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (l *jsiiProxy_Listener) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_Listener) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_Listener) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_Listener) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_Listener) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_Listener) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (l *jsiiProxy_Listener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_Listener) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// construct properties for Listener.
// Experimental.
type ListenerProps struct {
	// The accelerator for this listener.
	// Experimental.
	Accelerator IAccelerator `json:"accelerator"`
	// The list of port ranges for the connections from clients to the accelerator.
	// Experimental.
	PortRanges *[]*PortRange `json:"portRanges"`
	// Client affinity to direct all requests from a user to the same endpoint.
	// Experimental.
	ClientAffinity ClientAffinity `json:"clientAffinity"`
	// Name of the listener.
	// Experimental.
	ListenerName *string `json:"listenerName"`
	// The protocol for the connections from clients to the accelerator.
	// Experimental.
	Protocol ConnectionProtocol `json:"protocol"`
}

// LoadBalancer Interface.
// Experimental.
type LoadBalancer struct {
	// The ARN of this load balancer.
	// Experimental.
	LoadBalancerArn *string `json:"loadBalancerArn"`
}

// The list of port ranges for the connections from clients to the accelerator.
// Experimental.
type PortRange struct {
	// The first port in the range of ports, inclusive.
	// Experimental.
	FromPort *float64 `json:"fromPort"`
	// The last port in the range of ports, inclusive.
	// Experimental.
	ToPort *float64 `json:"toPort"`
}

