package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrafficMirrorFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorFilterProps := &CfnTrafficMirrorFilterProps{
//   	Description: jsii.String("description"),
//   	NetworkServices: []*string{
//   		jsii.String("networkServices"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTrafficMirrorFilterProps struct {
	// The description of the Traffic Mirror filter.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The network service traffic that is associated with the Traffic Mirror filter.
	//
	// Valid values are `amazon-dns` .
	NetworkServices *[]*string `field:"optional" json:"networkServices" yaml:"networkServices"`
	// The tags to assign to a Traffic Mirror filter.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

