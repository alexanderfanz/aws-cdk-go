package awskinesisanalyticsv2


// The configuration parameters for the default Amazon Glue database.
//
// You use this database for SQL queries that you write in a Kinesis Data Analytics Studio notebook.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogConfigurationProperty := &CatalogConfigurationProperty{
//   	GlueDataCatalogConfiguration: &GlueDataCatalogConfigurationProperty{
//   		DatabaseArn: jsii.String("databaseArn"),
//   	},
//   }
//
type CfnApplication_CatalogConfigurationProperty struct {
	// The configuration parameters for the default Amazon Glue database.
	//
	// You use this database for Apache Flink SQL queries and table API transforms that you write in a Kinesis Data Analytics Studio notebook.
	GlueDataCatalogConfiguration interface{} `field:"optional" json:"glueDataCatalogConfiguration" yaml:"glueDataCatalogConfiguration"`
}
