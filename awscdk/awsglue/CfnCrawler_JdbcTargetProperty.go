package awsglue


// Specifies a JDBC data store to crawl.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jdbcTargetProperty := &JdbcTargetProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	Exclusions: []*string{
//   		jsii.String("exclusions"),
//   	},
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html
//
type CfnCrawler_JdbcTargetProperty struct {
	// The name of the connection to use to connect to the JDBC target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-connectionname
	//
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-exclusions
	//
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The path of the JDBC target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-crawler-jdbctarget.html#cfn-glue-crawler-jdbctarget-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

