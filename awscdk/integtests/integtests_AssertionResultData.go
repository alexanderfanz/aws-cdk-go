package integtests


// The result of an assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assertionResultData := &assertionResultData{
//   	status: awscdk.Integ_tests.status_PASS,
//
//   	// the properties below are optional
//   	message: jsii.String("message"),
//   }
//
// Experimental.
type AssertionResultData struct {
	// The status of the assertion, i.e. pass or fail.
	// Experimental.
	Status Status `field:"required" json:"status" yaml:"status"`
	// Any message returned with the assertion result typically this will be the diff if there is any.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}
