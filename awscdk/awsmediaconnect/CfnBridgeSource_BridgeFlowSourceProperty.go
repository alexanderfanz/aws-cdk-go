package awsmediaconnect


// The source of the bridge.
//
// A flow source originates in MediaConnect as an existing cloud flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeFlowSourceProperty := &BridgeFlowSourceProperty{
//   	FlowArn: jsii.String("flowArn"),
//
//   	// the properties below are optional
//   	FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
type CfnBridgeSource_BridgeFlowSourceProperty struct {
	// The ARN of the cloud flow used as a source of this bridge.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the VPC interface attachment to use for this source.
	FlowVpcInterfaceAttachment interface{} `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
}

