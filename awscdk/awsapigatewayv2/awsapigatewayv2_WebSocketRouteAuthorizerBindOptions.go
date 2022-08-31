package awsapigatewayv2

import (
	"github.com/aws/constructs-go/constructs/v3"
)

// Input to the bind() operation, that binds an authorizer to a route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import constructs "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//   var webSocketRoute webSocketRoute
//
//   webSocketRouteAuthorizerBindOptions := &webSocketRouteAuthorizerBindOptions{
//   	route: webSocketRoute,
//   	scope: construct,
//   }
//
// Experimental.
type WebSocketRouteAuthorizerBindOptions struct {
	// The route to which the authorizer is being bound.
	// Experimental.
	Route IWebSocketRoute `field:"required" json:"route" yaml:"route"`
	// The scope for any constructs created as part of the bind.
	// Experimental.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}
