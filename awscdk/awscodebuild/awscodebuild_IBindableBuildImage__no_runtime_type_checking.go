//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBindableBuildImage) validateBindParameters(scope awscdk.Construct, project IProject, options *BuildImageBindOptions) error {
	return nil
}

