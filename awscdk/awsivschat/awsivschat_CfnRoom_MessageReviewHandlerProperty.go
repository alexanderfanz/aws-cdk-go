package awsivschat


// The MessageReviewHandler property type specifies configuration information for optional message review.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageReviewHandlerProperty := &messageReviewHandlerProperty{
//   	fallbackResult: jsii.String("fallbackResult"),
//   	uri: jsii.String("uri"),
//   }
//
type CfnRoom_MessageReviewHandlerProperty struct {
	// Specifies the fallback behavior (whether the message is allowed or denied) if the handler does not return a valid response, encounters an error, or times out.
	//
	// (For the timeout period, see [Service Quotas](https://docs.aws.amazon.com/ivs/latest/userguide/service-quotas.html) .) If allowed, the message is delivered with returned content to all users connected to the room. If denied, the message is not delivered to any user.
	//
	// *Default* : `ALLOW`.
	FallbackResult *string `field:"optional" json:"fallbackResult" yaml:"fallbackResult"`
	// Identifier of the message review handler.
	//
	// Currently this must be an ARN of a lambda function.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

