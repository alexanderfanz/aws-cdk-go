package awsiot


// A key-value pair that you define in the header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPropertyProperty := &userPropertyProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnTopicRule_UserPropertyProperty struct {
	// A key to be specified in `UserProperty` .
	Key *string `field:"required" json:"key" yaml:"key"`
	// A value to be specified in `UserProperty` .
	Value *string `field:"required" json:"value" yaml:"value"`
}
