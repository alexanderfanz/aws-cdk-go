// The CDK Construct Library for AWS::AppRunner
package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The App Runner VPC Connector.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &VpcConnectorProps{
//   	Vpc: Vpc,
//   	VpcSubnets: vpc.selectSubnets(&SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	}),
//   	VpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	VpcConnector: VpcConnector,
//   })
//
// Experimental.
type VpcConnector interface {
	awscdk.Resource
	IVpcConnector
	// Allows specifying security group connections for the VPC connector.
	// Experimental.
	Connections() awsec2.Connections
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn() *string
	// The name of the VPC connector.
	// Experimental.
	VpcConnectorName() *string
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision() *float64
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpcConnector
type jsiiProxy_VpcConnector struct {
	internal.Type__awscdkResource
	jsiiProxy_IVpcConnector
}

func (j *jsiiProxy_VpcConnector) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) VpcConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) VpcConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpcConnector) VpcConnectorRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpcConnectorRevision",
		&returns,
	)
	return returns
}


// Experimental.
func NewVpcConnector(scope constructs.Construct, id *string, props *VpcConnectorProps) VpcConnector {
	_init_.Initialize()

	if err := validateNewVpcConnectorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpcConnector{}

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVpcConnector_Override(v VpcConnector, scope constructs.Construct, id *string, props *VpcConnectorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import from VPC connector attributes.
// Experimental.
func VpcConnector_FromVpcConnectorAttributes(scope constructs.Construct, id *string, attrs *VpcConnectorAttributes) IVpcConnector {
	_init_.Initialize()

	if err := validateVpcConnector_FromVpcConnectorAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IVpcConnector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"fromVpcConnectorAttributes",
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
// Experimental.
func VpcConnector_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpcConnector_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func VpcConnector_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVpcConnector_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func VpcConnector_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVpcConnector_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.VpcConnector",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := v.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (v *jsiiProxy_VpcConnector) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := v.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) GetResourceNameAttribute(nameAttr *string) *string {
	if err := v.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpcConnector) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

