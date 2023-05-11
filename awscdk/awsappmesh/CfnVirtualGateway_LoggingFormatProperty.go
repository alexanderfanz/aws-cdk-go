package awsappmesh


// An object that represents the format for the logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingFormatProperty := &LoggingFormatProperty{
//   	Json: []interface{}{
//   		&JsonFormatRefProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Text: jsii.String("text"),
//   }
//
type CfnVirtualGateway_LoggingFormatProperty struct {
	// The logging format for JSON.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// The logging format for text.
	Text *string `field:"optional" json:"text" yaml:"text"`
}
