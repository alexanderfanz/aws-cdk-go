package awslex


// Specifies settings that are unique to a locale.
//
// For example, you can use different Lambda function depending on the bot's locale.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   botAliasLocaleSettingsProperty := &BotAliasLocaleSettingsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	CodeHookSpecification: &CodeHookSpecificationProperty{
//   		LambdaCodeHook: &LambdaCodeHookProperty{
//   			CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   			LambdaArn: jsii.String("lambdaArn"),
//   		},
//   	},
//   }
//
type CfnBot_BotAliasLocaleSettingsProperty struct {
	// Determines whether the locale is enabled for the bot.
	//
	// If the value is `false` , the locale isn't available for use.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the Lambda function that should be used in the locale.
	CodeHookSpecification interface{} `field:"optional" json:"codeHookSpecification" yaml:"codeHookSpecification"`
}

