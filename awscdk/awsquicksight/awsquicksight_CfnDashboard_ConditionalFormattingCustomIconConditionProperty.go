package awsquicksight


// Determines the custom condition for an icon set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionalFormattingCustomIconConditionProperty := &ConditionalFormattingCustomIconConditionProperty{
//   	Expression: jsii.String("expression"),
//   	IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   		Icon: jsii.String("icon"),
//   		UnicodeIcon: jsii.String("unicodeIcon"),
//   	},
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   	DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   		IconDisplayOption: jsii.String("iconDisplayOption"),
//   	},
//   }
//
type CfnDashboard_ConditionalFormattingCustomIconConditionProperty struct {
	// The expression that determines the condition of the icon set.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Custom icon options for an icon set.
	IconOptions interface{} `field:"required" json:"iconOptions" yaml:"iconOptions"`
	// Determines the color of the icon.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Determines the icon display configuration.
	DisplayConfiguration interface{} `field:"optional" json:"displayConfiguration" yaml:"displayConfiguration"`
}
