//go:build no_runtime_type_checking

package awsiot

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAction) validateBindParameters(topicRule ITopicRule) error {
	return nil
}

