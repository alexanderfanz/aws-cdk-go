//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunBatchJob) validateBindParameters(_task awsstepfunctions.Task) error {
	return nil
}

func validateNewRunBatchJobParameters(props *RunBatchJobProps) error {
	return nil
}

