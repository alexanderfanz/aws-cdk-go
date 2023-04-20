package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateSubnetProps := &PrivateSubnetProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	MapPublicIpOnLaunch: jsii.Boolean(false),
//   }
//
// Experimental.
type PrivateSubnetProps struct {
	// The availability zone for the subnet.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The CIDR notation for this subnet.
	// Experimental.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The VPC which this subnet is part of.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Controls if a public IP is associated to an instance at launch.
	// Experimental.
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
}

