//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRefElement) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnRefElement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnRefElement) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnRefElement_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnRefElement_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnRefElementParameters(scope constructs.Construct, id *string) error {
	return nil
}

