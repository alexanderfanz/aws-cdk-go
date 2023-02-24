package awsdatabrew


// Selector of a column from a dataset for profile job configuration.
//
// One selector includes either a column name or a regular expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSelectorProperty := &ColumnSelectorProperty{
//   	Name: jsii.String("name"),
//   	Regex: jsii.String("regex"),
//   }
//
type CfnJob_ColumnSelectorProperty struct {
	// The name of a column from a dataset.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A regular expression for selecting a column from a dataset.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

