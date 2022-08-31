//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RateLimitedApiKey) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateGrantReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_RateLimitedApiKey) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateRateLimitedApiKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRateLimitedApiKey_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewRateLimitedApiKeyParameters(scope constructs.Construct, id *string, props *RateLimitedApiKeyProps) error {
	return nil
}
