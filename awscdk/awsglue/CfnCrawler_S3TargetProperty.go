package awsglue


// Specifies a data store in Amazon Simple Storage Service (Amazon S3).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TargetProperty := &S3TargetProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   	EventQueueArn: jsii.String("eventQueueArn"),
//   	Exclusions: []*string{
//   		jsii.String("exclusions"),
//   	},
//   	Path: jsii.String("path"),
//   	SampleSize: jsii.Number(123),
//   }
//
type CfnCrawler_S3TargetProperty struct {
	// The name of a connection which allows a job or crawler to access data in Amazon S3 within an Amazon Virtual Private Cloud environment (Amazon VPC).
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A valid Amazon dead-letter SQS ARN.
	//
	// For example, `arn:aws:sqs:region:account:deadLetterQueue` .
	DlqEventQueueArn *string `field:"optional" json:"dlqEventQueueArn" yaml:"dlqEventQueueArn"`
	// A valid Amazon SQS ARN.
	//
	// For example, `arn:aws:sqs:region:account:sqs` .
	EventQueueArn *string `field:"optional" json:"eventQueueArn" yaml:"eventQueueArn"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see [Catalog Tables with a Crawler](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html) .
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The path to the Amazon S3 target.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Sets the number of files in each leaf folder to be crawled when crawling sample files in a dataset.
	//
	// If not set, all the files are crawled. A valid value is an integer between 1 and 249.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
}
