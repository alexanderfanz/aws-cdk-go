package awsgreengrassv2


// Contains information about the timeout configuration for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobTimeoutConfigProperty := &ioTJobTimeoutConfigProperty{
//   	inProgressTimeoutInMinutes: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobTimeoutConfigProperty struct {
	// The amount of time, in minutes, that devices have to complete the job.
	//
	// The timer starts when the job status is set to `IN_PROGRESS` . If the job status doesn't change to a terminal state before the time expires, then the job status is set to `TIMED_OUT` .
	//
	// The timeout interval must be between 1 minute and 7 days (10080 minutes).
	InProgressTimeoutInMinutes *float64 `field:"optional" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

