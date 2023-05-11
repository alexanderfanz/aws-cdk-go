package awsquicksight


// A physical table type built from the results of the custom SQL query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customSqlProperty := &CustomSqlProperty{
//   	Columns: []interface{}{
//   		&InputColumnProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	Name: jsii.String("name"),
//   	SqlQuery: jsii.String("sqlQuery"),
//   }
//
type CfnDataSet_CustomSqlProperty struct {
	// The column schema from the SQL query result set.
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// The Amazon Resource Name (ARN) of the data source.
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// A display name for the SQL query result.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The SQL query.
	SqlQuery *string `field:"required" json:"sqlQuery" yaml:"sqlQuery"`
}
