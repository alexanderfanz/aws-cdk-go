// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents the "actual" results to compare.
//
// Example:
//   var myCustomResource customResource
//   var stack stack
//   var app app
//
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
//   	testCases: []*stack{
//   		stack,
//   	},
//   })
//   integ.assertions.expect(jsii.String("CustomAssertion"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"foo": jsii.String("bar"),
//   }), awscdkintegtestsalpha.ActualResult.fromCustomResource(myCustomResource, jsii.String("data")))
//
// Experimental.
type ActualResult interface {
	// The actual results as a string.
	// Experimental.
	Result() *string
	// Experimental.
	SetResult(val *string)
}

// The jsii proxy struct for ActualResult
type jsiiProxy_ActualResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ActualResult) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewActualResult_Override(a ActualResult) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		nil, // no parameters
		a,
	)
}

func (j *jsiiProxy_ActualResult) SetResult(val *string) {
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

// Get the actual results from a AwsApiCall.
// Experimental.
func ActualResult_FromAwsApiCall(query IAwsApiCall, attribute *string) ActualResult {
	_init_.Initialize()

	var returns ActualResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		"fromAwsApiCall",
		[]interface{}{query, attribute},
		&returns,
	)

	return returns
}

// Get the actual results from a CustomResource.
// Experimental.
func ActualResult_FromCustomResource(customResource awscdk.CustomResource, attribute *string) ActualResult {
	_init_.Initialize()

	var returns ActualResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		"fromCustomResource",
		[]interface{}{customResource, attribute},
		&returns,
	)

	return returns
}
