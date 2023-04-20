package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var service service
//
//   aliasTargetInstanceProps := &AliasTargetInstanceProps{
//   	DnsName: jsii.String("dnsName"),
//   	Service: service,
//
//   	// the properties below are optional
//   	CustomAttributes: map[string]*string{
//   		"customAttributesKey": jsii.String("customAttributes"),
//   	},
//   	InstanceId: jsii.String("instanceId"),
//   }
//
// Experimental.
type AliasTargetInstanceProps struct {
	// Custom attributes of the instance.
	// Experimental.
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	// Experimental.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// DNS name of the target.
	// Experimental.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The Cloudmap service this resource is registered to.
	// Experimental.
	Service IService `field:"required" json:"service" yaml:"service"`
}

