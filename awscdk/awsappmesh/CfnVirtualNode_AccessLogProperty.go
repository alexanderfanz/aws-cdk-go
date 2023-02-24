package awsappmesh


// An object that represents the access logging information for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogProperty := &AccessLogProperty{
//   	File: &FileAccessLogProperty{
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		Format: &LoggingFormatProperty{
//   			Json: []interface{}{
//   				&JsonFormatRefProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Text: jsii.String("text"),
//   		},
//   	},
//   }
//
type CfnVirtualNode_AccessLogProperty struct {
	// The file object to send virtual node access logs to.
	File interface{} `field:"optional" json:"file" yaml:"file"`
}

