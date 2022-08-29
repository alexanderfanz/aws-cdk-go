// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Synthesizer that reuses bootstrap roles from a different region.
//
// A special synthesizer that behaves similarly to `DefaultStackSynthesizer`,
// but doesn't require bootstrapping the environment it operates in. Instead,
// it will re-use the Roles that were created for a different region (which
// is possible because IAM is a global service).
//
// However, it will not assume asset buckets or repositories have been created,
// and therefore does not support assets.
//
// Used by the CodePipeline construct for the support stacks needed for
// cross-region replication S3 buckets. App builders do not need to use this
// synthesizer directly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   bootstraplessSynthesizer := cdk.NewBootstraplessSynthesizer(&bootstraplessSynthesizerProps{
//   	cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   	deployRoleArn: jsii.String("deployRoleArn"),
//   })
//
type BootstraplessSynthesizer interface {
	DefaultStackSynthesizer
	// Returns the ARN of the CFN execution Role.
	CloudFormationExecutionRoleArn() *string
	// Returns the ARN of the deploy Role.
	DeployRoleArn() *string
	Stack() Stack
	// Register a Docker Image Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation
	// Register a File Asset.
	//
	// Returns the parameters that can be used to refer to the asset inside the template.
	AddFileAsset(_asset *FileAssetSource) *FileAssetLocation
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

// The jsii proxy struct for BootstraplessSynthesizer
type jsiiProxy_BootstraplessSynthesizer struct {
	jsiiProxy_DefaultStackSynthesizer
}

func (j *jsiiProxy_BootstraplessSynthesizer) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) DeployRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deployRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BootstraplessSynthesizer) Stack() Stack {
	var returns Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBootstraplessSynthesizer(props *BootstraplessSynthesizerProps) BootstraplessSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_BootstraplessSynthesizer{}

	_jsii_.Create(
		"aws-cdk-lib.BootstraplessSynthesizer",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBootstraplessSynthesizer_Override(b BootstraplessSynthesizer, props *BootstraplessSynthesizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.BootstraplessSynthesizer",
		[]interface{}{props},
		b,
	)
}

func BootstraplessSynthesizer_DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_BOOTSTRAP_STACK_VERSION_SSM_PARAMETER",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_DOCKER_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_DOCKER_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_KEY_ARN_EXPORT_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PREFIX() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PREFIX",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_FILE_ASSETS_BUCKET_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_FILE_ASSETS_BUCKET_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSET_PUBLISHING_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func BootstraplessSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.BootstraplessSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddDockerImageAsset(_asset *DockerImageAssetSource) *DockerImageAssetLocation {
	var returns *DockerImageAssetLocation

	_jsii_.Invoke(
		b,
		"addDockerImageAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) AddFileAsset(_asset *FileAssetSource) *FileAssetLocation {
	var returns *FileAssetLocation

	_jsii_.Invoke(
		b,
		"addFileAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BootstraplessSynthesizer) Bind(stack Stack) {
	_jsii_.InvokeVoid(
		b,
		"bind",
		[]interface{}{stack},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) EmitStackArtifact(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) {
	_jsii_.InvokeVoid(
		b,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BootstraplessSynthesizer) SynthesizeStackTemplate(stack Stack, session ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}
