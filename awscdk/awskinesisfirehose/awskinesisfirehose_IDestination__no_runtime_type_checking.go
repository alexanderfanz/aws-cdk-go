//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDestination) validateBindParameters(scope constructs.Construct, options *DestinationBindOptions) error {
	return nil
}
