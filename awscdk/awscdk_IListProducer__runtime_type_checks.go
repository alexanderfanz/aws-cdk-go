//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"
)

func (i *jsiiProxy_IListProducer) validateProduceParameters(context IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

