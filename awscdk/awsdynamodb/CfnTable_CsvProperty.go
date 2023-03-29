package awsdynamodb


// The options for imported source files in CSV format.
//
// The values are Delimiter and HeaderList.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvProperty := &CsvProperty{
//   	Delimiter: jsii.String("delimiter"),
//   	HeaderList: []*string{
//   		jsii.String("headerList"),
//   	},
//   }
//
type CfnTable_CsvProperty struct {
	// The delimiter used for separating items in the CSV file being imported.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// List of the headers used to specify a common header for all source CSV files being imported.
	//
	// If this field is specified then the first line of each CSV file is treated as data instead of the header. If this field is not specified the the first line of each CSV file is treated as the header.
	HeaderList *[]*string `field:"optional" json:"headerList" yaml:"headerList"`
}

