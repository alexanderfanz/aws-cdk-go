package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// A GameLift script, that is installed and runs on instances in an Amazon GameLift fleet.
//
// It consists of
// a zip file with all of the components of the realtime game server script.
//
// Example:
//   var bucket bucket
//
//   gamelift.NewScript(this, jsii.String("Script"), &ScriptProps{
//   	Content: gamelift.Content_FromBucket(bucket, jsii.String("sample-asset-key")),
//   })
//
// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-script-uploading.html
//
// Experimental.
type Script interface {
	ScriptBase
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
	// The principal this GameLift script is using.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The IAM role GameLift assumes to acccess server script content.
	// Experimental.
	Role() awsiam.IRole
	// The ARN of the realtime server script.
	// Experimental.
	ScriptArn() *string
	// The Identifier of the realtime server script.
	// Experimental.
	ScriptId() *string
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Script
type jsiiProxy_Script struct {
	jsiiProxy_ScriptBase
}

func (j *jsiiProxy_Script) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) ScriptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) ScriptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Script) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewScript(scope constructs.Construct, id *string, props *ScriptProps) Script {
	_init_.Initialize()

	if err := validateNewScriptParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Script{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Script",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewScript_Override(s Script, scope constructs.Construct, id *string, props *ScriptProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Script",
		[]interface{}{scope, id, props},
		s,
	)
}

// Create a new realtime server script from asset content.
// Experimental.
func Script_FromAsset(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) Script {
	_init_.Initialize()

	if err := validateScript_FromAssetParameters(scope, id, path, options); err != nil {
		panic(err)
	}
	var returns Script

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"fromAsset",
		[]interface{}{scope, id, path, options},
		&returns,
	)

	return returns
}

// Create a new realtime server script from s3 content.
// Experimental.
func Script_FromBucket(scope constructs.Construct, id *string, bucket awss3.IBucket, key *string, objectVersion *string) Script {
	_init_.Initialize()

	if err := validateScript_FromBucketParameters(scope, id, bucket, key); err != nil {
		panic(err)
	}
	var returns Script

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"fromBucket",
		[]interface{}{scope, id, bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Import a script into CDK using its ARN.
// Experimental.
func Script_FromScriptArn(scope constructs.Construct, id *string, scriptArn *string) IScript {
	_init_.Initialize()

	if err := validateScript_FromScriptArnParameters(scope, id, scriptArn); err != nil {
		panic(err)
	}
	var returns IScript

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"fromScriptArn",
		[]interface{}{scope, id, scriptArn},
		&returns,
	)

	return returns
}

// Import an existing realtime server script from its attributes.
// Experimental.
func Script_FromScriptAttributes(scope constructs.Construct, id *string, attrs *ScriptAttributes) IScript {
	_init_.Initialize()

	if err := validateScript_FromScriptAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IScript

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"fromScriptAttributes",
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
func Script_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScript_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Script_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateScript_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Script_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateScript_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Script",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Script) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_Script) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Script) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Script) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Script) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

