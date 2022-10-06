//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RemoveTag) validateApplyTagParameters(resource ITaggable) error {
	return nil
}

func (r *jsiiProxy_RemoveTag) validateVisitParameters(construct IConstruct) error {
	return nil
}

func validateNewRemoveTagParameters(key *string, props *TagProps) error {
	return nil
}

