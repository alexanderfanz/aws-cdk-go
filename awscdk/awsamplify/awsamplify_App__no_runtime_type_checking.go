//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_App) validateAddAutoBranchEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (a *jsiiProxy_App) validateAddBranchParameters(id *string, options *BranchOptions) error {
	return nil
}

func (a *jsiiProxy_App) validateAddCustomRuleParameters(rule CustomRule) error {
	return nil
}

func (a *jsiiProxy_App) validateAddDomainParameters(id *string, options *DomainOptions) error {
	return nil
}

func (a *jsiiProxy_App) validateAddEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (a *jsiiProxy_App) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_App) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_App) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_App) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_App) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApp_FromAppIdParameters(scope constructs.Construct, id *string, appId *string) error {
	return nil
}

func validateApp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApp_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewAppParameters(scope constructs.Construct, id *string, props *AppProps) error {
	return nil
}
