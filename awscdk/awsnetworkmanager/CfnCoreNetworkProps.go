package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCoreNetwork`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnCoreNetworkProps := &CfnCoreNetworkProps{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PolicyDocument: policyDocument,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCoreNetworkProps struct {
	// The ID of the global network that your core network is a part of.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The description of a core network.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Describes a core network policy. For more information, see [Core network policies](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) .
	//
	// If you update the policy document, CloudFormation will apply the core network change set generated from the updated policy document, and then set it as the LIVE policy.
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The list of key-value tags associated with a core network.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
