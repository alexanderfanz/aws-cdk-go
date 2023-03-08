//go:build !no_runtime_type_checking

package awsstepfunctionstasks

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

func (i *jsiiProxy_InvokeFunction) validateBindParameters(_task awsstepfunctions.Task) error {
	if _task == nil {
		return fmt.Errorf("parameter _task is required, but nil was provided")
	}

	return nil
}

func validateNewInvokeFunctionParameters(lambdaFunction awslambda.IFunction, props *InvokeFunctionProps) error {
	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
