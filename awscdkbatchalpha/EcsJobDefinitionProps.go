// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for EcsJobDefinition.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import efs "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFileSystem iFileSystem
//
//
//   jobDefn := batch.NewEcsJobDefinition(this, jsii.String("JobDefn"), &EcsJobDefinitionProps{
//   	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("containerDefn"), &EcsEc2ContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/amazonlinux/amazonlinux:latest")),
//   		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
//   		Cpu: jsii.Number(256),
//   		Volumes: []ecsVolume{
//   			batch.*ecsVolume_Efs(&EfsVolumeOptions{
//   				Name: jsii.String("myVolume"),
//   				FileSystem: myFileSystem,
//   				ContainerPath: jsii.String("/Volumes/myVolume"),
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type EcsJobDefinitionProps struct {
	// The name of this job definition.
	// Experimental.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The default parameters passed to the container These parameters can be referenced in the `command` that you give to the container.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html#parameters
	//
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The number of times to retry a job.
	//
	// The job is retried on failure the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Defines the retry behavior for this job.
	// Experimental.
	RetryStrategies *[]RetryStrategy `field:"optional" json:"retryStrategies" yaml:"retryStrategies"`
	// The priority of this Job.
	//
	// Only used in Fairshare Scheduling
	// to decide which job to run first when there are multiple jobs
	// with the same share identifier.
	// Experimental.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes,
	// Batch terminates your jobs if they aren't finished.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The container that this job will run.
	// Experimental.
	Container IEcsContainerDefinition `field:"required" json:"container" yaml:"container"`
	// Whether to propogate tags from the JobDefinition to the ECS task that Batch spawns.
	// Experimental.
	PropagateTags *bool `field:"optional" json:"propagateTags" yaml:"propagateTags"`
}
