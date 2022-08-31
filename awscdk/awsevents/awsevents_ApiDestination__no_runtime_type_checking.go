//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiDestination) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ApiDestination) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ApiDestination) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_ApiDestination) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_ApiDestination) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApiDestination_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApiDestination_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewApiDestinationParameters(scope constructs.Construct, id *string, props *ApiDestinationProps) error {
	return nil
}
