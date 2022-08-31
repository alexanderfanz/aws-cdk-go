//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogGroup) validateAddMetricFilterParameters(id *string, props *MetricFilterOptions) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateAddStreamParameters(id *string, props *StreamOptions) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateAddSubscriptionFilterParameters(id *string, props *SubscriptionFilterOptions) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateExtractMetricParameters(jsonField *string, metricNamespace *string, metricName *string) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_LogGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateLogGroup_FromLogGroupArnParameters(scope constructs.Construct, id *string, logGroupArn *string) error {
	return nil
}

func validateLogGroup_FromLogGroupNameParameters(scope constructs.Construct, id *string, logGroupName *string) error {
	return nil
}

func validateLogGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLogGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewLogGroupParameters(scope constructs.Construct, id *string, props *LogGroupProps) error {
	return nil
}
