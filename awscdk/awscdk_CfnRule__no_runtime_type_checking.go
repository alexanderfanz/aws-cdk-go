//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRule) validateAddAssertionParameters(condition ICfnConditionExpression, description *string) error {
	return nil
}

func (c *jsiiProxy_CfnRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnRule) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnRule) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnRule_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnRuleParameters(scope constructs.Construct, id *string, props *CfnRuleProps) error {
	return nil
}
