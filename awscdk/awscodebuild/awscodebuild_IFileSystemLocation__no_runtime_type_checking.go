//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IFileSystemLocation) validateBindParameters(scope awscdk.Construct, project IProject) error {
	return nil
}
