//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateOwner_AccountRootParameters(accountId *string) error {
	return nil
}

func validateOwner_UserParameters(user awsiam.IUser) error {
	return nil
}
