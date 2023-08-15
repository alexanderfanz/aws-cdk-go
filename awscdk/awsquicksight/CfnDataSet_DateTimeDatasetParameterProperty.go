package awsquicksight


// <p>A parameter created in the dataset of date time data type.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeDatasetParameterProperty := &DateTimeDatasetParameterProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ValueType: jsii.String("valueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &DateTimeDatasetParameterDefaultValuesProperty{
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html
//
type CfnDataSet_DateTimeDatasetParameterProperty struct {
	// <p>Identifier of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// <p>Name of the parameter created in the dataset.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>Every parameter value could be either a single value or multi value which helps to validate before evaluation.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-valuetype
	//
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// <p>List of default values defined for a given string date time parameter type.
	//
	// Currently only static values are supported.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}
