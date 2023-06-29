package awsec2


// Specifies the hibernation options for the instance.
//
// `HibernationOptions` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hibernationOptionsProperty := &HibernationOptionsProperty{
//   	Configured: jsii.Boolean(false),
//   }
//
type CfnInstance_HibernationOptionsProperty struct {
	// Set to `true` to enable your instance for hibernation.
	//
	// Default: `false`.
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

