//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

func (a *jsiiProxy_ApplicationScalingAction) validateBindParameters(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _alarm == nil {
		return fmt.Errorf("parameter _alarm is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationScalingActionParameters(stepScalingAction awsapplicationautoscaling.StepScalingAction) error {
	if stepScalingAction == nil {
		return fmt.Errorf("parameter stepScalingAction is required, but nil was provided")
	}

	return nil
}

