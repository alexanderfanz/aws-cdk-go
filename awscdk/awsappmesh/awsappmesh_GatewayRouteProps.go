package awsappmesh


// Properties to define a new GatewayRoute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var gatewayRouteSpec gatewayRouteSpec
//   var virtualGateway virtualGateway
//
//   gatewayRouteProps := &gatewayRouteProps{
//   	routeSpec: gatewayRouteSpec,
//   	virtualGateway: virtualGateway,
//
//   	// the properties below are optional
//   	gatewayRouteName: jsii.String("gatewayRouteName"),
//   }
//
// Experimental.
type GatewayRouteProps struct {
	// What protocol the route uses.
	// Experimental.
	RouteSpec GatewayRouteSpec `field:"required" json:"routeSpec" yaml:"routeSpec"`
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName *string `field:"optional" json:"gatewayRouteName" yaml:"gatewayRouteName"`
	// The VirtualGateway this GatewayRoute is associated with.
	// Experimental.
	VirtualGateway IVirtualGateway `field:"required" json:"virtualGateway" yaml:"virtualGateway"`
}

