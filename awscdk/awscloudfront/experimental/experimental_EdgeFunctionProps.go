package experimental

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Properties for creating a Lambda@Edge function.
//
// Example:
//   var myBucket bucket
//   // A Lambda@Edge function added to default behavior of a Distribution
//   // and triggered on every request
//   myFunc := #error#.NewEdgeFunction(this, jsii.String("MyFunction"), &edgeFunctionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(myBucket),
//   		edgeLambdas: []edgeLambda{
//   			&edgeLambda{
//   				functionVersion: myFunc.currentVersion,
//   				eventType: cloudfront.lambdaEdgeEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type EdgeFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	// Experimental.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	// Experimental.
	OnFailure awslambda.IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	// Experimental.
	OnSuccess awslambda.IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	// Experimental.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	// Experimental.
	AllowPublicSubnet *bool `field:"optional" json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	// Experimental.
	Architecture awslambda.Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// DEPRECATED.
	// Deprecated: use `architecture`.
	Architectures *[]awslambda.Architecture `field:"optional" json:"architectures" yaml:"architectures"`
	// Code signing config associated with this function.
	// Experimental.
	CodeSigningConfig awslambda.ICodeSigningConfig `field:"optional" json:"codeSigningConfig" yaml:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	// Experimental.
	CurrentVersionOptions *awslambda.VersionOptions `field:"optional" json:"currentVersionOptions" yaml:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	//
	// If SNS topic is desired, specify `deadLetterTopic` property instead.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	// Experimental.
	DeadLetterQueueEnabled *bool `field:"optional" json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// The SNS topic to use as a DLQ.
	//
	// Note that if `deadLetterQueueEnabled` is set to `true`, an SQS queue will be created
	// rather than an SNS topic. Using an SNS topic as a DLQ requires this property to be set explicitly.
	// Experimental.
	DeadLetterTopic awssns.ITopic `field:"optional" json:"deadLetterTopic" yaml:"deadLetterTopic"`
	// A description of the function.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	// Experimental.
	EnvironmentEncryption awskms.IKey `field:"optional" json:"environmentEncryption" yaml:"environmentEncryption"`
	// The size of the function’s /tmp directory in MiB.
	// Experimental.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	// Experimental.
	Events *[]awslambda.IEventSource `field:"optional" json:"events" yaml:"events"`
	// The filesystem configuration for the lambda function.
	// Experimental.
	Filesystem awslambda.FileSystem `field:"optional" json:"filesystem" yaml:"filesystem"`
	// A name for the function.
	// Experimental.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	// Experimental.
	InitialPolicy *[]awsiam.PolicyStatement `field:"optional" json:"initialPolicy" yaml:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	// Experimental.
	InsightsVersion awslambda.LambdaInsightsVersion `field:"optional" json:"insightsVersion" yaml:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	// Experimental.
	Layers *[]awslambda.ILayerVersion `field:"optional" json:"layers" yaml:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	// Experimental.
	LogRetentionRetryOptions *awslambda.LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	// Experimental.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	// Experimental.
	Profiling *bool `field:"optional" json:"profiling" yaml:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	// Experimental.
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `field:"optional" json:"profilingGroup" yaml:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	// Experimental.
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
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// What security group to associate with the Lambda's network interfaces. This property is being deprecated, consider using securityGroups instead.
	//
	// Only used if 'vpc' is supplied.
	//
	// Use securityGroups property instead.
	// Function constructor will throw an error if both are specified.
	// Deprecated: - This property is deprecated, use securityGroups instead.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	// Experimental.
	Tracing awslambda.Tracing `field:"optional" json:"tracing" yaml:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	// Experimental.
	Code awslambda.Code `field:"required" json:"code" yaml:"code"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include
	// namespaces and other qualifiers, depending on the runtime.
	// For more information, see https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-features.html#gettingstarted-features-programmingmodel.
	//
	// Use `Handler.FROM_IMAGE` when defining a function from a Docker image.
	//
	// NOTE: If you specify your source code as inline text by specifying the
	// ZipFile property within the Code property, specify index.function_name as
	// the handler.
	// Experimental.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// The runtime environment for the Lambda function that you are uploading.
	//
	// For valid values, see the Runtime property in the AWS Lambda Developer
	// Guide.
	//
	// Use `Runtime.FROM_IMAGE` when when defining a function from a Docker image.
	// Experimental.
	Runtime awslambda.Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// The stack ID of Lambda@Edge function.
	// Experimental.
	StackId *string `field:"optional" json:"stackId" yaml:"stackId"`
}

