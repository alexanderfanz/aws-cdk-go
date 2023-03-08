package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnACL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnACLProps := &CfnACLProps{
//   	AclName: jsii.String("aclName"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserNames: []*string{
//   		jsii.String("userNames"),
//   	},
//   }
//
type CfnACLProps struct {
	// The name of the Access Control List.
	AclName *string `field:"required" json:"aclName" yaml:"aclName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The list of users that belong to the Access Control List.
	UserNames *[]*string `field:"optional" json:"userNames" yaml:"userNames"`
}

