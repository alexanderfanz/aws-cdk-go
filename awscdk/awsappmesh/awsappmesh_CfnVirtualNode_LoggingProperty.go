package awsappmesh


// An object that represents the logging information for a virtual node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProperty := &loggingProperty{
//   	accessLog: &accessLogProperty{
//   		file: &fileAccessLogProperty{
//   			path: jsii.String("path"),
//   		},
//   	},
//   }
//
type CfnVirtualNode_LoggingProperty struct {
	// The access log configuration for a virtual node.
	AccessLog interface{} `field:"optional" json:"accessLog" yaml:"accessLog"`
}
