package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Defines HTTP route matching based on the URL path of the request.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http(&HttpRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   				Weight: jsii.Number(50),
//   			},
//   			&weightedTarget{
//   				VirtualNode: node,
//   				Weight: jsii.Number(50),
//   			},
//   		},
//   		Match: &HttpRouteMatch{
//   			Path: appmesh.HttpRoutePathMatch_StartsWith(jsii.String("/path-to-app")),
//   		},
//   	}),
//   })
//
// Experimental.
type HttpRoutePathMatch interface {
	// Returns the route path match configuration.
	// Experimental.
	Bind(scope awscdk.Construct) *HttpRoutePathMatchConfig
}

// The jsii proxy struct for HttpRoutePathMatch
type jsiiProxy_HttpRoutePathMatch struct {
	_ byte // padding
}

// Experimental.
func NewHttpRoutePathMatch_Override(h HttpRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
// Experimental.
func HttpRoutePathMatch_Exactly(path *string) HttpRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpRoutePathMatch_ExactlyParameters(path); err != nil {
		panic(err)
	}
	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"exactly",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
// Experimental.
func HttpRoutePathMatch_Regex(regex *string) HttpRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpRoutePathMatch_RegexParameters(regex); err != nil {
		panic(err)
	}
	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"regex",
		[]interface{}{regex},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
// Experimental.
func HttpRoutePathMatch_StartsWith(prefix *string) HttpRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpRoutePathMatch_StartsWithParameters(prefix); err != nil {
		panic(err)
	}
	var returns HttpRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpRoutePathMatch",
		"startsWith",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpRoutePathMatch) Bind(scope awscdk.Construct) *HttpRoutePathMatchConfig {
	if err := h.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *HttpRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

