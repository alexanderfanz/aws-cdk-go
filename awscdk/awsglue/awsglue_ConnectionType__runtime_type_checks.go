//go:build !no_runtime_type_checking

package awsglue

import (
	"fmt"
)

func validateNewConnectionTypeParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

