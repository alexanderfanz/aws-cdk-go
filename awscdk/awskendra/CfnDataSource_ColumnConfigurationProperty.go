package awskendra


// Provides information about how Amazon Kendra should use the columns of a database in an index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnConfigurationProperty := &ColumnConfigurationProperty{
//   	ChangeDetectingColumns: []*string{
//   		jsii.String("changeDetectingColumns"),
//   	},
//   	DocumentDataColumnName: jsii.String("documentDataColumnName"),
//   	DocumentIdColumnName: jsii.String("documentIdColumnName"),
//
//   	// the properties below are optional
//   	DocumentTitleColumnName: jsii.String("documentTitleColumnName"),
//   	FieldMappings: []interface{}{
//   		&DataSourceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ColumnConfigurationProperty struct {
	// One to five columns that indicate when a document in the database has changed.
	ChangeDetectingColumns *[]*string `field:"required" json:"changeDetectingColumns" yaml:"changeDetectingColumns"`
	// The column that contains the contents of the document.
	DocumentDataColumnName *string `field:"required" json:"documentDataColumnName" yaml:"documentDataColumnName"`
	// The column that provides the document's identifier.
	DocumentIdColumnName *string `field:"required" json:"documentIdColumnName" yaml:"documentIdColumnName"`
	// The column that contains the title of the document.
	DocumentTitleColumnName *string `field:"optional" json:"documentTitleColumnName" yaml:"documentTitleColumnName"`
	// An array of objects that map database column names to the corresponding fields in an index.
	//
	// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}
