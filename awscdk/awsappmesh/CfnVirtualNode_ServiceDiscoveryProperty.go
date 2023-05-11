package awsappmesh


// An object that represents the service discovery information for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceDiscoveryProperty := &ServiceDiscoveryProperty{
//   	AwsCloudMap: &AwsCloudMapServiceDiscoveryProperty{
//   		NamespaceName: jsii.String("namespaceName"),
//   		ServiceName: jsii.String("serviceName"),
//
//   		// the properties below are optional
//   		Attributes: []interface{}{
//   			&AwsCloudMapInstanceAttributeProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		IpPreference: jsii.String("ipPreference"),
//   	},
//   	Dns: &DnsServiceDiscoveryProperty{
//   		Hostname: jsii.String("hostname"),
//
//   		// the properties below are optional
//   		IpPreference: jsii.String("ipPreference"),
//   		ResponseType: jsii.String("responseType"),
//   	},
//   }
//
type CfnVirtualNode_ServiceDiscoveryProperty struct {
	// Specifies any AWS Cloud Map information for the virtual node.
	AwsCloudMap interface{} `field:"optional" json:"awsCloudMap" yaml:"awsCloudMap"`
	// Specifies the DNS information for the virtual node.
	Dns interface{} `field:"optional" json:"dns" yaml:"dns"`
}
