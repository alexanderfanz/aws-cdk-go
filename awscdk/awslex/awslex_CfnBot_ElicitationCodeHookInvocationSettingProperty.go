package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   elicitationCodeHookInvocationSettingProperty := &elicitationCodeHookInvocationSettingProperty{
//   	enableCodeHookInvocation: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	invocationLabel: jsii.String("invocationLabel"),
//   }
//
type CfnBot_ElicitationCodeHookInvocationSettingProperty struct {
	// `CfnBot.ElicitationCodeHookInvocationSettingProperty.EnableCodeHookInvocation`.
	EnableCodeHookInvocation interface{} `field:"required" json:"enableCodeHookInvocation" yaml:"enableCodeHookInvocation"`
	// `CfnBot.ElicitationCodeHookInvocationSettingProperty.InvocationLabel`.
	InvocationLabel *string `field:"optional" json:"invocationLabel" yaml:"invocationLabel"`
}
