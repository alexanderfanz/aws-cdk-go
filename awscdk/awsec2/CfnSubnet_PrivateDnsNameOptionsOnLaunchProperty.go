package awsec2


// Describes the options for instance hostnames.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateDnsNameOptionsOnLaunchProperty := &PrivateDnsNameOptionsOnLaunchProperty{
//   	EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   	EnableResourceNameDnsARecord: jsii.Boolean(false),
//   	HostnameType: jsii.String("hostnameType"),
//   }
//
type CfnSubnet_PrivateDnsNameOptionsOnLaunchProperty struct {
	// Indicates whether to respond to DNS queries for instance hostname with DNS AAAA records.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// The type of hostname for EC2 instances.
	//
	// For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 only subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

