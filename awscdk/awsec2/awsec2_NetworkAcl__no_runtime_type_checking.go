//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkAcl) validateAddEntryParameters(id *string, options *CommonNetworkAclEntryOptions) error {
	return nil
}

func (n *jsiiProxy_NetworkAcl) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NetworkAcl) validateAssociateWithSubnetParameters(id *string, selection *SubnetSelection) error {
	return nil
}

func (n *jsiiProxy_NetworkAcl) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NetworkAcl) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NetworkAcl) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NetworkAcl) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateNetworkAcl_FromNetworkAclIdParameters(scope constructs.Construct, id *string, networkAclId *string) error {
	return nil
}

func validateNetworkAcl_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkAcl_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewNetworkAclParameters(scope constructs.Construct, id *string, props *NetworkAclProps) error {
	return nil
}
