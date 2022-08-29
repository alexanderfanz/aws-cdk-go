// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Uses conventionally named roles and asset storage locations.
//
// This synthesizer:
//
// - Supports cross-account deployments (the CLI can have credentials to one
//    account, and you can still deploy to another account by assuming roles with
//    well-known names in the other account).
// - Supports the **CDK Pipelines** library.
//
// Requires the environment to have been bootstrapped with Bootstrap Stack V2
// (also known as "modern bootstrap stack"). The synthesizer adds a version
// check to the template, to make sure the bootstrap stack is recent enough
// to support all features expected by this synthesizer.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewMyStack(app, jsii.String("MyStack"), &stackProps{
//   	synthesizer: awscdk.NewDefaultStackSynthesizer(&defaultStackSynthesizerProps{
//   		fileAssetsBucketName: jsii.String("my-orgs-asset-bucket"),
//   	}),
//   })
//
type DefaultStackSynthesizer interface {
	StackSynthesizer
	// Returns the ARN of the CFN execution Role.
	CloudFormationExecutionRoleArn() *string
	// Returns the ARN of the deploy Role.
	DeployRoleArn() *string
	Stack() Stack
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(asset *FileAssetSource) *FileAssetLocation
	// Bind to the stack this environment is going to be used on.
	//
	// Must be called before any of the other methods are called.
	Bind(stack Stack)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions)
	// Synthesize the associated stack to the session.
	Synthesize(session ISynthesisSession)
	// Have the stack write out its template.
	SynthesizeStackTemplate(stack Stack, session ISynthesisSession)
}

// The jsii proxy struct for DefaultStackSynthesizer
type jsiiProxy_DefaultStackSynthesizer struct {
	jsiiProxy_StackSynthesizer
}

func (j *jsiiProxy_DefaultStackSynthesizer) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultStackSynthesizer) DeployRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultStackSynthesizer) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewDefaultStackSynthesizer(props *DefaultStackSynthesizerProps) DefaultStackSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_DefaultStackSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.DefaultStackSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewDefaultStackSynthesizer_Override(d DefaultStackSynthesizer, props *DefaultStackSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DefaultStackSynthesizer",
		[]interface{}{props},
		d,
	)
}

func DefaultStackSynthesizer_DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_DOCKER_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_DOCKER_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_FILE_ASSETS_BUCKET_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_FILE_ASSETS_BUCKET_NAME",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func DefaultStackSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.DefaultStackSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultStackSynthesizer) AddDockerImageAsset(asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		d,
		"addDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultStackSynthesizer) AddFileAsset(asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		d,
		"addFileAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultStackSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		d,
		"bind",
		[]interface{}{stack},
	)
}

func (d *jsiiProxy_DefaultStackSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		d,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (d *jsiiProxy_DefaultStackSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DefaultStackSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}
