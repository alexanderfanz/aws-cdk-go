//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (j *jsiiProxy_JournaldLogDriver) validateBindParameters(_scope awscdk.Construct, _containerDefinition ContainerDefinition) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _containerDefinition == nil {
		return fmt.Errorf("parameter _containerDefinition is required, but nil was provided")
	}

	return nil
}

func validateJournaldLogDriver_AwsLogsParameters(props *AwsLogDriverProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewJournaldLogDriverParameters(props *JournaldLogDriverProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
