package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// The Nodegroup resource class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var instanceType instanceType
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//
//   nodegroup := awscdk.Aws_eks.NewNodegroup(this, jsii.String("MyNodegroup"), &nodegroupProps{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	amiType: awscdk.*Aws_eks.nodegroupAmiType_AL2_X86_64,
//   	capacityType: awscdk.*Aws_eks.capacityType_SPOT,
//   	desiredSize: jsii.Number(123),
//   	diskSize: jsii.Number(123),
//   	forceUpdate: jsii.Boolean(false),
//   	instanceTypes: []*instanceType{
//   		instanceType,
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	launchTemplateSpec: &launchTemplateSpec{
//   		id: jsii.String("id"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	maxSize: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	nodegroupName: jsii.String("nodegroupName"),
//   	nodeRole: role,
//   	releaseVersion: jsii.String("releaseVersion"),
//   	remoteAccess: &nodegroupRemoteAccess{
//   		sshKeyName: jsii.String("sshKeyName"),
//
//   		// the properties below are optional
//   		sourceSecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   	},
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	taints: []taintSpec{
//   		&taintSpec{
//   			effect: awscdk.*Aws_eks.taintEffect_NO_SCHEDULE,
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type Nodegroup interface {
	awscdk.Resource
	INodegroup
	// the Amazon EKS cluster resource.
	Cluster() ICluster
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// ARN of the nodegroup.
	NodegroupArn() *string
	// Nodegroup name.
	NodegroupName() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// IAM role of the instance profile for the nodegroup.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
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

// The jsii proxy struct for Nodegroup
type jsiiProxy_Nodegroup struct {
	internal.Type__awscdkResource
	jsiiProxy_INodegroup
}

func (j *jsiiProxy_Nodegroup) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) NodegroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Nodegroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewNodegroup(scope constructs.Construct, id *string, props *NodegroupProps) Nodegroup {
	_init_.Initialize()

	if err := validateNewNodegroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Nodegroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.Nodegroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNodegroup_Override(n Nodegroup, scope constructs.Construct, id *string, props *NodegroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.Nodegroup",
		[]interface{}{scope, id, props},
		n,
	)
}

// Import the Nodegroup from attributes.
func Nodegroup_FromNodegroupName(scope constructs.Construct, id *string, nodegroupName *string) INodegroup {
	_init_.Initialize()

	if err := validateNodegroup_FromNodegroupNameParameters(scope, id, nodegroupName); err != nil {
		panic(err)
	}
	var returns INodegroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Nodegroup",
		"fromNodegroupName",
		[]interface{}{scope, id, nodegroupName},
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
func Nodegroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNodegroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Nodegroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Nodegroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNodegroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Nodegroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Nodegroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNodegroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.Nodegroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Nodegroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_Nodegroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Nodegroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (n *jsiiProxy_Nodegroup) GetResourceNameAttribute(nameAttr *string) *string {
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

func (n *jsiiProxy_Nodegroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

