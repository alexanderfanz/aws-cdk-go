package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a new custom network ACL.
//
// By default, will deny all inbound and outbound traffic unless entries are
// added explicitly allowing it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   networkAcl := awscdk.Aws_ec2.NewNetworkAcl(this, jsii.String("MyNetworkAcl"), &NetworkAclProps{
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	NetworkAclName: jsii.String("networkAclName"),
//   	SubnetSelection: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		SubnetName: jsii.String("subnetName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.*Aws_ec2.SubnetType_ISOLATED,
//   	},
//   })
//
// Experimental.
type NetworkAcl interface {
	awscdk.Resource
	INetworkAcl
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
	// The ID of the NetworkACL.
	// Experimental.
	NetworkAclId() *string
	// The VPC ID for this NetworkACL.
	// Experimental.
	NetworkAclVpcId() *string
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
	// Add a new entry to the ACL.
	// Experimental.
	AddEntry(id *string, options *CommonNetworkAclEntryOptions) NetworkAclEntry
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
	// Associate the ACL with a given set of subnets.
	// Experimental.
	AssociateWithSubnet(id *string, selection *SubnetSelection)
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

// The jsii proxy struct for NetworkAcl
type jsiiProxy_NetworkAcl struct {
	internal.Type__awscdkResource
	jsiiProxy_INetworkAcl
}

func (j *jsiiProxy_NetworkAcl) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAcl) NetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAcl) NetworkAclVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAclVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAcl) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAcl) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkAcl) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewNetworkAcl(scope constructs.Construct, id *string, props *NetworkAclProps) NetworkAcl {
	_init_.Initialize()

	if err := validateNewNetworkAclParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkAcl{}

	_jsii_.Create(
		"monocdk.aws_ec2.NetworkAcl",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkAcl_Override(n NetworkAcl, scope constructs.Construct, id *string, props *NetworkAclProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.NetworkAcl",
		[]interface{}{scope, id, props},
		n,
	)
}

// Import an existing NetworkAcl into this app.
// Experimental.
func NetworkAcl_FromNetworkAclId(scope constructs.Construct, id *string, networkAclId *string) INetworkAcl {
	_init_.Initialize()

	if err := validateNetworkAcl_FromNetworkAclIdParameters(scope, id, networkAclId); err != nil {
		panic(err)
	}
	var returns INetworkAcl

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.NetworkAcl",
		"fromNetworkAclId",
		[]interface{}{scope, id, networkAclId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func NetworkAcl_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkAcl_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.NetworkAcl",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func NetworkAcl_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkAcl_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.NetworkAcl",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) AddEntry(id *string, options *CommonNetworkAclEntryOptions) NetworkAclEntry {
	if err := n.validateAddEntryParameters(id, options); err != nil {
		panic(err)
	}
	var returns NetworkAclEntry

	_jsii_.Invoke(
		n,
		"addEntry",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NetworkAcl) AssociateWithSubnet(id *string, selection *SubnetSelection) {
	if err := n.validateAssociateWithSubnetParameters(id, selection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"associateWithSubnet",
		[]interface{}{id, selection},
	)
}

func (n *jsiiProxy_NetworkAcl) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := n.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) GetResourceNameAttribute(nameAttr *string) *string {
	if err := n.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkAcl) OnSynthesize(session constructs.ISynthesisSession) {
	if err := n.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NetworkAcl) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkAcl) Synthesize(session awscdk.ISynthesisSession) {
	if err := n.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NetworkAcl) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkAcl) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

