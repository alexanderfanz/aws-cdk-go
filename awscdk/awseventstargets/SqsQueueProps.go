package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the SQS Queue Event Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var ruleTargetInput ruleTargetInput
//
//   sqsQueueProps := &SqsQueueProps{
//   	DeadLetterQueue: queue,
//   	MaxEventAge: cdk.Duration_Minutes(jsii.Number(30)),
//   	Message: ruleTargetInput,
//   	MessageGroupId: jsii.String("messageGroupId"),
//   	RetryAttempts: jsii.Number(123),
//   }
//
type SqsQueueProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The message to send to the queue.
	//
	// Must be a valid JSON text passed to the target queue.
	// Default: the entire EventBridge event.
	//
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
	// Message Group ID for messages sent to this queue.
	//
	// Required for FIFO queues, leave empty for regular queues.
	// Default: - no message group ID (regular queue).
	//
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

