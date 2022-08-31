//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDynamicReference) validateNewErrorParameters(message *string) error {
	return nil
}

func (c *jsiiProxy_CfnDynamicReference) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func validateNewCfnDynamicReferenceParameters(service CfnDynamicReferenceService, key *string) error {
	return nil
}
