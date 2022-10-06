//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BundlingDockerImage) validateCpParameters(imagePath *string) error {
	return nil
}

func (b *jsiiProxy_BundlingDockerImage) validateRunParameters(options *DockerRunOptions) error {
	return nil
}

func validateBundlingDockerImage_FromAssetParameters(path *string, options *DockerBuildOptions) error {
	return nil
}

func validateBundlingDockerImage_FromRegistryParameters(image *string) error {
	return nil
}

func validateNewBundlingDockerImageParameters(image *string) error {
	return nil
}

