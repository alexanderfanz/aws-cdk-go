//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReceiptRuleSet) validateAddRuleParameters(id *string, options *ReceiptRuleOptions) error {
	return nil
}

func (r *jsiiProxy_ReceiptRuleSet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ReceiptRuleSet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ReceiptRuleSet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_ReceiptRuleSet) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_ReceiptRuleSet) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateReceiptRuleSet_FromReceiptRuleSetNameParameters(scope constructs.Construct, id *string, receiptRuleSetName *string) error {
	return nil
}

func validateReceiptRuleSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateReceiptRuleSet_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewReceiptRuleSetParameters(scope constructs.Construct, id *string, props *ReceiptRuleSetProps) error {
	return nil
}
