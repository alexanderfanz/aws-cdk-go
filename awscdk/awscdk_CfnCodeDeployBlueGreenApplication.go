// An experiment to bundle the entire CDK into a single module
package awscdk


// The application actually being deployed.
//
// Type of the {@link CfnCodeDeployBlueGreenHookProps.applications} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenApplication := &CfnCodeDeployBlueGreenApplication{
//   	EcsAttributes: &CfnCodeDeployBlueGreenEcsAttributes{
//   		TaskDefinitions: []*string{
//   			jsii.String("taskDefinitions"),
//   		},
//   		TaskSets: []*string{
//   			jsii.String("taskSets"),
//   		},
//   		TrafficRouting: &CfnTrafficRouting{
//   			ProdTrafficRoute: &CfnTrafficRoute{
//   				LogicalId: jsii.String("logicalId"),
//   				Type: jsii.String("type"),
//   			},
//   			TargetGroups: []*string{
//   				jsii.String("targetGroups"),
//   			},
//   			TestTrafficRoute: &CfnTrafficRoute{
//   				LogicalId: jsii.String("logicalId"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Target: &CfnCodeDeployBlueGreenApplicationTarget{
//   		LogicalId: jsii.String("logicalId"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// Experimental.
type CfnCodeDeployBlueGreenApplication struct {
	// The detailed attributes of the deployed target.
	// Experimental.
	EcsAttributes *CfnCodeDeployBlueGreenEcsAttributes `field:"required" json:"ecsAttributes" yaml:"ecsAttributes"`
	// The target that is being deployed.
	// Experimental.
	Target *CfnCodeDeployBlueGreenApplicationTarget `field:"required" json:"target" yaml:"target"`
}

