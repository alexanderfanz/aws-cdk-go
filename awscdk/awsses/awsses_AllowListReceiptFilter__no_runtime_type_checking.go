//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AllowListReceiptFilter) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AllowListReceiptFilter) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAllowListReceiptFilter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAllowListReceiptFilterParameters(scope constructs.Construct, id *string, props *AllowListReceiptFilterProps) error {
	return nil
}

