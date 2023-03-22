package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   datasetFormatProperty := &datasetFormatProperty{
//   	csv: &csvProperty{
//   		header: jsii.Boolean(false),
//   	},
//   	json: json,
//   	parquet: jsii.Boolean(false),
//   }
//
type CfnModelBiasJobDefinition_DatasetFormatProperty struct {
	// `CfnModelBiasJobDefinition.DatasetFormatProperty.Csv`.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// `CfnModelBiasJobDefinition.DatasetFormatProperty.Json`.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// `CfnModelBiasJobDefinition.DatasetFormatProperty.Parquet`.
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

