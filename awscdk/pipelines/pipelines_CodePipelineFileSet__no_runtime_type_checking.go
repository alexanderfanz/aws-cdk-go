//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func validateCodePipelineFileSet_FromArtifactParameters(artifact awscodepipeline.Artifact) error {
	return nil
}
