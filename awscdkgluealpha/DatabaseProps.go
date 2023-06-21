package awscdkgluealpha


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   databaseProps := &DatabaseProps{
//   	DatabaseName: jsii.String("databaseName"),
//   	LocationUri: jsii.String("locationUri"),
//   }
//
// Experimental.
type DatabaseProps struct {
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The location of the database (for example, an HDFS path).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html
	//
	// Experimental.
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
}

