package awsscheduler


// A `RetryPolicy` object that includes information about the retry policy settings, including the maximum age of an event, and the maximum number of times EventBridge Scheduler will try to deliver the event to a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retryPolicyProperty := &retryPolicyProperty{
//   	maximumEventAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   }
//
type CfnSchedule_RetryPolicyProperty struct {
	// The maximum amount of time, in seconds, to continue to make retry attempts.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// The maximum number of retry attempts to make before the request fails.
	//
	// Retry attempts with exponential backoff continue until either the maximum number of attempts is made or until the duration of the `MaximumEventAgeInSeconds` is reached.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

