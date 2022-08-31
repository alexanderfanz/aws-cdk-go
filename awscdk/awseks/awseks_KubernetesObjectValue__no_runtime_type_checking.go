//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesObjectValue) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (k *jsiiProxy_KubernetesObjectValue) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateKubernetesObjectValue_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKubernetesObjectValueParameters(scope constructs.Construct, id *string, props *KubernetesObjectValueProps) error {
	return nil
}
