package awscloudfront


// The traffic configuration of your continuous deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficConfigProperty := &TrafficConfigProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	SingleHeaderConfig: &SingleHeaderConfigProperty{
//   		Header: jsii.String("header"),
//   		Value: jsii.String("value"),
//   	},
//   	SingleWeightConfig: &SingleWeightConfigProperty{
//   		Weight: jsii.Number(123),
//
//   		// the properties below are optional
//   		SessionStickinessConfig: &SessionStickinessConfigProperty{
//   			IdleTtl: jsii.Number(123),
//   			MaximumTtl: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnContinuousDeploymentPolicy_TrafficConfigProperty struct {
	// The type of traffic configuration.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Determines which HTTP requests are sent to the staging distribution.
	SingleHeaderConfig interface{} `field:"optional" json:"singleHeaderConfig" yaml:"singleHeaderConfig"`
	// Contains the percentage of traffic to send to the staging distribution.
	SingleWeightConfig interface{} `field:"optional" json:"singleWeightConfig" yaml:"singleWeightConfig"`
}
