//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBucketNotificationDestination) validateBindParameters(scope awscdk.Construct, bucket IBucket) error {
	return nil
}

