//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaApplication) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LambdaApplication) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LambdaApplication) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (l *jsiiProxy_LambdaApplication) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_LambdaApplication) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateLambdaApplication_FromLambdaApplicationNameParameters(scope constructs.Construct, id *string, lambdaApplicationName *string) error {
	return nil
}

func validateLambdaApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLambdaApplication_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewLambdaApplicationParameters(scope constructs.Construct, id *string, props *LambdaApplicationProps) error {
	return nil
}
