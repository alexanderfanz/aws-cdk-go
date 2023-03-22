package awsbackup


// Includes information about tags you define to assign tagged resources to a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionParameterProperty := &conditionParameterProperty{
//   	conditionKey: jsii.String("conditionKey"),
//   	conditionValue: jsii.String("conditionValue"),
//   }
//
type CfnBackupSelection_ConditionParameterProperty struct {
	// The key in a key-value pair.
	//
	// For example, in the tag `Department: Accounting` , `Department` is the key.
	ConditionKey *string `field:"optional" json:"conditionKey" yaml:"conditionKey"`
	// The value in a key-value pair.
	//
	// For example, in the tag `Department: Accounting` , `Accounting` is the value.
	ConditionValue *string `field:"optional" json:"conditionValue" yaml:"conditionValue"`
}

