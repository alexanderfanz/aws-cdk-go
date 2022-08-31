//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"fmt"
)

func (c *jsiiProxy_CfnDynamicReference) validateNewErrorParameters(message *string) error {
	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnDynamicReference) validateResolveParameters(_context IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func validateNewCfnDynamicReferenceParameters(service CfnDynamicReferenceService, key *string) error {
	if service == "" {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}
