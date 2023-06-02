package awscdk


// Docker build options.
//
// Example:
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(jsii.String("/path/to/handler"), &AssetOptions{
//   		Bundling: &BundlingOptions{
//   			Image: awscdk.DockerImage_FromBuild(jsii.String("/path/to/dir/with/DockerFile"), &DockerBuildOptions{
//   				BuildArgs: map[string]*string{
//   					"ARG1": jsii.String("value1"),
//   				},
//   			}),
//   			Command: []*string{
//   				jsii.String("my"),
//   				jsii.String("cool"),
//   				jsii.String("command"),
//   			},
//   		},
//   	}),
//   	Runtime: lambda.Runtime_PYTHON_3_9(),
//   	Handler: jsii.String("index.handler"),
//   })
//
type DockerBuildOptions struct {
	// Build args.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Set build target for multi-stage container builds. Any stage defined afterwards will be ignored.
	//
	// Example value: `build-env`.
	TargetStage *string `field:"optional" json:"targetStage" yaml:"targetStage"`
}

