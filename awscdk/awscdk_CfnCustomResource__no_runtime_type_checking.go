//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomResource) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnCustomResource_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnCustomResource_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnCustomResource) validateSetServiceTokenParameters(val *string) error {
	return nil
}

func validateNewCfnCustomResourceParameters(scope Construct, id *string, props *CfnCustomResourceProps) error {
	return nil
}

