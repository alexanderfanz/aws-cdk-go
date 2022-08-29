package awslambdanodejs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for a NodejsFunction.
//
// Example:
//   lambda.NewNodejsFunction(this, jsii.String("my-handler"), &nodejsFunctionProps{
//   	bundling: &bundlingOptions{
//   		minify: jsii.Boolean(true),
//   		 // minify code, defaults to false
//   		sourceMap: jsii.Boolean(true),
//   		 // include source map, defaults to false
//   		sourceMapMode: lambda.sourceMapMode_INLINE,
//   		 // defaults to SourceMapMode.DEFAULT
//   		sourcesContent: jsii.Boolean(false),
//   		 // do not include original source into source map, defaults to true
//   		target: jsii.String("es2020"),
//   		 // target environment for the generated JavaScript code
//   		loader: map[string]*string{
//   			 // Use the 'dataurl' loader for '.png' files
//   			".png": jsii.String("dataurl"),
//   		},
//   		define: map[string]*string{
//   			 // Replace strings during build time
//   			"process.env.API_KEY": JSON.stringify(jsii.String("xxx-xxxx-xxx")),
//   			"process.env.PRODUCTION": JSON.stringify(jsii.Boolean(true)),
//   			"process.env.NUMBER": JSON.stringify(jsii.Number(123)),
//   		},
//   		logLevel: lambda.logLevel_SILENT,
//   		 // defaults to LogLevel.WARNING
//   		keepNames: jsii.Boolean(true),
//   		 // defaults to false
//   		tsconfig: jsii.String("custom-tsconfig.json"),
//   		 // use custom-tsconfig.json instead of default,
//   		metafile: jsii.Boolean(true),
//   		 // include meta file, defaults to false
//   		banner: jsii.String("/* comments */"),
//   		 // requires esbuild >= 0.9.0, defaults to none
//   		footer: jsii.String("/* comments */"),
//   		 // requires esbuild >= 0.9.0, defaults to none
//   		charset: lambda.charset_UTF8,
//   		 // do not escape non-ASCII characters, defaults to Charset.ASCII
//   		format: lambda.outputFormat_ESM,
//   		 // ECMAScript module output format, defaults to OutputFormat.CJS (OutputFormat.ESM requires Node.js 14.x)
//   		mainFields: []*string{
//   			jsii.String("module"),
//   			jsii.String("main"),
//   		},
//   		 // prefer ECMAScript versions of dependencies
//   		inject: []*string{
//   			jsii.String("./my-shim.js"),
//   			jsii.String("./other-shim.js"),
//   		},
//   		 // allows to automatically replace a global variable with an import from another file
//   		esbuildArgs: map[string]interface{}{
//   			 // Pass additional arguments to esbuild
//   			"--log-limit": jsii.String("0"),
//   			"--splitting": jsii.Boolean(true),
//   		},
//   	},
//   })
//
type NodejsFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure awslambda.IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	OnSuccess awslambda.IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `field:"optional" json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture awslambda.Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig awslambda.ICodeSigningConfig `field:"optional" json:"codeSigningConfig" yaml:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *awslambda.VersionOptions `field:"optional" json:"currentVersionOptions" yaml:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	//
	// If SNS topic is desired, specify `deadLetterTopic` property instead.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `field:"optional" json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// The SNS topic to use as a DLQ.
	//
	// Note that if `deadLetterQueueEnabled` is set to `true`, an SQS queue will be created
	// rather than an SNS topic. Using an SNS topic as a DLQ requires this property to be set explicitly.
	DeadLetterTopic awssns.ITopic `field:"optional" json:"deadLetterTopic" yaml:"deadLetterTopic"`
	// A description of the function.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `field:"optional" json:"environmentEncryption" yaml:"environmentEncryption"`
	// The size of the function’s /tmp directory in MiB.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]awslambda.IEventSource `field:"optional" json:"events" yaml:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem awslambda.FileSystem `field:"optional" json:"filesystem" yaml:"filesystem"`
	// A name for the function.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `field:"optional" json:"initialPolicy" yaml:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion awslambda.LambdaInsightsVersion `field:"optional" json:"insightsVersion" yaml:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]awslambda.ILayerVersion `field:"optional" json:"layers" yaml:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *awslambda.LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `field:"optional" json:"profiling" yaml:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `field:"optional" json:"profilingGroup" yaml:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing awslambda.Tracing `field:"optional" json:"tracing" yaml:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	// This is required when `vpcSubnets` is specified.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// This requires `vpc` to be specified in order for interfaces to actually be
	// placed in the subnets. If `vpc` is not specify, this will raise an error.
	//
	// Note: Internet access for Lambda Functions requires a NAT Gateway, so picking
	// public subnets is not allowed (unless `allowPublicSubnet` is set to `true`).
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Whether to automatically reuse TCP connections when working with the AWS SDK for JavaScript.
	//
	// This sets the `AWS_NODEJS_CONNECTION_REUSE_ENABLED` environment variable
	// to `1`.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/node-reusing-connections.html
	//
	AwsSdkConnectionReuse *bool `field:"optional" json:"awsSdkConnectionReuse" yaml:"awsSdkConnectionReuse"`
	// Bundling options.
	Bundling *BundlingOptions `field:"optional" json:"bundling" yaml:"bundling"`
	// The path to the dependencies lock file (`yarn.lock` or `package-lock.json`).
	//
	// This will be used as the source for the volume mounted in the Docker
	// container.
	//
	// Modules specified in `nodeModules` will be installed using the right
	// installer (`npm` or `yarn`) along with this lock file.
	DepsLockFilePath *string `field:"optional" json:"depsLockFilePath" yaml:"depsLockFilePath"`
	// Path to the entry file (JavaScript or TypeScript).
	Entry *string `field:"optional" json:"entry" yaml:"entry"`
	// The name of the exported handler in the entry file.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// The path to the directory containing project config files (`package.json` or `tsconfig.json`).
	ProjectRoot *string `field:"optional" json:"projectRoot" yaml:"projectRoot"`
	// The runtime environment.
	//
	// Only runtimes of the Node.js family are
	// supported.
	Runtime awslambda.Runtime `field:"optional" json:"runtime" yaml:"runtime"`
}
