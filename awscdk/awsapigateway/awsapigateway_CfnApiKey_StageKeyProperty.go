package awsapigateway


// `StageKey` is a property of the [AWS::ApiGateway::ApiKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html) resource that specifies the stage to associate with the API key. This association allows only clients with the key to make requests to methods in that stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageKeyProperty := &stageKeyProperty{
//   	restApiId: jsii.String("restApiId"),
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnApiKey_StageKeyProperty struct {
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"optional" json:"restApiId" yaml:"restApiId"`
	// The stage name associated with the stage key.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

