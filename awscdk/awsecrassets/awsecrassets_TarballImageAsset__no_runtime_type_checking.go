//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecrassets

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TarballImageAsset) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_TarballImageAsset) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateTarballImageAsset_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TarballImageAsset) validateSetImageUriParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TarballImageAsset) validateSetRepositoryParameters(val awsecr.IRepository) error {
	return nil
}

func validateNewTarballImageAssetParameters(scope constructs.Construct, id *string, props *TarballImageAssetProps) error {
	return nil
}
