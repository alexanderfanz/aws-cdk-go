package awsquicksight


// A transform operation that tags a column with additional information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagColumnOperationProperty := &TagColumnOperationProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Tags: []columnTagProperty{
//   		&columnTagProperty{
//   			ColumnDescription: &ColumnDescriptionProperty{
//   				Text: jsii.String("text"),
//   			},
//   			ColumnGeographicRole: jsii.String("columnGeographicRole"),
//   		},
//   	},
//   }
//
type CfnDataSet_TagColumnOperationProperty struct {
	// The column that this operation acts on.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The dataset column tag, currently only used for geospatial type tagging.
	//
	// > This is not tags for the AWS tagging feature.
	Tags *[]*CfnDataSet_ColumnTagProperty `field:"required" json:"tags" yaml:"tags"`
}
