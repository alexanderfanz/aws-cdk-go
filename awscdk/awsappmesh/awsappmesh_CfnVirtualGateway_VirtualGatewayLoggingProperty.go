package awsappmesh


// An object that represents logging information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayLoggingProperty := &virtualGatewayLoggingProperty{
//   	accessLog: &virtualGatewayAccessLogProperty{
//   		file: &virtualGatewayFileAccessLogProperty{
//   			path: jsii.String("path"),
//   		},
//   	},
//   }
//
type CfnVirtualGateway_VirtualGatewayLoggingProperty struct {
	// The access log configuration.
	AccessLog interface{} `field:"optional" json:"accessLog" yaml:"accessLog"`
}
