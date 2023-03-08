package awssagemaker


// The meta data of the Glue table which serves as data catalog for the `OfflineStore` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCatalogConfigProperty := &DataCatalogConfigProperty{
//   	Catalog: jsii.String("catalog"),
//   	Database: jsii.String("database"),
//   	TableName: jsii.String("tableName"),
//   }
//
type CfnFeatureGroup_DataCatalogConfigProperty struct {
	// The name of the Glue table catalog.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// The name of the Glue table database.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The name of the Glue table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

