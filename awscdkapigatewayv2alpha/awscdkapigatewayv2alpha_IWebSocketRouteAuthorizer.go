// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An authorizer that can attach to an WebSocket Route.
// Experimental.
type IWebSocketRouteAuthorizer interface {
	// Bind this authorizer to a specified WebSocket route.
	// Experimental.
	Bind(options *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig
}

// The jsii proxy for IWebSocketRouteAuthorizer
type jsiiProxy_IWebSocketRouteAuthorizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IWebSocketRouteAuthorizer) Bind(options *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig {
	var returns *WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}
