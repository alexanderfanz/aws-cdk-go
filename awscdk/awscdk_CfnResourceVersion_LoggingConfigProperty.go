// An experiment to bundle the entire CDK into a single module
package awscdk


// Logging configuration information for a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigProperty := &loggingConfigProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   	logRoleArn: jsii.String("logRoleArn"),
//   }
//
type CfnResourceVersion_LoggingConfigProperty struct {
	// The Amazon CloudWatch logs group to which CloudFormation sends error logging information when invoking the type's handlers.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The ARN of the role that CloudFormation should assume when sending log entries to CloudWatch logs.
	LogRoleArn *string `field:"optional" json:"logRoleArn" yaml:"logRoleArn"`
}

