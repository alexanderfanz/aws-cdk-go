package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionProps := &CfnFunctionProps{
//   	Code: &CodeProperty{
//   		ImageUri: jsii.String("imageUri"),
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		ZipFile: jsii.String("zipFile"),
//   	},
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	CodeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   	Description: jsii.String("description"),
//   	Environment: &EnvironmentProperty{
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		Size: jsii.Number(123),
//   	},
//   	FileSystemConfigs: []interface{}{
//   		&FileSystemConfigProperty{
//   			Arn: jsii.String("arn"),
//   			LocalMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
//   	Handler: jsii.String("handler"),
//   	ImageConfig: &ImageConfigProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EntryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	MemorySize: jsii.Number(123),
//   	PackageType: jsii.String("packageType"),
//   	ReservedConcurrentExecutions: jsii.Number(123),
//   	Runtime: jsii.String("runtime"),
//   	RuntimeManagementConfig: &RuntimeManagementConfigProperty{
//   		UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//
//   		// the properties below are optional
//   		RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   	},
//   	SnapStart: &SnapStartProperty{
//   		ApplyOn: jsii.String("applyOn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   	TracingConfig: &TracingConfigProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnFunctionProps struct {
	// The code for the function.
	Code interface{} `field:"required" json:"code" yaml:"code"`
	// The Amazon Resource Name (ARN) of the function's execution role.
	Role *string `field:"required" json:"role" yaml:"role"`
	// The instruction set architecture that the function supports.
	//
	// Enter a string array with one of the valid values (arm64 or x86_64). The default value is `x86_64` .
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// To enable code signing for this function, specify the ARN of a code-signing configuration.
	//
	// A code-signing configuration
	// includes a set of signing profiles, which define the trusted publishers for this function.
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.
	//
	// For more information, see [Dead-letter queues](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq) .
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// A description of the function.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Environment variables that are accessible from function code during execution.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The size of the function's `/tmp` directory in MB.
	//
	// The default value is 512, but it can be any whole number between 512 and 10,240 MB.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Connection settings for an Amazon EFS file system.
	//
	// To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) resource, you must also specify a `DependsOn` attribute to ensure that the mount target is created or updated before the function.
	//
	// For more information about using the `DependsOn` attribute, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	FileSystemConfigs interface{} `field:"optional" json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// The name of the Lambda function, up to 64 characters in length.
	//
	// If you don't specify a name, AWS CloudFormation generates one.
	//
	// If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The name of the method within your code that Lambda calls to run your function.
	//
	// Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Lambda programming model](https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html) .
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Configuration values that override the container image Dockerfile settings.
	//
	// For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
	ImageConfig interface{} `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// The ARN of the AWS Key Management Service ( AWS KMS ) customer managed key that's used to encrypt your function's [environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption) . When [Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html) is activated, this key is also used to encrypt your function's snapshot. If you don't provide a customer managed key, Lambda uses a default service key.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// The type of deployment package.
	//
	// Set to `Image` for container image and set `Zip` for .zip file archive.
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// The number of simultaneous executions to reserve for the function.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) . Runtime is required if the deployment package is a .zip file archive.
	//
	// The following list includes deprecated runtimes. For more information, see [Runtime deprecation policy](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy) .
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Sets the runtime management configuration for a function's version.
	//
	// For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html) .
	RuntimeManagementConfig interface{} `field:"optional" json:"runtimeManagementConfig" yaml:"runtimeManagementConfig"`
	// The function's [AWS Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.
	SnapStart interface{} `field:"optional" json:"snapStart" yaml:"snapStart"`
	// A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The amount of time (in seconds) that Lambda allows a function to run before stopping it.
	//
	// The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html) .
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Set `Mode` to `Active` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) .
	TracingConfig interface{} `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
	// For network connectivity to AWS resources in a [VPC](https://docs.aws.amazon.com/lambda/latest/dg/configuration-network.html) , specify a list of security groups and subnets in the VPC.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

