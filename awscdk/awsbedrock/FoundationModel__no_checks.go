//go:build no_runtime_type_checking

package awsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateFoundationModel_FromFoundationModelIdParameters(scope constructs.Construct, _id *string, foundationModelId FoundationModelIdentifier) error {
	return nil
}
