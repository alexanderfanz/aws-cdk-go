package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define an EventBridge Connection.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
//   	Authorization: events.Authorization_ApiKey(jsii.String("x-api-key"), awscdk.SecretValue_SecretsManager(jsii.String("ApiSecretName"))),
//   	Description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &ApiDestinationProps{
//   	Connection: Connection,
//   	Endpoint: jsii.String("https://example.com"),
//   	Description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Minutes(jsii.Number(1))),
//   	Targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
// Experimental.
type Connection interface {
	awscdk.Resource
	IConnection
	// The ARN of the connection created.
	// Experimental.
	ConnectionArn() *string
	// The Name for the connection.
	// Experimental.
	ConnectionName() *string
	// The ARN for the secret created for the connection.
	// Experimental.
	ConnectionSecretArn() *string
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

// The jsii proxy struct for Connection
type jsiiProxy_Connection struct {
	internal.Type__awscdkResource
	jsiiProxy_IConnection
}

func (j *jsiiProxy_Connection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) ConnectionSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnection(scope constructs.Construct, id *string, props *ConnectionProps) Connection {
	_init_.Initialize()

	if err := validateNewConnectionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Connection{}

	_jsii_.Create(
		"monocdk.aws_events.Connection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewConnection_Override(c Connection, scope constructs.Construct, id *string, props *ConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events.Connection",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing connection resource.
// Experimental.
func Connection_FromConnectionAttributes(scope constructs.Construct, id *string, attrs *ConnectionAttributes) IConnection {
	_init_.Initialize()

	if err := validateConnection_FromConnectionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IConnection

	_jsii_.StaticInvoke(
		"monocdk.aws_events.Connection",
		"fromConnectionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing connection resource.
// Experimental.
func Connection_FromEventBusArn(scope constructs.Construct, id *string, connectionArn *string, connectionSecretArn *string) IConnection {
	_init_.Initialize()

	if err := validateConnection_FromEventBusArnParameters(scope, id, connectionArn, connectionSecretArn); err != nil {
		panic(err)
	}
	var returns IConnection

	_jsii_.StaticInvoke(
		"monocdk.aws_events.Connection",
		"fromEventBusArn",
		[]interface{}{scope, id, connectionArn, connectionSecretArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Connection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_events.Connection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Connection_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateConnection_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_events.Connection",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Connection) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connection) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Connection) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Connection) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_Connection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Connection) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

