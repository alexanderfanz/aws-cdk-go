package awspinpoint


// Specifies the configuration of an in-app message, including its header, body, buttons, colors, and images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inAppMessageContentProperty := &InAppMessageContentProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	BodyConfig: &BodyConfigProperty{
//   		Alignment: jsii.String("alignment"),
//   		Body: jsii.String("body"),
//   		TextColor: jsii.String("textColor"),
//   	},
//   	HeaderConfig: &HeaderConfigProperty{
//   		Alignment: jsii.String("alignment"),
//   		Header: jsii.String("header"),
//   		TextColor: jsii.String("textColor"),
//   	},
//   	ImageUrl: jsii.String("imageUrl"),
//   	PrimaryBtn: &ButtonConfigProperty{
//   		Android: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		DefaultConfig: &DefaultButtonConfigurationProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BorderRadius: jsii.Number(123),
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   			Text: jsii.String("text"),
//   			TextColor: jsii.String("textColor"),
//   		},
//   		Ios: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		Web: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   	},
//   	SecondaryBtn: &ButtonConfigProperty{
//   		Android: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		DefaultConfig: &DefaultButtonConfigurationProperty{
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BorderRadius: jsii.Number(123),
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   			Text: jsii.String("text"),
//   			TextColor: jsii.String("textColor"),
//   		},
//   		Ios: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   		Web: &OverrideButtonConfigurationProperty{
//   			ButtonAction: jsii.String("buttonAction"),
//   			Link: jsii.String("link"),
//   		},
//   	},
//   }
//
type CfnInAppTemplate_InAppMessageContentProperty struct {
	// The background color for an in-app message banner, expressed as a hex color code (such as #000000 for black).
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// An object that contains configuration information about the header or title text of the in-app message.
	BodyConfig interface{} `field:"optional" json:"bodyConfig" yaml:"bodyConfig"`
	// An object that contains configuration information about the header or title text of the in-app message.
	HeaderConfig interface{} `field:"optional" json:"headerConfig" yaml:"headerConfig"`
	// The URL of the image that appears on an in-app message banner.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// An object that contains configuration information about the primary button in an in-app message.
	PrimaryBtn interface{} `field:"optional" json:"primaryBtn" yaml:"primaryBtn"`
	// An object that contains configuration information about the secondary button in an in-app message.
	SecondaryBtn interface{} `field:"optional" json:"secondaryBtn" yaml:"secondaryBtn"`
}
