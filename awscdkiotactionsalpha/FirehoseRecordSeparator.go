// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha


// Record Separator to be used to separate records.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &FirehosePutRecordActionProps{
//   			BatchMode: jsii.Boolean(true),
//   			RecordSeparator: actions.FirehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type FirehoseRecordSeparator string

const (
	// Separate by a new line.
	// Experimental.
	FirehoseRecordSeparator_NEWLINE FirehoseRecordSeparator = "NEWLINE"
	// Separate by a tab.
	// Experimental.
	FirehoseRecordSeparator_TAB FirehoseRecordSeparator = "TAB"
	// Separate by a windows new line.
	// Experimental.
	FirehoseRecordSeparator_WINDOWS_NEWLINE FirehoseRecordSeparator = "WINDOWS_NEWLINE"
	// Separate by a commma.
	// Experimental.
	FirehoseRecordSeparator_COMMA FirehoseRecordSeparator = "COMMA"
)

