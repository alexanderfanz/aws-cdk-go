//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	return nil
}

func validateNewCliCredentialsStackSynthesizerParameters(props *CliCredentialsStackSynthesizerProps) error {
	return nil
}

