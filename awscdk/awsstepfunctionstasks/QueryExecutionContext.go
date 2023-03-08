package awsstepfunctionstasks


// Database and data catalog context in which the query execution occurs.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &AthenaStartQueryExecutionProps{
//   	QueryString: sfn.JsonPath_StringAt(jsii.String("$.queryString")),
//   	QueryExecutionContext: &QueryExecutionContext{
//   		DatabaseName: jsii.String("mydatabase"),
//   	},
//   	ResultConfiguration: &ResultConfiguration{
//   		EncryptionConfiguration: &EncryptionConfiguration{
//   			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
//   		},
//   		OutputLocation: &Location{
//   			BucketName: jsii.String("query-results-bucket"),
//   			ObjectKey: jsii.String("folder"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecutionContext.html
//
type QueryExecutionContext struct {
	// Name of catalog used in query execution.
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Name of database used in query execution.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

