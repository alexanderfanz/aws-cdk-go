package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

// An image that will be built from a local directory with a Dockerfile.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   import ecsPatterns "github.com/aws/aws-cdk-go/awscdk"
//   import cxapi "github.com/aws/aws-cdk-go/awscdk"
//   import path "github.com/aws-samples/dummy/path"
//
//   app := awscdk.NewApp()
//
//   stack := awscdk.NewStack(app, jsii.String("aws-ecs-patterns-queue"))
//   stack.node.setContext(cxapi.eCS_REMOVE_DEFAULT_DESIRED_COUNT, jsii.Boolean(true))
//
//   vpc := ec2.NewVpc(stack, jsii.String("VPC"), &vpcProps{
//   	maxAzs: jsii.Number(2),
//   })
//
//   ecsPatterns.NewQueueProcessingFargateService(stack, jsii.String("QueueProcessingService"), &queueProcessingFargateServiceProps{
//   	vpc: vpc,
//   	memoryLimitMiB: jsii.Number(512),
//   	image: ecs.NewAssetImage(path.join(__dirname, jsii.String(".."), jsii.String("sqs-reader"))),
//   })
//
type AssetImage interface {
	ContainerImage
	// Called when the image is used by a ContainerDefinition.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for AssetImage
type jsiiProxy_AssetImage struct {
	jsiiProxy_ContainerImage
}

// Constructs a new instance of the AssetImage class.
func NewAssetImage(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	j := jsiiProxy_AssetImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetImage",
		[]interface{}{directory, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AssetImage class.
func NewAssetImage_Override(a AssetImage, directory *string, props *AssetImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.AssetImage",
		[]interface{}{directory, props},
		a,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func AssetImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func AssetImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func AssetImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func AssetImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func AssetImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.AssetImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}
