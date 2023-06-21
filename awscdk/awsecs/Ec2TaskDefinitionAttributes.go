package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Attributes used to import an existing EC2 task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   ec2TaskDefinitionAttributes := &Ec2TaskDefinitionAttributes{
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	ExecutionRole: role,
//   	NetworkMode: awscdk.Aws_ecs.NetworkMode_NONE,
//   	TaskRole: role,
//   }
//
type Ec2TaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The IAM role that grants containers and Fargate agents permission to make AWS API calls on your behalf.
	//
	// Some tasks do not have an execution role.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

