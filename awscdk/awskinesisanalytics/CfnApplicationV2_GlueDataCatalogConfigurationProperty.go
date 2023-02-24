package awskinesisanalytics


// The configuration of the Glue Data Catalog that you use for Apache Flink SQL queries and table API transforms that you write in an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   glueDataCatalogConfigurationProperty := &GlueDataCatalogConfigurationProperty{
//   	DatabaseArn: jsii.String("databaseArn"),
//   }
//
type CfnApplicationV2_GlueDataCatalogConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the database.
	DatabaseArn *string `field:"optional" json:"databaseArn" yaml:"databaseArn"`
}

