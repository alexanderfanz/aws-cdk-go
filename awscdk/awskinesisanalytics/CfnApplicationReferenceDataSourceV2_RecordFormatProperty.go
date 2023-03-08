package awskinesisanalytics


// For a SQL-based Kinesis Data Analytics application, describes the record format and relevant mapping information that should be applied to schematize the records on the stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordFormatProperty := &RecordFormatProperty{
//   	RecordFormatType: jsii.String("recordFormatType"),
//
//   	// the properties below are optional
//   	MappingParameters: &MappingParametersProperty{
//   		CsvMappingParameters: &CSVMappingParametersProperty{
//   			RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   			RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   		},
//   		JsonMappingParameters: &JSONMappingParametersProperty{
//   			RecordRowPath: jsii.String("recordRowPath"),
//   		},
//   	},
//   }
//
type CfnApplicationReferenceDataSourceV2_RecordFormatProperty struct {
	// The type of record format.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
	// When you configure application input at the time of creating or updating an application, provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	MappingParameters interface{} `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

