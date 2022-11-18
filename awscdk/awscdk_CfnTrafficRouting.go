// An experiment to bundle the entire CDK into a single module
package awscdk


// Type of the {@link CfnCodeDeployBlueGreenEcsAttributes.trafficRouting} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRouting := &cfnTrafficRouting{
//   	prodTrafficRoute: &cfnTrafficRoute{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   	targetGroups: []*string{
//   		jsii.String("targetGroups"),
//   	},
//   	testTrafficRoute: &cfnTrafficRoute{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   }
//
// Experimental.
type CfnTrafficRouting struct {
	// The listener to be used by your load balancer to direct traffic to your target groups.
	// Experimental.
	ProdTrafficRoute *CfnTrafficRoute `field:"required" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// The logical IDs of the blue and green, respectively, AWS::ElasticLoadBalancingV2::TargetGroup target groups.
	// Experimental.
	TargetGroups *[]*string `field:"required" json:"targetGroups" yaml:"targetGroups"`
	// The listener to be used by your load balancer to direct traffic to your target groups.
	// Experimental.
	TestTrafficRoute *CfnTrafficRoute `field:"required" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

