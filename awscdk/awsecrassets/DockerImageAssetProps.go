package awsecrassets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for DockerImageAssets.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &DockerImageAssetProps{
//   	Directory: path.join(__dirname, jsii.String("my-image")),
//   	BuildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	Invalidation: &DockerImageAssetInvalidationOptions{
//   		BuildArgs: jsii.Boolean(false),
//   	},
//   })
//
type DockerImageAssetProps struct {
	// File paths matching the patterns will be excluded.
	//
	// See `ignoreMode` to set the matching behavior.
	// Has no effect on Assets bundled using the `bundling` property.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for `exclude` patterns.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Unique identifier of the docker image asset and its potential revisions.
	//
	// Required if using AppScopedStagingSynthesizer.
	AssetName *string `field:"optional" json:"assetName" yaml:"assetName"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Build secrets.
	//
	// Docker BuildKit must be enabled to use build secrets.
	//
	// Example:
	//   import { DockerBuildSecret } from 'aws-cdk-lib';
	//
	//   const buildSecrets = {
	//     'MY_SECRET': DockerBuildSecret.fromSrc('file.txt')
	//   };
	//
	// See: https://docs.docker.com/build/buildkit/
	//
	BuildSecrets *map[string]*string `field:"optional" json:"buildSecrets" yaml:"buildSecrets"`
	// Cache from options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	CacheFrom *[]*DockerCacheOption `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache to options to pass to the `docker build` command.
	// See: https://docs.docker.com/build/cache/backends/
	//
	CacheTo *DockerCacheOption `field:"optional" json:"cacheTo" yaml:"cacheTo"`
	// Path to the Dockerfile (relative to the directory).
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Outputs to pass to the `docker build` command.
	// See: https://docs.docker.com/engine/reference/commandline/build/#custom-build-outputs
	//
	Outputs *[]*string `field:"optional" json:"outputs" yaml:"outputs"`
	// Platform to build for.
	//
	// _Requires Docker Buildx_.
	Platform Platform `field:"optional" json:"platform" yaml:"platform"`
	// Docker target to build to.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The directory where the Dockerfile is stored.
	//
	// Any directory inside with a name that matches the CDK output folder (cdk.out by default) will be excluded from the asset
	Directory *string `field:"required" json:"directory" yaml:"directory"`
}

