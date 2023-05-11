package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFirewall`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallProps := &CfnFirewallProps{
//   	FirewallName: jsii.String("firewallName"),
//   	FirewallPolicyArn: jsii.String("firewallPolicyArn"),
//   	SubnetMappings: []interface{}{
//   		&SubnetMappingProperty{
//   			SubnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			IpAddressType: jsii.String("ipAddressType"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	DeleteProtection: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	FirewallPolicyChangeProtection: jsii.Boolean(false),
//   	SubnetChangeProtection: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallProps struct {
	// The descriptive name of the firewall.
	//
	// You can't change the name of a firewall after you create it.
	FirewallName *string `field:"required" json:"firewallName" yaml:"firewallName"`
	// The Amazon Resource Name (ARN) of the firewall policy.
	//
	// The relationship of firewall to firewall policy is many to one. Each firewall requires one firewall policy association, and you can use the same firewall policy for multiple firewalls.
	FirewallPolicyArn *string `field:"required" json:"firewallPolicyArn" yaml:"firewallPolicyArn"`
	// The public subnets that Network Firewall is using for the firewall.
	//
	// Each subnet must belong to a different Availability Zone.
	SubnetMappings interface{} `field:"required" json:"subnetMappings" yaml:"subnetMappings"`
	// The unique identifier of the VPC where the firewall is in use.
	//
	// You can't change the VPC of a firewall after you create the firewall.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// A flag indicating whether it is possible to delete the firewall.
	//
	// A setting of `TRUE` indicates that the firewall is protected against deletion. Use this setting to protect against accidentally deleting a firewall that is in use. When you create a firewall, the operation initializes this flag to `TRUE` .
	DeleteProtection interface{} `field:"optional" json:"deleteProtection" yaml:"deleteProtection"`
	// A description of the firewall.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A setting indicating whether the firewall is protected against a change to the firewall policy association.
	//
	// Use this setting to protect against accidentally modifying the firewall policy for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE` .
	FirewallPolicyChangeProtection interface{} `field:"optional" json:"firewallPolicyChangeProtection" yaml:"firewallPolicyChangeProtection"`
	// A setting indicating whether the firewall is protected against changes to the subnet associations.
	//
	// Use this setting to protect against accidentally modifying the subnet associations for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE` .
	SubnetChangeProtection interface{} `field:"optional" json:"subnetChangeProtection" yaml:"subnetChangeProtection"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
