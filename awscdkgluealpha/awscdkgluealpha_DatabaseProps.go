// The CDK Construct Library for AWS::Glue
package awscdkgluealpha


// Example:
//   glue.NewDatabase(this, jsii.String("MyDatabase"), &databaseProps{
//   	databaseName: jsii.String("my_database"),
//   })
//
// Experimental.
type DatabaseProps struct {
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The location of the database (for example, an HDFS path).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html
	//
	// Experimental.
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
}

