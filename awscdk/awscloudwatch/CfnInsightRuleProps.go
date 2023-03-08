package awscloudwatch


// Properties for defining a `CfnInsightRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInsightRuleProps := &CfnInsightRuleProps{
//   	RuleBody: jsii.String("ruleBody"),
//   	RuleName: jsii.String("ruleName"),
//   	RuleState: jsii.String("ruleState"),
//
//   	// the properties below are optional
//   	Tags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnInsightRuleProps struct {
	// The definition of the rule, as a JSON object.
	//
	// For details about the syntax, see [Contributor Insights Rule Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights-RuleSyntax.html) in the *Amazon CloudWatch User Guide* .
	RuleBody *string `field:"required" json:"ruleBody" yaml:"ruleBody"`
	// The name of the rule.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The current state of the rule.
	//
	// Valid values are `ENABLED` and `DISABLED` .
	RuleState *string `field:"required" json:"ruleState" yaml:"ruleState"`
	// A list of key-value pairs to associate with the Contributor Insights rule.
	//
	// You can associate as many as 50 tags with a rule.
	//
	// Tags can help you organize and categorize your resources. For more information, see [Tagging Your Amazon CloudWatch Resources](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Tagging.html) .
	//
	// To be able to associate tags with a rule, you must have the `cloudwatch:TagResource` permission in addition to the `cloudwatch:PutInsightRule` permission.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

