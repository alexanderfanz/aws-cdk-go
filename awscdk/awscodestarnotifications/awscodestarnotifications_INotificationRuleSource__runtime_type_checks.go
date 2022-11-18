//go:build !no_runtime_type_checking

package awscodestarnotifications

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v3"
)

func (i *jsiiProxy_INotificationRuleSource) validateBindAsNotificationRuleSourceParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

