package awsmediaconnect


// Properties for defining a `CfnFlowOutput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlowOutputProps := &CfnFlowOutputProps{
//   	FlowArn: jsii.String("flowArn"),
//   	Protocol: jsii.String("protocol"),
//
//   	// the properties below are optional
//   	CidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	Description: jsii.String("description"),
//   	Destination: jsii.String("destination"),
//   	Encryption: &EncryptionProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//
//   		// the properties below are optional
//   		Algorithm: jsii.String("algorithm"),
//   		KeyType: jsii.String("keyType"),
//   	},
//   	MaxLatency: jsii.Number(123),
//   	MinLatency: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Port: jsii.Number(123),
//   	RemoteId: jsii.String("remoteId"),
//   	SmoothingLatency: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
type CfnFlowOutputProps struct {
	// The Amazon Resource Name (ARN) of the flow this output is attached to.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The protocol to use for the output.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The range of IP addresses that are allowed to initiate output requests to this flow.
	//
	// Format the IP addresses as a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// This description is not visible outside of the current AWS account even if the account grants entitlements to other accounts.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP address where you want to send the output.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The encryption credentials that you want to use for the output.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds.
	//
	// This parameter applies only to RIST-based, Zixi-based, and Fujitsu-based streams.
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds for SRT-based streams.
	//
	// In streams that use the SRT protocol, this value that you set on your MediaConnect source or output represents the minimal potential latency of that connection. The latency of the stream is set to the highest number between the sender’s minimum latency and the receiver’s minimum latency.
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the VPC interface.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port to use when MediaConnect distributes content to the output.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The identifier that is assigned to the Zixi receiver.
	//
	// This parameter applies only to outputs that use Zixi pull.
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *float64 `field:"optional" json:"smoothingLatency" yaml:"smoothingLatency"`
	// The stream ID that you want to use for this transport.
	//
	// This parameter applies only to Zixi and SRT caller-based streams.
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The VPC interface that you want to send your output to.
	VpcInterfaceAttachment interface{} `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

