//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudwatchactions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationScalingAction) validateBindParameters(_scope awscdk.Construct, _alarm awscloudwatch.IAlarm) error {
	return nil
}

func validateNewApplicationScalingActionParameters(stepScalingAction awsapplicationautoscaling.StepScalingAction) error {
	return nil
}
