package awsecs


// The Service Connect service object configuration.
//
// For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectServiceProperty := &serviceConnectServiceProperty{
//   	portName: jsii.String("portName"),
//
//   	// the properties below are optional
//   	clientAliases: []interface{}{
//   		&serviceConnectClientAliasProperty{
//   			port: jsii.Number(123),
//
//   			// the properties below are optional
//   			dnsName: jsii.String("dnsName"),
//   		},
//   	},
//   	discoveryName: jsii.String("discoveryName"),
//   	ingressPortOverride: jsii.Number(123),
//   }
//
type CfnService_ServiceConnectServiceProperty struct {
	// The `portName` must match the name of one of the `portMappings` from all the containers in the task definition of this Amazon ECS service.
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// The list of client aliases for this Service Connect service.
	//
	// You use these to assign names that can be used by client applications. The maximum number of client aliases that you can have in this list is 1.
	//
	// Each alias ("endpoint") is a fully-qualified name and port number that other Amazon ECS tasks ("clients") can use to connect to this service.
	//
	// Each name and port mapping must be unique within the namespace.
	//
	// For each `ServiceConnectService` , you must provide at least one `clientAlias` with one `port` .
	ClientAliases interface{} `field:"optional" json:"clientAliases" yaml:"clientAliases"`
	// The `discoveryName` is the name of the new AWS Cloud Map service that Amazon ECS creates for this Amazon ECS service.
	//
	// This must be unique within the AWS Cloud Map namespace. The name can contain up to 64 characters. The name can include lowercase letters, numbers, underscores (_), and hyphens (-). The name can't start with a hyphen.
	//
	// If this parameter isn't specified, the default value of `discoveryName.namespace` is used. If the `discoveryName` isn't specified, the port mapping name from the task definition is used in `portName.namespace` .
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// The port number for the Service Connect proxy to listen on.
	//
	// Use the value of this field to bypass the proxy for traffic on the port number specified in the named `portMapping` in the task definition of this application, and then use it in your VPC security groups to allow traffic into the proxy for this Amazon ECS service.
	//
	// In `awsvpc` mode and Fargate, the default value is the container port number. The container port number is in the `portMapping` in the task definition. In bridge mode, the default value is the ephemeral port of the Service Connect proxy.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
}
