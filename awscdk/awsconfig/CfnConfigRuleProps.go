package awsconfig


// Properties for defining a `CfnConfigRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inputParameters interface{}
//
//   cfnConfigRuleProps := &CfnConfigRuleProps{
//   	Source: &SourceProperty{
//   		Owner: jsii.String("owner"),
//
//   		// the properties below are optional
//   		CustomPolicyDetails: &CustomPolicyDetailsProperty{
//   			EnableDebugLogDelivery: jsii.Boolean(false),
//   			PolicyRuntime: jsii.String("policyRuntime"),
//   			PolicyText: jsii.String("policyText"),
//   		},
//   		SourceDetails: []interface{}{
//   			&SourceDetailProperty{
//   				EventSource: jsii.String("eventSource"),
//   				MessageType: jsii.String("messageType"),
//
//   				// the properties below are optional
//   				MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   			},
//   		},
//   		SourceIdentifier: jsii.String("sourceIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	ConfigRuleName: jsii.String("configRuleName"),
//   	Description: jsii.String("description"),
//   	InputParameters: inputParameters,
//   	MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   	Scope: &ScopeProperty{
//   		ComplianceResourceId: jsii.String("complianceResourceId"),
//   		ComplianceResourceTypes: []*string{
//   			jsii.String("complianceResourceTypes"),
//   		},
//   		TagKey: jsii.String("tagKey"),
//   		TagValue: jsii.String("tagValue"),
//   	},
//   }
//
type CfnConfigRuleProps struct {
	// Provides the rule owner ( `AWS` for managed rules, `CUSTOM_POLICY` for Custom Policy rules, and `CUSTOM_LAMBDA` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// A name for the AWS Config rule.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// The description that you provide for the AWS Config rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to the AWS Config rule Lambda function.
	InputParameters interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// You can specify a value for `MaximumExecutionFrequency` when:
	//
	// - You are using an AWS managed rule that is triggered at a periodic frequency.
	// - Your custom rule is triggered when AWS Config delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html) .
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources can trigger an evaluation for the rule.
	//
	// The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.
	//
	// > The scope can be empty.
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
}
