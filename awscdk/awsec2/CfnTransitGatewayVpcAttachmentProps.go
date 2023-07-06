package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGatewayVpcAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   cfnTransitGatewayVpcAttachmentProps := &CfnTransitGatewayVpcAttachmentProps{
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	AddSubnetIds: []*string{
//   		jsii.String("addSubnetIds"),
//   	},
//   	Options: options,
//   	RemoveSubnetIds: []*string{
//   		jsii.String("removeSubnetIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayVpcAttachmentProps struct {
	// The IDs of the subnets.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the transit gateway.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The IDs of one or more subnets to add.
	//
	// You can specify at most one subnet per Availability Zone.
	AddSubnetIds *[]*string `field:"optional" json:"addSubnetIds" yaml:"addSubnetIds"`
	// The VPC attachment options.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The IDs of one or more subnets to remove.
	RemoveSubnetIds *[]*string `field:"optional" json:"removeSubnetIds" yaml:"removeSubnetIds"`
	// The tags for the VPC attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

