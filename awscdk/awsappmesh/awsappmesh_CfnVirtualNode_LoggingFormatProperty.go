package awsappmesh


// An object that represents the format for the logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFormatProperty := &loggingFormatProperty{
//   	json: []interface{}{
//   		&jsonFormatRefProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	text: jsii.String("text"),
//   }
//
type CfnVirtualNode_LoggingFormatProperty struct {
	// The logging format for JSON.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// The logging format for text.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

