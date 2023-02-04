package awsiot


// The criteria that determines how many retries are allowed for each failure type for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retryCriteriaProperty := &retryCriteriaProperty{
//   	failureType: jsii.String("failureType"),
//   	numberOfRetries: jsii.Number(123),
//   }
//
type CfnJobTemplate_RetryCriteriaProperty struct {
	// The type of job execution failures that can initiate a job retry.
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// The number of retries allowed for a failure type for the job.
	NumberOfRetries *float64 `field:"optional" json:"numberOfRetries" yaml:"numberOfRetries"`
}
