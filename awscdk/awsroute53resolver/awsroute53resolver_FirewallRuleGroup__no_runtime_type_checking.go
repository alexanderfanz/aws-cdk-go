//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsroute53resolver

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallRuleGroup) validateAddRuleParameters(rule *FirewallRule) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroup) validateAssociateParameters(id *string, props *FirewallRuleGroupAssociationOptions) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (f *jsiiProxy_FirewallRuleGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateFirewallRuleGroup_FromFirewallRuleGroupIdParameters(scope constructs.Construct, id *string, firewallRuleGroupId *string) error {
	return nil
}

func validateFirewallRuleGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFirewallRuleGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewFirewallRuleGroupParameters(scope constructs.Construct, id *string, props *FirewallRuleGroupProps) error {
	return nil
}
