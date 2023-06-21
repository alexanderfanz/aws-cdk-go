package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a new AppMesh mesh.
//
// Example:
//   // This is the ARN for the mesh from different AWS IAM account ID.
//   // Ensure mesh is properly shared with your account. For more details, see: https://github.com/aws/aws-cdk/issues/15404
//   arn := "arn:aws:appmesh:us-east-1:123456789012:mesh/testMesh"
//   sharedMesh := appmesh.Mesh_FromMeshArn(this, jsii.String("imported-mesh"), arn)
//
//   // This VirtualNode resource can communicate with the resources in the mesh from different AWS IAM account ID.
//   // This VirtualNode resource can communicate with the resources in the mesh from different AWS IAM account ID.
//   appmesh.NewVirtualNode(this, jsii.String("test-node"), &VirtualNodeProps{
//   	Mesh: sharedMesh,
//   })
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/meshes.html
//
type Mesh interface {
	awscdk.Resource
	IMesh
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The Amazon Resource Name (ARN) of the AppMesh mesh.
	MeshArn() *string
	// The name of the AppMesh mesh.
	MeshName() *string
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
	// Adds a VirtualGateway to the Mesh.
	AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway
	// Adds a VirtualNode to the Mesh.
	AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode
	// Adds a VirtualRouter to the Mesh with the given id and props.
	AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter
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

// The jsii proxy struct for Mesh
type jsiiProxy_Mesh struct {
	internal.Type__awscdkResource
	jsiiProxy_IMesh
}

func (j *jsiiProxy_Mesh) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) MeshArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) MeshName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meshName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Mesh) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewMesh(scope constructs.Construct, id *string, props *MeshProps) Mesh {
	_init_.Initialize()

	if err := validateNewMeshParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Mesh{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Mesh",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMesh_Override(m Mesh, scope constructs.Construct, id *string, props *MeshProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Mesh",
		[]interface{}{scope, id, props},
		m,
	)
}

// Import an existing mesh by arn.
func Mesh_FromMeshArn(scope constructs.Construct, id *string, meshArn *string) IMesh {
	_init_.Initialize()

	if err := validateMesh_FromMeshArnParameters(scope, id, meshArn); err != nil {
		panic(err)
	}
	var returns IMesh

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"fromMeshArn",
		[]interface{}{scope, id, meshArn},
		&returns,
	)

	return returns
}

// Import an existing mesh by name.
func Mesh_FromMeshName(scope constructs.Construct, id *string, meshName *string) IMesh {
	_init_.Initialize()

	if err := validateMesh_FromMeshNameParameters(scope, id, meshName); err != nil {
		panic(err)
	}
	var returns IMesh

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"fromMeshName",
		[]interface{}{scope, id, meshName},
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
func Mesh_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMesh_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Mesh_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMesh_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Mesh_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMesh_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Mesh",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) AddVirtualGateway(id *string, props *VirtualGatewayBaseProps) VirtualGateway {
	if err := m.validateAddVirtualGatewayParameters(id, props); err != nil {
		panic(err)
	}
	var returns VirtualGateway

	_jsii_.Invoke(
		m,
		"addVirtualGateway",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) AddVirtualNode(id *string, props *VirtualNodeBaseProps) VirtualNode {
	if err := m.validateAddVirtualNodeParameters(id, props); err != nil {
		panic(err)
	}
	var returns VirtualNode

	_jsii_.Invoke(
		m,
		"addVirtualNode",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) AddVirtualRouter(id *string, props *VirtualRouterBaseProps) VirtualRouter {
	if err := m.validateAddVirtualRouterParameters(id, props); err != nil {
		panic(err)
	}
	var returns VirtualRouter

	_jsii_.Invoke(
		m,
		"addVirtualRouter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_Mesh) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := m.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) GetResourceNameAttribute(nameAttr *string) *string {
	if err := m.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mesh) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

