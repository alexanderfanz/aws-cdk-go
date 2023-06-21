package awsquicksight


// The refresh properties of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetRefreshPropertiesProperty := &DataSetRefreshPropertiesProperty{
//   	RefreshConfiguration: &RefreshConfigurationProperty{
//   		IncrementalRefresh: &IncrementalRefreshProperty{
//   			LookbackWindow: &LookbackWindowProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Size: jsii.Number(123),
//   				SizeUnit: jsii.String("sizeUnit"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_DataSetRefreshPropertiesProperty struct {
	// The refresh configuration for a dataset.
	RefreshConfiguration interface{} `field:"optional" json:"refreshConfiguration" yaml:"refreshConfiguration"`
}

