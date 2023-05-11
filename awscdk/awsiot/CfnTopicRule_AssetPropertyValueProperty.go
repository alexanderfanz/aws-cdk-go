package awsiot


// An asset property value entry containing the following information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyValueProperty := &AssetPropertyValueProperty{
//   	Timestamp: &AssetPropertyTimestampProperty{
//   		TimeInSeconds: jsii.String("timeInSeconds"),
//
//   		// the properties below are optional
//   		OffsetInNanos: jsii.String("offsetInNanos"),
//   	},
//   	Value: &AssetPropertyVariantProperty{
//   		BooleanValue: jsii.String("booleanValue"),
//   		DoubleValue: jsii.String("doubleValue"),
//   		IntegerValue: jsii.String("integerValue"),
//   		StringValue: jsii.String("stringValue"),
//   	},
//
//   	// the properties below are optional
//   	Quality: jsii.String("quality"),
//   }
//
type CfnTopicRule_AssetPropertyValueProperty struct {
	// The asset property value timestamp.
	Timestamp interface{} `field:"required" json:"timestamp" yaml:"timestamp"`
	// The value of the asset property.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// Optional.
	//
	// A string that describes the quality of the value. Accepts substitution templates. Must be `GOOD` , `BAD` , or `UNCERTAIN` .
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
}
