//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) validateBoundParameters(_scope awscdk.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (s *jsiiProxy_ServiceCatalogDeployActionBeta1) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewServiceCatalogDeployActionBeta1Parameters(props *ServiceCatalogDeployActionBeta1Props) error {
	return nil
}

