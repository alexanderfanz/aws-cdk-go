package awsroute53resolver


// A single firewall rule in a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallRuleProperty := &FirewallRuleProperty{
//   	Action: jsii.String("action"),
//   	FirewallDomainListId: jsii.String("firewallDomainListId"),
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	BlockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   	BlockOverrideDomain: jsii.String("blockOverrideDomain"),
//   	BlockOverrideTtl: jsii.Number(123),
//   	BlockResponse: jsii.String("blockResponse"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html
//
type CfnFirewallRuleGroup_FirewallRuleProperty struct {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list:  - `ALLOW` - Permit the request to go through.
	//
	// - `ALERT` - Permit the request to go through but send an alert to the logs.
	// - `BLOCK` - Disallow the request. If this is specified,then `BlockResponse` must also be specified.
	//
	// if `BlockResponse` is `OVERRIDE` , then all of the following `OVERRIDE` attributes must be specified:
	//
	// - `BlockOverrideDnsType`
	// - `BlockOverrideDomain`
	// - `BlockOverrideTtl`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The ID of the domain list that's used in the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-firewalldomainlistid
	//
	FirewallDomainListId *string `field:"required" json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// The priority of the rule in the rule group.
	//
	// This value must be unique within the rule group. DNS Firewall processes the rules in a rule group by order of priority, starting from the lowest setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The DNS record's type.
	//
	// This determines the format of the record value that you provided in `BlockOverrideDomain` . Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-blockoverridednstype
	//
	BlockOverrideDnsType *string `field:"optional" json:"blockOverrideDnsType" yaml:"blockOverrideDnsType"`
	// The custom DNS record to send back in response to the query.
	//
	// Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-blockoverridedomain
	//
	BlockOverrideDomain *string `field:"optional" json:"blockOverrideDomain" yaml:"blockOverrideDomain"`
	// The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
	//
	// Used for the rule action `BLOCK` with a `BlockResponse` setting of `OVERRIDE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-blockoverridettl
	//
	BlockOverrideTtl *float64 `field:"optional" json:"blockOverrideTtl" yaml:"blockOverrideTtl"`
	// The way that you want DNS Firewall to block the request. Used for the rule action setting `BLOCK` .
	//
	// - `NODATA` - Respond indicating that the query was successful, but no response is available for it.
	// - `NXDOMAIN` - Respond indicating that the domain name that's in the query doesn't exist.
	// - `OVERRIDE` - Provide a custom override in the response. This option requires custom handling details in the rule's `BlockOverride*` settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallrule.html#cfn-route53resolver-firewallrulegroup-firewallrule-blockresponse
	//
	BlockResponse *string `field:"optional" json:"blockResponse" yaml:"blockResponse"`
}

