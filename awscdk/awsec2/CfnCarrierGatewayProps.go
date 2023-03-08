package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCarrierGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCarrierGatewayProps := &CfnCarrierGatewayProps{
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCarrierGatewayProps struct {
	// The ID of the VPC associated with the carrier gateway.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The tags assigned to the carrier gateway.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

