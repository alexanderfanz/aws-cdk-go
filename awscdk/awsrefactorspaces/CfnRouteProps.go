package awsrefactorspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteProps := &CfnRouteProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	RouteType: jsii.String("routeType"),
//   	ServiceIdentifier: jsii.String("serviceIdentifier"),
//
//   	// the properties below are optional
//   	DefaultRoute: &DefaultRouteInputProperty{
//   		ActivationState: jsii.String("activationState"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UriPathRoute: &UriPathRouteInputProperty{
//   		ActivationState: jsii.String("activationState"),
//
//   		// the properties below are optional
//   		IncludeChildPaths: jsii.Boolean(false),
//   		Methods: []*string{
//   			jsii.String("methods"),
//   		},
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   }
//
type CfnRouteProps struct {
	// The unique identifier of the application.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The unique identifier of the environment.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The route type of the route.
	RouteType *string `field:"required" json:"routeType" yaml:"routeType"`
	// The unique identifier of the service.
	ServiceIdentifier *string `field:"required" json:"serviceIdentifier" yaml:"serviceIdentifier"`
	// Configuration for the default route type.
	DefaultRoute interface{} `field:"optional" json:"defaultRoute" yaml:"defaultRoute"`
	// The tags assigned to the route.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for the URI path route type.
	UriPathRoute interface{} `field:"optional" json:"uriPathRoute" yaml:"uriPathRoute"`
}
