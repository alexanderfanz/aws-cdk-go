package awschatbot


// Properties for defining a `CfnMicrosoftTeamsChannelConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMicrosoftTeamsChannelConfigurationProps := &cfnMicrosoftTeamsChannelConfigurationProps{
//   	configurationName: jsii.String("configurationName"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	teamId: jsii.String("teamId"),
//   	teamsChannelId: jsii.String("teamsChannelId"),
//   	teamsTenantId: jsii.String("teamsTenantId"),
//
//   	// the properties below are optional
//   	guardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	loggingLevel: jsii.String("loggingLevel"),
//   	snsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	userRoleRequired: jsii.Boolean(false),
//   }
//
type CfnMicrosoftTeamsChannelConfigurationProps struct {
	// The name of the configuration.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot .
	//
	// This is a user-defined role that AWS Chatbot will assume. This is not the service-linked role. For more information, see [IAM Policies for AWS Chatbot](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html) .
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The ID of the Microsoft Team authorized with AWS Chatbot .
	//
	// To get the team ID, you must perform the initial authorization flow with Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team ID from the console. For more details, see steps 1-4 in [Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the *AWS Chatbot Administrator Guide* .
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The ID of the Microsoft Teams channel.
	//
	// To get the channel ID, open Microsoft Teams, right click on the channel name in the left pane, then choose Copy. An example of the channel ID syntax is: `19%3ab6ef35dc342d56ba5654e6fc6d25a071%40thread.tacv2` .
	TeamsChannelId *string `field:"required" json:"teamsChannelId" yaml:"teamsChannelId"`
	// The ID of the Microsoft Teams tenant.
	//
	// To get the tenant ID, you must perform the initial authorization flow with Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the tenant ID from the console. For more details, see steps 1-4 in [Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the *AWS Chatbot Administrator Guide* .
	TeamsTenantId *string `field:"required" json:"teamsTenantId" yaml:"teamsTenantId"`
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies *[]*string `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Logging levels include `ERROR` , `INFO` , or `NONE` .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot .
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Enables use of a user role requirement in your chat configuration.
	UserRoleRequired interface{} `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

