package awsmediaconnect


// The bridge's source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeSourceProperty := &BridgeSourceProperty{
//   	FlowSource: &BridgeFlowSourceProperty{
//   		FlowArn: jsii.String("flowArn"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	NetworkSource: &BridgeNetworkSourceProperty{
//   		MulticastIp: jsii.String("multicastIp"),
//   		Name: jsii.String("name"),
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
type CfnBridge_BridgeSourceProperty struct {
	// The source of the bridge.
	//
	// A flow source originates in MediaConnect as an existing cloud flow.
	FlowSource interface{} `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the bridge.
	//
	// A network source originates at your premises.
	NetworkSource interface{} `field:"optional" json:"networkSource" yaml:"networkSource"`
}

