//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Distribution) validateAddBehaviorParameters(pathPattern *string, origin IOrigin, behaviorOptions *AddBehaviorOptions) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_Distribution) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDistribution_FromDistributionAttributesParameters(scope constructs.Construct, id *string, attrs *DistributionAttributes) error {
	return nil
}

func validateDistribution_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDistribution_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDistributionParameters(scope constructs.Construct, id *string, props *DistributionProps) error {
	return nil
}
