package awsquicksight


// The image configuration of a table field URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableFieldImageConfigurationProperty := &TableFieldImageConfigurationProperty{
//   	SizingOptions: &TableCellImageSizingConfigurationProperty{
//   		TableCellImageScalingConfiguration: jsii.String("tableCellImageScalingConfiguration"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablefieldimageconfiguration.html
//
type CfnAnalysis_TableFieldImageConfigurationProperty struct {
	// The sizing options for the table image configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablefieldimageconfiguration.html#cfn-quicksight-analysis-tablefieldimageconfiguration-sizingoptions
	//
	SizingOptions interface{} `field:"optional" json:"sizingOptions" yaml:"sizingOptions"`
}

