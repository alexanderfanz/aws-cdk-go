package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Construction properties for the BaseRunTaskProps.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var parameters interface{}
//   var taskDefinition taskDefinition
//
//   ecsRunTaskBaseProps := &EcsRunTaskBaseProps{
//   	Cluster: cluster,
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: jsii.Number(123),
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MemoryLimit: jsii.Number(123),
//   			MemoryReservation: jsii.Number(123),
//   		},
//   	},
//   	IntegrationPattern: awscdk.Aws_stepfunctions.ServiceIntegrationPattern_FIRE_AND_FORGET,
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   }
//
// Deprecated: No replacement.
type EcsRunTaskBaseProps struct {
	// The topic to run the task on.
	// Deprecated: No replacement.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions.
	// Deprecated: No replacement.
	TaskDefinition awsecs.TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Deprecated: No replacement.
	ContainerOverrides *[]*ContainerOverride `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// The service integration pattern indicates different ways to call RunTask in ECS.
	//
	// The valid value for Lambda is FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
	// Deprecated: No replacement.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Additional parameters to pass to the base task.
	// Deprecated: No replacement.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

