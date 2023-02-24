package awsiotwireless


// LoRaWANUpdateGatewayTaskEntry object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANUpdateGatewayTaskEntryProperty := &LoRaWANUpdateGatewayTaskEntryProperty{
//   	CurrentVersion: &LoRaWANGatewayVersionProperty{
//   		Model: jsii.String("model"),
//   		PackageVersion: jsii.String("packageVersion"),
//   		Station: jsii.String("station"),
//   	},
//   	UpdateVersion: &LoRaWANGatewayVersionProperty{
//   		Model: jsii.String("model"),
//   		PackageVersion: jsii.String("packageVersion"),
//   		Station: jsii.String("station"),
//   	},
//   }
//
type CfnTaskDefinition_LoRaWANUpdateGatewayTaskEntryProperty struct {
	// The version of the gateways that should receive the update.
	CurrentVersion interface{} `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// The firmware version to update the gateway to.
	UpdateVersion interface{} `field:"optional" json:"updateVersion" yaml:"updateVersion"`
}

