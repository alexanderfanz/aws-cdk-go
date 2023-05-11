package awsappmesh


// An object representing the gateway route to rewrite.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpGatewayRouteRewriteProperty := &HttpGatewayRouteRewriteProperty{
//   	Hostname: &GatewayRouteHostnameRewriteProperty{
//   		DefaultTargetHostname: jsii.String("defaultTargetHostname"),
//   	},
//   	Path: &HttpGatewayRoutePathRewriteProperty{
//   		Exact: jsii.String("exact"),
//   	},
//   	Prefix: &HttpGatewayRoutePrefixRewriteProperty{
//   		DefaultPrefix: jsii.String("defaultPrefix"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
type CfnGatewayRoute_HttpGatewayRouteRewriteProperty struct {
	// The host name to rewrite.
	Hostname interface{} `field:"optional" json:"hostname" yaml:"hostname"`
	// The path to rewrite.
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// The specified beginning characters to rewrite.
	Prefix interface{} `field:"optional" json:"prefix" yaml:"prefix"`
}
