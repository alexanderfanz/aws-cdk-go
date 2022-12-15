//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) validateBoundParameters(scope awscdk.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeployStackInstancesAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCloudFormationDeployStackInstancesActionParameters(props *CloudFormationDeployStackInstancesActionProps) error {
	return nil
}

