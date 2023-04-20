package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the Batch Job Event Target.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), &JobQueueProps{
//   	ComputeEnvironments: []orderedComputeEnvironment{
//   		&orderedComputeEnvironment{
//   			ComputeEnvironment: batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
//   				"managed": jsii.Boolean(false),
//   			}),
//   			Order: jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
//   	"container": map[string]repositoryImage{
//   		"image": awscdk.ContainerImage_fromRegistry(jsii.String("test-repo")),
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewBatchJob(jobQueue.JobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &BatchJobProps{
//   	DeadLetterQueue: queue,
//   	Event: events.RuleTargetInput_FromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	RetryAttempts: jsii.Number(2),
//   	MaxEventAge: awscdk.Duration_*Hours(jsii.Number(2)),
//   }))
//
type BatchJobProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The number of times to attempt to retry, if the job fails.
	//
	// Valid values are 1–10.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The name of the submitted job.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// The size of the array, if this is an array batch job.
	//
	// Valid values are integers between 2 and 10,000.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

