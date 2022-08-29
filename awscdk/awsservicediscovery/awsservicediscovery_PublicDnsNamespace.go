package awsservicediscovery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Public DNS Namespace.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
//   	name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
//   	name: jsii.String("foo"),
//   	dnsRecordType: servicediscovery.dnsRecordType_A,
//   	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
//   	healthCheck: &healthCheckConfig{
//   		type: servicediscovery.healthCheckType_HTTPS,
//   		resourcePath: jsii.String("/healthcheck"),
//   		failureThreshold: jsii.Number(2),
//   	},
//   })
//
//   service.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
//   	ipv4: jsii.String("54.239.25.192"),
//   	port: jsii.Number(443),
//   })
//
//   app.synth()
//
type PublicDnsNamespace interface {
	awscdk.Resource
	IPublicDnsNamespace
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Namespace Arn for the namespace.
	NamespaceArn() *string
	// ID of hosted zone created by namespace.
	NamespaceHostedZoneId() *string
	// Namespace Id for the namespace.
	NamespaceId() *string
	// A name for the namespace.
	NamespaceName() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	PublicDnsNamespaceArn() *string
	PublicDnsNamespaceId() *string
	PublicDnsNamespaceName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Type of the namespace.
	Type() NamespaceType
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Creates a service within the namespace.
	CreateService(id *string, props *DnsServiceProps) Service
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PublicDnsNamespace
type jsiiProxy_PublicDnsNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IPublicDnsNamespace
}

func (j *jsiiProxy_PublicDnsNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewPublicDnsNamespace(scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) PublicDnsNamespace {
	_init_.Initialize()

	j := jsiiProxy_PublicDnsNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPublicDnsNamespace_Override(p PublicDnsNamespace, scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		[]interface{}{scope, id, props},
		p,
	)
}

func PublicDnsNamespace_FromPublicDnsNamespaceAttributes(scope constructs.Construct, id *string, attrs *PublicDnsNamespaceAttributes) IPublicDnsNamespace {
	_init_.Initialize()

	var returns IPublicDnsNamespace

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"fromPublicDnsNamespaceAttributes",
		[]interface{}{scope, id, attrs},
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
func PublicDnsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func PublicDnsNamespace_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PublicDnsNamespace_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.PublicDnsNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PublicDnsNamespace) CreateService(id *string, props *DnsServiceProps) Service {
	var returns Service

	_jsii_.Invoke(
		p,
		"createService",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
