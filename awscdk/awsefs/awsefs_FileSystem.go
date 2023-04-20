package awsefs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsefs/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// The Elastic File System implementation of IFileSystem.
//
// It creates a new, empty file system in Amazon Elastic File System (Amazon EFS).
// It also creates mount target (AWS::EFS::MountTarget) implicitly to mount the
// EFS file system on an Amazon Elastic Compute Cloud (Amazon EC2) instance or another resource.
//
// Example:
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	LifecyclePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
//   	 // files are not transitioned to infrequent access (IA) storage by default
//   	PerformanceMode: efs.PerformanceMode_GENERAL_PURPOSE,
//   	 // default
//   	OutOfInfrequentAccessPolicy: efs.OutOfInfrequentAccessPolicy_AFTER_1_ACCESS,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
//
// Experimental.
type FileSystem interface {
	awscdk.Resource
	IFileSystem
	// The security groups/rules used to allow network connections to the file system.
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
	// The ARN of the file system.
	// Experimental.
	FileSystemArn() *string
	// The ID of the file system, assigned by Amazon EFS.
	// Experimental.
	FileSystemId() *string
	// Dependable that can be depended upon to ensure the mount targets of the filesystem are ready.
	// Experimental.
	MountTargetsAvailable() awscdk.IDependable
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
	// create access point from this filesystem.
	// Experimental.
	AddAccessPoint(id *string, accessPointOptions *AccessPointOptions) AccessPoint
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
	// Grant the actions defined in actions to the given grantee on this File System resource.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
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

// The jsii proxy struct for FileSystem
type jsiiProxy_FileSystem struct {
	internal.Type__awscdkResource
	jsiiProxy_IFileSystem
}

func (j *jsiiProxy_FileSystem) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) FileSystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) MountTargetsAvailable() awscdk.IDependable {
	var returns awscdk.IDependable
	_jsii_.Get(
		j,
		"mountTargetsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Constructor for creating a new EFS FileSystem.
// Experimental.
func NewFileSystem(scope constructs.Construct, id *string, props *FileSystemProps) FileSystem {
	_init_.Initialize()

	if err := validateNewFileSystemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileSystem{}

	_jsii_.Create(
		"monocdk.aws_efs.FileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructor for creating a new EFS FileSystem.
// Experimental.
func NewFileSystem_Override(f FileSystem, scope constructs.Construct, id *string, props *FileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_efs.FileSystem",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing File System from the given properties.
// Experimental.
func FileSystem_FromFileSystemAttributes(scope constructs.Construct, id *string, attrs *FileSystemAttributes) IFileSystem {
	_init_.Initialize()

	if err := validateFileSystem_FromFileSystemAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IFileSystem

	_jsii_.StaticInvoke(
		"monocdk.aws_efs.FileSystem",
		"fromFileSystemAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func FileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_efs.FileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FileSystem_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFileSystem_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_efs.FileSystem",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func FileSystem_DEFAULT_PORT() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"monocdk.aws_efs.FileSystem",
		"DEFAULT_PORT",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FileSystem) AddAccessPoint(id *string, accessPointOptions *AccessPointOptions) AccessPoint {
	if err := f.validateAddAccessPointParameters(id, accessPointOptions); err != nil {
		panic(err)
	}
	var returns AccessPoint

	_jsii_.Invoke(
		f,
		"addAccessPoint",
		[]interface{}{id, accessPointOptions},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FileSystem) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := f.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileSystem) OnSynthesize(session constructs.ISynthesisSession) {
	if err := f.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FileSystem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileSystem) Synthesize(session awscdk.ISynthesisSession) {
	if err := f.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

