package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// A reference to an existing task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   taskDefinitionAttributes := &taskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	compatibility: awscdk.Aws_ecs.compatibility_EC2,
//   	networkMode: awscdk.*Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
// Experimental.
type TaskDefinitionAttributes struct {
	// The arn of the task definition.
	// Experimental.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	// Experimental.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// What launch types this task definition should be compatible with.
	// Experimental.
	Compatibility Compatibility `field:"optional" json:"compatibility" yaml:"compatibility"`
}

