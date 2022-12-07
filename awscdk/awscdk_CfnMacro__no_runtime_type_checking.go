//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMacro) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnMacro) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnMacro_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnMacro_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMacro_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnMacro) validateSetFunctionNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CfnMacro) validateSetNameParameters(val *string) error {
	return nil
}

func validateNewCfnMacroParameters(scope Construct, id *string, props *CfnMacroProps) error {
	return nil
}

