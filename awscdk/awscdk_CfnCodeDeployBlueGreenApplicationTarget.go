// An experiment to bundle the entire CDK into a single module
package awscdk


// Type of the {@link CfnCodeDeployBlueGreenApplication.target} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenApplicationTarget := &cfnCodeDeployBlueGreenApplicationTarget{
//   	logicalId: jsii.String("logicalId"),
//   	type: jsii.String("type"),
//   }
//
// Experimental.
type CfnCodeDeployBlueGreenApplicationTarget struct {
	// The logical id of the target resource.
	// Experimental.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The resource type of the target being deployed.
	//
	// Right now, the only allowed value is 'AWS::ECS::Service'.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

