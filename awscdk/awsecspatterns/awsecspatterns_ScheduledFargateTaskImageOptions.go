package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// The properties for the ScheduledFargateTask using an image.
//
// Example:
//   var cluster cluster
//
//   scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &scheduledFargateTaskProps{
//   	cluster: cluster,
//   	scheduledFargateTaskImageOptions: &scheduledFargateTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		memoryLimitMiB: jsii.Number(512),
//   	},
//   	schedule: appscaling.schedule.expression(jsii.String("rate(1 minute)")),
//   	platformVersion: ecs.fargatePlatformVersion_LATEST,
//   })
//
type ScheduledFargateTaskImageOptions struct {
	// The image used to start a container.
	//
	// Image or taskDefinition must be specified, but not both.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The log driver to use.
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The secret to expose to the container as an environment variable.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The runtime platform of the task definition.
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition awsecs.FargateTaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
}
