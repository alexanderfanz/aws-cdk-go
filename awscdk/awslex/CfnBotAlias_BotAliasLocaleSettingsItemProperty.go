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
//   botAliasLocaleSettingsItemProperty := &BotAliasLocaleSettingsItemProperty{
//   	BotAliasLocaleSetting: &BotAliasLocaleSettingsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		CodeHookSpecification: &CodeHookSpecificationProperty{
//   			LambdaCodeHook: &LambdaCodeHookProperty{
//   				CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   				LambdaArn: jsii.String("lambdaArn"),
//   			},
//   		},
//   	},
//   	LocaleId: jsii.String("localeId"),
//   }
//
type CfnBotAlias_BotAliasLocaleSettingsItemProperty struct {
	// Specifies settings that are unique to a locale.
	BotAliasLocaleSetting interface{} `field:"required" json:"botAliasLocaleSetting" yaml:"botAliasLocaleSetting"`
	// The unique identifier of the locale.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}
