package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContactFlow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactFlowProps := &CfnContactFlowProps{
//   	Content: jsii.String("content"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	State: jsii.String("state"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnContactFlowProps struct {
	// The content of the flow.
	//
	// For more information, see [Amazon Connect Flow language](https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html) in the *Amazon Connect Administrator Guide* .
	Content *string `field:"required" json:"content" yaml:"content"`
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the flow.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the flow.
	//
	// For descriptions of the available types, see [Choose a flow type](https://docs.aws.amazon.com/connect/latest/adminguide/create-contact-flow.html#contact-flow-types) in the *Amazon Connect Administrator Guide* .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the flow.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state of the flow.
	State *string `field:"optional" json:"state" yaml:"state"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
