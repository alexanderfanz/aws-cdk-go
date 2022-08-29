// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for creating a custom resource that will perform an API call using the AWS SDK.
// Experimental.
type IAwsApiCall interface {
	constructs.IConstruct
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall at the given path.
	//
	// For example the SQS.receiveMessage api response would look
	// like:
	//
	// If you wanted to assert the value of `Body` you could do.
	//
	// Example:
	//   var integ integTest
	//   actual := map[string][]map[string]interface{}{
	//   	"Messages": []map[string]interface{}{
	//   		map[string]interface{}{
	//   			"MessageId": jsii.String(""),
	//   			"ReceiptHandle": jsii.String(""),
	//   			"MD5OfBody": jsii.String(""),
	//   			"Body": jsii.String("hello"),
	//   			"Attributes": map[string]interface{}{
	//   			},
	//   			"MD5OfMessageAttributes": map[string]interface{}{
	//   			},
	//   			"MessageAttributes": map[string]interface{}{
	//   			},
	//   		},
	//   	},
	//   }
	//   message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"))
	//
	//   message.assertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult.stringLikeRegexp(jsii.String("hello")))
	//
	// Experimental.
	AssertAtPath(path *string, expected ExpectedResult)
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall.
	//
	// Example:
	//   var integ integTest
	//
	//   invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
	//   	functionName: jsii.String("my-func"),
	//   })
	//   invoke.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Payload": jsii.String("OK"),
	//   }))
	//
	// Experimental.
	Expect(expected ExpectedResult)
	// Returns the value of an attribute of the custom resource of an arbitrary type.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt`. Use `Token.asXxx` to encode the returned `Reference` as a specific type or
	// use the convenience `getAttString` for string attributes.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Returns the value of an attribute of the custom resource of type string.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt` encoded as a string.
	// Experimental.
	GetAttString(attributeName *string) *string
	// access the AssertionsProvider.
	//
	// This can be used to add additional IAM policies
	// the the provider role policy.
	//
	// Example:
	//   var apiCall awsApiCall
	//
	//   apiCall.provider.addToRolePolicy(map[string]interface{}{
	//   	"Effect": jsii.String("Allow"),
	//   	"Action": []*string{
	//   		jsii.String("s3:GetObject"),
	//   	},
	//   	"Resource": []*string{
	//   		jsii.String("*"),
	//   	},
	//   })
	//
	// Experimental.
	Provider() AssertionsProvider
}

// The jsii proxy for IAwsApiCall
type jsiiProxy_IAwsApiCall struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IAwsApiCall) AssertAtPath(path *string, expected ExpectedResult) {
	_jsii_.InvokeVoid(
		i,
		"assertAtPath",
		[]interface{}{path, expected},
	)
}

func (i *jsiiProxy_IAwsApiCall) Expect(expected ExpectedResult) {
	_jsii_.InvokeVoid(
		i,
		"expect",
		[]interface{}{expected},
	)
}

func (i *jsiiProxy_IAwsApiCall) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		i,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAwsApiCall) GetAttString(attributeName *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAwsApiCall) Provider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}
