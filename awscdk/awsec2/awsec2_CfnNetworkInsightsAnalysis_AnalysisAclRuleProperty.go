package awsec2


// Describes a network access control (ACL) rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisAclRuleProperty := &AnalysisAclRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   	Egress: jsii.Boolean(false),
//   	PortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	Protocol: jsii.String("protocol"),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisAclRuleProperty struct {
	// The IPv4 address range, in CIDR notation.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Indicates whether the rule is an outbound rule.
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// The range of ports.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// The protocol.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether to allow or deny traffic that matches the rule.
	RuleAction *string `field:"optional" json:"ruleAction" yaml:"ruleAction"`
	// The rule number.
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
}

