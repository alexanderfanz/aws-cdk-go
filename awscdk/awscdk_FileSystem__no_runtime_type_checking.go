//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateFileSystem_CopyDirectoryParameters(srcDir *string, destDir *string, options *CopyOptions) error {
	return nil
}

func validateFileSystem_FingerprintParameters(fileOrDirectory *string, options *FingerprintOptions) error {
	return nil
}

func validateFileSystem_IsEmptyParameters(dir *string) error {
	return nil
}

func validateFileSystem_MkdtempParameters(prefix *string) error {
	return nil
}
