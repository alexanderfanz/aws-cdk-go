package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNetworkInsightsAccessScopeAnalysis`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInsightsAccessScopeAnalysisProps := &CfnNetworkInsightsAccessScopeAnalysisProps{
//   	NetworkInsightsAccessScopeId: jsii.String("networkInsightsAccessScopeId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAccessScopeAnalysisProps struct {
	// The ID of the Network Access Scope.
	NetworkInsightsAccessScopeId *string `field:"required" json:"networkInsightsAccessScopeId" yaml:"networkInsightsAccessScopeId"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

