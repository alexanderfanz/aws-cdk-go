package cloudassemblyschema


// Options to use with cdk destroy.
//
// Example:
//   app := awscdk.NewApp()
//
//   stackUnderTest := awscdk.NewStack(app, jsii.String("StackUnderTest"))
//
//   stack := awscdk.NewStack(app, jsii.String("stack"))
//
//   testCase := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("CustomizedDeploymentWorkflow"), &IntegTestProps{
//   	TestCases: []stack{
//   		stackUnderTest,
//   	},
//   	DiffAssets: jsii.Boolean(true),
//   	StackUpdateWorkflow: jsii.Boolean(true),
//   	CdkCommandOptions: &CdkCommands{
//   		Deploy: &DeployCommand{
//   			Args: &DeployOptions{
//   				RequireApproval: awscdk.RequireApproval_NEVER,
//   				Json: jsii.Boolean(true),
//   			},
//   		},
//   		Destroy: &DestroyCommand{
//   			Args: &DestroyOptions{
//   				Force: jsii.Boolean(true),
//   			},
//   		},
//   	},
//   })
//
type DestroyOptions struct {
	// Deploy all stacks.
	//
	// Requried if `stacks` is not set.
	// Default: - false.
	//
	All *bool `field:"optional" json:"all" yaml:"all"`
	// command-line for executing your app or a cloud assembly directory e.g. "node bin/my-app.js" or "cdk.out".
	// Default: - read from cdk.json
	//
	App *string `field:"optional" json:"app" yaml:"app"`
	// Include "aws:asset:*" CloudFormation metadata for resources that use assets.
	// Default: true.
	//
	AssetMetadata *bool `field:"optional" json:"assetMetadata" yaml:"assetMetadata"`
	// Path to CA certificate to use when validating HTTPS requests.
	// Default: - read from AWS_CA_BUNDLE environment variable.
	//
	CaBundlePath *string `field:"optional" json:"caBundlePath" yaml:"caBundlePath"`
	// Show colors and other style from console output.
	// Default: true.
	//
	Color *bool `field:"optional" json:"color" yaml:"color"`
	// Additional context.
	// Default: - no additional context.
	//
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// enable emission of additional debugging information, such as creation stack traces of tokens.
	// Default: false.
	//
	Debug *bool `field:"optional" json:"debug" yaml:"debug"`
	// Force trying to fetch EC2 instance credentials.
	// Default: - guess EC2 instance status.
	//
	Ec2Creds *bool `field:"optional" json:"ec2Creds" yaml:"ec2Creds"`
	// Ignores synthesis errors, which will likely produce an invalid output.
	// Default: false.
	//
	IgnoreErrors *bool `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// Use JSON output instead of YAML when templates are printed to STDOUT.
	// Default: false.
	//
	Json *bool `field:"optional" json:"json" yaml:"json"`
	// Perform context lookups.
	//
	// Synthesis fails if this is disabled and context lookups need
	// to be performed.
	// Default: true.
	//
	Lookups *bool `field:"optional" json:"lookups" yaml:"lookups"`
	// Show relevant notices.
	// Default: true.
	//
	Notices *bool `field:"optional" json:"notices" yaml:"notices"`
	// Emits the synthesized cloud assembly into a directory.
	// Default: cdk.out
	//
	Output *string `field:"optional" json:"output" yaml:"output"`
	// Include "aws:cdk:path" CloudFormation metadata for each resource.
	// Default: true.
	//
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// Use the indicated AWS profile as the default environment.
	// Default: - no profile is used.
	//
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Use the indicated proxy.
	//
	// Will read from
	// HTTPS_PROXY environment if specified.
	// Default: - no proxy.
	//
	Proxy *string `field:"optional" json:"proxy" yaml:"proxy"`
	// Role to pass to CloudFormation for deployment.
	// Default: - use the bootstrap cfn-exec role.
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// List of stacks to deploy.
	//
	// Requried if `all` is not set.
	// Default: - [].
	//
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
	// Copy assets to the output directory.
	//
	// Needed for local debugging the source files with SAM CLI.
	// Default: false.
	//
	Staging *bool `field:"optional" json:"staging" yaml:"staging"`
	// Do not construct stacks with warnings.
	// Default: false.
	//
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// Print trace for stack warnings.
	// Default: false.
	//
	Trace *bool `field:"optional" json:"trace" yaml:"trace"`
	// show debug logs.
	// Default: false.
	//
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Include "AWS::CDK::Metadata" resource in synthesized templates.
	// Default: true.
	//
	VersionReporting *bool `field:"optional" json:"versionReporting" yaml:"versionReporting"`
	// Only destroy the given stack.
	// Default: false.
	//
	Exclusively *bool `field:"optional" json:"exclusively" yaml:"exclusively"`
	// Do not ask for permission before destroying stacks.
	// Default: false.
	//
	Force *bool `field:"optional" json:"force" yaml:"force"`
}

