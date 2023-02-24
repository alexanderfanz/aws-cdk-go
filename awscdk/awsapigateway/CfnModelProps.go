package awsapigateway


// Properties for defining a `CfnModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   cfnModelProps := &CfnModelProps{
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Schema: schema,
//   }
//
type CfnModelProps struct {
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The content-type for the model.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the model.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the model.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema for the model.
	//
	// For `application/json` models, this should be JSON schema draft 4 model. Do not include "\* /" characters in the description of any properties because such "\* /" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

