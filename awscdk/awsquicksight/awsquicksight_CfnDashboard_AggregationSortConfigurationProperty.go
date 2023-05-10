package awsquicksight


// The configuration options to sort aggregated values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationSortConfigurationProperty := &AggregationSortConfigurationProperty{
//   	AggregationFunction: &AggregationFunctionProperty{
//   		CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   		DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   		NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   			PercentileAggregation: &PercentileAggregationProperty{
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   		},
//   	},
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SortDirection: jsii.String("sortDirection"),
//   }
//
type CfnDashboard_AggregationSortConfigurationProperty struct {
	// The function that aggregates the values in `Column` .
	AggregationFunction interface{} `field:"required" json:"aggregationFunction" yaml:"aggregationFunction"`
	// The column that determines the sort order of aggregated values.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The sort direction of values.
	//
	// - `ASC` : Sort in ascending order.
	// - `DESC` : Sort in descending order.
	SortDirection *string `field:"required" json:"sortDirection" yaml:"sortDirection"`
}
