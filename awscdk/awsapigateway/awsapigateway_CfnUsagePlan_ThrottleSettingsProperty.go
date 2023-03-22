package awsapigateway


// `ThrottleSettings` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throttleSettingsProperty := &throttleSettingsProperty{
//   	burstLimit: jsii.Number(123),
//   	rateLimit: jsii.Number(123),
//   }
//
type CfnUsagePlan_ThrottleSettingsProperty struct {
	// The API target request burst rate limit.
	//
	// This allows more requests through for a period of time than the target rate limit.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API target request rate limit.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

