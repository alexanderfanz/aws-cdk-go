//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Authorizer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Authorizer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Authorizer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Authorizer) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_Authorizer) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAuthorizer_IsAuthorizerParameters(x interface{}) error {
	return nil
}

func validateAuthorizer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAuthorizer_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewAuthorizerParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

