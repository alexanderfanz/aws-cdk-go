package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// The FSx for Lustre File System implementation of IFileSystem.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//   var bucket bucket
//
//
//   lustreConfiguration := map[string]interface{}{
//   	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
//   	"exportPath": bucket.s3UrlForObject(),
//   	"importPath": bucket.s3UrlForObject(),
//   	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
//   }
//
//   fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
//   	Vpc: vpc,
//   	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
//   	StorageCapacityGiB: jsii.Number(1200),
//   	LustreConfiguration: LustreConfiguration,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html
//
type LustreFileSystem interface {
	FileSystemBase
	// The security groups/rules used to allow network connections to the file system.
	Connections() awsec2.Connections
	// The DNS name assigned to this file system.
	DnsName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ID that AWS assigns to the file system.
	FileSystemId() *string
	// The mount name of the file system, generated by FSx.
	MountName() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
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

// The jsii proxy struct for LustreFileSystem
type jsiiProxy_LustreFileSystem struct {
	jsiiProxy_FileSystemBase
}

func (j *jsiiProxy_LustreFileSystem) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LustreFileSystem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewLustreFileSystem(scope constructs.Construct, id *string, props *LustreFileSystemProps) LustreFileSystem {
	_init_.Initialize()

	if err := validateNewLustreFileSystemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LustreFileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLustreFileSystem_Override(l LustreFileSystem, scope constructs.Construct, id *string, props *LustreFileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		[]interface{}{scope, id, props},
		l,
	)
}

// Import an existing FSx for Lustre file system from the given properties.
func LustreFileSystem_FromLustreFileSystemAttributes(scope constructs.Construct, id *string, attrs *FileSystemAttributes) IFileSystem {
	_init_.Initialize()

	if err := validateLustreFileSystem_FromLustreFileSystemAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IFileSystem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"fromLustreFileSystemAttributes",
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
func LustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLustreFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func LustreFileSystem_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateLustreFileSystem_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LustreFileSystem_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateLustreFileSystem_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_fsx.LustreFileSystem",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := l.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LustreFileSystem) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := l.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) GetResourceNameAttribute(nameAttr *string) *string {
	if err := l.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

