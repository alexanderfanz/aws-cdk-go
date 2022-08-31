//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapplicationautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepScalingAction) validateAddAdjustmentParameters(adjustment *AdjustmentTier) error {
	return nil
}

func (s *jsiiProxy_StepScalingAction) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StepScalingAction) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateStepScalingAction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStepScalingActionParameters(scope constructs.Construct, id *string, props *StepScalingActionProps) error {
	return nil
}
